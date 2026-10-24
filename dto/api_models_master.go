// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Master
package dto

import (
	"time"
)

type AccessoryEffectFilterMaster struct {
	Id         int64       `json:"id" msgpack:"id"`                 // Key: 0
	EffectType EffectTypes `json:"effectType" msgpack:"effectType"` // Key: 1
	Name       string      `json:"name" msgpack:"name"`             // Key: 2
	StartDate  time.Time   `json:"startDate" msgpack:"startDate"`   // Key: 3
	Order      int         `json:"order" msgpack:"order"`           // Key: 4
}

type AccessoryEffectMaster struct {
	Id             int64  `json:"id" msgpack:"id"`                         // Key: 0
	EffectMasterId int64  `json:"effectMasterId" msgpack:"effectMasterId"` // Key: 1
	Name           string `json:"name" msgpack:"name"`                     // Key: 2
	Description    string `json:"description" msgpack:"description"`       // Key: 3
	Variety        *int   `json:"variety" msgpack:"variety"`               // Key: 4
}

type AccessoryLevelPatternGroupMaster struct {
	Id       int64                         `json:"id" msgpack:"id"`             // Key: 0
	Patterns []AccessoryLevelPatternMaster `json:"patterns" msgpack:"patterns"` // Key: 1
}

type AccessoryLevelPatternMaster struct {
	Id                                 int64                             `json:"id" msgpack:"id"`                                                                 // Key: 0
	AccessoryLevelPatternGroupMasterId int64                             `json:"accessoryLevelPatternGroupMasterId" msgpack:"accessoryLevelPatternGroupMasterId"` // Key: 1
	Level                              int                               `json:"level" msgpack:"level"`                                                           // Key: 2
	RequiredCoin                       int                               `json:"requiredCoin" msgpack:"requiredCoin"`                                             // Key: 3
	Items                              []AccessoryLevelPatternItemMaster `json:"items" msgpack:"items"`                                                           // Key: 4
}

type AccessoryLevelPatternItemMaster struct {
	ItemMasterId int64 `json:"itemMasterId" msgpack:"itemMasterId"` // Key: 0
	Quantity     int   `json:"quantity" msgpack:"quantity"`         // Key: 1
}

type AccessoryMaster struct {
	Id                           int64              `json:"id" msgpack:"id"`                                                     // Key: 0
	Name                         string             `json:"name" msgpack:"name"`                                                 // Key: 1
	Description                  string             `json:"description" msgpack:"description"`                                   // Key: 2
	Rarity                       PossessionRarities `json:"rarity" msgpack:"rarity"`                                             // Key: 3
	AccessoryLevelPatternGroupId int64              `json:"accessoryLevelPatternGroupId" msgpack:"accessoryLevelPatternGroupId"` // Key: 4
	FixedAccessoryEffects        []int64            `json:"fixedAccessoryEffects" msgpack:"fixedAccessoryEffects"`               // Key: 5
	RandomEffectGroups           []int64            `json:"randomEffectGroups" msgpack:"randomEffectGroups"`                     // Key: 6
	PronounceName                string             `json:"pronounceName" msgpack:"pronounceName"`                               // Key: 7
	Series                       int                `json:"series" msgpack:"series"`                                             // Key: 8
	MaxLevel                     int                `json:"maxLevel" msgpack:"maxLevel"`                                         // Key: 9
}

type AchivementRateRewardMaster struct {
	ThingType      ThingTypes        `json:"thingType" msgpack:"thingType"`           // Key: 0
	ThingId        int64             `json:"thingId" msgpack:"thingId"`               // Key: 1
	ThingQuantity  int               `json:"thingQuantity" msgpack:"thingQuantity"`   // Key: 2
	Difficulty     MusicDifficulties `json:"difficulty" msgpack:"difficulty"`         // Key: 3
	AchivementRate float64           `json:"achivementRate" msgpack:"achivementRate"` // Key: 4
}

type ActivityLogMessageTemplateMaster struct {
	Id              int64               `json:"id" msgpack:"id"`                           // Key: 0
	LogType         ApiActivityLogTypes `json:"logType" msgpack:"logType"`                 // Key: 1
	MessageTemplate string              `json:"messageTemplate" msgpack:"messageTemplate"` // Key: 2
}

type AdditionalRewardPackageMaster struct {
	Id                  int64                         `json:"id" msgpack:"id"`                                   // Key: 0
	Description         string                        `json:"description" msgpack:"description"`                 // Key: 1
	LiveSettingMasterId int64                         `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"` // Key: 2
	Rewards             []AdditionalRewardThingMaster `json:"rewards" msgpack:"rewards"`                         // Key: 3
	Etcetera            bool                          `json:"etcetera" msgpack:"etcetera"`                       // Key: 4
	StartDate           *time.Time                    `json:"startDate" msgpack:"startDate"`                     // Key: 5
	EndDate             *time.Time                    `json:"endDate" msgpack:"endDate"`                         // Key: 6
	Order               int                           `json:"order" msgpack:"order"`                             // Key: 7
}

type AdditionalRewardThingMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type AlbumEffectMaster struct {
	Id             int64  `json:"id" msgpack:"id"`                         // Key: 0
	Level          int    `json:"level" msgpack:"level"`                   // Key: 1
	Description    string `json:"description" msgpack:"description"`       // Key: 2
	EffectMasterId int64  `json:"effectMasterId" msgpack:"effectMasterId"` // Key: 3
}

type AlbumThemeMaster struct {
	Id          int64     `json:"id" msgpack:"id"`                   // Key: 0
	Order       int       `json:"order" msgpack:"order"`             // Key: 1
	Name        string    `json:"name" msgpack:"name"`               // Key: 2
	Description string    `json:"description" msgpack:"description"` // Key: 3
	IsHide      bool      `json:"isHide" msgpack:"isHide"`           // Key: 4
	IsDefault   bool      `json:"isDefault" msgpack:"isDefault"`     // Key: 5
	StartDate   time.Time `json:"startDate" msgpack:"startDate"`     // Key: 6
	EndDate     time.Time `json:"endDate" msgpack:"endDate"`         // Key: 7
}

type AnotherNotationMaster struct {
	Id                  int64                `json:"id" msgpack:"id"`                                   // Key: 0
	MusicMasterId       int64                `json:"musicMasterId" msgpack:"musicMasterId"`             // Key: 1
	VocalVersion        int                  `json:"vocalVersion" msgpack:"vocalVersion"`               // Key: 2
	NotationPath        string               `json:"notationPath" msgpack:"notationPath"`               // Key: 3
	Difficulty          MusicDifficulties    `json:"difficulty" msgpack:"difficulty"`                   // Key: 4
	Level               int                  `json:"level" msgpack:"level"`                             // Key: 5
	AnotherNotationType AnotherNotationTypes `json:"anotherNotationType" msgpack:"anotherNotationType"` // Key: 6
	StartDate           time.Time            `json:"startDate" msgpack:"startDate"`                     // Key: 7
	EndDate             time.Time            `json:"endDate" msgpack:"endDate"`                         // Key: 8
}

type AuditionMaster struct {
	Id                    int64     `json:"id" msgpack:"id"`                                       // Key: 0
	MusicMasterId         int64     `json:"musicMasterId" msgpack:"musicMasterId"`                 // Key: 1
	RecommendedCompany    Companies `json:"recommendedCompany" msgpack:"recommendedCompany"`       // Key: 2
	CanSkip               bool      `json:"canSkip" msgpack:"canSkip"`                             // Key: 3
	SenseNotationMasterId int64     `json:"senseNotationMasterId" msgpack:"senseNotationMasterId"` // Key: 4
	MaxPhase              string    `json:"maxPhase" msgpack:"maxPhase"`                           // Key: 5
	DisplayStartAt        time.Time `json:"displayStartAt" msgpack:"displayStartAt"`               // Key: 6
	DisplayEndAt          time.Time `json:"displayEndAt" msgpack:"displayEndAt"`                   // Key: 7
	VocalVersion          int       `json:"vocalVersion" msgpack:"vocalVersion"`                   // Key: 8
	AuditionGroupNumber   int       `json:"auditionGroupNumber" msgpack:"auditionGroupNumber"`     // Key: 9
	SkipStartAt           time.Time `json:"skipStartAt" msgpack:"skipStartAt"`                     // Key: 10
}

type AuditionPhaseMaster struct {
	Id                            int64 `json:"id" msgpack:"id"`                                                       // Key: 0
	AuditionasterId               int64 `json:"auditionasterId" msgpack:"auditionasterId"`                             // Key: 1
	Phase                         byte  `json:"phase" msgpack:"phase"`                                                 // Key: 2
	RecommendedPlayerRank         int   `json:"recommendedPlayerRank" msgpack:"recommendedPlayerRank"`                 // Key: 3
	ClearScore                    int64 `json:"clearScore" msgpack:"clearScore"`                                       // Key: 4
	StarActCount                  *int  `json:"starActCount" msgpack:"starActCount"`                                   // Key: 5
	AuditionRewardPackageMasterId int64 `json:"auditionRewardPackageMasterId" msgpack:"auditionRewardPackageMasterId"` // Key: 6
}

type AuditionRewardPackageMaster struct {
	Id      int64                 `json:"id" msgpack:"id"`           // Key: 0
	Rewards []AuditionRewardThing `json:"rewards" msgpack:"rewards"` // Key: 1
}

type AuditionRewardThing struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type BannerMaster struct {
	Id                   int64                      `json:"id" msgpack:"id"`                                     // Key: 0
	Order                int                        `json:"order" msgpack:"order"`                               // Key: 1
	Name                 string                     `json:"name" msgpack:"name"`                                 // Key: 2
	JumpType             JumpTypes                  `json:"jumpType" msgpack:"jumpType"`                         // Key: 3
	JumpPageId           string                     `json:"jumpPageId" msgpack:"jumpPageId"`                     // Key: 4
	WebLinkType          *WebLinkTypes              `json:"webLinkType" msgpack:"webLinkType"`                   // Key: 5
	LinkUrl              string                     `json:"linkUrl" msgpack:"linkUrl"`                           // Key: 6
	DeleteConditionType  BannerDeleteConditionTypes `json:"deleteConditionType" msgpack:"deleteConditionType"`   // Key: 7
	DeleteConditionValue string                     `json:"deleteConditionValue" msgpack:"deleteConditionValue"` // Key: 8
	ImagePath            string                     `json:"imagePath" msgpack:"imagePath"`                       // Key: 9
	StartDate            time.Time                  `json:"startDate" msgpack:"startDate"`                       // Key: 10
	EndDate              time.Time                  `json:"endDate" msgpack:"endDate"`                           // Key: 11
}

type BodyMotionMaster struct {
	Id         int64  `json:"id" msgpack:"id"`                 // Key: 0
	MotionName string `json:"motionName" msgpack:"motionName"` // Key: 1
}

type BombMaster struct {
	Id          int64  `json:"id" msgpack:"id"`                   // Key: 0
	Name        string `json:"name" msgpack:"name"`               // Key: 1
	Description string `json:"description" msgpack:"description"` // Key: 2
	Order       int    `json:"order" msgpack:"order"`             // Key: 3
	Hidden      bool   `json:"hidden" msgpack:"hidden"`           // Key: 4
	IsDefault   bool   `json:"isDefault" msgpack:"isDefault"`     // Key: 5
}

type BonusLiveMaster struct {
	Id                           int64                         `json:"id" msgpack:"id"`                                                     // Key: 0
	EventMasterId                int64                         `json:"eventMasterId" msgpack:"eventMasterId"`                               // Key: 1
	FirstClearEventPointQuantity int                           `json:"firstClearEventPointQuantity" msgpack:"firstClearEventPointQuantity"` // Key: 2
	ClearEventPointQuantity      int                           `json:"clearEventPointQuantity" msgpack:"clearEventPointQuantity"`           // Key: 3
	UnlockConditionType          BonusLiveUnlockConditionTypes `json:"unlockConditionType" msgpack:"unlockConditionType"`                   // Key: 4
	UnlockConditionValue         *int64                        `json:"unlockConditionValue" msgpack:"unlockConditionValue"`                 // Key: 5
	Stages                       []BonusLiveStageMaster        `json:"stages" msgpack:"stages"`                                             // Key: 6
}

type BonusLiveStageMaster struct {
	Id                    int64                             `json:"id" msgpack:"id"`                                       // Key: 0
	BonusLiveMasterId     int64                             `json:"bonusLiveMasterId" msgpack:"bonusLiveMasterId"`         // Key: 1
	Order                 int                               `json:"order" msgpack:"order"`                                 // Key: 2
	MusicMasterId         int64                             `json:"musicMasterId" msgpack:"musicMasterId"`                 // Key: 3
	VocalVersion          int                               `json:"vocalVersion" msgpack:"vocalVersion"`                   // Key: 4
	SenseNotationMasterId int64                             `json:"senseNotationMasterId" msgpack:"senseNotationMasterId"` // Key: 5
	RecommendedCompany    *Companies                        `json:"recommendedCompany" msgpack:"recommendedCompany"`       // Key: 6
	RecommendedLevel      int                               `json:"recommendedLevel" msgpack:"recommendedLevel"`           // Key: 7
	ClearScore            int64                             `json:"clearScore" msgpack:"clearScore"`                       // Key: 8
	UnlockConditionValue  *int64                            `json:"unlockConditionValue" msgpack:"unlockConditionValue"`   // Key: 9
	StartDate             time.Time                         `json:"startDate" msgpack:"startDate"`                         // Key: 10
	Rewards               []BonusLiveStageRewardThingMaster `json:"rewards" msgpack:"rewards"`                             // Key: 11
	ClearStarActCount     *int                              `json:"clearStarActCount" msgpack:"clearStarActCount"`         // Key: 12
}

type BonusLiveStageRewardThingMaster struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 2
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 3
	IsFirstClear  bool       `json:"isFirstClear" msgpack:"isFirstClear"`   // Key: 4
}

type EffectOrderMaster struct {
	Order          int   `json:"order" msgpack:"order"`                   // Key: 0
	EffectMasterId int64 `json:"effectMasterId" msgpack:"effectMasterId"` // Key: 1
}

type BranchMaster struct {
	Id            int64               `json:"id" msgpack:"id"`                       // Key: 6
	Order         int                 `json:"order" msgpack:"order"`                 // Key: 0
	BranchEffects []EffectOrderMaster `json:"branchEffects" msgpack:"branchEffects"` // Key: 1
	JudgeType1    *BranchJudgeTypes   `json:"judgeType1" msgpack:"judgeType1"`       // Key: 2
	Parameter1    *int64              `json:"parameter1" msgpack:"parameter1"`       // Key: 3
	JudgeType2    *BranchJudgeTypes   `json:"judgeType2" msgpack:"judgeType2"`       // Key: 4
	Parameter2    *int64              `json:"parameter2" msgpack:"parameter2"`       // Key: 5
}

type BuffItemMaster struct {
	Id                int64               `json:"id" msgpack:"id"`                               // Key: 0
	ItemMasterId      int64               `json:"itemMasterId" msgpack:"itemMasterId"`           // Key: 1
	EffectType        CampaignEffectTypes `json:"effectType" msgpack:"effectType"`               // Key: 2
	EffectValue       int                 `json:"effectValue" msgpack:"effectValue"`             // Key: 3
	ValidDay          int                 `json:"validDay" msgpack:"validDay"`                   // Key: 4
	DialogTitle       string              `json:"dialogTitle" msgpack:"dialogTitle"`             // Key: 5
	DialogDescription string              `json:"dialogDescription" msgpack:"dialogDescription"` // Key: 6
}

type CampaignMaster struct {
	Id                       int64               `json:"id" msgpack:"id"`                                             // Key: 0
	Title                    string              `json:"title" msgpack:"title"`                                       // Key: 1
	Description              string              `json:"description" msgpack:"description"`                           // Key: 2
	IconImagePath            string              `json:"iconImagePath" msgpack:"iconImagePath"`                       // Key: 3
	Order                    int                 `json:"order" msgpack:"order"`                                       // Key: 4
	StartDate                time.Time           `json:"startDate" msgpack:"startDate"`                               // Key: 5
	EndDate                  time.Time           `json:"endDate" msgpack:"endDate"`                                   // Key: 6
	ComebackCampaignMasterId *int64              `json:"comebackCampaignMasterId" msgpack:"comebackCampaignMasterId"` // Key: 7
	CampaignEffectType       CampaignEffectTypes `json:"campaignEffectType" msgpack:"campaignEffectType"`             // Key: 8
	CampaignEffectValue      int                 `json:"campaignEffectValue" msgpack:"campaignEffectValue"`           // Key: 9
}

type CategoryGroupMaster struct {
	Id        int64  `json:"id" msgpack:"id"`               // Key: 0
	Order     int    `json:"order" msgpack:"order"`         // Key: 1
	Name      string `json:"name" msgpack:"name"`           // Key: 2
	ColorCode string `json:"colorCode" msgpack:"colorCode"` // Key: 3
}

type CategoryMaster struct {
	Id                    int64  `json:"id" msgpack:"id"`                                       // Key: 0
	CategoryGroupMasterId int64  `json:"categoryGroupMasterId" msgpack:"categoryGroupMasterId"` // Key: 1
	Order                 int    `json:"order" msgpack:"order"`                                 // Key: 2
	Name                  string `json:"name" msgpack:"name"`                                   // Key: 3
}

type ChangeBodyMotionMaster struct {
	Id               int64   `json:"id" msgpack:"id"`                             // Key: 0
	BeforeMotionName string  `json:"beforeMotionName" msgpack:"beforeMotionName"` // Key: 1
	AfterMotionName  string  `json:"afterMotionName" msgpack:"afterMotionName"`   // Key: 2
	Second           float32 `json:"second" msgpack:"second"`                     // Key: 3
}

type CharacterAwakeningItemGroupMaster struct {
	Id    int64                          `json:"id" msgpack:"id"`       // Key: 0
	Items []CharacterAwakeningItemMaster `json:"items" msgpack:"items"` // Key: 1
}

type CharacterAwakeningItemMaster struct {
	Id               int64 `json:"id" msgpack:"id"`                             // Key: 0
	AwakeningPhase   int   `json:"awakeningPhase" msgpack:"awakeningPhase"`     // Key: 3
	ItemMasterId     int64 `json:"itemMasterId" msgpack:"itemMasterId"`         // Key: 4
	RequiredQuantity int   `json:"requiredQuantity" msgpack:"requiredQuantity"` // Key: 5
}

type CharacterBaseBloomGenericItemMaster struct {
	Id                    int64                `json:"id" msgpack:"id"`                                       // Key: 0
	ItemMasterId          int64                `json:"itemMasterId" msgpack:"itemMasterId"`                   // Key: 1
	CharacterBaseMasterId int64                `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 2
	TalentBloomItemType   TalentBloomItemTypes `json:"talentBloomItemType" msgpack:"talentBloomItemType"`     // Key: 3
}

type CharacterBaseMaster struct {
	Id                     int64              `json:"id" msgpack:"id"`                                         // Key: 0
	Name                   string             `json:"name" msgpack:"name"`                                     // Key: 1
	Description            string             `json:"description" msgpack:"description"`                       // Key: 2
	School                 string             `json:"school" msgpack:"school"`                                 // Key: 3
	Grade                  *int               `json:"grade" msgpack:"grade"`                                   // Key: 4
	BirthMonth             *int               `json:"birthMonth" msgpack:"birthMonth"`                         // Key: 5
	BirthDay               *int               `json:"birthDay" msgpack:"birthDay"`                             // Key: 6
	Height                 int                `json:"height" msgpack:"height"`                                 // Key: 7
	Hobby                  string             `json:"hobby" msgpack:"hobby"`                                   // Key: 8
	CompanyMasterId        int64              `json:"companyMasterId" msgpack:"companyMasterId"`               // Key: 9
	NameRomanization       string             `json:"nameRomanization" msgpack:"nameRomanization"`             // Key: 10
	SenseName              string             `json:"senseName" msgpack:"senseName"`                           // Key: 11
	SenseEffect            string             `json:"senseEffect" msgpack:"senseEffect"`                       // Key: 12
	CharacterVoice         string             `json:"characterVoice" msgpack:"characterVoice"`                 // Key: 13
	ProfileImageAssetId    string             `json:"profileImageAssetId" msgpack:"profileImageAssetId"`       // Key: 14
	Age                    *int               `json:"age" msgpack:"age"`                                       // Key: 15
	FamilyNameRomanization string             `json:"familyNameRomanization" msgpack:"familyNameRomanization"` // Key: 16
	FirstNameRomanization  string             `json:"firstNameRomanization" msgpack:"firstNameRomanization"`   // Key: 17
	PronounceFamilyName    string             `json:"pronounceFamilyName" msgpack:"pronounceFamilyName"`       // Key: 18
	PronounceFirstName     string             `json:"pronounceFirstName" msgpack:"pronounceFirstName"`         // Key: 19
	FamilyName             string             `json:"familyName" msgpack:"familyName"`                         // Key: 20
	FirstName              string             `json:"firstName" msgpack:"firstName"`                           // Key: 21
	EvoSenseName           string             `json:"evoSenseName" msgpack:"evoSenseName"`                     // Key: 22
	EvoSenseEffect         string             `json:"evoSenseEffect" msgpack:"evoSenseEffect"`                 // Key: 23
	DefaultCostumeMasterId int64              `json:"defaultCostumeMasterId" msgpack:"defaultCostumeMasterId"` // Key: 24
	CharacterBaseType      CharacterBaseTypes `json:"characterBaseType" msgpack:"characterBaseType"`           // Key: 25
}

type CharacterBloomBonusGroupMaster struct {
	Id           int64                        `json:"id" msgpack:"id"`                     // Key: 0
	BloomBonuses []CharacterBloomBonusMaster  `json:"bloomBonuses" msgpack:"bloomBonuses"` // Key: 1
	BloomRewards []CharacterBloomRewardMaster `json:"bloomRewards" msgpack:"bloomRewards"` // Key: 2
}

type CharacterBloomBonusMaster struct {
	BloomBonusType BloomBonusTypes `json:"bloomBonusType" msgpack:"bloomBonusType"` // Key: 1
	Description    string          `json:"description" msgpack:"description"`       // Key: 2
	Phase          int             `json:"phase" msgpack:"phase"`                   // Key: 3
	EffectMasterId int64           `json:"effectMasterId" msgpack:"effectMasterId"` // Key: 5
	IconPath       string          `json:"iconPath" msgpack:"iconPath"`             // Key: 6
}

type CharacterBloomRewardMaster struct {
	Phase         int        `json:"phase" msgpack:"phase"`                 // Key: 4
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type CharacterBloomDetailMaster struct {
	CharacterMasterId int64  `json:"characterMasterId" msgpack:"characterMasterId"` // Key: 0
	Text              string `json:"text" msgpack:"text"`                           // Key: 1
	VoicePath         string `json:"voicePath" msgpack:"voicePath"`                 // Key: 2
}

type CharacterBloomItemMaster struct {
	Rarity                   CharacterRarities    `json:"rarity" msgpack:"rarity"`                                     // Key: 0
	CurrentStage             int                  `json:"currentStage" msgpack:"currentStage"`                         // Key: 1
	RequiredPieceAmount      int                  `json:"requiredPieceAmount" msgpack:"requiredPieceAmount"`           // Key: 2
	TalentBloomItemType      TalentBloomItemTypes `json:"talentBloomItemType" msgpack:"talentBloomItemType"`           // Key: 3
	GenericBloomItemMasterId *int64               `json:"genericBloomItemMasterId" msgpack:"genericBloomItemMasterId"` // Key: 4
	RequiredItemMasterId     *int64               `json:"requiredItemMasterId" msgpack:"requiredItemMasterId"`         // Key: 5
	RequiredItemAmount       *int                 `json:"requiredItemAmount" msgpack:"requiredItemAmount"`             // Key: 6
}

type CharacterCategoryMaster struct {
	CategoryMasterId int64 `json:"categoryMasterId" msgpack:"categoryMasterId"` // Key: 0
	IsAwaken         bool  `json:"isAwaken" msgpack:"isAwaken"`                 // Key: 1
}

type CharacterEpisodeMaster struct {
	Id                     int64                 `json:"id" msgpack:"id"`                                         // Key: 0
	CharacterMasterId      int64                 `json:"characterMasterId" msgpack:"characterMasterId"`           // Key: 1
	EpisodeMasterId        int64                 `json:"episodeMasterId" msgpack:"episodeMasterId"`               // Key: 2
	EpisodeOrder           CharacterEpisodeOrder `json:"episodeOrder" msgpack:"episodeOrder"`                     // Key: 3
	RequiredCharacterLevel int                   `json:"requiredCharacterLevel" msgpack:"requiredCharacterLevel"` // Key: 4
}

type CharacterEpisodeRelationMaster struct {
	CharacterMasterid           int64   `json:"characterMasterid" msgpack:"characterMasterid"`                     // Key: 0
	PreviousStoryMasterIds      []int64 `json:"previousStoryMasterIds" msgpack:"previousStoryMasterIds"`           // Key: 1
	NextStoryMasterIds          []int64 `json:"nextStoryMasterIds" msgpack:"nextStoryMasterIds"`                   // Key: 2
	SideStoryCharacterMasterIds []int64 `json:"sideStoryCharacterMasterIds" msgpack:"sideStoryCharacterMasterIds"` // Key: 3
	RelatedStoryMasterIds       []int64 `json:"relatedStoryMasterIds" msgpack:"relatedStoryMasterIds"`             // Key: 4
}

type CharacterEpisodeReleaseItemGroupMaster struct {
	Id    int64                               `json:"id" msgpack:"id"`       // Key: 0
	Items []CharacterEpisodeReleaseItemMaster `json:"items" msgpack:"items"` // Key: 1
}

type CharacterEpisodeReleaseItemMaster struct {
	ItemMasterId     int64 `json:"itemMasterId" msgpack:"itemMasterId"`         // Key: 0
	RequiredQuantity int   `json:"requiredQuantity" msgpack:"requiredQuantity"` // Key: 1
	Order            int   `json:"order" msgpack:"order"`                       // Key: 2
}

type CharacterExperienceItemMaster struct {
	Id                        int64   `json:"id" msgpack:"id"`                                               // Key: 0
	ItemMasterId              int64   `json:"itemMasterId" msgpack:"itemMasterId"`                           // Key: 1
	AcquirableExperience      int64   `json:"acquirableExperience" msgpack:"acquirableExperience"`           // Key: 2
	AcquirableExperienceBonus float32 `json:"acquirableExperienceBonus" msgpack:"acquirableExperienceBonus"` // Key: 3
}

type CharacterKeyMissionMaster struct {
	Id                    int64     `json:"id" msgpack:"id"`                                       // Key: 0
	Level                 int       `json:"level" msgpack:"level"`                                 // Key: 1
	Color                 string    `json:"color" msgpack:"color"`                                 // Key: 2
	GiveStarPoint         int       `json:"giveStarPoint" msgpack:"giveStarPoint"`                 // Key: 3
	StartDate             time.Time `json:"startDate" msgpack:"startDate"`                         // Key: 4
	EndDate               time.Time `json:"endDate" msgpack:"endDate"`                             // Key: 5
	RequiredCategoryCount int       `json:"requiredCategoryCount" msgpack:"requiredCategoryCount"` // Key: 6
}

type CharacterLessonScoreRewardMaster struct {
	Id                       int64 `json:"id" msgpack:"id"`                                             // Key: 0
	CharacterBaseMasterId    int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`       // Key: 1
	RequiredScore            int64 `json:"requiredScore" msgpack:"requiredScore"`                       // Key: 2
	LessonScoreGroupMasterId int64 `json:"lessonScoreGroupMasterId" msgpack:"lessonScoreGroupMasterId"` // Key: 3
}

