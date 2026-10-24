// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Player
package dto

import (
	"time"
)

type Currency struct {
	Id        int64 `json:"id" msgpack:"id"`               // Key: 0
	Coin      int   `json:"coin" msgpack:"coin"`           // Key: 1
	FreeJewel int   `json:"freeJewel" msgpack:"freeJewel"` // Key: 2
	PaidJewel int   `json:"paidJewel" msgpack:"paidJewel"` // Key: 3
}

type GameHint struct {
	Id             int64          `json:"id" msgpack:"id"`                         // Key: 0
	PageCategory   PageCategories `json:"pageCategory" msgpack:"pageCategory"`     // Key: 1
	HasAlreadyRead bool           `json:"hasAlreadyRead" msgpack:"hasAlreadyRead"` // Key: 2
}

type HomeDisplayPreference struct {
	Id                              int64                     `json:"id" msgpack:"id"`                                                           // Key: 0
	HomeCharacterBaseMasterId       *int64                    `json:"homeCharacterBaseMasterId" msgpack:"homeCharacterBaseMasterId"`             // Key: 1
	MemberCharacterBaseMasterId     *int64                    `json:"memberCharacterBaseMasterId" msgpack:"memberCharacterBaseMasterId"`         // Key: 3
	StoryCharacterBaseMasterId      *int64                    `json:"storyCharacterBaseMasterId" msgpack:"storyCharacterBaseMasterId"`           // Key: 4
	ShopCharacterBaseMasterId       *int64                    `json:"shopCharacterBaseMasterId" msgpack:"shopCharacterBaseMasterId"`             // Key: 5
	HomeCostumeMasterId             *int64                    `json:"homeCostumeMasterId" msgpack:"homeCostumeMasterId"`                         // Key: 6
	MemberCostumeMasterId           *int64                    `json:"memberCostumeMasterId" msgpack:"memberCostumeMasterId"`                     // Key: 8
	StoryCostumeMasterId            *int64                    `json:"storyCostumeMasterId" msgpack:"storyCostumeMasterId"`                       // Key: 9
	ShopCostumeMasterId             *int64                    `json:"shopCostumeMasterId" msgpack:"shopCostumeMasterId"`                         // Key: 10
	IllustCharacterMasterId         int64                     `json:"illustCharacterMasterId" msgpack:"illustCharacterMasterId"`                 // Key: 11
	DisplayAwakeningStatus          bool                      `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"`                   // Key: 12
	HomeCharacterDisplayType        HomeCharacterDisplayTypes `json:"homeCharacterDisplayType" msgpack:"homeCharacterDisplayType"`               // Key: 13
	LoginBonusCharacterBaseMasterId *int64                    `json:"loginBonusCharacterBaseMasterId" msgpack:"loginBonusCharacterBaseMasterId"` // Key: 14
	LoginBonusCostumeMasterId       *int64                    `json:"loginBonusCostumeMasterId" msgpack:"loginBonusCostumeMasterId"`             // Key: 15
}

type Limit struct {
	Id                                     int64 `json:"id" msgpack:"id"`                                                                         // Key: 0
	AdditionalAcquirablePhotoLimit         int   `json:"additionalAcquirablePhotoLimit" msgpack:"additionalAcquirablePhotoLimit"`                 // Key: 1
	AcquirablePhotoLimitIncreasedTimes     int   `json:"acquirablePhotoLimitIncreasedTimes" msgpack:"acquirablePhotoLimitIncreasedTimes"`         // Key: 2
	AdditionalAcquirableAccessoryLimit     int   `json:"additionalAcquirableAccessoryLimit" msgpack:"additionalAcquirableAccessoryLimit"`         // Key: 3
	AcquirableAccessoryLimitIncreasedTimes int   `json:"acquirableAccessoryLimitIncreasedTimes" msgpack:"acquirableAccessoryLimitIncreasedTimes"` // Key: 4
}

