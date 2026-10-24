package handlers

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ymst_server/api"
	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/logger"
	"ymst_server/middleware"
	"ymst_server/services"
)

type InboxHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
	Svc  *services.InboxService
}

func NewInboxHandler(log *logger.Logger, game config.GameConfig, svc *services.InboxService) *InboxHandler {
	return &InboxHandler{Log: log, Game: game, Svc: svc}
}

func (h *InboxHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/Inbox/CheckPackages", h.CheckPackagesAsync)
	r.POST("/api/Inbox/CheckPackages", h.CheckPackagesAsync)
	r.GET("/api/Inboxes/CheckPackages", h.CheckPackagesAsync)
	r.POST("/api/Inboxes/CheckPackages", h.CheckPackagesAsync)
	r.GET("/api/Inboxes/CheckPackagesAsync", h.CheckPackagesAsync)
	r.POST("/api/Inboxes/CheckPackagesAsync", h.CheckPackagesAsync)
	r.GET("/api/Inbox/GetPackages", h.GetPackages)
	r.POST("/api/Inbox/GetPackages", h.GetPackages)
	r.GET("/api/Inboxes/GetPackages", h.GetPackages)
	r.POST("/api/Inboxes/GetPackages", h.GetPackages)
	r.POST("/api/Inbox/ReceivePackage", h.ReceivePackage)
	r.POST("/api/Inboxes/ReceivePackage", h.ReceivePackage)
	r.POST("/api/Inbox/ReceiveAllPackages", h.ReceiveAllPackages)
	r.POST("/api/Inboxes/ReceiveAllPackages", h.ReceiveAllPackages)
}

