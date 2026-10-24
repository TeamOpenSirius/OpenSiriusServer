package handlers

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"ymst_server/api"
	"ymst_server/apiwire"
	"ymst_server/core/config"
	"ymst_server/core/database"
	"ymst_server/core/logger"
	"ymst_server/dto"
	"ymst_server/middleware"
	"ymst_server/services"
)

type AccountHandler struct {
	Log       *logger.Logger
	Game      config.GameConfig
	DB        *gorm.DB
	Secret    []byte
	TokenUser string
	Accounts  *services.AccountService
	Extra     *services.AccountExtraService
}

const defaultGameUserKeyForAccount = "default"
const accountTokenStorePath = "data/account_tokens.json"

func NewAccountHandler(log *logger.Logger, game config.GameConfig, db *gorm.DB, jwtSecret string, accounts *services.AccountService, extra *services.AccountExtraService) *AccountHandler {
	if jwtSecret == "" {
		jwtSecret = "change-me"
	}
	return &AccountHandler{
		Log:       log,
		Game:      game,
		DB:        db,
		Secret:    []byte(jwtSecret),
		TokenUser: "1",
		Accounts:  accounts,
		Extra:     extra,
	}
}

func (h *AccountHandler) RegisterRoutes(r gin.IRouter) {
	r.POST("/api/Account/Register", h.Register)
	r.POST("/api/Account/Authenticate", h.Authenticate)
	r.POST("/api/Account/GetConfirmationCode", h.GetConfirmationCode)
	r.POST("/api/Account/Delete", h.DeleteAccount)
}

