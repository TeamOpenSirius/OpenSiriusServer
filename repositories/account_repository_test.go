package repositories

import (
	"context"
	"strings"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ymst_server/core/database"
)

func newAuthRepoTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file:authrepo_"+time.Now().Format("150405.000000000")+"?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestAccountRepositoryCreateAndFind(t *testing.T) {
	db := newAuthRepoTestDB(t)
	repo := NewAccountRepository(db)
	account, err := repo.CreateAccount(context.Background(), "Alice")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(account.UserKey, "usr_") {
		t.Fatalf("unexpected user key: %s", account.UserKey)
	}
	if !strings.HasPrefix(account.LoginToken, "login_") {
		t.Fatalf("unexpected login token: %s", account.LoginToken)
	}
	found, err := repo.FindByLoginToken(context.Background(), account.LoginToken)
	if err != nil {
		t.Fatal(err)
	}
	if found.UserKey != account.UserKey {
		t.Fatalf("user key mismatch")
	}
	if err := repo.MarkAuthenticated(context.Background(), account.UserKey, time.Now().UTC()); err != nil {
		t.Fatal(err)
	}
	found, err = repo.FindByUserKey(context.Background(), account.UserKey)
	if err != nil {
		t.Fatal(err)
	}
	if found.LastAuthenticatedAt == nil {
		t.Fatalf("last authenticated not set")
	}
}

func TestSessionRepositoryCreateFindTouchRevoke(t *testing.T) {
	db := newAuthRepoTestDB(t)
	repo := NewSessionRepository(db)
	session, err := repo.CreateSession(context.Background(), "usr_test", time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(session.SessionToken, "sess_") {
		t.Fatalf("unexpected session token: %s", session.SessionToken)
	}
	found, err := repo.FindByToken(context.Background(), session.SessionToken)
	if err != nil {
		t.Fatal(err)
	}
	if found.UserKey != "usr_test" {
		t.Fatalf("user key mismatch")
	}
	if err := repo.Touch(context.Background(), session.SessionToken, time.Now().UTC()); err != nil {
		t.Fatal(err)
	}
	if err := repo.Revoke(context.Background(), session.SessionToken); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.FindByToken(context.Background(), session.SessionToken); err == nil {
		t.Fatalf("revoked session found")
	}
}
