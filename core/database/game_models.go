package database

import (
	"time"

	"gorm.io/gorm"
)

type GameUser struct {
	ID                             int64 `gorm:"primaryKey"`
	PlayerRank                     int
	CurrentRankPoint               int
	CurrentStamina                 int
	MaxStaminaRestoredAt           time.Time
	PaidJewel                      int
	FreeJewel                      int
	Coin                           int
	PlayerRankLimit                int
	StaminaRecoverTimesWithJewel   int
	CircleUsageRestrictionsEndTime time.Time
	CircleID                       string
	GameStartAt                    time.Time
	HashUserID                     string
	BanLevel                       int
	TutorialStatus                 int
	MonthlyPayment                 int
	SplashLastDisplayedAt          time.Time
	IsCapedPlayerRank              bool
	RequireCapedPlayerRankAnnounce bool
	ObjectIndex                    int `gorm:"index"`
	CreatedAt                      time.Time
	UpdatedAt                      time.Time
}

type GameUserProfile struct {
	ID                     int64 `gorm:"primaryKey"`
	Name                   string
	Introduction           string
	MainUCharacterID       int64
	MNameplateID           *int64
	MNameColorID           int64
	MTrophyID1             *int64
	MTrophyID2             *int64
	MTrophyID3             *int64
	PlayerRate             float64
	IsPublicPlayerRate     bool
	LeagueClass            int
	TotalSPCount           int
	IsPublicAlbumMainPage  bool
	MNameplateDetailID     *int64
	MainCharacterMasterID  int64
	DisplayAwakeningStatus bool
	IsPublicActivityLog    bool
	NameBaseColorMasterID  int64
	IconFrameMasterID      int64
	HomeSkinMasterID       int64
	ObjectIndex            int `gorm:"index"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

type GameUserPreference struct {
	ID           int64 `gorm:"primaryKey"`
	MultiPartyID int64
	BirthDate    *time.Time
	ObjectIndex  int `gorm:"index"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type GameHomeDisplayPreference struct {
	ID                              int64 `gorm:"primaryKey"`
	HomeCharacterBaseMasterID       *int64
	MemberCharacterBaseMasterID     *int64
	StoryCharacterBaseMasterID      *int64
	ShopCharacterBaseMasterID       *int64
	HomeCostumeMasterID             *int64
	MemberCostumeMasterID           *int64
	StoryCostumeMasterID            *int64
	ShopCostumeMasterID             *int64
	IllustCharacterMasterID         int64
	DisplayAwakeningStatus          bool
	HomeCharacterDisplayType        int
	LoginBonusCharacterBaseMasterID *int64
	LoginBonusCostumeMasterID       *int64
	ObjectIndex                     int `gorm:"index"`
	CreatedAt                       time.Time
	UpdatedAt                       time.Time
}

type GameCharacter struct {
	ID                       int64 `gorm:"primaryKey"`
	CharacterMasterID        int64 `gorm:"index"`
	Level                    int
	CurrentExperience        int
	TalentStage              int
	AwakeningPhase           int
	CharacterBaseID          int64 `gorm:"index"`
	SenseLevel               int
	ReadEpisodeOrder         int
	ReleasedEpisodeOrder     int
	DisplayAwakeningStatus   bool
	SecondaryCharacterBaseID *int64
	SecondarySenseLevel      int
	SelectionType            int
	IsFavorite               bool
	ObjectIndex              int `gorm:"index"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

type GameCharacterBase struct {
	ID                           int64 `gorm:"primaryKey"`
	CharacterBaseMasterID        int64 `gorm:"index"`
	StarRank                     int
	TotalStarPoint               int
	CostumeMasterID              *int64
	KeyMissionLevel              int
	PortalCharacterID            int64
	PortalDisplayAwakeningStatus bool
	ObjectIndex                  int `gorm:"index"`
	CreatedAt                    time.Time
	UpdatedAt                    time.Time
}

type GameParty struct {
	ID             int64 `gorm:"primaryKey"`
	Order          int
	Name           string
	LeaderPosition int
	ObjectIndex    int `gorm:"index"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type GamePartySlot struct {
	ID                      int64 `gorm:"primaryKey"`
	PartyID                 int64 `gorm:"index"`
	Position                int
	CharacterID             int64 `gorm:"index"`
	PosterID                *int64
	AccessoryID             *int64
	BonusAbilityEnableFlags int
	ObjectIndex             int `gorm:"index"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

type GameHint struct {
	ID             int64 `gorm:"primaryKey"`
	PageCategory   int64 `gorm:"uniqueIndex"`
	HasAlreadyRead bool
	ObjectIndex    int `gorm:"index"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type GameViewedShop struct {
	ID                   int64 `gorm:"primaryKey"`
	ExchangeShopMasterID *int64
	LastViewedAt         time.Time
	ViewedShopCategory   int64 `gorm:"index"`
	ObjectIndex          int   `gorm:"index"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type GameUserBonus struct {
	ID                  int64 `gorm:"primaryKey"`
	ExperienceBonus     float64
	LessonStarRankBonus float64
	ObjectIndex         int `gorm:"index"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type GameCurrency struct {
	ID          int64 `gorm:"primaryKey"`
	Coin        int
	FreeJewel   int
	PaidJewel   int
	ObjectIndex int `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GameRestriction struct {
	ID                             uint  `gorm:"primaryKey"`
	RestrictionID                  int64 `gorm:"index"`
	MultiLiveRestrictionFinishedAt *time.Time
	ReadMultiLiveRestrictionDialog *bool
	ObjectIndex                    int `gorm:"index"`
	CreatedAt                      time.Time
	UpdatedAt                      time.Time
}

type GameAlbumPreset struct {
	ID          int64 `gorm:"primaryKey"`
	Name        string
	Order       int
	ObjectIndex int `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GameFriendInvitation struct {
	ID                          int64 `gorm:"primaryKey"`
	InvitationCode              string
	HasInputOtherInvitationCode bool
	ObjectIndex                 int `gorm:"index"`
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}

type GameFriendInvitationMission struct {
	ID                                   int64 `gorm:"primaryKey"`
	FriendInvitationMissionMasterID      int64
	FriendInvitationMissionStageMasterID int64
	CurrentCount                         int64
	IsCleared                            bool
	IsRewardReceived                     bool
	ObjectIndex                          int `gorm:"index"`
	CreatedAt                            time.Time
	UpdatedAt                            time.Time
}

type GameRawDataObject struct {
	ID          uint      `gorm:"primaryKey"`
	UserKey     string    `gorm:"size:64;not null;uniqueIndex:idx_game_raw_object_key;index"`
	ObjectKey   string    `gorm:"size:128;not null;uniqueIndex:idx_game_raw_object_key"`
	UnionID     int64     `gorm:"not null;index"`
	ObjectIndex int       `gorm:"not null;index"`
	DataJSON    string    `gorm:"type:text"`
	IsNull      bool      `gorm:"not null;default:false"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

type GameUserStateMeta struct {
	ID            uint      `gorm:"primaryKey"`
	UserKey       string    `gorm:"size:64;not null;uniqueIndex"`
	SeedMode      string    `gorm:"size:32;not null"`
	StateVersion  int64     `gorm:"not null;default:1"`
	ObjectCount   int       `gorm:"not null;default:0"`
	LastBuiltAt   time.Time `gorm:"not null"`
	LastUpdatedAt time.Time `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type GameActiveLive struct {
	ID                      int64  `gorm:"primaryKey"`
	UserKey                 string `gorm:"size:64;not null;index"`
	PartyID                 int64  `gorm:"not null;index"`
	LiveMasterID            int64  `gorm:"not null;index"`
	LiveSettingMasterID     int64
	AuditionMasterID        *int64
	LeagueMasterID          *int64
	UseStamina              bool
	StaminaConsumptionRatio int
	IsAutoPlay              bool
	Status                  int `gorm:"not null;default:0;index"`
	StartedAt               time.Time
	FinishedAt              *time.Time
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

type GameConfirmationCode struct {
	ID        uint      `gorm:"primaryKey"`
	UserKey   string    `gorm:"size:64;not null;index"`
	Code      string    `gorm:"size:32;not null;index"`
	Purpose   string    `gorm:"size:32;not null;default:'account'"`
	ExpiresAt time.Time `gorm:"not null;index"`
	UsedAt    *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GameAccount struct {
	ID                  uint   `gorm:"primaryKey"`
	UserKey             string `gorm:"size:64;not null;uniqueIndex"`
	LoginToken          string `gorm:"size:128;not null;uniqueIndex"`
	Name                string `gorm:"size:64;not null"`
	Status              int    `gorm:"not null;default:0"`
	BanLevel            int    `gorm:"not null;default:0"`
	WarnedUntilUnix     *int64
	LastAuthenticatedAt *time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type GameSession struct {
	ID           uint   `gorm:"primaryKey"`
	UserKey      string `gorm:"size:64;not null;index"`
	SessionToken string `gorm:"size:128;not null;uniqueIndex"`
	ExpiresAt    *time.Time
	LastSeenAt   time.Time
	RevokedAt    *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type GameNotification struct {
	ID             uint      `gorm:"primaryKey"`
	NotificationID int64     `gorm:"not null;uniqueIndex"`
	Title          string    `gorm:"size:256;not null"`
	Body           string    `gorm:"type:text"`
	BannerPath     string    `gorm:"size:512"`
	TabCategory    int       `gorm:"not null;default:0"`
	Category       int       `gorm:"not null;default:0"`
	IsConfirmation bool      `gorm:"not null;default:false"`
	IsAnonymous    bool      `gorm:"not null;default:false"`
	SortOrder      int       `gorm:"not null;default:0"`
	PostingAt      time.Time `gorm:"not null"`
	LastUpdatedAt  time.Time `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type GameNotificationReadState struct {
	ID          uint      `gorm:"primaryKey"`
	UserKey     string    `gorm:"size:64;not null;uniqueIndex:idx_notice_read"`
	TabCategory int       `gorm:"not null;uniqueIndex:idx_notice_read"`
	ReadAt      time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GameInboxPackage struct {
	ID          uint   `gorm:"primaryKey"`
	UserKey     string `gorm:"size:64;not null;index;uniqueIndex:idx_inbox_package"`
	PackageKey  string `gorm:"size:128;not null;uniqueIndex:idx_inbox_package"`
	Title       string `gorm:"size:256"`
	Message     string `gorm:"type:text"`
	SenderName  string `gorm:"size:128"`
	PayloadJSON string `gorm:"type:text"`
	ReceivedAt  *time.Time
	ExpiresAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GameFriendRequest struct {
	ID          uint   `gorm:"primaryKey"`
	RequestKey  string `gorm:"size:128;not null;uniqueIndex"`
	FromUserKey string `gorm:"size:64;not null;index"`
	ToUserKey   string `gorm:"size:64;not null;index"`
	Status      int    `gorm:"not null;default:0"`
	Message     string `gorm:"size:256"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GameFriendRelation struct {
	ID            uint   `gorm:"primaryKey"`
	UserKey       string `gorm:"size:64;not null;uniqueIndex:idx_friend_relation"`
	FriendUserKey string `gorm:"size:64;not null;uniqueIndex:idx_friend_relation"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type GameCircleInvite struct {
	ID         uint   `gorm:"primaryKey"`
	InviteKey  string `gorm:"size:128;not null;uniqueIndex"`
	UserKey    string `gorm:"size:64;not null;index"`
	CircleID   int64  `gorm:"not null;index"`
	CircleName string `gorm:"size:128;not null"`
	Status     int    `gorm:"not null;default:0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type GameCircle struct {
	ID           uint   `gorm:"primaryKey"`
	CircleID     int64  `gorm:"not null;uniqueIndex"`
	CircleName   string `gorm:"size:128;not null"`
	OwnerUserKey string `gorm:"size:64;not null;index"`
	Description  string `gorm:"size:512"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type GameCircleMember struct {
	ID        uint      `gorm:"primaryKey"`
	CircleID  int64     `gorm:"not null;uniqueIndex:idx_circle_member"`
	UserKey   string    `gorm:"size:64;not null;uniqueIndex:idx_circle_member"`
	Role      int       `gorm:"not null;default:0"`
	JoinedAt  time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GameSupportCompanyState struct {
	ID                  uint   `gorm:"primaryKey"`
	UserKey             string `gorm:"size:64;not null;uniqueIndex:idx_support_company"`
	Company             int    `gorm:"not null;uniqueIndex:idx_support_company"`
	Level               int64  `gorm:"not null;default:1"`
	CurrentSupportPoint int64  `gorm:"not null;default:0"`
	LastLevelUppedAt    *time.Time
	LevelLimit          int64 `gorm:"not null;default:0"`
	LevelLimitStatus    int64 `gorm:"not null;default:0"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