type User struct {
	Id                             int64          `json:"id" msgpack:"id"`                                                         // Key: 0
	PlayerRank                     int            `json:"playerRank" msgpack:"playerRank"`                                         // Key: 1
	CurrentRankPoint               int            `json:"currentRankPoint" msgpack:"currentRankPoint"`                             // Key: 2
	CurrentStamina                 int            `json:"currentStamina" msgpack:"currentStamina"`                                 // Key: 3
	MaxStaminaRestoredAt           time.Time      `json:"maxStaminaRestoredAt" msgpack:"maxStaminaRestoredAt"`                     // Key: 4
	PaidJewel                      int            `json:"paidJewel" msgpack:"paidJewel"`                                           // Key: 5
	FreeJewel                      int            `json:"freeJewel" msgpack:"freeJewel"`                                           // Key: 6
	Coin                           int            `json:"coin" msgpack:"coin"`                                                     // Key: 7
	PlayerRankLimit                int            `json:"playerRankLimit" msgpack:"playerRankLimit"`                               // Key: 8
	StaminaRecoverTimesWithJewel   int            `json:"staminaRecoverTimesWithJewel" msgpack:"staminaRecoverTimesWithJewel"`     // Key: 9
	CircleUsageRestrictionsEndTime time.Time      `json:"circleUsageRestrictionsEndTime" msgpack:"circleUsageRestrictionsEndTime"` // Key: 10
	CircleId                       string         `json:"circleId" msgpack:"circleId"`                                             // Key: 11
	GameStartAt                    time.Time      `json:"gameStartAt" msgpack:"gameStartAt"`                                       // Key: 12
	HashUserId                     string         `json:"hashUserId" msgpack:"hashUserId"`                                         // Key: 13
	BanLevel                       BanLevels      `json:"banLevel" msgpack:"banLevel"`                                             // Key: 14
	TutorialStatus                 TutorialStatus `json:"tutorialStatus" msgpack:"tutorialStatus"`                                 // Key: 15
	MonthlyPayment                 int            `json:"monthlyPayment" msgpack:"monthlyPayment"`                                 // Key: 16
	SplashLastDisplayedAt          time.Time      `json:"splashLastDisplayedAt" msgpack:"splashLastDisplayedAt"`                   // Key: 17
	IsCapedPlayerRank              bool           `json:"isCapedPlayerRank" msgpack:"isCapedPlayerRank"`                           // Key: 18
	RequireCapedPlayerRankAnnounce bool           `json:"requireCapedPlayerRankAnnounce" msgpack:"requireCapedPlayerRankAnnounce"` // Key: 19
}

type UserBonus struct {
	Id                  int64   `json:"id" msgpack:"id"`                                   // Key: 0
	ExperienceBonus     float32 `json:"experienceBonus" msgpack:"experienceBonus"`         // Key: 1
	LessonStarRankBonus float32 `json:"lessonStarRankBonus" msgpack:"lessonStarRankBonus"` // Key: 2
}

type UserPreference struct {
	Id           int64      `json:"id" msgpack:"id"`                     // Key: 0
	MultiPartyId int64      `json:"multiPartyId" msgpack:"multiPartyId"` // Key: 1
	BirthDate    *time.Time `json:"birthDate" msgpack:"birthDate"`       // Key: 2
}

type ApiPlayerUserProfile struct {
	Id                     int64            `json:"id" msgpack:"id"`                                         // Key: 0
	Name                   string           `json:"name" msgpack:"name"`                                     // Key: 1
	Introduction           string           `json:"introduction" msgpack:"introduction"`                     // Key: 2
	MainUCharacterId       int64            `json:"mainUCharacterId" msgpack:"mainUCharacterId"`             // Key: 3
	MNameplateId           *int64           `json:"mNameplateId" msgpack:"mNameplateId"`                     // Key: 4
	MNameColorId           int64            `json:"mNameColorId" msgpack:"mNameColorId"`                     // Key: 5
	MTrophyId1             *int64           `json:"mTrophyId1" msgpack:"mTrophyId1"`                         // Key: 6
	MTrophyId2             *int64           `json:"mTrophyId2" msgpack:"mTrophyId2"`                         // Key: 7
	MTrophyId3             *int64           `json:"mTrophyId3" msgpack:"mTrophyId3"`                         // Key: 8
	PlayerRate             float64          `json:"playerRate" msgpack:"playerRate"`                         // Key: 9
	IsPublicPlayerRate     bool             `json:"isPublicPlayerRate" msgpack:"isPublicPlayerRate"`         // Key: 11
	LeagueClass            LeagueClassTypes `json:"leagueClass" msgpack:"leagueClass"`                       // Key: 12
	TotalSpCount           int              `json:"totalSpCount" msgpack:"totalSpCount"`                     // Key: 13
	IsPublicAlbumMainPage  bool             `json:"isPublicAlbumMainPage" msgpack:"isPublicAlbumMainPage"`   // Key: 14
	MNameplateDetailId     *int64           `json:"mNameplateDetailId" msgpack:"mNameplateDetailId"`         // Key: 15
	MainCharacterMasterId  int64            `json:"mainCharacterMasterId" msgpack:"mainCharacterMasterId"`   // Key: 16
	DisplayAwakeningStatus bool             `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 17
	IsPublicActivityLog    bool             `json:"isPublicActivityLog" msgpack:"isPublicActivityLog"`       // Key: 18
	NameBaseColorMasterId  int64            `json:"nameBaseColorMasterId" msgpack:"nameBaseColorMasterId"`   // Key: 19
	IconFrameMasterId      int64            `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`           // Key: 20
	HomeSkinMasterId       int64            `json:"homeSkinMasterId" msgpack:"homeSkinMasterId"`             // Key: 21
}

type ViewPage struct {
	Id           int64         `json:"id" msgpack:"id"`                     // Key: 0
	PageType     ViewPageTypes `json:"pageType" msgpack:"pageType"`         // Key: 1
	LastViewedAt time.Time     `json:"lastViewedAt" msgpack:"lastViewedAt"` // Key: 2
}
