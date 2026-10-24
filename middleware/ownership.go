package middleware

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ymst_server/core/auth"
	"ymst_server/core/database"
)

// RequireUser allows any authenticated application user role to proceed.
func RequireUser() gin.HandlerFunc {
	return Authorize(auth.RoleAdmin, auth.RoleSubscriber, auth.RoleCreator)
}

// RequireSelfOrAdmin allows admins or the user identified by :id to proceed.
func RequireSelfOrAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if db == nil {
			abortWithError(c, http.StatusInternalServerError, "database unavailable")
			return
		}

		userID, role, ok := UserFromContext(c)
		if !ok {
			abortWithError(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		if role == auth.RoleAdmin {
			c.Next()
			return
		}

		targetID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			abortWithError(c, http.StatusBadRequest, "invalid id")
			return
		}
		currentUserID, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			abortWithError(c, http.StatusUnauthorized, "invalid user id")
			return
		}
		if currentUserID != targetID {
			abortWithError(c, http.StatusForbidden, "forbidden")
			return
		}

		var user database.User
		if err := db.WithContext(c.Request.Context()).Where("id = ?", targetID).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				abortWithError(c, http.StatusNotFound, "user not found")
				return
			}
			abortWithError(c, http.StatusInternalServerError, "permission check failed")
			return
		}
		c.Next()
	}
}
