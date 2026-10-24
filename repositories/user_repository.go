package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"

	"ymst_server/api"
	"ymst_server/apiwire"
	"ymst_server/core/database"
	"ymst_server/userdata"
)

const (
	UnionUser                    int64 = 0
	UnionUserProfile             int64 = 1
	UnionUserPreference          int64 = 2
	UnionHomeDisplayPreference   int64 = 3
	UnionCharacter               int64 = 4
	UnionCharacterBase           int64 = 5
	UnionParty                   int64 = 6
	UnionPartySlot               int64 = 7
	UnionLive                    int64 = 24
	UnionMusic                   int64 = 25
	UnionViewedShop              int64 = 65
	UnionGameHint                int64 = 66
	UnionUserBonus               int64 = 67
	UnionAuditionClear           int64 = 90
	UnionSpRate                  int64 = 93
	UnionEpisode                 int64 = 95
	UnionCurrency                int64 = 128
	UnionLiveAchievement         int64 = 130
	UnionMusicBookmark           int64 = 146
	UnionLiveDropLimit           int64 = 147
	UnionRestriction             int64 = 148
	UnionTripleCastParty         int64 = 166
	UnionTripleCastPartySlot     int64 = 167
	UnionAlbumPreset             int64 = 174
	UnionFriendInvitation        int64 = 179
	UnionFriendInvitationMission int64 = 180
)

const defaultGameUserKey = "default"
const minimalHomeSeedMode = "minimal_home"

var tutorialRelatedPageCategories = []int64{19, 21, 29, 34}
var minimalHomeUnionIDs = []int64{
	UnionUser,
	UnionUserProfile,
	UnionUserPreference,
	UnionHomeDisplayPreference,
	UnionCharacter,
	UnionCharacterBase,
	UnionParty,
	UnionPartySlot,
	UnionLive,
	UnionMusic,
	UnionViewedShop,
	UnionGameHint,
	UnionUserBonus,
	UnionCurrency,
	UnionRestriction,
	UnionAlbumPreset,
	UnionFriendInvitation,
	UnionFriendInvitationMission,
}

type UserRepository struct {
	db               *gorm.DB
	userKey          string
	seedPath         string
	seedMode         string
	profileLevel     string
	includeUnions    []int64
	excludeUnions    []int64
	templates        *userdata.UserDataTemplateRegistry
	generated        userdata.GeneratedUserOptions
	characterBaseIDs []int64
	characterIDs     []int64
	liveIDs          []int64
	musicIDs         []int64
	episodeIDs       []int64
}

type UserDataStats struct {
	ObjectCount        int
	GameHintCount      int
	CharacterCount     int
	CharacterBaseCount int
	PartyCount         int
	PartySlotCount     int
	Coin               int
	TutorialStatus     int
	NilSlotCount       int
	SeedMode           string
	UnionCounts        map[int64]int
}

type capturedUserData struct {
	Payload []any `json:"payload"`
}

type indexedObject struct {
	Index  int
	Object any
}

func NewUserRepository(db *gorm.DB, seedPath string) *UserRepository {
	return &UserRepository{db: db, userKey: defaultGameUserKey, seedPath: seedPath, seedMode: userdata.SeedModeGenerated, profileLevel: userdata.ProfileTemplateMinimal, generated: userdata.DefaultGeneratedUserOptions()}
}

func (r *UserRepository) WithSeedMode(seedMode string) *UserRepository {
	if strings.TrimSpace(seedMode) == "" {
		seedMode = userdata.SeedModeGenerated
	}
	return &UserRepository{db: r.db, userKey: r.userKey, seedPath: r.seedPath, seedMode: seedMode, profileLevel: r.profileLevel, includeUnions: r.includeUnions, excludeUnions: r.excludeUnions, templates: r.templates, generated: r.generated, characterBaseIDs: r.characterBaseIDs, characterIDs: r.characterIDs, liveIDs: r.liveIDs, musicIDs: r.musicIDs, episodeIDs: r.episodeIDs}
}

func (r *UserRepository) WithProfileLevel(level string) *UserRepository {
	return &UserRepository{db: r.db, userKey: r.userKey, seedPath: r.seedPath, seedMode: r.seedMode, profileLevel: level, includeUnions: r.includeUnions, excludeUnions: r.excludeUnions, templates: r.templates, generated: r.generated, characterBaseIDs: r.characterBaseIDs, characterIDs: r.characterIDs, liveIDs: r.liveIDs, musicIDs: r.musicIDs, episodeIDs: r.episodeIDs}
}

func (r *UserRepository) WithIncludeUnions(ids []int64) *UserRepository {
	return &UserRepository{db: r.db, userKey: r.userKey, seedPath: r.seedPath, seedMode: r.seedMode, profileLevel: r.profileLevel, includeUnions: ids, excludeUnions: r.excludeUnions, templates: r.templates, generated: r.generated, characterBaseIDs: r.characterBaseIDs, characterIDs: r.characterIDs, liveIDs: r.liveIDs, musicIDs: r.musicIDs, episodeIDs: r.episodeIDs}
}

func (r *UserRepository) WithExcludeUnions(ids []int64) *UserRepository {
	return &UserRepository{db: r.db, userKey: r.userKey, seedPath: r.seedPath, seedMode: r.seedMode, profileLevel: r.profileLevel, includeUnions: r.includeUnions, excludeUnions: ids, templates: r.templates, generated: r.generated, characterBaseIDs: r.characterBaseIDs, characterIDs: r.characterIDs, liveIDs: r.liveIDs, musicIDs: r.musicIDs, episodeIDs: r.episodeIDs}
}

func (r *UserRepository) WithTemplates(t *userdata.UserDataTemplateRegistry) *UserRepository {
	return &UserRepository{db: r.db, userKey: r.userKey, seedPath: r.seedPath, seedMode: r.seedMode, profileLevel: r.profileLevel, includeUnions: r.includeUnions, excludeUnions: r.excludeUnions, templates: t, generated: r.generated, characterBaseIDs: r.characterBaseIDs, characterIDs: r.characterIDs, liveIDs: r.liveIDs, musicIDs: r.musicIDs, episodeIDs: r.episodeIDs}
}

func (r *UserRepository) WithSeedCharacterIDs(baseIDs, charIDs, liveIDs, musicIDs, episodeIDs []int64) *UserRepository {
	return &UserRepository{db: r.db, userKey: r.userKey, seedPath: r.seedPath, seedMode: r.seedMode, profileLevel: r.profileLevel, includeUnions: r.includeUnions, excludeUnions: r.excludeUnions, templates: r.templates, generated: r.generated, characterBaseIDs: baseIDs, characterIDs: charIDs, liveIDs: liveIDs, musicIDs: musicIDs, episodeIDs: episodeIDs}
}

func (r *UserRepository) WithGeneratedOptions(opts userdata.GeneratedUserOptions) *UserRepository {
	if opts.PlayerRank <= 0 {
		opts = userdata.DefaultGeneratedUserOptions()
	}
	return &UserRepository{db: r.db, userKey: r.userKey, seedPath: r.seedPath, seedMode: r.seedMode, profileLevel: r.profileLevel, includeUnions: r.includeUnions, excludeUnions: r.excludeUnions, templates: r.templates, generated: opts, characterBaseIDs: r.characterBaseIDs, characterIDs: r.characterIDs, liveIDs: r.liveIDs, musicIDs: r.musicIDs, episodeIDs: r.episodeIDs}
}

func (r *UserRepository) WithUserKey(userKey string) *UserRepository {
	if strings.TrimSpace(userKey) == "" {
		userKey = defaultGameUserKey
	}
	return &UserRepository{db: r.db, seedPath: r.seedPath, userKey: userKey, seedMode: r.seedMode, profileLevel: r.profileLevel, includeUnions: r.includeUnions, excludeUnions: r.excludeUnions, templates: r.templates, generated: r.generated, characterBaseIDs: r.characterBaseIDs, characterIDs: r.characterIDs, liveIDs: r.liveIDs, musicIDs: r.musicIDs, episodeIDs: r.episodeIDs}
}

