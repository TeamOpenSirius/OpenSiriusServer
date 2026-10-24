package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ymst_server/api"
	"ymst_server/core/config"
	"ymst_server/core/logger"
)

type EnvironmentHandler struct {
	Game config.GameConfig
	Log  *logger.Logger
}

func NewEnvironmentHandler(game config.GameConfig) *EnvironmentHandler {
	return &EnvironmentHandler{Game: game}
}

func NewEnvironmentHandlerWithLogger(game config.GameConfig, log *logger.Logger) *EnvironmentHandler {
	return &EnvironmentHandler{Game: game, Log: log}
}

func (h *EnvironmentHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/Environment", h.GetEnvironment)
	r.POST("/api/Environment", h.GetEnvironment)
	r.GET("/api/Environment/Ping", h.Ping)
}

func (h *EnvironmentHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (h *EnvironmentHandler) GetEnvironment(c *gin.Context) {
	appVersion := strings.TrimSpace(c.Query("applicationVersion"))
	if appVersion == "" {
		appVersion = strings.TrimSpace(h.Game.ApplicationVersion)
	}
	if appVersion == "" {
		appVersion = "2.30.1"
	}

	assetVersion := strings.TrimSpace(h.Game.AssetVersion)
	if assetVersion == "" {
		assetVersion = "1.91.0"
	}

	baseURL := strings.TrimRight(strings.TrimSpace(h.Game.BaseURL), "/")
	if baseURL == "" {
		baseURL = "http://127.0.0.1:8080"
	}

	realtimeURL := strings.TrimRight(strings.TrimSpace(h.Game.RealtimeURL), "/")
	if realtimeURL == "" {
		realtimeURL = baseURL
	}

	masterDataURL := strings.TrimRight(strings.TrimSpace(h.Game.MasterDataURL), "/")
	if masterDataURL == "" {
		masterDataURL = baseURL + "/master-data/production"
	}
	staticContentURL := baseURL
	assetURL := strings.TrimRight(strings.TrimSpace(h.Game.AssetURL), "/")
	if assetURL == "" {
		assetURL = baseURL + "/production"
	}

	c.Header("X-Client-Version", appVersion)
	c.Header("X-Assets-Version", assetVersion)
	c.Header("X-FM", "0")

	payload := []any{
		appVersion,
		assetVersion,
		baseURL,
		baseURL,
		baseURL,
		false,
		masterDataURL,
		staticContentURL,
		assetURL,
		false,
		baseURL,
		realtimeURL,
		baseURL + "/game/web-shop",
	}

	if h.Log != nil {
		h.Log.GetZap().Info("environment",
			zap.String("app_version", appVersion),
			zap.String("asset_version", assetVersion),
			zap.String("base_url", baseURL),
			zap.String("master_data_url", masterDataURL),
			zap.String("static_content_url", staticContentURL),
			zap.String("asset_url", assetURL),
		)
	}

	if err := api.WriteApiMsgpackResult(c, payload); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
