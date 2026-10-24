package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"ymst_server/api"
	"ymst_server/core/config"
	"ymst_server/core/database"
	"ymst_server/core/logger"
	"ymst_server/dto"
)

const defaultGameUserKeyForLive = "default"
const tripleCastOrderBase = 1

type LiveHandler struct {
	Log  *logger.Logger
	Game config.GameConfig
	DB   *gorm.DB
}

type liveStatus struct {
	Concentration int
	Expression    int
	Vocal         int
}

type liveActor struct {
	ID                             int64
	CharacterBaseMasterID          int64
	CharacterMasterID              int64
	AwakeningPhase                 int
	TalentStage                    int
	Position                       int
	SenseLevel                     int
	BaseStatus                     liveStatus
	CurrentStatus                  liveStatus
	DisplayAwakeningStatus         bool
	SecondaryCharacterBaseMasterID *int64
	SecondarySenseLevel            int
	SelectionType                  int
}

type liveTimingEvent struct {
	TotalSenseLights []int
	GrantSenseLights []int
	IsStarAct        bool
	LostLights       bool
	AcquirableLights []int
	SenseIDs         []int64
	SenseVoiceIDs    []int64
	SenseMasterID    int64
}

type liveSense struct {
	ID                 int64
	ActorID            int64
	CoolTime           int
	SenseType          int
	AcquirableLights   []int
	PreSenseEffect     []liveEffect
	SenseEffect        liveSenseEffect
	PosterEffect       []liveEffect
	AccessoryEffect    []liveEffect
	CombinationSenseID []int64
	SenseMasterID      int64
	OriginalActorID    int64
}

type liveSenseEffect struct {
	ScoreFactor     int
	Principal       int
	BranchCondition int
	EffectBranches  []liveEffectBranch
}

type liveEffectBranch struct {
	ConditionValue *int64
	JudgeTypes     *int
	SenseEffects   []liveEffect
}

type liveEffect struct {
	Order       int
	MasterID    int64
	EffectTypes int
	Triggers    []liveEffectTrigger
	Targets     []liveEffectTargetValue
	Duration    int
}

type liveEffectTrigger struct {
	Type  int
	Value int64
}

type liveEffectTargetValue struct {
	TargetActorID *int64
	Value         float64
}

type liveStarAct struct {
	ScoreFactor     int
	ActorID         int64
	BranchCondition int
	EffectBranches  []liveEffectBranch
}

type liveUnit struct {
	Actors                             map[int]liveActor
	TimeEvents                         map[int]liveTimingEvent
	PossibleSenses                     []liveSense
	StartEffects                       []liveEffect
	StarAct                            liveStarAct
	TotalStatus                        int
	StarActSenseLightCount             int
	MaxPrincipal                       int
	IsFirstPlayOlivier                 bool
	BaseScorePercentage                int
	BaseScoreDifficultyAutoCoefficient float64
	UActiveLiveID                      int64
}

type liveUnitWithOrder struct {
	LiveUnit *liveUnit
	Order    int
}

type startTripleCastLivePayload struct {
	MLiveID       int64
	MTripleCastID int64
}

func NewLiveHandler(log *logger.Logger, game config.GameConfig, db *gorm.DB) *LiveHandler {
	return &LiveHandler{Log: log, Game: game, DB: db}
}

func (h *LiveHandler) RegisterRoutes(r gin.IRouter) {
	r.POST("/api/Lives/Start", h.Start)
	r.POST("/api/Lives/StartTripleCastLive", h.StartTripleCastLive)
	r.POST("/api/Lives/FinishAndValidate", h.FinishAndValidate)
	r.POST("/api/Lives/FinishAnotherNotationLive", h.FinishAnotherNotationLive)
}

