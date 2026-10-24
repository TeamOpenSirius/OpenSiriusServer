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
	"ymst_server/core/config"
	"ymst_server/core/logger"
	"ymst_server/dto"
)

type TripleCastHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
}

func NewTripleCastHandler(log *logger.Logger, game config.GameConfig) *TripleCastHandler {
	return &TripleCastHandler{
		Log:  log,
		Game: game,
	}
}

func (h *TripleCastHandler) RegisterRoutes(r gin.IRouter) {
	r.POST("/api/TripleCast/TopMenuInformation", h.TopMenuInformation)
	r.POST("/api/Leagues/TripleCastTopMenuInformation", h.TopMenuInformation)
}

func (h *TripleCastHandler) TopMenuInformation(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)

	if h.Log != nil {
		h.Log.GetZap().Info("triple cast top menu information",
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := writeTripleCastTopMenuInformationResult(c); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func writeTripleCastTopMenuInformationResult(c *gin.Context) error {
	var body bytes.Buffer
	enc := msgpack.NewEncoder(&body)

	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := encodeLeagueTopMenuInformationResult(enc, seedTripleCastTopMenuInformationResult()); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}

	c.Data(http.StatusOK, api.MsgpackContentType, body.Bytes())
	return nil
}

func seedTripleCastTopMenuInformationResult() dto.LeagueTopMenuInformationResult {
	upClassBorderScore := int64(0)
	upClassBorderRanking := 0
	groupLowestRank := 6
	enrolledGroupNumber := 19

	return dto.LeagueTopMenuInformationResult{
		DisplayStartAt:       time.Date(2026, 5, 15, 5, 0, 0, 0, time.UTC),
		CountingStartAt:      time.Date(2026, 5, 28, 0, 0, 0, 0, time.UTC),
		DisplayEndAt:         time.Date(2026, 6, 1, 5, 0, 0, 0, time.UTC),
		MusicMasterId:        10030,
		LeagueGroupRanking:   0,
		UpClassBorderScore:   &upClassBorderScore,
		UpClassBorderRanking: &upClassBorderRanking,
		GroupLowestRank:      &groupLowestRank,
		IsParticipated:       false,
		EnrolledGroupNumber:  &enrolledGroupNumber,
	}
}
