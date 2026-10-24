package services

import (
	"context"
	"time"

	"ymst_server/apiwire"
	"ymst_server/repositories"
)

type PlayerService struct{ users *repositories.UserRepository }

func NewPlayerService(users *repositories.UserRepository) *PlayerService {
	return &PlayerService{users: users}
}
func (s *PlayerService) UpdateTutorial(ctx context.Context, userKey string, status int64) ([]any, error) {
	return s.users.WithUserKey(userKey).UpdateTutorialStatusContext(ctx, status)
}
func (s *PlayerService) UpdateGameHintRead(ctx context.Context, userKey string, categories []int64) ([]any, error) {
	return s.users.WithUserKey(userKey).MarkGameHintsReadContext(ctx, categories)
}
func (s *PlayerService) UpdateSplashLastDisplayTime(ctx context.Context, userKey string, t time.Time) ([]any, error) {
	return s.users.WithUserKey(userKey).UpdateSplashLastDisplayedAtContext(ctx, t)
}

type ShopService struct{ users *repositories.UserRepository }

func NewShopService(users *repositories.UserRepository) *ShopService {
	return &ShopService{users: users}
}
func (s *ShopService) ViewPage(ctx context.Context, userKey string, category int64, exchangeShopMasterID *int64) ([]any, []any, error) {
	updates, err := s.users.WithUserKey(userKey).MarkShopViewedContext(ctx, category, exchangeShopMasterID)
	if err != nil {
		return nil, nil, err
	}
	return apiwire.ViewShopResult([]any{}), updates, nil
}
