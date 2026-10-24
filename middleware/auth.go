package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"ymst_server/core/auth"
	"ymst_server/repositories"
)

type contextKey string

const (
	contextUserID contextKey = "user_id"
	contextRole   contextKey = "role"
)

const (
	ContextUserKey      = "ymst_user_key"
	ContextSessionToken = "ymst_session_token"
)

// Authenticate validates access tokens and sets user info in context.
func Authenticate(authSvc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := extractBearer(c.GetHeader("Authorization"))
		if tokenStr == "" {
			abortWithError(c, http.StatusUnauthorized, "missing token")
			return
		}

		claims, err := authSvc.ParseAccess(tokenStr)
		if err != nil {
			status := http.StatusUnauthorized
			if errors.Is(err, auth.ErrExpiredToken) {
				status = http.StatusUnauthorized
			}
			abortWithError(c, status, err.Error())
			return
		}

		if store := authSvc.Store(); store != nil {
			if err := store.TouchSessionByAccessToken(tokenStr); err != nil {
				abortWithError(c, http.StatusUnauthorized, "invalid session")
				return
			}
		}

		c.Set(string(contextUserID), claims.Subject)
		c.Set(string(contextRole), claims.Role)
		c.Next()
	}
}

// Authorize checks if the user has one of the required roles.
func Authorize(roles ...string) gin.HandlerFunc {
	roleSet := make(map[string]struct{}, len(roles))
	for _, r := range roles {
		roleSet[r] = struct{}{}
	}

	return func(c *gin.Context) {
		roleVal, ok := c.Get(string(contextRole))
		if !ok {
			abortWithError(c, http.StatusForbidden, "missing role")
			return
		}
		role, _ := roleVal.(string)
		if _, exists := roleSet[role]; !exists {
			abortWithError(c, http.StatusForbidden, "forbidden")
			return
		}
		c.Next()
	}
}

// UserFromContext returns the user id and role if set.
func UserFromContext(c *gin.Context) (string, string, bool) {
	userID, okID := c.Get(string(contextUserID))
	role, okRole := c.Get(string(contextRole))
	uid, _ := userID.(string)
	rol, _ := role.(string)
	return uid, rol, okID && okRole
}

func extractBearer(header string) string {
	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 {
		return ""
	}
	if !strings.EqualFold(parts[0], "Bearer") {
		return ""
	}
	return strings.TrimSpace(parts[1])
}

func AuthMiddleware(sessions *repositories.SessionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessions == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token := extractGameSessionToken(c)
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		session, err := sessions.FindByToken(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		now := time.Now().UTC()
		if now.Sub(session.LastSeenAt) > 30*time.Second {
			_ = sessions.Touch(c.Request.Context(), token, now)
		}
		c.Set(ContextUserKey, session.UserKey)
		c.Set(ContextSessionToken, token)
		c.Next()
	}
}

func AuthMiddlewareWithPublicPaths(sessions *repositories.SessionRepository, publicPaths map[string]bool) gin.HandlerFunc {
	base := AuthMiddleware(sessions)
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if publicPaths[path] || !strings.HasPrefix(path, "/api/") || strings.HasPrefix(path, "/api/v1/") {
			c.Next()
			return
		}
		base(c)
	}
}

func UserKeyFromContext(c *gin.Context) (string, bool) {
	value, ok := c.Get(ContextUserKey)
	userKey, _ := value.(string)
	return userKey, ok && strings.TrimSpace(userKey) != ""
}

func MustUserKey(c *gin.Context) string {
	userKey, _ := UserKeyFromContext(c)
	return userKey
}

func extractGameSessionToken(c *gin.Context) string {
	authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
	if authHeader != "" {
		if token := extractBearer(authHeader); token != "" {
			return token
		}
		if !strings.Contains(authHeader, " ") {
			return authHeader
		}
	}
	if token := strings.TrimSpace(c.GetHeader("X-Session-Token")); token != "" {
		return token
	}
	if token := strings.TrimSpace(c.GetHeader("X-Auth-Token")); token != "" {
		return token
	}
	return ""
}