type CharacterLevelMaster struct {
	Level                int       `json:"level" msgpack:"level"`                               // Key: 0
	ExperienceToLevelUp  int64     `json:"experienceToLevelUp" msgpack:"experienceToLevelUp"`   // Key: 1
	CharacterStatusLevel int       `json:"characterStatusLevel" msgpack:"characterStatusLevel"` // Key: 2
	StartDate            time.Time `json:"startDate" msgpack:"startDate"`                       // Key: 3
}

type CharacterMaster struct {
	Id                                  int64                     `json:"id" msgpack:"id"`                                       // Key: 0
	CharacterBaseMasterId               int64                     `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	CharacterBaseMaster                 *CharacterBaseMaster      `json:"characterBaseMaster" msgpack:"characterBaseMaster"`
	Name                                string                    `json:"name" msgpack:"name"`                                                               // Key: 2
	Description                         string                    `json:"description" msgpack:"description"`                                                 // Key: 3
	AssetId                             string                    `json:"assetId" msgpack:"assetId"`                                                         // Key: 4
	Rarity                              CharacterRarities         `json:"rarity" msgpack:"rarity"`                                                           // Key: 5
	Attribute                           Attributes                `json:"attribute" msgpack:"attribute"`                                                     // Key: 6
	SecondaryAttribute                  *Attributes               `json:"secondaryAttribute" msgpack:"secondaryAttribute"`                                   // Key: 27
	MinLevelStatus                      *Status                   `json:"minLevelStatus" msgpack:"minLevelStatus"`                                           // Key: 7
	StarActMasterId                     int64                     `json:"starActMasterId" msgpack:"starActMasterId"`                                         // Key: 9
	AwakenStarActMasterId               *int64                    `json:"awakenStarActMasterId" msgpack:"awakenStarActMasterId"`                             // Key: 10
	SenseMasterId                       int64                     `json:"senseMasterId" msgpack:"senseMasterId"`                                             // Key: 11
	ForbidGenericItemBloom              bool                      `json:"forbidGenericItemBloom" msgpack:"forbidGenericItemBloom"`                           // Key: 12
	BloomBonusGroupMasterId             int64                     `json:"bloomBonusGroupMasterId" msgpack:"bloomBonusGroupMasterId"`                         // Key: 13
	SenseEnhanceItemGroupMasterId       int64                     `json:"senseEnhanceItemGroupMasterId" msgpack:"senseEnhanceItemGroupMasterId"`             // Key: 14
	FirstEpisodeReleaseItemGroupId      int64                     `json:"firstEpisodeReleaseItemGroupId" msgpack:"firstEpisodeReleaseItemGroupId"`           // Key: 15
	SecondEpisodeReleaseItemGroupId     int64                     `json:"secondEpisodeReleaseItemGroupId" msgpack:"secondEpisodeReleaseItemGroupId"`         // Key: 16
	CharacterAwakeningItemGroupMasterId *int64                    `json:"characterAwakeningItemGroupMasterId" msgpack:"characterAwakeningItemGroupMasterId"` // Key: 17
	DisplayStartAt                      time.Time                 `json:"displayStartAt" msgpack:"displayStartAt"`                                           // Key: 18
	DisplayEndAt                        time.Time                 `json:"displayEndAt" msgpack:"displayEndAt"`                                               // Key: 19
	UnlockText                          string                    `json:"unlockText" msgpack:"unlockText"`                                                   // Key: 20
	Categories                          []CharacterCategoryMaster `json:"categories" msgpack:"categories"`                                                   // Key: 21
	LeaderSenseMasterId                 *int64                    `json:"leaderSenseMasterId" msgpack:"leaderSenseMasterId"`                                 // Key: 22
	MaxTalentStage                      int                       `json:"maxTalentStage" msgpack:"maxTalentStage"`                                           // Key: 23
	MaxTalentStageReleaseDate           *time.Time                `json:"maxTalentStageReleaseDate" msgpack:"maxTalentStageReleaseDate"`                     // Key: 24
	SecondaryCharacterBaseMasterId      *int64                    `json:"secondaryCharacterBaseMasterId" msgpack:"secondaryCharacterBaseMasterId"`           // Key: 25
	SecondaryCharacterBaseMaster        *CharacterBaseMaster      `json:"secondaryCharacterBaseMaster" msgpack:"secondaryCharacterBaseMaster"`
	SecondarySenseMasterId              *int64                    `json:"secondarySenseMasterId" msgpack:"secondarySenseMasterId"` // Key: 26
}

type CharacterMissionCategoryLevelMaster struct {
	Id            int64                         `json:"id" msgpack:"id"`                       // Key: 0
	CategoryType  CharacterMissionCategoryTypes `json:"categoryType" msgpack:"categoryType"`   // Key: 1
	Level         int                           `json:"level" msgpack:"level"`                 // Key: 2
	GiveStarPoint int                           `json:"giveStarPoint" msgpack:"giveStarPoint"` // Key: 3
	StartDate     time.Time                     `json:"startDate" msgpack:"startDate"`         // Key: 4
	EndDate       time.Time                     `json:"endDate" msgpack:"endDate"`             // Key: 5
}

type CharacterMissionItemMaster struct {
	Id                       int64 `json:"id" msgpack:"id"`                                             // Key: 0
	ItemMasterId             int64 `json:"itemMasterId" msgpack:"itemMasterId"`                         // Key: 1
	CharacterBaseMasterId    int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`       // Key: 2
	CharacterMissionMasterId int64 `json:"characterMissionMasterId" msgpack:"characterMissionMasterId"` // Key: 3
}

type CharacterMissionMaster struct {
	Id        int64                         `json:"id" msgpack:"id"`               // Key: 0
	Title     string                        `json:"title" msgpack:"title"`         // Key: 1
	JumpType  *JumpTypes                    `json:"jumpType" msgpack:"jumpType"`   // Key: 2
	JumpValue *int64                        `json:"jumpValue" msgpack:"jumpValue"` // Key: 3
	Stages    []CharacterMissionStageMaster `json:"stages" msgpack:"stages"`       // Key: 4
}

type CharacterMissionStageMaster struct {
	Id                                    int64 `json:"id" msgpack:"id"`                                                                       // Key: 0
	CharacterMissionCategoryLevelMasterId int64 `json:"characterMissionCategoryLevelMasterId" msgpack:"characterMissionCategoryLevelMasterId"` // Key: 1
	CharacterMissionMasterId              int64 `json:"characterMissionMasterId" msgpack:"characterMissionMasterId"`                           // Key: 2
	ExclusionNoSenseCharacter             bool  `json:"exclusionNoSenseCharacter" msgpack:"exclusionNoSenseCharacter"`                         // Key: 3
	Order                                 int   `json:"order" msgpack:"order"`                                                                 // Key: 4
	StageOrder                            int   `json:"stageOrder" msgpack:"stageOrder"`                                                       // Key: 5
	GoalCount                             int64 `json:"goalCount" msgpack:"goalCount"`                                                         // Key: 6
}

type CharacterPieceMaster struct {
	ItemMasterId         int64                `json:"itemMasterId" msgpack:"itemMasterId"`                 // Key: 0
	CharacterMasterId    int64                `json:"characterMasterId" msgpack:"characterMasterId"`       // Key: 1
	DugongRequiredAmount int                  `json:"dugongRequiredAmount" msgpack:"dugongRequiredAmount"` // Key: 2
	TalentBloomItemType  TalentBloomItemTypes `json:"talentBloomItemType" msgpack:"talentBloomItemType"`   // Key: 3
}

type CharacterPointEventCharacterRankingRewardMaster struct {
	Id                                                           int64 `json:"id" msgpack:"id"`                                                                                                                     // Key: 0
	CharacterPointEventMasterId                                  int64 `json:"characterPointEventMasterId" msgpack:"characterPointEventMasterId"`                                                                   // Key: 1
	CharacterBaseMasterId                                        int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`                                                                               // Key: 2
	TargetMinRank                                                int   `json:"targetMinRank" msgpack:"targetMinRank"`                                                                                               // Key: 3
	TargetMaxRank                                                int   `json:"targetMaxRank" msgpack:"targetMaxRank"`                                                                                               // Key: 4
	CharacterPointEventCharacterRankingRewardItemPackageMasterId int64 `json:"characterPointEventCharacterRankingRewardItemPackageMasterId" msgpack:"characterPointEventCharacterRankingRewardItemPackageMasterId"` // Key: 5
}

type CharacterPointEventCharacterRankingRewardItemPackageMaster struct {
	Id      int64                                                 `json:"id" msgpack:"id"`           // Key: 0
	Rewards []CharacterPointEventCharacterRankingRewardItemMaster `json:"rewards" msgpack:"rewards"` // Key: 1
}

type CharacterPointEventCharacterRankingRewardItemMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type CharacterPointEventMaster struct {
	Id                     int64   `json:"id" msgpack:"id"`                                         // Key: 0
	EventMasterId          int64   `json:"eventMasterId" msgpack:"eventMasterId"`                   // Key: 1
	Title                  string  `json:"title" msgpack:"title"`                                   // Key: 2
	ExchangeShopMasterId   int64   `json:"exchangeShopMasterId" msgpack:"exchangeShopMasterId"`     // Key: 3
	EventPointItemMasterId int64   `json:"eventPointItemMasterId" msgpack:"eventPointItemMasterId"` // Key: 4
	CharacterBaseMasterIds []int64 `json:"characterBaseMasterIds" msgpack:"characterBaseMasterIds"` // Key: 5
}

type CharacterProfileRestrictionMaster struct {
	Id                      int64                           `json:"id" msgpack:"id"`                                       // Key: 0
	CharacterBaseMasterId   int64                           `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	CharacterBaseMaster     *CharacterBaseMaster            `json:"characterBaseMaster" msgpack:"characterBaseMaster"`
	ProfileRestrictionItem  CharacterProfileRestrictionItem `json:"profileRestrictionItem" msgpack:"profileRestrictionItem"`   // Key: 2
	UnlockConditionMasterId int64                           `json:"unlockConditionMasterId" msgpack:"unlockConditionMasterId"` // Key: 3
	ProfileUnlockValue      int64                           `json:"profileUnlockValue" msgpack:"profileUnlockValue"`           // Key: 4
}

type CharacterSenseEnhanceItemGroupMaster struct {
	Id    int64                             `json:"id" msgpack:"id"`       // Key: 0
	Items []CharacterSenseEnhanceItemMaster `json:"items" msgpack:"items"` // Key: 1
}

type CharacterSenseEnhanceItemMaster struct {
	CurrentLevel     int   `json:"currentLevel" msgpack:"currentLevel"`         // Key: 0
	ItemMasterId     int64 `json:"itemMasterId" msgpack:"itemMasterId"`         // Key: 1
	RequiredQuantity int   `json:"requiredQuantity" msgpack:"requiredQuantity"` // Key: 2
}

type CharacterStarRankMaster struct {
	Rank                int     `json:"rank" msgpack:"rank"`                               // Key: 0
	NextRankPoint       int64   `json:"nextRankPoint" msgpack:"nextRankPoint"`             // Key: 1
	RequiredLessonScore int64   `json:"requiredLessonScore" msgpack:"requiredLessonScore"` // Key: 2
	StatusBonus         float32 `json:"statusBonus" msgpack:"statusBonus"`                 // Key: 3
}

type CharacterStarRankRewardGroupMaster struct {
	Id      int64                           `json:"id" msgpack:"id"`           // Key: 0
	Rewards []CharacterStarRankRewardMaster `json:"rewards" msgpack:"rewards"` // Key: 1
}

