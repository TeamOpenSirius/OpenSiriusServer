package services

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"gorm.io/gorm"

	"ymst_server/apiwire"
	"ymst_server/repositories"
)

const AccountDeleteErrorUnknown int64 = 1

type AccountExtraService struct {
	accounts *repositories.AccountRepository
	sessions *repositories.SessionRepository
	codes    *repositories.ConfirmationCodeRepository
}

func NewAccountExtraService(accounts *repositories.AccountRepository, sessions *repositories.SessionRepository, codes *repositories.ConfirmationCodeRepository) *AccountExtraService {
	return &AccountExtraService{accounts: accounts, sessions: sessions, codes: codes}
}
func (s *AccountExtraService) GetConfirmationCode(ctx context.Context, userKey string) (string, int64, error) {
	if s == nil || s.codes == nil {
		return "", 0, errors.New("account extra service not initialized")
	}
	if userKey == "" {
		userKey = "anonymous"
	}
	now := time.Now().UTC()
	ttl := 5 * time.Minute
	row, err := s.codes.FindActive(ctx, userKey, "account", now)
	if err == nil {
		return row.Code, int64(row.ExpiresAt.Sub(now).Seconds()), nil
	}
	if err != gorm.ErrRecordNotFound {
		return "", 0, err
	}
	code, err := generateDigits(6)
	if err != nil {
		return "", 0, err
	}
	row, err = s.codes.Create(ctx, userKey, "account", code, now.Add(ttl))
	if err != nil {
		return "", 0, err
	}
	return row.Code, int64(ttl.Seconds()), nil
}
func (s *AccountExtraService) DeleteAccount(ctx context.Context, userKey string) (bool, any, error) {
	if userKey == "" {
		return false, AccountDeleteErrorUnknown, nil
	}
	if s.accounts != nil {
		if err := s.accounts.MarkDeleted(ctx, userKey); err != nil {
			return false, AccountDeleteErrorUnknown, err
		}
	}
	if s.sessions != nil {
		if err := s.sessions.RevokeAllForUser(ctx, userKey); err != nil {
			return false, AccountDeleteErrorUnknown, err
		}
	}
	return true, nil, nil
}
func generateDigits(n int) (string, error) {
	const digits = "0123456789"
	out := make([]byte, n)
	for i := range out {
		v, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		out[i] = digits[v.Int64()]
	}
	return string(out), nil
}

func TimedConfirmationCodePayload(code string, remaining int64) []any {
	return apiwire.TimedConfirmationCode(code, remaining)
}