func (r *UserRepository) UserKey() string { return r.userKey }

func (r *UserRepository) BuildGeneratedState(ctx context.Context, displayName string) ([]any, error) {
	opts := userdata.BuildInitialStateOptions{
		UserKey:                r.userKey,
		DisplayName:            displayName,
		Now:                    time.Now().UTC(),
		Generated:              r.generated,
		ProfileLevel:           r.profileLevel,
		IncludeUnions:          r.includeUnions,
		ExcludeUnions:          r.excludeUnions,
		TemplateRegistry:       r.templates,
		CharacterBaseMasterIDs: r.characterBaseIDs,
		CharacterMasterIDs:     r.characterIDs,
		LiveMasterIDs:          r.liveIDs,
		MusicMasterIDs:         r.musicIDs,
		EpisodeMasterIDs:       r.episodeIDs,
	}
	if r.profileLevel == userdata.ProfileTemplateMinimal {
		return userdata.BuildTemplateMinimalInitialState(ctx, opts)
	}
	return userdata.BuildGeneratedInitialState(ctx, opts)
}

func (r *UserRepository) EnsureSeeded() error {
	return r.EnsureState(context.Background())
}

func (r *UserRepository) EnsureSeededForUser(userKey string) error {
	return r.WithUserKey(userKey).EnsureState(context.Background())
}

func (r *UserRepository) EnsureState(ctxs ...context.Context) error {
	if r == nil || r.db == nil {
		return fmt.Errorf("database not initialized")
	}
	ctx := context.Background()
	if len(ctxs) > 0 && ctxs[0] != nil {
		ctx = ctxs[0]
	}
	seedMode := strings.TrimSpace(r.seedMode)
	if seedMode == "" {
		seedMode = userdata.SeedModeGenerated
	}
	var count int64
	if err := r.db.WithContext(ctx).Where("user_key = ?", r.userKey).Model(&database.GameRawDataObject{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		rebuild, err := r.rawStateNeedsRebuild(ctx, seedMode, count)
		if err != nil {
			return err
		}
		if !rebuild {
			return nil
		}
	}
	var payload []any
	var err error
	switch seedMode {
	case userdata.SeedModeCapturedJSON:
		payload, err = userdata.BuildCapturedInitialState(ctx, r.seedPath)
	default:
		opts := userdata.BuildInitialStateOptions{
			UserKey:                r.userKey,
			DisplayName:            "Player",
			Now:                    time.Now().UTC(),
			Generated:              r.generated,
			ProfileLevel:           r.profileLevel,
			IncludeUnions:          r.includeUnions,
			ExcludeUnions:          r.excludeUnions,
			TemplateRegistry:       r.templates,
			CharacterBaseMasterIDs: r.characterBaseIDs,
			CharacterMasterIDs:     r.characterIDs,
			LiveMasterIDs:          r.liveIDs,
			MusicMasterIDs:         r.musicIDs,
			EpisodeMasterIDs:       r.episodeIDs,
		}
		if r.profileLevel == userdata.ProfileTemplateMinimal {
			payload, err = userdata.BuildTemplateMinimalInitialState(ctx, opts)
		} else {
			payload, err = userdata.BuildGeneratedInitialState(ctx, opts)
		}
	}
	if err != nil {
		return err
	}
	return r.ReplaceRawState(ctx, payload, seedMode)
}

func (r *UserRepository) rawStateNeedsRebuild(ctx context.Context, seedMode string, objectCount int64) (bool, error) {
	var meta database.GameUserStateMeta
	err := r.db.WithContext(ctx).Where("user_key = ?", r.userKey).First(&meta).Error
	if err == nil && strings.TrimSpace(meta.SeedMode) != "" && !strings.EqualFold(meta.SeedMode, seedMode) {
		return true, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if !strings.EqualFold(seedMode, userdata.SeedModeCapturedJSON) {
		return false, nil
	}

	var malformedCount int64
	if err := r.db.WithContext(ctx).Where("user_key = ? AND union_id = ?", r.userKey, userdata.UnionMalformed).Model(&database.GameRawDataObject{}).Count(&malformedCount).Error; err != nil {
		return false, err
	}
	if malformedCount > 0 {
		return true, nil
	}
	rebuild, err := r.rawCapturedStateLostZeroDateTime(ctx)
	if err != nil {
		return false, err
	}
	if rebuild {
		return true, nil
	}

	return objectCount < 100, nil
}

func (r *UserRepository) rawCapturedStateLostZeroDateTime(ctx context.Context) (bool, error) {
	var row database.GameRawDataObject
	err := r.db.WithContext(ctx).
		Where("user_key = ? AND union_id = ? AND is_null = ?", r.userKey, UnionUser, false).
		Order("object_index ASC").
		First(&row).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	data, err := rawRowData(row)
	if err != nil {
		return false, err
	}
	return len(data) > 10 && data[10] == nil, nil
}

func (r *UserRepository) CloneSeedToUser(userKey string, patch func(payload []any) []any) error {
	repo := r.WithUserKey(userKey)
	ctx := context.Background()
	seedMode := strings.TrimSpace(repo.seedMode)
	if seedMode == "" {
		seedMode = userdata.SeedModeGenerated
	}
	var payload []any
	var err error
	switch seedMode {
	case userdata.SeedModeCapturedJSON:
		payload, err = userdata.BuildCapturedInitialState(ctx, repo.seedPath)
	default:
		opts := userdata.BuildInitialStateOptions{
			UserKey:                repo.userKey,
			DisplayName:            "Player",
			Now:                    time.Now().UTC(),
			Generated:              repo.generated,
			ProfileLevel:           repo.profileLevel,
			IncludeUnions:          repo.includeUnions,
			ExcludeUnions:          repo.excludeUnions,
			TemplateRegistry:       repo.templates,
			CharacterBaseMasterIDs: repo.characterBaseIDs,
			CharacterMasterIDs:     repo.characterIDs,
			LiveMasterIDs:          repo.liveIDs,
			MusicMasterIDs:         repo.musicIDs,
			EpisodeMasterIDs:       repo.episodeIDs,
		}
		if repo.profileLevel == userdata.ProfileTemplateMinimal {
			payload, err = userdata.BuildTemplateMinimalInitialState(ctx, opts)
		} else {
			payload, err = userdata.BuildGeneratedInitialState(ctx, opts)
		}
	}
	if err != nil {
		return err
	}
	if patch != nil {
		payload = patch(payload)
	}
	return repo.ReplaceRawState(ctx, payload, seedMode)
}

func (r *UserRepository) ensureCapturedObjectCounts() error {
	var zeroIndexUserCount int64
	if err := r.db.Model(&database.GameUserProfile{}).Where("object_index = 0").Count(&zeroIndexUserCount).Error; err != nil {
		return err
	}
	var restrictionCount int64
	if err := r.db.Model(&database.GameRestriction{}).Count(&restrictionCount).Error; err != nil {
		return err
	}
	var syntheticHintCount int64
	if err := r.db.Model(&database.GameHint{}).Where("id = ? AND page_category = ?", int64(17000033), int64(33)).Count(&syntheticHintCount).Error; err != nil {
		return err
	}
	var zeroNullableIDCount int64
	if err := r.db.Model(&database.GameUserProfile{}).Where("m_nameplate_id = 0 OR m_nameplate_detail_id = 0").Count(&zeroNullableIDCount).Error; err != nil {
		return err
	}
	var zeroHomeRefCount int64
	if err := r.db.Model(&database.GameHomeDisplayPreference{}).Where("home_character_base_master_id = 0 OR home_costume_master_id = 0").Count(&zeroHomeRefCount).Error; err != nil {
		return err
	}
	if restrictionCount >= 2 && syntheticHintCount == 0 && zeroIndexUserCount == 0 && zeroNullableIDCount == 0 && zeroHomeRefCount == 0 {
		return nil
	}
	fresh, err := loadCapturedUserDataPayload(r.seedPath)
	if err != nil {
		return err
	}
	patchSeedDefaults(fresh)
	return r.replaceState(fresh)
}

func (r *UserRepository) BuildDataObjects(ctxs ...context.Context) ([]any, error) {
	ctx := context.Background()
	if len(ctxs) > 0 && ctxs[0] != nil {
		ctx = ctxs[0]
	}
	if err := r.EnsureState(ctx); err != nil {
		return nil, err
	}
	rawObjects, ok, err := r.BuildRawDataObjects(ctx)
	if err != nil {
		return nil, err
	}
	if ok {
		return rawObjects, nil
	}
	return nil, fmt.Errorf("user raw state empty after ensure: user_key=%s", r.userKey)
}

func (r *UserRepository) EnsureMinimalHomeState(ctx context.Context) ([]any, error) {
	if r == nil || r.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	seed := r.WithSeedMode(userdata.SeedModeGenerated).WithProfileLevel(userdata.ProfileTemplateMinimal)
	defaults, err := seed.BuildGeneratedState(ctx, "Player")
	if err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		inserted := 0
		nextIndex := -1
		for i, raw := range defaults {
			unionID, data, ok := splitDataObject(raw)
			if !ok {
				continue
			}
			key := rawObjectKey(unionID, data, i)
			var existing database.GameRawDataObject
			err := tx.Where("user_key = ? AND object_key = ?", r.userKey, key).First(&existing).Error
			if err == nil {
				continue
			}
			if err != gorm.ErrRecordNotFound {
				return err
			}
			if nextIndex < 0 {
				var value int
				if err := tx.Model(&database.GameRawDataObject{}).Where("user_key = ?", r.userKey).Select("COALESCE(MAX(object_index), -1) + 1").Scan(&value).Error; err != nil {
					return err
				}
				nextIndex = value
			}
			if err := r.upsertRawObjectWithKey(tx, key, nextIndex, unionID, data); err != nil {
				return err
			}
			inserted++
			nextIndex++
		}
		if inserted == 0 {
			return nil
		}
		return r.updateStateMetaMode(tx, minimalHomeSeedMode)
	}); err != nil {
		return nil, err
	}

	return r.BuildMinimalHomeDataObjects(ctx, defaults)
}