type CharacterStarRankRewardMaster struct {
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type CircleEventCirclePointRewardMaster struct {
	Id                  int64      `json:"id" msgpack:"id"`                                   // Key: 0
	CircleEventMasterId int64      `json:"circleEventMasterId" msgpack:"circleEventMasterId"` // Key: 1
	Point               int64      `json:"point" msgpack:"point"`                             // Key: 2
	ThingType           ThingTypes `json:"thingType" msgpack:"thingType"`                     // Key: 3
	ThingId             int64      `json:"thingId" msgpack:"thingId"`                         // Key: 4
	ThingQuantity       int        `json:"thingQuantity" msgpack:"thingQuantity"`             // Key: 5
	Order               int        `json:"order" msgpack:"order"`                             // Key: 6
}

type CircleEventCircleRankingRewardMaster struct {
	Id                                                int64 `json:"id" msgpack:"id"`                                                                                               // Key: 0
	CircleEventMasterId                               int64 `json:"circleEventMasterId" msgpack:"circleEventMasterId"`                                                             // Key: 1
	CircleEventCircleRankingRewardItemPackageMasterId int64 `json:"circleEventCircleRankingRewardItemPackageMasterId" msgpack:"circleEventCircleRankingRewardItemPackageMasterId"` // Key: 2
	TargetMinRank                                     int   `json:"targetMinRank" msgpack:"targetMinRank"`                                                                         // Key: 3
	TargetMaxRank                                     int   `json:"targetMaxRank" msgpack:"targetMaxRank"`                                                                         // Key: 4
}

type CircleEventCircleRankingRewardItemPackageMaster struct {
	Id      int64                                `json:"id" msgpack:"id"`           // Key: 0
	Rewards []CircleEventCircleRankingRewardItem `json:"rewards" msgpack:"rewards"` // Key: 1
}

type CircleEventCircleRankingRewardItem struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type CircleEventMaster struct {
	Id                                            int64 `json:"id" msgpack:"id"`                                                                                       // Key: 0
	EventMasterId                                 int64 `json:"eventMasterId" msgpack:"eventMasterId"`                                                                 // Key: 1
	CircleEventMissionRefreshSettingGroupMasterId int64 `json:"circleEventMissionRefreshSettingGroupMasterId" msgpack:"circleEventMissionRefreshSettingGroupMasterId"` // Key: 2
}

type CircleEventMissionMaster struct {
	Id                    int64                   `json:"id" msgpack:"id"`                                       // Key: 0
	CircleEventMasterId   int64                   `json:"circleEventMasterId" msgpack:"circleEventMasterId"`     // Key: 1
	MissionCategory       CircleEventMissionTypes `json:"missionCategory" msgpack:"missionCategory"`             // Key: 2
	Title                 string                  `json:"title" msgpack:"title"`                                 // Key: 3
	JumpType              JumpTypes               `json:"jumpType" msgpack:"jumpType"`                           // Key: 4
	JumpTargetId          *int64                  `json:"jumpTargetId" msgpack:"jumpTargetId"`                   // Key: 5
	GoalCount             int64                   `json:"goalCount" msgpack:"goalCount"`                         // Key: 6
	GivePointAtGoal       int                     `json:"givePointAtGoal" msgpack:"givePointAtGoal"`             // Key: 7
	CircleGoalCount       int64                   `json:"circleGoalCount" msgpack:"circleGoalCount"`             // Key: 8
	CircleGivePointAtGoal int                     `json:"circleGivePointAtGoal" msgpack:"circleGivePointAtGoal"` // Key: 9
	IsDefault             bool                    `json:"isDefault" msgpack:"isDefault"`                         // Key: 10
}

type CircleEventMissionRefreshSettingGroupMaster struct {
	Id       int64                              `json:"id" msgpack:"id"`             // Key: 0
	Settings []CircleEventMissionRefreshSetting `json:"settings" msgpack:"settings"` // Key: 1
}

type CircleEventMissionRefreshSetting struct {
	MinRefreshCount int `json:"minRefreshCount" msgpack:"minRefreshCount"` // Key: 0
	MaxRefreshCount int `json:"maxRefreshCount" msgpack:"maxRefreshCount"` // Key: 1
	RequiredCoin    int `json:"requiredCoin" msgpack:"requiredCoin"`       // Key: 2
}

type CircleSupportCompanyLevelDetailMaster struct {
	Id                   int64     `json:"id" msgpack:"id"`                                     // Key: 0
	Company              Companies `json:"company" msgpack:"company"`                           // Key: 1
	Level                int       `json:"level" msgpack:"level"`                               // Key: 2
	RequiredSupportPoint int       `json:"requiredSupportPoint" msgpack:"requiredSupportPoint"` // Key: 3
	EffectMasterId       int64     `json:"effectMasterId" msgpack:"effectMasterId"`             // Key: 4
	Description          string    `json:"description" msgpack:"description"`                   // Key: 5
	ReleaseDate          time.Time `json:"releaseDate" msgpack:"releaseDate"`                   // Key: 6
}

type CircleSupportCompanyLevelLimitMaster struct {
	Id             int64                                        `json:"id" msgpack:"id"`                         // Key: 0
	NextLevelLimit int                                          `json:"nextLevelLimit" msgpack:"nextLevelLimit"` // Key: 1
	RequiredCoin   int64                                        `json:"requiredCoin" msgpack:"requiredCoin"`     // Key: 2
	ReleaseDate    time.Time                                    `json:"releaseDate" msgpack:"releaseDate"`       // Key: 3
	Details        []CircleSupportCompanyLevelLimitDetailMaster `json:"details" msgpack:"details"`               // Key: 4
}

type CircleSupportCompanyLevelLimitDetailMaster struct {
	Order            int   `json:"order" msgpack:"order"`                       // Key: 0
	ItemMasterId     int64 `json:"itemMasterId" msgpack:"itemMasterId"`         // Key: 1
	RequiredQuantity int64 `json:"requiredQuantity" msgpack:"requiredQuantity"` // Key: 2
}

type CircleTheaterLevelMaster struct {
	Id               int64  `json:"id" msgpack:"id"`                             // Key: 0
	Level            int    `json:"level" msgpack:"level"`                       // Key: 1
	StaminaBaseValue int    `json:"staminaBaseValue" msgpack:"staminaBaseValue"` // Key: 2
	Description      string `json:"description" msgpack:"description"`           // Key: 3
}

type ComboRewardMaster struct {
	ThingType         ThingTypes        `json:"thingType" msgpack:"thingType"`                 // Key: 0
	ThingId           int64             `json:"thingId" msgpack:"thingId"`                     // Key: 1
	ThingQuantity     int               `json:"thingQuantity" msgpack:"thingQuantity"`         // Key: 2
	Difficulty        MusicDifficulties `json:"difficulty" msgpack:"difficulty"`               // Key: 3
	ComboCountPercent int               `json:"comboCountPercent" msgpack:"comboCountPercent"` // Key: 4
}

type ComebackCampaignMaster struct {
	Id                      int64 `json:"id" msgpack:"id"`                                           // Key: 0
	MissionActiveDays       int   `json:"missionActiveDays" msgpack:"missionActiveDays"`             // Key: 1
	JewelShopItemActiveDays int   `json:"jewelShopItemActiveDays" msgpack:"jewelShopItemActiveDays"` // Key: 2
	LoginBonusActiveDays    int   `json:"loginBonusActiveDays" msgpack:"loginBonusActiveDays"`       // Key: 3
	BuffCampaignActiveDays  int   `json:"buffCampaignActiveDays" msgpack:"buffCampaignActiveDays"`   // Key: 4
}

type ComicMaster struct {
	Id        int64                `json:"id" msgpack:"id"`               // Key: 0
	GroupName string               `json:"groupName" msgpack:"groupName"` // Key: 1
	Episodes  []ComicEpisodeMaster `json:"episodes" msgpack:"episodes"`   // Key: 2
}

type ComicEpisodeMaster struct {
	Id        int64     `json:"id" msgpack:"id"`               // Key: 0
	Order     int       `json:"order" msgpack:"order"`         // Key: 1
	Title     string    `json:"title" msgpack:"title"`         // Key: 2
	Body      string    `json:"body" msgpack:"body"`           // Key: 3
	StartDate time.Time `json:"startDate" msgpack:"startDate"` // Key: 4
}

type CompanyMaster struct {
	Id          int64     `json:"id" msgpack:"id"`                   // Key: 0
	Name        string    `json:"name" msgpack:"name"`               // Key: 1
	Companies   Companies `json:"companies" msgpack:"companies"`     // Key: 2
	Description string    `json:"description" msgpack:"description"` // Key: 3
	IsOther     bool      `json:"isOther" msgpack:"isOther"`         // Key: 4
}

type ConcertMaster struct {
	Id            int64     `json:"id" msgpack:"id"`                       // Key: 0
	StartDate     time.Time `json:"startDate" msgpack:"startDate"`         // Key: 2
	EndDate       time.Time `json:"endDate" msgpack:"endDate"`             // Key: 3
	Name          string    `json:"name" msgpack:"name"`                   // Key: 4
	EventMasterId int64     `json:"eventMasterId" msgpack:"eventMasterId"` // Key: 5
}

type ConcertStageMaster struct {
	Id                    int64                 `json:"id" msgpack:"id"`                                       // Key: 0
	ConcertMasterId       int64                 `json:"concertMasterId" msgpack:"concertMasterId"`             // Key: 1
	MusicMasterId         int64                 `json:"musicMasterId" msgpack:"musicMasterId"`                 // Key: 2
	VocalVersion          int                   `json:"vocalVersion" msgpack:"vocalVersion"`                   // Key: 3
	SenseNotationMasterId int64                 `json:"senseNotationMasterId" msgpack:"senseNotationMasterId"` // Key: 4
	RecommendedCompany    *Companies            `json:"recommendedCompany" msgpack:"recommendedCompany"`       // Key: 5
	RecommendedLevel      int                   `json:"recommendedLevel" msgpack:"recommendedLevel"`           // Key: 6
	ClearScore            int64                 `json:"clearScore" msgpack:"clearScore"`                       // Key: 7
	Rewards               []ConcertRewardMaster `json:"rewards" msgpack:"rewards"`                             // Key: 8
}

type ConcertRewardMaster struct {
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type ConcoursMaster struct {
	Id            int64                  `json:"id" msgpack:"id"`                       // Key: 0
	EventMasterId int64                  `json:"eventMasterId" msgpack:"eventMasterId"` // Key: 1
	BgmPath       string                 `json:"bgmPath" msgpack:"bgmPath"`             // Key: 2
	Details       []ConcoursDetailMaster `json:"details" msgpack:"details"`             // Key: 3
	Points        []ConcoursPointMaster  `json:"points" msgpack:"points"`               // Key: 4
}

type ConcoursDetailMaster struct {
	Id        int64     `json:"id" msgpack:"id"`               // Key: 0
	Name      string    `json:"name" msgpack:"name"`           // Key: 1
	StartDate time.Time `json:"startDate" msgpack:"startDate"` // Key: 2
	EndDate   time.Time `json:"endDate" msgpack:"endDate"`     // Key: 3
}

type ConcoursPointMaster struct {
	TargetHighestRank int `json:"targetHighestRank" msgpack:"targetHighestRank"` // Key: 0
	TargetLowestRank  int `json:"targetLowestRank" msgpack:"targetLowestRank"`   // Key: 1
	Point             int `json:"point" msgpack:"point"`                         // Key: 2
}

type ConcoursPointRewardMaster struct {
	Id               int64                            `json:"id" msgpack:"id"`                             // Key: 0
	ConcoursMasterId int64                            `json:"concoursMasterId" msgpack:"concoursMasterId"` // Key: 1
	RequiredPoint    int                              `json:"requiredPoint" msgpack:"requiredPoint"`       // Key: 2
	Things           []ConcoursPointRewardThingMaster `json:"things" msgpack:"things"`                     // Key: 3
}

type ConcoursPointRewardThingMaster struct {
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type CostumeGroupMaster struct {
	Id                                    int64     `json:"id" msgpack:"id"`                                                                       // Key: 0
	Name                                  string    `json:"name" msgpack:"name"`                                                                   // Key: 1
	Description                           string    `json:"description" msgpack:"description"`                                                     // Key: 2
	Order                                 int       `json:"order" msgpack:"order"`                                                                 // Key: 3
	CostumeMasterIds                      []int64   `json:"costumeMasterIds" msgpack:"costumeMasterIds"`                                           // Key: 4
	CostumeWearableCharacterGroupMasterId *int64    `json:"costumeWearableCharacterGroupMasterId" msgpack:"costumeWearableCharacterGroupMasterId"` // Key: 5
	PerformanceGroupId                    *int      `json:"performanceGroupId" msgpack:"performanceGroupId"`                                       // Key: 6
	IsHairChange                          bool      `json:"isHairChange" msgpack:"isHairChange"`                                                   // Key: 7
	DisplayStartAt                        time.Time `json:"displayStartAt" msgpack:"displayStartAt"`                                               // Key: 8
	DisplayEndAt                          time.Time `json:"displayEndAt" msgpack:"displayEndAt"`                                                   // Key: 9
}

type CostumeMaster struct {
	Id                   int64  `json:"id" msgpack:"id"`                                     // Key: 0
	Name                 string `json:"name" msgpack:"name"`                                 // Key: 1
	Order                int    `json:"order" msgpack:"order"`                               // Key: 2
	IsDefault            bool   `json:"isDefault" msgpack:"isDefault"`                       // Key: 3
	CostumeGroupMasterId int64  `json:"costumeGroupMasterId" msgpack:"costumeGroupMasterId"` // Key: 4
	Description          string `json:"description" msgpack:"description"`                   // Key: 5
}

type CostumeWearableCharacterGroupMaster struct {
	Id                     int64   `json:"id" msgpack:"id"`                                         // Key: 0
	CharacterBaseMasterIds []int64 `json:"characterBaseMasterIds" msgpack:"characterBaseMasterIds"` // Key: 1
}

type CourseRankingRewardMaster struct {
	Id                  int64                            `json:"id" msgpack:"id"`                                   // Key: 0
	MusicCourseMasterId int64                            `json:"musicCourseMasterId" msgpack:"musicCourseMasterId"` // Key: 1
	TargetHighestRank   int                              `json:"targetHighestRank" msgpack:"targetHighestRank"`     // Key: 2
	TargetLowestRank    int                              `json:"targetLowestRank" msgpack:"targetLowestRank"`       // Key: 3
	Rewards             []CourseRankingRewardThingMaster `json:"rewards" msgpack:"rewards"`                         // Key: 4
}

type CourseRankingRewardThingMaster struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 2
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 3
}

type DecorationMaster struct {
	Id         int64                `json:"id" msgpack:"id"`                 // Key: 0
	Order      int                  `json:"order" msgpack:"order"`           // Key: 1
	IsDefault  bool                 `json:"isDefault" msgpack:"isDefault"`   // Key: 2
	Name       string               `json:"name" msgpack:"name"`             // Key: 3
	Category   DecorationCategories `json:"category" msgpack:"category"`     // Key: 4
	WidhtPx    int                  `json:"widhtPx" msgpack:"widhtPx"`       // Key: 5
	HeightPx   int                  `json:"heightPx" msgpack:"heightPx"`     // Key: 6
	MinWidthPx int                  `json:"minWidthPx" msgpack:"minWidthPx"` // Key: 7
	MaxWidthPx int                  `json:"maxWidthPx" msgpack:"maxWidthPx"` // Key: 8
}

type DugongRunCourseMaster struct {
	Id                     int64                    `json:"id" msgpack:"id"`                                         // Key: 0
	EventMasterId          *int64                   `json:"eventMasterId" msgpack:"eventMasterId"`                   // Key: 1
	Difficulty             DugongRunDifficultyTypes `json:"difficulty" msgpack:"difficulty"`                         // Key: 2
	MusicMasterId          int64                    `json:"musicMasterId" msgpack:"musicMasterId"`                   // Key: 3
	NotationPath           string                   `json:"notationPath" msgpack:"notationPath"`                     // Key: 4
	ScrollSpeed            int                      `json:"scrollSpeed" msgpack:"scrollSpeed"`                       // Key: 5
	DugongRunCourseGroupId int64                    `json:"dugongRunCourseGroupId" msgpack:"dugongRunCourseGroupId"` // Key: 6
}

type DugongRunRewardMaster struct {
	Id                       int64      `json:"id" msgpack:"id"`                                             // Key: 0
	RequiredClearCourseCount int        `json:"requiredClearCourseCount" msgpack:"requiredClearCourseCount"` // Key: 1
	ThingType                ThingTypes `json:"thingType" msgpack:"thingType"`                               // Key: 2
	ThingId                  int64      `json:"thingId" msgpack:"thingId"`                                   // Key: 3
	ThingQuantity            int        `json:"thingQuantity" msgpack:"thingQuantity"`                       // Key: 4
	DugongRunCourseGroupId   int64      `json:"dugongRunCourseGroupId" msgpack:"dugongRunCourseGroupId"`     // Key: 5
}

type EffectConditionMaster struct {
	Condition EffectConditions `json:"condition" msgpack:"condition"` // Key: 0
	Value     *int64           `json:"value" msgpack:"value"`         // Key: 1
}

type EffectDetailMaster struct {
	Level int     `json:"level" msgpack:"level"` // Key: 0
	Value float32 `json:"value" msgpack:"value"` // Key: 1
}

type EffectDurationGroupMaster struct {
	Id        int64                  `json:"id" msgpack:"id"`               // Key: 0
	Durations []EffectDurationMaster `json:"durations" msgpack:"durations"` // Key: 1
}

type EffectDurationMaster struct {
	Level           int     `json:"level" msgpack:"level"`                     // Key: 0
	DurationSeconds float32 `json:"durationSeconds" msgpack:"durationSeconds"` // Key: 1
}

type EffectMaster struct {
	Id              int64                   `json:"id" msgpack:"id"`                           // Key: 0
	Type            EffectTypes             `json:"type" msgpack:"type"`                       // Key: 1
	Range           EffectTargetRanges      `json:"range" msgpack:"range"`                     // Key: 2
	CalculationType CalculationTypes        `json:"calculationType" msgpack:"calculationType"` // Key: 3
	Details         []EffectDetailMaster    `json:"details" msgpack:"details"`                 // Key: 4
	Conditions      []EffectConditionMaster `json:"conditions" msgpack:"conditions"`           // Key: 5
	DurationSecond  int                     `json:"durationSecond" msgpack:"durationSecond"`   // Key: 6
	Triggers        []EffectTriggerMaster   `json:"triggers" msgpack:"triggers"`               // Key: 7
	FireTimingType  FireTimingTypes         `json:"fireTimingType" msgpack:"fireTimingType"`   // Key: 8
}

type EffectTriggerCharacterBaseGroupMaster struct {
	Id                     int64   `json:"id" msgpack:"id"`                                         // Key: 0
	CharacterBaseMasterIds []int64 `json:"characterBaseMasterIds" msgpack:"characterBaseMasterIds"` // Key: 1
}

type EffectTriggerMaster struct {
	Trigger TriggerType `json:"trigger" msgpack:"trigger"` // Key: 0
	Value   *int64      `json:"value" msgpack:"value"`     // Key: 1
}

type EpisodeMaster struct {
	Id                           int64                     `json:"id" msgpack:"id"`                       // Key: 0
	StoryMasterId                int64                     `json:"storyMasterId" msgpack:"storyMasterId"` // Key: 1
	StoryMaster                  *StoryMaster              `json:"storyMaster" msgpack:"storyMaster"`
	Title                        string                    `json:"title" msgpack:"title"`                                               // Key: 2
	Order                        int                       `json:"order" msgpack:"order"`                                               // Key: 3
	EpisodeRewardPackageMasterId int64                     `json:"episodeRewardPackageMasterId" msgpack:"episodeRewardPackageMasterId"` // Key: 4
	Conditions                   []EpisodeReleaseCondition `json:"conditions" msgpack:"conditions"`                                     // Key: 5
	PreEpisodeMasterId           *int64                    `json:"preEpisodeMasterId" msgpack:"preEpisodeMasterId"`                     // Key: 6
	DisplayStartDate             *time.Time                `json:"displayStartDate" msgpack:"displayStartDate"`                         // Key: 7
	DisplayEndDate               *time.Time                `json:"displayEndDate" msgpack:"displayEndDate"`                             // Key: 8
}

type EpisodeReleaseCondition struct {
	ConditionType EpisodeReleaseConditionTypes `json:"conditionType" msgpack:"conditionType"` // Key: 0
	Value1        *int64                       `json:"value1" msgpack:"value1"`               // Key: 1
	Value2        *int64                       `json:"value2" msgpack:"value2"`               // Key: 2
	Value3        *int64                       `json:"value3" msgpack:"value3"`               // Key: 3
}

type EpisodeRewardPackageMaster struct {
	Id      int64                `json:"id" msgpack:"id"`           // Key: 0
	Rewards []EpisodeRewardThing `json:"rewards" msgpack:"rewards"` // Key: 1
}

type EpisodeRewardThing struct {
	ThingId                int64      `json:"thingId" msgpack:"thingId"`                               // Key: 0
	ThingType              ThingTypes `json:"thingType" msgpack:"thingType"`                           // Key: 1
	ThingQuantity          int        `json:"thingQuantity" msgpack:"thingQuantity"`                   // Key: 2
	IsOnlyStoryEventTerm   bool       `json:"isOnlyStoryEventTerm" msgpack:"isOnlyStoryEventTerm"`     // Key: 3
	IsChapterAllReadReward bool       `json:"isChapterAllReadReward" msgpack:"isChapterAllReadReward"` // Key: 4
}

type EventBonusMaster struct {
	Id              int64                `json:"id" msgpack:"id"`                           // Key: 0
	BonusTargetType StoryEventBonusTypes `json:"bonusTargetType" msgpack:"bonusTargetType"` // Key: 1
	BonusTargetId   int64                `json:"bonusTargetId" msgpack:"bonusTargetId"`     // Key: 2
}

type EventBoxGachaBoxMaster struct {
	Id                    int64                         `json:"id" msgpack:"id"`                                       // Key: 0
	EventBoxGachaMasterId int64                         `json:"eventBoxGachaMasterId" msgpack:"eventBoxGachaMasterId"` // Key: 1
	Order                 int                           `json:"order" msgpack:"order"`                                 // Key: 2
	BoxThings             []EventBoxGachaBoxThingMaster `json:"boxThings" msgpack:"boxThings"`                         // Key: 3
}

type EventBoxGachaBoxThingMaster struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 2
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 3
	HitLimit      int        `json:"hitLimit" msgpack:"hitLimit"`           // Key: 4
	IsResettable  bool       `json:"isResettable" msgpack:"isResettable"`   // Key: 5
}

type EventBoxGachaMaster struct {
	Id                                int64                       `json:"id" msgpack:"id"`                                                               // Key: 0
	Name                              string                      `json:"name" msgpack:"name"`                                                           // Key: 1
	EventMasterId                     int64                       `json:"eventMasterId" msgpack:"eventMasterId"`                                         // Key: 2
	RequiredItemMasterId              int64                       `json:"requiredItemMasterId" msgpack:"requiredItemMasterId"`                           // Key: 3
	RequiredQuantity                  int                         `json:"requiredQuantity" msgpack:"requiredQuantity"`                                   // Key: 4
	Details                           []EventBoxGachaDetailMaster `json:"details" msgpack:"details"`                                                     // Key: 5
	EventBoxGachaTextTemplateMasterId int64                       `json:"eventBoxGachaTextTemplateMasterId" msgpack:"eventBoxGachaTextTemplateMasterId"` // Key: 6
}

type EventBoxGachaDetailMaster struct {
	Id         int64 `json:"id" msgpack:"id"`                 // Key: 0
	Order      int   `json:"order" msgpack:"order"`           // Key: 1
	PrizeCount int   `json:"prizeCount" msgpack:"prizeCount"` // Key: 2
}

type EventBoxGachaTextTemplateMaster struct {
	Id          int64  `json:"id" msgpack:"id"`                   // Key: 0
	Description string `json:"description" msgpack:"description"` // Key: 1
}

type EventCampClassMaster struct {
	Id                 int64 `json:"id" msgpack:"id"`                                 // Key: 0
	EventMasterId      int64 `json:"eventMasterId" msgpack:"eventMasterId"`           // Key: 1
	TargetMaxRank      int   `json:"targetMaxRank" msgpack:"targetMaxRank"`           // Key: 2
	TargetMinRank      int   `json:"targetMinRank" msgpack:"targetMinRank"`           // Key: 3
	RewardReceiveRatio int   `json:"rewardReceiveRatio" msgpack:"rewardReceiveRatio"` // Key: 4
}

type EventCampMaster struct {
	Id                         int64  `json:"id" msgpack:"id"`                                                 // Key: 0
	EventMasterId              int64  `json:"eventMasterId" msgpack:"eventMasterId"`                           // Key: 1
	WinCampRatio               int    `json:"winCampRatio" msgpack:"winCampRatio"`                             // Key: 2
	LoseCampRatio              int    `json:"loseCampRatio" msgpack:"loseCampRatio"`                           // Key: 3
	NoneCampRatio              int    `json:"noneCampRatio" msgpack:"noneCampRatio"`                           // Key: 4
	Camp1TeamName              string `json:"camp1TeamName" msgpack:"camp1TeamName"`                           // Key: 5
	Camp2TeamName              string `json:"camp2TeamName" msgpack:"camp2TeamName"`                           // Key: 6
	CampSelectMessage          string `json:"campSelectMessage" msgpack:"campSelectMessage"`                   // Key: 7
	BgmPath                    string `json:"bgmPath" msgpack:"bgmPath"`                                       // Key: 8
	CharacterMasterIdIconLeft  int64  `json:"characterMasterIdIconLeft" msgpack:"characterMasterIdIconLeft"`   // Key: 9
	CharacterMasterIdIconRight int64  `json:"characterMasterIdIconRight" msgpack:"characterMasterIdIconRight"` // Key: 10
}

type EventCampSupportPointRewardMaster struct {
	Id                        int64 `json:"id" msgpack:"id"`                                               // Key: 0
	EventMasterId             int64 `json:"eventMasterId" msgpack:"eventMasterId"`                         // Key: 1
	TargetMinRank             int   `json:"targetMinRank" msgpack:"targetMinRank"`                         // Key: 2
	TargetMaxRank             int   `json:"targetMaxRank" msgpack:"targetMaxRank"`                         // Key: 3
	Camp1RewardTrophyMasterId int64 `json:"camp1RewardTrophyMasterId" msgpack:"camp1RewardTrophyMasterId"` // Key: 4
	Camp2RewardTrophyMasterId int64 `json:"camp2RewardTrophyMasterId" msgpack:"camp2RewardTrophyMasterId"` // Key: 5
}

type EventMaster struct {
	Id                      int64                      `json:"id" msgpack:"id"`                                           // Key: 0
	EventType               EventTypes                 `json:"eventType" msgpack:"eventType"`                             // Key: 1
	StartDate               time.Time                  `json:"startDate" msgpack:"startDate"`                             // Key: 2
	EndDate                 time.Time                  `json:"endDate" msgpack:"endDate"`                                 // Key: 3
	ForceEndDate            time.Time                  `json:"forceEndDate" msgpack:"forceEndDate"`                       // Key: 4
	Name                    string                     `json:"name" msgpack:"name"`                                       // Key: 5
	IsAsideLiveButton       bool                       `json:"isAsideLiveButton" msgpack:"isAsideLiveButton"`             // Key: 6
	EventPointItemMasterId  *int64                     `json:"eventPointItemMasterId" msgpack:"eventPointItemMasterId"`   // Key: 7
	BonusAttribute          *Attributes                `json:"bonusAttribute" msgpack:"bonusAttribute"`                   // Key: 8
	BonusCategoryMasterId10 *int64                     `json:"bonusCategoryMasterId10" msgpack:"bonusCategoryMasterId10"` // Key: 9
	BonusCategoryMasterId20 *int64                     `json:"bonusCategoryMasterId20" msgpack:"bonusCategoryMasterId20"` // Key: 10
	Bonuses                 []EventBonusMaster         `json:"bonuses" msgpack:"bonuses"`                                 // Key: 11
	PointRewards            []EventPointRewardMaster   `json:"pointRewards" msgpack:"pointRewards"`                       // Key: 12
	RankingRewards          []EventRankingRewardMaster `json:"rankingRewards" msgpack:"rankingRewards"`                   // Key: 13
	ExchangeShopMasterId    *int64                     `json:"exchangeShopMasterId" msgpack:"exchangeShopMasterId"`       // Key: 14
	AsideLiveButtonType     AsideLiveButtonTypes       `json:"asideLiveButtonType" msgpack:"asideLiveButtonType"`         // Key: 15
	SecondaryBonusAttribute *Attributes                `json:"secondaryBonusAttribute" msgpack:"secondaryBonusAttribute"` // Key: 16
}

type EventPointRewardMaster struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	RequiredPoint int        `json:"requiredPoint" msgpack:"requiredPoint"` // Key: 1
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 2
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 3
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 4
}

type EventRankingRewardMaster struct {
	Id                int64                           `json:"id" msgpack:"id"`                               // Key: 0
	TargetHighestRank int                             `json:"targetHighestRank" msgpack:"targetHighestRank"` // Key: 1
	TargetLowestRank  int                             `json:"targetLowestRank" msgpack:"targetLowestRank"`   // Key: 2
	Rewards           []EventRankingRewardThingMaster `json:"rewards" msgpack:"rewards"`                     // Key: 3
}

type EventRankingRewardThingMaster struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 2
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 3
	IsCampReward  bool       `json:"isCampReward" msgpack:"isCampReward"`   // Key: 4
}

type ExchangeShopMaster struct {
	Id                          int64               `json:"id" msgpack:"id"`                                                   // Key: 0
	IsDisplayRequiredHavingItem bool                `json:"isDisplayRequiredHavingItem" msgpack:"isDisplayRequiredHavingItem"` // Key: 1
	Category                    ShopCategories      `json:"category" msgpack:"category"`                                       // Key: 2
	Name                        string              `json:"name" msgpack:"name"`                                               // Key: 3
	DisplayThingType            ThingTypes          `json:"displayThingType" msgpack:"displayThingType"`                       // Key: 4
	DisplayItemMasterId         *int64              `json:"displayItemMasterId" msgpack:"displayItemMasterId"`                 // Key: 5
	DisplayItemMaster           *ItemMaster         `json:"displayItemMaster" msgpack:"displayItemMaster"`
	BannerPath                  string              `json:"bannerPath" msgpack:"bannerPath"`           // Key: 6
	StartDate                   *time.Time          `json:"startDate" msgpack:"startDate"`             // Key: 7
	EndDate                     *time.Time          `json:"endDate" msgpack:"endDate"`                 // Key: 8
	LastRefreshedAt             time.Time           `json:"lastRefreshedAt" msgpack:"lastRefreshedAt"` // Key: 9
	Lineup                      []ExchangeShopThing `json:"lineup" msgpack:"lineup"`                   // Key: 10
	Order                       int                 `json:"order" msgpack:"order"`                     // Key: 11
}

type ExchangeShopThing struct {
	Id                   int64             `json:"id" msgpack:"id"`                                     // Key: 0
	ThingId              int64             `json:"thingId" msgpack:"thingId"`                           // Key: 1
	ThingType            ThingTypes        `json:"thingType" msgpack:"thingType"`                       // Key: 2
	ThingQuantity        int               `json:"thingQuantity" msgpack:"thingQuantity"`               // Key: 3
	ExchangeLimit        *int              `json:"exchangeLimit" msgpack:"exchangeLimit"`               // Key: 4
	Order                int               `json:"order" msgpack:"order"`                               // Key: 5
	ReplaceType          *ShopReplaceTypes `json:"replaceType" msgpack:"replaceType"`                   // Key: 6
	IsDisplayLocked      bool              `json:"isDisplayLocked" msgpack:"isDisplayLocked"`           // Key: 7
	UnlockType           *ShopUnlockTypes  `json:"unlockType" msgpack:"unlockType"`                     // Key: 8
	UnlockValue          *int64            `json:"unlockValue" msgpack:"unlockValue"`                   // Key: 9
	StartDate            *time.Time        `json:"startDate" msgpack:"startDate"`                       // Key: 10
	EndDate              *time.Time        `json:"endDate" msgpack:"endDate"`                           // Key: 11
	RequiredItemMasterId int64             `json:"requiredItemMasterId" msgpack:"requiredItemMasterId"` // Key: 12
	RequiredQuantity     int               `json:"requiredQuantity" msgpack:"requiredQuantity"`         // Key: 13
}

type FacialExpressionMaster struct {
	Id       int64  `json:"id" msgpack:"id"`             // Key: 0
	EyeBrow  string `json:"eyeBrow" msgpack:"eyeBrow"`   // Key: 1
	Eye      string `json:"eye" msgpack:"eye"`           // Key: 2
	EyeBlink string `json:"eyeBlink" msgpack:"eyeBlink"` // Key: 3
	Cheek    string `json:"cheek" msgpack:"cheek"`       // Key: 4
	Mouth    string `json:"mouth" msgpack:"mouth"`       // Key: 5
}

type FilmItemMaster struct {
	Id                          int64  `json:"id" msgpack:"id"`                                                   // Key: 0
	ItemMasterId                int64  `json:"itemMasterId" msgpack:"itemMasterId"`                               // Key: 1
	TargetCharacterBaseMasterId *int64 `json:"targetCharacterBaseMasterId" msgpack:"targetCharacterBaseMasterId"` // Key: 2
}

type FriendInvitationMissionMaster struct {
	Id           int64                                `json:"id" msgpack:"id"`                     // Key: 0
	Description  string                               `json:"description" msgpack:"description"`   // Key: 1
	Order        int                                  `json:"order" msgpack:"order"`               // Key: 2
	StartDate    time.Time                            `json:"startDate" msgpack:"startDate"`       // Key: 3
	EndDate      time.Time                            `json:"endDate" msgpack:"endDate"`           // Key: 4
	JumpType     JumpTypes                            `json:"jumpType" msgpack:"jumpType"`         // Key: 5
	JumpTargetId *int64                               `json:"jumpTargetId" msgpack:"jumpTargetId"` // Key: 6
	Stages       []FriendInvitationMissionStageMaster `json:"stages" msgpack:"stages"`             // Key: 7
}

type FriendInvitationMissionStageMaster struct {
	Id                int64                                 `json:"id" msgpack:"id"`                               // Key: 0
	GoalCount         int                                   `json:"goalCount" msgpack:"goalCount"`                 // Key: 1
	MissionStageOrder int                                   `json:"missionStageOrder" msgpack:"missionStageOrder"` // Key: 2
	Rewards           []FriendInvitationMissionRewardMaster `json:"rewards" msgpack:"rewards"`                     // Key: 3
}

type FriendInvitationMissionRewardMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type GachaBonusThing struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 2
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 3
}

type GachaDetailBonusThing struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 2
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 3
}

type GachaDetailMaster struct {
	Id                     int64                   `json:"id" msgpack:"id"`                                         // Key: 0
	RequiredTicketMItemId  *int64                  `json:"requiredTicketMItemId" msgpack:"requiredTicketMItemId"`   // Key: 1
	RequiredTicketQuantity *int                    `json:"requiredTicketQuantity" msgpack:"requiredTicketQuantity"` // Key: 2
	FreeJewelAmount        *int                    `json:"freeJewelAmount" msgpack:"freeJewelAmount"`               // Key: 3
	PaidJewelAmount        *int                    `json:"paidJewelAmount" msgpack:"paidJewelAmount"`               // Key: 4
	IsFree                 bool                    `json:"isFree" msgpack:"isFree"`                                 // Key: 5
	DailyRollLimit         *int                    `json:"dailyRollLimit" msgpack:"dailyRollLimit"`                 // Key: 6
	OverallRollLimit       *int                    `json:"overallRollLimit" msgpack:"overallRollLimit"`             // Key: 7
	PrizeCount             int                     `json:"prizeCount" msgpack:"prizeCount"`                         // Key: 8
	FixedPrizeCount        int                     `json:"fixedPrizeCount" msgpack:"fixedPrizeCount"`               // Key: 9
	ButtonType             GachaButtonTypes        `json:"buttonType" msgpack:"buttonType"`                         // Key: 10
	DetailBonusThings      []GachaDetailBonusThing `json:"detailBonusThings" msgpack:"detailBonusThings"`           // Key: 11
}

type GachaMaster struct {
	Id                                 int64                         `json:"id" msgpack:"id"`                                                                 // Key: 0
	Name                               string                        `json:"name" msgpack:"name"`                                                             // Key: 1
	CardType                           GachaCardTypes                `json:"cardType" msgpack:"cardType"`                                                     // Key: 2
	GachaType                          GachaTypes                    `json:"gachaType" msgpack:"gachaType"`                                                   // Key: 3
	GachaTextTemplateMasterId          int64                         `json:"gachaTextTemplateMasterId" msgpack:"gachaTextTemplateMasterId"`                   // Key: 4
	StartDate                          time.Time                     `json:"startDate" msgpack:"startDate"`                                                   // Key: 5
	EndDate                            time.Time                     `json:"endDate" msgpack:"endDate"`                                                       // Key: 6
	BonusThings                        []GachaBonusThing             `json:"bonusThings" msgpack:"bonusThings"`                                               // Key: 7
	GachaDetails                       []GachaDetailMaster           `json:"gachaDetails" msgpack:"gachaDetails"`                                             // Key: 8
	Things                             []GachaThing                  `json:"things" msgpack:"things"`                                                         // Key: 9
	HasMovie                           bool                          `json:"hasMovie" msgpack:"hasMovie"`                                                     // Key: 10
	StoryEventMasterId                 *int64                        `json:"storyEventMasterId" msgpack:"storyEventMasterId"`                                 // Key: 11
	AttentionGachaTextTemplateMasterId int64                         `json:"attentionGachaTextTemplateMasterId" msgpack:"attentionGachaTextTemplateMasterId"` // Key: 12
	IsForceDisplay                     bool                          `json:"isForceDisplay" msgpack:"isForceDisplay"`                                         // Key: 13
	IsHideEndDate                      bool                          `json:"isHideEndDate" msgpack:"isHideEndDate"`                                           // Key: 14
	UnlockType                         GachaUnlockTypes              `json:"unlockType" msgpack:"unlockType"`                                                 // Key: 15
	UnlockValue                        *int64                        `json:"unlockValue" msgpack:"unlockValue"`                                               // Key: 16
	ExchangeShopBannerPath             string                        `json:"exchangeShopBannerPath" msgpack:"exchangeShopBannerPath"`                         // Key: 17
	Order                              int64                         `json:"order" msgpack:"order"`                                                           // Key: 18
	DisplayFooterCondition             *GachaDisplayFooterConditions `json:"displayFooterCondition" msgpack:"displayFooterCondition"`                         // Key: 19
	GroupType                          GachaGroupTypes               `json:"groupType" msgpack:"groupType"`                                                   // Key: 20
	RollBonuses                        []GachaRollBonus              `json:"rollBonuses" msgpack:"rollBonuses"`                                               // Key: 21
	ReRollLimit                        *int                          `json:"reRollLimit" msgpack:"reRollLimit"`                                               // Key: 22
	ReplaceAssetName                   string                        `json:"replaceAssetName" msgpack:"replaceAssetName"`                                     // Key: 23
}

type GachaRollBonus struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	RollCount     int        `json:"rollCount" msgpack:"rollCount"`         // Key: 1
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 2
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 3
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 4
}

type GachaTextTemplateMaster struct {
	Id          int64  `json:"id" msgpack:"id"`                   // Key: 0
	Description string `json:"description" msgpack:"description"` // Key: 1
}

type GachaThing struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 2
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 3
	PickupOrder   *int       `json:"pickupOrder" msgpack:"pickupOrder"`     // Key: 4
	IsSelectable  bool       `json:"isSelectable" msgpack:"isSelectable"`   // Key: 5
}

type GameHintMaster struct {
	Id           int64          `json:"id" msgpack:"id"`                     // Key: 0
	PageCategory PageCategories `json:"pageCategory" msgpack:"pageCategory"` // Key: 1
	PageCount    int            `json:"pageCount" msgpack:"pageCount"`       // Key: 2
}

type GhostLiveMaster struct {
	Id                    int64     `json:"id" msgpack:"id"`                                       // Key: 0
	MusicMasterId         int64     `json:"musicMasterId" msgpack:"musicMasterId"`                 // Key: 1
	VocalVersion          int       `json:"vocalVersion" msgpack:"vocalVersion"`                   // Key: 2
	SenseNotationMasterId int64     `json:"senseNotationMasterId" msgpack:"senseNotationMasterId"` // Key: 3
	StartDate             time.Time `json:"startDate" msgpack:"startDate"`                         // Key: 4
	EndDate               time.Time `json:"endDate" msgpack:"endDate"`                             // Key: 5
}

type GradualMissionGroupMaster struct {
	Id        int64                  `json:"id" msgpack:"id"`               // Key: 0
	StartDate time.Time              `json:"startDate" msgpack:"startDate"` // Key: 1
	EndDate   time.Time              `json:"endDate" msgpack:"endDate"`     // Key: 2
	Missions  []GradualMissionMaster `json:"missions" msgpack:"missions"`   // Key: 3
}

type GradualMissionMaster struct {
	Days             int64   `json:"days" msgpack:"days"`                         // Key: 0
	MissionMasterIds []int64 `json:"missionMasterIds" msgpack:"missionMasterIds"` // Key: 2
}

type HeadDirectionMaster struct {
	Id            int64  `json:"id" msgpack:"id"`                       // Key: 0
	DirectionName string `json:"directionName" msgpack:"directionName"` // Key: 1
}

type HeadMotionMaster struct {
	Id         int64  `json:"id" msgpack:"id"`                 // Key: 0
	MotionName string `json:"motionName" msgpack:"motionName"` // Key: 1
}

type HomeBackgroundMaster struct {
	Id        int64      `json:"id" msgpack:"id"`               // Key: 0
	Name      string     `json:"name" msgpack:"name"`           // Key: 1
	StartDate *time.Time `json:"startDate" msgpack:"startDate"` // Key: 2
	EndDate   *time.Time `json:"endDate" msgpack:"endDate"`     // Key: 3
}

type HomeBGMDetailMaster struct {
	Id              int64     `json:"id" msgpack:"id"`                           // Key: 0
	HomeBGMMasterId int64     `json:"homeBGMMasterId" msgpack:"homeBGMMasterId"` // Key: 1
	BGMName         string    `json:"bGMName" msgpack:"bGMName"`                 // Key: 2
	BGMPath         string    `json:"bGMPath" msgpack:"bGMPath"`                 // Key: 3
	StartDate       time.Time `json:"startDate" msgpack:"startDate"`             // Key: 4
	EndDate         time.Time `json:"endDate" msgpack:"endDate"`                 // Key: 5
	MusicMasterId   int64     `json:"musicMasterId" msgpack:"musicMasterId"`     // Key: 6
}

type HomeBGMMaster struct {
	Id             int64                 `json:"id" msgpack:"id"`                         // Key: 0
	Priority       HomeBGMPriorityTypes  `json:"priority" msgpack:"priority"`             // Key: 1
	SelectionType  HomeBGMSelectionTypes `json:"selectionType" msgpack:"selectionType"`   // Key: 2
	DefaultBGMName string                `json:"defaultBGMName" msgpack:"defaultBGMName"` // Key: 3
	DefaultBGMPath string                `json:"defaultBGMPath" msgpack:"defaultBGMPath"` // Key: 4
	StartDate      time.Time             `json:"startDate" msgpack:"startDate"`           // Key: 5
	EndDate        time.Time             `json:"endDate" msgpack:"endDate"`               // Key: 6
}

type HomeCharacterMaster struct {
	Id                    int64     `json:"id" msgpack:"id"`                                       // Key: 0
	CharacterBaseMasterId int64     `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	StartDate             time.Time `json:"startDate" msgpack:"startDate"`                         // Key: 2
	EndDate               time.Time `json:"endDate" msgpack:"endDate"`                             // Key: 3
}

type HomeCharacterVoiceMaster struct {
	Id                           int64   `json:"id" msgpack:"id"`                                                     // Key: 0
	CharacterBaseMasterId        int64   `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`               // Key: 1
	Text                         string  `json:"text" msgpack:"text"`                                                 // Key: 2
	Weight                       int     `json:"weight" msgpack:"weight"`                                             // Key: 3
	CharacterVoicePeriodMasterId *int64  `json:"characterVoicePeriodMasterId" msgpack:"characterVoicePeriodMasterId"` // Key: 5
	IsPlayerBirthDateVoice       bool    `json:"isPlayerBirthDateVoice" msgpack:"isPlayerBirthDateVoice"`             // Key: 7
	VoiceFileName1               string  `json:"voiceFileName1" msgpack:"voiceFileName1"`                             // Key: 8
	VoiceFileName2               string  `json:"voiceFileName2" msgpack:"voiceFileName2"`                             // Key: 9
	VoiceFileName3               string  `json:"voiceFileName3" msgpack:"voiceFileName3"`                             // Key: 10
	VoiceFileName4               string  `json:"voiceFileName4" msgpack:"voiceFileName4"`                             // Key: 11
	VoiceInterval1               float32 `json:"voiceInterval1" msgpack:"voiceInterval1"`                             // Key: 13
	VoiceInterval2               float32 `json:"voiceInterval2" msgpack:"voiceInterval2"`                             // Key: 14
	VoiceInterval3               float32 `json:"voiceInterval3" msgpack:"voiceInterval3"`                             // Key: 15
	MouthMotionId1               string  `json:"mouthMotionId1" msgpack:"mouthMotionId1"`                             // Key: 16
	MouthMotionId2               string  `json:"mouthMotionId2" msgpack:"mouthMotionId2"`                             // Key: 17
	MouthMotionId3               string  `json:"mouthMotionId3" msgpack:"mouthMotionId3"`                             // Key: 18
	MouthMotionId4               string  `json:"mouthMotionId4" msgpack:"mouthMotionId4"`                             // Key: 19
	BodyMotionId1                string  `json:"bodyMotionId1" msgpack:"bodyMotionId1"`                               // Key: 20
	BodyMotionId2                string  `json:"bodyMotionId2" msgpack:"bodyMotionId2"`                               // Key: 21
	BodyMotionId3                string  `json:"bodyMotionId3" msgpack:"bodyMotionId3"`                               // Key: 22
	BodyMotionId4                string  `json:"bodyMotionId4" msgpack:"bodyMotionId4"`                               // Key: 23
}

type HomeCharacterVoicePeriodMaster struct {
	Id         int64  `json:"id" msgpack:"id"`                 // Key: 0
	StartMonth int    `json:"startMonth" msgpack:"startMonth"` // Key: 1
	StartDay   int    `json:"startDay" msgpack:"startDay"`     // Key: 2
	StartTime  string `json:"startTime" msgpack:"startTime"`   // Key: 3
	EndMonth   int    `json:"endMonth" msgpack:"endMonth"`     // Key: 4
	EndDay     int    `json:"endDay" msgpack:"endDay"`         // Key: 5
	EndTime    string `json:"endTime" msgpack:"endTime"`       // Key: 6
}

type HomePosterMaster struct {
	Id               int64     `json:"id" msgpack:"id"`                             // Key: 0
	PosterMasterId   int64     `json:"posterMasterId" msgpack:"posterMasterId"`     // Key: 1
	GachaMasterId    *int64    `json:"gachaMasterId" msgpack:"gachaMasterId"`       // Key: 2
	DisplayStartDate time.Time `json:"displayStartDate" msgpack:"displayStartDate"` // Key: 3
	DisplayEndDate   time.Time `json:"displayEndDate" msgpack:"displayEndDate"`     // Key: 4
}

type HomeSkinMaster struct {
	Id                  int64      `json:"id" msgpack:"id"`                                   // Key: 0
	Name                string     `json:"name" msgpack:"name"`                               // Key: 1
	Order               int        `json:"order" msgpack:"order"`                             // Key: 2
	UnlockConditionText string     `json:"unlockConditionText" msgpack:"unlockConditionText"` // Key: 3
	Hidden              bool       `json:"hidden" msgpack:"hidden"`                           // Key: 4
	IsDefault           bool       `json:"isDefault" msgpack:"isDefault"`                     // Key: 5
	DisplayStartDate    *time.Time `json:"displayStartDate" msgpack:"displayStartDate"`       // Key: 6
}

type IconFrameMaster struct {
	Id                  int64  `json:"id" msgpack:"id"`                                   // Key: 0
	Name                string `json:"name" msgpack:"name"`                               // Key: 1
	Order               int    `json:"order" msgpack:"order"`                             // Key: 2
	UnlockConditionText string `json:"unlockConditionText" msgpack:"unlockConditionText"` // Key: 3
	Hidden              bool   `json:"hidden" msgpack:"hidden"`                           // Key: 4
	IsDefault           bool   `json:"isDefault" msgpack:"isDefault"`                     // Key: 5
}

type ItemMaster struct {
	Id             int64              `json:"id" msgpack:"id"`                         // Key: 0
	Name           string             `json:"name" msgpack:"name"`                     // Key: 1
	Description    string             `json:"description" msgpack:"description"`       // Key: 2
	DisplayOrder   int64              `json:"displayOrder" msgpack:"displayOrder"`     // Key: 3
	DisplayEndDate *time.Time         `json:"displayEndDate" msgpack:"displayEndDate"` // Key: 4
	MaxStock       int                `json:"maxStock" msgpack:"maxStock"`             // Key: 5
	Category       ItemCategories     `json:"category" msgpack:"category"`             // Key: 6
	Consumable     bool               `json:"consumable" msgpack:"consumable"`         // Key: 7
	JumpType       *JumpTypes         `json:"jumpType" msgpack:"jumpType"`             // Key: 8
	JumpTargetId   *int64             `json:"jumpTargetId" msgpack:"jumpTargetId"`     // Key: 9
	TabCategory    TabCategories      `json:"tabCategory" msgpack:"tabCategory"`       // Key: 10
	Rarity         PossessionRarities `json:"rarity" msgpack:"rarity"`                 // Key: 11
}

type JewelShopCategoryMaster struct {
	Id               int64  `json:"id" msgpack:"id"`                             // Key: 0
	Name             string `json:"name" msgpack:"name"`                         // Key: 1
	LayoutType       int    `json:"layoutType" msgpack:"layoutType"`             // Key: 2
	IsDisplayLocking *int   `json:"isDisplayLocking" msgpack:"isDisplayLocking"` // Key: 3
	BannerPath       string `json:"bannerPath" msgpack:"bannerPath"`             // Key: 4
	Order            int    `json:"order" msgpack:"order"`                       // Key: 5
}

type JewelShopItemMaster struct {
	Id                       int64                    `json:"id" msgpack:"id"`                                     // Key: 0
	Order                    int                      `json:"order" msgpack:"order"`                               // Key: 1
	AppStoreProductId        string                   `json:"appStoreProductId" msgpack:"appStoreProductId"`       // Key: 2
	GooglePlayProductId      string                   `json:"googlePlayProductId" msgpack:"googlePlayProductId"`   // Key: 3
	Name                     string                   `json:"name" msgpack:"name"`                                 // Key: 4
	PackText                 string                   `json:"packText" msgpack:"packText"`                         // Key: 5
	MJewelShopCategoryId     int64                    `json:"mJewelShopCategoryId" msgpack:"mJewelShopCategoryId"` // Key: 6
	MJewelShopCategoryMaster *JewelShopCategoryMaster `json:"mJewelShopCategoryMaster" msgpack:"mJewelShopCategoryMaster"`
	DialogType               int                      `json:"dialogType" msgpack:"dialogType"`                             // Key: 7
	PurchaseType             ShopPurchaseTypes        `json:"purchaseType" msgpack:"purchaseType"`                         // Key: 8
	PurchaseValue            *int                     `json:"purchaseValue" msgpack:"purchaseValue"`                       // Key: 9
	SubscriptionFlag         bool                     `json:"subscriptionFlag" msgpack:"subscriptionFlag"`                 // Key: 10
	GivePaidJewel            *int                     `json:"givePaidJewel" msgpack:"givePaidJewel"`                       // Key: 11
	ReplaceTypes             *ShopReplaceTypes        `json:"replaceTypes" msgpack:"replaceTypes"`                         // Key: 12
	ReplaceValue             *int                     `json:"replaceValue" msgpack:"replaceValue"`                         // Key: 13
	PurchaseLimit            *int                     `json:"purchaseLimit" msgpack:"purchaseLimit"`                       // Key: 14
	ExpiredDay               *int                     `json:"expiredDay" msgpack:"expiredDay"`                             // Key: 15
	CanPurchaseDayOfWeek     string                   `json:"canPurchaseDayOfWeek" msgpack:"canPurchaseDayOfWeek"`         // Key: 16
	IsBuyNotDisplay          bool                     `json:"isBuyNotDisplay" msgpack:"isBuyNotDisplay"`                   // Key: 17
	SaleType                 *SaleTypes               `json:"saleType" msgpack:"saleType"`                                 // Key: 18
	SaleValue                *int                     `json:"saleValue" msgpack:"saleValue"`                               // Key: 19
	IsDisplayLocking         bool                     `json:"isDisplayLocking" msgpack:"isDisplayLocking"`                 // Key: 20
	UnlockType               *JewelShopUnlockTypes    `json:"unlockType" msgpack:"unlockType"`                             // Key: 21
	UnlockValue              *int64                   `json:"unlockValue" msgpack:"unlockValue"`                           // Key: 22
	StartDate                *time.Time               `json:"startDate" msgpack:"startDate"`                               // Key: 23
	EndDate                  *time.Time               `json:"endDate" msgpack:"endDate"`                                   // Key: 24
	Lineup                   []JewelShopThing         `json:"lineup" msgpack:"lineup"`                                     // Key: 25
	ShopItemType             ShopItemTypes            `json:"shopItemType" msgpack:"shopItemType"`                         // Key: 26
	ShopItemValue            *int64                   `json:"shopItemValue" msgpack:"shopItemValue"`                       // Key: 27
	ComebackCampaignMasterId *int64                   `json:"comebackCampaignMasterId" msgpack:"comebackCampaignMasterId"` // Key: 28
	ItemIconBodyImagePath    string                   `json:"itemIconBodyImagePath" msgpack:"itemIconBodyImagePath"`       // Key: 29
	ItemIconDetailImagePath  string                   `json:"itemIconDetailImagePath" msgpack:"itemIconDetailImagePath"`   // Key: 30
	ItemIconBadgeImagePath   string                   `json:"itemIconBadgeImagePath" msgpack:"itemIconBadgeImagePath"`     // Key: 31
}

type JewelShopThing struct {
	Id                   int64                `json:"id" msgpack:"id"`                             // Key: 0
	MJewelShopItemId     int64                `json:"mJewelShopItemId" msgpack:"mJewelShopItemId"` // Key: 1
	MJewelShopItemMaster *JewelShopItemMaster `json:"mJewelShopItemMaster" msgpack:"mJewelShopItemMaster"`
	ThingId              int64                `json:"thingId" msgpack:"thingId"`             // Key: 2
	ThingType            ThingTypes           `json:"thingType" msgpack:"thingType"`         // Key: 3
	ThingQuantity        int                  `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 4
}

type LeaderSenseMaster struct {
	Id          int64                     `json:"id" msgpack:"id"`                   // Key: 0
	Description string                    `json:"description" msgpack:"description"` // Key: 1
	Details     []LeaderSenseDetailMaster `json:"details" msgpack:"details"`         // Key: 2
}

type LeaderSenseDetailMaster struct {
	EffectMasterId int64                              `json:"effectMasterId" msgpack:"effectMasterId"` // Key: 0
	Conditions     []LeaderSenseDetailConditionMaster `json:"conditions" msgpack:"conditions"`         // Key: 1
}

type LeaderSenseDetailConditionMaster struct {
	CategoryMasterId1 *int64 `json:"categoryMasterId1" msgpack:"categoryMasterId1"` // Key: 0
	CategoryMasterId2 *int64 `json:"categoryMasterId2" msgpack:"categoryMasterId2"` // Key: 1
	CategoryMasterId3 *int64 `json:"categoryMasterId3" msgpack:"categoryMasterId3"` // Key: 2
	CategoryMasterId4 *int64 `json:"categoryMasterId4" msgpack:"categoryMasterId4"` // Key: 3
	CategoryMasterId5 *int64 `json:"categoryMasterId5" msgpack:"categoryMasterId5"` // Key: 4
}

type LeagueAllClassGlobalRankingRewardPackageMaster struct {
	Id             int64                                          `json:"id" msgpack:"id"`                         // Key: 0
	LeagueMasterId int64                                          `json:"leagueMasterId" msgpack:"leagueMasterId"` // Key: 1
	Rewards        []LeagueAllClassGlobalRankingRewardThingMaster `json:"rewards" msgpack:"rewards"`               // Key: 2
	MinRank        int                                            `json:"minRank" msgpack:"minRank"`               // Key: 3
	MaxRank        int                                            `json:"maxRank" msgpack:"maxRank"`               // Key: 4
	RewardTitle    string                                         `json:"rewardTitle" msgpack:"rewardTitle"`       // Key: 5
}

type LeagueAllClassGlobalRankingRewardThingMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type LeagueClassGroupMaster struct {
	Id      int64               `json:"id" msgpack:"id"`           // Key: 0
	Classes []LeagueClassMaster `json:"classes" msgpack:"classes"` // Key: 1
}

type LeagueClassMaster struct {
	Id                                int64            `json:"id" msgpack:"id"`                                                               // Key: 0
	UpUserAmount                      *int             `json:"upUserAmount" msgpack:"upUserAmount"`                                           // Key: 1
	KeepUserAmount                    *int             `json:"keepUserAmount" msgpack:"keepUserAmount"`                                       // Key: 2
	IsForceDownNotPlayingLeague       bool             `json:"isForceDownNotPlayingLeague" msgpack:"isForceDownNotPlayingLeague"`             // Key: 3
	UpRewardPackageMasterId           *int64           `json:"upRewardPackageMasterId" msgpack:"upRewardPackageMasterId"`                     // Key: 4
	KeepRewardPackageMasterId         *int64           `json:"keepRewardPackageMasterId" msgpack:"keepRewardPackageMasterId"`                 // Key: 5
	DownRewardPackageMasterId         *int64           `json:"downRewardPackageMasterId" msgpack:"downRewardPackageMasterId"`                 // Key: 6
	FirstAchieveRewardPackageMasterId *int64           `json:"firstAchieveRewardPackageMasterId" msgpack:"firstAchieveRewardPackageMasterId"` // Key: 7
	NameplateMasterId                 *int64           `json:"nameplateMasterId" msgpack:"nameplateMasterId"`                                 // Key: 8
	ClassType                         LeagueClassTypes `json:"classType" msgpack:"classType"`                                                 // Key: 9
	LeagueClassGroupMasterId          int64            `json:"leagueClassGroupMasterId" msgpack:"leagueClassGroupMasterId"`                   // Key: 10
}