func (h *LiveHandler) Start(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	payload, err := decodeStartLivePayload(body)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Warn("decode start live payload failed", zap.Error(err), zap.Int("body_size", len(body)))
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	activeLive, err := h.createActiveLive(payload)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("create active live failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	unit := h.seedLiveUnit(payload, activeLive.ID)
	if h.Log != nil {
		h.Log.GetZap().Info("live start",
			zap.Int64("party_id", payload.PartyId),
			zap.Int64("live_master_id", payload.LiveMasterId),
			zap.Int64("active_live_id", activeLive.ID),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", c.GetHeader("Authorization") != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := writeStartLiveResult(c, unit, nil); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *LiveHandler) StartTripleCastLive(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	payload, err := decodeStartTripleCastLivePayload(body)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Warn("decode start triple cast live payload failed",
				zap.Error(err),
				zap.Int("body_size", len(body)),
			)
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	activeLives, err := h.createTripleCastActiveLives(payload)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("create triple cast active lives failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	units := make([]*liveUnitWithOrder, 0, len(activeLives))
	for i, active := range activeLives {
		order := tripleCastOrderBase + i
		units = append(units, &liveUnitWithOrder{
			LiveUnit: h.seedTripleCastLiveUnit(payload.MLiveID, active.ID, order),
			Order:    order,
		})
	}

	if h.Log != nil {
		h.Log.GetZap().Info("start triple cast live",
			zap.Int64("m_live_id", payload.MLiveID),
			zap.Int64("m_triple_cast_id", payload.MTripleCastID),
			zap.Int("units", len(units)),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", c.GetHeader("Authorization") != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := writeStartTripleCastLiveResult(c, units); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *LiveHandler) FinishAndValidate(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	activeLive := h.findLatestSeedActiveLive()
	h.markSeedActiveLiveFinished(activeLive)

	if h.Log != nil {
		h.Log.GetZap().Info("live finish seed",
			zap.Int64("active_live_id", activeLive.ID),
			zap.Int64("live_master_id", activeLive.LiveMasterID),
			zap.Int64("party_id", activeLive.PartyID),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", c.GetHeader("Authorization") != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := writeFinishLiveSeedResult(c, activeLive); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *LiveHandler) FinishAnotherNotationLive(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	anotherNotationMasterID := decodeFinishAnotherNotationMasterID(body)
	activeLive := h.findLatestSeedActiveLive()
	h.markSeedActiveLiveFinished(activeLive)

	if h.Log != nil {
		h.Log.GetZap().Info("another notation live finish seed",
			zap.Int64("active_live_id", activeLive.ID),
			zap.Int64("live_master_id", activeLive.LiveMasterID),
			zap.Int64("party_id", activeLive.PartyID),
			zap.Int64("another_notation_master_id", anotherNotationMasterID),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", c.GetHeader("Authorization") != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := writeFinishLiveSeedResult(c, activeLive); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *LiveHandler) findLatestSeedActiveLive() *database.GameActiveLive {
	if h.DB != nil {
		var active database.GameActiveLive
		err := h.DB.
			Where("user_key = ? AND status = ?", defaultGameUserKeyForLive, 0).
			Order("started_at DESC, id DESC").
			First(&active).Error
		if err == nil {
			return &active
		}
		if h.Log != nil {
			h.Log.GetZap().Warn("active live not found, use seed fallback", zap.Error(err))
		}
	}

	now := time.Now().UTC()
	return &database.GameActiveLive{
		ID:                      now.UnixNano(),
		UserKey:                 defaultGameUserKeyForLive,
		PartyID:                 65394969,
		LiveMasterID:            10030,
		LiveSettingMasterID:     1,
		UseStamina:              false,
		StaminaConsumptionRatio: 0,
		IsAutoPlay:              false,
		Status:                  0,
		StartedAt:               now.Add(-3 * time.Minute),
	}
}

func (h *LiveHandler) markSeedActiveLiveFinished(activeLive *database.GameActiveLive) {
	if h.DB == nil || activeLive == nil || activeLive.ID == 0 {
		return
	}
	now := time.Now().UTC()
	q := h.DB.Model(&database.GameActiveLive{}).
		Where("user_key = ? AND status = ?", activeLive.UserKey, 0)
	if activeLive.LeagueMasterID != nil {
		q = q.Where("league_master_id = ?", *activeLive.LeagueMasterID)
	} else {
		q = q.Where("id = ?", activeLive.ID)
	}
	if err := q.Updates(map[string]any{
		"status":      1,
		"finished_at": now,
		"updated_at":  now,
	}).Error; err != nil && h.Log != nil {
		h.Log.GetZap().Warn("mark active live finished failed", zap.Error(err), zap.Int64("active_live_id", activeLive.ID))
	}
}

func decodeStartTripleCastLivePayload(body []byte) (*startTripleCastLivePayload, error) {
	var values []any
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&values); err != nil {
		return nil, err
	}
	if len(values) < 2 {
		return nil, fmt.Errorf("invalid StartTripleCastLivePayload length: %d", len(values))
	}
	return &startTripleCastLivePayload{
		MLiveID:       int64ValueLocal(values[0]),
		MTripleCastID: int64ValueLocal(values[1]),
	}, nil
}

func decodeStartLivePayload(body []byte) (*dto.StartLivePayload, error) {
	var values []any
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&values); err != nil {
		return nil, err
	}
	if len(values) < 15 {
		return nil, fmt.Errorf("invalid StartLivePayload length: %d", len(values))
	}
	return &dto.StartLivePayload{
		PartyId:                      int64ValueLocal(values[0]),
		LiveMasterId:                 int64ValueLocal(values[1]),
		AuditionMasterId:             nullableInt64Local(values[2]),
		LeagueMasterId:               nullableInt64Local(values[3]),
		UseStamina:                   boolValueLocal(values[4]),
		StaminaConsumptionRatio:      int(int64ValueLocal(values[5])),
		LiveSettingMasterId:          int64ValueLocal(values[6]),
		IsAutoPlay:                   boolValueLocal(values[7]),
		IsStoryEventChallenge:        boolValueLocal(values[8]),
		ConcertStageMasterId:         nullableInt64Local(values[9]),
		BonusLiveStageMasterId:       nullableInt64Local(values[10]),
		MusicCourseDetailMasterId:    nullableInt64Local(values[11]),
		MusicCourseGaugeType:         dto.MusicCourseGaugeType(int64ValueLocal(values[12])),
		GhostLiveMasterId:            nullableInt64Local(values[13]),
		TrialPartyEventStageMasterId: nullableInt64Local(values[14]),
	}, nil
}

func decodeFinishAnotherNotationMasterID(body []byte) int64 {
	var values []any
	if err := msgpack.NewDecoder(bytes.NewReader(body)).Decode(&values); err != nil || len(values) < 5 {
		return 0
	}
	return int64ValueLocal(values[4])
}

func (h *LiveHandler) createActiveLive(payload *dto.StartLivePayload) (*database.GameActiveLive, error) {
	now := time.Now().UTC()
	activeLive := &database.GameActiveLive{
		ID:                      now.UnixNano(),
		UserKey:                 defaultGameUserKeyForLive,
		PartyID:                 payload.PartyId,
		LiveMasterID:            payload.LiveMasterId,
		LiveSettingMasterID:     payload.LiveSettingMasterId,
		AuditionMasterID:        payload.AuditionMasterId,
		LeagueMasterID:          payload.LeagueMasterId,
		UseStamina:              payload.UseStamina,
		StaminaConsumptionRatio: payload.StaminaConsumptionRatio,
		IsAutoPlay:              payload.IsAutoPlay,
		Status:                  0,
		StartedAt:               now,
	}
	if h.DB == nil {
		return activeLive, nil
	}
	return activeLive, h.DB.Create(activeLive).Error
}

func (h *LiveHandler) createTripleCastActiveLives(payload *startTripleCastLivePayload) ([]*database.GameActiveLive, error) {
	now := time.Now().UTC()
	tripleCastID := payload.MTripleCastID
	result := make([]*database.GameActiveLive, 0, 3)

	for i := 0; i < 3; i++ {
		order := tripleCastOrderBase + i
		active := &database.GameActiveLive{
			ID:                      now.UnixNano() + int64(i+1),
			UserKey:                 defaultGameUserKeyForLive,
			PartyID:                 int64(order),
			LiveMasterID:            payload.MLiveID,
			LiveSettingMasterID:     1,
			LeagueMasterID:          &tripleCastID,
			UseStamina:              false,
			StaminaConsumptionRatio: 0,
			IsAutoPlay:              false,
			Status:                  0,
			StartedAt:               now,
		}
		if h.DB != nil {
			if err := h.DB.Create(active).Error; err != nil {
				return nil, err
			}
		}
		result = append(result, active)
	}

	return result, nil
}

func (h *LiveHandler) seedLiveUnit(payload *dto.StartLivePayload, activeLiveID int64) *liveUnit {
	actors := h.seedActors(payload.PartyId)
	possibleSenses := make([]liveSense, 0, len(actors))
	totalStatus := 0

	positions := make([]int, 0, len(actors))
	for position := range actors {
		positions = append(positions, position)
	}
	sort.Ints(positions)

	for _, position := range positions {
		actor := actors[position]
		totalStatus += actor.CurrentStatus.total()
		possibleSenses = append(possibleSenses, seedSenseForActor(actor))
	}

	starActorID := int64(210365700)
	if actor, ok := actors[5]; ok {
		starActorID = actor.ID
	}

	return &liveUnit{
		Actors:                             actors,
		TimeEvents:                         seedTimingEvents(payload.LiveMasterId),
		PossibleSenses:                     possibleSenses,
		StartEffects:                       []liveEffect{},
		StarAct:                            liveStarAct{ScoreFactor: 470, ActorID: starActorID, BranchCondition: 0, EffectBranches: []liveEffectBranch{}},
		TotalStatus:                        totalStatus,
		StarActSenseLightCount:             5,
		MaxPrincipal:                       1000,
		IsFirstPlayOlivier:                 false,
		BaseScorePercentage:                0,
		BaseScoreDifficultyAutoCoefficient: calcBaseScoreDifficultyAutoCoefficient(payload.LiveMasterId),
		UActiveLiveID:                      activeLiveID,
	}
}

func (h *LiveHandler) seedTripleCastLiveUnit(liveMasterID int64, activeLiveID int64, order int) *liveUnit {
	actors := seedTripleCastActors(order)
	possibleSenses := make([]liveSense, 0, len(actors))
	totalStatus := 0

	positions := make([]int, 0, len(actors))
	for position := range actors {
		positions = append(positions, position)
	}
	sort.Ints(positions)

	for _, position := range positions {
		actor := actors[position]
		totalStatus += actor.CurrentStatus.total()
		possibleSenses = append(possibleSenses, seedTripleCastSenseForActor(actor, order))
	}

	starActorID := int64(0)
	if actor, ok := actors[5]; ok {
		starActorID = actor.ID
	}

	return &liveUnit{
		Actors:                             actors,
		TimeEvents:                         seedTripleCastTimingEvents(liveMasterID, actors),
		PossibleSenses:                     possibleSenses,
		StartEffects:                       []liveEffect{},
		StarAct:                            liveStarAct{ScoreFactor: 470, ActorID: starActorID, BranchCondition: 0, EffectBranches: []liveEffectBranch{}},
		TotalStatus:                        totalStatus,
		StarActSenseLightCount:             5,
		MaxPrincipal:                       1000,
		IsFirstPlayOlivier:                 false,
		BaseScorePercentage:                0,
		BaseScoreDifficultyAutoCoefficient: calcBaseScoreDifficultyAutoCoefficient(liveMasterID),
		UActiveLiveID:                      activeLiveID,
	}
}

func tripleCastActorOffset(order int) int64 {
	if order <= 1 {
		return 0
	}
	return int64(order-1) * 1000000
}

func seedTripleCastActors(order int) map[int]liveActor {
	offset := tripleCastActorOffset(order)
	return map[int]liveActor{
		1: {ID: 210365704 + offset, CharacterBaseMasterID: 301, CharacterMasterID: 140970, AwakeningPhase: 0, TalentStage: 0, Position: 1, SenseLevel: 3, BaseStatus: liveStatus{Concentration: 3122, Expression: 2698, Vocal: 1917}, CurrentStatus: liveStatus{Concentration: 3122, Expression: 2698, Vocal: 1917}, DisplayAwakeningStatus: false, SecondarySenseLevel: 0, SelectionType: 1},
		2: {ID: 210365705 + offset, CharacterBaseMasterID: 103, CharacterMasterID: 140270, AwakeningPhase: 0, TalentStage: 0, Position: 2, SenseLevel: 4, BaseStatus: liveStatus{Concentration: 3252, Expression: 3096, Vocal: 2549}, CurrentStatus: liveStatus{Concentration: 3252, Expression: 3096, Vocal: 2549}, DisplayAwakeningStatus: false, SecondarySenseLevel: 0, SelectionType: 1},
		3: {ID: 210365709 + offset, CharacterBaseMasterID: 405, CharacterMasterID: 130210, AwakeningPhase: 0, TalentStage: 0, Position: 3, SenseLevel: 1, BaseStatus: liveStatus{Concentration: 24, Expression: 20, Vocal: 23}, CurrentStatus: liveStatus{Concentration: 24, Expression: 20, Vocal: 23}, DisplayAwakeningStatus: false, SecondarySenseLevel: 0, SelectionType: 1},
		4: {ID: 210365711 + offset, CharacterBaseMasterID: 404, CharacterMasterID: 130200, AwakeningPhase: 0, TalentStage: 0, Position: 4, SenseLevel: 1, BaseStatus: liveStatus{Concentration: 20, Expression: 25, Vocal: 22}, CurrentStatus: liveStatus{Concentration: 20, Expression: 25, Vocal: 22}, DisplayAwakeningStatus: false, SecondarySenseLevel: 0, SelectionType: 1},
		5: {ID: 210365700 + offset, CharacterBaseMasterID: 402, CharacterMasterID: 130180, AwakeningPhase: 0, TalentStage: 0, Position: 5, SenseLevel: 1, BaseStatus: liveStatus{Concentration: 27, Expression: 24, Vocal: 27}, CurrentStatus: liveStatus{Concentration: 27, Expression: 24, Vocal: 27}, DisplayAwakeningStatus: false, SecondarySenseLevel: 0, SelectionType: 1},
	}
}

func (h *LiveHandler) seedActors(partyID int64) map[int]liveActor {
	if h.DB != nil {
		var slots []database.GamePartySlot
		if err := h.DB.Where("party_id = ?", partyID).Order("position ASC").Find(&slots).Error; err == nil && len(slots) == 5 {
			actors := map[int]liveActor{}
			for _, slot := range slots {
				if slot.Position < 1 || slot.Position > 5 {
					continue
				}
				var ch database.GameCharacter
				if err := h.DB.First(&ch, "id = ?", slot.CharacterID).Error; err == nil {
					actors[slot.Position] = actorFromCharacter(slot.Position, ch)
				}
			}
			if len(actors) == 5 {
				return actors
			}
		}
	}

	return seedPacketActors()
}

func actorFromCharacter(position int, ch database.GameCharacter) liveActor {
	status := liveStatus{
		Concentration: 1000 + ch.Level*10,
		Expression:    1000 + ch.TalentStage*20,
		Vocal:         1000 + ch.SenseLevel*30,
	}
	return liveActor{
		ID:                             ch.ID,
		CharacterBaseMasterID:          ch.CharacterBaseID,
		CharacterMasterID:              ch.CharacterMasterID,
		AwakeningPhase:                 ch.AwakeningPhase,
		TalentStage:                    ch.TalentStage,
		Position:                       position,
		SenseLevel:                     ch.SenseLevel,
		BaseStatus:                     status,
		CurrentStatus:                  status,
		DisplayAwakeningStatus:         ch.DisplayAwakeningStatus,
		SecondaryCharacterBaseMasterID: ch.SecondaryCharacterBaseID,
		SecondarySenseLevel:            ch.SecondarySenseLevel,
		SelectionType:                  ch.SelectionType,
	}
}

func basicSenseEffect(actorID int64) liveSenseEffect {
	return liveSenseEffect{
		ScoreFactor:     700,
		Principal:       0,
		BranchCondition: 0,
		EffectBranches: []liveEffectBranch{
			{
				SenseEffects: []liveEffect{
					{
						Order:       0,
						MasterID:    actorID%100000 + 20000,
						EffectTypes: 41,
						Triggers:    []liveEffectTrigger{},
						Targets:     []liveEffectTargetValue{{Value: 50.0}},
						Duration:    0,
					},
				},
			},
		},
	}
}

func seedPacketActors() map[int]liveActor {
	return map[int]liveActor{
		1: seedPacketActor(210365704, 301, 140970, 1, 3, liveStatus{Concentration: 3122, Expression: 2698, Vocal: 1917}),
		2: seedPacketActor(210365705, 103, 140270, 2, 4, liveStatus{Concentration: 3252, Expression: 3096, Vocal: 2549}),
		3: seedPacketActor(210365709, 405, 130210, 3, 1, liveStatus{Concentration: 24, Expression: 20, Vocal: 23}),
		4: seedPacketActor(210365711, 404, 130200, 4, 1, liveStatus{Concentration: 20, Expression: 25, Vocal: 22}),
		5: seedPacketActor(210365700, 402, 130180, 5, 1, liveStatus{Concentration: 27, Expression: 24, Vocal: 27}),
	}
}

func seedPacketActor(id int64, characterBaseMasterID int64, characterMasterID int64, position int, senseLevel int, status liveStatus) liveActor {
	return liveActor{
		ID:                     id,
		CharacterBaseMasterID:  characterBaseMasterID,
		CharacterMasterID:      characterMasterID,
		AwakeningPhase:         0,
		TalentStage:            0,
		Position:               position,
		SenseLevel:             senseLevel,
		BaseStatus:             status,
		CurrentStatus:          status,
		DisplayAwakeningStatus: false,
		SecondarySenseLevel:    0,
		SelectionType:          1,
	}
}

func seedSenseForActor(actor liveActor) liveSense {
	switch actor.ID {
	case 210365704:
		return liveSense{
			ID:               21036570400,
			ActorID:          210365704,
			CoolTime:         40,
			SenseType:        4,
			AcquirableLights: []int{4},
			PreSenseEffect:   []liveEffect{},
			SenseEffect: liveSenseEffect{
				ScoreFactor:     700,
				Principal:       0,
				BranchCondition: 0,
				EffectBranches: []liveEffectBranch{
					{
						SenseEffects: []liveEffect{
							{
								Order:       0,
								MasterID:    20095,
								EffectTypes: 41,
								Triggers:    []liveEffectTrigger{},
								Targets:     []liveEffectTargetValue{{Value: 50.0}},
								Duration:    0,
							},
						},
					},
				},
			},
			PosterEffect:       []liveEffect{},
			AccessoryEffect:    []liveEffect{},
			CombinationSenseID: []int64{},
			SenseMasterID:      140970,
			OriginalActorID:    210365704,
		}
	case 210365705:
		return packetSense(actor, 21036570500, 40, 3, []int{3}, 750, 0, 140270)
	case 210365709:
		return packetSense(actor, 21036570900, 30, 2, []int{2}, 360, 0, 130210)
	case 210365711:
		return packetSense(actor, 21036571100, 50, 4, []int{4}, 140, 80, 130200)
	case 210365700:
		targetA := int64(210365709)
		targetB := int64(210365711)
		targetC := int64(210365700)
		return liveSense{
			ID:               21036570000,
			ActorID:          210365700,
			CoolTime:         40,
			SenseType:        3,
			AcquirableLights: []int{3},
			PreSenseEffect:   []liveEffect{},
			SenseEffect: liveSenseEffect{
				ScoreFactor:     140,
				Principal:       0,
				BranchCondition: 0,
				EffectBranches: []liveEffectBranch{
					{
						SenseEffects: []liveEffect{
							{
								Order:       0,
								MasterID:    20063,
								EffectTypes: 30,
								Triggers:    []liveEffectTrigger{},
								Targets: []liveEffectTargetValue{
									{TargetActorID: &targetA, Value: 21.0},
									{TargetActorID: &targetB, Value: 21.0},
									{TargetActorID: &targetC, Value: 21.0},
								},
								Duration: 25,
							},
						},
					},
				},
			},
			PosterEffect:       []liveEffect{},
			AccessoryEffect:    []liveEffect{},
			CombinationSenseID: []int64{},
			SenseMasterID:      130180,
			OriginalActorID:    210365700,
		}
	default:
		return liveSense{
			ID:                 actor.ID * 100,
			ActorID:            actor.ID,
			CoolTime:           40,
			SenseType:          4,
			AcquirableLights:   []int{4},
			PreSenseEffect:     []liveEffect{},
			SenseEffect:        basicSenseEffect(actor.ID),
			PosterEffect:       []liveEffect{},
			AccessoryEffect:    []liveEffect{},
			CombinationSenseID: []int64{},
			SenseMasterID:      actor.CharacterMasterID,
			OriginalActorID:    actor.ID,
		}
	}
}

func seedTripleCastSenseForActor(actor liveActor, order int) liveSense {
	_ = order
	senseID := actor.ID * 100
	effect := liveSenseEffect{ScoreFactor: 700, Principal: 0, BranchCondition: 0, EffectBranches: []liveEffectBranch{}}
	coolTime := 40
	senseType := 4
	lights := []int{4}

	switch actor.Position {
	case 1:
		effect = basicSenseEffect(actor.ID)
	case 2:
		effect = liveSenseEffect{ScoreFactor: 750, Principal: 0, BranchCondition: 0, EffectBranches: []liveEffectBranch{}}
		senseType = 3
		lights = []int{3}
	case 3:
		effect = liveSenseEffect{ScoreFactor: 360, Principal: 0, BranchCondition: 0, EffectBranches: []liveEffectBranch{}}
		coolTime = 30
		senseType = 2
		lights = []int{2}
	case 4:
		effect = liveSenseEffect{ScoreFactor: 140, Principal: 80, BranchCondition: 0, EffectBranches: []liveEffectBranch{}}
		coolTime = 50
	case 5:
		effect = liveSenseEffect{ScoreFactor: 140, Principal: 0, BranchCondition: 0, EffectBranches: []liveEffectBranch{}}
		senseType = 3
		lights = []int{3}
	}

	return liveSense{
		ID:                 senseID,
		ActorID:            actor.ID,
		CoolTime:           coolTime,
		SenseType:          senseType,
		AcquirableLights:   lights,
		PreSenseEffect:     []liveEffect{},
		SenseEffect:        effect,
		PosterEffect:       []liveEffect{},
		AccessoryEffect:    []liveEffect{},
		CombinationSenseID: []int64{},
		SenseMasterID:      actor.CharacterMasterID,
		OriginalActorID:    actor.ID,
	}
}

func packetSense(actor liveActor, id int64, coolTime int, senseType int, lights []int, scoreFactor int, principal int, senseMasterID int64) liveSense {
	return liveSense{
		ID:                 id,
		ActorID:            actor.ID,
		CoolTime:           coolTime,
		SenseType:          senseType,
		AcquirableLights:   lights,
		PreSenseEffect:     []liveEffect{},
		SenseEffect:        liveSenseEffect{ScoreFactor: scoreFactor, Principal: principal, BranchCondition: 0, EffectBranches: []liveEffectBranch{}},
		PosterEffect:       []liveEffect{},
		AccessoryEffect:    []liveEffect{},
		CombinationSenseID: []int64{},
		SenseMasterID:      senseMasterID,
		OriginalActorID:    actor.ID,
	}
}

func seedTimingEvents(liveMasterID int64) map[int]liveTimingEvent {
	seconds := []int{11, 22, 33, 44, 55, 66, 77, 88}
	switch liveMasterID {
	case 23704:
		seconds = []int{10, 20, 30, 40, 50, 60, 70, 80}
	case 11304:
		seconds = []int{14, 28, 42, 56, 70, 84, 98, 112}
	case 25603, 25604:
		seconds = []int{12, 24, 36, 48, 60, 72, 84, 96}
	}
	return map[int]liveTimingEvent{
		seconds[0]: packetTimingEvent([]int{4}, []int{4}, []int{4}, []int64{21036570400}, []int64{30199911, 99940221}),
		seconds[1]: packetTimingEvent([]int{4, 3}, []int{3}, []int{3}, []int64{21036570500}, []int64{10399912, 99940222}),
		seconds[2]: packetTimingEvent([]int{4, 3}, []int{}, []int{2}, []int64{21036570900}, []int64{40599912, 99910321}),
		seconds[3]: packetTimingEvent([]int{4, 3}, []int{}, []int{4}, []int64{21036571100}, []int64{40499912, 99910322}),
		seconds[4]: packetTimingEvent([]int{4, 3, 3}, []int{3}, []int{3}, []int64{21036570000}, []int64{40240411, 40240421}),
		seconds[5]: packetTimingEvent([]int{4, 3, 3}, []int{}, []int{2}, []int64{21036570900}, []int64{40599912, 99910322}),
		seconds[6]: packetTimingEvent([]int{4, 3, 3, 3}, []int{3}, []int{3}, []int64{21036570500}, []int64{10399911, 99940522}),
		seconds[7]: packetTimingEvent([]int{4, 3, 3, 3}, []int{}, []int{4}, []int64{21036570400}, []int64{30199912, 99940521}),
	}
}

func seedTripleCastTimingEvents(liveMasterID int64, actors map[int]liveActor) map[int]liveTimingEvent {
	seconds := []int{11, 22, 33, 44, 55, 66, 77, 88}
	switch liveMasterID {
	case 23704:
		seconds = []int{10, 20, 30, 40, 50, 60, 70, 80}
	case 11304:
		seconds = []int{14, 28, 42, 56, 70, 84, 98, 112}
	case 25603, 25604:
		seconds = []int{12, 24, 36, 48, 60, 72, 84, 96}
	}

	senseID := func(position int) int64 {
		if actor, ok := actors[position]; ok {
			return actor.ID * 100
		}
		return 0
	}
	baseID := func(position int) int64 {
		if actor, ok := actors[position]; ok {
			return actor.CharacterBaseMasterID
		}
		return 999
	}

	return map[int]liveTimingEvent{
		seconds[0]: packetTimingEvent([]int{4}, []int{4}, []int{4}, []int64{senseID(1)}, []int64{baseID(1)*100000 + 99911, 99940221}),
		seconds[1]: packetTimingEvent([]int{4, 3}, []int{3}, []int{3}, []int64{senseID(2)}, []int64{baseID(2)*100000 + 99912, 99940222}),
		seconds[2]: packetTimingEvent([]int{4, 3}, []int{}, []int{2}, []int64{senseID(3)}, []int64{baseID(3)*100000 + 99912, 99910321}),
		seconds[3]: packetTimingEvent([]int{4, 3}, []int{}, []int{4}, []int64{senseID(4)}, []int64{baseID(4)*100000 + 99912, 99910322}),
		seconds[4]: packetTimingEvent([]int{4, 3, 3}, []int{3}, []int{3}, []int64{senseID(5)}, []int64{baseID(5)*100000 + 40411, baseID(5)*100000 + 40421}),
		seconds[5]: packetTimingEvent([]int{4, 3, 3}, []int{}, []int{2}, []int64{senseID(3)}, []int64{baseID(3)*100000 + 99912, 99910322}),
		seconds[6]: packetTimingEvent([]int{4, 3, 3, 3}, []int{3}, []int{3}, []int64{senseID(2)}, []int64{baseID(2)*100000 + 99911, 99940522}),
		seconds[7]: packetTimingEvent([]int{4, 3, 3, 3}, []int{}, []int{4}, []int64{senseID(1)}, []int64{baseID(1)*100000 + 99912, 99940521}),
	}
}

func packetTimingEvent(total []int, grant []int, acquirable []int, senseIDs []int64, senseVoiceIDs []int64) liveTimingEvent {
	return liveTimingEvent{
		TotalSenseLights: total,
		GrantSenseLights: grant,
		IsStarAct:        false,
		LostLights:       false,
		AcquirableLights: acquirable,
		SenseIDs:         senseIDs,
		SenseVoiceIDs:    senseVoiceIDs,
		SenseMasterID:    0,
	}
}

func calcBaseScoreDifficultyAutoCoefficient(liveMasterID int64) float64 {
	if liveMasterID%10 == 4 {
		return 1.05
	}
	return 1.0
}

func writeStartLiveResult(c *gin.Context, unit *liveUnit, updates []any) error {
	var body bytes.Buffer
	enc := msgpack.NewEncoder(&body)
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := encodeLiveUnit(enc, unit); err != nil {
		return err
	}
	if updates == nil {
		if err := enc.EncodeArrayLen(0); err != nil {
			return err
		}
	} else if err := enc.Encode(updates); err != nil {
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

func writeStartTripleCastLiveResult(c *gin.Context, units []*liveUnitWithOrder) error {
	var body bytes.Buffer
	enc := msgpack.NewEncoder(&body)

	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := encodeStartTripleCastLiveResult(enc, units); err != nil {
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

func encodeStartTripleCastLiveResult(enc *msgpack.Encoder, units []*liveUnitWithOrder) error {
	if err := enc.EncodeArrayLen(1); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(len(units)); err != nil {
		return err
	}
	for _, unit := range units {
		if err := encodeLiveUnitWithOrder(enc, unit); err != nil {
			return err
		}
	}
	return nil
}

func writeFinishLiveSeedResult(c *gin.Context, activeLive *database.GameActiveLive) error {
	var body bytes.Buffer
	enc := msgpack.NewEncoder(&body)

	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := encodeFinishLiveSeedResult(enc, activeLive); err != nil {
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

func encodeFinishLiveSeedResult(enc *msgpack.Encoder, activeLive *database.GameActiveLive) error {
	if err := enc.EncodeArrayLen(31); err != nil {
		return err
	}
	if err := enc.EncodeNil(); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := encodeSeedRateResult(enc); err != nil {
		return err
	}
	if err := encodeSeedPlayerRankPointResult(enc); err != nil {
		return err
	}
	if err := enc.EncodeNil(); err != nil {
		return err
	}
	if err := enc.EncodeInt(6); err != nil {
		return err
	}
	if err := enc.EncodeInt(9); err != nil {
		return err
	}
	if err := enc.EncodeNil(); err != nil {
		return err
	}
	if err := enc.EncodeBool(false); err != nil {
		return err
	}
	if err := enc.EncodeUint8(0); err != nil {
		return err
	}
	if err := enc.EncodeUint8(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeFloat64(97.8138); err != nil {
		return err
	}
	for i := 0; i < 6; i++ {
		if err := enc.EncodeNil(); err != nil {
			return err
		}
	}
	if err := enc.EncodeInt(0); err != nil {
		return err
	}
	if err := enc.EncodeInt(0); err != nil {
		return err
	}
	if err := encodeSeedEventResult(enc); err != nil {
		return err
	}
	if err := encodeSeedLiveDropThings(enc); err != nil {
		return err
	}
	for i := 0; i < 6; i++ {
		if err := enc.EncodeNil(); err != nil {
			return err
		}
	}
	return enc.EncodeArrayLen(0)
}

func encodeSeedRateResult(enc *msgpack.Encoder) error {
	if err := enc.EncodeArrayLen(4); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(2); err != nil {
		return err
	}
	if err := enc.EncodeFloat64(0.0); err != nil {
		return err
	}
	// 达成率
	if err := enc.EncodeFloat64(999.0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(2); err != nil {
		return err
	}
	if err := enc.EncodeFloat64(0.0); err != nil {
		return err
	}
	if err := enc.EncodeFloat64(28.05); err != nil {
		return err
	}
	// rating前
	if err := enc.EncodeFloat64(1000.0); err != nil {
		return err
	}
	// rating后
	return enc.EncodeFloat64(1000.0)
}

func encodeSeedPlayerRankPointResult(enc *msgpack.Encoder) error {
	if err := enc.EncodeArrayLen(6); err != nil {
		return err
	}
	for _, v := range []int{15, 15, 282, 456, 174, 849} {
		if err := enc.EncodeInt(int64(v)); err != nil {
			return err
		}
	}
	return nil
}

func encodeSeedEventResult(enc *msgpack.Encoder) error {
	if err := enc.EncodeArrayLen(14); err != nil {
		return err
	}
	if err := enc.EncodeInt(10143); err != nil {
		return err
	}
	for i := 0; i < 6; i++ {
		if err := enc.EncodeNil(); err != nil {
			return err
		}
	}
	if err := enc.EncodeInt(6774); err != nil {
		return err
	}
	if err := enc.EncodeInt(6710); err != nil {
		return err
	}
	if err := enc.EncodeNil(); err != nil {
		return err
	}
	for _, v := range []int{26004, 16874, 3734, 3693} {
		if err := enc.EncodeInt(int64(v)); err != nil {
			return err
		}
	}
	return nil
}

type seedDrop struct {
	Type         int
	ID           int64
	Quantity     int64
	Order        int
	LiveDropType int
	IsFixedDrop  bool
}

func encodeSeedLiveDropThings(enc *msgpack.Encoder) error {
	drops := []seedDrop{
		{Type: 1, ID: 130021, Quantity: 69, Order: 101, LiveDropType: 0, IsFixedDrop: false},
		{Type: 1, ID: 130024, Quantity: 104, Order: 102, LiveDropType: 0, IsFixedDrop: false},
		{Type: 1, ID: 120001, Quantity: 34, Order: 103, LiveDropType: 0, IsFixedDrop: false},
		{Type: 1, ID: 120001, Quantity: 34, Order: 104, LiveDropType: 0, IsFixedDrop: false},
		{Type: 1, ID: 130033, Quantity: 34, Order: 105, LiveDropType: 0, IsFixedDrop: false},
		{Type: 1, ID: 130081, Quantity: 34, Order: 106, LiveDropType: 0, IsFixedDrop: false},
		{Type: 1, ID: 130062, Quantity: 34, Order: 108, LiveDropType: 0, IsFixedDrop: false},
		{Type: 12, ID: 0, Quantity: 3674, Order: 0, LiveDropType: 0, IsFixedDrop: false},
	}
	if err := enc.EncodeArrayLen(len(drops)); err != nil {
		return err
	}
	for _, drop := range drops {
		if err := encodeSeedLiveDropThing(enc, drop); err != nil {
			return err
		}
	}
	return nil
}

func encodeSeedLiveDropThing(enc *msgpack.Encoder, drop seedDrop) error {
	if err := enc.EncodeArrayLen(4); err != nil {
		return err
	}
	if err := encodeSeedReceivedThing(enc, drop.Type, drop.ID, drop.Quantity); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(drop.Order)); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(drop.LiveDropType)); err != nil {
		return err
	}
	return enc.EncodeBool(drop.IsFixedDrop)
}

func encodeSeedReceivedThing(enc *msgpack.Encoder, typ int, id int64, quantity int64) error {
	if err := enc.EncodeArrayLen(7); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(typ)); err != nil {
		return err
	}
	if err := enc.EncodeInt64(id); err != nil {
		return err
	}
	if err := enc.EncodeInt64(quantity); err != nil {
		return err
	}
	for i := 0; i < 3; i++ {
		if err := enc.EncodeNil(); err != nil {
			return err
		}
	}
	return enc.EncodeBool(false)
}

func encodeLiveUnit(enc *msgpack.Encoder, unit *liveUnit) error {
	if err := enc.EncodeArrayLen(12); err != nil {
		return err
	}
	if err := encodeActors(enc, unit.Actors); err != nil {
		return err
	}
	if err := encodeTimingEvents(enc, unit.TimeEvents); err != nil {
		return err
	}
	if err := encodeSenses(enc, unit.PossibleSenses); err != nil {
		return err
	}
	if err := encodeEffects(enc, unit.StartEffects); err != nil {
		return err
	}
	if err := encodeStarAct(enc, unit.StarAct); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(unit.TotalStatus)); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(unit.StarActSenseLightCount)); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(unit.MaxPrincipal)); err != nil {
		return err
	}
	if err := enc.EncodeBool(unit.IsFirstPlayOlivier); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(unit.BaseScorePercentage)); err != nil {
		return err
	}
	if err := enc.EncodeFloat64(unit.BaseScoreDifficultyAutoCoefficient); err != nil {
		return err
	}
	return enc.EncodeInt64(unit.UActiveLiveID)
}

func encodeLiveUnitWithOrder(enc *msgpack.Encoder, u *liveUnitWithOrder) error {
	if u == nil || u.LiveUnit == nil {
		return enc.EncodeNil()
	}
	unit := u.LiveUnit
	if err := enc.EncodeArrayLen(13); err != nil {
		return err
	}
	if err := encodeActors(enc, unit.Actors); err != nil {
		return err
	}
	if err := encodeTimingEvents(enc, unit.TimeEvents); err != nil {
		return err
	}
	if err := encodeSenses(enc, unit.PossibleSenses); err != nil {
		return err
	}
	if err := encodeEffects(enc, unit.StartEffects); err != nil {
		return err
	}
	if err := encodeStarAct(enc, unit.StarAct); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(unit.TotalStatus)); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(unit.StarActSenseLightCount)); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(unit.MaxPrincipal)); err != nil {
		return err
	}
	if err := enc.EncodeBool(unit.IsFirstPlayOlivier); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(unit.BaseScorePercentage)); err != nil {
		return err
	}
	if err := enc.EncodeFloat64(unit.BaseScoreDifficultyAutoCoefficient); err != nil {
		return err
	}
	if err := enc.EncodeInt64(unit.UActiveLiveID); err != nil {
		return err
	}
	return enc.EncodeInt(int64(u.Order))
}

func encodeActors(enc *msgpack.Encoder, actors map[int]liveActor) error {
	keys := make([]int, 0, len(actors))
	for key := range actors {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	if err := enc.EncodeMapLen(len(keys)); err != nil {
		return err
	}
	for _, key := range keys {
		if err := enc.EncodeInt(int64(key)); err != nil {
			return err
		}
		if err := encodeActor(enc, actors[key]); err != nil {
			return err
		}
	}
	return nil
}

func encodeActor(enc *msgpack.Encoder, actor liveActor) error {
	if err := enc.EncodeArrayLen(13); err != nil {
		return err
	}
	values := []any{actor.ID, actor.CharacterBaseMasterID, actor.CharacterMasterID, actor.AwakeningPhase, actor.TalentStage, actor.Position, actor.SenseLevel}
	for _, value := range values {
		if err := enc.Encode(value); err != nil {
			return err
		}
	}
	if err := encodeStatus(enc, actor.BaseStatus); err != nil {
		return err
	}
	if err := encodeStatus(enc, actor.CurrentStatus); err != nil {
		return err
	}
	if err := enc.EncodeBool(actor.DisplayAwakeningStatus); err != nil {
		return err
	}
	if actor.SecondaryCharacterBaseMasterID == nil {
		if err := enc.EncodeNil(); err != nil {
			return err
		}
	} else if err := enc.EncodeInt64(*actor.SecondaryCharacterBaseMasterID); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(actor.SecondarySenseLevel)); err != nil {
		return err
	}
	return enc.EncodeInt(int64(actor.SelectionType))
}

func encodeStatus(enc *msgpack.Encoder, status liveStatus) error {
	return enc.Encode([]any{status.Concentration, status.Expression, status.Vocal, status.total()})
}

func encodeTimingEvents(enc *msgpack.Encoder, events map[int]liveTimingEvent) error {
	keys := make([]int, 0, len(events))
	for key := range events {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	if err := enc.EncodeMapLen(len(keys)); err != nil {
		return err
	}
	for _, key := range keys {
		if err := enc.EncodeInt(int64(key)); err != nil {
			return err
		}
		if err := enc.Encode([]any{events[key].TotalSenseLights, events[key].GrantSenseLights, events[key].IsStarAct, events[key].LostLights, events[key].AcquirableLights, events[key].SenseIDs, events[key].SenseVoiceIDs, events[key].SenseMasterID}); err != nil {
			return err
		}
	}
	return nil
}

func encodeSenses(enc *msgpack.Encoder, senses []liveSense) error {
	if err := enc.EncodeArrayLen(len(senses)); err != nil {
		return err
	}
	for _, sense := range senses {
		if err := encodeSense(enc, sense); err != nil {
			return err
		}
	}
	return nil
}

func encodeSense(enc *msgpack.Encoder, sense liveSense) error {
	if err := enc.EncodeArrayLen(13); err != nil {
		return err
	}
	if err := enc.EncodeInt64(sense.ID); err != nil {
		return err
	}
	if err := enc.EncodeInt64(sense.ActorID); err != nil {
		return err
	}
	if err := enc.EncodeNil(); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(sense.CoolTime)); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(sense.SenseType)); err != nil {
		return err
	}
	if err := enc.Encode(sense.AcquirableLights); err != nil {
		return err
	}
	if err := encodeEffects(enc, sense.PreSenseEffect); err != nil {
		return err
	}
	if err := encodeSenseEffect(enc, sense.SenseEffect); err != nil {
		return err
	}
	if err := encodeEffects(enc, sense.PosterEffect); err != nil {
		return err
	}
	if err := encodeEffects(enc, sense.AccessoryEffect); err != nil {
		return err
	}
	if err := enc.Encode(sense.CombinationSenseID); err != nil {
		return err
	}
	if err := enc.EncodeInt64(sense.SenseMasterID); err != nil {
		return err
	}
	return enc.EncodeInt64(sense.OriginalActorID)
}

func encodeSenseEffect(enc *msgpack.Encoder, effect liveSenseEffect) error {
	if err := enc.EncodeArrayLen(4); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(effect.ScoreFactor)); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(effect.Principal)); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(effect.BranchCondition)); err != nil {
		return err
	}
	return encodeEffectBranches(enc, effect.EffectBranches)
}