func (r *UserRepository) BuildMinimalHomeDataObjects(ctx context.Context, defaults []any) ([]any, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	keys := make([]string, 0, len(defaults))
	for i, raw := range defaults {
		unionID, data, ok := splitDataObject(raw)
		if !ok {
			continue
		}
		keys = append(keys, rawObjectKey(unionID, data, i))
	}
	if len(keys) == 0 {
		return nil, nil
	}
	var raw []database.GameRawDataObject
	if err := r.db.WithContext(ctx).
		Where("user_key = ? AND object_key IN ? AND union_id IN ? AND is_null = ?", r.userKey, keys, minimalHomeUnionIDs, false).
		Order("object_index ASC").
		Find(&raw).Error; err != nil {
		return nil, err
	}
	out := make([]any, 0, len(raw))
	for _, row := range raw {
		data, err := rawRowData(row)
		if err != nil {
			return nil, err
		}
		out = append(out, apiwire.DataObject(row.UnionID, data))
	}
	return out, nil
}

func (r *UserRepository) BuildDataObjectsFromTypedTables() ([]any, error) {
	out := make([]indexedObject, 0, 800)
	appendObj := func(index int, obj any) {
		out = append(out, indexedObject{Index: index, Object: obj})
	}

	var users []database.GameUser
	if err := r.db.Order("object_index ASC").Find(&users).Error; err != nil {
		return nil, err
	}
	for _, v := range users {
		appendObj(v.ObjectIndex, api.DataObject(UnionUser, []any{
			v.ID, v.PlayerRank, v.CurrentRankPoint, v.CurrentStamina,
			v.MaxStaminaRestoredAt, v.PaidJewel, v.FreeJewel, v.Coin,
			v.PlayerRankLimit, v.StaminaRecoverTimesWithJewel,
			v.CircleUsageRestrictionsEndTime, nullableString(v.CircleID),
			v.GameStartAt, v.HashUserID, v.BanLevel, v.TutorialStatus,
			v.MonthlyPayment, v.SplashLastDisplayedAt, v.IsCapedPlayerRank,
			v.RequireCapedPlayerRankAnnounce,
		}))
	}

	var profiles []database.GameUserProfile
	if err := r.db.Order("object_index ASC").Find(&profiles).Error; err != nil {
		return nil, err
	}
	for _, v := range profiles {
		appendObj(v.ObjectIndex, api.DataObject(UnionUserProfile, api.ArrayByKey(22, map[int]any{
			0: v.ID, 1: v.Name, 2: v.Introduction, 3: v.MainUCharacterID,
			4: nullableInt64Ptr(v.MNameplateID), 5: v.MNameColorID, 6: nullableInt64Ptr(v.MTrophyID1),
			7: nullableInt64Ptr(v.MTrophyID2), 8: nullableInt64Ptr(v.MTrophyID3), 9: v.PlayerRate,
			11: v.IsPublicPlayerRate, 12: v.LeagueClass, 13: v.TotalSPCount,
			14: v.IsPublicAlbumMainPage, 15: nullableInt64Ptr(v.MNameplateDetailID),
			16: v.MainCharacterMasterID, 17: v.DisplayAwakeningStatus,
			18: v.IsPublicActivityLog, 19: v.NameBaseColorMasterID,
			20: v.IconFrameMasterID, 21: v.HomeSkinMasterID,
		})))
	}

	var prefs []database.GameUserPreference
	if err := r.db.Order("object_index ASC").Find(&prefs).Error; err != nil {
		return nil, err
	}
	for _, v := range prefs {
		appendObj(v.ObjectIndex, api.DataObject(UnionUserPreference, []any{v.ID, v.MultiPartyID, v.BirthDate}))
	}

	var homePrefs []database.GameHomeDisplayPreference
	if err := r.db.Order("object_index ASC").Find(&homePrefs).Error; err != nil {
		return nil, err
	}
	for _, v := range homePrefs {
		appendObj(v.ObjectIndex, api.DataObject(UnionHomeDisplayPreference, api.ArrayByKey(16, map[int]any{
			0: v.ID, 1: nullableInt64Ptr(v.HomeCharacterBaseMasterID), 3: nullableInt64Ptr(v.MemberCharacterBaseMasterID),
			4: nullableInt64Ptr(v.StoryCharacterBaseMasterID), 5: nullableInt64Ptr(v.ShopCharacterBaseMasterID),
			6: nullableInt64Ptr(v.HomeCostumeMasterID), 8: nullableInt64Ptr(v.MemberCostumeMasterID),
			9: nullableInt64Ptr(v.StoryCostumeMasterID), 10: nullableInt64Ptr(v.ShopCostumeMasterID),
			11: v.IllustCharacterMasterID, 12: v.DisplayAwakeningStatus,
			13: v.HomeCharacterDisplayType, 14: nullableInt64Ptr(v.LoginBonusCharacterBaseMasterID),
			15: nullableInt64Ptr(v.LoginBonusCostumeMasterID),
		})))
	}

	var bonuses []database.GameUserBonus
	if err := r.db.Order("object_index ASC").Find(&bonuses).Error; err != nil {
		return nil, err
	}
	for _, v := range bonuses {
		appendObj(v.ObjectIndex, api.DataObject(UnionUserBonus, []any{v.ID, v.ExperienceBonus, v.LessonStarRankBonus}))
	}

	var currencies []database.GameCurrency
	if err := r.db.Order("object_index ASC").Find(&currencies).Error; err != nil {
		return nil, err
	}
	for _, v := range currencies {
		appendObj(v.ObjectIndex, api.DataObject(UnionCurrency, []any{v.ID, v.Coin, v.FreeJewel, v.PaidJewel}))
	}

	var restrictions []database.GameRestriction
	if err := r.db.Order("object_index ASC").Find(&restrictions).Error; err != nil {
		return nil, err
	}
	for _, v := range restrictions {
		appendObj(v.ObjectIndex, api.DataObject(UnionRestriction, []any{v.RestrictionID, v.MultiLiveRestrictionFinishedAt, v.ReadMultiLiveRestrictionDialog}))
	}

	var hints []database.GameHint
	if err := r.db.Order("object_index ASC").Find(&hints).Error; err != nil {
		return nil, err
	}
	for _, v := range hints {
		appendObj(v.ObjectIndex, api.DataObject(UnionGameHint, []any{v.ID, v.PageCategory, v.HasAlreadyRead}))
	}

	var shops []database.GameViewedShop
	if err := r.db.Order("object_index ASC").Find(&shops).Error; err != nil {
		return nil, err
	}
	for _, v := range shops {
		appendObj(v.ObjectIndex, viewedShopObject(v))
	}

	var bases []database.GameCharacterBase
	if err := r.db.Order("object_index ASC").Find(&bases).Error; err != nil {
		return nil, err
	}
	for _, v := range bases {
		appendObj(v.ObjectIndex, api.DataObject(UnionCharacterBase, []any{v.ID, v.CharacterBaseMasterID, v.StarRank, v.TotalStarPoint, nullableInt64Ptr(v.CostumeMasterID), v.KeyMissionLevel, v.PortalCharacterID, v.PortalDisplayAwakeningStatus}))
	}

	var chars []database.GameCharacter
	if err := r.db.Order("object_index ASC").Find(&chars).Error; err != nil {
		return nil, err
	}
	for _, v := range chars {
		appendObj(v.ObjectIndex, api.DataObject(UnionCharacter, []any{v.ID, v.CharacterMasterID, v.Level, v.CurrentExperience, v.TalentStage, v.AwakeningPhase, v.CharacterBaseID, v.SenseLevel, v.ReadEpisodeOrder, v.ReleasedEpisodeOrder, v.DisplayAwakeningStatus, nullableInt64Ptr(v.SecondaryCharacterBaseID), v.SecondarySenseLevel, v.SelectionType, v.IsFavorite}))
	}

	var parties []database.GameParty
	if err := r.db.Order("object_index ASC").Find(&parties).Error; err != nil {
		return nil, err
	}
	for _, v := range parties {
		appendObj(v.ObjectIndex, api.DataObject(UnionParty, []any{v.ID, v.Order, v.Name, v.LeaderPosition}))
	}

	var slots []database.GamePartySlot
	if err := r.db.Order("object_index ASC").Find(&slots).Error; err != nil {
		return nil, err
	}
	for _, v := range slots {
		appendObj(v.ObjectIndex, api.DataObject(UnionPartySlot, []any{v.ID, v.PartyID, v.Position, v.CharacterID, nullableInt64Ptr(v.PosterID), nullableInt64Ptr(v.AccessoryID), v.BonusAbilityEnableFlags}))
	}

	var albums []database.GameAlbumPreset
	if err := r.db.Order("object_index ASC").Find(&albums).Error; err != nil {
		return nil, err
	}
	for _, v := range albums {
		appendObj(v.ObjectIndex, api.DataObject(UnionAlbumPreset, []any{v.ID, v.Name, v.Order}))
	}

	var invitations []database.GameFriendInvitation
	if err := r.db.Order("object_index ASC").Find(&invitations).Error; err != nil {
		return nil, err
	}
	for _, v := range invitations {
		appendObj(v.ObjectIndex, api.DataObject(UnionFriendInvitation, []any{v.ID, v.InvitationCode, v.HasInputOtherInvitationCode}))
	}

	var missions []database.GameFriendInvitationMission
	if err := r.db.Order("object_index ASC").Find(&missions).Error; err != nil {
		return nil, err
	}
	for _, v := range missions {
		appendObj(v.ObjectIndex, api.DataObject(UnionFriendInvitationMission, []any{v.ID, v.FriendInvitationMissionMasterID, v.FriendInvitationMissionStageMasterID, v.CurrentCount, v.IsCleared, v.IsRewardReceived}))
	}

	return sortIndexedObjects(out), nil
}

