package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/logger"
	"ymst_server/middleware"
	"ymst_server/services"
)

type HomeHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
	Svc  *services.HomeService
}

func NewHomeHandler(log *logger.Logger, game config.GameConfig, svc *services.HomeService) *HomeHandler {
	return &HomeHandler{
		Log:  log,
		Game: game,
		Svc:  svc,
	}
}

func (h *HomeHandler) RegisterRoutes(r gin.IRouter) {
	r.POST("/api/Home/CheckReceiveLoginBonus", h.CheckReceiveLoginBonus)
	r.GET("/api/Home/CheckReceiveLoginBonus", h.CheckReceiveLoginBonus)
	r.POST("/api/Home/CheckEexternalPayment", h.CheckEexternalPayment)
	r.GET("/api/Home/CheckEexternalPayment", h.CheckEexternalPayment)
}

func (h *HomeHandler) CheckReceiveLoginBonus(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)

	if h.Log != nil {
		h.Log.GetZap().Info("home check receive login bonus",
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
			zap.String("user_key", userKey),
		)
	}

	c.Header("X-Client-Version", h.Game.ClientVersion())
	c.Header("X-Assets-Version", h.Game.APIAssetsVersion())
	c.Header("X-FM", "0")

	payload, updates, err := h.Svc.CheckReceiveLoginBonus(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err := apiwire.WriteResult(c, payload, updates, nil, nil, apiwire.DefaultResponseOptions()); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *HomeHandler) CheckEexternalPayment(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)

	if h.Log != nil {
		h.Log.GetZap().Info("home check eexternal payment",
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
			zap.String("user_key", userKey),
		)
	}

	writeGameHeaders(c, h.Game)

	payload, updates, err := h.Svc.CheckExternalPayment(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err := apiwire.WriteResult(c, payload, updates, nil, nil, apiwire.DefaultResponseOptions()); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func firstNonEmpty(value string, fallback string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}