func (h *InboxHandler) CheckPackagesAsync(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	h.logRequest(c, "inbox check packages")
	hasPackages, err := h.Svc.CheckPackages(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	if err := api.WriteBooleanResult(c, hasPackages, nil); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *InboxHandler) GetPackages(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	payload, err := h.Svc.GetPackages(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	if err := apiwire.WritePayload(c, payload); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *InboxHandler) ReceivePackage(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	arr, _ := apiwire.DecodeRequestArray(body)
	payload, updates, err := h.Svc.ReceivePackage(c.Request.Context(), userKey, apiwire.StringAt(arr, 0))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	if err := apiwire.WriteResult(c, payload, updates, nil, nil, apiwire.DefaultResponseOptions()); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *InboxHandler) ReceiveAllPackages(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	payload, updates, err := h.Svc.ReceiveAllPackages(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	if err := apiwire.WriteResult(c, payload, updates, nil, nil, apiwire.DefaultResponseOptions()); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *InboxHandler) logRequest(c *gin.Context, message string) {
	body, _ := io.ReadAll(c.Request.Body)
	if h.Log != nil {
		h.Log.GetZap().Info(message,
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}
}

type FriendHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
	Svc  *services.FriendService
}

func NewFriendHandler(log *logger.Logger, game config.GameConfig, svc *services.FriendService) *FriendHandler {
	return &FriendHandler{Log: log, Game: game, Svc: svc}
}

func (h *FriendHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/Friends/GetReceivedRequest", h.ReceivedRequest)
	r.POST("/api/Friends/GetReceivedRequest", h.ReceivedRequest)
	r.GET("/api/Friends/ReceivedRequest", h.ReceivedRequest)
	r.POST("/api/Friends/ReceivedRequest", h.ReceivedRequest)
	r.GET("/api/ReceivedRequest", h.ReceivedRequest)
	r.POST("/api/ReceivedRequest", h.ReceivedRequest)
	r.GET("/api/Friends/GetFriends", h.GetFriends)
	r.POST("/api/Friends/GetFriends", h.GetFriends)
	r.POST("/api/Friends/SendRequest", h.SendRequest)
	r.POST("/api/Friends/AcceptRequest", h.AcceptRequest)
	r.POST("/api/Friends/RejectRequest", h.RejectRequest)
	r.POST("/api/Friends/DeleteFriend", h.DeleteFriend)
}

func (h *FriendHandler) ReceivedRequest(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	logGameRequest(h.Log, c, "friend received request")
	results, count, err := h.Svc.GetReceivedRequests(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	if err := apiwire.WritePayload(c, apiwire.FriendListResult(results, count)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *FriendHandler) GetFriends(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	results, count, err := h.Svc.GetFriends(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	if err := apiwire.WritePayload(c, apiwire.FriendListResult(results, count)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
func (h *FriendHandler) SendRequest(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	arr, _ := apiwire.DecodeRequestArray(body)
	payload, err := h.Svc.SendRequest(c.Request.Context(), userKey, apiwire.StringAt(arr, 0), apiwire.StringAt(arr, 1))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	_ = apiwire.WritePayload(c, payload)
}
func (h *FriendHandler) AcceptRequest(c *gin.Context) { h.friendRequestAction(c, h.Svc.AcceptRequest) }
func (h *FriendHandler) RejectRequest(c *gin.Context) { h.friendRequestAction(c, h.Svc.RejectRequest) }
func (h *FriendHandler) DeleteFriend(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	arr, _ := apiwire.DecodeRequestArray(body)
	payload, err := h.Svc.DeleteFriend(c.Request.Context(), userKey, apiwire.StringAt(arr, 0))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	_ = apiwire.WritePayload(c, payload)
}
func (h *FriendHandler) friendRequestAction(c *gin.Context, fn func(context.Context, string, string) ([]any, error)) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	arr, _ := apiwire.DecodeRequestArray(body)
	payload, err := fn(c.Request.Context(), userKey, apiwire.StringAt(arr, 0))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	_ = apiwire.WritePayload(c, payload)
}

type CircleHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
	Svc  *services.CircleService
}

func NewCircleHandler(log *logger.Logger, game config.GameConfig, svc *services.CircleService) *CircleHandler {
	return &CircleHandler{Log: log, Game: game, Svc: svc}
}

func (h *CircleHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/Circles/GetInvitedCircles", h.Invited)
	r.POST("/api/Circles/GetInvitedCircles", h.Invited)
	r.GET("/api/Circles/Invited", h.Invited)
	r.POST("/api/Circles/Invited", h.Invited)
	r.GET("/api/Circles/GetSupportAndTheaterLevelInformation", h.GetSupportAndTheaterLevelInformation)
	r.POST("/api/Circles/GetSupportAndTheaterLevelInformation", h.GetSupportAndTheaterLevelInformation)
	r.GET("/api/Circles/GetMyCircle", h.GetMyCircle)
	r.POST("/api/Circles/GetMyCircle", h.GetMyCircle)
	r.POST("/api/Circles/Create", h.Create)
	r.POST("/api/Circles/Join", h.AcceptInvite)
	r.POST("/api/Circles/AcceptInvite", h.AcceptInvite)
	r.POST("/api/Circles/RejectInvite", h.RejectInvite)
	r.POST("/api/Circles/Leave", h.Leave)
}

func (h *CircleHandler) Invited(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	logGameRequest(h.Log, c, "circle invited")
	payload, err := h.Svc.GetInvitedCircles(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	if err := api.WriteApiMsgpackResult(c, payload); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *CircleHandler) GetMyCircle(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	payload, err := h.Svc.GetMyCircle(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	_ = apiwire.WritePayload(c, payload)
}
func (h *CircleHandler) Create(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	arr, _ := apiwire.DecodeRequestArray(body)
	payload, err := h.Svc.CreateCircle(c.Request.Context(), userKey, apiwire.StringAt(arr, 0), apiwire.StringAt(arr, 1))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	_ = apiwire.WritePayload(c, payload)
}
func (h *CircleHandler) AcceptInvite(c *gin.Context) { h.circleInviteAction(c, h.Svc.AcceptInvite) }
func (h *CircleHandler) RejectInvite(c *gin.Context) { h.circleInviteAction(c, h.Svc.RejectInvite) }
func (h *CircleHandler) Leave(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	payload, err := h.Svc.LeaveCircle(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	_ = apiwire.WritePayload(c, payload)
}
func (h *CircleHandler) circleInviteAction(c *gin.Context, fn func(context.Context, string, string) ([]any, error)) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	arr, _ := apiwire.DecodeRequestArray(body)
	payload, err := fn(c.Request.Context(), userKey, apiwire.StringAt(arr, 0))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	writeGameHeaders(c, h.Game)
	_ = apiwire.WritePayload(c, payload)
}

func (h *CircleHandler) GetSupportAndTheaterLevelInformation(c *gin.Context) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	logGameRequest(h.Log, c, "circle support and theater level information")
	writeGameHeaders(c, h.Game)

	payload, err := h.Svc.GetSupportAndTheaterLevelInformation(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err := api.WriteApiMsgpackResult(c, payload); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func logGameRequest(log *logger.Logger, c *gin.Context, message string) {
	body, _ := io.ReadAll(c.Request.Body)
	if log != nil {
		log.GetZap().Info(message,
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}
}

func writeGameHeaders(c *gin.Context, game config.GameConfig) {
	c.Header("X-Client-Version", game.ClientVersion())
	c.Header("X-Assets-Version", game.APIAssetsVersion())
	c.Header("X-FM", "0")
}
