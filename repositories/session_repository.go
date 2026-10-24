package repositories

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"ymst_server/core/database"
)

type SessionRepository struct{ db *gorm.DB }

func NewSessionRepository(db *gorm.DB) *SessionRepository { return &SessionRepository{db: db} }

func (r *SessionRepository) CreateSession(ctx context.Context, userKey string, ttl time.Duration) (*database.GameSession, error) {
	userKey = strings.TrimSpace(userKey)
	if userKey == "" {
		return nil, errors.New("missing user key")
	}
	var expiresAt *time.Time
	if ttl != 0 {
		t := time.Now().UTC().Add(ttl)
		expiresAt = &t
	}
	var lastErr error
	for i := 0; i < 8; i++ {
		token, err := NewToken("sess")
		if err != nil {
			return nil, err
		}
		now := time.Now().UTC()
		session := &database.GameSession{UserKey: userKey, SessionToken: token, ExpiresAt: expiresAt, LastSeenAt: now}
		if err := r.db.WithContext(ctx).Create(session).Error; err != nil {
			lastErr = err
			continue
		}
		return session, nil
	}
	if lastErr == nil {
		lastErr = errors.New("create session failed")
	}
	return nil, lastErr
}

func (r *SessionRepository) FindByToken(ctx context.Context, sessionToken string) (*database.GameSession, error) {
	sessionToken = strings.TrimSpace(sessionToken)
	if sessionToken == "" {
		return nil, gorm.ErrRecordNotFound
	}
	var session database.GameSession
	err := r.db.WithContext(ctx).Where("session_token = ? AND revoked_at IS NULL", sessionToken).First(&session).Error
	if err != nil {
		return nil, err
	}
	if session.ExpiresAt != nil && !session.ExpiresAt.After(time.Now().UTC()) {
		return nil, gorm.ErrRecordNotFound
	}
	return &session, nil
}

func (r *SessionRepository) Touch(ctx context.Context, sessionToken string, at time.Time) error {
	return r.db.WithContext(ctx).Model(&database.GameSession{}).Where("session_token = ? AND revoked_at IS NULL", strings.TrimSpace(sessionToken)).Update("last_seen_at", at).Error
}

func (r *SessionRepository) Revoke(ctx context.Context, sessionToken string) error {
	now := time.Now().UTC()
	return r.db.WithContext(ctx).Model(&database.GameSession{}).Where("session_token = ? AND revoked_at IS NULL", strings.TrimSpace(sessionToken)).Update("revoked_at", &now).Error
}

func (r *SessionRepository) RevokeAllForUser(ctx context.Context, userKey string) error {
	now := time.Now().UTC()
	return r.db.WithContext(ctx).Model(&database.GameSession{}).Where("user_key = ? AND revoked_at IS NULL", strings.TrimSpace(userKey)).Update("revoked_at", &now).Error
}
