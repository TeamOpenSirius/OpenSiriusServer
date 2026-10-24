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

type LeagueHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
}

func NewLeagueHandler(log *logger.Logger, game config.GameConfig) *LeagueHandler {
	return &LeagueHandler{
		Log:  log,
		Game: game,
	}
}

func (h *LeagueHandler) RegisterRoutes(r gin.IRouter) {
	r.POST("/api/Leagues/TopMenuInformation", h.TopMenuInformation)
}

func (h *LeagueHandler) TopMenuInformation(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)

	if h.Log != nil {
		h.Log.GetZap().Info("league top menu information",
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := writeLeagueTopMenuInformationResult(c, true); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func writeLeagueTopMenuInformationResult(c *gin.Context, participated bool) error {
	var body bytes.Buffer
	enc := msgpack.NewEncoder(&body)

	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := encodeLeagueTopMenuInformationResult(enc, seedLeagueTopMenuInformationResult(participated)); err != nil {
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

func seedLeagueTopMenuInformationResult(participated bool) dto.LeagueTopMenuInformationResult {
	groupRanking := 0
	groupLowestRank := 20
	if participated {
		groupRanking = 15
		groupLowestRank = 21
	}

	upClassBorderScore := int64(0)
	upClassBorderRanking := 0
	enrolledGroupNumber := 29

	return dto.LeagueTopMenuInformationResult{
		DisplayStartAt:       time.Date(2026, 5, 12, 5, 0, 0, 0, time.UTC),
		CountingStartAt:      time.Date(2026, 5, 18, 0, 0, 0, 0, time.UTC),
		DisplayEndAt:         time.Date(2026, 5, 19, 5, 0, 0, 0, time.UTC),
		MusicMasterId:        10028,
		LeagueGroupRanking:   groupRanking,
		UpClassBorderScore:   &upClassBorderScore,
		UpClassBorderRanking: &upClassBorderRanking,
		GroupLowestRank:      &groupLowestRank,
		IsParticipated:       participated,
		EnrolledGroupNumber:  &enrolledGroupNumber,
	}
}

func encodeLeagueTopMenuInformationResult(enc *msgpack.Encoder, result dto.LeagueTopMenuInformationResult) error {
	if err := enc.EncodeArrayLen(15); err != nil {
		return err
	}
	if err := enc.Encode(result.DisplayStartAt); err != nil {
		return err
	}
	if err := enc.Encode(result.CountingStartAt); err != nil {
		return err
	}
	if err := enc.Encode(result.DisplayEndAt); err != nil {
		return err
	}
	if err := enc.EncodeInt64(result.MusicMasterId); err != nil {
		return err
	}
	if err := enc.EncodeNil(); err != nil {
		return err
	}
	if err := enc.EncodeNil(); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(result.LeagueGroupRanking)); err != nil {
		return err
	}
	if err := enc.EncodeNil(); err != nil {
		return err
	}
	if err := encodeOptionalInt64(enc, result.UpClassBorderScore); err != nil {
		return err
	}
	if err := encodeOptionalInt(enc, result.UpClassBorderRanking); err != nil {
		return err
	}
	if err := encodeOptionalInt64(enc, result.KeepClassBorderScore); err != nil {
		return err
	}
	if err := encodeOptionalInt(enc, result.KeepClassBorderRanking); err != nil {
		return err
	}
	if err := encodeOptionalInt(enc, result.GroupLowestRank); err != nil {
		return err
	}
	if err := enc.EncodeBool(result.IsParticipated); err != nil {
		return err
	}
	return encodeOptionalInt(enc, result.EnrolledGroupNumber)
}

func encodeOptionalInt64(enc *msgpack.Encoder, value *int64) error {
	if value == nil {
		return enc.EncodeNil()
	}
	return enc.EncodeInt64(*value)
}

func encodeOptionalInt(enc *msgpack.Encoder, value *int) error {
	if value == nil {
		return enc.EncodeNil()
	}
	return enc.EncodeInt(int64(*value))
}
