package handlers

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"

	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/logger"
	"ymst_server/middleware"
	"ymst_server/services"
)

type EpisodeHandler struct {
	Log      *logger.Logger
	Game     config.GameConfig
	Episodes *services.EpisodeService
}

func NewEpisodeHandler(log *logger.Logger, game config.GameConfig, episodes *services.EpisodeService) *EpisodeHandler {
	return &EpisodeHandler{Log: log, Game: game, Episodes: episodes}
}

func (h *EpisodeHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/Episodes/:episodeMasterId/GetDetails", h.GetDetails)
	r.POST("/api/Episodes/:episodeMasterId/GetDetails", h.GetDetails)
	r.GET("/api/Episodes/:episodeMasterId/GetEpisodeDetail", h.GetDetails)
	r.POST("/api/Episodes/:episodeMasterId/GetEpisodeDetail", h.GetDetails)
	r.GET("/api/Episodes/GetEpisodeDetail", h.GetDetails)
	r.POST("/api/Episodes/GetEpisodeDetail", h.GetDetails)
	r.GET("/api/Episodes/:episodeMasterId/DetailAsset", h.GetDetailAsset)
	r.GET("/episode-details/:episodeMasterId", h.GetDetailAsset)
}

func (h *EpisodeHandler) GetDetails(c *gin.Context) {
	if _, ok := middleware.UserKeyFromContext(c); !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	episodeMasterID, ok := parseEpisodeMasterID(c)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	payload, err := h.Episodes.GetEpisodeDetail(c.Request.Context(), episodeMasterID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	responseBody, err := msgpack.Marshal(payload)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if h.Log != nil {
		h.Log.GetZap().Info("episode details",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("body_size", len(body)),
			zap.Int("response_size", len(responseBody)),
			zap.Int64("episode_master_id", episodeMasterID),
			zap.Any("episode_title", apiwire.ValueAt(payload, 0)),
			zap.Any("story_type", apiwire.ValueAt(payload, 1)),
			zap.Any("episode_order", apiwire.ValueAt(payload, 2)),
			zap.Any("episode_detail_asset_source", apiwire.ValueAt(payload, 3)),
			zap.String("response_format", "raw_episode_result"),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}
	writeRawGameHeaders(c, h.Game)
	c.Header("Content-Type", apiwire.MsgpackContentType)
	c.Data(http.StatusOK, apiwire.MsgpackContentType, responseBody)
}

func (h *EpisodeHandler) GetDetailAsset(c *gin.Context) {
	episodeMasterID, ok := parseEpisodeMasterID(c)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	payload, err := h.Episodes.GetEpisodeDetailAsset(c.Request.Context(), episodeMasterID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	body, err := msgpack.Marshal(payload)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if h.Log != nil {
		resultCount := len(payload)
		var first []any
		if resultCount > 0 {
			first, _ = payload[0].([]any)
		}
		h.Log.GetZap().Info("episode detail asset",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int64("episode_master_id", episodeMasterID),
			zap.Int("result_count", resultCount),
			zap.Any("first_id", apiwire.ValueAt(first, 0)),
			zap.Any("first_episode_master_id", apiwire.ValueAt(first, 1)),
			zap.Any("first_order", apiwire.ValueAt(first, 2)),
			zap.Any("first_phrase", apiwire.ValueAt(first, 7)),
			zap.Any("first_title", apiwire.ValueAt(first, 8)),
			zap.Int("size", len(body)),
		)
	}
	writeRawGameHeaders(c, h.Game)
	c.Header("Content-Type", apiwire.MsgpackContentType)
	c.Data(http.StatusOK, apiwire.MsgpackContentType, body)
}

func writeRawGameHeaders(c *gin.Context, game config.GameConfig) {
	c.Header("X-Client-Version", game.ClientVersion())
	c.Header("X-Assets-Version", game.APIAssetsVersion())
}

func parseEpisodeMasterID(c *gin.Context) (int64, bool) {
	idStr := strings.TrimSpace(c.Query("episodeMasterId"))
	if idStr == "" {
		idStr = strings.TrimSpace(c.Param("episodeMasterId"))
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, false
	}
	return id, true
}