func encodeStarAct(enc *msgpack.Encoder, starAct liveStarAct) error {
	if err := enc.EncodeArrayLen(4); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(starAct.ScoreFactor)); err != nil {
		return err
	}
	if err := enc.EncodeInt64(starAct.ActorID); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(starAct.BranchCondition)); err != nil {
		return err
	}
	return encodeEffectBranches(enc, starAct.EffectBranches)
}

func encodeEffectBranches(enc *msgpack.Encoder, branches []liveEffectBranch) error {
	if err := enc.EncodeArrayLen(len(branches)); err != nil {
		return err
	}
	for _, branch := range branches {
		if err := enc.EncodeArrayLen(3); err != nil {
			return err
		}
		if branch.ConditionValue == nil {
			if err := enc.EncodeNil(); err != nil {
				return err
			}
		} else if err := enc.EncodeInt64(*branch.ConditionValue); err != nil {
			return err
		}
		if branch.JudgeTypes == nil {
			if err := enc.EncodeNil(); err != nil {
				return err
			}
		} else if err := enc.EncodeInt(int64(*branch.JudgeTypes)); err != nil {
			return err
		}
		if err := encodeEffects(enc, branch.SenseEffects); err != nil {
			return err
		}
	}
	return nil
}

func encodeEffects(enc *msgpack.Encoder, effects []liveEffect) error {
	if err := enc.EncodeArrayLen(len(effects)); err != nil {
		return err
	}
	for _, effect := range effects {
		if err := enc.EncodeArrayLen(6); err != nil {
			return err
		}
		if err := enc.Encode(effect.Order); err != nil {
			return err
		}
		if err := enc.Encode(effect.MasterID); err != nil {
			return err
		}
		if err := enc.Encode(effect.EffectTypes); err != nil {
			return err
		}
		if err := encodeEffectTriggers(enc, effect.Triggers); err != nil {
			return err
		}
		if err := encodeEffectTargets(enc, effect.Targets); err != nil {
			return err
		}
		if err := enc.Encode(effect.Duration); err != nil {
			return err
		}
	}
	return nil
}

