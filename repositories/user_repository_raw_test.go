package repositories

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ymst_server/apiwire"
	"ymst_server/core/database"
)

func testUserRepo(t *testing.T) *UserRepository {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	return NewUserRepository(db, "")
}

func TestNormalizeJSONArrayKeepsNil(t *testing.T) {
	out := normalizeJSONArray([]any{[]any{int64(0), []any{int64(1)}}, nil, []any{int64(66), []any{int64(2)}}})
	if len(out) != 3 || out[1] != nil {
		t.Fatalf("nil slot not preserved: %#v", out)
	}
}

func TestSaveRawPayloadBuildRawDataObjectsKeepsOrderAndNil(t *testing.T) {
	repo := testUserRepo(t)
	payload := []any{
		apiwire.DataObject(UnionUser, []any{int64(1), int64(1), int64(0), int64(10), time.Time{}, int64(0), int64(0), int64(0), int64(100), int64(0), time.Time{}, nil, time.Time{}, "hash", int64(0), int64(1), int64(0), time.Time{}, false, false}),
		nil,
		apiwire.DataObject(UnionGameHint, []any{int64(2), int64(33), false}),
	}
	if err := repo.db.Transaction(func(tx *gorm.DB) error { return repo.SaveRawPayload(tx, payload, time.Now()) }); err != nil {
		t.Fatal(err)
	}
	got, ok, err := repo.BuildRawDataObjects()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("raw table not used")
	}
	if len(got) != 3 || got[1] != nil {
		t.Fatalf("unexpected raw payload: %#v", got)
	}
}

func TestUpdateTutorialStatusSyncsRaw(t *testing.T) {
	repo := testUserRepo(t)
	now := time.Now()
	user := database.GameUser{ID: 1, ObjectIndex: 0, TutorialStatus: 1, SplashLastDisplayedAt: now, CreatedAt: now, UpdatedAt: now}
	if err := repo.db.Create(&user).Error; err != nil {
		t.Fatal(err)
	}
	if err := repo.UpsertRawObject(repo.db, user.ObjectIndex, UnionUser, []any{user.ID, int64(0), int64(0), int64(0), time.Time{}, int64(0), int64(0), int64(0), int64(0), int64(0), time.Time{}, nil, time.Time{}, "", int64(0), int64(1), int64(0), now, false, false}); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.UpdateTutorialStatus(42); err != nil {
		t.Fatal(err)
	}
	got, ok, err := repo.BuildRawDataObjects()
	if err != nil {
		t.Fatal(err)
	}
	if !ok || len(got) == 0 {
		t.Fatal("raw objects missing")
	}
	_, data, ok := splitDataObject(got[0])
	if !ok {
		t.Fatalf("invalid data object: %#v", got[0])
	}
	if int64At(data, 15) != 42 {
		t.Fatalf("tutorial status not synced: %#v", data)
	}
}

func TestMarkGameHintsReadAddsRawObject(t *testing.T) {
	repo := testUserRepo(t)
	updates, err := repo.MarkGameHintsRead([]int64{33})
	if err != nil {
		t.Fatal(err)
	}
	if len(updates) == 0 {
		t.Fatal("expected updates")
	}
	got, ok, err := repo.BuildRawDataObjects()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("raw table not used")
	}
	found := false
	for _, raw := range got {
		unionID, data, ok := splitDataObject(raw)
		if ok && unionID == UnionGameHint && int64At(data, 1) == 33 && boolAt(data, 2) {
			found = true
		}
	}
	if !found {
		t.Fatalf("new game hint not found in raw payload: %#v", got)
	}
}
