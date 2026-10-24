package services

import (
	"context"
	"sort"
	"sync"
	"time"

	"gorm.io/gorm"

	"ymst_server/apiwire"
	"ymst_server/core/database"
	"ymst_server/repositories"
)

type Notification struct {
	ID             int64
	BannerPath     string
	Title          string
	Body           string
	PostingAt      time.Time
	LastUpdatedAt  time.Time
	TabCategory    int64
	Category       int64
	IsConfirmation bool
	Order          int64
}

type NotificationService struct {
	mu                sync.RWMutex
	notifications     map[int64]Notification
	readAtByTab       map[int64]time.Time
	notificationsRepo *repositories.NotificationRepository
	readsRepo         *repositories.NotificationReadRepository
}

func NewNotificationService(enableLocal bool) *NotificationService {
	s := &NotificationService{
		notifications: make(map[int64]Notification),
		readAtByTab:   make(map[int64]time.Time),
	}

	if enableLocal {
		now := time.Now()
		s.notifications[1] = Notification{
			ID:             1,
			BannerPath:     "",
			Title:          "测试的公告",
			Body:           "本地wds服务器 by cubevlmu",
			PostingAt:      now,
			LastUpdatedAt:  now,
			TabCategory:    1,
			Category:       1,
			IsConfirmation: false,
			Order:          1,
		}
	}

	return s
}

func NewNotificationDBService(notifications *repositories.NotificationRepository, reads *repositories.NotificationReadRepository) *NotificationService {
	return &NotificationService{notifications: map[int64]Notification{}, readAtByTab: map[int64]time.Time{}, notificationsRepo: notifications, readsRepo: reads}
}

func (s *NotificationService) GetNotifications(ctx context.Context, userKey string) ([]any, error) {
	if s.notificationsRepo == nil {
		items := s.List()
		out := make([]any, 0, len(items))
		for _, n := range items {
			out = append(out, notificationResultCodec(n))
		}
		return out, nil
	}
	rows, err := s.notificationsRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(rows))
	for _, row := range rows {
		out = append(out, notificationRowResult(row))
	}
	return out, nil
}

func (s *NotificationService) GetNotificationContent(ctx context.Context, userKey string, id int64) ([]any, error) {
	if s.notificationsRepo == nil {
		n, ok := s.Get(id)
		if !ok {
			return notificationResultContentFallback(id), nil
		}
		return notificationContentCodec(n), nil
	}
	row, err := s.notificationsRepo.Find(ctx, id)
	if err == gorm.ErrRecordNotFound {
		return notificationResultContentFallback(id), nil
	}
	if err != nil {
		return nil, err
	}
	return notificationRowContent(*row), nil
}

func (s *NotificationService) UpdateReadTime(ctx context.Context, userKey string, tabCategory int, readAt time.Time) error {
	if readAt.IsZero() {
		readAt = time.Now()
	}
	if s.readsRepo == nil {
		s.MarkRead(int64(tabCategory), readAt)
		return nil
	}
	return s.readsRepo.Upsert(ctx, userKey, tabCategory, readAt)
}

func notificationRowResult(row database.GameNotification) []any {
	return apiwire.NotificationResult(row.NotificationID, row.BannerPath, row.Title, row.PostingAt, row.LastUpdatedAt, int64(row.TabCategory), int64(row.Category), row.IsConfirmation, int64(row.SortOrder))
}

func notificationRowContent(row database.GameNotification) []any {
	return apiwire.NotificationContentResult(row.NotificationID, normalizeNotificationBodyText(row.Body), row.Title, row.PostingAt, row.LastUpdatedAt, int64(row.TabCategory), int64(row.Category), row.BannerPath)
}

func notificationResultCodec(n Notification) []any {
	return apiwire.NotificationResult(n.ID, n.BannerPath, n.Title, n.PostingAt, n.LastUpdatedAt, n.TabCategory, n.Category, n.IsConfirmation, n.Order)
}

func notificationContentCodec(n Notification) []any {
	return apiwire.NotificationContentResult(n.ID, normalizeNotificationBodyText(n.Body), n.Title, n.PostingAt, n.LastUpdatedAt, n.TabCategory, n.Category, n.BannerPath)
}

func notificationResultContentFallback(id int64) []any {
	if id <= 0 {
		id = time.Now().Unix()
	}
	now := time.Now()
	return apiwire.NotificationContentResult(id, normalizeNotificationBodyText(""), "お知らせ", now, now, 2, 1, "")
}

func normalizeNotificationBodyText(body string) string {
	if body == "" {
		return `[{\"Element\":4,\"Data\":\"お知らせ内容です。\"}]`
	}
	return body
}

func (s *NotificationService) List() []Notification {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]Notification, 0, len(s.notifications))
	for _, n := range s.notifications {
		out = append(out, n)
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].Order == out[j].Order {
			return out[i].PostingAt.After(out[j].PostingAt)
		}
		return out[i].Order < out[j].Order
	})

	return out
}

func (s *NotificationService) Get(id int64) (Notification, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	n, ok := s.notifications[id]
	return n, ok
}

func (s *NotificationService) First() (Notification, bool) {
	items := s.List()
	if len(items) == 0 {
		return Notification{}, false
	}
	return items[0], true
}

func (s *NotificationService) MarkRead(tabCategory int64, readAt time.Time) {
	if tabCategory <= 0 {
		return
	}
	if readAt.IsZero() {
		readAt = time.Now()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.readAtByTab[tabCategory] = readAt
}