func encodeEffectTriggers(enc *msgpack.Encoder, triggers []liveEffectTrigger) error {
	if err := enc.EncodeArrayLen(len(triggers)); err != nil {
		return err
	}
	for _, trigger := range triggers {
		if err := enc.EncodeArrayLen(2); err != nil {
			return err
		}
		if err := enc.EncodeInt(int64(trigger.Type)); err != nil {
			return err
		}
		if err := enc.EncodeInt64(trigger.Value); err != nil {
			return err
		}
	}
	return nil
}

func encodeEffectTargets(enc *msgpack.Encoder, targets []liveEffectTargetValue) error {
	if err := enc.EncodeArrayLen(len(targets)); err != nil {
		return err
	}
	for _, target := range targets {
		if err := enc.EncodeArrayLen(2); err != nil {
			return err
		}
		if target.TargetActorID == nil {
			if err := enc.EncodeNil(); err != nil {
				return err
			}
		} else if err := enc.EncodeInt64(*target.TargetActorID); err != nil {
			return err
		}
		if err := enc.EncodeFloat64(target.Value); err != nil {
			return err
		}
	}
	return nil
}

func (s liveStatus) total() int {
	return s.Concentration + s.Expression + s.Vocal
}

func nullableInt64Local(value any) *int64 {
	if value == nil {
		return nil
	}
	v := int64ValueLocal(value)
	if v == 0 {
		return nil
	}
	return &v
}

func int64ValueLocal(value any) int64 {
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
		return int64(v)
	case uint:
		return int64(v)
	default:
		return 0
	}
}

func boolValueLocal(value any) bool {
	v, _ := value.(bool)
	return v
}
