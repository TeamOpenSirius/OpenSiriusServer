package repositories

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ymst_server/core/database"
)

func TestSessionRepositoryExpiredSessionNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:session_expired?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	repo := NewSessionRepository(db)
	session, err := repo.CreateSession(context.Background(), "usr_expired", -time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := repo.FindByToken(context.Background(), session.SessionToken); err == nil {
		t.Fatalf("expired session found")
	}
}
