package userdata

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	"ymst_server/apiwire"
)

const (
	SeedModeGenerated    = "generated"
	SeedModeCapturedJSON = "captured_json"
	SeedModeRawFixture   = "raw_fixture"
)

type GeneratedUserOptions struct {
	PlayerRank     int64
	CurrentStamina int64
	Coin           int64
	FreeJewel      int64
	PaidJewel      int64
	TutorialStatus int64
}

type BuildInitialStateOptions struct {
	UserKey                string
	DisplayName            string
	Now                    time.Time
	Generated              GeneratedUserOptions
	ProfileLevel           string
	IncludeUnions          []int64
	ExcludeUnions          []int64
	TemplateRegistry       *UserDataTemplateRegistry
	CharacterBaseMasterIDs []int64
	CharacterMasterIDs     []int64
	LiveMasterIDs          []int64
	MusicMasterIDs         []int64
	EpisodeMasterIDs       []int64
}

const (
	ProfileMinimal = "minimal"
	ProfileHome    = "home"
	ProfileFull    = "full"
)

type capturedUserData struct {
	Payload []any `json:"payload"`
}

func DefaultGeneratedUserOptions() GeneratedUserOptions {
	return GeneratedUserOptions{PlayerRank: 1, CurrentStamina: 100, Coin: 100000, FreeJewel: 3000, PaidJewel: 0, TutorialStatus: 99}
}