func (r *UserRepository) BuildRawDataObjects(ctxs ...context.Context) ([]any, bool, error) {
	ctx := context.Background()
	if len(ctxs) > 0 && ctxs[0] != nil {
		ctx = ctxs[0]
	}
	var raw []database.GameRawDataObject
	if err := r.db.WithContext(ctx).Where("user_key = ?", r.userKey).Order("object_index ASC").Find(&raw).Error; err != nil {
		return nil, false, err
	}
	if len(raw) == 0 {
		return nil, false, nil
	}
	out := make([]any, 0, len(raw))
	for _, row := range raw {
		if row.IsNull {
			out = append(out, nil)
			continue
		}
		if row.UnionID == userdata.UnionMalformed {
			var malformed any
			if strings.TrimSpace(row.DataJSON) != "" {
				if err := json.Unmarshal([]byte(row.DataJSON), &malformed); err != nil {
					return nil, false, err
				}
			}
			out = append(out, userdata.NormalizeJSONValue(malformed))
			continue
		}
		var data []any
		if strings.TrimSpace(row.DataJSON) != "" {
			if err := json.Unmarshal([]byte(row.DataJSON), &data); err != nil {
				return nil, false, err
			}
		}
		normalized, _ := userdata.NormalizeJSONValue(data).([]any)
		data = normalized
		out = append(out, apiwire.DataObject(row.UnionID, data))
	}
	return out, true, nil
}

func (r *UserRepository) Stats(ctxs ...context.Context) (UserDataStats, error) {
	ctx := context.Background()
	if len(ctxs) > 0 && ctxs[0] != nil {
		ctx = ctxs[0]
	}
	payload, err := r.BuildDataObjects(ctx)
	if err != nil {
		return UserDataStats{}, err
	}
	stats := UserDataStats{ObjectCount: len(payload), TutorialStatus: -1, UnionCounts: map[int64]int{}}
	for _, raw := range payload {
		if raw == nil {
			stats.NilSlotCount++
			continue
		}
		unionID, data, ok := splitDataObject(raw)
		if !ok {
			continue
		}
		stats.UnionCounts[unionID]++
		switch unionID {
		case UnionUser:
			stats.TutorialStatus = int(int64At(data, 15))
		case UnionGameHint:
			stats.GameHintCount++
		case UnionCharacter:
			stats.CharacterCount++
		case UnionCharacterBase:
			stats.CharacterBaseCount++
		case UnionParty:
			stats.PartyCount++
		case UnionPartySlot:
			stats.PartySlotCount++
		case UnionCurrency:
			stats.Coin = int(int64At(data, 1))
		}
	}
	var meta database.GameUserStateMeta
	if err := r.db.WithContext(ctx).Where("user_key = ?", r.userKey).First(&meta).Error; err == nil {
		stats.SeedMode = meta.SeedMode
	}
	return stats, nil
}

func gameUserObject(v database.GameUser) any {
	return api.DataObject(UnionUser, []any{
		v.ID, v.PlayerRank, v.CurrentRankPoint, v.CurrentStamina,
		v.MaxStaminaRestoredAt, v.PaidJewel, v.FreeJewel, v.Coin,
		v.PlayerRankLimit, v.StaminaRecoverTimesWithJewel,
		v.CircleUsageRestrictionsEndTime, nullableString(v.CircleID),
		v.GameStartAt, v.HashUserID, v.BanLevel, v.TutorialStatus,
		v.MonthlyPayment, v.SplashLastDisplayedAt, v.IsCapedPlayerRank,
		v.RequireCapedPlayerRankAnnounce,
	})
}

