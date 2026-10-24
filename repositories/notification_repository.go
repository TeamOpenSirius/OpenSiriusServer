package repositories

import (
	"context"
	"time"

	"gorm.io/gorm"

	"ymst_server/core/database"
)

type NotificationRepository struct{ db *gorm.DB }

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) EnsureSeed(ctx context.Context) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&database.GameNotification{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	now := time.Now()
	return r.db.WithContext(ctx).Create(&database.GameNotification{NotificationID: now.Unix(), Title: "お知らせ", Body: `[{"Element":2,"Data":"お知らせ"},{"Element":4,"Data":"これはテスト用のお知らせ本文です。"}]`, TabCategory: 2, Category: 1, IsConfirmation: true, SortOrder: 1, PostingAt: now, LastUpdatedAt: now}).Error
}

func (r *NotificationRepository) List(ctx context.Context) ([]database.GameNotification, error) {
	if err := r.EnsureSeed(ctx); err != nil {
		return nil, err
	}
	var rows []database.GameNotification
	err := r.db.WithContext(ctx).Order("sort_order ASC, posting_at DESC").Find(&rows).Error
	return rows, err
}

func (r *NotificationRepository) Find(ctx context.Context, id int64) (*database.GameNotification, error) {
	if err := r.EnsureSeed(ctx); err != nil {
		return nil, err
	}
	var row database.GameNotification
	err := r.db.WithContext(ctx).Where("notification_id = ?", id).First(&row).Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}

type NotificationReadRepository struct{ db *gorm.DB }

func NewNotificationReadRepository(db *gorm.DB) *NotificationReadRepository {
	return &NotificationReadRepository{db: db}
}
func (r *NotificationReadRepository) Upsert(ctx context.Context, userKey string, tabCategory int, readAt time.Time) error {
	var row database.GameNotificationReadState
	err := r.db.WithContext(ctx).Where("user_key = ? AND tab_category = ?", userKey, tabCategory).First(&row).Error
	if err == nil {
		return r.db.WithContext(ctx).Model(&row).Update("read_at", readAt).Error
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	return r.db.WithContext(ctx).Create(&database.GameNotificationReadState{UserKey: userKey, TabCategory: tabCategory, ReadAt: readAt}).Error
}
