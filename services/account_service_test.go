package services

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
	"ymst_server/repositories"
)

func newAccountServiceTest(t *testing.T) (*AccountService, *repositories.UserRepository, *gorm.DB) {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file:accountsvc_"+time.Now().Format("150405.000000000")+"?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	seedPath := filepath.Join(t.TempDir(), "seed.json")
	seed := map[string]any{"payload": []any{
		apiwire.DataObject(repositories.UnionUser, []any{int64(1), int64(1), int64(0), int64(10), time.Time{}, int64(0), int64(0), int64(0), int64(100), int64(0), time.Time{}, nil, time.Time{}, "hash", int64(0), int64(1), int64(0), time.Time{}, false, false}),
		apiwire.DataObject(repositories.UnionUserProfile, []any{int64(1), "Seed", "", int64(1)}),
	}}
	body, err := json.Marshal(seed)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(seedPath, body, 0644); err != nil {
		t.Fatal(err)
	}
	users := repositories.NewUserRepository(db, seedPath)
	service := NewAccountService(repositories.NewAccountRepository(db), repositories.NewSessionRepository(db), users)
	return service, users, db
}

func TestAccountServiceRegisterCreatesAccountAndUserState(t *testing.T) {
	svc, users, db := newAccountServiceTest(t)
	result, err := svc.Register(context.Background(), "Alice")
	if err != nil {
		t.Fatal(err)
	}
	if result.LoginToken == "" {
		t.Fatalf("missing login token")
	}
	account, err := svc.accounts.FindByLoginToken(context.Background(), result.LoginToken)
	if err != nil {
		t.Fatal(err)
	}
	var rawCount int64
	if err := db.Model(&database.GameRawDataObject{}).Where("user_key = ?", account.UserKey).Count(&rawCount).Error; err != nil {
		t.Fatal(err)
	}
	if rawCount == 0 {
		t.Fatalf("raw state not cloned")
	}
	payload, err := users.WithUserKey(account.UserKey).BuildDataObjects()
	if err != nil {
		t.Fatal(err)
	}
	if !payloadHasProfileName(payload, "Alice") {
		t.Fatalf("profile name not patched: %#v", payload)
	}
}

func TestAccountServiceAuthenticateCreatesSession(t *testing.T) {
	svc, _, _ := newAccountServiceTest(t)
	registered, err := svc.Register(context.Background(), "Alice")
	if err != nil {
		t.Fatal(err)
	}
	authenticated, err := svc.Authenticate(context.Background(), apiwire.AuthenticatePayload{LoginToken: registered.LoginToken})
	if err != nil {
		t.Fatal(err)
	}
	if authenticated.SessionToken == "" || authenticated.SessionToken == registered.LoginToken {
		t.Fatalf("invalid session token")
	}
	session, err := svc.sessions.FindByToken(context.Background(), authenticated.SessionToken)
	if err != nil {
		t.Fatal(err)
	}
	account, err := svc.accounts.FindByLoginToken(context.Background(), registered.LoginToken)
	if err != nil {
		t.Fatal(err)
	}
	if session.UserKey != account.UserKey {
		t.Fatalf("session user key mismatch")
	}
}

func TestAccountServiceTwoUsersIsolatedRawState(t *testing.T) {
	svc, users, _ := newAccountServiceTest(t)
	alice, err := svc.Register(context.Background(), "Alice")
	if err != nil {
		t.Fatal(err)
	}
	bob, err := svc.Register(context.Background(), "Bob")
	if err != nil {
		t.Fatal(err)
	}
	aliceAccount, err := svc.accounts.FindByLoginToken(context.Background(), alice.LoginToken)
	if err != nil {
		t.Fatal(err)
	}
	bobAccount, err := svc.accounts.FindByLoginToken(context.Background(), bob.LoginToken)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := users.WithUserKey(aliceAccount.UserKey).UpdateTutorialStatus(10); err != nil {
		t.Fatal(err)
	}
	if _, err := users.WithUserKey(bobAccount.UserKey).UpdateTutorialStatus(20); err != nil {
		t.Fatal(err)
	}
	alicePayload, err := users.WithUserKey(aliceAccount.UserKey).BuildDataObjects()
	if err != nil {
		t.Fatal(err)
	}
	bobPayload, err := users.WithUserKey(bobAccount.UserKey).BuildDataObjects()
	if err != nil {
		t.Fatal(err)
	}
	if tutorialStatus(alicePayload) != 10 {
		t.Fatalf("alice tutorial = %d", tutorialStatus(alicePayload))
	}
	if tutorialStatus(bobPayload) != 20 {
		t.Fatalf("bob tutorial = %d", tutorialStatus(bobPayload))
	}
	if _, err := users.WithUserKey(aliceAccount.UserKey).MarkGameHintsRead([]int64{33}); err != nil {
		t.Fatal(err)
	}
	bobPayload, err = users.WithUserKey(bobAccount.UserKey).BuildDataObjects()
	if err != nil {
		t.Fatal(err)
	}
	if payloadHasGameHintCategory(bobPayload, 33) {
		t.Fatalf("bob has alice game hint update")
	}
}

func payloadHasProfileName(payload []any, name string) bool {
	for _, raw := range payload {
		obj, ok := raw.([]any)
		if !ok || len(obj) < 2 || apiwire.Int64Value(obj[0]) != repositories.UnionUserProfile {
			continue
		}
		arr, ok := obj[1].([]any)
		if ok && len(arr) > 1 && arr[1] == name {
			return true
		}
	}
	return false
}

func tutorialStatus(payload []any) int64 {
	for _, raw := range payload {
		obj, ok := raw.([]any)
		if !ok || len(obj) < 2 || apiwire.Int64Value(obj[0]) != repositories.UnionUser {
			continue
		}
		arr, ok := obj[1].([]any)
		if ok {
			return apiwire.Int64At(arr, 15)
		}
	}
	return 0
}

func payloadHasGameHintCategory(payload []any, category int64) bool {
	for _, raw := range payload {
		obj, ok := raw.([]any)
		if !ok || len(obj) < 2 || apiwire.Int64Value(obj[0]) != repositories.UnionGameHint {
			continue
		}
		arr, ok := obj[1].([]any)
		if ok && apiwire.Int64At(arr, 1) == category {
			return true
		}
	}
	return false
}
