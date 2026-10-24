package repositories

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ymst_server/apiwire"
	"ymst_server/core/database"
	"ymst_server/userdata"
)

func newGeneratedUserRepoTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file:user_generated_"+time.Now().Format("150405.000000000")+"?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestGeneratedStateCreatesRequiredUnions(t *testing.T) {
	payload, err := userdata.BuildGeneratedInitialState(context.Background(), userdata.BuildInitialStateOptions{UserKey: "usr_generated", DisplayName: "Alice", Now: time.Now().UTC(), Generated: userdata.DefaultGeneratedUserOptions(), ProfileLevel: userdata.ProfileFull})
	if err != nil {
		t.Fatal(err)
	}
	counts := unionCounts(payload)
	requiredExact := map[int64]int{UnionUser: 1, UnionUserProfile: 1, UnionUserPreference: 1, UnionHomeDisplayPreference: 1, UnionUserBonus: 1, UnionCurrency: 1, UnionRestriction: 1, UnionFriendInvitation: 1}
	for unionID, want := range requiredExact {
		if counts[unionID] != want {
			t.Fatalf("union %d count = %d want %d", unionID, counts[unionID], want)
		}
	}
	requiredMin := map[int64]int{UnionCharacter: 5, UnionCharacterBase: 5, UnionParty: 1, UnionPartySlot: 5, UnionLive: 3, UnionMusic: 3, UnionViewedShop: 1, UnionGameHint: 1, UnionAlbumPreset: 1}
	for unionID, wantMin := range requiredMin {
		if counts[unionID] < wantMin {
			t.Fatalf("union %d count = %d want >= %d", unionID, counts[unionID], wantMin)
		}
	}
}

func TestGeneratedStateReferencesAreValid(t *testing.T) {
	payload, err := userdata.BuildGeneratedInitialState(context.Background(), userdata.BuildInitialStateOptions{UserKey: "usr_refs", DisplayName: "Alice", Now: time.Now().UTC(), Generated: userdata.DefaultGeneratedUserOptions(), ProfileLevel: userdata.ProfileFull})
	if err != nil {
		t.Fatal(err)
	}
	characters := map[int64]bool{}
	characterBases := map[int64]bool{}
	parties := map[int64]bool{}
	var profileMainCharacterID int64
	for _, raw := range payload {
		unionID, data, ok := splitDataObject(raw)
		if !ok {
			continue
		}
		switch unionID {
		case UnionCharacter:
			characters[int64At(data, 0)] = true
			if !characterBases[int64At(data, 6)] {
				t.Fatalf("invalid character base ref: %#v", data)
			}
		case UnionCharacterBase:
			characterBases[int64At(data, 0)] = true
		case UnionParty:
			parties[int64At(data, 0)] = true
		case UnionPartySlot:
			if !parties[int64At(data, 1)] {
				t.Fatalf("party slot party ref missing: %#v", data)
			}
			if !characters[int64At(data, 3)] {
				t.Fatalf("party slot character ref missing: %#v", data)
			}
		case UnionUserProfile:
			profileMainCharacterID = int64At(data, 3)
		}
	}
	if !characters[profileMainCharacterID] {
		t.Fatalf("profile main character ref missing: %d", profileMainCharacterID)
	}
}