func gameHintObject(v database.GameHint) any {
	return api.DataObject(UnionGameHint, []any{v.ID, v.PageCategory, v.HasAlreadyRead})
}

func viewedShopObject(v database.GameViewedShop) any {
	return api.DataObject(UnionViewedShop, []any{v.ID, nullableInt64Ptr(v.ExchangeShopMasterID), v.LastViewedAt, v.ViewedShopCategory})
}

func nextObjectIndex(db *gorm.DB) int {
	maxIndex := 0
	for _, table := range []string{
		"game_users", "game_user_profiles", "game_user_preferences", "game_home_display_preferences",
		"game_characters", "game_character_bases", "game_parties", "game_party_slots", "game_hints",
		"game_viewed_shops", "game_user_bonus", "game_currencies", "game_restrictions",
		"game_album_presets", "game_friend_invitations", "game_friend_invitation_missions",
		"game_raw_data_objects",
	} {
		var value int
		_ = db.Table(table).Select("COALESCE(MAX(object_index), -1)").Scan(&value).Error
		if value >= maxIndex {
			maxIndex = value + 1
		}
	}
	return maxIndex
}

func (r *UserRepository) MarkGameHintsRead(categories []int64) ([]any, error) {
	if err := r.EnsureState(context.Background()); err != nil {
		return nil, err
	}
	return r.markGameHintsReadRaw(context.Background(), categories)
}

func (r *UserRepository) MarkGameHintsReadContext(ctx context.Context, categories []int64) ([]any, error) {
	if err := r.EnsureState(ctx); err != nil {
		return nil, err
	}
	return r.markGameHintsReadRaw(ctx, categories)
}

func (r *UserRepository) UpdateTutorialStatus(status int64) ([]any, error) {
	if err := r.EnsureState(context.Background()); err != nil {
		return nil, err
	}
	return r.updateTutorialStatusRaw(context.Background(), status)
}

func (r *UserRepository) UpdateTutorialStatusContext(ctx context.Context, status int64) ([]any, error) {
	if err := r.EnsureState(ctx); err != nil {
		return nil, err
	}
	return r.updateTutorialStatusRaw(ctx, status)
}

func (r *UserRepository) UpdateSplashLastDisplayedAt(t time.Time) ([]any, error) {
	if err := r.EnsureState(context.Background()); err != nil {
		return nil, err
	}
	return r.updateSplashLastDisplayedAtRaw(context.Background(), t)
}

func (r *UserRepository) UpdateSplashLastDisplayedAtContext(ctx context.Context, t time.Time) ([]any, error) {
	if err := r.EnsureState(ctx); err != nil {
		return nil, err
	}
	return r.updateSplashLastDisplayedAtRaw(ctx, t)
}

func (r *UserRepository) MarkShopViewed(category int64) ([]any, error) {
	return r.MarkShopViewedWithExchangeShopMasterID(category, nil)
}

func (r *UserRepository) MarkShopViewedWithExchangeShopMasterID(category int64, exchangeShopMasterID *int64) ([]any, error) {
	if err := r.EnsureState(context.Background()); err != nil {
		return nil, err
	}
	return r.markShopViewedRaw(context.Background(), category, exchangeShopMasterID)
}

func (r *UserRepository) MarkShopViewedContext(ctx context.Context, category int64, exchangeShopMasterID *int64) ([]any, error) {
	if err := r.EnsureState(ctx); err != nil {
		return nil, err
	}
	return r.markShopViewedRaw(ctx, category, exchangeShopMasterID)
}

func (r *UserRepository) replaceState(payload []any) error {
	return r.ReplaceRawState(context.Background(), payload, r.seedMode)
}

func (r *UserRepository) updateStateMetaMode(tx *gorm.DB, seedMode string) error {
	if strings.TrimSpace(seedMode) == "" {
		seedMode = userdata.SeedModeGenerated
	}
	now := time.Now().UTC()
	var count int64
	if err := tx.Model(&database.GameRawDataObject{}).Where("user_key = ?", r.userKey).Count(&count).Error; err != nil {
		return err
	}
	var meta database.GameUserStateMeta
	err := tx.Where("user_key = ?", r.userKey).First(&meta).Error
	if err == nil {
		return tx.Model(&meta).Updates(map[string]any{
			"seed_mode":       seedMode,
			"object_count":    int(count),
			"last_updated_at": now,
		}).Error
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	return tx.Create(&database.GameUserStateMeta{UserKey: r.userKey, SeedMode: seedMode, StateVersion: 1, ObjectCount: int(count), LastBuiltAt: now, LastUpdatedAt: now}).Error
}

func (r *UserRepository) ReplaceRawState(ctx context.Context, payload []any, seedMode string) error {
	if r == nil || r.db == nil {
		return fmt.Errorf("database not initialized")
	}
	if strings.TrimSpace(seedMode) == "" {
		seedMode = userdata.SeedModeGenerated
	}
	now := time.Now().UTC()
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_key = ?", r.userKey).Delete(&database.GameRawDataObject{}).Error; err != nil {
			return err
		}
		if err := r.SaveRawPayload(tx, payload, now); err != nil {
			return err
		}
		meta := database.GameUserStateMeta{UserKey: r.userKey, SeedMode: seedMode, StateVersion: 1, ObjectCount: len(payload), LastBuiltAt: now, LastUpdatedAt: now}
		var existing database.GameUserStateMeta
		err := tx.Where("user_key = ?", r.userKey).First(&existing).Error
		if err == nil {
			return tx.Model(&existing).Updates(map[string]any{"seed_mode": seedMode, "state_version": int64(1), "object_count": len(payload), "last_built_at": now, "last_updated_at": now}).Error
		}
		if err != gorm.ErrRecordNotFound {
			return err
		}
		return tx.Create(&meta).Error
	})
}

func (r *UserRepository) insertPayload(tx *gorm.DB, payload []any) error {
	now := time.Now()
	if err := r.SaveRawPayload(tx, payload, now); err != nil {
		return err
	}
	for i, raw := range payload {
		if raw == nil {
			continue
		}
		unionID, data, ok := splitDataObject(raw)
		if !ok {
			continue
		}
		if err := r.insertTypedObject(tx, unionID, data, i, now); err != nil {
			return err
		}
	}
	return nil
}

