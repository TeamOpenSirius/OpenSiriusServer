package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ymst_server/api"
	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/logger"
	"ymst_server/middleware"
	"ymst_server/repositories"
)

type UserDataHandler struct {
	Log   *logger.Logger
	Game  config.GameConfig
	Debug config.DebugConfig
	Users *repositories.UserRepository
}

func NewUserDataHandler(log *logger.Logger, game config.GameConfig, debug config.DebugConfig, users *repositories.UserRepository) *UserDataHandler {
	return &UserDataHandler{
		Log:   log,
		Game:  game,
		Debug: debug,
		Users: users,
	}
}

func (h *UserDataHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/data/user", h.GetUserData)
	r.POST("/api/data/user", h.GetUserData)
}

func (h *UserDataHandler) RegisterDebugRoutes(r gin.IRouter) {
	if !h.Debug.DumpAPIResponse {
		return
	}
	r.GET("/debug/user-state/stats", h.DebugStats)
	r.GET("/debug/user-state/objects", h.DebugObjects)
	r.POST("/debug/user-state/rebuild", h.DebugRebuild)
	r.GET("/debug/user-state/generated", h.DebugGenerated)
	r.GET("/debug/user-state/generated-msgpack", h.DebugGeneratedMsgpack)
}

func (h *UserDataHandler) GetUserData(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Header("Content-Type", api.MsgpackContentType)
	c.Header("X-Client-Version", h.Game.ClientVersion())
	c.Header("X-Assets-Version", h.Game.APIAssetsVersion())
	c.Header("X-FM", "0")

	if h.Log != nil {
		h.Log.GetZap().Debug("game user data request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("user_key", userKey),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	repo := h.Users.WithUserKey(userKey)
	payload, err := repo.BuildDataObjects(c.Request.Context())
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("build game user data failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	body, err := apiwire.EncodeResult(payload, nil, nil, nil, apiwire.ResponseOptions{
		CompressPayload:   true,
		CompressThreshold: 0,
		LZ4BlockSize:      64 * 1024,
	})
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("encode compressed user data failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Debug("game user data response",
			zap.Int("object_count", len(payload)),
			zap.Int("body_size", len(body)),
			zap.String("user_key", userKey),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	c.Data(http.StatusOK, api.MsgpackContentType, body)
}

func (h *UserDataHandler) DebugStats(c *gin.Context) {
	repo, ok := h.debugRepo(c)
	if !ok {
		return
	}
	stats, err := repo.Stats(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	unionCounts := map[string]int{}
	for unionID, count := range stats.UnionCounts {
		unionCounts[fmt.Sprintf("%d", unionID)] = count
	}
	c.JSON(http.StatusOK, gin.H{"userKey": repo.UserKey(), "seedMode": stats.SeedMode, "objectCount": stats.ObjectCount, "nilSlotCount": stats.NilSlotCount, "unionCounts": unionCounts})
}

func (h *UserDataHandler) DebugObjects(c *gin.Context) {
	repo, ok := h.debugRepo(c)
	if !ok {
		return
	}
	payload, err := repo.BuildDataObjects(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"userKey": repo.UserKey(), "objects": payload})
}

func (h *UserDataHandler) DebugRebuild(c *gin.Context) {
	repo, ok := h.debugRepo(c)
	if !ok {
		return
	}
	payload, err := repo.BuildGeneratedState(c.Request.Context(), "Player")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err := repo.ReplaceRawState(c.Request.Context(), payload, "generated"); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"userKey": repo.UserKey(), "objectCount": len(payload)})
}

func (h *UserDataHandler) debugRepo(c *gin.Context) (*repositories.UserRepository, bool) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, false
	}
	return h.Users.WithUserKey(userKey), true
}

func (h *UserDataHandler) DebugGenerated(c *gin.Context) {
	unionsParam := strings.TrimSpace(c.Query("unions"))
	profileParam := strings.TrimSpace(c.Query("profile"))
	if profileParam == "" {
		profileParam = h.Game.GeneratedUserProfile
	}
	repo := h.Users.WithProfileLevel(profileParam)
	if unionsParam != "" {
		ids := parseInt64CSV(unionsParam)
		repo = repo.WithIncludeUnions(ids)
	}
	payload, err := repo.BuildGeneratedState(c.Request.Context(), "DebugPlayer")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userKey":     repo.UserKey(),
		"profile":     profileParam,
		"unions":      unionsParam,
		"objectCount": len(payload),
		"payload":     payload,
	})
}

func parseInt64CSV(val string) []int64 {
	parts := strings.Split(val, ",")
	out := make([]int64, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		var n int64
		if _, err := fmt.Sscanf(part, "%d", &n); err == nil {
			out = append(out, n)
		}
	}
	return out
}

func (h *UserDataHandler) DebugGeneratedMsgpack(c *gin.Context) {
	unionsParam := strings.TrimSpace(c.Query("unions"))
	profileParam := strings.TrimSpace(c.Query("profile"))
	if profileParam == "" {
		profileParam = h.Game.GeneratedUserProfile
	}
	repo := h.Users.WithProfileLevel(profileParam)
	if unionsParam != "" {
		ids := parseInt64CSV(unionsParam)
		repo = repo.WithIncludeUnions(ids)
	}
	payload, err := repo.BuildGeneratedState(c.Request.Context(), "DebugPlayer")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	body, err := apiwire.EncodeResult(payload, nil, nil, nil, apiwire.ResponseOptions{
		CompressPayload:   true,
		CompressThreshold: 0,
		LZ4BlockSize:      64 * 1024,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Content-Type", api.MsgpackContentType)
	c.Header("X-FM", "0")
	c.Data(http.StatusOK, api.MsgpackContentType, body)
}

func (h *UserDataHandler) writeRawFixture(c *gin.Context) {
	body, err := os.ReadFile(h.Game.UserDataFixture)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("read raw user data fixture failed",
				zap.String("path", h.Game.UserDataFixture),
				zap.Error(err),
			)
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Debug("game user data raw fixture",
			zap.String("path", h.Game.UserDataFixture),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	c.Data(http.StatusOK, api.MsgpackContentType, body)
}