func BuildGeneratedInitialState(ctx context.Context, opts BuildInitialStateOptions) ([]any, error) {
	_ = ctx
	now := opts.Now
	if now.IsZero() {
		now = time.Now().UTC()
	}
	gen := opts.Generated
	if gen.PlayerRank <= 0 {
		gen = DefaultGeneratedUserOptions()
	}
	name := strings.TrimSpace(opts.DisplayName)
	if name == "" {
		name = "Player"
	}
	userKey := strings.TrimSpace(opts.UserKey)
	if userKey == "" {
		userKey = "default"
	}

	profileLevel := normalizeProfileLevel(opts.ProfileLevel)
	allowFn := makeUnionFilter(profileLevel, opts.IncludeUnions, opts.ExcludeUnions, opts.TemplateRegistry)

	characterBaseMasterIDs := normalizeInt64Slice(opts.CharacterBaseMasterIDs, []int64{1, 2, 3, 4, 5})
	characterMasterIDs := normalizeInt64Slice(opts.CharacterMasterIDs, []int64{1, 2, 3, 4, 5})
	liveMasterIDs := normalizeInt64Slice(opts.LiveMasterIDs, []int64{1001, 1002, 1003, 1004, 1005})
	musicMasterIDs := normalizeInt64Slice(opts.MusicMasterIDs, []int64{1001, 1002, 1003, 1004, 1005})
	episodeMasterIDs := normalizeInt64Slice(opts.EpisodeMasterIDs, []int64{1, 2, 3})

	payload := make([]any, 0, 64)
	add := func(unionID int64, data []any) {
		if allowFn(unionID) {
			payload = append(payload, apiwire.DataObject(unionID, data))
		}
	}

	addWithTemplate := func(unionID int64, defaultData []any) {
		if !allowFn(unionID) {
			return
		}
		if opts.TemplateRegistry != nil && opts.TemplateRegistry.HasTemplate(unionID) {
			if template, ok := opts.TemplateRegistry.GetTemplate(unionID); ok {
				payload = append(payload, apiwire.DataObject(unionID, template))
				return
			}
		}
		payload = append(payload, apiwire.DataObject(unionID, defaultData))
	}

	add(UnionUser, []any{int64(1), gen.PlayerRank, int64(0), gen.CurrentStamina, now.Add(24 * time.Hour), gen.PaidJewel, gen.FreeJewel, gen.Coin, int64(999), int64(0), nil, nil, now, hashUserID(userKey), int64(0), gen.TutorialStatus, int64(0), now, false, false})

	mainCharacterID := int64(1)
	mainCharacterMasterID := characterMasterIDs[0]
	mainCharacterBaseMasterID := characterBaseMasterIDs[0]
	add(UnionUserProfile, apiwire.ArrayByKey(22, map[int]any{0: int64(1), 1: name, 2: "", 3: mainCharacterID, 4: nil, 5: int64(1), 6: nil, 7: nil, 8: nil, 9: float64(0), 10: nil, 11: true, 12: int64(0), 13: int64(0), 14: true, 15: nil, 16: mainCharacterMasterID, 17: false, 18: true, 19: int64(1), 20: int64(1), 21: int64(1)}))
	add(UnionUserPreference, []any{int64(1), int64(1), nil})
	add(UnionHomeDisplayPreference, apiwire.ArrayByKey(16, map[int]any{0: int64(1), 1: mainCharacterBaseMasterID, 2: nil, 3: mainCharacterBaseMasterID, 4: mainCharacterBaseMasterID, 5: mainCharacterBaseMasterID, 6: nil, 7: nil, 8: nil, 9: nil, 10: nil, 11: mainCharacterMasterID, 12: false, 13: int64(0), 14: nil, 15: nil}))

	for i := 0; i < len(characterBaseMasterIDs); i++ {
		id := int64(i + 1)
		add(UnionCharacterBase, []any{id, characterBaseMasterIDs[i], int64(1), int64(0), nil, int64(1), id, false})
	}
	for i := 0; i < len(characterMasterIDs); i++ {
		id := int64(i + 1)
		add(UnionCharacter, []any{id, characterMasterIDs[i], int64(1), int64(0), int64(0), int64(0), id, int64(1), int64(0), int64(0), false, nil, int64(0), int64(0), false})
	}
	add(UnionParty, []any{int64(1), int64(1), "Main", int64(1)})
	for i := 0; i < len(characterMasterIDs); i++ {
		pos := int64(i + 1)
		add(UnionPartySlot, []any{pos, int64(1), pos, pos, nil, nil, int64(0)})
	}

	for i := 0; i < len(liveMasterIDs); i++ {
		id := int64(i + 1)
		add(UnionLive, []any{int64(24000000) + id, liveMasterIDs[i], int64(0), nil, float64(0), float64(0), int64(0), int64(1), int64(0)})
		add(UnionMusic, []any{int64(25000000) + id, nil, musicMasterIDs[i], nil, nil, true, nil, int64(0), int64(1), true})
		add(UnionMusicBookmark, []any{musicMasterIDs[i], int64(0)})
	}
	add(UnionLiveAchievement, []any{int64(1), int64(0), int64(0)})
	add(UnionLiveDropLimit, []any{int64(1), int64(0), int64(0)})

	add(UnionViewedShop, []any{int64(18000001), nil, now, int64(1)})
	add(UnionGameHint, []any{int64(17000001), int64(1), false})
	add(UnionUserBonus, []any{int64(1), float64(0), float64(0)})
	add(UnionCurrency, []any{int64(1), gen.Coin, gen.FreeJewel, gen.PaidJewel})
	add(UnionRestriction, []any{int64(1), nil, nil})
	add(UnionAlbumPreset, []any{int64(1), "Preset 1", int64(1)})
	add(UnionFriendInvitation, []any{int64(1), hashInvitationCode(userKey), false})
	add(UnionFriendInvitationMission, []any{int64(1), int64(1), int64(1), int64(0), false, false})
	addWithTemplate(UnionAuditionClear, []any{int64(1), int64(1), int64(0), int64(1), hashUserID(userKey), int64(0)})
	addWithTemplate(UnionSpRate, []any{int64(1), int64(0), int64(0), int64(0)})
	for _, id := range episodeMasterIDs {
		addWithTemplate(UnionEpisode, []any{id, false})
	}
	addWithTemplate(UnionTripleCastParty, []any{int64(1), int64(1), "TripleCast", int64(1)})
	for i := 0; i < len(liveMasterIDs); i++ {
		pos := int64(i + 1)
		addWithTemplate(UnionTripleCastPartySlot, []any{pos, int64(1), pos, pos, nil, nil, int64(0)})
	}

	return payload, nil
}

