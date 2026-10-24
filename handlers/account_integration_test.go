package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/database"
	"ymst_server/repositories"
	"ymst_server/services"
)

func TestAccountRegisterAuthenticateFlow(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open("file:handler_account_flow?mode=memory&cache=shared"), &gorm.Config{})
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
	sessionRepo := repositories.NewSessionRepository(db)
	accountRepo := repositories.NewAccountRepository(db)
	accountSvc := services.NewAccountService(accountRepo, sessionRepo, users)
	extraSvc := services.NewAccountExtraService(accountRepo, sessionRepo, repositories.NewConfirmationCodeRepository(db))
	h := NewAccountHandler(nil, config.GameConfig{}, db, "test-secret", accountSvc, extraSvc)
	r := gin.New()
	h.RegisterRoutes(r)

	registerBody, err := msgpack.Marshal([]any{"Alice"})
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/Account/Register", bytes.NewReader(registerBody))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("register status = %d", w.Code)
	}
	loginToken := decodePayloadToken(t, w.Body.Bytes())
	if loginToken == "" {
		t.Fatalf("missing login token")
	}

	authBody, err := msgpack.Marshal([]any{loginToken, int64(0), "", "", ""})
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/api/Account/Authenticate", bytes.NewReader(authBody))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("authenticate status = %d", w.Code)
	}
	sessionToken := decodePayloadToken(t, w.Body.Bytes())
	if sessionToken == "" || sessionToken == loginToken {
		t.Fatalf("invalid session token")
	}
	if _, err := sessionRepo.FindByToken(context.Background(), sessionToken); err != nil {
		t.Fatal(err)
	}
}

func decodePayloadToken(t *testing.T, body []byte) string {
	t.Helper()
	dec := msgpack.NewDecoder(bytes.NewReader(body))
	var common []any
	var payload []any
	if err := dec.Decode(&common); err != nil {
		t.Fatal(err)
	}
	if err := dec.Decode(&payload); err != nil {
		t.Fatal(err)
	}
	if len(payload) == 0 {
		return ""
	}
	token, _ := payload[0].(string)
	return token
}
