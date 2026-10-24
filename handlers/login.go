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

type GameLoginHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
	Svc  *services.LoginService
}

func NewGameLoginHandler(log *logger.Logger, game config.GameConfig, svc *services.LoginService) *GameLoginHandler {
	return &GameLoginHandler{
		Log:  log,
		Game: game,
		Svc:  svc,
	}
}

func (h *GameLoginHandler) RegisterRoutes(r gin.IRouter) {
	r.POST("/api/Login", h.Login)
}

func (h *GameLoginHandler) Login(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	auth := c.GetHeader("Authorization")

	if h.Log != nil {
		h.Log.GetZap().Info("game login",
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(auth) != ""),
			zap.String("user_key", userKey),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}

	c.Header("X-Client-Version", h.Game.ClientVersion())
	c.Header("X-Assets-Version", h.Game.APIAssetsVersion())

	arr, _ := apiwire.DecodeRequestArray(body)
	pushToken := apiwire.StringAt(arr, 0)
	payload, updates, err := h.Svc.Login(c.Request.Context(), userKey, pushToken)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := apiwire.WriteResult(c, payload, updates, nil, nil, apiwire.DefaultResponseOptions()); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