func (h *AccountHandler) Authenticate(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	values, err := apiwire.DecodeRequestArray(body)
	if err != nil && len(body) > 0 {
		if h.Log != nil {
			h.Log.GetZap().Warn("decode authenticate payload failed",
				zap.Error(err),
				zap.Int("body_size", len(body)),
			)
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	payload := apiwire.DecodeAuthenticatePayload(values)
	payload.LoginToken = strings.TrimSpace(payload.LoginToken)

	if strings.TrimSpace(payload.LoginToken) == "" || h.Accounts == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("game account authenticate",
			zap.Int("body_size", len(body)),
			zap.String("login_token", payload.LoginToken),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("platform", c.GetHeader("X-Platform")),
		)
	}

	result, err := h.Accounts.Authenticate(c.Request.Context(), payload)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Warn("account authenticate failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := api.WriteApiMsgpackResult(c, apiwire.AuthenticateResult(result.SessionToken, result.BanLevel, result.WarnedUntil)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *AccountHandler) Register(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	values, err := apiwire.DecodeRequestArray(body)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Warn("decode register payload failed",
				zap.Error(err),
				zap.Int("body_size", len(body)),
			)
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	payload := apiwire.DecodeRegisterPayload(values)
	payload.Name = strings.TrimSpace(payload.Name)

	if h.Accounts == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	result, err := h.Accounts.Register(c.Request.Context(), payload.Name)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("register account failed",
				zap.Error(err),
				zap.String("name", payload.Name),
			)
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("account register",
			zap.String("name", payload.Name),
			zap.String("token", result.LoginToken),
			zap.Int("error_type", int(result.ErrorType)),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := api.WriteApiMsgpackResult(c, apiwire.AccountRegistResult(result.LoginToken, result.ErrorType)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *AccountHandler) makeGameToken() (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"uid": h.TokenUser,
		"lc":  "1",
		"pf":  "GooglePlay",
		"gv":  "GooglePlay",
		"ld":  now.Format("01/02/2006 15:04:05"),
		"nbf": now.Unix(),
		"exp": now.Add(24 * time.Hour).Unix(),
		"iat": now.Unix(),
		"iss": "sirius.kms3.com",
		"aud": "sirius",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(h.Secret)
}

func (h *AccountHandler) GetConfirmationCode(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)

	userKey, _ := middleware.UserKeyFromContext(c)
	code, remaining, err := h.Extra.GetConfirmationCode(c.Request.Context(), userKey)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Error("get confirmation code failed", zap.Error(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("account get confirmation code",
			zap.String("user_key", userKey),
			zap.String("code", code),
			zap.Int64("remaining_seconds", remaining),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", c.GetHeader("Authorization") != ""),
		)
	}

	writeGameHeaders(c, h.Game)

	if err := apiwire.WritePayload(c, apiwire.TimedConfirmationCode(code, int64(remaining))); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *AccountHandler) getOrCreateConfirmationCode(userKey string) (string, int, error) {
	now := time.Now().UTC()
	ttl := 10 * time.Minute

	if h.DB != nil {
		var existing database.GameConfirmationCode
		err := h.DB.
			Where("user_key = ? AND used_at IS NULL AND expires_at > ?", userKey, now).
			Order("expires_at DESC").
			First(&existing).Error
		if err == nil {
			remaining := int(existing.ExpiresAt.Sub(now).Seconds())
			if remaining > 30 {
				return existing.Code, remaining, nil
			}
		} else if err != gorm.ErrRecordNotFound {
			return "", 0, err
		}

		code, err := generateConfirmationCode8()
		if err != nil {
			return "", 0, err
		}

		row := &database.GameConfirmationCode{
			UserKey:   userKey,
			Code:      code,
			Purpose:   "account",
			ExpiresAt: now.Add(ttl),
		}
		if err := h.DB.Create(row).Error; err != nil {
			return "", 0, err
		}

		return code, int(ttl.Seconds()), nil
	}

	return "12345678", int(ttl.Seconds()), nil
}

func (h *AccountHandler) registerAccount(name string) (string, dto.AccountRegisterErrorTypes, error) {
	if strings.TrimSpace(name) == "" {
		name = "Player"
	}

	token, err := generateLoginToken()
	if err != nil {
		return "", dto.AccountRegisterErrorTypesNone, err
	}

	store, err := loadAccountTokenStore()
	if err != nil {
		return "", dto.AccountRegisterErrorTypesNone, err
	}
	store.Tokens[token] = accountTokenRecord{
		UserKey:    defaultGameUserKeyForAccount,
		Name:       name,
		Token:      token,
		CreatedAt:  time.Now().UTC(),
		LastUsedAt: nil,
	}
	if err := saveAccountTokenStore(store); err != nil {
		return "", dto.AccountRegisterErrorTypesNone, err
	}

	return token, dto.AccountRegisterErrorTypesNone, nil
}

func generateLoginToken() (string, error) {
	buf := make([]byte, 24)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return "login_" + hex.EncodeToString(buf), nil
}

type accountTokenStore struct {
	Tokens map[string]accountTokenRecord `json:"tokens"`
}

type accountTokenRecord struct {
	UserKey    string     `json:"user_key"`
	Name       string     `json:"name"`
	Token      string     `json:"token"`
	CreatedAt  time.Time  `json:"created_at"`
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
}

func loadAccountTokenStore() (accountTokenStore, error) {
	store := accountTokenStore{Tokens: map[string]accountTokenRecord{}}
	body, err := os.ReadFile(accountTokenStorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil
		}
		return store, err
	}
	if len(bytes.TrimSpace(body)) == 0 {
		return store, nil
	}
	if err := json.Unmarshal(body, &store); err != nil {
		return store, err
	}
	if store.Tokens == nil {
		store.Tokens = map[string]accountTokenRecord{}
	}
	return store, nil
}

func saveAccountTokenStore(store accountTokenStore) error {
	if err := os.MkdirAll(filepath.Dir(accountTokenStorePath), 0755); err != nil {
		return err
	}
	body, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(accountTokenStorePath, body, 0644)
}

func (h *AccountHandler) validateLoginToken(token string) (bool, error) {
	store, err := loadAccountTokenStore()
	if err != nil {
		return false, err
	}
	record, ok := store.Tokens[token]
	if !ok {
		return false, nil
	}
	now := time.Now().UTC()
	record.LastUsedAt = &now
	store.Tokens[token] = record
	if err := saveAccountTokenStore(store); err != nil {
		return false, err
	}
	return true, nil
}

func generateConfirmationCode8() (string, error) {
	const digits = "0123456789"

	buf := make([]byte, 8)
	for i := range buf {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		buf[i] = digits[n.Int64()]
	}

	return string(buf), nil
}

func writeTimedConfirmationCodeResult(c *gin.Context, result dto.TimedConfirmationCode) error {
	var body bytes.Buffer
	enc := msgpack.NewEncoder(&body)

	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := encodeTimedConfirmationCode(enc, result); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}

	c.Data(http.StatusOK, api.MsgpackContentType, body.Bytes())
	return nil
}

func writeAccountRegisterResult(c *gin.Context, result dto.AccountRegistResult) error {
	var body bytes.Buffer
	enc := msgpack.NewEncoder(&body)

	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := encodeAccountRegistResult(enc, result); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}

	c.Data(http.StatusOK, api.MsgpackContentType, body.Bytes())
	return nil
}

