package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ymst_server/core/database"
	"ymst_server/repositories"
)

func TestAuthMiddlewareSetsUserKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open("file:middleware_auth?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	sessions := repositories.NewSessionRepository(db)
	session, err := sessions.CreateSession(context.Background(), "usr_middleware", time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	r := gin.New()
	r.Use(AuthMiddleware(sessions))
	r.GET("/api/data/user", func(c *gin.Context) {
		userKey, ok := UserKeyFromContext(c)
		if !ok {
			t.Fatalf("missing user key")
		}
		c.String(http.StatusOK, userKey)
	})

	req := httptest.NewRequest(http.MethodGet, "/api/data/user", nil)
	req.Header.Set("Authorization", "Bearer "+session.SessionToken)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if w.Body.String() != "usr_middleware" {
		t.Fatalf("body = %s", w.Body.String())
	}
}

func TestAuthMiddlewareRejectsMissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open("file:middleware_auth_missing?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	r := gin.New()
	r.Use(AuthMiddleware(repositories.NewSessionRepository(db)))
	r.GET("/api/data/user", func(c *gin.Context) { c.Status(http.StatusOK) })
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/data/user", nil))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", w.Code)
	}
}
