package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"ymst_server/apiwire"
	"ymst_server/repositories"
)

type AccountService struct {
	accounts *repositories.AccountRepository
	sessions *repositories.SessionRepository
	users    *repositories.UserRepository
}

type RegisterResult struct {
	LoginToken string
	ErrorType  int64
}

type AuthenticateResult struct {
	SessionToken string
	BanLevel     int64
	WarnedUntil  any
}

func NewAccountService(accounts *repositories.AccountRepository, sessions *repositories.SessionRepository, users *repositories.UserRepository) *AccountService {
	return &AccountService{accounts: accounts, sessions: sessions, users: users}
}

func (s *AccountService) Register(ctx context.Context, name string) (*RegisterResult, error) {
	if s == nil || s.accounts == nil || s.users == nil {
		return nil, errors.New("account service not initialized")
	}
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Player"
	}
	account, err := s.accounts.CreateAccount(ctx, name)
	if err != nil {
		return nil, err
	}
	if err := s.users.CloneSeedToUser(account.UserKey, func(payload []any) []any { return patchRegisteredUserName(payload, name) }); err != nil {
		return nil, err
	}
	return &RegisterResult{LoginToken: account.LoginToken, ErrorType: 0}, nil
}

func (s *AccountService) Authenticate(ctx context.Context, payload apiwire.AuthenticatePayload) (*AuthenticateResult, error) {
	if s == nil || s.accounts == nil || s.sessions == nil {
		return nil, errors.New("account service not initialized")
	}
	loginToken := strings.TrimSpace(payload.LoginToken)
	if loginToken == "" {
		return nil, gorm.ErrRecordNotFound
	}
	account, err := s.accounts.FindByLoginToken(ctx, loginToken)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		// Temporary compatibility path: older/local clients may present a
		// login_token that predates GameAccount persistence. Provision an
		// isolated account for that token, then continue with normal session auth.
		account, err = s.accounts.CreateAccountWithLoginToken(ctx, "Player", loginToken)
		if err != nil {
			return nil, err
		}
		if s.users != nil {
			if err := s.users.CloneSeedToUser(account.UserKey, func(payload []any) []any { return patchRegisteredUserName(payload, account.Name) }); err != nil {
				return nil, err
			}
		}
	}
	if account.Status == repositories.GameAccountStatusDeleted {
		return nil, gorm.ErrRecordNotFound
	}
	if account.Status == repositories.GameAccountStatusBanned {
		return nil, errors.New("account banned")
	}
	session, err := s.sessions.CreateSession(ctx, account.UserKey, 24*time.Hour)
	if err != nil {
		return nil, err
	}
	_ = s.accounts.MarkAuthenticated(ctx, account.UserKey, time.Now().UTC())
	var warnedUntil any
	if account.WarnedUntilUnix != nil {
		warnedUntil = *account.WarnedUntilUnix
	}
	return &AuthenticateResult{SessionToken: session.SessionToken, BanLevel: int64(account.BanLevel), WarnedUntil: warnedUntil}, nil
}

func patchRegisteredUserName(payload []any, name string) []any {
	for _, raw := range payload {
		obj, ok := raw.([]any)
		if !ok || len(obj) < 2 || apiwire.Int64Value(obj[0]) != repositories.UnionUserProfile {
			continue
		}
		arr, ok := obj[1].([]any)
		if !ok || len(arr) < 2 {
			continue
		}
		arr[1] = name
	}
	return payload
}
