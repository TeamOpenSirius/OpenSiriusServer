package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"

	"ymst_server/api"
	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/logger"
	"ymst_server/dto"
	"ymst_server/middleware"
	"ymst_server/services"
)

type NotificationHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
	Svc  *services.NotificationService
}

func NewNotificationHandler(log *logger.Logger, game config.GameConfig, svc *services.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		Log:  log,
		Game: game,
		Svc:  svc,
	}
}

func (h *NotificationHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/Home/GetNotifications", h.GetNotifications)
	r.POST("/api/Home/GetNotifications", h.GetNotifications)
	r.GET("/api/Home/GetNotificationsAsync", h.GetNotifications)
	r.POST("/api/Home/GetNotificationsAsync", h.GetNotifications)
	r.GET("/api/Home/GetNotificationsAsync/:id", h.GetNotificationContent)
	r.POST("/api/Home/GetNotificationsAsync/:id", h.GetNotificationContent)

	r.GET("/api/Home/GetNotificationsAnonymous", h.GetNotifications)
	r.POST("/api/Home/GetNotificationsAnonymous", h.GetNotifications)
	r.GET("/api/Home/GetNotificationsAnonymousAsync", h.GetNotifications)
	r.POST("/api/Home/GetNotificationsAnonymousAsync", h.GetNotifications)
	r.GET("/api/Home/GetNotificationsAnonymousAsync/:id", h.GetNotificationContent)
	r.POST("/api/Home/GetNotificationsAnonymousAsync/:id", h.GetNotificationContent)

	r.GET("/api/Home/GetNotificationContent", h.GetNotificationContent)
	r.POST("/api/Home/GetNotificationContent", h.GetNotificationContent)
	r.GET("/api/Home/GetNotificationContentAsync", h.GetNotificationContent)
	r.POST("/api/Home/GetNotificationContentAsync", h.GetNotificationContent)
	r.GET("/api/Home/GetNotificationContentAnonymous", h.GetNotificationContent)
	r.POST("/api/Home/GetNotificationContentAnonymous", h.GetNotificationContent)
	r.GET("/api/Home/GetNotificationContentAnonymousAsync", h.GetNotificationContent)
	r.POST("/api/Home/GetNotificationContentAnonymousAsync", h.GetNotificationContent)

	r.GET("/api/Home/UpdateNotificationReadTime", h.UpdateNotificationReadTime)
	r.POST("/api/Home/UpdateNotificationReadTime", h.UpdateNotificationReadTime)
	r.GET("/api/Home/UpdateNotificationReadTimeAsync", h.UpdateNotificationReadTime)
	r.POST("/api/Home/UpdateNotificationReadTimeAsync", h.UpdateNotificationReadTime)
	r.GET("/api/Home/GetMultiLiveRestrictionNotification", h.GetMultiLiveRestrictionNotification)
	r.POST("/api/Home/GetMultiLiveRestrictionNotification", h.GetMultiLiveRestrictionNotification)
	r.GET("/api/Home/GetMultiLiveRestrictionNotificationAsync", h.GetMultiLiveRestrictionNotification)
	r.POST("/api/Home/GetMultiLiveRestrictionNotificationAsync", h.GetMultiLiveRestrictionNotification)
}