type LeagueGroupRankingRewardPackageMaster struct {
	Id          int64                                 `json:"id" msgpack:"id"`                   // Key: 0
	Rewards     []LeagueGroupRankingRewardThingMaster `json:"rewards" msgpack:"rewards"`         // Key: 1
	ClassType   LeagueClassTypes                      `json:"classType" msgpack:"classType"`     // Key: 2
	MinRank     int                                   `json:"minRank" msgpack:"minRank"`         // Key: 3
	MaxRank     int                                   `json:"maxRank" msgpack:"maxRank"`         // Key: 4
	RewardTitle string                                `json:"rewardTitle" msgpack:"rewardTitle"` // Key: 5
}

type LeagueGroupRankingRewardThingMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type LeagueMaster struct {
	Id                       int64             `json:"id" msgpack:"id"`                                             // Key: 0
	MusicMasterId            int64             `json:"musicMasterId" msgpack:"musicMasterId"`                       // Key: 1
	MaxDifficulty            MusicDifficulties `json:"maxDifficulty" msgpack:"maxDifficulty"`                       // Key: 2
	DisplayStartAt           time.Time         `json:"displayStartAt" msgpack:"displayStartAt"`                     // Key: 3
	CountingStartAt          time.Time         `json:"countingStartAt" msgpack:"countingStartAt"`                   // Key: 4
	DisplayEndAt             time.Time         `json:"displayEndAt" msgpack:"displayEndAt"`                         // Key: 5
	LeagueClassGroupMasterId int64             `json:"leagueClassGroupMasterId" msgpack:"leagueClassGroupMasterId"` // Key: 6
	SenseNotationMasterId    int64             `json:"senseNotationMasterId" msgpack:"senseNotationMasterId"`       // Key: 7
	VocalVersion             int               `json:"vocalVersion" msgpack:"vocalVersion"`                         // Key: 8
	LeagueSeasonMasterId     int64             `json:"leagueSeasonMasterId" msgpack:"leagueSeasonMasterId"`         // Key: 9
}

type LeaguePlayRewardPackageMaster struct {
	Id      int64                         `json:"id" msgpack:"id"`           // Key: 0
	Rewards []LeaguePlayRewardThingMaster `json:"rewards" msgpack:"rewards"` // Key: 1
}

type LeaguePlayRewardThingMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type LeagueRewardPackageMaster struct {
	Id      int64                     `json:"id" msgpack:"id"`           // Key: 0
	Rewards []LeagueRewardThingMaster `json:"rewards" msgpack:"rewards"` // Key: 1
}

type LeagueRewardThingMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type LessonScoreRewardGroupMaster struct {
	Id      int64                     `json:"id" msgpack:"id"`           // Key: 0
	Rewards []LessonScoreRewardMaster `json:"rewards" msgpack:"rewards"` // Key: 1
}

type LessonScoreRewardMaster struct {
	Id            int64      `json:"id" msgpack:"id"`                       // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 2
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 3
}

type LightLoadSplitEffectMaster struct {
	Id              int64 `json:"id" msgpack:"id"`                           // Key: 0
	BaseEffect      int   `json:"baseEffect" msgpack:"baseEffect"`           // Key: 1
	LightLoadEffect int   `json:"lightLoadEffect" msgpack:"lightLoadEffect"` // Key: 2
}

type LipSyncMaster struct {
	Id         int64  `json:"id" msgpack:"id"`                 // Key: 0
	MotionName string `json:"motionName" msgpack:"motionName"` // Key: 1
}

type LiveDropFrameGroupMaster struct {
	Id         int64                 `json:"id" msgpack:"id"`                 // Key: 0
	DropFrames []LiveDropFrameMaster `json:"dropFrames" msgpack:"dropFrames"` // Key: 1
}

type LiveDropFrameMaster struct {
	Id                         int64                  `json:"id" msgpack:"id"`                                                 // Key: 0
	LiveDropFrameGroupMasterId int64                  `json:"liveDropFrameGroupMasterId" msgpack:"liveDropFrameGroupMasterId"` // Key: 1
	Order                      int                    `json:"order" msgpack:"order"`                                           // Key: 2
	ChangeType                 ProbabilityChangeTypes `json:"changeType" msgpack:"changeType"`                                 // Key: 3
	FrameLotCondition          FrameLotConditionTypes `json:"frameLotCondition" msgpack:"frameLotCondition"`                   // Key: 4
	UnaffectedIncreaseEffect   bool                   `json:"unaffectedIncreaseEffect" msgpack:"unaffectedIncreaseEffect"`     // Key: 5
	Rewards                    []LiveDropThingMaster  `json:"rewards" msgpack:"rewards"`                                       // Key: 6
	FrameLotConditionValue     *int                   `json:"frameLotConditionValue" msgpack:"frameLotConditionValue"`         // Key: 7
	StartDate                  *time.Time             `json:"startDate" msgpack:"startDate"`                                   // Key: 8
	EndDate                    *time.Time             `json:"endDate" msgpack:"endDate"`                                       // Key: 9
}

type LiveDropThingMaster struct {
	ThingId                   int64      `json:"thingId" msgpack:"thingId"`                                     // Key: 0
	ThingType                 ThingTypes `json:"thingType" msgpack:"thingType"`                                 // Key: 1
	ThingQuantity             int        `json:"thingQuantity" msgpack:"thingQuantity"`                         // Key: 2
	IsSpecialFallCountUpThing bool       `json:"isSpecialFallCountUpThing" msgpack:"isSpecialFallCountUpThing"` // Key: 3
	IsSpecialFallThing        bool       `json:"isSpecialFallThing" msgpack:"isSpecialFallThing"`               // Key: 4
}

type LiveMaster struct {
	Id              int64                    `json:"id" msgpack:"id"`                           // Key: 0
	Difficulty      MusicDifficulties        `json:"difficulty" msgpack:"difficulty"`           // Key: 1
	MusicMasterId   int64                    `json:"musicMasterId" msgpack:"musicMasterId"`     // Key: 2
	Level           int                      `json:"level" msgpack:"level"`                     // Key: 3
	NoteCount       int                      `json:"noteCount" msgpack:"noteCount"`             // Key: 4
	UnlockCondition LiveUnlockConditionTypes `json:"unlockCondition" msgpack:"unlockCondition"` // Key: 5
	UnlockValue     *int64                   `json:"unlockValue" msgpack:"unlockValue"`         // Key: 6
	StartDate       time.Time                `json:"startDate" msgpack:"startDate"`             // Key: 7
	EndDate         time.Time                `json:"endDate" msgpack:"endDate"`                 // Key: 8
}

type LiveSettingMaster struct {
	Id                         int64     `json:"id" msgpack:"id"`                                                 // Key: 0
	Name                       string    `json:"name" msgpack:"name"`                                             // Key: 1
	LiveType                   LiveTypes `json:"liveType" msgpack:"liveType"`                                     // Key: 2
	LiveDropFrameGroupMasterId int64     `json:"liveDropFrameGroupMasterId" msgpack:"liveDropFrameGroupMasterId"` // Key: 3
}

type LoginBonusSpineCostumeMaster struct {
	Id                    int64     `json:"id" msgpack:"id"`                                       // Key: 0
	SpineId               int64     `json:"spineId" msgpack:"spineId"`                             // Key: 1
	CharacterBaseMasterId int64     `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 2
	TargetValue           int64     `json:"targetValue" msgpack:"targetValue"`                     // Key: 3
	ReleaseDate           time.Time `json:"releaseDate" msgpack:"releaseDate"`                     // Key: 4
}

type LoopMotionMaster struct {
	Id                    int64              `json:"id" msgpack:"id"`                                       // Key: 0
	TargetCharacterBaseId string             `json:"targetCharacterBaseId" msgpack:"targetCharacterBaseId"` // Key: 1
	LoopSpeed             float32            `json:"loopSpeed" msgpack:"loopSpeed"`                         // Key: 2
	Height                int                `json:"height" msgpack:"height"`                               // Key: 3
	Size                  SpineBodySizeTypes `json:"size" msgpack:"size"`                                   // Key: 4
}

type MarketFrameThingMaster struct {
	Id                    int64              `json:"id" msgpack:"id"`                                       // Key: 0
	ThingId               int64              `json:"thingId" msgpack:"thingId"`                             // Key: 1
	ThingType             ThingTypes         `json:"thingType" msgpack:"thingType"`                         // Key: 2
	ThingQuantity         int                `json:"thingQuantity" msgpack:"thingQuantity"`                 // Key: 3
	RequiredThingType     ThingTypes         `json:"requiredThingType" msgpack:"requiredThingType"`         // Key: 6
	RequiredThingId       int64              `json:"requiredThingId" msgpack:"requiredThingId"`             // Key: 7
	RequiredThingQuantity int                `json:"requiredThingQuantity" msgpack:"requiredThingQuantity"` // Key: 8
	DisplayType           MarketDisplayTypes `json:"displayType" msgpack:"displayType"`                     // Key: 9
}

type MissionMaster struct {
	Id                       int64                `json:"id" msgpack:"id"`                                             // Key: 0
	MissionCategory          MissionCategories    `json:"missionCategory" msgpack:"missionCategory"`                   // Key: 1
	MissionViewOrder         int                  `json:"missionViewOrder" msgpack:"missionViewOrder"`                 // Key: 2
	Title                    string               `json:"title" msgpack:"title"`                                       // Key: 3
	Description              string               `json:"description" msgpack:"description"`                           // Key: 4
	StartDate                time.Time            `json:"startDate" msgpack:"startDate"`                               // Key: 14
	EndDate                  time.Time            `json:"endDate" msgpack:"endDate"`                                   // Key: 15
	EventMasterId            *int64               `json:"eventMasterId" msgpack:"eventMasterId"`                       // Key: 11
	JumpType                 JumpTypes            `json:"jumpType" msgpack:"jumpType"`                                 // Key: 12
	JumpTargetId             *int64               `json:"jumpTargetId" msgpack:"jumpTargetId"`                         // Key: 13
	Stages                   []MissionStageMaster `json:"stages" msgpack:"stages"`                                     // Key: 16
	ComebackCampaignMasterId *int64               `json:"comebackCampaignMasterId" msgpack:"comebackCampaignMasterId"` // Key: 17
}

type MissionPassDetailMaster struct {
	Id                  int64                    `json:"id" msgpack:"id"`                                   // Key: 0
	Phase               int                      `json:"phase" msgpack:"phase"`                             // Key: 1
	MissionPassMasterId *int64                   `json:"missionPassMasterId" msgpack:"missionPassMasterId"` // Key: 2
	ClearPoint          int                      `json:"clearPoint" msgpack:"clearPoint"`                   // Key: 3
	StartDate           time.Time                `json:"startDate" msgpack:"startDate"`                     // Key: 4
	EndDate             time.Time                `json:"endDate" msgpack:"endDate"`                         // Key: 5
	Rewards             []MissionPassRewardThing `json:"rewards" msgpack:"rewards"`                         // Key: 6
}

type MissionPassLoopRewardMaster struct {
	Id        int64                              `json:"id" msgpack:"id"`               // Key: 0
	Phase     int                                `json:"phase" msgpack:"phase"`         // Key: 1
	StartDate time.Time                          `json:"startDate" msgpack:"startDate"` // Key: 2
	EndDate   time.Time                          `json:"endDate" msgpack:"endDate"`     // Key: 3
	Rewards   []MissionPassLoopRewardThingMaster `json:"rewards" msgpack:"rewards"`     // Key: 4
}

type MissionPassLoopRewardThingMaster struct {
	Id                            int64      `json:"id" msgpack:"id"`                                                       // Key: 0
	MissionPassLoopRewardMasterId int64      `json:"missionPassLoopRewardMasterId" msgpack:"missionPassLoopRewardMasterId"` // Key: 1
	ThingId                       int64      `json:"thingId" msgpack:"thingId"`                                             // Key: 2
	ThingType                     ThingTypes `json:"thingType" msgpack:"thingType"`                                         // Key: 3
	ThingQuantity                 int        `json:"thingQuantity" msgpack:"thingQuantity"`                                 // Key: 4
	IsSp                          bool       `json:"isSp" msgpack:"isSp"`                                                   // Key: 5
}

type MissionPassMaster struct {
	Id           int64     `json:"id" msgpack:"id"`                     // Key: 0
	ItemMasterId int64     `json:"itemMasterId" msgpack:"itemMasterId"` // Key: 1
	StartDate    time.Time `json:"startDate" msgpack:"startDate"`       // Key: 2
	EndDate      time.Time `json:"endDate" msgpack:"endDate"`           // Key: 3
}

type MissionPassRewardThing struct {
	Id                        int64      `json:"id" msgpack:"id"`                                               // Key: 0
	ThingId                   int64      `json:"thingId" msgpack:"thingId"`                                     // Key: 1
	ThingType                 ThingTypes `json:"thingType" msgpack:"thingType"`                                 // Key: 2
	ThingQuantity             int        `json:"thingQuantity" msgpack:"thingQuantity"`                         // Key: 3
	MissionPassDetailMasterId int64      `json:"missionPassDetailMasterId" msgpack:"missionPassDetailMasterId"` // Key: 4
	IsSp                      bool       `json:"isSp" msgpack:"isSp"`                                           // Key: 5
}

type MissionRewardMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type MissionStageMaster struct {
	Id                int64                 `json:"id" msgpack:"id"`                               // Key: 0
	MissionStageOrder int                   `json:"missionStageOrder" msgpack:"missionStageOrder"` // Key: 1
	StageGoalValue    int64                 `json:"stageGoalValue" msgpack:"stageGoalValue"`       // Key: 2
	StartDate         *time.Time            `json:"startDate" msgpack:"startDate"`                 // Key: 3
	Rewards           []MissionRewardMaster `json:"rewards" msgpack:"rewards"`                     // Key: 4
}

type MultiLiveScheduleMaster struct {
	Id                   int64                          `json:"id" msgpack:"id"`                                     // Key: 0
	StartDate            time.Time                      `json:"startDate" msgpack:"startDate"`                       // Key: 1
	EndDate              time.Time                      `json:"endDate" msgpack:"endDate"`                           // Key: 2
	LiveSettingMasterId1 int64                          `json:"liveSettingMasterId1" msgpack:"liveSettingMasterId1"` // Key: 3
	LiveSettingMasterId2 int64                          `json:"liveSettingMasterId2" msgpack:"liveSettingMasterId2"` // Key: 4
	AdditionalEndDate    *time.Time                     `json:"additionalEndDate" msgpack:"additionalEndDate"`       // Key: 5
	EndDateDisplayType1  LiveScheduleEndDateDisplayType `json:"endDateDisplayType1" msgpack:"endDateDisplayType1"`   // Key: 6
	EndDateDisplayType2  LiveScheduleEndDateDisplayType `json:"endDateDisplayType2" msgpack:"endDateDisplayType2"`   // Key: 7
}

type MusicCourseMaster struct {
	Id                           int64                           `json:"id" msgpack:"id"`                                                     // Key: 0
	Name                         string                          `json:"name" msgpack:"name"`                                                 // Key: 1
	MusicCourseType              MusicCourseType                 `json:"musicCourseType" msgpack:"musicCourseType"`                           // Key: 2
	Order                        int64                           `json:"order" msgpack:"order"`                                               // Key: 3
	InitialLife                  int                             `json:"initialLife" msgpack:"initialLife"`                                   // Key: 4
	RequiredItemMasterId         *int64                          `json:"requiredItemMasterId" msgpack:"requiredItemMasterId"`                 // Key: 5
	RequiredAmount               int                             `json:"requiredAmount" msgpack:"requiredAmount"`                             // Key: 6
	StartDate                    time.Time                       `json:"startDate" msgpack:"startDate"`                                       // Key: 7
	EndDate                      *time.Time                      `json:"endDate" msgpack:"endDate"`                                           // Key: 8
	UnlockCondition              MusicCourseUnlockConditionTypes `json:"unlockCondition" msgpack:"unlockCondition"`                           // Key: 9
	UnlockConditionValue         *int64                          `json:"unlockConditionValue" msgpack:"unlockConditionValue"`                 // Key: 10
	TournamentQualifyingMasterId *int64                          `json:"tournamentQualifyingMasterId" msgpack:"tournamentQualifyingMasterId"` // Key: 11
	Details                      []MusicCourseDetailMaster       `json:"details" msgpack:"details"`                                           // Key: 12
}

type MusicCourseDetailMaster struct {
	Id                      int64              `json:"id" msgpack:"id"`                                           // Key: 0
	SetListNumber           int                `json:"setListNumber" msgpack:"setListNumber"`                     // Key: 1
	LiveMasterId            *int64             `json:"liveMasterId" msgpack:"liveMasterId"`                       // Key: 2
	VocalVersion            int                `json:"vocalVersion" msgpack:"vocalVersion"`                       // Key: 3
	AnotherNotationMasterId int64              `json:"anotherNotationMasterId" msgpack:"anotherNotationMasterId"` // Key: 4
	RandomDifficulty        *MusicDifficulties `json:"randomDifficulty" msgpack:"randomDifficulty"`               // Key: 5
	RandomLevel             *int               `json:"randomLevel" msgpack:"randomLevel"`                         // Key: 6
}

type MusicCourseRewardGroupMaster struct {
	Id                         int64                         `json:"id" msgpack:"id"`                                                 // Key: 0
	MusicCourseMasterId        int64                         `json:"musicCourseMasterId" msgpack:"musicCourseMasterId"`               // Key: 1
	RequiredCertificationGrade MusicCourseCertificationGrade `json:"requiredCertificationGrade" msgpack:"requiredCertificationGrade"` // Key: 2
	Rewards                    []MusicCourseRewardThing      `json:"rewards" msgpack:"rewards"`                                       // Key: 3
}

type MusicCourseRewardThing struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type MusicGroupDetailMaster struct {
	Id                 int64 `json:"id" msgpack:"id"`                                 // Key: 0
	MusicGroupMasterId int64 `json:"musicGroupMasterId" msgpack:"musicGroupMasterId"` // Key: 1
	MusicMasterId      int64 `json:"musicMasterId" msgpack:"musicMasterId"`           // Key: 2
}

type MusicGroupMaster struct {
	Id                int64 `json:"id" msgpack:"id"`                               // Key: 0
	MainMusicMasterId int64 `json:"mainMusicMasterId" msgpack:"mainMusicMasterId"` // Key: 1
}

type MusicMaster struct {
	Id                   int64                     `json:"id" msgpack:"id"`                                     // Key: 0
	Name                 string                    `json:"name" msgpack:"name"`                                 // Key: 1
	Description          string                    `json:"description" msgpack:"description"`                   // Key: 2
	RewardRuleMasterId   int64                     `json:"rewardRuleMasterId" msgpack:"rewardRuleMasterId"`     // Key: 4
	PronounceName        string                    `json:"pronounceName" msgpack:"pronounceName"`               // Key: 5
	LyricWriter          string                    `json:"lyricWriter" msgpack:"lyricWriter"`                   // Key: 6
	Composer             string                    `json:"composer" msgpack:"composer"`                         // Key: 7
	Arranger             string                    `json:"arranger" msgpack:"arranger"`                         // Key: 8
	UnlockText           string                    `json:"unlockText" msgpack:"unlockText"`                     // Key: 10
	IsLongVersion        bool                      `json:"isLongVersion" msgpack:"isLongVersion"`               // Key: 11
	ReleasedAt           *time.Time                `json:"releasedAt" msgpack:"releasedAt"`                     // Key: 12
	StaminaConsumption   int                       `json:"staminaConsumption" msgpack:"staminaConsumption"`     // Key: 13
	MusicTimeSecond      int                       `json:"musicTimeSecond" msgpack:"musicTimeSecond"`           // Key: 14
	Invisible            bool                      `json:"invisible" msgpack:"invisible"`                       // Key: 15
	SampleStartSeconds   float32                   `json:"sampleStartSeconds" msgpack:"sampleStartSeconds"`     // Key: 17
	SampleEndSeconds     float32                   `json:"sampleEndSeconds" msgpack:"sampleEndSeconds"`         // Key: 18
	DelaySeconds         float32                   `json:"delaySeconds" msgpack:"delaySeconds"`                 // Key: 19
	VocalVersions        []MusicVocalVersionMaster `json:"vocalVersions" msgpack:"vocalVersions"`               // Key: 20
	UnlockConditionType  MusicUnlockConditionTypes `json:"unlockConditionType" msgpack:"unlockConditionType"`   // Key: 21
	UnlockConditionValue *int64                    `json:"unlockConditionValue" msgpack:"unlockConditionValue"` // Key: 22
	MusicVideoType       MusicVideoTypes           `json:"musicVideoType" msgpack:"musicVideoType"`             // Key: 23
	MusicCoverType       MusicCoverTypes           `json:"musicCoverType" msgpack:"musicCoverType"`             // Key: 24
	StoryEventMasterId   *int64                    `json:"storyEventMasterId" msgpack:"storyEventMasterId"`     // Key: 25
	StoryMasterId        *int64                    `json:"storyMasterId" msgpack:"storyMasterId"`               // Key: 26
	EventMasterId        *int64                    `json:"eventMasterId" msgpack:"eventMasterId"`               // Key: 27
}

type MusicVideoDefaultCostumeGroupMaster struct {
	Id              int64                            `json:"id" msgpack:"id"`                           // Key: 0
	DefaultCostumes []MusicVideoDefaultCostumeMaster `json:"defaultCostumes" msgpack:"defaultCostumes"` // Key: 1
}

type MusicVideoDefaultCostumeMaster struct {
	Id                    int64 `json:"id" msgpack:"id"`                                       // Key: 0
	CharacterBaseMasterId int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	CostumeMasterId       int64 `json:"costumeMasterId" msgpack:"costumeMasterId"`             // Key: 2
}

type MusicVideoMaster struct {
	Id                                    int64                               `json:"id" msgpack:"id"`                                                                       // Key: 0
	MusicMasterId                         int64                               `json:"musicMasterId" msgpack:"musicMasterId"`                                                 // Key: 1
	MemberAmount                          int                                 `json:"memberAmount" msgpack:"memberAmount"`                                                   // Key: 2
	DisplayStartAt                        time.Time                           `json:"displayStartAt" msgpack:"displayStartAt"`                                               // Key: 3
	DisplayEndAt                          *time.Time                          `json:"displayEndAt" msgpack:"displayEndAt"`                                                   // Key: 4
	Characters                            []MusicVideoOriginalCharacterMaster `json:"characters" msgpack:"characters"`                                                       // Key: 5
	DelaySeconds                          float32                             `json:"delaySeconds" msgpack:"delaySeconds"`                                                   // Key: 6
	MovieDelaySeconds                     float32                             `json:"movieDelaySeconds" msgpack:"movieDelaySeconds"`                                         // Key: 7
	MusicVideoDefaultCostumeGroupMasterId *int64                              `json:"musicVideoDefaultCostumeGroupMasterId" msgpack:"musicVideoDefaultCostumeGroupMasterId"` // Key: 8
	IsFixedMember                         bool                                `json:"isFixedMember" msgpack:"isFixedMember"`                                                 // Key: 9
	Order                                 int                                 `json:"order" msgpack:"order"`                                                                 // Key: 10
}

type MusicVideoOriginalCharacterMaster struct {
	Order                 int   `json:"order" msgpack:"order"`                                 // Key: 0
	CharacterBaseMasterId int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	CostumeMasterId       int64 `json:"costumeMasterId" msgpack:"costumeMasterId"`             // Key: 2
}

type MusicVocalVersionMaster struct {
	Id                 int64           `json:"id" msgpack:"id"`                                 // Key: 0
	MusicMasterId      int64           `json:"musicMasterId" msgpack:"musicMasterId"`           // Key: 1
	VocalVersion       int             `json:"vocalVersion" msgpack:"vocalVersion"`             // Key: 2
	Singer             string          `json:"singer" msgpack:"singer"`                         // Key: 3
	Name               string          `json:"name" msgpack:"name"`                             // Key: 4
	MusicTimeSecond    int             `json:"musicTimeSecond" msgpack:"musicTimeSecond"`       // Key: 5
	SampleStartSeconds float32         `json:"sampleStartSeconds" msgpack:"sampleStartSeconds"` // Key: 6
	SampleEndSeconds   float32         `json:"sampleEndSeconds" msgpack:"sampleEndSeconds"`     // Key: 7
	MusicVideoType     MusicVideoTypes `json:"musicVideoType" msgpack:"musicVideoType"`         // Key: 8
	Characters         []int64         `json:"characters" msgpack:"characters"`                 // Key: 9
}

type NameBaseColorMaster struct {
	Id                  int64  `json:"id" msgpack:"id"`                                   // Key: 0
	Name                string `json:"name" msgpack:"name"`                               // Key: 1
	Order               int    `json:"order" msgpack:"order"`                             // Key: 2
	UnlockConditionText string `json:"unlockConditionText" msgpack:"unlockConditionText"` // Key: 3
	Hidden              bool   `json:"hidden" msgpack:"hidden"`                           // Key: 4
	IsDefault           bool   `json:"isDefault" msgpack:"isDefault"`                     // Key: 5
}

type NameColorMaster struct {
	Id          int64  `json:"id" msgpack:"id"`                   // Key: 0
	Name        string `json:"name" msgpack:"name"`               // Key: 1
	Description string `json:"description" msgpack:"description"` // Key: 2
	Order       int    `json:"order" msgpack:"order"`             // Key: 3
	Hidden      bool   `json:"hidden" msgpack:"hidden"`           // Key: 4
	IsDefault   bool   `json:"isDefault" msgpack:"isDefault"`     // Key: 5
	UnlockText  string `json:"unlockText" msgpack:"unlockText"`   // Key: 6
}

type NameplateDetailMaster struct {
	Id           int64            `json:"id" msgpack:"id"`                     // Key: 0
	Name         string           `json:"name" msgpack:"name"`                 // Key: 1
	ClassType    LeagueClassTypes `json:"classType" msgpack:"classType"`       // Key: 2
	EnrollCount  int              `json:"enrollCount" msgpack:"enrollCount"`   // Key: 3
	ChangeValue1 *int64           `json:"changeValue1" msgpack:"changeValue1"` // Key: 4
	ChangeValue2 *int64           `json:"changeValue2" msgpack:"changeValue2"` // Key: 5
	Priority     int              `json:"priority" msgpack:"priority"`         // Key: 6
}

type NameplateMaster struct {
	Id           int64                   `json:"id" msgpack:"id"`                     // Key: 0
	Name         string                  `json:"name" msgpack:"name"`                 // Key: 1
	Description  string                  `json:"description" msgpack:"description"`   // Key: 2
	Order        int                     `json:"order" msgpack:"order"`               // Key: 3
	Hidden       bool                    `json:"hidden" msgpack:"hidden"`             // Key: 4
	IsDefault    bool                    `json:"isDefault" msgpack:"isDefault"`       // Key: 5
	Details      []NameplateDetailMaster `json:"details" msgpack:"details"`           // Key: 6
	UnlockText   string                  `json:"unlockText" msgpack:"unlockText"`     // Key: 7
	ChangeType   NamePlateChangeTypes    `json:"changeType" msgpack:"changeType"`     // Key: 8
	ChangeValue1 *int64                  `json:"changeValue1" msgpack:"changeValue1"` // Key: 9
	ChangeValue2 *int64                  `json:"changeValue2" msgpack:"changeValue2"` // Key: 10
}

type NgWordMaster struct {
	Id      int64  `json:"id" msgpack:"id"`           // Key: 0
	KeyWord string `json:"keyWord" msgpack:"keyWord"` // Key: 1
}

type NoteMaster struct {
	Id          int64  `json:"id" msgpack:"id"`                   // Key: 0
	Name        string `json:"name" msgpack:"name"`               // Key: 1
	Description string `json:"description" msgpack:"description"` // Key: 2
	Order       int    `json:"order" msgpack:"order"`             // Key: 3
	Hidden      bool   `json:"hidden" msgpack:"hidden"`           // Key: 4
	IsDefault   bool   `json:"isDefault" msgpack:"isDefault"`     // Key: 5
}

type PermanentMarketThingMaster struct {
	Id                    int64            `json:"id" msgpack:"id"`                                       // Key: 0
	ThingId               int64            `json:"thingId" msgpack:"thingId"`                             // Key: 1
	ThingType             ThingTypes       `json:"thingType" msgpack:"thingType"`                         // Key: 2
	ThingQuantity         int              `json:"thingQuantity" msgpack:"thingQuantity"`                 // Key: 3
	ExchangeLimit         *int             `json:"exchangeLimit" msgpack:"exchangeLimit"`                 // Key: 4
	Order                 int              `json:"order" msgpack:"order"`                                 // Key: 5
	UnlockType            *ShopUnlockTypes `json:"unlockType" msgpack:"unlockType"`                       // Key: 6
	UnlockValue           *int64           `json:"unlockValue" msgpack:"unlockValue"`                     // Key: 7
	StartDate             *time.Time       `json:"startDate" msgpack:"startDate"`                         // Key: 8
	EndDate               *time.Time       `json:"endDate" msgpack:"endDate"`                             // Key: 9
	RequiredThingType     ThingTypes       `json:"requiredThingType" msgpack:"requiredThingType"`         // Key: 11
	RequiredThingId       int64            `json:"requiredThingId" msgpack:"requiredThingId"`             // Key: 12
	RequiredThingQuantity int              `json:"requiredThingQuantity" msgpack:"requiredThingQuantity"` // Key: 13
}

type PhotoEffectChangeItemMaster struct {
	Id                           int64 `json:"id" msgpack:"id"`                                                     // Key: 0
	ItemMasterId                 int64 `json:"itemMasterId" msgpack:"itemMasterId"`                                 // Key: 1
	PhotoEffectTypeGroupMasterId int64 `json:"photoEffectTypeGroupMasterId" msgpack:"photoEffectTypeGroupMasterId"` // Key: 2
	RequiredQuantity             int   `json:"requiredQuantity" msgpack:"requiredQuantity"`                         // Key: 3
}

type PhotoEffectMaster struct {
	Id                           int64  `json:"id" msgpack:"id"`                                                     // Key: 0
	Name                         string `json:"name" msgpack:"name"`                                                 // Key: 1
	Description                  string `json:"description" msgpack:"description"`                                   // Key: 2
	EffectMasterId               int64  `json:"effectMasterId" msgpack:"effectMasterId"`                             // Key: 3
	Variety                      *int   `json:"variety" msgpack:"variety"`                                           // Key: 4
	PhotoEffectTypeGroupMasterId *int64 `json:"photoEffectTypeGroupMasterId" msgpack:"photoEffectTypeGroupMasterId"` // Key: 5
	PhotoEffectGroupMasterId     int64  `json:"photoEffectGroupMasterId" msgpack:"photoEffectGroupMasterId"`         // Key: 6
}

type PhotoEffectTypeGroupMaster struct {
	Id                    int64  `json:"id" msgpack:"id"`                                       // Key: 0
	PhotoEffectMaxVariety int    `json:"photoEffectMaxVariety" msgpack:"photoEffectMaxVariety"` // Key: 1
	Description           string `json:"description" msgpack:"description"`                     // Key: 2
	CanUseItem            bool   `json:"canUseItem" msgpack:"canUseItem"`                       // Key: 3
}

type PhotoEffectVarietyChangeDetailMaster struct {
	Id                          int64 `json:"id" msgpack:"id"`                                                   // Key: 0
	BeforePhotoEffectMaxVariety int   `json:"beforePhotoEffectMaxVariety" msgpack:"beforePhotoEffectMaxVariety"` // Key: 1
	AfterPhotoEffectMaxVariety  int   `json:"afterPhotoEffectMaxVariety" msgpack:"afterPhotoEffectMaxVariety"`   // Key: 2
	BeforeVariety               int   `json:"beforeVariety" msgpack:"beforeVariety"`                             // Key: 3
	AfterVariety                int   `json:"afterVariety" msgpack:"afterVariety"`                               // Key: 4
}

type PhotoEffectVarietyUpDetailMaster struct {
	Id                           int64 `json:"id" msgpack:"id"`                                                     // Key: 0
	PhotoEffectTypeGroupMasterId int64 `json:"photoEffectTypeGroupMasterId" msgpack:"photoEffectTypeGroupMasterId"` // Key: 1
	CurrentVariety               int   `json:"currentVariety" msgpack:"currentVariety"`                             // Key: 2
	RequiredQuantity             int   `json:"requiredQuantity" msgpack:"requiredQuantity"`                         // Key: 3
}

type PhotoLevelUpItemGroupMaster struct {
	Id           int64                    `json:"id" msgpack:"id"`                     // Key: 0
	Rarity       PhotoRarities            `json:"rarity" msgpack:"rarity"`             // Key: 1
	BaseLevel    int                      `json:"baseLevel" msgpack:"baseLevel"`       // Key: 2
	RequiredCoin int                      `json:"requiredCoin" msgpack:"requiredCoin"` // Key: 3
	Items        []PhotoLevelUpItemMaster `json:"items" msgpack:"items"`               // Key: 4
}

type PhotoLevelUpItemMaster struct {
	ItemMasterId     int64 `json:"itemMasterId" msgpack:"itemMasterId"`         // Key: 0
	RequiredQuantity int   `json:"requiredQuantity" msgpack:"requiredQuantity"` // Key: 1
}

type PhotoSpotMaster struct {
	Id                int64     `json:"id" msgpack:"id"`                               // Key: 0
	Name              string    `json:"name" msgpack:"name"`                           // Key: 1
	ReleasePlayerRank int       `json:"releasePlayerRank" msgpack:"releasePlayerRank"` // Key: 2
	DisplayStartAt    time.Time `json:"displayStartAt" msgpack:"displayStartAt"`       // Key: 3
	DisplayEndAt      time.Time `json:"displayEndAt" msgpack:"displayEndAt"`           // Key: 4
}

type PickupCharacterMissionDetailGroupMaster struct {
	Id      int64                                `json:"id" msgpack:"id"`           // Key: 0
	Details []PickupCharacterMissionDetailMaster `json:"details" msgpack:"details"` // Key: 1
}

type PickupCharacterMissionDetailMaster struct {
	Id             int64                                      `json:"id" msgpack:"id"`                         // Key: 0
	Order          int                                        `json:"order" msgpack:"order"`                   // Key: 1
	CheckCondition PickupCharacterMissionCheckCondition       `json:"checkCondition" msgpack:"checkCondition"` // Key: 2
	Description    string                                     `json:"description" msgpack:"description"`       // Key: 3
	GoalValue      int                                        `json:"goalValue" msgpack:"goalValue"`           // Key: 4
	Rewards        []PickupCharacterMissionDetailRewardMaster `json:"rewards" msgpack:"rewards"`               // Key: 5
}

type PickupCharacterMissionDetailRewardMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type PickupCharacterMissionMaster struct {
	Id                                        int64     `json:"id" msgpack:"id"`                                                                               // Key: 0
	CharacterMasterId                         int64     `json:"characterMasterId" msgpack:"characterMasterId"`                                                 // Key: 1
	Order                                     int       `json:"order" msgpack:"order"`                                                                         // Key: 2
	PickupCharacterMissionDetailGroupMasterId int64     `json:"pickupCharacterMissionDetailGroupMasterId" msgpack:"pickupCharacterMissionDetailGroupMasterId"` // Key: 3
	StartDate                                 time.Time `json:"startDate" msgpack:"startDate"`                                                                 // Key: 4
	EndDate                                   time.Time `json:"endDate" msgpack:"endDate"`                                                                     // Key: 5
}

type PickupSelectionGachaMaster struct {
	Id                   int64 `json:"id" msgpack:"id"`                                     // Key: 0
	GachaMasterId        int64 `json:"gachaMasterId" msgpack:"gachaMasterId"`               // Key: 1
	MinSelectCount       int   `json:"minSelectCount" msgpack:"minSelectCount"`             // Key: 2
	MaxSelectCount       int   `json:"maxSelectCount" msgpack:"maxSelectCount"`             // Key: 3
	FixedPrizePickupOnly bool  `json:"fixedPrizePickupOnly" msgpack:"fixedPrizePickupOnly"` // Key: 4
}

type PlayerRankCapMaster struct {
	Id      int64                       `json:"id" msgpack:"id"`           // Key: 0
	Coin    int                         `json:"coin" msgpack:"coin"`       // Key: 1
	Details []PlayerRankCapDetailMaster `json:"details" msgpack:"details"` // Key: 2
}

type PlayerRankCapDetailMaster struct {
	ItemMasterId int64 `json:"itemMasterId" msgpack:"itemMasterId"` // Key: 0
	Quantity     int   `json:"quantity" msgpack:"quantity"`         // Key: 1
}

type PlayerRankMaster struct {
	Rank                  int    `json:"rank" msgpack:"rank"`                                   // Key: 0
	PointToLevelUp        int    `json:"pointToLevelUp" msgpack:"pointToLevelUp"`               // Key: 1
	MaxStamina            int    `json:"maxStamina" msgpack:"maxStamina"`                       // Key: 2
	IsReleasedRank        bool   `json:"isReleasedRank" msgpack:"isReleasedRank"`               // Key: 4
	PlayerRankCapMasterId *int64 `json:"playerRankCapMasterId" msgpack:"playerRankCapMasterId"` // Key: 5
}

type PosterAbilityMaster struct {
	Id                   int64               `json:"id" msgpack:"id"`                                     // Key: 0
	Name                 string              `json:"name" msgpack:"name"`                                 // Key: 1
	Description          string              `json:"description" msgpack:"description"`                   // Key: 2
	PosterMasterId       int64               `json:"posterMasterId" msgpack:"posterMasterId"`             // Key: 3
	Type                 PosterEffectTypes   `json:"type" msgpack:"type"`                                 // Key: 4
	FrameNumber          int                 `json:"frameNumber" msgpack:"frameNumber"`                   // Key: 6
	ReleaseLevelAt       int                 `json:"releaseLevelAt" msgpack:"releaseLevelAt"`             // Key: 7
	Hidden               bool                `json:"hidden" msgpack:"hidden"`                             // Key: 8
	Branches             []BranchMaster      `json:"branches" msgpack:"branches"`                         // Key: 9
	BranchConditionType1 BranchConditionType `json:"branchConditionType1" msgpack:"branchConditionType1"` // Key: 10
	ConditionValue1      *int64              `json:"conditionValue1" msgpack:"conditionValue1"`           // Key: 11
	BranchConditionType2 BranchConditionType `json:"branchConditionType2" msgpack:"branchConditionType2"` // Key: 12
	ConditionValue2      *int64              `json:"conditionValue2" msgpack:"conditionValue2"`           // Key: 13
}

type PosterCharacterMaster struct {
	EpisodeType           PosterEpisodeTypes `json:"episodeType" msgpack:"episodeType"`                     // Key: 2
	CharacterBaseMasterId *int64             `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 3
	Description           string             `json:"description" msgpack:"description"`                     // Key: 4
}

