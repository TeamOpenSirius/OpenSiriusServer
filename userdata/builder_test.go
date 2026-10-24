package userdata

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"
	"time"
)

func TestGeneratedMinimalUserDataPayload(t *testing.T) {
	opts := BuildInitialStateOptions{
		UserKey:      "test_user",
		DisplayName:  "Tester",
		Now:          time.Now().UTC(),
		Generated:    DefaultGeneratedUserOptions(),
		ProfileLevel: ProfileMinimal,
	}
	payload, err := BuildGeneratedInitialState(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if len(payload) == 0 {
		t.Fatal("empty minimal payload")
	}

	counts := map[int64]int{}
	unionArrayLens := map[int64]int{}
	for _, raw := range payload {
		obj, ok := raw.([]any)
		if !ok || len(obj) < 2 {
			t.Fatalf("malformed data object: %#v", raw)
		}
		unionID := Int64Value(obj[0])
		data, ok := obj[1].([]any)
		if !ok {
			t.Fatalf("malformed data array for union %d: %#v", unionID, obj[1])
		}
		counts[unionID]++
		unionArrayLens[unionID] = len(data)
	}

	minimalRequired := []int64{UnionUser, UnionUserProfile, UnionUserPreference, UnionHomeDisplayPreference, UnionCharacter, UnionCharacterBase, UnionParty, UnionPartySlot, UnionCurrency}
	for _, uid := range minimalRequired {
		if counts[uid] == 0 {
			t.Fatalf("minimal union %d missing", uid)
		}
	}

	knownSizes := map[int64]int{
		UnionUser:                  20,
		UnionUserProfile:           22,
		UnionUserPreference:        3,
		UnionHomeDisplayPreference: 16,
		UnionCharacter:             15,
		UnionCharacterBase:         8,
		UnionParty:                 4,
		UnionPartySlot:             7,
		UnionCurrency:              4,
	}
	for unionID, wantLen := range knownSizes {
		if got, ok := unionArrayLens[unionID]; ok && got != wantLen {
			t.Fatalf("union %d arrayLen = %d want %d", unionID, got, wantLen)
		}
	}

	unexpectedInMinimal := []int64{UnionAuditionClear, UnionSpRate, UnionEpisode, UnionLiveAchievement, UnionMusicBookmark, UnionLiveDropLimit, UnionTripleCastParty, UnionTripleCastPartySlot, UnionFriendInvitationMission, UnionRestriction, UnionAlbumPreset, UnionFriendInvitation}
	for _, uid := range unexpectedInMinimal {
		if counts[uid] > 0 {
			t.Errorf("union %d should not be in minimal profile, got %d objects", uid, counts[uid])
		}
	}
}

func TestGeneratedPayloadTimeRoundTrip(t *testing.T) {
	now := time.Now().UTC()
	opts := BuildInitialStateOptions{
		UserKey:      "time_test",
		DisplayName:  "TimeTester",
		Now:          now,
		Generated:    DefaultGeneratedUserOptions(),
		ProfileLevel: ProfileMinimal,
	}
	payload, err := BuildGeneratedInitialState(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}

	for _, raw := range payload {
		unionID, data, ok := SplitDataObject(raw)
		if !ok {
			continue
		}
		for idx, field := range data {
			if _, isTime := field.(time.Time); isTime && field.(time.Time).IsZero() {
				t.Errorf("union %d field %d is zero time", unionID, idx)
			}
			if s, isString := field.(string); isString {
				if _, ok := parseCapturedTime(s); ok {
					t.Errorf("union %d field %d is time-like string %q, not time.Time", unionID, idx, s)
				}
			}
		}
	}

	normalized := normalizeForJSONTest(payload)
	raw, err := json.Marshal(normalized)
	if err != nil {
		t.Fatal(err)
	}

	var decoded []any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	restored := NormalizeJSONArray(decoded)

	for i, orig := range payload {
		if i >= len(restored) {
			t.Fatalf("restored payload shorter at index %d", i)
		}
		if orig == nil && restored[i] == nil {
			continue
		}
		origUnion, origData, _ := SplitDataObject(orig)
		restUnion, restData, _ := SplitDataObject(restored[i])
		if origUnion != restUnion {
			t.Fatalf("union mismatch at %d: %d != %d", i, origUnion, restUnion)
		}
		if len(origData) != len(restData) {
			t.Fatalf("array len mismatch at union %d index %d: %d != %d", origUnion, i, len(origData), len(restData))
		}
		for j, origField := range origData {
			restField := restData[j]
			_, origIsTime := origField.(time.Time)
			_, restIsTime := restField.(time.Time)
			if origIsTime && !restIsTime {
				t.Errorf("union %d[%d] orig is time.Time but restored is %T", origUnion, j, restField)
			}
		}
	}
}

func TestNormalizeCapturedPayloadUnwrapsExportedValue(t *testing.T) {
	payload := NormalizeCapturedPayload([]any{
		map[string]any{
			"value": []any{float64(UnionUser), []any{float64(1), "player"}},
			"Count": float64(2),
		},
		nil,
	})

	if len(payload) != 2 {
		t.Fatalf("payload len = %d want 2", len(payload))
	}
	unionID, data, ok := SplitDataObject(payload[0])
	if !ok {
		t.Fatalf("payload[0] is not a data object: %#v", payload[0])
	}
	if unionID != UnionUser {
		t.Fatalf("unionID = %d want %d", unionID, UnionUser)
	}
	if got := Int64Value(data[0]); got != 1 {
		t.Fatalf("data[0] = %d want 1", got)
	}
	if payload[1] != nil {
		t.Fatalf("payload[1] = %#v want nil", payload[1])
	}
}

func TestGeneratedIncludeExcludeUnions(t *testing.T) {
	opts := BuildInitialStateOptions{
		UserKey:       "filter_test",
		DisplayName:   "Filter",
		Now:           time.Now().UTC(),
		Generated:     DefaultGeneratedUserOptions(),
		ProfileLevel:  ProfileFull,
		IncludeUnions: []int64{0, 128},
	}
	payload, err := BuildGeneratedInitialState(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	counts := unionCounts(payload)
	if counts[UnionUser] != 1 || counts[UnionCurrency] != 1 {
		t.Fatalf("include filter failed: counts=%#v", counts)
	}
	if len(counts) != 2 {
		t.Fatalf("include filter should have exactly 2 union types, got %d: %#v", len(counts), counts)
	}

	opts.IncludeUnions = nil
	opts.ExcludeUnions = []int64{128}
	payload, err = BuildGeneratedInitialState(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	counts = unionCounts(payload)
	if counts[UnionCurrency] > 0 {
		t.Fatalf("exclude filter failed: union 128 should be excluded, got %d", counts[UnionCurrency])
	}
}

func TestGeneratedPayloadAgainstCapturedTemplateShape(t *testing.T) {
	capturedPath := findTestCapturedFile(t)
	if capturedPath == "" {
		t.Skip("no captured file found for template test")
	}
	templates, err := LoadTemplateRegistry(capturedPath)
	if err != nil {
		t.Fatal(err)
	}
	if templates.TemplateCount() == 0 {
		t.Skip("no templates loaded")
	}

	opts := BuildInitialStateOptions{
		UserKey:          "tmpl_test",
		DisplayName:      "Templater",
		Now:              time.Now().UTC(),
		Generated:        DefaultGeneratedUserOptions(),
		ProfileLevel:     ProfileFull,
		TemplateRegistry: templates,
	}
	payload, err := BuildGeneratedInitialState(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if len(payload) == 0 {
		t.Fatal("empty template-generated payload")
	}

	for _, raw := range payload {
		unionID, data, ok := SplitDataObject(raw)
		if !ok {
			continue
		}
		if !templates.HasTemplate(unionID) {
			continue
		}
		firstTemplate, _ := templates.GetTemplate(unionID)
		if len(data) != len(firstTemplate) {
			t.Errorf("union %d: generated arrayLen=%d, template arrayLen=%d", unionID, len(data), len(firstTemplate))
		}
	}
}

func TestTemplateRegistryLoadsTemplates(t *testing.T) {
	capturedPath := findTestCapturedFile(t)
	if capturedPath == "" {
		t.Skip("no captured file found")
	}
	reg, err := LoadTemplateRegistry(capturedPath)
	if err != nil {
		t.Fatal(err)
	}
	if reg.TemplateCount() == 0 {
		t.Fatal("template registry loaded zero templates")
	}
	_, ok := reg.GetTemplate(UnionUser)
	if !ok {
		t.Error("missing User template")
	}
	if !reg.HasTemplate(UnionUser) {
		t.Error("HasTemplate(0) should return true")
	}
}

func findTestCapturedFile(t *testing.T) string {
	t.Helper()
	patterns := []string{
		"../references/captured/*api_data_user*.api.json",
		"references/captured/*api_data_user*.api.json",
		"../../references/captured/*api_data_user*.api.json",
	}
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err == nil && len(matches) > 0 {
			return matches[0]
		}
	}
	return ""
}

func normalizeForJSONTest(value any) any {
	switch v := value.(type) {
	case time.Time:
		if v.IsZero() {
			return nil
		}
		return map[string]any{
			"__timestamp_seconds__":     v.Unix(),
			"__timestamp_nanoseconds__": int64(v.Nanosecond()),
		}
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = normalizeForJSONTest(item)
		}
		return out
	default:
		return v
	}
}

func unionCounts(payload []any) map[int64]int {
	out := map[int64]int{}
	for _, raw := range payload {
		unionID, _, ok := SplitDataObject(raw)
		if ok {
			out[unionID]++
		}
	}
	return out
}

func loadTestTemplates(t *testing.T) *UserDataTemplateRegistry {
	t.Helper()
	reg, err := LoadTemplateRegistry(findTestCapturedFile(t))
	if err != nil {
		t.Fatalf("load templates: %v", err)
	}
	if reg.TemplateCount() == 0 {
		t.Skip("no captured templates available")
	}
	return reg
}

func TestTemplateMinimalStateBuildsCoherentParty(t *testing.T) {
	templates := loadTestTemplates(t)
	opts := BuildInitialStateOptions{
		UserKey:          "tmpl_test",
		DisplayName:      "Tester",
		Now:              time.Now().UTC(),
		Generated:        DefaultGeneratedUserOptions(),
		ProfileLevel:     ProfileTemplateMinimal,
		TemplateRegistry: templates,
	}
	payload, err := BuildTemplateMinimalInitialState(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if len(payload) == 0 {
		t.Fatal("empty payload")
	}

	partyIDs := map[int64]bool{}
	characterIDs := map[int64]bool{}
	characterBaseIDs := map[int64]bool{}
	var partySlotRefs []struct {
		partyID     int64
		characterID int64
	}

	for _, raw := range payload {
		unionID, data, ok := SplitDataObject(raw)
		if !ok {
			continue
		}
		switch unionID {
		case UnionParty:
			partyIDs[fieldAt(data, 0)] = true
		case UnionPartySlot:
			pid := fieldAt(data, 1)
			cid := fieldAt(data, 3)
			partySlotRefs = append(partySlotRefs, struct {
				partyID     int64
				characterID int64
			}{pid, cid})
		case UnionCharacter:
			characterIDs[fieldAt(data, 0)] = true
			baseID := fieldAt(data, 6)
			if baseID == 0 {
				t.Errorf("character %d has zero CharacterBaseId", fieldAt(data, 0))
			}
		case UnionCharacterBase:
			characterBaseIDs[fieldAt(data, 0)] = true
		}
	}

	for _, ref := range partySlotRefs {
		if !partyIDs[ref.partyID] {
			t.Errorf("PartySlot references missing PartyId=%d", ref.partyID)
		}
		if !characterIDs[ref.characterID] {
			t.Errorf("PartySlot references missing CharacterId=%d", ref.characterID)
		}
	}

	if len(partyIDs) == 0 {
		t.Error("no Party objects in template_minimal")
	}
	if len(partySlotRefs) == 0 {
		t.Error("no PartySlot objects in template_minimal")
	}
}

func TestTemplateMinimalKeepsCapturedMasterIds(t *testing.T) {
	templates := loadTestTemplates(t)
	opts := BuildInitialStateOptions{
		UserKey:          "tmpl_master",
		DisplayName:      "Tester",
		Now:              time.Now().UTC(),
		Generated:        DefaultGeneratedUserOptions(),
		ProfileLevel:     ProfileTemplateMinimal,
		TemplateRegistry: templates,
	}
	payload, err := BuildTemplateMinimalInitialState(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}

	for _, raw := range payload {
		unionID, data, ok := SplitDataObject(raw)
		if !ok {
			continue
		}
		switch unionID {
		case UnionCharacter:
			masterID := fieldAt(data, 1)
			if masterID >= 1 && masterID <= 5 {
				t.Errorf("Character %d has hardcoded master id %d", fieldAt(data, 0), masterID)
			}
			if masterID == 0 {
				t.Errorf("Character %d has zero CharacterMasterId", fieldAt(data, 0))
			}
		case UnionCharacterBase:
			masterID := fieldAt(data, 1)
			if masterID >= 1 && masterID <= 5 {
				t.Errorf("CharacterBase %d has hardcoded master id %d", fieldAt(data, 0), masterID)
			}
			if masterID == 0 {
				t.Errorf("CharacterBase %d has zero CharacterBaseMasterId", fieldAt(data, 0))
			}
		case UnionHomeDisplayPreference:
			homeBase := fieldAt(data, 1)
			memberBase := fieldAt(data, 3)
			storyBase := fieldAt(data, 4)
			shopBase := fieldAt(data, 5)
			illustMaster := fieldAt(data, 11)
			if homeBase >= 1 && homeBase <= 5 {
				t.Errorf("HomeDisplayPreference.HomeCharacterBaseMasterId is hardcoded %d", homeBase)
			}
			if memberBase >= 1 && memberBase <= 5 {
				t.Errorf("HomeDisplayPreference.MemberCharacterBaseMasterId is hardcoded %d", memberBase)
			}
			if storyBase >= 1 && storyBase <= 5 {
				t.Errorf("HomeDisplayPreference.StoryCharacterBaseMasterId is hardcoded %d", storyBase)
			}
			if shopBase >= 1 && shopBase <= 5 {
				t.Errorf("HomeDisplayPreference.ShopCharacterBaseMasterId is hardcoded %d", shopBase)
			}
			if illustMaster >= 1 && illustMaster <= 5 {
				t.Errorf("HomeDisplayPreference.IllustCharacterMasterId is hardcoded %d", illustMaster)
			}
		case UnionUserProfile:
			mainMasterID := fieldAt(data, 16)
			if mainMasterID >= 1 && mainMasterID <= 5 {
				t.Errorf("UserProfile.MainCharacterMasterId is hardcoded %d", mainMasterID)
			}
		}
	}
}

func TestTemplateMinimalKeepsRequiredUnions(t *testing.T) {
	templates := loadTestTemplates(t)
	opts := BuildInitialStateOptions{
		UserKey:          "tmpl_req",
		DisplayName:      "Tester",
		Now:              time.Now().UTC(),
		Generated:        DefaultGeneratedUserOptions(),
		ProfileLevel:     ProfileTemplateMinimal,
		TemplateRegistry: templates,
	}
	payload, err := BuildTemplateMinimalInitialState(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}

	counts := unionCounts(payload)
	required := []int64{
		UnionUser, UnionUserProfile, UnionUserPreference,
		UnionHomeDisplayPreference, UnionCharacter, UnionCharacterBase,
		UnionParty, UnionPartySlot,
		UnionViewedShop, UnionGameHint, UnionUserBonus,
		UnionCurrency, UnionRestriction, UnionAlbumPreset,
		UnionFriendInvitation,
	}
	for _, uid := range required {
		if counts[uid] == 0 {
			t.Errorf("required union %d missing in template_minimal", uid)
		}
	}
	if counts[UnionPartySlot] < 5 {
		t.Errorf("expected >=5 PartySlots, got %d", counts[UnionPartySlot])
	}
}
