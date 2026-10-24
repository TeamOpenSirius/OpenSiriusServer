package database

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"ymst_server/utils/cryptoutil"
)

func seedDefaultUsers(db *gorm.DB) error {
	if db == nil {
		return errors.New("database not initialized")
	}

	type seedUser struct {
		Username string
		Nickname string
		Email    string
		State    int
		Role     int
	}

	defaultUsers := []seedUser{
		{
			Username: "admin",
			Nickname: "Administrator",
			Email:    "admin@ymst.local",
			State:    UserStateVerified,
			Role:     UserRoleAdmin,
		},
		{
			Username: "user",
			Nickname: "Default User",
			Email:    "user@ymst.local",
			State:    UserStateVerified,
			Role:     UserRoleUser,
		},
	}

	passwordSHA256 := cryptoutil.SHA256Hex("123456")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordSHA256), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	for _, item := range defaultUsers {
		var existing User
		err := db.Where("username = ?", item.Username).First(&existing).Error
		if err == nil {
			updates := map[string]any{
				"state":      item.State,
				"role":       item.Role,
				"nickname":   item.Nickname,
				"email":      item.Email,
				"updated_at": now,
			}
			if bcrypt.CompareHashAndPassword([]byte(existing.PasswordHash), []byte(passwordSHA256)) != nil {
				updates["password_hash"] = string(hashedPassword)
			}
			if err := db.Model(&existing).Updates(updates).Error; err != nil {
				return err
			}
			continue
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		user := User{
			State:         item.State,
			Role:          item.Role,
			Username:      item.Username,
			Nickname:      item.Nickname,
			PasswordHash:  string(hashedPassword),
			Gender:        GenderUnknown,
			Email:         item.Email,
			RegisterDate:  now,
			EmailVerified: true,
			CreatedAt:     now,
			UpdatedAt:     now,
		}
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
