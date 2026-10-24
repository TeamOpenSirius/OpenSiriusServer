package router

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ymst_server/core/auth"
	"ymst_server/core/config"
	"ymst_server/core/database"
	"ymst_server/core/logger"
	"ymst_server/handlers"
	"ymst_server/masterdata"
	"ymst_server/middleware"
	"ymst_server/repositories"
	"ymst_server/services"
	"ymst_server/userdata"
)

type SecurityOptions struct {
	Enabled bool
	Guard   middleware.GuardConfig
	Login   services.LoginGuardConfig
}

// New builds the HTTP router.
func New(log *logger.Logger, authSvc *auth.Service, db *database.Client, game config.GameConfig, debug config.DebugConfig, jwtSecret string, sec SecurityOptions) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		c.Header("Date", "Sat, 16 May 2026 12:00:00 GMT")
		c.Next()
	})
	if sec.Enabled {
		securityGuard := middleware.NewSecurityGuard(log, sec.Guard)
		r.Use(securityGuard.Handler())
	}
	r.Use(middleware.ZapLogger(log))
	if syncedGame, err := masterdata.SyncFromBackend(context.Background(), game, log.GetZap()); err != nil {
		log.GetZap().Warn("sync master data from backend failed", zap.Error(err))
	} else {
		game = syncedGame
	}

	userRepo := repositories.NewUserRepository(db.DB(), game.UserDataSeedJSON).
		WithSeedMode(game.UserDataSeedMode).
		WithProfileLevel(game.GeneratedUserProfile).
		WithIncludeUnions(game.GeneratedIncludeUnions).
		WithExcludeUnions(game.GeneratedExcludeUnions).
		WithSeedCharacterIDs(game.SeedCharacterBaseMasterIDs, game.SeedCharacterMasterIDs, nil, nil, nil).
		WithGeneratedOptions(userdataGeneratedOptions(game))

	if templates, err := userdata.LoadTemplateRegistry(game.UserDataSeedJSON); err != nil {
		log.GetZap().Warn("load user data template registry failed", zap.Error(err))
	} else if templates.TemplateCount() > 0 {
		userRepo = userRepo.WithTemplates(templates)
		log.GetZap().Info("user data template registry loaded", zap.Int("template_count", templates.TemplateCount()))
	}
	if !strings.EqualFold(strings.TrimSpace(game.UserDataSeedMode), userdata.SeedModeRawFixture) {
		if err := userRepo.EnsureSeeded(); err != nil {
			log.GetZap().Error("initialize game user repository failed", zap.Error(err))
			panic(err)
		} else if userSeedStats, err := userRepo.Stats(); err != nil {
			log.GetZap().Warn("load game user data stats failed", zap.Error(err))
		} else {
			log.GetZap().Info("game user repository initialized",
				zap.Int("object_count", userSeedStats.ObjectCount),
				zap.Int("game_hint_count", userSeedStats.GameHintCount),
				zap.Int("character_count", userSeedStats.CharacterCount),
				zap.Int("party_count", userSeedStats.PartyCount),
				zap.Int("party_slot_count", userSeedStats.PartySlotCount),
				zap.Int("tutorial_status", userSeedStats.TutorialStatus),
			)
		}
	}
	accountRepo := repositories.NewAccountRepository(db.DB())
	sessionRepo := repositories.NewSessionRepository(db.DB())
	accountSvc := services.NewAccountService(accountRepo, sessionRepo, userRepo)
	accountExtraSvc := services.NewAccountExtraService(accountRepo, sessionRepo, repositories.NewConfirmationCodeRepository(db.DB()))
	loginSvc := services.NewLoginService(accountRepo, userRepo)
	homeSvc := services.NewHomeService()
	playerSvc := services.NewPlayerService(userRepo)
	shopSvc := services.NewShopService(userRepo)
	notificationSvc := services.NewNotificationDBService(repositories.NewNotificationRepository(db.DB()), repositories.NewNotificationReadRepository(db.DB()))
	inboxSvc := services.NewInboxService(repositories.NewInboxRepository(db.DB()))
	friendSvc := services.NewFriendService(repositories.NewFriendRepository(db.DB()))
	circleSvc := services.NewCircleService(repositories.NewCircleRepository(db.DB()), repositories.NewSupportCompanyRepository(db.DB()))
	masterRepo := masterdata.Open(filepath.Join(game.MasterDataRoot, game.MasterDataURI), game.MasterDataRoot)
	episodeSvc := services.NewEpisodeService(masterRepo, db.DB(), game.BaseURL)
	r.Use(middleware.AuthMiddlewareWithPublicPaths(sessionRepo, gamePublicPaths()))

	healthHandler := &handlers.HealthHandler{DB: db}
	environmentHandler := handlers.NewEnvironmentHandlerWithLogger(game, log)
	environmentHandler.RegisterRoutes(r)
	accountHandler := handlers.NewAccountHandler(log, game, db.DB(), jwtSecret, accountSvc, accountExtraSvc)
	accountHandler.RegisterRoutes(r)
	gameLoginHandler := handlers.NewGameLoginHandler(log, game, loginSvc)
	gameLoginHandler.RegisterRoutes(r)
	masterDataHandler := handlers.NewMasterDataHandler(log, game)
	masterDataHandler.RegisterRoutes(r)
	episodeHandler := handlers.NewEpisodeHandler(log, game, episodeSvc)
	episodeHandler.RegisterRoutes(r)
	liveHandler := handlers.NewLiveHandler(log, game, db.DB())
	liveHandler.RegisterRoutes(r)
	leagueHandler := handlers.NewLeagueHandler(log, game)
	leagueHandler.RegisterRoutes(r)
	tripleCastHandler := handlers.NewTripleCastHandler(log, game)
	tripleCastHandler.RegisterRoutes(r)
	userDataHandler := handlers.NewUserDataHandler(log, game, debug, userRepo)
	userDataHandler.RegisterRoutes(r)
	userDataHandler.RegisterDebugRoutes(r)
	homeHandler := handlers.NewHomeHandler(log, game, homeSvc)
	homeHandler.RegisterRoutes(r)
	playerHandler := handlers.NewPlayerHandler(log, game, userRepo, playerSvc)
	playerHandler.RegisterRoutes(r)
	inboxHandler := handlers.NewInboxHandler(log, game, inboxSvc)
	inboxHandler.RegisterRoutes(r)
	friendHandler := handlers.NewFriendHandler(log, game, friendSvc)
	friendHandler.RegisterRoutes(r)
	friendInvitationHandler := handlers.NewFriendInvitationHandler(log, game)
	friendInvitationHandler.RegisterRoutes(r)
	circleHandler := handlers.NewCircleHandler(log, game, circleSvc)
	circleHandler.RegisterRoutes(r)
	shopHandler := handlers.NewShopHandler(log, game, userRepo, shopSvc)
	shopHandler.RegisterRoutes(r)
	notificationHandler := handlers.NewNotificationHandler(log, game, notificationSvc)
	notificationHandler.RegisterRoutes(r)

	v1 := r.Group("/api/v1")

	// Public routes.
	v1.GET("/health", healthHandler.Health)

	v1Protected := v1.Group("/")
	v1Protected.Use(middleware.Authenticate(authSvc))

	adminUsers := v1Protected.Group("/users")
	adminUsers.Use(middleware.Authorize(auth.RoleAdmin))

	return r
}

func userdataGeneratedOptions(game config.GameConfig) userdata.GeneratedUserOptions {
	return userdata.GeneratedUserOptions{
		PlayerRank:     game.GeneratedUser.PlayerRank,
		CurrentStamina: game.GeneratedUser.CurrentStamina,
		Coin:           game.GeneratedUser.Coin,
		FreeJewel:      game.GeneratedUser.FreeJewel,
		PaidJewel:      game.GeneratedUser.PaidJewel,
		TutorialStatus: game.GeneratedUser.TutorialStatus,
	}
}

func gamePublicPaths() map[string]bool {
	return map[string]bool{
		"/api/Environment":                 true,
		"/api/Environment/Ping":            true,
		"/api/Account/Register":            true,
		"/api/Account/Authenticate":        true,
		"/api/Account/GetConfirmationCode": true,
		"/api/data/master":                 true,
	}
}
