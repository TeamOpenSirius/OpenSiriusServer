package repositories

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"ymst_server/core/database"
)

const (
	GameAccountStatusNormal  = 0
	GameAccountStatusDeleted = 1
	GameAccountStatusBanned  = 2
)

type AccountRepository struct{ db *gorm.DB }

func NewAccountRepository(db *gorm.DB) *AccountRepository { return &AccountRepository{db: db} }

func (r *AccountRepository) CreateAccount(ctx context.Context, name string) (*database.GameAccount, error) {
	if strings.TrimSpace(name) == "" {
		name = "Player"
	}
	var lastErr error
	for i := 0; i < 8; i++ {
		userKey, err := NewToken("usr")
		if err != nil {
			return nil, err
		}
		loginToken, err := NewToken("login")
		if err != nil {
			return nil, err
		}
		account := &database.GameAccount{UserKey: userKey, LoginToken: loginToken, Name: name, Status: GameAccountStatusNormal, BanLevel: 0}
		if err := r.db.WithContext(ctx).Create(account).Error; err != nil {
			lastErr = err
			continue
		}
		return account, nil
	}
	if lastErr == nil {
		lastErr = errors.New("create account failed")
	}
	return nil, lastErr
}

func (r *AccountRepository) CreateAccountWithLoginToken(ctx context.Context, name string, loginToken string) (*database.GameAccount, error) {
	if strings.TrimSpace(name) == "" {
		name = "Player"
	}
	loginToken = strings.TrimSpace(loginToken)
	if loginToken == "" {
		return nil, errors.New("missing login token")
	}
	var lastErr error
	for i := 0; i < 8; i++ {
		userKey, err := NewToken("usr")
		if err != nil {
			return nil, err
		}
		account := &database.GameAccount{UserKey: userKey, LoginToken: loginToken, Name: name, Status: GameAccountStatusNormal, BanLevel: 0}
		if err := r.db.WithContext(ctx).Create(account).Error; err != nil {
			lastErr = err
			continue
		}
		return account, nil
	}
	if lastErr == nil {
		lastErr = errors.New("create account failed")
	}
	return nil, lastErr
}

func (r *AccountRepository) FindByLoginToken(ctx context.Context, loginToken string) (*database.GameAccount, error) {
	var account database.GameAccount
	err := r.db.WithContext(ctx).Where("login_token = ?", strings.TrimSpace(loginToken)).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) FindByUserKey(ctx context.Context, userKey string) (*database.GameAccount, error) {
	var account database.GameAccount
	err := r.db.WithContext(ctx).Where("user_key = ?", strings.TrimSpace(userKey)).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) MarkAuthenticated(ctx context.Context, userKey string, at time.Time) error {
	return r.db.WithContext(ctx).Model(&database.GameAccount{}).Where("user_key = ?", strings.TrimSpace(userKey)).Update("last_authenticated_at", &at).Error
}

func (r *AccountRepository) MarkDeleted(ctx context.Context, userKey string) error {
	return r.db.WithContext(ctx).Model(&database.GameAccount{}).Where("user_key = ?", strings.TrimSpace(userKey)).Update("status", GameAccountStatusDeleted).Error
}
