package repositories

import (
	"context"
	"time"

	"gorm.io/gorm"

	"ymst_server/core/database"
)

type ConfirmationCodeRepository struct{ db *gorm.DB }

func NewConfirmationCodeRepository(db *gorm.DB) *ConfirmationCodeRepository {
	return &ConfirmationCodeRepository{db: db}
}

func (r *ConfirmationCodeRepository) FindActive(ctx context.Context, userKey string, purpose string, now time.Time) (*database.GameConfirmationCode, error) {
	var row database.GameConfirmationCode
	err := r.db.WithContext(ctx).Where("user_key = ? AND purpose = ? AND used_at IS NULL AND expires_at > ?", userKey, purpose, now).Order("expires_at DESC").First(&row).Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *ConfirmationCodeRepository) Create(ctx context.Context, userKey string, purpose string, code string, expiresAt time.Time) (*database.GameConfirmationCode, error) {
	row := &database.GameConfirmationCode{UserKey: userKey, Purpose: purpose, Code: code, ExpiresAt: expiresAt}
	if err := r.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, err
	}
	return row, nil
}
