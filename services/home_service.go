package services

import (
	"context"
	"time"

	"ymst_server/apiwire"
	"ymst_server/repositories"
)

type LoginService struct {
	accounts *repositories.AccountRepository
	users    *repositories.UserRepository
}

func NewLoginService(accounts *repositories.AccountRepository, users ...*repositories.UserRepository) *LoginService {
	var userRepo *repositories.UserRepository
	if len(users) > 0 {
		userRepo = users[0]
	}
	return &LoginService{accounts: accounts, users: userRepo}
}
func (s *LoginService) Login(ctx context.Context, userKey string, pushToken string) ([]any, []any, error) {
	if s != nil && s.accounts != nil && userKey != "" {
		_ = s.accounts.MarkAuthenticated(ctx, userKey, time.Now().UTC())
	}

	updates := []any{}
	if s != nil && s.users != nil && userKey != "" {
		objects, err := s.users.WithUserKey(userKey).BuildDataObjects(ctx)
		if err != nil {
			return nil, nil, err
		}
		updates = dataObjectUpdatesByUnion(objects, 0, 1, 2, 3)
	}

	return apiwire.EmptyLoginResult(), updates, nil
}

func dataObjectUpdatesByUnion(objects []any, unionIDs ...int64) []any {
	if len(objects) == 0 || len(unionIDs) == 0 {
		return []any{}
	}
	wanted := make(map[int64]struct{}, len(unionIDs))
	for _, unionID := range unionIDs {
		wanted[unionID] = struct{}{}
	}
	updates := make([]any, 0, len(unionIDs))
	for _, raw := range objects {
		object, ok := raw.([]any)
		if !ok || len(object) < 2 {
			continue
		}
		if _, ok := wanted[apiwire.Int64Value(object[0])]; ok {
			updates = append(updates, raw)
		}
	}
	return updates
}

type HomeService struct{}

func NewHomeService() *HomeService { return &HomeService{} }
func (s *HomeService) CheckReceiveLoginBonus(ctx context.Context, userKey string) ([]any, []any, error) {
	return []any{}, []any{}, nil
}
func (s *HomeService) CheckExternalPayment(ctx context.Context, userKey string) ([]any, []any, error) {
	return apiwire.ExternalPaymentResult([]any{}), []any{}, nil
}
