package core

import (
	"fmt"

	"go.uber.org/zap"

	"ymst_server/core/auth"
	"ymst_server/middleware"
	"ymst_server/router"
	"ymst_server/services"
)

// Run starts the HTTP server component.
func Run() error {
	if err := InitCommon(); err != nil {
		return err
	}
	defer ShutdownCommon()

	common := Common
	if common == nil {
		return fmt.Errorf("common is not initialized")
	}

	authSvc := auth.NewService(
		common.Config.JWTSecret,
		common.Config.JWTIssuer,
		common.Config.AccessTTL,
		common.Config.RefreshTTL,
	)
	authSvc.BindPersistence(auth.NewPersistence(common.Database.DB()))
	authSvc.BindUserRepository(auth.NewUserRepository(common.Database.DB()))

	r := router.New(common.Logger, authSvc, common.Database, common.Config.Game, common.Config.Debug, common.Config.JWTSecret, router.SecurityOptions{
		Enabled: common.Security.IsEnabled(),
		Guard: middleware.GuardConfig{
			MinRequestInterval: common.Security.MinRequestInterval,
			WindowSize:         common.Security.WindowSize,
			MaxRequestsPerIP:   common.Security.MaxRequestsPerIP,
			DuplicateTTL:       common.Security.DuplicateTTL,
			MaxBodyBytes:       common.Security.MaxBodyBytes,
		},
		Login: services.LoginGuardConfig{
			MinInterval:   common.Security.LoginMinInterval,
			WindowSize:    common.Security.LoginWindowSize,
			MaxPerIP:      common.Security.MaxLoginPerIP,
			MaxFailures:   common.Security.MaxLoginFailures,
			MaxUsersPerIP: common.Security.MaxLoginUsersPerIP,
		},
	})

	common.Logger.GetZap().Info("server starting",
		zap.String("addr", common.Config.Addr),
		zap.String("env", common.Config.Env),
		zap.String("gin_mode", common.GinMode),
		zap.String("db_mode", common.Database.Mode()),
		zap.String("log_path", common.Config.LogPath),
		zap.String("jwt_issuer", common.Config.JWTIssuer),
		zap.Bool("security_enabled", common.Security.IsEnabled()),
	)

	if err := r.Run(common.Config.Addr); err != nil {
		common.Logger.GetZap().Error("server stopped", zap.Error(err))
		return err
	}
	return nil
}