func TestEmptyDatabaseAutoGeneratesRawState(t *testing.T) {
	db := newGeneratedUserRepoTestDB(t)
	repo := NewUserRepository(db, "").WithUserKey("usr_empty")
	var before int64
	if err := db.Model(&database.GameRawDataObject{}).Where("user_key = ?", "usr_empty").Count(&before).Error; err != nil {
		t.Fatal(err)
	}
	if before != 0 {
		t.Fatalf("raw rows before build = %d", before)
	}
	payload, err := repo.BuildDataObjects(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var after int64
	if err := db.Model(&database.GameRawDataObject{}).Where("user_key = ?", "usr_empty").Count(&after).Error; err != nil {
		t.Fatal(err)
	}
	if after != int64(len(payload)) {
		t.Fatalf("raw row count = %d want %d", after, len(payload))
	}
	var meta database.GameUserStateMeta
	if err := db.Where("user_key = ?", "usr_empty").First(&meta).Error; err != nil {
		t.Fatal(err)
	}
	if meta.SeedMode != userdata.SeedModeGenerated || meta.ObjectCount != len(payload) {
		t.Fatalf("bad state meta: %#v", meta)
	}
}

func TestCapturedJSONRoundTripKeepsNilAndUnionCounts(t *testing.T) {
	db := newGeneratedUserRepoTestDB(t)
	seedPayload := []any{
		apiwire.DataObject(UnionUser, []any{int64(1), int64(1), int64(0), int64(100), time.Now().UTC(), int64(0), int64(3000), int64(100000), int64(999), int64(0), nil, nil, time.Now().UTC(), "hash", int64(0), int64(99), int64(0), time.Now().UTC(), false, false}),
		nil,
		apiwire.DataObject(UnionGameHint, []any{int64(17000033), int64(33), true}),
	}
	body, err := json.Marshal(map[string]any{"payload": seedPayload})
	if err != nil {
		t.Fatal(err)
	}
	seedPath := filepath.Join(t.TempDir(), "api_data_user.bin.api.json")
	if err := os.WriteFile(seedPath, body, 0644); err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepository(db, seedPath).WithSeedMode(userdata.SeedModeCapturedJSON).WithUserKey("usr_captured")
	payload, err := repo.BuildDataObjects(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(payload) != len(seedPayload) || payload[1] != nil {
		t.Fatalf("captured payload shape changed: len=%d nil=%v", len(payload), payload[1] == nil)
	}
	counts := unionCounts(payload)
	if counts[UnionUser] != 1 || counts[UnionGameHint] != 1 {
		t.Fatalf("captured union counts mismatch: %#v", counts)
	}
	stats, err := repo.Stats(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if stats.NilSlotCount != 1 || stats.SeedMode != userdata.SeedModeCapturedJSON {
		t.Fatalf("bad captured stats: %#v", stats)
	}
}

func TestGeneratedStatePersistsReloadsAndIgnoresTypedTables(t *testing.T) {
	db := newGeneratedUserRepoTestDB(t)
	repo := NewUserRepository(db, "").WithUserKey("usr_persist")
	first, err := repo.BuildDataObjects(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	second, err := repo.BuildDataObjects(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(first) != len(second) {
		t.Fatalf("object count changed: %d -> %d", len(first), len(second))
	}
	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&database.GameUser{}).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&database.GameCharacter{}).Error; err != nil {
		t.Fatal(err)
	}
	afterClear, err := repo.BuildDataObjects(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(afterClear) != len(first) {
		t.Fatalf("typed clear affected raw payload: %d -> %d", len(first), len(afterClear))
	}
}

func TestMutatorsModifyRawState(t *testing.T) {
	db := newGeneratedUserRepoTestDB(t)
	repo := NewUserRepository(db, "").WithUserKey("usr_mutator")
	if _, err := repo.BuildDataObjects(context.Background()); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.UpdateTutorialStatus(42); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.MarkGameHintsRead([]int64{33}); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.MarkShopViewedWithExchangeShopMasterID(2, nil); err != nil {
		t.Fatal(err)
	}
	payload, err := repo.BuildDataObjects(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	foundHint := false
	foundShop := false
	for _, raw := range payload {
		unionID, data, ok := splitDataObject(raw)
		if !ok {
			continue
		}
		if unionID == UnionUser && int64At(data, 15) != 42 {
			t.Fatalf("tutorial not updated: %#v", data)
		}
		if unionID == UnionGameHint && int64At(data, 1) == 33 && boolAt(data, 2) {
			foundHint = true
		}
		if unionID == UnionViewedShop && int64At(data, 3) == 2 {
			foundShop = true
		}
	}
	if !foundHint || !foundShop {
		t.Fatalf("mutated objects missing: hint=%v shop=%v", foundHint, foundShop)
	}
}

func unionCounts(payload []any) map[int64]int {
	out := map[int64]int{}
	for _, raw := range payload {
		unionID, _, ok := splitDataObject(raw)
		if ok {
			out[unionID]++
		}
	}
	return out
}