func (r *UserRepository) SaveRawPayload(tx *gorm.DB, payload []any, now time.Time) error {
	seenKeys := make(map[string]int, len(payload))
	for i, raw := range payload {
		if raw == nil {
			objectKey := fmt.Sprintf("null:%d", i)
			row := database.GameRawDataObject{
				UserKey:     r.userKey,
				ObjectKey:   objectKey,
				UnionID:     0,
				ObjectIndex: i,
				DataJSON:    "",
				IsNull:      true,
				CreatedAt:   now,
				UpdatedAt:   now,
			}
			if err := tx.Create(&row).Error; err != nil {
				return err
			}
			continue
		}
		unionID, data, ok := splitDataObject(raw)
		if !ok {
			dataJSON, err := json.Marshal(raw)
			if err != nil {
				return err
			}
			objectKey := fmt.Sprintf("malformed:%d", i)
			row := database.GameRawDataObject{
				UserKey:     r.userKey,
				ObjectKey:   objectKey,
				UnionID:     userdata.UnionMalformed,
				ObjectIndex: i,
				DataJSON:    string(dataJSON),
				IsNull:      false,
				CreatedAt:   now,
				UpdatedAt:   now,
			}
			if err := tx.Create(&row).Error; err != nil {
				return err
			}
			continue
		}
		dataJSON, err := marshalObjectData(data)
		if err != nil {
			return err
		}
		objectKey := rawObjectKey(unionID, data, i)
		if seenKeys[objectKey] > 0 {
			objectKey = fmt.Sprintf("%d:idx:%d", unionID, i)
		}
		seenKeys[objectKey]++
		row := database.GameRawDataObject{
			UserKey:     r.userKey,
			ObjectKey:   objectKey,
			UnionID:     unionID,
			ObjectIndex: i,
			DataJSON:    dataJSON,
			IsNull:      false,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		if err := tx.Create(&row).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *UserRepository) UpsertRawObject(tx *gorm.DB, objectIndex int, unionID int64, data []any) error {
	return r.upsertRawObjectWithKey(tx, rawObjectKey(unionID, data, objectIndex), objectIndex, unionID, data)
}

func (r *UserRepository) UpsertRawDataObjectByID(tx *gorm.DB, unionID int64, objectID int64, data []any) error {
	key := fmt.Sprintf("%d:%d", unionID, objectID)
	var existing database.GameRawDataObject
	err := tx.Where("user_key = ? AND object_key = ?", r.userKey, key).First(&existing).Error
	if err == nil {
		return r.upsertRawObjectWithKey(tx, key, existing.ObjectIndex, unionID, data)
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	return r.upsertRawObjectWithKey(tx, key, nextObjectIndex(tx), unionID, data)
}

func (r *UserRepository) upsertRawObjectWithKey(tx *gorm.DB, key string, objectIndex int, unionID int64, data []any) error {
	dataJSON, err := marshalObjectData(data)
	if err != nil {
		return err
	}
	now := time.Now()
	var existing database.GameRawDataObject
	err = tx.Where("user_key = ? AND object_key = ?", r.userKey, key).First(&existing).Error
	if err == nil {
		if err := tx.Model(&existing).Updates(map[string]any{
			"union_id":     unionID,
			"object_index": objectIndex,
			"data_json":    dataJSON,
			"is_null":      false,
			"updated_at":   now,
		}).Error; err != nil {
			return err
		}
		return r.touchStateMeta(tx, 0)
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	if err := tx.Create(&database.GameRawDataObject{UserKey: r.userKey, ObjectKey: key, UnionID: unionID, ObjectIndex: objectIndex, DataJSON: dataJSON, IsNull: false, CreatedAt: now, UpdatedAt: now}).Error; err != nil {
		return err
	}
	return r.touchStateMeta(tx, 1)
}

func (r *UserRepository) nextRawObjectIndex(ctx context.Context) (int, error) {
	var value int
	err := r.db.WithContext(ctx).Model(&database.GameRawDataObject{}).Where("user_key = ?", r.userKey).Select("COALESCE(MAX(object_index), -1) + 1").Scan(&value).Error
	return value, err
}

func rawObjectKey(unionID int64, data []any, objectIndex int) string {
	return userdata.MakeRawObjectKey(unionID, data, objectIndex)
}

func (r *UserRepository) updateTutorialStatusRaw(ctx context.Context, status int64) ([]any, error) {
	if status <= 0 {
		status = 99
	}
	now := time.Now()
	return r.Mutator().UpdateObjectByUnion(ctx, UnionUser, func(data []any) bool { return true }, func(data []any) ([]any, error) {
		data = ensureArraySize(data, 20)
		data[15] = status
		data[17] = now
		return data, nil
	})
}

func (r *UserRepository) updateSplashLastDisplayedAtRaw(ctx context.Context, t time.Time) ([]any, error) {
	if t.IsZero() {
		t = time.Now()
	}
	return r.Mutator().UpdateObjectByUnion(ctx, UnionUser, func(data []any) bool { return true }, func(data []any) ([]any, error) {
		data = ensureArraySize(data, 20)
		data[17] = t
		return data, nil
	})
}

func (r *UserRepository) markGameHintsReadRaw(ctx context.Context, categories []int64) ([]any, error) {
	all := append([]int64{}, categories...)
	if intersects(categories, tutorialRelatedPageCategories) {
		all = append(all, tutorialRelatedPageCategories...)
	}
	updates := make([]any, 0, len(all))
	seenCategories := map[int64]bool{}
	for _, category := range all {
		if category <= 0 || seenCategories[category] {
			continue
		}
		seenCategories[category] = true
		updated, err := r.Mutator().UpdateObjectByUnion(ctx, UnionGameHint, func(candidate []any) bool { return int64At(candidate, 1) == category }, func(candidate []any) ([]any, error) {
			candidate = ensureArraySize(candidate, 3)
			candidate[2] = true
			return candidate, nil
		})
		if err == gorm.ErrRecordNotFound {
			data := []any{int64(17000000) + category, category, true}
			inserted, insertErr := r.Mutator().InsertObject(ctx, UnionGameHint, data)
			if insertErr != nil {
				return nil, insertErr
			}
			updated = []any{inserted}
		} else if err != nil {
			return nil, err
		}
		updates = append(updates, updated...)
	}
	return updates, nil
}

func (r *UserRepository) markShopViewedRaw(ctx context.Context, category int64, exchangeShopMasterID *int64) ([]any, error) {
	if category <= 0 {
		category = 1
	}
	now := time.Now()
	updated, err := r.Mutator().UpdateObjectByUnion(ctx, UnionViewedShop, func(candidate []any) bool { return int64At(candidate, 3) == category }, func(candidate []any) ([]any, error) {
		candidate = ensureArraySize(candidate, 4)
		candidate[1] = nullableInt64Ptr(exchangeShopMasterID)
		candidate[2] = now
		return candidate, nil
	})
	if err == gorm.ErrRecordNotFound {
		inserted, insertErr := r.Mutator().InsertObject(ctx, UnionViewedShop, []any{int64(18000000) + category, nullableInt64Ptr(exchangeShopMasterID), now, category})
		if insertErr != nil {
			return nil, insertErr
		}
		return []any{inserted}, nil
	}
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (r *UserRepository) findRawObjectByUnion(tx *gorm.DB, unionID int64) (database.GameRawDataObject, []any, error) {
	var row database.GameRawDataObject
	err := tx.Where("user_key = ? AND union_id = ? AND is_null = ?", r.userKey, unionID, false).Order("object_index ASC").First(&row).Error
	if err != nil {
		return row, nil, err
	}
	data, err := rawRowData(row)
	return row, data, err
}

func (r *UserRepository) findRawGameHintByCategory(tx *gorm.DB, category int64) (database.GameRawDataObject, []any, bool, error) {
	return r.findRawByUnionPredicate(tx, UnionGameHint, func(data []any) bool { return int64At(data, 1) == category })
}

func (r *UserRepository) findRawViewedShopByCategory(tx *gorm.DB, category int64) (database.GameRawDataObject, []any, bool, error) {
	return r.findRawByUnionPredicate(tx, UnionViewedShop, func(data []any) bool { return int64At(data, 3) == category })
}

func (r *UserRepository) findRawByUnionPredicate(tx *gorm.DB, unionID int64, match func([]any) bool) (database.GameRawDataObject, []any, bool, error) {
	var rows []database.GameRawDataObject
	if err := tx.Where("user_key = ? AND union_id = ? AND is_null = ?", r.userKey, unionID, false).Order("object_index ASC").Find(&rows).Error; err != nil {
		return database.GameRawDataObject{}, nil, false, err
	}
	for _, row := range rows {
		data, err := rawRowData(row)
		if err != nil {
			return row, nil, false, err
		}
		if match(data) {
			return row, data, true, nil
		}
	}
	return database.GameRawDataObject{}, nil, false, nil
}

func rawRowData(row database.GameRawDataObject) ([]any, error) {
	var data []any
	if strings.TrimSpace(row.DataJSON) != "" {
		if err := json.Unmarshal([]byte(row.DataJSON), &data); err != nil {
			return nil, err
		}
	}
	normalized, _ := normalizeJSONValue(data).([]any)
	return normalized, nil
}

func ensureArraySize(data []any, size int) []any {
	if len(data) >= size {
		return data
	}
	out := make([]any, size)
	copy(out, data)
	return out
}

func (r *UserRepository) insertTypedObject(tx *gorm.DB, unionID int64, data []any, idx int, now time.Time) error {
	switch unionID {
	case UnionUser:
		return tx.Create(&database.GameUser{ID: int64At(data, 0), PlayerRank: int(int64At(data, 1)), CurrentRankPoint: int(int64At(data, 2)), CurrentStamina: int(int64At(data, 3)), MaxStaminaRestoredAt: timeAt(data, 4), PaidJewel: int(int64At(data, 5)), FreeJewel: int(int64At(data, 6)), Coin: int(int64At(data, 7)), PlayerRankLimit: int(int64At(data, 8)), StaminaRecoverTimesWithJewel: int(int64At(data, 9)), CircleUsageRestrictionsEndTime: timeAt(data, 10), CircleID: stringAt(data, 11), GameStartAt: timeAt(data, 12), HashUserID: stringAt(data, 13), BanLevel: int(int64At(data, 14)), TutorialStatus: int(int64At(data, 15)), MonthlyPayment: int(int64At(data, 16)), SplashLastDisplayedAt: timeAt(data, 17), IsCapedPlayerRank: boolAt(data, 18), RequireCapedPlayerRankAnnounce: boolAt(data, 19), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionUserProfile:
		return tx.Create(&database.GameUserProfile{ID: int64At(data, 0), Name: stringAt(data, 1), Introduction: stringAt(data, 2), MainUCharacterID: int64At(data, 3), MNameplateID: ptrFromAny(valueAt(data, 4)), MNameColorID: int64At(data, 5), MTrophyID1: ptrFromAny(valueAt(data, 6)), MTrophyID2: ptrFromAny(valueAt(data, 7)), MTrophyID3: ptrFromAny(valueAt(data, 8)), PlayerRate: float64Value(valueAt(data, 9)), IsPublicPlayerRate: boolAt(data, 11), LeagueClass: int(int64At(data, 12)), TotalSPCount: int(int64At(data, 13)), IsPublicAlbumMainPage: boolAt(data, 14), MNameplateDetailID: ptrFromAny(valueAt(data, 15)), MainCharacterMasterID: int64At(data, 16), DisplayAwakeningStatus: boolAt(data, 17), IsPublicActivityLog: boolAt(data, 18), NameBaseColorMasterID: int64At(data, 19), IconFrameMasterID: int64At(data, 20), HomeSkinMasterID: int64At(data, 21), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionUserPreference:
		return tx.Create(&database.GameUserPreference{ID: int64At(data, 0), MultiPartyID: int64At(data, 1), BirthDate: timePtrFromAny(valueAt(data, 2)), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionHomeDisplayPreference:
		return tx.Create(&database.GameHomeDisplayPreference{ID: int64At(data, 0), HomeCharacterBaseMasterID: ptrFromAny(valueAt(data, 1)), MemberCharacterBaseMasterID: ptrFromAny(valueAt(data, 3)), StoryCharacterBaseMasterID: ptrFromAny(valueAt(data, 4)), ShopCharacterBaseMasterID: ptrFromAny(valueAt(data, 5)), HomeCostumeMasterID: ptrFromAny(valueAt(data, 6)), MemberCostumeMasterID: ptrFromAny(valueAt(data, 8)), StoryCostumeMasterID: ptrFromAny(valueAt(data, 9)), ShopCostumeMasterID: ptrFromAny(valueAt(data, 10)), IllustCharacterMasterID: int64At(data, 11), DisplayAwakeningStatus: boolAt(data, 12), HomeCharacterDisplayType: int(int64At(data, 13)), LoginBonusCharacterBaseMasterID: ptrFromAny(valueAt(data, 14)), LoginBonusCostumeMasterID: ptrFromAny(valueAt(data, 15)), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionCharacter:
		return tx.Create(&database.GameCharacter{ID: int64At(data, 0), CharacterMasterID: int64At(data, 1), Level: int(int64At(data, 2)), CurrentExperience: int(int64At(data, 3)), TalentStage: int(int64At(data, 4)), AwakeningPhase: int(int64At(data, 5)), CharacterBaseID: int64At(data, 6), SenseLevel: int(int64At(data, 7)), ReadEpisodeOrder: int(int64At(data, 8)), ReleasedEpisodeOrder: int(int64At(data, 9)), DisplayAwakeningStatus: boolAt(data, 10), SecondaryCharacterBaseID: ptrFromAny(valueAt(data, 11)), SecondarySenseLevel: int(int64At(data, 12)), SelectionType: int(int64At(data, 13)), IsFavorite: boolAt(data, 14), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionCharacterBase:
		return tx.Create(&database.GameCharacterBase{ID: int64At(data, 0), CharacterBaseMasterID: int64At(data, 1), StarRank: int(int64At(data, 2)), TotalStarPoint: int(int64At(data, 3)), CostumeMasterID: ptrFromAny(valueAt(data, 4)), KeyMissionLevel: int(int64At(data, 5)), PortalCharacterID: int64At(data, 6), PortalDisplayAwakeningStatus: boolAt(data, 7), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionParty:
		return tx.Create(&database.GameParty{ID: int64At(data, 0), Order: int(int64At(data, 1)), Name: stringAt(data, 2), LeaderPosition: int(int64At(data, 3)), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionPartySlot:
		return tx.Create(&database.GamePartySlot{ID: int64At(data, 0), PartyID: int64At(data, 1), Position: int(int64At(data, 2)), CharacterID: int64At(data, 3), PosterID: ptrFromAny(valueAt(data, 4)), AccessoryID: ptrFromAny(valueAt(data, 5)), BonusAbilityEnableFlags: int(int64At(data, 6)), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionGameHint:
		return tx.Create(&database.GameHint{ID: int64At(data, 0), PageCategory: int64At(data, 1), HasAlreadyRead: boolAt(data, 2), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionViewedShop:
		return tx.Create(&database.GameViewedShop{ID: int64At(data, 0), ExchangeShopMasterID: ptrFromAny(valueAt(data, 1)), LastViewedAt: timeAt(data, 2), ViewedShopCategory: int64At(data, 3), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionUserBonus:
		return tx.Create(&database.GameUserBonus{ID: int64At(data, 0), ExperienceBonus: float64Value(valueAt(data, 1)), LessonStarRankBonus: float64Value(valueAt(data, 2)), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionCurrency:
		return tx.Create(&database.GameCurrency{ID: int64At(data, 0), Coin: int(int64At(data, 1)), FreeJewel: int(int64At(data, 2)), PaidJewel: int(int64At(data, 3)), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionRestriction:
		return tx.Create(&database.GameRestriction{RestrictionID: int64At(data, 0), MultiLiveRestrictionFinishedAt: timePtrFromAny(valueAt(data, 1)), ReadMultiLiveRestrictionDialog: boolPtrFromAny(valueAt(data, 2)), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionAlbumPreset:
		return tx.Create(&database.GameAlbumPreset{ID: int64At(data, 0), Name: stringAt(data, 1), Order: int(int64At(data, 2)), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionFriendInvitation:
		return tx.Create(&database.GameFriendInvitation{ID: int64At(data, 0), InvitationCode: stringAt(data, 1), HasInputOtherInvitationCode: boolAt(data, 2), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	case UnionFriendInvitationMission:
		return tx.Create(&database.GameFriendInvitationMission{ID: int64At(data, 0), FriendInvitationMissionMasterID: int64At(data, 1), FriendInvitationMissionStageMasterID: int64At(data, 2), CurrentCount: int64At(data, 3), IsCleared: boolAt(data, 4), IsRewardReceived: boolAt(data, 5), ObjectIndex: idx, CreatedAt: now, UpdatedAt: now}).Error
	default:
		return nil
	}
}

func loadCapturedUserDataPayload(path string) ([]any, error) {
	if strings.TrimSpace(path) == "" {
		matches, err := filepath.Glob("references/captured/*api_data_user*.api.json")
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, os.ErrNotExist
		}
		path = matches[0]
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var decoded capturedUserData
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return nil, err
	}
	return userdata.NormalizeCapturedPayload(decoded.Payload), nil
}

func patchSeedDefaults(payload []any) {
	patchUser(payload, time.Now())
	categories := make([]int64, 0, 32)
	for _, raw := range payload {
		unionID, data, ok := splitDataObject(raw)
		if ok && unionID == UnionGameHint {
			categories = append(categories, int64At(data, 1))
		}
	}
	patchGameHintsRead(payload, categories)
}

func patchUser(payload []any, splashAt time.Time) []any {
	if splashAt.IsZero() {
		splashAt = time.Now()
	}
	for i, raw := range payload {
		unionID, data, ok := splitDataObject(raw)
		if !ok || unionID != UnionUser {
			continue
		}
		if len(data) > 15 {
			data[15] = int64(99)
		}
		if len(data) > 17 {
			data[17] = splashAt
		}
		payload[i] = api.DataObject(UnionUser, data)
		return []any{cloneAny(payload[i])}
	}
	return nil
}

func patchGameHintsRead(payload []any, categories []int64) []any {
	updates := make([]any, 0, len(categories))
	seenCategories := map[int64]bool{}
	for _, category := range categories {
		if category <= 0 || seenCategories[category] {
			continue
		}
		seenCategories[category] = true
		found := false
		for i, raw := range payload {
			unionID, data, ok := splitDataObject(raw)
			if !ok || unionID != UnionGameHint || int64At(data, 1) != category {
				continue
			}
			if len(data) > 2 {
				data[2] = true
			}
			payload[i] = api.DataObject(UnionGameHint, data)
			updates = append(updates, cloneAny(payload[i]))
			found = true
			break
		}
		if !found {
			obj := api.DataObject(UnionGameHint, []any{int64(17000000) + category, category, true})
			payload = append(payload, obj)
			updates = append(updates, cloneAny(obj))
		}
	}
	return updates
}

func patchShopViewed(payload []any, category int64) []any {
	if category <= 0 {
		category = 1
	}
	now := time.Now()
	for i, raw := range payload {
		unionID, data, ok := splitDataObject(raw)
		if !ok || unionID != UnionViewedShop || int64At(data, 3) != category {
			continue
		}
		if len(data) > 2 {
			data[2] = now
		}
		payload[i] = api.DataObject(UnionViewedShop, data)
		return []any{cloneAny(payload[i])}
	}
	obj := api.DataObject(UnionViewedShop, []any{int64(18000000) + category, nil, now, category})
	payload = append(payload, obj)
	return []any{cloneAny(obj)}
}

func clearGameTables(tx *gorm.DB) error {
	for _, model := range []any{&database.GameUser{}, &database.GameUserProfile{}, &database.GameUserPreference{}, &database.GameHomeDisplayPreference{}, &database.GameCharacter{}, &database.GameCharacterBase{}, &database.GameParty{}, &database.GamePartySlot{}, &database.GameHint{}, &database.GameViewedShop{}, &database.GameUserBonus{}, &database.GameCurrency{}, &database.GameRestriction{}, &database.GameAlbumPreset{}, &database.GameFriendInvitation{}, &database.GameFriendInvitationMission{}, &database.GameRawDataObject{}} {
		if err := tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(model).Error; err != nil {
			return err
		}
	}
	return nil
}

func normalizeJSONArray(in []any) []any {
	out := make([]any, len(in))
	for i, value := range in {
		out[i] = normalizeJSONValue(value)
	}
	return out
}

func normalizeJSONValue(value any) any {
	switch v := value.(type) {
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = normalizeJSONValue(item)
		}
		return out
	case map[string]any:
		if sec, ok := v["__timestamp_seconds__"]; ok {
			return time.Unix(int64Value(sec), int64Value(v["__timestamp_nanoseconds__"])).UTC()
		}
		out := make(map[string]any, len(v))
		for key, item := range v {
			out[key] = normalizeJSONValue(item)
		}
		return out
	case float64:
		if v == float64(int64(v)) {
			return int64(v)
		}
		return v
	case string:
		if t, ok := parseCapturedTime(v); ok {
			return t
		}
		return v
	default:
		return v
	}
}

func parseCapturedTime(value string) (time.Time, bool) {
	layouts := []string{"2006-01-02 15:04:05.999999999", "2006-01-02 15:04:05", time.RFC3339Nano, time.RFC3339}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, value, time.Local); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

func splitDataObject(raw any) (int64, []any, bool) {
	obj, ok := raw.([]any)
	if !ok || len(obj) < 2 {
		return 0, nil, false
	}
	data, ok := obj[1].([]any)
	if !ok {
		return 0, nil, false
	}
	return int64Value(obj[0]), data, true
}

func sortIndexedObjects(items []indexedObject) []any {
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			if items[j].Index < items[i].Index {
				items[i], items[j] = items[j], items[i]
			}
		}
	}
	out := make([]any, 0, len(items))
	for _, item := range items {
		out = append(out, item.Object)
	}
	return out
}

func cloneAny(value any) any {
	switch v := value.(type) {
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = cloneAny(item)
		}
		return out
	case map[string]any:
		out := make(map[string]any, len(v))
		for key, item := range v {
			out[key] = cloneAny(item)
		}
		return out
	default:
		return value
	}
}

func nullableString(v string) any {
	if v == "" {
		return nil
	}
	return v
}

func nullableInt64Ptr(v *int64) any {
	if v == nil || *v == 0 {
		return nil
	}
	return *v
}

func valueAt(data []any, idx int) any {
	if idx < 0 || idx >= len(data) {
		return nil
	}
	return data[idx]
}

func int64At(data []any, idx int) int64 { return int64Value(valueAt(data, idx)) }

func int64Value(v any) int64 {
	switch n := v.(type) {
	case int:
		return int64(n)
	case int8:
		return int64(n)
	case int16:
		return int64(n)
	case int32:
		return int64(n)
	case int64:
		return n
	case uint:
		return int64(n)
	case uint8:
		return int64(n)
	case uint16:
		return int64(n)
	case uint32:
		return int64(n)
	case uint64:
		return int64(n)
	default:
		return 0
	}
}

func stringAt(data []any, idx int) string {
	v, _ := valueAt(data, idx).(string)
	return v
}

func boolAt(data []any, idx int) bool {
	v, _ := valueAt(data, idx).(bool)
	return v
}

func timeAt(data []any, idx int) time.Time {
	v, _ := valueAt(data, idx).(time.Time)
	return v
}

func ptrFromAny(v any) *int64 {
	if v == nil {
		return nil
	}
	n := int64Value(v)
	if n == 0 {
		return nil
	}
	return &n
}

func timePtrFromAny(v any) *time.Time {
	t, ok := v.(time.Time)
	if !ok || t.IsZero() {
		return nil
	}
	return &t
}

func boolPtrFromAny(v any) *bool {
	b, ok := v.(bool)
	if !ok {
		return nil
	}
	return &b
}

func float64Value(v any) float64 {
	switch n := v.(type) {
	case float64:
		return n
	case float32:
		return float64(n)
	case int64:
		return float64(n)
	case int:
		return float64(n)
	default:
		return 0
	}
}

func intersects(left []int64, right []int64) bool {
	seen := make(map[int64]bool, len(left))
	for _, value := range left {
		seen[value] = true
	}
	for _, value := range right {
		if seen[value] {
			return true
		}
	}
	return false
}

func marshalObjectData(data []any) (string, error) {
	normalized := normalizeForJSON(data)
	raw, err := json.Marshal(normalized)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func normalizeForJSON(value any) any {
	switch v := value.(type) {
	case time.Time:
		return map[string]any{
			"__timestamp_seconds__":     v.Unix(),
			"__timestamp_nanoseconds__": int64(v.Nanosecond()),
		}
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = normalizeForJSON(item)
		}
		return out
	default:
		return v
	}
}
