package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ymst_server/api"
	"ymst_server/core/config"
	"ymst_server/core/logger"
)

type FriendInvitationHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
}

func NewFriendInvitationHandler(log *logger.Logger, game config.GameConfig) *FriendInvitationHandler {
	return &FriendInvitationHandler{Log: log, Game: game}
}

func (h *FriendInvitationHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/FriendInvitation/Update", h.Update)
	r.POST("/api/FriendInvitation/Update", h.Update)
}

func (h *FriendInvitationHandler) Update(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	if h.Log != nil {
		h.Log.GetZap().Info("friend invitation update",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := api.WriteBooleanResult(c, true, nil); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