type PosterCostumeMaster struct {
	Phase              int    `json:"phase" msgpack:"phase"`                           // Key: 0
	CostumeMasterId    *int64 `json:"costumeMasterId" msgpack:"costumeMasterId"`       // Key: 1
	ItemMasterId       *int64 `json:"itemMasterId" msgpack:"itemMasterId"`             // Key: 2
	Quantity           int    `json:"quantity" msgpack:"quantity"`                     // Key: 3
	DecorationMasterId *int64 `json:"decorationMasterId" msgpack:"decorationMasterId"` // Key: 4
}

type PosterLevelPatternGroupMaster struct {
	Id       int64                      `json:"id" msgpack:"id"`             // Key: 0
	Patterns []PosterLevelPatternMaster `json:"patterns" msgpack:"patterns"` // Key: 1
}

type PosterLevelPatternMaster struct {
	Id                  int64 `json:"id" msgpack:"id"`                                   // Key: 0
	LevelPatternGroupId int64 `json:"levelPatternGroupId" msgpack:"levelPatternGroupId"` // Key: 1
	Level               int   `json:"level" msgpack:"level"`                             // Key: 2
	ItemMasterId        int64 `json:"itemMasterId" msgpack:"itemMasterId"`               // Key: 3
	Quantity            int   `json:"quantity" msgpack:"quantity"`                       // Key: 4
}

type PosterMaster struct {
	Id                                     int64                           `json:"id" msgpack:"id"`                                                                         // Key: 0
	Name                                   string                          `json:"name" msgpack:"name"`                                                                     // Key: 1
	OrganizeRestrictGroupId                *int                            `json:"organizeRestrictGroupId" msgpack:"organizeRestrictGroupId"`                               // Key: 2
	Rarity                                 PossessionRarities              `json:"rarity" msgpack:"rarity"`                                                                 // Key: 3
	LevelPatternGroupMasterId              int64                           `json:"levelPatternGroupMasterId" msgpack:"levelPatternGroupMasterId"`                           // Key: 5
	SubTitlePositionX1                     *float32                        `json:"subTitlePositionX1" msgpack:"subTitlePositionX1"`                                         // Key: 6
	SubTitlePositionY1                     *float32                        `json:"subTitlePositionY1" msgpack:"subTitlePositionY1"`                                         // Key: 7
	SubTitlePositionX2                     *float32                        `json:"subTitlePositionX2" msgpack:"subTitlePositionX2"`                                         // Key: 8
	SubTitlePositionY2                     *float32                        `json:"subTitlePositionY2" msgpack:"subTitlePositionY2"`                                         // Key: 9
	SubTitlePositionX3                     *float32                        `json:"subTitlePositionX3" msgpack:"subTitlePositionX3"`                                         // Key: 10
	SubTitlePositionY3                     *float32                        `json:"subTitlePositionY3" msgpack:"subTitlePositionY3"`                                         // Key: 11
	ReleaseItemGroupId                     int64                           `json:"releaseItemGroupId" msgpack:"releaseItemGroupId"`                                         // Key: 12
	PronounceName                          string                          `json:"pronounceName" msgpack:"pronounceName"`                                                   // Key: 15
	Costumes                               []PosterCostumeMaster           `json:"costumes" msgpack:"costumes"`                                                             // Key: 16
	AppearanceCharacterBaseMasterIds       []int64                         `json:"appearanceCharacterBaseMasterIds" msgpack:"appearanceCharacterBaseMasterIds"`             // Key: 17
	IsRestrictItemBreakThrough             bool                            `json:"isRestrictItemBreakThrough" msgpack:"isRestrictItemBreakThrough"`                         // Key: 18
	DisplayStartAt                         time.Time                       `json:"displayStartAt" msgpack:"displayStartAt"`                                                 // Key: 19
	DisplayEndAt                           time.Time                       `json:"displayEndAt" msgpack:"displayEndAt"`                                                     // Key: 20
	UnlockText                             string                          `json:"unlockText" msgpack:"unlockText"`                                                         // Key: 21
	Orientation                            PosterOrientation               `json:"orientation" msgpack:"orientation"`                                                       // Key: 22
	SubTitleDisplayCondition               PosterSubTitleDisplayConditions `json:"subTitleDisplayCondition" msgpack:"subTitleDisplayCondition"`                             // Key: 23
	SubTitleDisplayConditionValue          *int                            `json:"subTitleDisplayConditionValue" msgpack:"subTitleDisplayConditionValue"`                   // Key: 24
	PosterBreakthroughMaxPhase             *int                            `json:"posterBreakthroughMaxPhase" msgpack:"posterBreakthroughMaxPhase"`                         // Key: 25
	PosterBreakthroughMaxPhaseReleaseDate  *time.Time                      `json:"posterBreakthroughMaxPhaseReleaseDate" msgpack:"posterBreakthroughMaxPhaseReleaseDate"`   // Key: 26
	SecondarySubTitleDisplayCondition      PosterSubTitleDisplayConditions `json:"secondarySubTitleDisplayCondition" msgpack:"secondarySubTitleDisplayCondition"`           // Key: 27
	SecondarySubTitleDisplayConditionValue *int                            `json:"secondarySubTitleDisplayConditionValue" msgpack:"secondarySubTitleDisplayConditionValue"` // Key: 28
	AlternateImagePositionX1               *float32                        `json:"alternateImagePositionX1" msgpack:"alternateImagePositionX1"`                             // Key: 29
	AlternateImagePositionY1               *float32                        `json:"alternateImagePositionY1" msgpack:"alternateImagePositionY1"`                             // Key: 30
	AlternateImageReleasePhase1            *int                            `json:"alternateImageReleasePhase1" msgpack:"alternateImageReleasePhase1"`                       // Key: 31
	AlternateImagePositionX2               *float32                        `json:"alternateImagePositionX2" msgpack:"alternateImagePositionX2"`                             // Key: 32
	AlternateImagePositionY2               *float32                        `json:"alternateImagePositionY2" msgpack:"alternateImagePositionY2"`                             // Key: 33
	AlternateImageReleasePhase2            *int                            `json:"alternateImageReleasePhase2" msgpack:"alternateImageReleasePhase2"`                       // Key: 34
	AlternateImagePositionX3               *float32                        `json:"alternateImagePositionX3" msgpack:"alternateImagePositionX3"`                             // Key: 35
	AlternateImagePositionY3               *float32                        `json:"alternateImagePositionY3" msgpack:"alternateImagePositionY3"`                             // Key: 36
	AlternateImageReleasePhase3            *int                            `json:"alternateImageReleasePhase3" msgpack:"alternateImageReleasePhase3"`                       // Key: 37
}

type PosterReleaseItemGroupMaster struct {
	Id                   int64                     `json:"id" msgpack:"id"`                                     // Key: 0
	Items                []PosterReleaseItemMaster `json:"items" msgpack:"items"`                               // Key: 1
	ItemConsumeApplyFlag bool                      `json:"itemConsumeApplyFlag" msgpack:"itemConsumeApplyFlag"` // Key: 2
}

type PosterReleaseItemMaster struct {
	CurrentPhase     int   `json:"currentPhase" msgpack:"currentPhase"`         // Key: 0
	ItemMasterId     int64 `json:"itemMasterId" msgpack:"itemMasterId"`         // Key: 1
	RequiredQuantity int   `json:"requiredQuantity" msgpack:"requiredQuantity"` // Key: 2
}

type PosterStoryMaster struct {
	Id                    int64              `json:"id" msgpack:"id"`                                       // Key: 0
	PosterMasterId        int64              `json:"posterMasterId" msgpack:"posterMasterId"`               // Key: 1
	EpisodeType           PosterEpisodeTypes `json:"episodeType" msgpack:"episodeType"`                     // Key: 2
	CharacterBaseMasterId *int64             `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 3
	Description           string             `json:"description" msgpack:"description"`                     // Key: 4
	Order                 int                `json:"order" msgpack:"order"`                                 // Key: 5
	CharacterIconId       *int64             `json:"characterIconId" msgpack:"characterIconId"`             // Key: 6
	CharacterName         string             `json:"characterName" msgpack:"characterName"`                 // Key: 7
}

type RandomEffectGroupMaster struct {
	Id               int64   `json:"id" msgpack:"id"`                             // Key: 0
	AccessoryEffects []int64 `json:"accessoryEffects" msgpack:"accessoryEffects"` // Key: 1
}

type ResultVoiceMaster struct {
	Id                    int64                     `json:"id" msgpack:"id"`                                       // Key: 0
	CharacterBaseMasterId int64                     `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	Order                 int                       `json:"order" msgpack:"order"`                                 // Key: 2
	ConditionType         ResultVoiceConditionTypes `json:"conditionType" msgpack:"conditionType"`                 // Key: 3
	VoiceAssetId          string                    `json:"voiceAssetId" msgpack:"voiceAssetId"`                   // Key: 4
	Motion                string                    `json:"motion" msgpack:"motion"`                               // Key: 5
}

type RewardRuleMaster struct {
	Id                    int64                        `json:"id" msgpack:"id"`                                       // Key: 0
	AchivementRateRewards []AchivementRateRewardMaster `json:"achivementRateRewards" msgpack:"achivementRateRewards"` // Key: 1
}

type RouletteEventMaster struct {
	Id                        int64 `json:"id" msgpack:"id"`                                               // Key: 0
	EventMasterId             int64 `json:"eventMasterId" msgpack:"eventMasterId"`                         // Key: 1
	RoulettePointItemMasterId int64 `json:"roulettePointItemMasterId" msgpack:"roulettePointItemMasterId"` // Key: 2
}

type RouletteMaster struct {
	Id                         int64                      `json:"id" msgpack:"id"`                                                 // Key: 0
	RouletteEventMasterId      int64                      `json:"rouletteEventMasterId" msgpack:"rouletteEventMasterId"`           // Key: 1
	Order                      int                        `json:"order" msgpack:"order"`                                           // Key: 2
	RouletteType               RouletteTypes              `json:"rouletteType" msgpack:"rouletteType"`                             // Key: 3
	RouletteTicketItemMasterId int64                      `json:"rouletteTicketItemMasterId" msgpack:"rouletteTicketItemMasterId"` // Key: 4
	RequiredTicketQuantity     int                        `json:"requiredTicketQuantity" msgpack:"requiredTicketQuantity"`         // Key: 5
	RouletteImagePath          string                     `json:"rouletteImagePath" msgpack:"rouletteImagePath"`                   // Key: 6
	Prizes                     []RoulettePrizeMaster      `json:"prizes" msgpack:"prizes"`                                         // Key: 7
	RollRewards                []RouletteRollRewardMaster `json:"rollRewards" msgpack:"rollRewards"`                               // Key: 8
	Name                       string                     `json:"name" msgpack:"name"`                                             // Key: 9
}

