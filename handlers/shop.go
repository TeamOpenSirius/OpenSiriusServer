package handlers

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ymst_server/api"
	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/logger"
	"ymst_server/middleware"
	"ymst_server/repositories"
	"ymst_server/services"
)

type ShopHandler struct {
	Log   *logger.Logger
	Game  config.GameConfig
	Users *repositories.UserRepository
	Svc   *services.ShopService
}

func NewShopHandler(log *logger.Logger, game config.GameConfig, users *repositories.UserRepository, svc *services.ShopService) *ShopHandler {
	return &ShopHandler{Log: log, Game: game, Users: users, Svc: svc}
}

func (h *ShopHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/Shops/ViewPage", h.ViewPage)
	r.POST("/api/Shops/ViewPage", h.ViewPage)
}

func (h *ShopHandler) ViewPage(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	values, err := apiwire.DecodeRequestArray(body)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Warn("decode shop view page request failed", zap.Error(err), zap.Int("body_size", len(body)))
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	viewPageType := apiwire.Int64At(values, 0)
	exchangeShopMasterID := apiwire.NullableInt64At(values, 1)
	payload, updates, err := h.Svc.ViewPage(c.Request.Context(), userKey, viewPageType, exchangeShopMasterID)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("mark shop viewed failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("shop view page",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("body_size", len(body)),
			zap.Int64("view_page_type", viewPageType),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
			zap.Int64("unix", time.Now().Unix()),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := api.WriteApiMsgpackResultFull(c, payload, updates, nil, nil); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