func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	payload, err := h.Svc.GetNotifications(c.Request.Context(), userKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("home get notifications",
			zap.Int("body_size", len(body)),
			zap.Int("count", len(payload)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := api.WriteApiMsgpackResult(c, payload); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *NotificationHandler) GetNotificationContent(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	id := parseNotificationID(c, body)
	payload, err := h.Svc.GetNotificationContent(c.Request.Context(), userKey, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("home get notification content",
			zap.Int64("request_id", id),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := api.WriteApiMsgpackResult(c, payload); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *NotificationHandler) UpdateNotificationReadTime(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	arr, _ := apiwire.DecodeRequestArray(body)
	tabCategory := apiwire.IntAt(arr, 0)
	readAt := apiwire.TimeAt(arr, 1)
	if tabCategory <= 0 {
		tabCategory = 1
	}
	if readAt.IsZero() {
		readAt = time.Now()
	}
	if err := h.Svc.UpdateReadTime(c.Request.Context(), userKey, tabCategory, readAt); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("home update notification read time",
			zap.Int("tab_category", tabCategory),
			zap.Time("read_at", readAt),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := api.WriteBooleanResult(c, true, nil); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *NotificationHandler) GetMultiLiveRestrictionNotification(c *gin.Context) {
	logGameRequest(h.Log, c, "home get multi live restriction notification")
	writeGameHeaders(c, h.Game)
	if err := api.WriteBooleanResult(c, false, nil); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func parseNotificationID(c *gin.Context, body []byte) int64 {
	if raw := strings.TrimSpace(c.Param("id")); raw != "" {
		if id, err := strconv.ParseInt(raw, 10, 64); err == nil {
			return id
		}
	}

	for _, key := range []string{"mNotificationId", "notificationId", "id"} {
		if raw := strings.TrimSpace(c.Query(key)); raw != "" {
			if id, err := strconv.ParseInt(raw, 10, 64); err == nil {
				return id
			}
		}
	}

	if id, ok := decodeMsgpackScalarInt64(body); ok {
		return id
	}

	var arr []int64
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&arr); err == nil && len(arr) > 0 {
		return arr[0]
	}

	var arrAny []any
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&arrAny); err == nil && len(arrAny) > 0 {
		return anyToInt64(arrAny[0])
	}

	return 0
}

func decodeMsgpackScalarInt64(body []byte) (int64, bool) {
	if len(body) == 0 {
		return 0, false
	}

	var value int64
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&value); err == nil && value > 0 {
		return value, true
	}

	var valueAny any
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&valueAny); err == nil {
		v := anyToInt64(valueAny)
		if v > 0 {
			return v, true
		}
	}

	return 0, false
}

type notificationBodyContentSeed struct {
	Element int    `json:"Element"`
	Data    string `json:"Data"`
}

func seedNotificationBodyJSON() string {
	raw, err := json.Marshal([]notificationBodyContentSeed{
		{Element: 2, Data: "お知らせ"},
		{Element: 4, Data: "これはテスト用のお知らせ本文です。"},
		{Element: 4, Data: "サーバーから配信された詳細内容を表示しています。"},
	})
	if err != nil {
		return `[{"Element":4,"Data":"お知らせ内容です。"}]`
	}
	return string(raw)
}

func normalizeNotificationBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return seedNotificationBodyJSON()
	}
	if strings.HasPrefix(body, "[") {
		return body
	}

	raw, err := json.Marshal([]notificationBodyContentSeed{
		{Element: 4, Data: body},
	})
	if err != nil {
		return `[{"Element":4,"Data":"お知らせ内容です。"}]`
	}
	return string(raw)
}

func notificationResultPayload(n services.Notification) []any {
	result := dto.NotificationResult{
		Id:                      n.ID,
		BannerPath:              n.BannerPath,
		Title:                   n.Title,
		PostingAt:               n.PostingAt,
		LastUpdatedAt:           n.LastUpdatedAt,
		NotificationTabCategory: dto.NotificationTabCategory(n.TabCategory),
		NotificationCategory:    dto.NotificationCategory(n.Category),
		IsConfirmation:          n.IsConfirmation,
		Order:                   int(n.Order),
	}

	return []any{
		result.Id,
		result.BannerPath,
		result.Title,
		result.PostingAt,
		result.LastUpdatedAt,
		result.NotificationTabCategory,
		result.NotificationCategory,
		result.IsConfirmation,
		result.Order,
	}
}

func notificationContentResultPayload(n services.Notification) []any {
	result := dto.NotificationContentResult{
		Id:                      n.ID,
		Body:                    normalizeNotificationBody(n.Body),
		BannerPath:              n.BannerPath,
		Title:                   n.Title,
		PostingAt:               n.PostingAt,
		LastUpdatedAt:           n.LastUpdatedAt,
		NotificationTabCategory: dto.NotificationTabCategory(n.TabCategory),
		NotificationCategory:    dto.NotificationCategory(n.Category),
	}

	return []any{
		result.Id,
		result.Body,
		result.Title,
		result.PostingAt,
		result.LastUpdatedAt,
		result.NotificationTabCategory,
		result.NotificationCategory,
		result.BannerPath,
	}
}

func decodeReadNotificationPayload(body []byte) (int64, time.Time) {
	tabCategory := int64(1)
	readAt := time.Now()

	if id, ok := decodeMsgpackScalarInt64(body); ok {
		return id, readAt
	}

	var arrAny []any
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&arrAny); err != nil {
		return tabCategory, readAt
	}

	if len(arrAny) > 0 {
		if value := anyToInt64(arrAny[0]); value > 0 {
			tabCategory = value
		}
	}
	if len(arrAny) > 1 {
		if t, ok := arrAny[1].(time.Time); ok {
			readAt = t
		}
	}

	return tabCategory, readAt
}

func anyToInt64(value any) int64 {
	switch v := value.(type) {
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case int:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		if v <= math.MaxInt64 {
			return int64(v)
		}
		return 0
	case uint:
		return int64(v)
	default:
		return 0
	}
}