type RoulettePrizeMaster struct {
	Id                   int64                     `json:"id" msgpack:"id"`                                     // Key: 0
	Slot                 int                       `json:"slot" msgpack:"slot"`                                 // Key: 1
	AcquirePointQuantity *int                      `json:"acquirePointQuantity" msgpack:"acquirePointQuantity"` // Key: 2
	AcquirePrizeThing    *RoulettePrizeThingMaster `json:"acquirePrizeThing" msgpack:"acquirePrizeThing"`       // Key: 3
}

type RoulettePrizeThingMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type RouletteRollRewardMaster struct {
	RollCount int                       `json:"rollCount" msgpack:"rollCount"` // Key: 0
	Rewards   []RouletteRollRewardThing `json:"rewards" msgpack:"rewards"`     // Key: 1
}

type RouletteRollRewardThing struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type SceneCameraMaster struct {
	Id                              int64 `json:"id" msgpack:"id"`                                                           // Key: 0
	StartPositionX                  int   `json:"startPositionX" msgpack:"startPositionX"`                                   // Key: 1
	StartPositionY                  int   `json:"startPositionY" msgpack:"startPositionY"`                                   // Key: 2
	StartZoomRatio                  int   `json:"startZoomRatio" msgpack:"startZoomRatio"`                                   // Key: 3
	EndPositionX                    int   `json:"endPositionX" msgpack:"endPositionX"`                                       // Key: 4
	EndPositionY                    int   `json:"endPositionY" msgpack:"endPositionY"`                                       // Key: 5
	EndZoomRatio                    int   `json:"endZoomRatio" msgpack:"endZoomRatio"`                                       // Key: 6
	CameraMoveTurnaroundTimeSeconds int   `json:"cameraMoveTurnaroundTimeSeconds" msgpack:"cameraMoveTurnaroundTimeSeconds"` // Key: 7
}

