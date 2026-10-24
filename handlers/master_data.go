package handlers

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ymst_server/api"
	"ymst_server/core/config"
	"ymst_server/core/logger"
)

type MasterDataHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
}

func NewMasterDataHandler(log *logger.Logger, game config.GameConfig) *MasterDataHandler {
	return &MasterDataHandler{
		Log:  log,
		Game: game,
	}
}

func (h *MasterDataHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/api/data/master", h.GetMasterData)
	r.POST("/api/data/master", h.GetMasterData)
	r.GET("/master-data/production/*filepath", h.GetMasterDataFile)
	r.GET("/production/*filepath", h.GetAssetFile)
}

func (h *MasterDataHandler) GetMasterData(c *gin.Context) {
	uri := strings.TrimSpace(h.Game.MasterDataURI)
	if uri == "" {
		uri = "2026-05-09/mastermemory_1778239615_1778239615.db"
	}
	version := strings.TrimSpace(h.Game.MasterDataVersion)
	if version == "" {
		version = "1778239615_1778239615"
	}
	publishTimestamp := h.Game.MasterDataPublishTimestamp
	if publishTimestamp <= 0 {
		publishTimestamp = 1778239615
	}

	c.Header("X-Client-Version", h.Game.ClientVersion())
	c.Header("X-Assets-Version", h.Game.APIAssetsVersion())
	c.Header("X-FM", "0")
	c.Header("Cache-Control", "no-cache")

	if h.Log != nil {
		h.Log.GetZap().Info("game master data manifest",
			zap.String("version", version),
			zap.String("uri", uri),
			zap.Int64("publish_timestamp", publishTimestamp),
		)
	}

	payload := []any{
		uri,
		"",
		version,
		publishTimestamp,
	}

	if err := api.WriteApiMsgpackResult(c, payload); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *MasterDataHandler) GetMasterDataFile(c *gin.Context) {
	h.serveFile(c, c.Param("filepath"), h.Game.MasterDataRoot, "data/master-data/production", "master data file")
}

func (h *MasterDataHandler) GetAssetFile(c *gin.Context) {
	h.serveFile(c, c.Param("filepath"), h.Game.AssetRoot, "data/assets/production", "asset file")
}

func (h *MasterDataHandler) serveFile(c *gin.Context, requestPath string, configuredRoot string, defaultRoot string, logName string) {
	rawPath := strings.TrimPrefix(c.Param("filepath"), "/")
	if requestPath != "" {
		rawPath = strings.TrimPrefix(requestPath, "/")
	}
	if rawPath == "" || strings.Contains(rawPath, "..") || filepath.IsAbs(rawPath) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rootDir := strings.TrimSpace(configuredRoot)
	if rootDir == "" {
		rootDir = defaultRoot
	}

	fullPath := filepath.Join(rootDir, filepath.FromSlash(rawPath))
	cleanRoot, err := filepath.Abs(rootDir)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	cleanFull, err := filepath.Abs(fullPath)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if cleanFull != cleanRoot && !strings.HasPrefix(cleanFull, cleanRoot+string(os.PathSeparator)) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info(logName+" request",
			zap.String("path", cleanFull),
		)
	}

	info, err := os.Stat(cleanFull)
	if err != nil || info.IsDir() {
		if h.Log != nil {
			h.Log.GetZap().Warn(logName+" not found",
				zap.String("path", cleanFull),
				zap.Error(err),
			)
		}
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Header("Accept-Ranges", "bytes")
	c.Header("Cache-Control", "no-cache")
	if contentType := mime.TypeByExtension(filepath.Ext(cleanFull)); contentType != "" {
		c.Header("Content-Type", contentType)
	} else {
		c.Header("Content-Type", "application/octet-stream")
	}
	c.File(cleanFull)
}
