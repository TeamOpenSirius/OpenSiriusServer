package database

import (
	"time"

	"gorm.io/gorm"
)

const (
	UserStateRegistered = 1
	UserStateVerified   = 2
	UserStateBanned     = 3
	UserStateDeleted    = 4
)

const (
	UserRoleUser       = 1
	UserRoleAdmin      = 2
	UserRoleSuperAdmin = 3
)

const (
	GenderUnknown = 0
	GenderMale    = 1
	GenderFemale  = 2
	GenderOther   = 3
)

const (
	UserTokenTypeEmailVerify   = 1
	UserTokenTypeResetPassword = 2
	UserTokenTypeChangeEmail   = 3
)

type User struct {
	ID            uint           `gorm:"primaryKey"`
	State         int            `gorm:"not null;default:1;index"`
	Role          int            `gorm:"not null;default:1;index"`
	Username      string         `gorm:"size:64;not null;uniqueIndex"`
	Nickname      string         `gorm:"size:128;not null"`
	PasswordHash  string         `gorm:"size:255;not null"`
	Avatar        string         `gorm:"size:1024"`
	Gender        int            `gorm:"not null;default:0"`
	Email         string         `gorm:"size:255;not null;uniqueIndex"`
	Bio           string         `gorm:"type:text"`
	Website       string         `gorm:"size:255"`
	RegisterDate  time.Time      `gorm:"not null;index"`
	LastLogin     *time.Time     `gorm:"index"`
	LastLoginIP   string         `gorm:"column:last_login_ip;size:64"`
	EmailVerified bool           `gorm:"not null;default:false"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type UserSession struct {
	ID               uint       `gorm:"primaryKey"`
	UserID           uint       `gorm:"not null;index"`
	TokenHash        string     `gorm:"size:255;not null;index"`
	RefreshTokenHash string     `gorm:"size:255;not null;index"`
	DeviceInfo       string     `gorm:"type:text"`
	IP               string     `gorm:"column:ip;size:64"`
	UserAgent        string     `gorm:"type:text"`
	ExpiresAt        time.Time  `gorm:"not null;index"`
	CreatedAt        time.Time  `gorm:"not null"`
	LastUsedAt       *time.Time `gorm:"index"`
	RevokedAt        *time.Time `gorm:"index"`
}

type UserToken struct {
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `gorm:"not null;index"`
	Type      int        `gorm:"not null;index"`
	TokenHash string     `gorm:"size:255;not null;uniqueIndex"`
	ExpiresAt time.Time  `gorm:"not null;index"`
	UsedAt    *time.Time `gorm:"index"`
	CreatedAt time.Time  `gorm:"not null"`
}

func AutoMigrate(db *gorm.DB) error {
	// Older dev schemas briefly had a unique index on ViewedShopCategory,
	// but captured data can contain multiple rows with the same category.
	if db.Migrator().HasIndex(&GameViewedShop{}, "idx_game_viewed_shops_viewed_shop_category") {
		if err := db.Migrator().DropIndex(&GameViewedShop{}, "idx_game_viewed_shops_viewed_shop_category"); err != nil {
			return err
		}
	}
	// The early GameRestriction schema used the game Id as the SQL primary key,
	// but captured data can contain multiple Restriction objects with the same
	// game Id. Recreate this table so SQLite does not keep the old primary key.
	if db.Migrator().HasTable(&GameRestriction{}) {
		if err := db.Migrator().DropTable(&GameRestriction{}); err != nil {
			return err
		}
	}
	// Stage 3 briefly used a package_key-only unique index. Inbox packages are
	// user-scoped, so rebuild it as (user_key, package_key).
	if db.Migrator().HasIndex(&GameInboxPackage{}, "idx_inbox_package") {
		if err := db.Migrator().DropIndex(&GameInboxPackage{}, "idx_inbox_package"); err != nil {
			return err
		}
	}

	return db.AutoMigrate(
		&User{},
		&UserSession{},
		&UserToken{},
		&GameUser{},
		&GameUserProfile{},
		&GameUserPreference{},
		&GameHomeDisplayPreference{},
		&GameCharacter{},
		&GameCharacterBase{},
		&GameParty{},
		&GamePartySlot{},
		&GameHint{},
		&GameViewedShop{},
		&GameUserBonus{},
		&GameCurrency{},
		&GameRestriction{},
		&GameAlbumPreset{},
		&GameFriendInvitation{},
		&GameFriendInvitationMission{},
		&GameRawDataObject{},
		&GameUserStateMeta{},
		&GameActiveLive{},
		&GameConfirmationCode{},
		&GameAccount{},
		&GameSession{},
		&GameNotification{},
		&GameNotificationReadState{},
		&GameInboxPackage{},
		&GameFriendRequest{},
		&GameFriendRelation{},
		&GameCircleInvite{},
		&GameCircle{},
		&GameCircleMember{},
		&GameSupportCompanyState{},
	)
}