func normalizeProfileLevel(level string) string {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case ProfileHome, ProfileFull, ProfileTemplateMinimal:
		return strings.ToLower(strings.TrimSpace(level))
	default:
		return ProfileMinimal
	}
}

func makeUnionFilter(profileLevel string, includeUnions, excludeUnions []int64, templates *UserDataTemplateRegistry) func(int64) bool {
	includeSet := toSet(includeUnions)
	excludeSet := toSet(excludeUnions)

	hasExplicitFilter := len(includeSet) > 0 || len(excludeSet) > 0

	return func(unionID int64) bool {
		if hasExplicitFilter {
			if len(includeSet) > 0 && !includeSet[unionID] {
				return false
			}
			if len(excludeSet) > 0 && excludeSet[unionID] {
				return false
			}
			return true
		}

		return profileAllowsUnion(profileLevel, unionID, templates)
	}
}

func profileAllowsUnion(profileLevel string, unionID int64, templates *UserDataTemplateRegistry) bool {
	switch profileLevel {
	case ProfileFull:
		return true
	case ProfileHome:
		return profileAllowsUnionMinimal(unionID) || homeOnlyUnionSet[unionID]
	default:
		return profileAllowsUnionMinimal(unionID)
	}
}

func profileAllowsUnionMinimal(unionID int64) bool {
	return minimalUnionSet[unionID]
}

var minimalUnionSet = setOf(
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
)

var homeOnlyUnionSet = setOf(
	UnionLiveAchievement,
	UnionLiveDropLimit,
	UnionMusicBookmark,
	UnionRestriction,
	UnionAlbumPreset,
	UnionFriendInvitation,
	UnionFriendInvitationMission,
)

func toSet(ids []int64) map[int64]bool {
	s := make(map[int64]bool, len(ids))
	for _, id := range ids {
		s[id] = true
	}
	return s
}

func setOf(ids ...int64) map[int64]bool {
	return toSet(ids)
}

func normalizeInt64Slice(vals []int64, defaults []int64) []int64 {
	if len(vals) > 0 {
		return vals
	}
	return defaults
}

func BuildCapturedInitialState(ctx context.Context, seedPath string) ([]any, error) {
	_ = ctx
	if strings.TrimSpace(seedPath) == "" {
		matches, err := filepath.Glob("references/captured/*api_data_user*.api.json")
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, os.ErrNotExist
		}
		seedPath = matches[0]
	}
	raw, err := os.ReadFile(seedPath)
	if err != nil {
		return nil, err
	}
	var decoded capturedUserData
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return nil, err
	}
	return NormalizeCapturedPayload(decoded.Payload), nil
}

func NormalizeJSONArray(in []any) []any {
	out := make([]any, len(in))
	for i, value := range in {
		out[i] = NormalizeJSONValue(value)
	}
	return out
}

func NormalizeCapturedPayload(in []any) []any {
	out := make([]any, 0, len(in))
	for _, value := range in {
		if wrapped, ok := value.(map[string]any); ok {
			if unwrapped, exists := wrapped["value"]; exists {
				value = unwrapped
			}
		}
		out = append(out, NormalizeJSONValue(value))
	}
	return out
}

func NormalizeJSONValue(value any) any {
	switch v := value.(type) {
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = NormalizeJSONValue(item)
		}
		return out
	case map[string]any:
		if sec, ok := v["__timestamp_seconds__"]; ok {
			return time.Unix(Int64Value(sec), Int64Value(v["__timestamp_nanoseconds__"])).UTC()
		}
		out := make(map[string]any, len(v))
		for key, item := range v {
			out[key] = NormalizeJSONValue(item)
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

func hashUserID(userKey string) string {
	sum := sha1.Sum([]byte(userKey))
	return hex.EncodeToString(sum[:])
}

func hashInvitationCode(userKey string) string {
	sum := sha1.Sum([]byte("invite:" + userKey))
	return strings.ToUpper(hex.EncodeToString(sum[:])[:10])
}
