package handlers

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"

	"ymst_server/api"
	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/logger"
	"ymst_server/middleware"
	"ymst_server/repositories"
	"ymst_server/services"
)

type PlayerHandler struct {
	Log   *logger.Logger
	Game  config.GameConfig
	Users *repositories.UserRepository
	Svc   *services.PlayerService
}

func NewPlayerHandler(log *logger.Logger, game config.GameConfig, users *repositories.UserRepository, svc *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		Log:   log,
		Game:  game,
		Users: users,
		Svc:   svc,
	}
}

func (h *PlayerHandler) RegisterRoutes(r gin.IRouter) {
	r.POST("/api/Player/UpdateGameHintRead", h.UpdateGameHintRead)
	r.GET("/api/Player/UpdateGameHintRead", h.UpdateGameHintRead)
	r.POST("/api/Player/UpdateTutorialStatus", h.UpdateTutorialStatus)
	r.GET("/api/Player/UpdateTutorialStatus", h.UpdateTutorialStatus)
	r.POST("/api/Player/UpdateTutorial", h.UpdateTutorialStatus)
	r.GET("/api/Player/UpdateTutorial", h.UpdateTutorialStatus)
	r.POST("/api/Player/UpdateSplashLastDisplayTime", h.UpdateSplashLastDisplayTime)
	r.GET("/api/Player/UpdateSplashLastDisplayTime", h.UpdateSplashLastDisplayTime)
}

func (h *PlayerHandler) UpdateGameHintRead(c *gin.Context) {
	userKey, ok := h.userKeyForRequest(c)
	if !ok {
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	values, err := apiwire.DecodeRequestArray(body)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Warn("decode game hint read request failed", zap.Error(err), zap.Int("body_size", len(body)))
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	categories := int64SliceFromAny(apiwire.ValueAt(values, 0))
	decoded := len(values) > 0
	updates, err := h.Svc.UpdateGameHintRead(c.Request.Context(), userKey, categories)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("update game hint read failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("player update game hint read",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
			zap.Any("categories", categories),
			zap.Bool("decoded", decoded),
		)
	}

	c.Header("X-Client-Version", h.Game.ClientVersion())
	c.Header("X-Assets-Version", h.Game.APIAssetsVersion())
	c.Header("X-FM", "0")

	if err := api.WriteBooleanResult(c, true, updates); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func decodeGameHintCategories(body []byte) ([]int64, bool) {
	var int64Payload [][]int64
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&int64Payload); err == nil {
		if len(int64Payload) > 0 && len(int64Payload[0]) > 0 {
			return int64Payload[0], true
		}
	}

	var intPayload [][]int
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&intPayload); err == nil {
		if len(intPayload) > 0 && len(intPayload[0]) > 0 {
			out := make([]int64, 0, len(intPayload[0]))
			for _, value := range intPayload[0] {
				out = append(out, int64(value))
			}
			return out, true
		}
	}

	return nil, false
}

func (h *PlayerHandler) UpdateTutorialStatus(c *gin.Context) {
	userKey, ok := h.userKeyForRequest(c)
	if !ok {
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	values, err := apiwire.DecodeRequestArray(body)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Warn("decode tutorial status request failed", zap.Error(err), zap.Int("body_size", len(body)))
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	status := apiwire.Int64At(values, 0)
	decoded := len(values) > 0
	updates, err := h.Svc.UpdateTutorial(c.Request.Context(), userKey, status)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("update tutorial status failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("player update tutorial status",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("body_size", len(body)),
			zap.Int64("status", status),
			zap.Bool("decoded", decoded),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := api.WriteBooleanResult(c, true, updates); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *PlayerHandler) UpdateSplashLastDisplayTime(c *gin.Context) {
	userKey, ok := h.userKeyForRequest(c)
	if !ok {
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	updates, err := h.Svc.UpdateSplashLastDisplayTime(c.Request.Context(), userKey, time.Now())
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("update splash last display time failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("player update splash last display time",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := api.WriteBooleanResult(c, true, updates); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *PlayerHandler) userKeyForRequest(c *gin.Context) (string, bool) {
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return "", false
	}
	return userKey, true
}

func decodeFirstInt64(body []byte) (int64, bool) {
	var nested [][]int64
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&nested); err == nil {
		if len(nested) > 0 && len(nested[0]) > 0 {
			return nested[0][0], true
		}
	}

	var values []int64
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&values); err == nil {
		if len(values) > 0 {
			return values[0], true
		}
	}

	var value int64
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&value); err == nil {
		return value, true
	}

	return 0, false
}

func int64SliceFromAny(value any) []int64 {
	arr, ok := value.([]any)
	if !ok {
		return nil
	}
	out := make([]int64, 0, len(arr))
	for _, item := range arr {
		out = append(out, apiwire.Int64Value(item))
	}
	return out
}