func encodeAccountRegistResult(enc *msgpack.Encoder, result dto.AccountRegistResult) error {
	if err := enc.EncodeArrayLen(2); err != nil {
		return err
	}
	if err := enc.EncodeString(result.Token); err != nil {
		return err
	}
	return enc.EncodeInt(int64(result.ErrorType))
}

func encodeTimedConfirmationCode(enc *msgpack.Encoder, result dto.TimedConfirmationCode) error {
	if err := enc.EncodeArrayLen(2); err != nil {
		return err
	}
	if err := enc.EncodeString(result.ConfirmationCode); err != nil {
		return err
	}
	return enc.EncodeInt(int64(result.RemainingSeconds))
}

func int64ValueForAccount(value any) int64 {
	switch v := value.(type) {
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case int:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case uint:
		return int64(v)
	default:
		return 0
	}
}

func stringValueForAccount(value any) string {
	v, _ := value.(string)
	return v
}

func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	userKey, ok := middleware.UserKeyFromContext(c)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	success, errorType, err := h.Extra.DeleteAccount(c.Request.Context(), userKey)
	if err != nil {
		if h.Log != nil {
			h.Log.GetZap().Warn("mark account delete requested failed",
				zap.Error(err),
				zap.String("user_key", userKey),
				zap.Int("body_size", len(body)),
			)
		}

		writeGameHeaders(c, h.Game)
		if err := apiwire.WritePayload(c, apiwire.AccountDeletionResult(false, int64(1))); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	if h.Log != nil {
		h.Log.GetZap().Info("account delete",
			zap.String("user_key", userKey),
			zap.Int("body_size", len(body)),
			zap.Bool("has_authorization", strings.TrimSpace(c.GetHeader("Authorization")) != ""),
		)
	}

	writeGameHeaders(c, h.Game)
	if err := apiwire.WritePayload(c, apiwire.AccountDeletionResult(success, errorType)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (h *AccountHandler) markAccountDeleteRequested(userKey string) error {
	if h.Accounts == nil {
		return nil
	}
	_ = userKey
	return nil
}

func ptrInt(v int) *int {
	return &v
}

func writeAccountDeletionResult(c *gin.Context, success bool, errType *int) error {
	var body bytes.Buffer
	enc := msgpack.NewEncoder(&body)

	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := encodeAccountDeletionResult(enc, success, errType); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}
	if err := enc.EncodeArrayLen(0); err != nil {
		return err
	}

	c.Data(http.StatusOK, api.MsgpackContentType, body.Bytes())
	return nil
}

func encodeAccountDeletionResult(enc *msgpack.Encoder, success bool, errType *int) error {
	if err := enc.EncodeArrayLen(2); err != nil {
		return err
	}
	if err := enc.EncodeBool(success); err != nil {
		return err
	}
	if errType == nil {
		return enc.EncodeNil()
	}
	return enc.EncodeInt(int64(*errType))
}