type ScoreRewardMaster struct {
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type SenseBranchMaster struct {
	Id      int64               `json:"id" msgpack:"id"`           // Key: 0
	Order   int                 `json:"order" msgpack:"order"`     // Key: 1
	Effects []SenseEffectMaster `json:"effects" msgpack:"effects"` // Key: 2
}

type SenseEffectMaster struct {
	Order          int   `json:"order" msgpack:"order"`                   // Key: 0
	EffectMasterId int64 `json:"effectMasterId" msgpack:"effectMasterId"` // Key: 1
}

type SenseMaster struct {
	Id                     int64               `json:"id" msgpack:"id"`                                         // Key: 0
	Name                   string              `json:"name" msgpack:"name"`                                     // Key: 1
	Description            string              `json:"description" msgpack:"description"`                       // Key: 2
	Type                   SenseTypes          `json:"type" msgpack:"type"`                                     // Key: 3
	PreEffects             []EffectOrderMaster `json:"preEffects" msgpack:"preEffects"`                         // Key: 4
	Branches               []BranchMaster      `json:"branches" msgpack:"branches"`                             // Key: 5
	AcquirableGauge        int                 `json:"acquirableGauge" msgpack:"acquirableGauge"`               // Key: 6
	AcquirableScorePercent int                 `json:"acquirableScorePercent" msgpack:"acquirableScorePercent"` // Key: 7
	ScoreUpPerLevel        int                 `json:"scoreUpPerLevel" msgpack:"scoreUpPerLevel"`               // Key: 8
	LightCount             int                 `json:"lightCount" msgpack:"lightCount"`                         // Key: 9
	CoolTime               int                 `json:"coolTime" msgpack:"coolTime"`                             // Key: 10
	BranchCondition1       BranchConditionType `json:"branchCondition1" msgpack:"branchCondition1"`             // Key: 11
	ConditionValue1        *int64              `json:"conditionValue1" msgpack:"conditionValue1"`               // Key: 12
	BranchCondition2       BranchConditionType `json:"branchCondition2" msgpack:"branchCondition2"`             // Key: 13
	ConditionValue2        *int64              `json:"conditionValue2" msgpack:"conditionValue2"`               // Key: 14
	SubTypes               []SenseTypes        `json:"subTypes" msgpack:"subTypes"`                             // Key: 15
}

type SenseNotationBuffMaster struct {
	Id          int64                    `json:"id" msgpack:"id"`                   // Key: 0
	Type        SenseNotationBuffTypes   `json:"type" msgpack:"type"`               // Key: 1
	TargetValue *int64                   `json:"targetValue" msgpack:"targetValue"` // Key: 2
	BuffValue   int                      `json:"buffValue" msgpack:"buffValue"`     // Key: 3
	StatusType  SenseNotationStatusTypes `json:"statusType" msgpack:"statusType"`   // Key: 4
}

type SenseNotationDetailMaster struct {
	Id           int64 `json:"id" msgpack:"id"`                     // Key: 0
	Position     int   `json:"position" msgpack:"position"`         // Key: 1
	TimingSecond int   `json:"timingSecond" msgpack:"timingSecond"` // Key: 2
}

type SenseNotationMaster struct {
	Id      int64                       `json:"id" msgpack:"id"`           // Key: 0
	Details []SenseNotationDetailMaster `json:"details" msgpack:"details"` // Key: 1
	Buffs   []SenseNotationBuffMaster   `json:"buffs" msgpack:"buffs"`     // Key: 2
}

type SensePerformanceCharacterMaster struct {
	Id                       int64  `json:"id" msgpack:"id"`                                             // Key: 0
	CharacterBaseMasterId    int64  `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`       // Key: 1
	SensePerformanceMasterId int64  `json:"sensePerformanceMasterId" msgpack:"sensePerformanceMasterId"` // Key: 2
	ExpressId                string `json:"expressId" msgpack:"expressId"`                               // Key: 3
}

type SensePerformanceMaster struct {
	Id               int64                             `json:"id" msgpack:"id"`                             // Key: 0
	LabelType        int                               `json:"labelType" msgpack:"labelType"`               // Key: 1
	PerformanceGroup *int                              `json:"performanceGroup" msgpack:"performanceGroup"` // Key: 2
	Characters       []SensePerformanceCharacterMaster `json:"characters" msgpack:"characters"`             // Key: 3
}

type SignMaster struct {
	Id                    int64  `json:"id" msgpack:"id"`                                       // Key: 0
	CharacterBaseMasterId int64  `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	SignImagePath         string `json:"signImagePath" msgpack:"signImagePath"`                 // Key: 2
}

type SpecialEpisodeMaster struct {
	Id                          int64  `json:"id" msgpack:"id"`                                                   // Key: 0
	StoryMasterId               int64  `json:"storyMasterId" msgpack:"storyMasterId"`                             // Key: 1
	Title                       string `json:"title" msgpack:"title"`                                             // Key: 2
	RequiredReadEpisodeMasterId *int64 `json:"requiredReadEpisodeMasterId" msgpack:"requiredReadEpisodeMasterId"` // Key: 3
}

type SpecialEventLayout struct {
	Id             int64                   `json:"id" msgpack:"id"`                         // Key: 0
	UIType         CustomLayoutUITypes     `json:"uIType" msgpack:"uIType"`                 // Key: 1
	UIValue        string                  `json:"uIValue" msgpack:"uIValue"`               // Key: 2
	AnchorXPercent int                     `json:"anchorXPercent" msgpack:"anchorXPercent"` // Key: 3
	AnchorYPercent int                     `json:"anchorYPercent" msgpack:"anchorYPercent"` // Key: 4
	PositionX      float32                 `json:"positionX" msgpack:"positionX"`           // Key: 5
	PositionY      float32                 `json:"positionY" msgpack:"positionY"`           // Key: 6
	SizeX          float32                 `json:"sizeX" msgpack:"sizeX"`                   // Key: 7
	SizeY          float32                 `json:"sizeY" msgpack:"sizeY"`                   // Key: 8
	ScaleX         float32                 `json:"scaleX" msgpack:"scaleX"`                 // Key: 9
	ScaleY         float32                 `json:"scaleY" msgpack:"scaleY"`                 // Key: 10
	Layer          int                     `json:"layer" msgpack:"layer"`                   // Key: 11
	ActionType     CustomLayoutActionTypes `json:"actionType" msgpack:"actionType"`         // Key: 12
	ActionValue    string                  `json:"actionValue" msgpack:"actionValue"`       // Key: 13
	JumpType       JumpTypes               `json:"jumpType" msgpack:"jumpType"`             // Key: 14
	JumpValue      string                  `json:"jumpValue" msgpack:"jumpValue"`           // Key: 15
	StartDate      time.Time               `json:"startDate" msgpack:"startDate"`           // Key: 16
	EndDate        time.Time               `json:"endDate" msgpack:"endDate"`               // Key: 17
}

type SpecialEventMaster struct {
	Id                int64                     `json:"id" msgpack:"id"`                               // Key: 0
	EventMaster       *EventMaster              `json:"eventMaster" msgpack:"eventMaster"`             // Key: 1
	EventName         string                    `json:"eventName" msgpack:"eventName"`                 // Key: 2
	CategoryType      SpecialEventCategoryTypes `json:"categoryType" msgpack:"categoryType"`           // Key: 3
	Layouts           []SpecialEventLayout      `json:"layouts" msgpack:"layouts"`                     // Key: 4
	GameHintPageCount int                       `json:"gameHintPageCount" msgpack:"gameHintPageCount"` // Key: 5
}

type SpecialStoryMaster struct {
	Id                  int64               `json:"id" msgpack:"id"`                                   // Key: 0
	StoryMasterId       int64               `json:"storyMasterId" msgpack:"storyMasterId"`             // Key: 1
	StartDate           time.Time           `json:"startDate" msgpack:"startDate"`                     // Key: 2
	EndDate             time.Time           `json:"endDate" msgpack:"endDate"`                         // Key: 3
	EventStoryListFlags EventStoryListFlags `json:"eventStoryListFlags" msgpack:"eventStoryListFlags"` // Key: 4
	Title               string              `json:"title" msgpack:"title"`                             // Key: 5
	Order               int                 `json:"order" msgpack:"order"`                             // Key: 6
}

type SplashMaster struct {
	Id                     int64                      `json:"id" msgpack:"id"`                                         // Key: 0
	Order                  int                        `json:"order" msgpack:"order"`                                   // Key: 1
	SplashType             SplashTypes                `json:"splashType" msgpack:"splashType"`                         // Key: 2
	SplashValue            string                     `json:"splashValue" msgpack:"splashValue"`                       // Key: 3
	StartDate              time.Time                  `json:"startDate" msgpack:"startDate"`                           // Key: 4
	EndDate                time.Time                  `json:"endDate" msgpack:"endDate"`                               // Key: 5
	AdditionalSplashValue0 SplashAdditionalValueTypes `json:"additionalSplashValue0" msgpack:"additionalSplashValue0"` // Key: 6
	AdditionalSplashValue1 SplashAdditionalValueTypes `json:"additionalSplashValue1" msgpack:"additionalSplashValue1"` // Key: 7
	UnlockCondition        SplashUnlockConditionTypes `json:"unlockCondition" msgpack:"unlockCondition"`               // Key: 8
	UnlockConditionValue   *int64                     `json:"unlockConditionValue" msgpack:"unlockConditionValue"`     // Key: 9
}

type SpotConversationMaster struct {
	Id              int64     `json:"id" msgpack:"id"`                           // Key: 0
	Spot            SpotTypes `json:"spot" msgpack:"spot"`                       // Key: 1
	CharacterId1    *int64    `json:"characterId1" msgpack:"characterId1"`       // Key: 2
	CharacterId2    *int64    `json:"characterId2" msgpack:"characterId2"`       // Key: 3
	CharacterId3    *int64    `json:"characterId3" msgpack:"characterId3"`       // Key: 4
	CharacterId4    *int64    `json:"characterId4" msgpack:"characterId4"`       // Key: 5
	CharacterId5    *int64    `json:"characterId5" msgpack:"characterId5"`       // Key: 8
	EpisodeMasterId int64     `json:"episodeMasterId" msgpack:"episodeMasterId"` // Key: 9
	CostumeId1      *int64    `json:"costumeId1" msgpack:"costumeId1"`           // Key: 10
	CostumeId2      *int64    `json:"costumeId2" msgpack:"costumeId2"`           // Key: 11
	CostumeId3      *int64    `json:"costumeId3" msgpack:"costumeId3"`           // Key: 12
	CostumeId4      *int64    `json:"costumeId4" msgpack:"costumeId4"`           // Key: 13
	CostumeId5      *int64    `json:"costumeId5" msgpack:"costumeId5"`           // Key: 14
	Title           string    `json:"title" msgpack:"title"`                     // Key: 15
}

type StaminaRecoveryItemMaster struct {
	ItemMasterId    int64      `json:"itemMasterId" msgpack:"itemMasterId"`       // Key: 0
	RecoveryAmount  int        `json:"recoveryAmount" msgpack:"recoveryAmount"`   // Key: 1
	ExpiredAt       *time.Time `json:"expiredAt" msgpack:"expiredAt"`             // Key: 2
	IsForceShowList bool       `json:"isForceShowList" msgpack:"isForceShowList"` // Key: 3
}

type StampMaster struct {
	Id                     int64     `json:"id" msgpack:"id"`                                         // Key: 0
	Order                  int       `json:"order" msgpack:"order"`                                   // Key: 1
	IsDefault              bool      `json:"isDefault" msgpack:"isDefault"`                           // Key: 2
	CharacterBaseMasterId  int64     `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`   // Key: 3
	Type                   StampType `json:"type" msgpack:"type"`                                     // Key: 4
	AssetId                string    `json:"assetId" msgpack:"assetId"`                               // Key: 5
	VoiceAssetId           string    `json:"voiceAssetId" msgpack:"voiceAssetId"`                     // Key: 6
	Name                   string    `json:"name" msgpack:"name"`                                     // Key: 7
	CharacterBaseMasterIds []int64   `json:"characterBaseMasterIds" msgpack:"characterBaseMasterIds"` // Key: 8
}

type StarActBranchMaster struct {
	Id      int64               `json:"id" msgpack:"id"`           // Key: 0
	Order   int                 `json:"order" msgpack:"order"`     // Key: 1
	Effects []SenseEffectMaster `json:"effects" msgpack:"effects"` // Key: 2
}

type StarActConditionMaster struct {
	Id                 int64 `json:"id" msgpack:"id"`                                 // Key: 0
	FreeLight          int   `json:"freeLight" msgpack:"freeLight"`                   // Key: 1
	SupportLight       int   `json:"supportLight" msgpack:"supportLight"`             // Key: 2
	ControlLight       int   `json:"controlLight" msgpack:"controlLight"`             // Key: 3
	AmplificationLight int   `json:"amplificationLight" msgpack:"amplificationLight"` // Key: 4
	SpecialLight       int   `json:"specialLight" msgpack:"specialLight"`             // Key: 5
}

type StarActMaster struct {
	Id                       int64               `json:"id" msgpack:"id"`                                             // Key: 0
	Name                     string              `json:"name" msgpack:"name"`                                         // Key: 1
	Description              string              `json:"description" msgpack:"description"`                           // Key: 2
	StarActConditionMasterId int64               `json:"starActConditionMasterId" msgpack:"starActConditionMasterId"` // Key: 4
	AcquirableScorePercent   int                 `json:"acquirableScorePercent" msgpack:"acquirableScorePercent"`     // Key: 5
	ScoreUpPerLevel          int                 `json:"scoreUpPerLevel" msgpack:"scoreUpPerLevel"`                   // Key: 6
	PreEffects               []EffectOrderMaster `json:"preEffects" msgpack:"preEffects"`                             // Key: 7
	Branches                 []BranchMaster      `json:"branches" msgpack:"branches"`                                 // Key: 8
	BranchCondition1         BranchConditionType `json:"branchCondition1" msgpack:"branchCondition1"`                 // Key: 9
	ConditionValue1          *int64              `json:"conditionValue1" msgpack:"conditionValue1"`                   // Key: 10
	BranchCondition2         BranchConditionType `json:"branchCondition2" msgpack:"branchCondition2"`                 // Key: 11
	ConditionValue2          *int64              `json:"conditionValue2" msgpack:"conditionValue2"`                   // Key: 12
}

type StarRankRewardMaster struct {
	Rank                                 int   `json:"rank" msgpack:"rank"`                                                                 // Key: 0
	CharacterBaseMasterId                int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`                               // Key: 1
	CharacterStarRankRewardGroupMasterId int64 `json:"characterStarRankRewardGroupMasterId" msgpack:"characterStarRankRewardGroupMasterId"` // Key: 2
}

type StepupGachaGroupMaster struct {
	Id           int64               `json:"id" msgpack:"id"`                     // Key: 0
	StepupGachas []StepupGachaMaster `json:"stepupGachas" msgpack:"stepupGachas"` // Key: 1
}

type StepupGachaMaster struct {
	GachaMasterId int64 `json:"gachaMasterId" msgpack:"gachaMasterId"` // Key: 0
	Step          int   `json:"step" msgpack:"step"`                   // Key: 1
}

type StoryEventBonusCharacterBaseMaster struct {
	Id                 int64                `json:"id" msgpack:"id"`                                 // Key: 0
	StoryEventMasterId int64                `json:"storyEventMasterId" msgpack:"storyEventMasterId"` // Key: 1
	BonusTargetType    StoryEventBonusTypes `json:"bonusTargetType" msgpack:"bonusTargetType"`       // Key: 2
	BonusTargetId      int64                `json:"bonusTargetId" msgpack:"bonusTargetId"`           // Key: 3
}

type StoryEventCircleHighScoreRewardMaster struct {
	Id                 int64      `json:"id" msgpack:"id"`                                 // Key: 0
	StoryEventMasterId int64      `json:"storyEventMasterId" msgpack:"storyEventMasterId"` // Key: 1
	Score              int64      `json:"score" msgpack:"score"`                           // Key: 2
	ThingId            int64      `json:"thingId" msgpack:"thingId"`                       // Key: 3
	ThingType          ThingTypes `json:"thingType" msgpack:"thingType"`                   // Key: 4
	ThingQuantity      int        `json:"thingQuantity" msgpack:"thingQuantity"`           // Key: 5
}

type StoryEventCircleMaster struct {
	Id                 int64 `json:"id" msgpack:"id"`                                 // Key: 0
	StoryEventMasterId int64 `json:"storyEventMasterId" msgpack:"storyEventMasterId"` // Key: 1
	MusicMasterId      int64 `json:"musicMasterId" msgpack:"musicMasterId"`           // Key: 2
}

type StoryEventCircleMissionMaster struct {
	Id                 int64  `json:"id" msgpack:"id"`                                 // Key: 0
	StoryEventMasterId int64  `json:"storyEventMasterId" msgpack:"storyEventMasterId"` // Key: 1
	Title              string `json:"title" msgpack:"title"`                           // Key: 2
	CountUnit          string `json:"countUnit" msgpack:"countUnit"`                   // Key: 3
	GoalCount          int64  `json:"goalCount" msgpack:"goalCount"`                   // Key: 4
	MaxPhase           int    `json:"maxPhase" msgpack:"maxPhase"`                     // Key: 5
	GivePointAtGoal    int    `json:"givePointAtGoal" msgpack:"givePointAtGoal"`       // Key: 6
}

type StoryEventCircleMissionRewardMaster struct {
	Id                  int64      `json:"id" msgpack:"id"`                                   // Key: 0
	StoryEventMasterId  int64      `json:"storyEventMasterId" msgpack:"storyEventMasterId"`   // Key: 1
	RequiredCirclePoint int        `json:"requiredCirclePoint" msgpack:"requiredCirclePoint"` // Key: 2
	RequiredUserPoint   int        `json:"requiredUserPoint" msgpack:"requiredUserPoint"`     // Key: 3
	ThingId             int64      `json:"thingId" msgpack:"thingId"`                         // Key: 4
	ThingType           ThingTypes `json:"thingType" msgpack:"thingType"`                     // Key: 5
	ThingQuantity       int        `json:"thingQuantity" msgpack:"thingQuantity"`             // Key: 6
}

type StoryEventEpisodeMaster struct {
	Id                           int64  `json:"id" msgpack:"id"`                                                     // Key: 0
	StoryMasterId                int64  `json:"storyMasterId" msgpack:"storyMasterId"`                               // Key: 1
	RequiredTotalAcquiredPoint   int64  `json:"requiredTotalAcquiredPoint" msgpack:"requiredTotalAcquiredPoint"`     // Key: 3
	Title                        string `json:"title" msgpack:"title"`                                               // Key: 4
	Order                        int    `json:"order" msgpack:"order"`                                               // Key: 5
	RequiredReadEpisodeMasterId  *int64 `json:"requiredReadEpisodeMasterId" msgpack:"requiredReadEpisodeMasterId"`   // Key: 6
	EpisodeRewardPackageMasterId int64  `json:"episodeRewardPackageMasterId" msgpack:"episodeRewardPackageMasterId"` // Key: 7
}

type StoryEventHighScoreBuffMaster struct {
	Id                                    int64  `json:"id" msgpack:"id"`                                                                       // Key: 0
	StoryEventHighScoreBuffPatternGroupId int64  `json:"storyEventHighScoreBuffPatternGroupId" msgpack:"storyEventHighScoreBuffPatternGroupId"` // Key: 1
	EffectMasterId                        int64  `json:"effectMasterId" msgpack:"effectMasterId"`                                               // Key: 2
	EffectName                            string `json:"effectName" msgpack:"effectName"`                                                       // Key: 3
	EffectDescription                     string `json:"effectDescription" msgpack:"effectDescription"`                                         // Key: 4
	IconImagePath                         string `json:"iconImagePath" msgpack:"iconImagePath"`                                                 // Key: 5
}

type StoryEventHighScoreBuffPatternGroupMaster struct {
	Id       int64                                  `json:"id" msgpack:"id"`             // Key: 0
	Patterns []StoryEventHighScoreBuffPatternMaster `json:"patterns" msgpack:"patterns"` // Key: 1
}

type StoryEventHighScoreBuffPatternMaster struct {
	NextLevel     int `json:"nextLevel" msgpack:"nextLevel"`         // Key: 0
	RequiredPoint int `json:"requiredPoint" msgpack:"requiredPoint"` // Key: 1
}

type StoryEventHighScoreBuffSettingMaster struct {
	Id                              int64                          `json:"id" msgpack:"id"`                                                           // Key: 0
	StoryEventHighScoreMasterId     int64                          `json:"storyEventHighScoreMasterId" msgpack:"storyEventHighScoreMasterId"`         // Key: 1
	StoryEventHighScoreBuffMasterId int64                          `json:"storyEventHighScoreBuffMasterId" msgpack:"storyEventHighScoreBuffMasterId"` // Key: 2
	StoryEventHighScoreBuffMaster   *StoryEventHighScoreBuffMaster `json:"storyEventHighScoreBuffMaster" msgpack:"storyEventHighScoreBuffMaster"`     // Key: 3
}

type StoryEventHighScoreMaster struct {
	Id                    int64   `json:"id" msgpack:"id"`                                       // Key: 0
	StoryEventMasterId    int64   `json:"storyEventMasterId" msgpack:"storyEventMasterId"`       // Key: 1
	MusicMasterId         int64   `json:"musicMasterId" msgpack:"musicMasterId"`                 // Key: 2
	BuffIds               []int64 `json:"buffIds" msgpack:"buffIds"`                             // Key: 3
	SenseNotationMasterId int64   `json:"senseNotationMasterId" msgpack:"senseNotationMasterId"` // Key: 4
}

type StoryEventHighScoreRewardMaster struct {
	Id                 int64      `json:"id" msgpack:"id"`                                 // Key: 0
	StoryEventMasterId int64      `json:"storyEventMasterId" msgpack:"storyEventMasterId"` // Key: 1
	Score              int64      `json:"score" msgpack:"score"`                           // Key: 2
	ThingId            int64      `json:"thingId" msgpack:"thingId"`                       // Key: 3
	ThingType          ThingTypes `json:"thingType" msgpack:"thingType"`                   // Key: 4
	ThingQuantity      int        `json:"thingQuantity" msgpack:"thingQuantity"`           // Key: 5
}

type StoryEventMaster struct {
	Id                      int64                   `json:"id" msgpack:"id"`                                           // Key: 0
	Category                StoryEventCategoryTypes `json:"category" msgpack:"category"`                               // Key: 1
	BonusAttribute          *Attributes             `json:"bonusAttribute" msgpack:"bonusAttribute"`                   // Key: 2
	Title                   string                  `json:"title" msgpack:"title"`                                     // Key: 3
	ExchangeShopMasterId    int64                   `json:"exchangeShopMasterId" msgpack:"exchangeShopMasterId"`       // Key: 4
	EventPointItemMasterId  int64                   `json:"eventPointItemMasterId" msgpack:"eventPointItemMasterId"`   // Key: 5
	StartDate               time.Time               `json:"startDate" msgpack:"startDate"`                             // Key: 6
	EndDate                 time.Time               `json:"endDate" msgpack:"endDate"`                                 // Key: 7
	ForceEndDate            time.Time               `json:"forceEndDate" msgpack:"forceEndDate"`                       // Key: 8
	BonusCategoryMasterId10 *int64                  `json:"bonusCategoryMasterId10" msgpack:"bonusCategoryMasterId10"` // Key: 9
	BonusCategoryMasterId20 *int64                  `json:"bonusCategoryMasterId20" msgpack:"bonusCategoryMasterId20"` // Key: 10
	SecondaryBonusAttribute *Attributes             `json:"secondaryBonusAttribute" msgpack:"secondaryBonusAttribute"` // Key: 11
}

type StoryEventRewardMaster struct {
	Id                            int64 `json:"id" msgpack:"id"`                                                       // Key: 0
	StoryEventMasterId            int64 `json:"storyEventMasterId" msgpack:"storyEventMasterId"`                       // Key: 1
	TargetMinRank                 int   `json:"targetMinRank" msgpack:"targetMinRank"`                                 // Key: 2
	TargetMaxRank                 int   `json:"targetMaxRank" msgpack:"targetMaxRank"`                                 // Key: 3
	StoryEventRewardItemPackageId int64 `json:"storyEventRewardItemPackageId" msgpack:"storyEventRewardItemPackageId"` // Key: 4
}

type StoryEventRewardItemPackage struct {
	Id      int64                   `json:"id" msgpack:"id"`           // Key: 0
	Rewards []StoryEventRewardThing `json:"rewards" msgpack:"rewards"` // Key: 1
}

type StoryEventRewardThing struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type StoryEventStoryBgmGroupMaster struct {
	StoryEventMasterId int64                        `json:"storyEventMasterId" msgpack:"storyEventMasterId"` // Key: 0
	Schedules          []StoryEventStoryBgmSchedule `json:"schedules" msgpack:"schedules"`                   // Key: 1
}

type StoryEventStoryBgmSchedule struct {
	AssetName string    `json:"assetName" msgpack:"assetName"` // Key: 0
	StartDate time.Time `json:"startDate" msgpack:"startDate"` // Key: 1
	EndDate   time.Time `json:"endDate" msgpack:"endDate"`     // Key: 2
}

type StoryEventStoryMaster struct {
	Id                  int64               `json:"id" msgpack:"id"`                                   // Key: 0
	StoryMasterId       int64               `json:"storyMasterId" msgpack:"storyMasterId"`             // Key: 1
	StoryEventMasterId  int64               `json:"storyEventMasterId" msgpack:"storyEventMasterId"`   // Key: 2
	ListedStartDate     time.Time           `json:"listedStartDate" msgpack:"listedStartDate"`         // Key: 3
	EventStoryListFlags EventStoryListFlags `json:"eventStoryListFlags" msgpack:"eventStoryListFlags"` // Key: 4
	IsKeyStory          bool                `json:"isKeyStory" msgpack:"isKeyStory"`                   // Key: 5
}

type StoryEventTotalPointRewardMaster struct {
	Id                 int64      `json:"id" msgpack:"id"`                                 // Key: 0
	StoryEventMasterId int64      `json:"storyEventMasterId" msgpack:"storyEventMasterId"` // Key: 1
	RequiredTotalPoint int        `json:"requiredTotalPoint" msgpack:"requiredTotalPoint"` // Key: 2
	ThingType          ThingTypes `json:"thingType" msgpack:"thingType"`                   // Key: 3
	ThingId            int64      `json:"thingId" msgpack:"thingId"`                       // Key: 4
	ThingQuantity      int        `json:"thingQuantity" msgpack:"thingQuantity"`           // Key: 5
}

type StoryMaster struct {
	Id              int64      `json:"id" msgpack:"id"`                           // Key: 0
	Type            StoryTypes `json:"type" msgpack:"type"`                       // Key: 1
	CompanyMasterId *int64     `json:"companyMasterId" msgpack:"companyMasterId"` // Key: 2
	EventMasterId   *int64     `json:"eventMasterId" msgpack:"eventMasterId"`     // Key: 3
	ChapterOrder    int        `json:"chapterOrder" msgpack:"chapterOrder"`       // Key: 4
	DisplayStartAt  time.Time  `json:"displayStartAt" msgpack:"displayStartAt"`   // Key: 5
	DisplayEndAt    time.Time  `json:"displayEndAt" msgpack:"displayEndAt"`       // Key: 6
}

type StoryRelationMaster struct {
	StoryMasterId               int64   `json:"storyMasterId" msgpack:"storyMasterId"`                             // Key: 0
	PreviousStoryMasterIds      []int64 `json:"previousStoryMasterIds" msgpack:"previousStoryMasterIds"`           // Key: 1
	NextStoryMasterIds          []int64 `json:"nextStoryMasterIds" msgpack:"nextStoryMasterIds"`                   // Key: 2
	SideStoryCharacterMasterIds []int64 `json:"sideStoryCharacterMasterIds" msgpack:"sideStoryCharacterMasterIds"` // Key: 3
	RelatedStoryMasterIds       []int64 `json:"relatedStoryMasterIds" msgpack:"relatedStoryMasterIds"`             // Key: 4
}

type TeamChallengeMaster struct {
	Id            int64                        `json:"id" msgpack:"id"`                       // Key: 0
	MusicMasterId int64                        `json:"musicMasterId" msgpack:"musicMasterId"` // Key: 1
	Difficulty    TeamChallengeDifficultyTypes `json:"difficulty" msgpack:"difficulty"`       // Key: 2
	GoalType      TeamChallengeGoalTypes       `json:"goalType" msgpack:"goalType"`           // Key: 3
	GoalMinValue  int64                        `json:"goalMinValue" msgpack:"goalMinValue"`   // Key: 4
	GoalMaxValue  int64                        `json:"goalMaxValue" msgpack:"goalMaxValue"`   // Key: 5
	StartDate     time.Time                    `json:"startDate" msgpack:"startDate"`         // Key: 6
	EndDate       time.Time                    `json:"endDate" msgpack:"endDate"`             // Key: 7
}

type TheaterChapterMaster struct {
	Id                          int64                `json:"id" msgpack:"id"`                                                   // Key: 0
	Name                        string               `json:"name" msgpack:"name"`                                               // Key: 1
	CompanyMasterId             int64                `json:"companyMasterId" msgpack:"companyMasterId"`                         // Key: 2
	DisplayStartAt              time.Time            `json:"displayStartAt" msgpack:"displayStartAt"`                           // Key: 3
	DisplayEndAt                time.Time            `json:"displayEndAt" msgpack:"displayEndAt"`                               // Key: 4
	Stories                     []TheaterStoryMaster `json:"stories" msgpack:"stories"`                                         // Key: 5
	EpisodeMasterId             *int64               `json:"episodeMasterId" msgpack:"episodeMasterId"`                         // Key: 6
	RequiredReadEpisodeMasterId int64                `json:"requiredReadEpisodeMasterId" msgpack:"requiredReadEpisodeMasterId"` // Key: 7
	TheaterStoryMasterIds       []int64              `json:"theaterStoryMasterIds" msgpack:"theaterStoryMasterIds"`             // Key: 8
}

type TheaterDetailMaster struct {
	Id           int64  `json:"id" msgpack:"id"`                     // Key: 0
	Speaker      string `json:"speaker" msgpack:"speaker"`           // Key: 1
	Phrase       string `json:"phrase" msgpack:"phrase"`             // Key: 2
	VoiceAssetId string `json:"voiceAssetId" msgpack:"voiceAssetId"` // Key: 3
}

type TheaterRoleMaster struct {
	RoleName              string `json:"roleName" msgpack:"roleName"`                           // Key: 0
	CharacterBaseMasterId int64  `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	CostumeMasterId       int64  `json:"costumeMasterId" msgpack:"costumeMasterId"`             // Key: 2
	Order                 int    `json:"order" msgpack:"order"`                                 // Key: 3
}

type TheaterStoryMaster struct {
	Id             int64                 `json:"id" msgpack:"id"`                         // Key: 0
	Name           string                `json:"name" msgpack:"name"`                     // Key: 1
	Order          int                   `json:"order" msgpack:"order"`                   // Key: 2
	HasSynopsis    bool                  `json:"hasSynopsis" msgpack:"hasSynopsis"`       // Key: 3
	Synopsis       string                `json:"synopsis" msgpack:"synopsis"`             // Key: 4
	DisplayStartAt time.Time             `json:"displayStartAt" msgpack:"displayStartAt"` // Key: 5
	DisplayEndAt   time.Time             `json:"displayEndAt" msgpack:"displayEndAt"`     // Key: 6
	Roles          []TheaterRoleMaster   `json:"roles" msgpack:"roles"`                   // Key: 7
	Details        []TheaterDetailMaster `json:"details" msgpack:"details"`               // Key: 8
}

type TimeLimitedControlMaster struct {
	Id      int64     `json:"id" msgpack:"id"`           // Key: 0
	EndDate time.Time `json:"endDate" msgpack:"endDate"` // Key: 1
}

type TipMaster struct {
	Id                    int64  `json:"id" msgpack:"id"`                                       // Key: 0
	CharacterBaseMasterId int64  `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	Description           string `json:"description" msgpack:"description"`                     // Key: 2
	Name                  string `json:"name" msgpack:"name"`                                   // Key: 3
}

type TitleBackgroundDetailMaster struct {
	Id               int64                    `json:"id" msgpack:"id"`                             // Key: 0
	PathType         TitleBackgroundPathTypes `json:"pathType" msgpack:"pathType"`                 // Key: 1
	PathValue        string                   `json:"pathValue" msgpack:"pathValue"`               // Key: 2
	IsTitleHide      bool                     `json:"isTitleHide" msgpack:"isTitleHide"`           // Key: 3
	IsDecorationHide bool                     `json:"isDecorationHide" msgpack:"isDecorationHide"` // Key: 4
}

type TitleBackgroundMaster struct {
	Id        int64                         `json:"id" msgpack:"id"`               // Key: 0
	Priority  TitleBackgroundPriorityTypes  `json:"priority" msgpack:"priority"`   // Key: 1
	StartDate time.Time                     `json:"startDate" msgpack:"startDate"` // Key: 2
	EndDate   time.Time                     `json:"endDate" msgpack:"endDate"`     // Key: 3
	Details   []TitleBackgroundDetailMaster `json:"details" msgpack:"details"`     // Key: 4
}

type TitleCallVoiceMaster struct {
	Id        int64                       `json:"id" msgpack:"id"`               // Key: 0
	Priority  TitleCallVoicePriorityTypes `json:"priority" msgpack:"priority"`   // Key: 1
	VoiceId   int64                       `json:"voiceId" msgpack:"voiceId"`     // Key: 2
	StartDate time.Time                   `json:"startDate" msgpack:"startDate"` // Key: 3
	EndDate   time.Time                   `json:"endDate" msgpack:"endDate"`     // Key: 4
}

type TitleDecorationMaster struct {
	Id             int64                `json:"id" msgpack:"id"`                         // Key: 0
	DecorationType TitleDecorationTypes `json:"decorationType" msgpack:"decorationType"` // Key: 1
	StartDate      time.Time            `json:"startDate" msgpack:"startDate"`           // Key: 2
	EndDate        time.Time            `json:"endDate" msgpack:"endDate"`               // Key: 3
}

type TotalPointEventMaster struct {
	Id                   int64 `json:"id" msgpack:"id"`                                     // Key: 0
	SpecialEventMasterId int64 `json:"specialEventMasterId" msgpack:"specialEventMasterId"` // Key: 1
}

type TotalPointEventRewardMaster struct {
	Id                      int64      `json:"id" msgpack:"id"`                                           // Key: 0
	TotalPointEventMasterId int64      `json:"totalPointEventMasterId" msgpack:"totalPointEventMasterId"` // Key: 1
	Point                   int64      `json:"point" msgpack:"point"`                                     // Key: 2
	ThingType               ThingTypes `json:"thingType" msgpack:"thingType"`                             // Key: 3
	ThingId                 int64      `json:"thingId" msgpack:"thingId"`                                 // Key: 4
	ThingQuantity           int        `json:"thingQuantity" msgpack:"thingQuantity"`                     // Key: 5
	Order                   int        `json:"order" msgpack:"order"`                                     // Key: 6
}

type TournamentDetailMaster struct {
	Id            int64             `json:"id" msgpack:"id"`                       // Key: 0
	MusicMasterId int64             `json:"musicMasterId" msgpack:"musicMasterId"` // Key: 1
	Difficulty    MusicDifficulties `json:"difficulty" msgpack:"difficulty"`       // Key: 2
}

type TournamentMaster struct {
	Id        int64                    `json:"id" msgpack:"id"`               // Key: 0
	StartDate time.Time                `json:"startDate" msgpack:"startDate"` // Key: 1
	EndDate   time.Time                `json:"endDate" msgpack:"endDate"`     // Key: 2
	Details   []TournamentDetailMaster `json:"details" msgpack:"details"`     // Key: 3
}

type TournamentQualifyingMaster struct {
	Id             int64     `json:"id" msgpack:"id"`                         // Key: 0
	EventMasterId  int64     `json:"eventMasterId" msgpack:"eventMasterId"`   // Key: 1
	EntryPageUrl   string    `json:"entryPageUrl" msgpack:"entryPageUrl"`     // Key: 2
	EntryStartDate time.Time `json:"entryStartDate" msgpack:"entryStartDate"` // Key: 3
	EntryEndDate   time.Time `json:"entryEndDate" msgpack:"entryEndDate"`     // Key: 4
	WebsiteUrl     string    `json:"websiteUrl" msgpack:"websiteUrl"`         // Key: 5
}

type TrialPartyAccessoryMaster struct {
	Id                 int64   `json:"id" msgpack:"id"`                                 // Key: 0
	TrialPartyMasterId int64   `json:"trialPartyMasterId" msgpack:"trialPartyMasterId"` // Key: 1
	AccessoryMasterId  int64   `json:"accessoryMasterId" msgpack:"accessoryMasterId"`   // Key: 2
	Level              int     `json:"level" msgpack:"level"`                           // Key: 3
	AccessoryEffects   []int64 `json:"accessoryEffects" msgpack:"accessoryEffects"`     // Key: 4
}

type TrialPartyCharacterMaster struct {
	Id                 int64                 `json:"id" msgpack:"id"`                                 // Key: 0
	TrialPartyMasterId int64                 `json:"trialPartyMasterId" msgpack:"trialPartyMasterId"` // Key: 1
	CharacterMasterId  int64                 `json:"characterMasterId" msgpack:"characterMasterId"`   // Key: 2
	Level              int                   `json:"level" msgpack:"level"`                           // Key: 3
	TalentStage        int                   `json:"talentStage" msgpack:"talentStage"`               // Key: 4
	AwakeningStatus    int                   `json:"awakeningStatus" msgpack:"awakeningStatus"`       // Key: 5
	SenseLevel         int                   `json:"senseLevel" msgpack:"senseLevel"`                 // Key: 6
	ReadEpisodeOrder   CharacterEpisodeOrder `json:"readEpisodeOrder" msgpack:"readEpisodeOrder"`     // Key: 7
}

type TrialPartyEventMaster struct {
	Id             int64     `json:"id" msgpack:"id"`                         // Key: 0
	MEventId       int64     `json:"mEventId" msgpack:"mEventId"`             // Key: 1
	IsHideComplate bool      `json:"isHideComplate" msgpack:"isHideComplate"` // Key: 2
	StartDate      time.Time `json:"startDate" msgpack:"startDate"`           // Key: 3
	EndDate        time.Time `json:"endDate" msgpack:"endDate"`               // Key: 4
	BgmPath        string    `json:"bgmPath" msgpack:"bgmPath"`               // Key: 5
	Name           string    `json:"name" msgpack:"name"`                     // Key: 6
}

type TrialPartyEventStageMaster struct {
	Id                       int64                              `json:"id" msgpack:"id"`                                             // Key: 0
	TrialPartyEventMasterId  int64                              `json:"trialPartyEventMasterId" msgpack:"trialPartyEventMasterId"`   // Key: 1
	MusicMasterId            int64                              `json:"musicMasterId" msgpack:"musicMasterId"`                       // Key: 2
	VocalVersion             int                                `json:"vocalVersion" msgpack:"vocalVersion"`                         // Key: 3
	SenseNotationMasterId    int64                              `json:"senseNotationMasterId" msgpack:"senseNotationMasterId"`       // Key: 4
	ClearScore               int64                              `json:"clearScore" msgpack:"clearScore"`                             // Key: 5
	ClearSenseCount          int                                `json:"clearSenseCount" msgpack:"clearSenseCount"`                   // Key: 6
	ClearStarActCount        int                                `json:"clearStarActCount" msgpack:"clearStarActCount"`               // Key: 7
	ClearPrincipalGaugeValue int                                `json:"clearPrincipalGaugeValue" msgpack:"clearPrincipalGaugeValue"` // Key: 8
	Title                    string                             `json:"title" msgpack:"title"`                                       // Key: 9
	Hint                     string                             `json:"hint" msgpack:"hint"`                                         // Key: 10
	RestrictionInfo          string                             `json:"restrictionInfo" msgpack:"restrictionInfo"`                   // Key: 11
	Rewards                  []TrialPartyEventStageRewardMaster `json:"rewards" msgpack:"rewards"`                                   // Key: 12
	TrialPartyMasterId       int64                              `json:"trialPartyMasterId" msgpack:"trialPartyMasterId"`             // Key: 13
	UnrecommendedAuto        bool                               `json:"unrecommendedAuto" msgpack:"unrecommendedAuto"`               // Key: 14
}

type TrialPartyEventStageRewardMaster struct {
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type TrialPartyMaster struct {
	Id             int64                  `json:"id" msgpack:"id"`                         // Key: 0
	LeaderPosition int                    `json:"leaderPosition" msgpack:"leaderPosition"` // Key: 1
	DefaultSlots   []TrialPartySlotMaster `json:"defaultSlots" msgpack:"defaultSlots"`     // Key: 2
}

type TrialPartySlotMaster struct {
	Position                    int                          `json:"position" msgpack:"position"`                                       // Key: 0
	IsPositionLock              bool                         `json:"isPositionLock" msgpack:"isPositionLock"`                           // Key: 1
	TrialPartyCharacterMasterId int64                        `json:"trialPartyCharacterMasterId" msgpack:"trialPartyCharacterMasterId"` // Key: 2
	CharacterLockType           TrialPartyCharacterLockTypes `json:"characterLockType" msgpack:"characterLockType"`                     // Key: 3
	TrialPartyPosterMasterId    *int64                       `json:"trialPartyPosterMasterId" msgpack:"trialPartyPosterMasterId"`       // Key: 4
	PosterLockype               TrialPartyEquipmentLockTypes `json:"posterLockype" msgpack:"posterLockype"`                             // Key: 5
	TrialPartyAccessoryMasterId *int64                       `json:"trialPartyAccessoryMasterId" msgpack:"trialPartyAccessoryMasterId"` // Key: 6
	AccessoryLockType           TrialPartyEquipmentLockTypes `json:"accessoryLockType" msgpack:"accessoryLockType"`                     // Key: 7
}

type TrialPartyPosterMaster struct {
	Id                 int64 `json:"id" msgpack:"id"`                                 // Key: 0
	TrialPartyMasterId int64 `json:"trialPartyMasterId" msgpack:"trialPartyMasterId"` // Key: 1
	PosterMasterId     int64 `json:"posterMasterId" msgpack:"posterMasterId"`         // Key: 2
	Level              int   `json:"level" msgpack:"level"`                           // Key: 3
	BreakthroughPhase  int   `json:"breakthroughPhase" msgpack:"breakthroughPhase"`   // Key: 4
}

type TripleCastAllClassGlobalRankingRewardPackageMaster struct {
	Id                 int64                                              `json:"id" msgpack:"id"`                                 // Key: 0
	TripleCastMasterId int64                                              `json:"tripleCastMasterId" msgpack:"tripleCastMasterId"` // Key: 1
	Rewards            []TripleCastAllClassGlobalRankingRewardThingMaster `json:"rewards" msgpack:"rewards"`                       // Key: 2
	MinRank            int                                                `json:"minRank" msgpack:"minRank"`                       // Key: 3
	MaxRank            int                                                `json:"maxRank" msgpack:"maxRank"`                       // Key: 4
	RewardTitle        string                                             `json:"rewardTitle" msgpack:"rewardTitle"`               // Key: 5
}

type TripleCastAllClassGlobalRankingRewardThingMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type TripleCastGroupRankingRewardPackageMaster struct {
	Id          int64                                     `json:"id" msgpack:"id"`                   // Key: 0
	Rewards     []TripleCastGroupRankingRewardThingMaster `json:"rewards" msgpack:"rewards"`         // Key: 1
	ClassType   LeagueClassTypes                          `json:"classType" msgpack:"classType"`     // Key: 2
	MinRank     int                                       `json:"minRank" msgpack:"minRank"`         // Key: 3
	MaxRank     int                                       `json:"maxRank" msgpack:"maxRank"`         // Key: 4
	RewardTitle string                                    `json:"rewardTitle" msgpack:"rewardTitle"` // Key: 5
}

type TripleCastGroupRankingRewardThingMaster struct {
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 0
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
}

type TripleCastMaster struct {
	Id                          int64             `json:"id" msgpack:"id"`                                                   // Key: 0
	MusicMasterId               int64             `json:"musicMasterId" msgpack:"musicMasterId"`                             // Key: 1
	VocalVersion                int               `json:"vocalVersion" msgpack:"vocalVersion"`                               // Key: 2
	MaxDifficulty               MusicDifficulties `json:"maxDifficulty" msgpack:"maxDifficulty"`                             // Key: 3
	SenseNotationMasterId1      int64             `json:"senseNotationMasterId1" msgpack:"senseNotationMasterId1"`           // Key: 4
	SenseNotationMasterId2      int64             `json:"senseNotationMasterId2" msgpack:"senseNotationMasterId2"`           // Key: 5
	SenseNotationMasterId3      int64             `json:"senseNotationMasterId3" msgpack:"senseNotationMasterId3"`           // Key: 6
	DisplayStartAt              time.Time         `json:"displayStartAt" msgpack:"displayStartAt"`                           // Key: 7
	CountingStartAt             time.Time         `json:"countingStartAt" msgpack:"countingStartAt"`                         // Key: 8
	DisplayEndAt                time.Time         `json:"displayEndAt" msgpack:"displayEndAt"`                               // Key: 9
	LeagueClassGroupMasterId    int64             `json:"leagueClassGroupMasterId" msgpack:"leagueClassGroupMasterId"`       // Key: 10
	TripleCastSeasonMasterId    int64             `json:"tripleCastSeasonMasterId" msgpack:"tripleCastSeasonMasterId"`       // Key: 11
	VocalEffectMasterId         int64             `json:"vocalEffectMasterId" msgpack:"vocalEffectMasterId"`                 // Key: 12
	ExpressionEffectMasterId    int64             `json:"expressionEffectMasterId" msgpack:"expressionEffectMasterId"`       // Key: 13
	ConcentrationEffectMasterId int64             `json:"concentrationEffectMasterId" msgpack:"concentrationEffectMasterId"` // Key: 14
}

type TrophyGroupMaster struct {
	Id       int64            `json:"id" msgpack:"id"`             // Key: 0
	Category TrophyCategories `json:"category" msgpack:"category"` // Key: 1
}

type TrophyMaster struct {
	Id                  int64              `json:"id" msgpack:"id"`                                   // Key: 0
	Name                string             `json:"name" msgpack:"name"`                               // Key: 1
	Description         string             `json:"description" msgpack:"description"`                 // Key: 2
	Rarity              PossessionRarities `json:"rarity" msgpack:"rarity"`                           // Key: 3
	Order               int                `json:"order" msgpack:"order"`                             // Key: 4
	TrophyGroupMasterId int64              `json:"trophyGroupMasterId" msgpack:"trophyGroupMasterId"` // Key: 5
	Hidden              bool               `json:"hidden" msgpack:"hidden"`                           // Key: 6
	UnlockText          string             `json:"unlockText" msgpack:"unlockText"`                   // Key: 7
}

type UnlockConditionMaster struct {
	Id                         int64                      `json:"id" msgpack:"id"`                                                 // Key: 0
	ConditionText              string                     `json:"conditionText" msgpack:"conditionText"`                           // Key: 1
	ConditionUsableTableFilter ConditionUsableTableFilter `json:"conditionUsableTableFilter" msgpack:"conditionUsableTableFilter"` // Key: 2
	UnlockConditionType        UnlockConditionTypes       `json:"unlockConditionType" msgpack:"unlockConditionType"`               // Key: 3
}
