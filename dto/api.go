// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared
package dto

import (
	"time"
)

type Market struct {
	Id              int64     `json:"id" msgpack:"id"`                           // Key: 0
	LastRefreshedAt time.Time `json:"lastRefreshedAt" msgpack:"lastRefreshedAt"` // Key: 1
	RefreshTimes    int       `json:"refreshTimes" msgpack:"refreshTimes"`       // Key: 2
}

type DeletedDataObject struct {
	TypeName string `json:"typeName" msgpack:"typeName"` // Key: 0
	Id       int64  `json:"id" msgpack:"id"`             // Key: 1
}

type DummyEditPositionPayloadArray struct {
	Payloads []EditPositionPayload `json:"payloads" msgpack:"payloads"` // Key: 0
}

type DummyFriendResultArray struct {
	Payloads []FriendResult `json:"payloads" msgpack:"payloads"` // Key: 0
}

type DummyResponse struct {
	Fault                            []Fault                           `json:"fault" msgpack:"fault"`                                                       // Key: 0
	Added                            []any                             `json:"added" msgpack:"added"`                                                       // Key: 1
	Modified                         []any                             `json:"modified" msgpack:"modified"`                                                 // Key: 2
	Deleted                          []DeletedDataObject               `json:"deleted" msgpack:"deleted"`                                                   // Key: 3
	Notifications                    []any                             `json:"notifications" msgpack:"notifications"`                                       // Key: 4
	GachaInfoResults                 []GachaInfoResult                 `json:"gachaInfoResults" msgpack:"gachaInfoResults"`                                 // Key: 5
	CircleInformationResults         []CircleInformationResult         `json:"circleInformationResults" msgpack:"circleInformationResults"`                 // Key: 6
	UseExperienceItemsPayloads       []UseExperienceItemsPayload       `json:"useExperienceItemsPayloads" msgpack:"useExperienceItemsPayloads"`             // Key: 7
	LoginBonusResults                []LoginBonusResult                `json:"loginBonusResults" msgpack:"loginBonusResults"`                               // Key: 8
	NotificationResults              []NotificationResult              `json:"notificationResults" msgpack:"notificationResults"`                           // Key: 9
	LeaguePartyAndRankingResults     []LeaguePartyAndRankingResult     `json:"leaguePartyAndRankingResults" msgpack:"leaguePartyAndRankingResults"`         // Key: 10
	GachaHistoryResults              []GachaHistoryResult              `json:"gachaHistoryResults" msgpack:"gachaHistoryResults"`                           // Key: 11
	EpisodeDetailResults             []EpisodeDetailResult             `json:"episodeDetailResults" msgpack:"episodeDetailResults"`                         // Key: 12
	HelpDialog                       *HelpDialog                       `json:"helpDialog" msgpack:"helpDialog"`                                             // Key: 13
	BulkLevelUpPayloads              []BulkLevelUpPayload              `json:"bulkLevelUpPayloads" msgpack:"bulkLevelUpPayloads"`                           // Key: 14
	CharacterBaseStarPointResults    []CharacterBaseStarPointResult    `json:"characterBaseStarPointResults" msgpack:"characterBaseStarPointResults"`       // Key: 15
	GeneratePhotoResults             []GeneratePhotoResult             `json:"generatePhotoResults" msgpack:"generatePhotoResults"`                         // Key: 16
	TripleCastPartyAndRankingResults []TripleCastPartyAndRankingResult `json:"tripleCastPartyAndRankingResults" msgpack:"tripleCastPartyAndRankingResults"` // Key: 17
	CourseRankingResults             []CourseRankingResult             `json:"courseRankingResults" msgpack:"courseRankingResults"`                         // Key: 18
	ScoreWithDateResults             []ScoreWithDateResult             `json:"scoreWithDateResults" msgpack:"scoreWithDateResults"`                         // Key: 19
	MultiRoomInformationResults      []MultiRoomInformationResult      `json:"multiRoomInformationResults" msgpack:"multiRoomInformationResults"`           // Key: 20
	MultiRoomJoinnedResults          []MultiRoomJoinnedResult          `json:"multiRoomJoinnedResults" msgpack:"multiRoomJoinnedResults"`                   // Key: 21
	MultiRoomInvitedResults          []MultiRoomInvitedResult          `json:"multiRoomInvitedResults" msgpack:"multiRoomInvitedResults"`                   // Key: 22
}

type BanLevels int32

const (
	BanLevelsNormal  BanLevels = 0
	BanLevelsWarning BanLevels = 1
	BanLevelsSuspend BanLevels = 2
	BanLevelsDelete  BanLevels = 3
)

type GameVersions int32

const (
	GameVersionsUnknown    GameVersions = 0
	GameVersionsAppStore   GameVersions = 1
	GameVersionsGooglePlay GameVersions = 2
)

type ThingTypes int32

const (
	ThingTypesItem          ThingTypes = 1
	ThingTypesCharacter     ThingTypes = 2
	ThingTypesPoster        ThingTypes = 3
	ThingTypesAccessory     ThingTypes = 4
	ThingTypesCostume       ThingTypes = 5
	ThingTypesTrophy        ThingTypes = 6
	ThingTypesStamp         ThingTypes = 7
	ThingTypesNameplate     ThingTypes = 8
	ThingTypesNameColor     ThingTypes = 9
	ThingTypesBomb          ThingTypes = 10
	ThingTypesNote          ThingTypes = 11
	ThingTypesCoin          ThingTypes = 12
	ThingTypesJewel         ThingTypes = 13
	ThingTypesMusic         ThingTypes = 14
	ThingTypesStamina       ThingTypes = 15
	ThingTypesDecoration    ThingTypes = 16
	ThingTypesAlbumTheme    ThingTypes = 17
	ThingTypesNameBaseColor ThingTypes = 18
	ThingTypesIconFrame     ThingTypes = 19
	ThingTypesHomeSkin      ThingTypes = 20
)

type MissionCategories int32

const (
	MissionCategoriesBeginner       MissionCategories = 1
	MissionCategoriesNormal         MissionCategories = 2
	MissionCategoriesLimit          MissionCategories = 3
	MissionCategoriesDaily          MissionCategories = 4
	MissionCategoriesWeekly         MissionCategories = 5
	MissionCategoriesSiriusDay      MissionCategories = 6
	MissionCategoriesTheaterSirius  MissionCategories = 11
	MissionCategoriesTheaterEden    MissionCategories = 12
	MissionCategoriesTheaterGingaza MissionCategories = 13
	MissionCategoriesTheaterDenki   MissionCategories = 14
)

type CircleEventMissionTypes int32

const (
	CircleEventMissionTypesIndividual          CircleEventMissionTypes = 1
	CircleEventMissionTypesIndividualAndCircle CircleEventMissionTypes = 2
	CircleEventMissionTypesDaily               CircleEventMissionTypes = 3
)

type TabCategories int32

const (
	TabCategoriesHidden TabCategories = 0
	TabCategoriesNormal TabCategories = 1
	TabCategoriesRaise  TabCategories = 2
	TabCategoriesTicket TabCategories = 3
	TabCategoriesPiece  TabCategories = 4
)

type ItemCategories int32

const (
	ItemCategoriesStaminaRecovery                 ItemCategories = 11
	ItemCategoriesCharacterLevel                  ItemCategories = 12
	ItemCategoriesConsumeItem                     ItemCategories = 13
	ItemCategoriesTalentBloom                     ItemCategories = 14
	ItemCategoriesCharacterMission                ItemCategories = 16
	ItemCategoriesGachaTicket                     ItemCategories = 21
	ItemCategoriesExchangeTicket                  ItemCategories = 22
	ItemCategoriesRouletteBall                    ItemCategories = 31
	ItemCategoriesScratchCoin                     ItemCategories = 32
	ItemCategoriesGachaPoint                      ItemCategories = 41
	ItemCategoriesEventPoint                      ItemCategories = 42
	ItemCategoriesMissionPassPoint                ItemCategories = 43
	ItemCategoriesFilm                            ItemCategories = 51
	ItemCategoriesAlbumSkin                       ItemCategories = 52
	ItemCategoriesPlayerRankAwakening             ItemCategories = 61
	ItemCategoriesLessonPartyQuantityLimitRelease ItemCategories = 62
	ItemCategoriesBuffItem                        ItemCategories = 71
)

type JumpTypes int32

const (
	JumpTypesNone                   JumpTypes = 0
	JumpTypesInformation            JumpTypes = 1
	JumpTypesShop                   JumpTypes = 2
	JumpTypesJewelShop              JumpTypes = 3
	JumpTypesGacha                  JumpTypes = 4
	JumpTypesEvent                  JumpTypes = 5
	JumpTypesLive                   JumpTypes = 6
	JumpTypesMember                 JumpTypes = 7
	JumpTypesPhoto                  JumpTypes = 8
	JumpTypesStory                  JumpTypes = 9
	JumpTypesWebLink                JumpTypes = 10
	JumpTypesBeginnerMission        JumpTypes = 11
	JumpTypesMusicShop              JumpTypes = 12
	JumpTypesSpot                   JumpTypes = 13
	JumpTypesMission                JumpTypes = 14
	JumpTypesConcert                JumpTypes = 15
	JumpTypesStarRank               JumpTypes = 16
	JumpTypesEpisode                JumpTypes = 17
	JumpTypesSpecialEvent           JumpTypes = 18
	JumpTypesTotalPointEventRanking JumpTypes = 19
	JumpTypesBoxGacha               JumpTypes = 20
	JumpTypesTournamentEventTop     JumpTypes = 21
	JumpTypesLesson                 JumpTypes = 22
	JumpTypesLeague                 JumpTypes = 23
)

type StampType int32

const (
	StampTypeDefault   StampType = 1
	StampTypeAnimation StampType = 2
)

type TrophyCategories int32

const (
	TrophyCategoriesCharacter   TrophyCategories = 1
	TrophyCategoriesAchievement TrophyCategories = 2
	TrophyCategoriesEvent       TrophyCategories = 3
	TrophyCategoriesOther       TrophyCategories = 4
)

type DecorationCategories int32

const (
	DecorationCategoriesSpecial      DecorationCategories = 3
	DecorationCategoriesSeal         DecorationCategories = 20
	DecorationCategoriesSticker      DecorationCategories = 30
	DecorationCategoriesSpeachBubble DecorationCategories = 40
	DecorationCategoriesEffect       DecorationCategories = 50
	DecorationCategoriesFrame        DecorationCategories = 60
	DecorationCategoriesComedy       DecorationCategories = 70
	DecorationCategoriesText         DecorationCategories = 90
	DecorationCategoriesOther        DecorationCategories = 900
)

type NamePlateChangeTypes int32

const (
	NamePlateChangeTypesNone                         NamePlateChangeTypes = 0
	NamePlateChangeTypesDaiStarEnrollCount           NamePlateChangeTypes = 1
	NamePlateChangeTypesTripleCastDaiStarEnrollCount NamePlateChangeTypes = 2
)

type CharacterRarities int32

const (
	CharacterRaritiesRare1 CharacterRarities = 1
	CharacterRaritiesRare2 CharacterRarities = 2
	CharacterRaritiesRare3 CharacterRarities = 3
	CharacterRaritiesRare4 CharacterRarities = 4
)

type CharacterMissionCategoryTypes int32

const (
	CharacterMissionCategoryTypesLive        CharacterMissionCategoryTypes = 1
	CharacterMissionCategoryTypesEnhancement CharacterMissionCategoryTypes = 2
	CharacterMissionCategoryTypesTalentBloom CharacterMissionCategoryTypes = 3
	CharacterMissionCategoryTypesPoster      CharacterMissionCategoryTypes = 4
	CharacterMissionCategoryTypesOther       CharacterMissionCategoryTypes = 5
)

type CharacterMissionTypes int32

const (
	CharacterMissionTypesExecutedLeaderSenseCount     CharacterMissionTypes = 101
	CharacterMissionTypesEnhancedCharacter            CharacterMissionTypes = 102
	CharacterMissionTypesGetInGacha                   CharacterMissionTypes = 103
	CharacterMissionTypesTotalLevel                   CharacterMissionTypes = 104
	CharacterMissionTypesCharacterEpisodeCount        CharacterMissionTypes = 105
	CharacterMissionTypesCharacterAppearedPosterCount CharacterMissionTypes = 106
	CharacterMissionTypesReleasedPosterStoryCount     CharacterMissionTypes = 107
)

type PickupCharacterMissionCheckCondition int32

const (
	PickupCharacterMissionCheckConditionLevel          PickupCharacterMissionCheckCondition = 1
	PickupCharacterMissionCheckConditionAwakeningPhase PickupCharacterMissionCheckCondition = 2
	PickupCharacterMissionCheckConditionTalentStage    PickupCharacterMissionCheckCondition = 3
	PickupCharacterMissionCheckConditionSenseLevel     PickupCharacterMissionCheckCondition = 4
	PickupCharacterMissionCheckConditionEpisodeOrder   PickupCharacterMissionCheckCondition = 5
)

type PossessionRarities int32

const (
	PossessionRaritiesR   PossessionRarities = 1
	PossessionRaritiesSR  PossessionRarities = 2
	PossessionRaritiesSSR PossessionRarities = 3
)

type PossessionRarityFlag int32

const (
	PossessionRarityFlagNone PossessionRarityFlag = 0
	PossessionRarityFlagR    PossessionRarityFlag = 2
	PossessionRarityFlagSR   PossessionRarityFlag = 4
	PossessionRarityFlagSSR  PossessionRarityFlag = 8
)

type PhotoRarities int32

const (
	PhotoRaritiesRare1 PhotoRarities = 1
	PhotoRaritiesRare2 PhotoRarities = 2
	PhotoRaritiesRare3 PhotoRarities = 3
	PhotoRaritiesRare4 PhotoRarities = 4
	PhotoRaritiesRare5 PhotoRarities = 5
)

type EditTypes int32

const (
	EditTypesSimpleEdit EditTypes = 1
	EditTypesDetailEdit EditTypes = 2
)

type AlbumItemTypes int32

const (
	AlbumItemTypesPhoto AlbumItemTypes = 1
	AlbumItemTypesDeco  AlbumItemTypes = 2
	AlbumItemTypesTheme AlbumItemTypes = 3
	AlbumItemTypesStamp AlbumItemTypes = 4
)

type UseDecoPageFlag int32

const (
	UseDecoPageFlagNone   UseDecoPageFlag = 0
	UseDecoPageFlagPage1  UseDecoPageFlag = 1
	UseDecoPageFlagPage2  UseDecoPageFlag = 2
	UseDecoPageFlagPage3  UseDecoPageFlag = 4
	UseDecoPageFlagPage4  UseDecoPageFlag = 8
	UseDecoPageFlagPage5  UseDecoPageFlag = 16
	UseDecoPageFlagPage6  UseDecoPageFlag = 32
	UseDecoPageFlagPage7  UseDecoPageFlag = 64
	UseDecoPageFlagPage8  UseDecoPageFlag = 128
	UseDecoPageFlagPage9  UseDecoPageFlag = 256
	UseDecoPageFlagPage10 UseDecoPageFlag = 512
)

type MissionPassRewardStatus int32

const (
	MissionPassRewardStatusNotReceived     MissionPassRewardStatus = 1
	MissionPassRewardStatusReceived        MissionPassRewardStatus = 2
	MissionPassRewardStatusNotPaidSpItem   MissionPassRewardStatus = 3
	MissionPassRewardStatusNotReachedPoint MissionPassRewardStatus = 4
	MissionPassRewardStatusReceiveSuccess  MissionPassRewardStatus = 5
)

type MusicDifficulties int32

const (
	MusicDifficultiesNone    MusicDifficulties = 0
	MusicDifficultiesNormal  MusicDifficulties = 1
	MusicDifficultiesHard    MusicDifficulties = 2
	MusicDifficultiesExtra   MusicDifficulties = 3
	MusicDifficultiesStella  MusicDifficulties = 4
	MusicDifficultiesOlivier MusicDifficulties = 5
)

type Attributes int32

const (
	AttributesCute     Attributes = 1
	AttributesCool     Attributes = 2
	AttributesColorful Attributes = 3
	AttributesCheerful Attributes = 4
)

type CharacterProfileRestrictionItem int32

const (
	CharacterProfileRestrictionItemBirthday       CharacterProfileRestrictionItem = 1
	CharacterProfileRestrictionItemHeight         CharacterProfileRestrictionItem = 2
	CharacterProfileRestrictionItemSchool         CharacterProfileRestrictionItem = 3
	CharacterProfileRestrictionItemGrade          CharacterProfileRestrictionItem = 4
	CharacterProfileRestrictionItemSenseName      CharacterProfileRestrictionItem = 5
	CharacterProfileRestrictionItemSenseEffect    CharacterProfileRestrictionItem = 6
	CharacterProfileRestrictionItemEvoSenseName   CharacterProfileRestrictionItem = 7
	CharacterProfileRestrictionItemEvoSenseEffect CharacterProfileRestrictionItem = 8
	CharacterProfileRestrictionItemHobby          CharacterProfileRestrictionItem = 9
	CharacterProfileRestrictionItemBackground     CharacterProfileRestrictionItem = 10
	CharacterProfileRestrictionItemCharacterVoice CharacterProfileRestrictionItem = 11
	CharacterProfileRestrictionItemDescription    CharacterProfileRestrictionItem = 12
)

type TriggerType int32

const (
	TriggerTypeOverLife                       TriggerType = 1
	TriggerTypeBelowLife                      TriggerType = 2
	TriggerTypeCharacterBase                  TriggerType = 3
	TriggerTypeCompany                        TriggerType = 4
	TriggerTypeAttribute                      TriggerType = 5
	TriggerTypeSenseType                      TriggerType = 6
	TriggerTypeCompanyCount                   TriggerType = 7
	TriggerTypeAttributeCount                 TriggerType = 8
	TriggerTypeCharacterBaseGroup             TriggerType = 9
	TriggerTypeAllMemberBelongingCompany      TriggerType = 10
	TriggerTypeMaxMemberBelongingCompanyCount TriggerType = 11
)

type AchievementRateGrades int32

const (
	AchievementRateGradesNone   AchievementRateGrades = 0
	AchievementRateGradesC      AchievementRateGrades = 1
	AchievementRateGradesB      AchievementRateGrades = 2
	AchievementRateGradesA      AchievementRateGrades = 3
	AchievementRateGradesAPlus  AchievementRateGrades = 4
	AchievementRateGradesS      AchievementRateGrades = 5
	AchievementRateGradesSPlus  AchievementRateGrades = 6
	AchievementRateGradesSS     AchievementRateGrades = 7
	AchievementRateGradesSSPlus AchievementRateGrades = 8
	AchievementRateGradesSSS    AchievementRateGrades = 9
)

type ClearLamps int32

const (
	ClearLampsNone             ClearLamps = 0
	ClearLampsClear            ClearLamps = 1
	ClearLampsFullCombo        ClearLamps = 2
	ClearLampsFullComboMulti1  ClearLamps = 3
	ClearLampsFullComboMulti2  ClearLamps = 4
	ClearLampsFullComboMulti3  ClearLamps = 5
	ClearLampsAllPerfect       ClearLamps = 6
	ClearLampsAllPerfectMulti1 ClearLamps = 7
	ClearLampsAllPerfectMulti2 ClearLamps = 8
	ClearLampsAllPerfectMulti3 ClearLamps = 9
)

type StoryTypes int32

const (
	StoryTypesNone      StoryTypes = 0
	StoryTypesMain      StoryTypes = 1
	StoryTypesEvent     StoryTypes = 2
	StoryTypesSide      StoryTypes = 3
	StoryTypesCharacter StoryTypes = 4
	StoryTypesSpecial   StoryTypes = 5
)

type EpisodeReleaseConditionTypes int32

const (
	EpisodeReleaseConditionTypesTotalRankInCompany EpisodeReleaseConditionTypes = 1
	EpisodeReleaseConditionTypesCharacterRank      EpisodeReleaseConditionTypes = 2
)

type EventStoryListFlags int32

const (
	EventStoryListFlagsSirius           EventStoryListFlags = 1
	EventStoryListFlagsEden             EventStoryListFlags = 2
	EventStoryListFlagsGingaza          EventStoryListFlags = 4
	EventStoryListFlagsDenki            EventStoryListFlags = 8
	EventStoryListFlagsLoveLiveSunshine EventStoryListFlags = 512
	EventStoryListFlagsYuruCamp         EventStoryListFlags = 1024
	EventStoryListFlagsGirlsUndPanzer   EventStoryListFlags = 2048
)

type BonusAbilityEnableFlags int32

const (
	BonusAbilityEnableFlagsNone   BonusAbilityEnableFlags = 0
	BonusAbilityEnableFlagsFirst  BonusAbilityEnableFlags = 1
	BonusAbilityEnableFlagsSecond BonusAbilityEnableFlags = 2
	BonusAbilityEnableFlagsThird  BonusAbilityEnableFlags = 4
	BonusAbilityEnableFlagsFourth BonusAbilityEnableFlags = 8
	BonusAbilityEnableFlagsFifth  BonusAbilityEnableFlags = 16
	BonusAbilityEnableFlagsAll    BonusAbilityEnableFlags = 31
)

type PageCategories int32

const (
	PageCategoriesTutorialIngame                PageCategories = 1
	PageCategoriesSpot                          PageCategories = 2
	PageCategoriesPhoto                         PageCategories = 3
	PageCategoriesCircle                        PageCategories = 4
	PageCategoriesLesson                        PageCategories = 5
	PageCategoriesAuditionTop                   PageCategories = 6
	PageCategoriesAuditionConfirm               PageCategories = 7
	PageCategoriesLeague                        PageCategories = 8
	PageCategoriesSpecialLive                   PageCategories = 9
	PageCategoriesMultiLive                     PageCategories = 10
	PageCategoriesLiveResult                    PageCategories = 11
	PageCategoriesOlivierLiveResult             PageCategories = 12
	PageCategoriesParty                         PageCategories = 13
	PageCategoriesEnhancement                   PageCategories = 14
	PageCategoriesCostume                       PageCategories = 15
	PageCategoriesStarRank                      PageCategories = 16
	PageCategoriesSideStory                     PageCategories = 17
	PageCategoriesTrophyAndProfile              PageCategories = 18
	PageCategoriesShop                          PageCategories = 19
	PageCategoriesMedalExchangeShop             PageCategories = 20
	PageCategoriesBeginnerMission               PageCategories = 21
	PageCategoriesStaminaConsumption            PageCategories = 22
	PageCategoriesCompanyIntroMovieSirius       PageCategories = 24
	PageCategoriesCompanyIntroMovieEden         PageCategories = 25
	PageCategoriesCompanyIntroMovieGingaza      PageCategories = 26
	PageCategoriesCompanyIntroMovieDenki        PageCategories = 27
	PageCategoriesSpotStoryArchive              PageCategories = 28
	PageCategoriesPhotoTop                      PageCategories = 29
	PageCategoriesTeamChallenge                 PageCategories = 30
	PageCategoriesActorPortal                   PageCategories = 31
	PageCategoriesMusicSelection                PageCategories = 32
	PageCategoriesConcertStageSelect            PageCategories = 33
	PageCategoriesAlbumTop                      PageCategories = 34
	PageCategoriesCircleEvent                   PageCategories = 35
	PageCategoriesTournamentEventTop            PageCategories = 36
	PageCategoriesCourseModeSelection           PageCategories = 37
	PageCategoriesTripleCastLeague              PageCategories = 38
	PageCategoriesTournamentCourseModeSelection PageCategories = 39
	PageCategoriesGhostLive                     PageCategories = 40
	PageCategoriesTrialPartyEvent               PageCategories = 41
	PageCategoriesStellaConcours                PageCategories = 42
	PageCategoriesMultiRoom                     PageCategories = 43
)

type ViewPageTypes int32

const (
	ViewPageTypesShop ViewPageTypes = 1
)

type AuthenticationProviders int32

const (
	AuthenticationProvidersGoogle AuthenticationProviders = 1
	AuthenticationProvidersApple  AuthenticationProviders = 2
)

type AccountDeletionErrorTypes int32

const (
	AccountDeletionErrorTypesNone           AccountDeletionErrorTypes = 0
	AccountDeletionErrorTypesEnrolledCircle AccountDeletionErrorTypes = 1
)

type AccountRegisterErrorTypes int32

const (
	AccountRegisterErrorTypesNone                 AccountRegisterErrorTypes = 0
	AccountRegisterErrorTypesNgWord               AccountRegisterErrorTypes = 1
	AccountRegisterErrorTypesTooManyRegistrations AccountRegisterErrorTypes = 2
)

type HomeCharacterVoiceSetTypes int32

const (
	HomeCharacterVoiceSetTypesHome  HomeCharacterVoiceSetTypes = 1
	HomeCharacterVoiceSetTypesLive  HomeCharacterVoiceSetTypes = 2
	HomeCharacterVoiceSetTypesStory HomeCharacterVoiceSetTypes = 3
	HomeCharacterVoiceSetTypesParty HomeCharacterVoiceSetTypes = 4
	HomeCharacterVoiceSetTypesShop  HomeCharacterVoiceSetTypes = 5
)

type HomeCharacterVoicePlayPeriod int32

const (
	HomeCharacterVoicePlayPeriodPlayerBirthDay    HomeCharacterVoicePlayPeriod = 0
	HomeCharacterVoicePlayPeriodCharacterBirthDay HomeCharacterVoicePlayPeriod = 1
	HomeCharacterVoicePlayPeriodNewYear           HomeCharacterVoicePlayPeriod = 2
	HomeCharacterVoicePlayPeriodValentineDay      HomeCharacterVoicePlayPeriod = 3
	HomeCharacterVoicePlayPeriodGoldenWeek        HomeCharacterVoicePlayPeriod = 4
	HomeCharacterVoicePlayPeriodHalloWeen         HomeCharacterVoicePlayPeriod = 5
	HomeCharacterVoicePlayPeriodChristmas         HomeCharacterVoicePlayPeriod = 6
	HomeCharacterVoicePlayPeriodEndOfYear         HomeCharacterVoicePlayPeriod = 7
	HomeCharacterVoicePlayPeriodSeasonWinter      HomeCharacterVoicePlayPeriod = 8
	HomeCharacterVoicePlayPeriodSeasonSpring      HomeCharacterVoicePlayPeriod = 9
	HomeCharacterVoicePlayPeriodSeasonRainy       HomeCharacterVoicePlayPeriod = 10
	HomeCharacterVoicePlayPeriodSeasonSummer      HomeCharacterVoicePlayPeriod = 11
	HomeCharacterVoicePlayPeriodSeasonAutumn      HomeCharacterVoicePlayPeriod = 12
)

type HomeCharacterMotion int32

const (
	HomeCharacterMotionIdle     HomeCharacterMotion = 0
	HomeCharacterMotionMotion01 HomeCharacterMotion = 1
	HomeCharacterMotionMotion02 HomeCharacterMotion = 2
	HomeCharacterMotionMotion03 HomeCharacterMotion = 3
	HomeCharacterMotionMotion04 HomeCharacterMotion = 4
	HomeCharacterMotionMotion05 HomeCharacterMotion = 5
	HomeCharacterMotionMotion06 HomeCharacterMotion = 6
	HomeCharacterMotionMotion07 HomeCharacterMotion = 7
	HomeCharacterMotionMotion08 HomeCharacterMotion = 8
	HomeCharacterMotionMotion09 HomeCharacterMotion = 9
	HomeCharacterMotionMotion10 HomeCharacterMotion = 10
)

type BannerDeleteConditionTypes int32

const (
	BannerDeleteConditionTypesBuyItem            BannerDeleteConditionTypes = 1
	BannerDeleteConditionTypesMissionClear       BannerDeleteConditionTypes = 2
	BannerDeleteConditionTypesTimeLimitedControl BannerDeleteConditionTypes = 3
)

type WebLinkTypes int32

const (
	WebLinkTypesNormal  WebLinkTypes = 0
	WebLinkTypesWebShop WebLinkTypes = 1
)

type LoginBonusTypes int32

const (
	LoginBonusTypesNormal   LoginBonusTypes = 1
	LoginBonusTypesSpecial  LoginBonusTypes = 2
	LoginBonusTypesComeback LoginBonusTypes = 3
)

type LoginBonusLayoutTypes int32

const (
	LoginBonusLayoutTypesNormal  LoginBonusLayoutTypes = 1
	LoginBonusLayoutTypesSpecial LoginBonusLayoutTypes = 2
)

type LoginBonusSpineSelectType int32

const (
	LoginBonusSpineSelectTypeRandom               LoginBonusSpineSelectType = 1
	LoginBonusSpineSelectTypeHomeCharacterBaseId  LoginBonusSpineSelectType = 2
	LoginBonusSpineSelectTypeSelectedSpineCostume LoginBonusSpineSelectType = 3
)

type SplashTypes int32

const (
	SplashTypesVideo         SplashTypes = 1
	SplashTypesPicture       SplashTypes = 2
	SplashTypesEpisode       SplashTypes = 3
	SplashTypesLotteryResult SplashTypes = 4
)

type SplashAdditionalValueTypes int32

const (
	SplashAdditionalValueTypesNone                  SplashAdditionalValueTypes = 0
	SplashAdditionalValueTypesHomeCharacterBaseId   SplashAdditionalValueTypes = 1
	SplashAdditionalValueTypesLotteryHighestPrizeId SplashAdditionalValueTypes = 2
)

type SplashUnlockConditionTypes int32

const (
	SplashUnlockConditionTypesNone             SplashUnlockConditionTypes = 0
	SplashUnlockConditionTypesLotteryUserCache SplashUnlockConditionTypes = 1
)

type CampaignEffectTypes int32

const (
	CampaignEffectTypesLiveReward        CampaignEffectTypes = 1
	CampaignEffectTypesLessonReward      CampaignEffectTypes = 2
	CampaignEffectTypesStarPoint         CampaignEffectTypes = 3
	CampaignEffectTypesRankPoint         CampaignEffectTypes = 4
	CampaignEffectTypesAccessoryDropRate CampaignEffectTypes = 5
	CampaignEffectTypesLessonCount       CampaignEffectTypes = 6
	CampaignEffectTypesForDisplay        CampaignEffectTypes = 10
)

type HomeCharacterDisplayTypes int32

const (
	HomeCharacterDisplayTypesCharacterModel HomeCharacterDisplayTypes = 0
	HomeCharacterDisplayTypesIllust         HomeCharacterDisplayTypes = 1
)

type HomeBGMSelectionTypes int32

const (
	HomeBGMSelectionTypesUserSelect HomeBGMSelectionTypes = 1
	HomeBGMSelectionTypesRandom     HomeBGMSelectionTypes = 2
)

type HomeBGMPriorityTypes int32

const (
	HomeBGMPriorityTypesLow    HomeBGMPriorityTypes = -1
	HomeBGMPriorityTypesNormal HomeBGMPriorityTypes = 0
	HomeBGMPriorityTypesHigh   HomeBGMPriorityTypes = 1
)

type TitleBackgroundPriorityTypes int32

const (
	TitleBackgroundPriorityTypesLow    TitleBackgroundPriorityTypes = -1
	TitleBackgroundPriorityTypesNormal TitleBackgroundPriorityTypes = 0
	TitleBackgroundPriorityTypesHigh   TitleBackgroundPriorityTypes = 1
)

type TitleCallVoicePriorityTypes int32

const (
	TitleCallVoicePriorityTypesLow    TitleCallVoicePriorityTypes = -1
	TitleCallVoicePriorityTypesNormal TitleCallVoicePriorityTypes = 0
	TitleCallVoicePriorityTypesHigh   TitleCallVoicePriorityTypes = 1
)

type TitleBackgroundPathTypes int32

const (
	TitleBackgroundPathTypesPath              TitleBackgroundPathTypes = 1
	TitleBackgroundPathTypesCharacterMasterId TitleBackgroundPathTypes = 2
	TitleBackgroundPathTypesPosterMasterId    TitleBackgroundPathTypes = 3
)

type TitleDecorationTypes int32

const (
	TitleDecorationTypesFirstAnniversary      TitleDecorationTypes = 1
	TitleDecorationTypesSecondAnniversary     TitleDecorationTypes = 2
	TitleDecorationTypesSecondHalfAnniversary TitleDecorationTypes = 3
)

type SpotTypes int32

const (
	SpotTypesUtagawaHighSchool     SpotTypes = 1
	SpotTypesHigashiUenoHighSchool SpotTypes = 2
	SpotTypesPark                  SpotTypes = 3
	SpotTypesCafe                  SpotTypes = 4
	SpotTypesElectricTown          SpotTypes = 5
	SpotTypesThemePark             SpotTypes = 6
)

type SpotPositions int32

const (
	SpotPositionsPosition1 SpotPositions = 1
	SpotPositionsPosition2 SpotPositions = 2
	SpotPositionsPosition3 SpotPositions = 3
	SpotPositionsPosition4 SpotPositions = 4
)

type SenseTypes int32

const (
	SenseTypesSupport       SenseTypes = 1
	SenseTypesControl       SenseTypes = 2
	SenseTypesAmplification SenseTypes = 3
	SenseTypesSpecial       SenseTypes = 4
	SenseTypesNone          SenseTypes = 9
	SenseTypesAlternative   SenseTypes = 10
)

type EffectConditions int32

const (
	EffectConditionsCharacterBase      EffectConditions = 1
	EffectConditionsCompany            EffectConditions = 2
	EffectConditionsAttribute          EffectConditions = 3
	EffectConditionsSenseType          EffectConditions = 4
	EffectConditionsCharacter          EffectConditions = 5
	EffectConditionsEquippedPoster     EffectConditions = 6
	EffectConditionsNeighborPosition   EffectConditions = 7
	EffectConditionsCharacterBaseGroup EffectConditions = 8
)

type EffectTypes int32

const (
	EffectTypesBaseVocalUp                         EffectTypes = 1
	EffectTypesBaseExpressionUp                    EffectTypes = 2
	EffectTypesBaseConcentrationUp                 EffectTypes = 3
	EffectTypesBaseCorrection                      EffectTypes = 4
	EffectTypesVocalUp                             EffectTypes = 5
	EffectTypesExpressionUp                        EffectTypes = 6
	EffectTypesConcentrationUp                     EffectTypes = 7
	EffectTypesVocalLimitUp                        EffectTypes = 8
	EffectTypesExpressionLimitUp                   EffectTypes = 9
	EffectTypesConcentrationLimitUp                EffectTypes = 10
	EffectTypesPerformanceUp                       EffectTypes = 11
	EffectTypesFinalPerformanceUpCancelSense       EffectTypes = 12
	EffectTypesAddSenseLightSelf                   EffectTypes = 13
	EffectTypesAddSenseLightVariable               EffectTypes = 14
	EffectTypesAddSenseLightSupport                EffectTypes = 15
	EffectTypesAddSenseLightControl                EffectTypes = 16
	EffectTypesAddSenseLightAmplification          EffectTypes = 17
	EffectTypesAddSenseLightSpecial                EffectTypes = 18
	EffectTypesChangeWrongLightToSpLight           EffectTypes = 19
	EffectTypesSenseRecastDown                     EffectTypes = 20
	EffectTypesSenseCoolTimeRecastDown             EffectTypes = 21
	EffectTypesSenseAlternative                    EffectTypes = 22
	EffectTypesScoreUpByHighLife                   EffectTypes = 23
	EffectTypesScoreUpByLowLife                    EffectTypes = 24
	EffectTypesScoreUpByBuff                       EffectTypes = 25
	EffectTypesBuffTimeExtend                      EffectTypes = 26
	EffectTypesLifeHealing                         EffectTypes = 27
	EffectTypesLifeFixedValue                      EffectTypes = 28
	EffectTypesLifeGuard                           EffectTypes = 29
	EffectTypesSenseScoreUp                        EffectTypes = 30
	EffectTypesStarActScoreUp                      EffectTypes = 31
	EffectTypesBaseScoreUp                         EffectTypes = 32
	EffectTypesScoreGainOnScore                    EffectTypes = 33
	EffectTypesScoreGainOnVocal                    EffectTypes = 34
	EffectTypesScoreGainOnExpression               EffectTypes = 35
	EffectTypesScoreGainOnConcentration            EffectTypes = 36
	EffectTypesPrincipalGaugeGain                  EffectTypes = 37
	EffectTypesRewardUp                            EffectTypes = 38
	EffectTypesLightGuard                          EffectTypes = 39
	EffectTypesPrincipalGaugeUp                    EffectTypes = 40
	EffectTypesPrincipalGaugeLimitUp               EffectTypes = 41
	EffectTypesDecreaseRequireSupportLight         EffectTypes = 42
	EffectTypesDecreaseRequireControlLight         EffectTypes = 43
	EffectTypesDecreaseRequireAmplificationLight   EffectTypes = 44
	EffectTypesDecreaseRequireSpecialLight         EffectTypes = 45
	EffectTypesPerformanceLimitUp                  EffectTypes = 46
	EffectTypesScoreGainOnPerformance              EffectTypes = 47
	EffectTypesPrincipalGaugeBonus                 EffectTypes = 48
	EffectTypesPerformanceDuplicateUp              EffectTypes = 49
	EffectTypesCombinationSense                    EffectTypes = 50
	EffectTypesPlayerRankPointUp                   EffectTypes = 51
	EffectTypesPrincipalGaugeGainPercentageOfLimit EffectTypes = 52
	EffectTypesStarActProcrastinate                EffectTypes = 53
)

type EffectTargetRanges int32

const (
	EffectTargetRangesNone EffectTargetRanges = 0
	EffectTargetRangesSelf EffectTargetRanges = 1
	EffectTargetRangesAll  EffectTargetRanges = 2
)

type CalculationTypes int32

const (
	CalculationTypesPercentageAddition CalculationTypes = 1
	CalculationTypesMultiplication     CalculationTypes = 2
	CalculationTypesFixedAddition      CalculationTypes = 3
)

type EffectBuffTypes int32

const (
	EffectBuffTypesInstant    EffectBuffTypes = 1
	EffectBuffTypesContinuity EffectBuffTypes = 2
)

type SenseNotationBuffTypes int32

const (
	SenseNotationBuffTypesNone          SenseNotationBuffTypes = 0
	SenseNotationBuffTypesAttribute     SenseNotationBuffTypes = 1
	SenseNotationBuffTypesCompany       SenseNotationBuffTypes = 2
	SenseNotationBuffTypesCharacterBase SenseNotationBuffTypes = 3
	SenseNotationBuffTypesCharacter     SenseNotationBuffTypes = 4
)

type SenseNotationStatusTypes int32

const (
	SenseNotationStatusTypesPerformance   SenseNotationStatusTypes = 0
	SenseNotationStatusTypesVocal         SenseNotationStatusTypes = 1
	SenseNotationStatusTypesExpression    SenseNotationStatusTypes = 2
	SenseNotationStatusTypesConcentration SenseNotationStatusTypes = 3
)

type BranchJudgeTypes int32

const (
	BranchJudgeTypesNone     BranchJudgeTypes = 0
	BranchJudgeTypesEqual    BranchJudgeTypes = 1
	BranchJudgeTypesMoreThan BranchJudgeTypes = 2
	BranchJudgeTypesLessThan BranchJudgeTypes = 3
)

type FireTimingTypes int32

const (
	FireTimingTypesStarAct   FireTimingTypes = 1
	FireTimingTypesSense     FireTimingTypes = 2
	FireTimingTypesStartLive FireTimingTypes = 3
	FireTimingTypesPassive   FireTimingTypes = 4
)

type BranchConditionType int32

const (
	BranchConditionTypeNone                         BranchConditionType = 0
	BranchConditionTypeStarActFireCount             BranchConditionType = 1
	BranchConditionTypePartyPosition                BranchConditionType = 2
	BranchConditionTypeCompanyMemberCount           BranchConditionType = 3
	BranchConditionTypeCompanyCount                 BranchConditionType = 4
	BranchConditionTypeAttributeMemberCount         BranchConditionType = 5
	BranchConditionTypeAttributeCount               BranchConditionType = 6
	BranchConditionTypePosterAppearanceMemberCount  BranchConditionType = 7
	BranchConditionTypeLifeGuardCount               BranchConditionType = 8
	BranchConditionTypePosterAppearanceCompanyCount BranchConditionType = 9
	BranchConditionTypeStorageSenseLightCount       BranchConditionType = 10
	BranchConditionTypeCharacterBaseGroup           BranchConditionType = 11
	BranchConditionTypeSenseTriggeredCount          BranchConditionType = 12
)

type StorageSenseLightConditionValue int32

const (
	StorageSenseLightConditionValueSupport       StorageSenseLightConditionValue = 1
	StorageSenseLightConditionValueControl       StorageSenseLightConditionValue = 2
	StorageSenseLightConditionValueAmplification StorageSenseLightConditionValue = 3
	StorageSenseLightConditionValueSpecial       StorageSenseLightConditionValue = 4
	StorageSenseLightConditionValueAll           StorageSenseLightConditionValue = 5
)

type PosterEffectTypes int32

const (
	PosterEffectTypesLeader PosterEffectTypes = 1
	PosterEffectTypesNormal PosterEffectTypes = 2
)

type PosterAbilityTypes int32

const (
	PosterAbilityTypesLeader PosterAbilityTypes = 1
	PosterAbilityTypesNormal PosterAbilityTypes = 2
)

type PosterEpisodeTypes int32

const (
	PosterEpisodeTypesInformation PosterEpisodeTypes = 0
	PosterEpisodeTypesChapter1    PosterEpisodeTypes = 1
	PosterEpisodeTypesChapter2    PosterEpisodeTypes = 2
	PosterEpisodeTypesChapter3    PosterEpisodeTypes = 3
	PosterEpisodeTypesChapter4    PosterEpisodeTypes = 4
	PosterEpisodeTypesAfterTalk   PosterEpisodeTypes = 5
)

type PosterOrientation int32

const (
	PosterOrientationPortrait  PosterOrientation = 1
	PosterOrientationLandscape PosterOrientation = 2
)

type PosterSubTitleDisplayConditions int32

const (
	PosterSubTitleDisplayConditionsNone              PosterSubTitleDisplayConditions = 0
	PosterSubTitleDisplayConditionsBreakThroughPhase PosterSubTitleDisplayConditions = 1
)

type ResultVoiceConditionTypes int32

const (
	ResultVoiceConditionTypesAllPerfect                     ResultVoiceConditionTypes = 1
	ResultVoiceConditionTypesFullCombo                      ResultVoiceConditionTypes = 2
	ResultVoiceConditionTypesClearedAndBestAchievementRate  ResultVoiceConditionTypes = 3
	ResultVoiceConditionTypesCleared                        ResultVoiceConditionTypes = 4
	ResultVoiceConditionTypesLifeZeroAndBestAchievementRate ResultVoiceConditionTypes = 5
	ResultVoiceConditionTypesLifeZero                       ResultVoiceConditionTypes = 6
)

type FrameLotConditionTypes int32

const (
	FrameLotConditionTypesNone                              FrameLotConditionTypes = 0
	FrameLotConditionTypesConsumeStamina                    FrameLotConditionTypes = 1
	FrameLotConditionTypesScoreWithConsumeStamina           FrameLotConditionTypes = 2
	FrameLotConditionTypesStarActCountWithConsumeStamina    FrameLotConditionTypes = 3
	FrameLotConditionTypesAchievementRateWithConsumeStamina FrameLotConditionTypes = 4
)

type ProbabilityChangeTypes int32

const (
	ProbabilityChangeTypesNone             ProbabilityChangeTypes = 0
	ProbabilityChangeTypesAccessoryFormula ProbabilityChangeTypes = 1
)

type TimingTypes int32

const (
	TimingTypesNone         TimingTypes = 0
	TimingTypesMISS         TimingTypes = 1
	TimingTypesBAD          TimingTypes = 2
	TimingTypesGOOD         TimingTypes = 3
	TimingTypesGREAT        TimingTypes = 4
	TimingTypesPERFECT      TimingTypes = 5
	TimingTypesPERFECT_STAR TimingTypes = 6
)

type HighScoreTypes int32

const (
	HighScoreTypesNormal        HighScoreTypes = 0
	HighScoreTypesMulti         HighScoreTypes = 1
	HighScoreTypesTheaterLeague HighScoreTypes = 2
	HighScoreTypesAudition      HighScoreTypes = 3
)

type LiveTypes int32

const (
	LiveTypesNormal               LiveTypes = 1
	LiveTypesMulti                LiveTypes = 2
	LiveTypesLesson               LiveTypes = 3
	LiveTypesAudition             LiveTypes = 4
	LiveTypesLeague               LiveTypes = 5
	LiveTypesMultiCollectionEvent LiveTypes = 6
	LiveTypesConcert              LiveTypes = 7
	LiveTypesBonusLive            LiveTypes = 8
	LiveTypesCourseMode           LiveTypes = 9
	LiveTypesTripleCast           LiveTypes = 10
	LiveTypesGhostLive            LiveTypes = 11
	LiveTypesTrial                LiveTypes = 12
	LiveTypesMultiRoom            LiveTypes = 13
)

type SenseLightTypes int32

const (
	SenseLightTypesVariable      SenseLightTypes = 0
	SenseLightTypesSupport       SenseLightTypes = 1
	SenseLightTypesControl       SenseLightTypes = 2
	SenseLightTypesAmplification SenseLightTypes = 3
	SenseLightTypesSpecial       SenseLightTypes = 4
)

type OlivierReleaseStatuses int32

const (
	OlivierReleaseStatusesNone          OlivierReleaseStatuses = 0
	OlivierReleaseStatusesChallengeable OlivierReleaseStatuses = 1
	OlivierReleaseStatusesPurchasable   OlivierReleaseStatuses = 2
	OlivierReleaseStatusesReleased      OlivierReleaseStatuses = 3
)

type LiveDropTypes int32

const (
	LiveDropTypesNormal  LiveDropTypes = 0
	LiveDropTypesSpecial LiveDropTypes = 1
)

type MusicVideoTypes int32

const (
	MusicVideoTypesNone              MusicVideoTypes = 0
	MusicVideoTypesRealTimeRendering MusicVideoTypes = 1
	MusicVideoTypesMovie             MusicVideoTypes = 2
)

type MusicCoverTypes int32

const (
	MusicCoverTypesOriginal MusicCoverTypes = 1
	MusicCoverTypesCover    MusicCoverTypes = 2
)

type TeamChallengeDifficultyTypes int32

const (
	TeamChallengeDifficultyTypesEasy   TeamChallengeDifficultyTypes = 1
	TeamChallengeDifficultyTypesNormal TeamChallengeDifficultyTypes = 2
	TeamChallengeDifficultyTypesHard   TeamChallengeDifficultyTypes = 3
)

type TeamChallengeGoalTypes int32

const (
	TeamChallengeGoalTypesTotalScore   TeamChallengeGoalTypes = 1
	TeamChallengeGoalTypesTotalCombo   TeamChallengeGoalTypes = 2
	TeamChallengeGoalTypesTotalStarAct TeamChallengeGoalTypes = 3
)

type AnotherNotationTypes int32

const (
	AnotherNotationTypesShelved      AnotherNotationTypes = 1
	AnotherNotationTypesDefect       AnotherNotationTypes = 2
	AnotherNotationTypesChange       AnotherNotationTypes = 3
	AnotherNotationTypesMysterious   AnotherNotationTypes = 4
	AnotherNotationTypesMadness      AnotherNotationTypes = 5
	AnotherNotationTypesReproduction AnotherNotationTypes = 6
	AnotherNotationTypesAnother      AnotherNotationTypes = 7
	AnotherNotationTypesEasy         AnotherNotationTypes = 8
)

type MusicBookmarkFlags int32

const (
	MusicBookmarkFlagsNone      MusicBookmarkFlags = 0
	MusicBookmarkFlagsBookmark1 MusicBookmarkFlags = 1
	MusicBookmarkFlagsBookmark2 MusicBookmarkFlags = 2
	MusicBookmarkFlagsBookmark3 MusicBookmarkFlags = 4
)

type LiveScheduleEndDateDisplayType int32

const (
	LiveScheduleEndDateDisplayTypeNone              LiveScheduleEndDateDisplayType = 0
	LiveScheduleEndDateDisplayTypeEndDate           LiveScheduleEndDateDisplayType = 1
	LiveScheduleEndDateDisplayTypeAdditionalEndDate LiveScheduleEndDateDisplayType = 2
)

type BonusLiveUnlockConditionTypes int32

const (
	BonusLiveUnlockConditionTypesRouletteEventPoint BonusLiveUnlockConditionTypes = 1
)

type MusicCourseType int32

const (
	MusicCourseTypeDanClass             MusicCourseType = 1
	MusicCourseTypeAdvancedDanClass     MusicCourseType = 2
	MusicCourseTypeTournamentQualifying MusicCourseType = 3
	MusicCourseTypeSpecial              MusicCourseType = 4
	MusicCourseTypeWeeklyCourseRanking  MusicCourseType = 5
)

type MusicCourseGaugeType int32

const (
	MusicCourseGaugeTypeNormal MusicCourseGaugeType = 0
	MusicCourseGaugeTypeHot    MusicCourseGaugeType = 1
)

type MusicCourseUnlockConditionTypes int32

const (
	MusicCourseUnlockConditionTypesNone                  MusicCourseUnlockConditionTypes = 0
	MusicCourseUnlockConditionTypesClearedMMusicCourseId MusicCourseUnlockConditionTypes = 1
	MusicCourseUnlockConditionTypesPlayerRank            MusicCourseUnlockConditionTypes = 2
)

type MusicCourseCertificationGrade int32

const (
	MusicCourseCertificationGradeNone                   MusicCourseCertificationGrade = 0
	MusicCourseCertificationGradeFailed                 MusicCourseCertificationGrade = 1
	MusicCourseCertificationGradeNormalGaugeCertificate MusicCourseCertificationGrade = 2
	MusicCourseCertificationGradeHotGaugeCertificate    MusicCourseCertificationGrade = 3
)

type MatchingResult int32

const (
	MatchingResultWin  MatchingResult = 0
	MatchingResultLose MatchingResult = 1
	MatchingResultDraw MatchingResult = 2
)

type Companies int32

const (
	CompaniesNone             Companies = 0
	CompaniesSirius           Companies = 1
	CompaniesEden             Companies = 2
	CompaniesGingaza          Companies = 3
	CompaniesDenki            Companies = 4
	CompaniesLoveLiveSunshine Companies = 900
)

type CharacterEpisodeOrder int32

const (
	CharacterEpisodeOrderNone   CharacterEpisodeOrder = 0
	CharacterEpisodeOrderFirst  CharacterEpisodeOrder = 1
	CharacterEpisodeOrderSecond CharacterEpisodeOrder = 2
)

type BloomBonusTypes int32

const (
	BloomBonusTypesNone             BloomBonusTypes = 0
	BloomBonusTypesEffectDuringLive BloomBonusTypes = 1
	BloomBonusTypesStatusBonus      BloomBonusTypes = 2
	BloomBonusTypesGrantItem        BloomBonusTypes = 3
)

type LiveReleaseStatus int32

const (
	LiveReleaseStatusNone            LiveReleaseStatus = 0
	LiveReleaseStatusAvailableInShop LiveReleaseStatus = 1
	LiveReleaseStatusPlayable        LiveReleaseStatus = 2
)

type TalentBloomItemTypes int32

const (
	TalentBloomItemTypesActorPiece    TalentBloomItemTypes = 1
	TalentBloomItemTypesActorDaiPiece TalentBloomItemTypes = 2
)

type CharacterBaseTypes int32

const (
	CharacterBaseTypesInitial       CharacterBaseTypes = 1
	CharacterBaseTypesCollaboration CharacterBaseTypes = 900
)

type SenseFirePriority int32

const (
	SenseFirePriorityPrimary   SenseFirePriority = 1
	SenseFirePrioritySecondary SenseFirePriority = 2
)

type CharacterSelectionTypes int32

const (
	CharacterSelectionTypesPrimary   CharacterSelectionTypes = 1
	CharacterSelectionTypesSecondary CharacterSelectionTypes = 2
	CharacterSelectionTypesRandom    CharacterSelectionTypes = 99
)

type FriendSearchResultStatus int32

const (
	FriendSearchResultStatusFriend          FriendSearchResultStatus = 1
	FriendSearchResultStatusRequest         FriendSearchResultStatus = 2
	FriendSearchResultStatusNone            FriendSearchResultStatus = 3
	FriendSearchResultStatusUser            FriendSearchResultStatus = 4
	FriendSearchResultStatusReceivedRequest FriendSearchResultStatus = 5
)

type FriendAcceptResultStatus int32

const (
	FriendAcceptResultStatusAcceptSuccess               FriendAcceptResultStatus = 1
	FriendAcceptResultStatusRequestUserFriendsLimitOver FriendAcceptResultStatus = 2
	FriendAcceptResultStatusAcceptUserFriendsLimitOver  FriendAcceptResultStatus = 3
	FriendAcceptResultStatusCanceled                    FriendAcceptResultStatus = 4
)

type FriendRequestResultStatus int32

const (
	FriendRequestResultStatusRequestSuccess     FriendRequestResultStatus = 1
	FriendRequestResultStatusIsFriends          FriendRequestResultStatus = 2
	FriendRequestResultStatusIsApplying         FriendRequestResultStatus = 3
	FriendRequestResultStatusIsMySelf           FriendRequestResultStatus = 4
	FriendRequestResultStatusDataNotFound       FriendRequestResultStatus = 5
	FriendRequestResultStatusIsFriendCountLimit FriendRequestResultStatus = 6
)

type LeagueClassTypes int32

const (
	LeagueClassTypesNone    LeagueClassTypes = 0
	LeagueClassTypesRookie  LeagueClassTypes = 1
	LeagueClassTypesHope    LeagueClassTypes = 2
	LeagueClassTypesCast    LeagueClassTypes = 3
	LeagueClassTypesElite   LeagueClassTypes = 4
	LeagueClassTypesVeteran LeagueClassTypes = 5
	LeagueClassTypesStar    LeagueClassTypes = 6
	LeagueClassTypesDaiStar LeagueClassTypes = 7
)

type LeagueClassChangeTypes int32

const (
	LeagueClassChangeTypesNone LeagueClassChangeTypes = 0
	LeagueClassChangeTypesUp   LeagueClassChangeTypes = 1
	LeagueClassChangeTypesKeep LeagueClassChangeTypes = 2
	LeagueClassChangeTypesDown LeagueClassChangeTypes = 3
)

type FontSizes int32

const (
	FontSizesSmall  FontSizes = 1
	FontSizesMiddle FontSizes = 2
	FontSizesLarge  FontSizes = 3
)

type CharacterPositions int32

const (
	CharacterPositionsNone       CharacterPositions = 0
	CharacterPositionsOuterLeft  CharacterPositions = 1
	CharacterPositionsInnerLeft  CharacterPositions = 2
	CharacterPositionsCenter     CharacterPositions = 3
	CharacterPositionsInnerRight CharacterPositions = 4
	CharacterPositionsOuterRight CharacterPositions = 5
)

type CharacterAppearanceTypes int32

const (
	CharacterAppearanceTypesFadeIn            CharacterAppearanceTypes = 0
	CharacterAppearanceTypesSlideInFromRight  CharacterAppearanceTypes = 1
	CharacterAppearanceTypesSlideInFromLeft   CharacterAppearanceTypes = 2
	CharacterAppearanceTypesSlideInFromBottom CharacterAppearanceTypes = 3
)

type SpineSizes int32

const (
	SpineSizesSmall  SpineSizes = 1
	SpineSizesMiddle SpineSizes = 2
	SpineSizesLarge  SpineSizes = 3
)

type SpineBodySizeTypes int32

const (
	SpineBodySizeTypesSmall  SpineBodySizeTypes = 1
	SpineBodySizeTypesMedium SpineBodySizeTypes = 2
)

type FadeTypes int32

const (
	FadeTypesBlackFadeOutFadeIn FadeTypes = 1
	FadeTypesWhiteFadeOutFadeIn FadeTypes = 2
	FadeTypesTimeElapsed        FadeTypes = 3
	FadeTypesCrossFade          FadeTypes = 4
)

type WindowEffects int32

const (
	WindowEffectsSepia     WindowEffects = 1
	WindowEffectsWhiteBlur WindowEffects = 2
)

type CharacterLayerTypes int32

const (
	CharacterLayerTypesNone   CharacterLayerTypes = 0
	CharacterLayerTypesLayer1 CharacterLayerTypes = 1
	CharacterLayerTypesLayer2 CharacterLayerTypes = 2
	CharacterLayerTypesLayer3 CharacterLayerTypes = 3
)

type CircleAuthorities int32

const (
	CircleAuthoritiesNone         CircleAuthorities = 0
	CircleAuthoritiesMember       CircleAuthorities = 1
	CircleAuthoritiesDeputyLeader CircleAuthorities = 2
	CircleAuthoritiesLeader       CircleAuthorities = 3
)

type PlayTimeTypes int32

const (
	PlayTimeTypesNone        PlayTimeTypes = -1
	PlayTimeTypesZero        PlayTimeTypes = 0
	PlayTimeTypesOne         PlayTimeTypes = 1
	PlayTimeTypesTwo         PlayTimeTypes = 2
	PlayTimeTypesThree       PlayTimeTypes = 3
	PlayTimeTypesFour        PlayTimeTypes = 4
	PlayTimeTypesFive        PlayTimeTypes = 5
	PlayTimeTypesSix         PlayTimeTypes = 6
	PlayTimeTypesSeven       PlayTimeTypes = 7
	PlayTimeTypesEight       PlayTimeTypes = 8
	PlayTimeTypesNine        PlayTimeTypes = 9
	PlayTimeTypesTen         PlayTimeTypes = 10
	PlayTimeTypesEleven      PlayTimeTypes = 11
	PlayTimeTypesTwelve      PlayTimeTypes = 12
	PlayTimeTypesThirteen    PlayTimeTypes = 13
	PlayTimeTypesFourteen    PlayTimeTypes = 14
	PlayTimeTypesFifteen     PlayTimeTypes = 15
	PlayTimeTypesSixteen     PlayTimeTypes = 16
	PlayTimeTypesSeventeen   PlayTimeTypes = 17
	PlayTimeTypesEighteen    PlayTimeTypes = 18
	PlayTimeTypesNineteen    PlayTimeTypes = 19
	PlayTimeTypesTwenty      PlayTimeTypes = 20
	PlayTimeTypesTwentyOne   PlayTimeTypes = 21
	PlayTimeTypesTwentyTwo   PlayTimeTypes = 22
	PlayTimeTypesTwentyThree PlayTimeTypes = 23
	PlayTimeTypesTwentyFour  PlayTimeTypes = 24
	PlayTimeTypesTwentyFive  PlayTimeTypes = 25
	PlayTimeTypesTwentySix   PlayTimeTypes = 26
	PlayTimeTypesTwentySeven PlayTimeTypes = 27
	PlayTimeTypesTwentyEight PlayTimeTypes = 28
	PlayTimeTypesTwentyNine  PlayTimeTypes = 29
)

type EntryTypes int32

const (
	EntryTypesNone     EntryTypes = 0
	EntryTypesFree     EntryTypes = 1
	EntryTypesApproval EntryTypes = 2
)

type SearchCircleMemberConditionTypes int32

const (
	SearchCircleMemberConditionTypesNone      SearchCircleMemberConditionTypes = 0
	SearchCircleMemberConditionTypesOneToFive SearchCircleMemberConditionTypes = 1
	SearchCircleMemberConditionTypesSixToNine SearchCircleMemberConditionTypes = 2
)

type CircleAuthorityResultStatus int32

const (
	CircleAuthorityResultStatusPermissionDenied     CircleAuthorityResultStatus = 1
	CircleAuthorityResultStatusDeputyLeaderFullJoin CircleAuthorityResultStatus = 2
	CircleAuthorityResultStatusChangeSuccess        CircleAuthorityResultStatus = 3
	CircleAuthorityResultStatusLeaderChangeSuccess  CircleAuthorityResultStatus = 4
	CircleAuthorityResultStatusDataNotFound         CircleAuthorityResultStatus = 5
)

type CircleResultStatus int32

const (
	CircleResultStatusRequestIllegal         CircleResultStatus = 1
	CircleResultStatusLessThanNecessaryRank  CircleResultStatus = 2
	CircleResultStatusCanJoinCircle          CircleResultStatus = 3
	CircleResultStatusAlreadyJoinAnyCircle   CircleResultStatus = 4
	CircleResultStatusAlreadyInvite          CircleResultStatus = 5
	CircleResultStatusFreeEntry              CircleResultStatus = 6
	CircleResultStatusAlreadyRequest         CircleResultStatus = 7
	CircleResultStatusInviteSuccess          CircleResultStatus = 8
	CircleResultStatusPreRequestSuccess      CircleResultStatus = 9
	CircleResultStatusRequestSuccess         CircleResultStatus = 10
	CircleResultStatusApproveSuccess         CircleResultStatus = 11
	CircleResultStatusJoinSuccess            CircleResultStatus = 12
	CircleResultStatusEditSuccess            CircleResultStatus = 13
	CircleResultStatusCreateSuccess          CircleResultStatus = 14
	CircleResultStatusExpulsionSuccess       CircleResultStatus = 15
	CircleResultStatusResignationSuccess     CircleResultStatus = 16
	CircleResultStatusReleaseSuccess         CircleResultStatus = 17
	CircleResultStatusError                  CircleResultStatus = 18
	CircleResultStatusDataNotFound           CircleResultStatus = 19
	CircleResultStatusSearchSuccess          CircleResultStatus = 20
	CircleResultStatusCancelSuccess          CircleResultStatus = 21
	CircleResultStatusRejectSuccess          CircleResultStatus = 22
	CircleResultStatusNgWord                 CircleResultStatus = 23
	CircleResultStatusMemberAmountUpperLimit CircleResultStatus = 24
	CircleResultStatusInvalidParameter       CircleResultStatus = 25
	CircleResultStatusReleaseFailureByEvent  CircleResultStatus = 26
)

type ApiActivityLogTypes int32

const (
	ApiActivityLogTypesNone                     ApiActivityLogTypes = 0
	ApiActivityLogTypesJoinCircle               ApiActivityLogTypes = 1
	ApiActivityLogTypesLeaveCircle              ApiActivityLogTypes = 2
	ApiActivityLogTypesChangedPosition          ApiActivityLogTypes = 3
	ApiActivityLogTypesGachaAcquiredCharacter   ApiActivityLogTypes = 4
	ApiActivityLogTypesGachaAcquiredPoster      ApiActivityLogTypes = 5
	ApiActivityLogTypesPlayerRate               ApiActivityLogTypes = 6
	ApiActivityLogTypesFirstClearedAllPerfect   ApiActivityLogTypes = 7
	ApiActivityLogTypesFirstClearedFullCombo    ApiActivityLogTypes = 8
	ApiActivityLogTypesDonateSupportLevelLimit  ApiActivityLogTypes = 9
	ApiActivityLogTypesReleaseSupportLevelLimit ApiActivityLogTypes = 10
)

type CircleAuthorityUpdateTypes int32

const (
	CircleAuthorityUpdateTypesNone      CircleAuthorityUpdateTypes = 0
	CircleAuthorityUpdateTypesPromotion CircleAuthorityUpdateTypes = 1
	CircleAuthorityUpdateTypesDemotion  CircleAuthorityUpdateTypes = 2
)

type CircleDonateSupportCompanyResult int32

const (
	CircleDonateSupportCompanyResultSuccess                  CircleDonateSupportCompanyResult = 0
	CircleDonateSupportCompanyResultExceededQuantityReturned CircleDonateSupportCompanyResult = 1
	CircleDonateSupportCompanyResultError                    CircleDonateSupportCompanyResult = 99
)

type NotificationTabCategory int32

const (
	NotificationTabCategoryImportant         NotificationTabCategory = 1
	NotificationTabCategoryUpdateInformation NotificationTabCategory = 2
	NotificationTabCategoryBugInformation    NotificationTabCategory = 3
)

type NotificationCategory int32

const (
	NotificationCategoryNotification NotificationCategory = 1
	NotificationCategoryUpdate       NotificationCategory = 2
	NotificationCategoryCampaign     NotificationCategory = 3
	NotificationCategoryEvent        NotificationCategory = 4
	NotificationCategoryGacha        NotificationCategory = 5
	NotificationCategoryBug          NotificationCategory = 6
)

type ShopCategories int32

const (
	ShopCategoriesNone    ShopCategories = 0
	ShopCategoriesNormal  ShopCategories = 1
	ShopCategoriesMedal   ShopCategories = 2
	ShopCategoriesEvent   ShopCategories = 3
	ShopCategoriesGacha   ShopCategories = 4
	ShopCategoriesSpecial ShopCategories = 5
	ShopCategoriesMusic   ShopCategories = 6
)

type ShopReplaceTypes int32

const (
	ShopReplaceTypesNone             ShopReplaceTypes = 0
	ShopReplaceTypesDaily            ShopReplaceTypes = 1
	ShopReplaceTypesWeekly           ShopReplaceTypes = 2
	ShopReplaceTypesMonthly          ShopReplaceTypes = 3
	ShopReplaceTypesDailyPassExpired ShopReplaceTypes = 10
)

type ShopUnlockTypes int32

const (
	ShopUnlockTypesReadStory          ShopUnlockTypes = 2
	ShopUnlockTypesRateReached        ShopUnlockTypes = 8
	ShopUnlockTypesClassReached       ShopUnlockTypes = 9
	ShopUnlockTypesBuyItem            ShopUnlockTypes = 10
	ShopUnlockTypesReadFirstMainStory ShopUnlockTypes = 11
	ShopUnlockTypesAuditionReached    ShopUnlockTypes = 12
	ShopUnlockTypesSelectedPickup     ShopUnlockTypes = 17
	ShopUnlockTypesNotInPossession    ShopUnlockTypes = 20
)

type JewelShopUnlockTypes int32

const (
	JewelShopUnlockTypesReadStory          JewelShopUnlockTypes = 2
	JewelShopUnlockTypesRateReached        JewelShopUnlockTypes = 8
	JewelShopUnlockTypesClassReached       JewelShopUnlockTypes = 9
	JewelShopUnlockTypesBuyJewelShopItem   JewelShopUnlockTypes = 15
	JewelShopUnlockTypesPlayerRank         JewelShopUnlockTypes = 16
	JewelShopUnlockTypesTimeLimitedControl JewelShopUnlockTypes = 18
	JewelShopUnlockTypesFlashSale          JewelShopUnlockTypes = 19
)

type SaleTypes int32

const (
	SaleTypesAccountRegister SaleTypes = 1
)

type MusicUnlockConditionTypes int32

const (
	MusicUnlockConditionTypesDefault           MusicUnlockConditionTypes = 1
	MusicUnlockConditionTypesDistribute        MusicUnlockConditionTypes = 2
	MusicUnlockConditionTypesBuy               MusicUnlockConditionTypes = 10
	MusicUnlockConditionTypesReadEpisodeAndBuy MusicUnlockConditionTypes = 11
	MusicUnlockConditionTypesClearAudition     MusicUnlockConditionTypes = 12
)

type LiveUnlockConditionTypes int32

const (
	LiveUnlockConditionTypesNone           LiveUnlockConditionTypes = 0
	LiveUnlockConditionTypesExtraGoodCount LiveUnlockConditionTypes = 13
	LiveUnlockConditionTypesNotationShop   LiveUnlockConditionTypes = 14
)

type ShopPurchaseTypes int32

const (
	ShopPurchaseTypesPaid      ShopPurchaseTypes = 1
	ShopPurchaseTypesMoney     ShopPurchaseTypes = 2
	ShopPurchaseTypesFree      ShopPurchaseTypes = 3
	ShopPurchaseTypesFreeJewel ShopPurchaseTypes = 4
)

type PaymentStatus int32

const (
	PaymentStatusStarted                PaymentStatus = 1
	PaymentStatusConfirmed              PaymentStatus = 2
	PaymentStatusCompleted              PaymentStatus = 3
	PaymentStatusErrorNotFound          PaymentStatus = 4
	PaymentStatusErrorWrongStatus       PaymentStatus = 5
	PaymentStatusErrorRetryConfirmation PaymentStatus = 6
	PaymentStatusVoided                 PaymentStatus = 7
)

type ShopItemTypes int32

const (
	ShopItemTypesNone                 ShopItemTypes = 0
	ShopItemTypesLoginPass            ShopItemTypes = 1
	ShopItemTypesStarPass             ShopItemTypes = 2
	ShopItemTypesDaiStarPass          ShopItemTypes = 3
	ShopItemTypesMultiLiveDropBoost   ShopItemTypes = 4
	ShopItemTypesMultiLiveDropLimitUp ShopItemTypes = 5
)

type ViewedShopCategoryTypes int32

const (
	ViewedShopCategoryTypesExchangeShop ViewedShopCategoryTypes = 0
	ViewedShopCategoryTypesMusic        ViewedShopCategoryTypes = 1
	ViewedShopCategoryTypesLive         ViewedShopCategoryTypes = 2
	ViewedShopCategoryTypesGacha        ViewedShopCategoryTypes = 3
	ViewedShopCategoryTypesMarket       ViewedShopCategoryTypes = 4
)

type StarPassTypes int32

const (
	StarPassTypesStarPass    StarPassTypes = 1
	StarPassTypesDaiStarPass StarPassTypes = 2
)

type LoginPassNotificationTypes int32

const (
	LoginPassNotificationTypesNone        LoginPassNotificationTypes = 0
	LoginPassNotificationTypesPurchasable LoginPassNotificationTypes = 1
	LoginPassNotificationTypesInvalided   LoginPassNotificationTypes = 2
)

type MarketDisplayTypes int32

const (
	MarketDisplayTypesDefault   MarketDisplayTypes = 0
	MarketDisplayTypesRecommend MarketDisplayTypes = 1
	MarketDisplayTypesRare      MarketDisplayTypes = 2
)

type ConditionUsableTableFilter int32

const (
	ConditionUsableTableFilterDisplayRestriction ConditionUsableTableFilter = 1
	ConditionUsableTableFilterCostume            ConditionUsableTableFilter = 2
	ConditionUsableTableFilterJewelShop          ConditionUsableTableFilter = 4
	ConditionUsableTableFilterExchangeItem       ConditionUsableTableFilter = 8
	ConditionUsableTableFilterMusicShop          ConditionUsableTableFilter = 16
	ConditionUsableTableFilterMusic              ConditionUsableTableFilter = 32
)

type UnlockConditionTypes int32

const (
	UnlockConditionTypesCharacterStarRank   UnlockConditionTypes = 1
	UnlockConditionTypesStory               UnlockConditionTypes = 2
	UnlockConditionTypesGetPoster           UnlockConditionTypes = 3
	UnlockConditionTypesReleasePoster1      UnlockConditionTypes = 4
	UnlockConditionTypesReleasePoster2      UnlockConditionTypes = 5
	UnlockConditionTypesReleasePoster3      UnlockConditionTypes = 6
	UnlockConditionTypesReleasePoster4      UnlockConditionTypes = 7
	UnlockConditionTypesPlayerRate          UnlockConditionTypes = 8
	UnlockConditionTypesLeagueClass         UnlockConditionTypes = 9
	UnlockConditionTypesBuyItem             UnlockConditionTypes = 10
	UnlockConditionTypesAnyMainStoryChapter UnlockConditionTypes = 11
	UnlockConditionTypesAuditionClear       UnlockConditionTypes = 12
	UnlockConditionTypesExtraGoodCount      UnlockConditionTypes = 13
	UnlockConditionTypesNotationShop        UnlockConditionTypes = 14
	UnlockConditionTypesBuyJewelShopItem    UnlockConditionTypes = 15
	UnlockConditionTypesPlayerRank          UnlockConditionTypes = 16
	UnlockConditionTypesSelectedPickup      UnlockConditionTypes = 17
	UnlockConditionTypesTimeLimitedControl  UnlockConditionTypes = 18
	UnlockConditionTypesFlashSale           UnlockConditionTypes = 19
	UnlockConditionTypesNotInPossession     UnlockConditionTypes = 20
)

type GachaCardTypes int32

const (
	GachaCardTypesCharacter GachaCardTypes = 1
	GachaCardTypesPoster    GachaCardTypes = 2
)

type GachaTypes int32

const (
	GachaTypesPickup                 GachaTypes = 1
	GachaTypesSpecialTicket          GachaTypes = 2
	GachaTypesRare4Fixed             GachaTypes = 3
	GachaTypesNotPossessedRare4Fixed GachaTypes = 4
	GachaTypesSelection              GachaTypes = 5
	GachaTypesStepup                 GachaTypes = 6
	GachaTypesPickupSelection        GachaTypes = 7
	GachaTypesStationary             GachaTypes = 8
	GachaTypesReRoll                 GachaTypes = 9
)

type GachaEmissionFlags int32

const (
	GachaEmissionFlagsRare2 GachaEmissionFlags = 1
	GachaEmissionFlagsRare3 GachaEmissionFlags = 2
	GachaEmissionFlagsRare4 GachaEmissionFlags = 4
)

type GachaButtonTypes int32

const (
	GachaButtonTypesPaidJewelTenSeries    GachaButtonTypes = 1
	GachaButtonTypesFreeJewelTenSeries    GachaButtonTypes = 2
	GachaButtonTypesPaidJewelSingle       GachaButtonTypes = 3
	GachaButtonTypesFreeJewelSingle       GachaButtonTypes = 4
	GachaButtonTypesSingleTicket          GachaButtonTypes = 5
	GachaButtonTypesSingleTicketTenSeries GachaButtonTypes = 6
	GachaButtonTypesTenSeriesTicket       GachaButtonTypes = 7
)

type GachaUnlockTypes int32

const (
	GachaUnlockTypesNone               GachaUnlockTypes = 0
	GachaUnlockTypesTimeLimitedControl GachaUnlockTypes = 1
	GachaUnlockTypesGachaComplete      GachaUnlockTypes = 2
)

type GachaDisplayFooterConditions int32

const (
	GachaDisplayFooterConditionsNone           GachaDisplayFooterConditions = 0
	GachaDisplayFooterConditionsRemainRollLeft GachaDisplayFooterConditions = 1
)

type GachaGroupTypes int32

const (
	GachaGroupTypesNone   GachaGroupTypes = 0
	GachaGroupTypesPickup GachaGroupTypes = 1
	GachaGroupTypesActor  GachaGroupTypes = 2
	GachaGroupTypesPoster GachaGroupTypes = 3
)

type EventTypes int32

const (
	EventTypesStoryEvent         EventTypes = 20
	EventTypesSpecialEvent       EventTypes = 21
	EventTypesCharacterBasePoint EventTypes = 22
	EventTypesAnotherNotation    EventTypes = 23
	EventTypesCircleEvent        EventTypes = 24
	EventTypesRouletteEvent      EventTypes = 25
	EventTypesCampEvent          EventTypes = 26
	EventTypesConcert            EventTypes = 30
	EventTypesCourseMode         EventTypes = 31
	EventTypesDugongRun          EventTypes = 32
	EventTypesTrial              EventTypes = 33
	EventTypesConcours           EventTypes = 34
)

type StoryEventCategoryTypes int32

const (
	StoryEventCategoryTypesNormal    StoryEventCategoryTypes = 1
	StoryEventCategoryTypesVoltage   StoryEventCategoryTypes = 2
	StoryEventCategoryTypesHighScore StoryEventCategoryTypes = 3
	StoryEventCategoryTypesCircle    StoryEventCategoryTypes = 4
	StoryEventCategoryTypesMainStory StoryEventCategoryTypes = 5
)

type SpecialEventCategoryTypes int32

const (
	SpecialEventCategoryTypesTotalPoint SpecialEventCategoryTypes = 1
)

type StoryEventBonusTypes int32

const (
	StoryEventBonusTypesCompany       StoryEventBonusTypes = 1
	StoryEventBonusTypesCharacterBase StoryEventBonusTypes = 2
	StoryEventBonusTypesPoster        StoryEventBonusTypes = 3
)

type DugongRunDifficultyTypes int32

const (
	DugongRunDifficultyTypesEasy   DugongRunDifficultyTypes = 1
	DugongRunDifficultyTypesNormal DugongRunDifficultyTypes = 2
	DugongRunDifficultyTypesHard   DugongRunDifficultyTypes = 3
)

type DugongRunClearTypes int32

const (
	DugongRunClearTypesFailed     DugongRunClearTypes = 0
	DugongRunClearTypesClear      DugongRunClearTypes = 1
	DugongRunClearTypesNoMisstake DugongRunClearTypes = 2
)

type AsideLiveButtonTypes int32

const (
	AsideLiveButtonTypesNone               AsideLiveButtonTypes = 0
	AsideLiveButtonTypesTinyDreamFest      AsideLiveButtonTypes = 1
	AsideLiveButtonTypesConcertStageSelect AsideLiveButtonTypes = 2
	AsideLiveButtonTypesAnotherNotation    AsideLiveButtonTypes = 3
	AsideLiveButtonTypesTrial              AsideLiveButtonTypes = 4
)

type CampTypes int32

const (
	CampTypesNone  CampTypes = 0
	CampTypesCamp1 CampTypes = 1
	CampTypesCamp2 CampTypes = 2
)

type CustomLayoutUITypes int32

const (
	CustomLayoutUITypesNone                          CustomLayoutUITypes = 0
	CustomLayoutUITypesImage                         CustomLayoutUITypes = 1
	CustomLayoutUITypesImageButton                   CustomLayoutUITypes = 2
	CustomLayoutUITypesGrayOutImageButton            CustomLayoutUITypes = 3
	CustomLayoutUITypesButton                        CustomLayoutUITypes = 10
	CustomLayoutUITypesRedButton                     CustomLayoutUITypes = 11
	CustomLayoutUITypesGrayOutButton                 CustomLayoutUITypes = 12
	CustomLayoutUITypesGrayOutRedButton              CustomLayoutUITypes = 13
	CustomLayoutUITypesSpecialEventLogoWithDate      CustomLayoutUITypes = 20
	CustomLayoutUITypesItemCountPanel                CustomLayoutUITypes = 21
	CustomLayoutUITypesBgm                           CustomLayoutUITypes = 90
	CustomLayoutUITypesTotalPointEvent_ScorePanel    CustomLayoutUITypes = 101
	CustomLayoutUITypesTotalPointEvent_RankingButton CustomLayoutUITypes = 102
)

type CustomLayoutActionTypes int32

const (
	CustomLayoutActionTypesNone     CustomLayoutActionTypes = 0
	CustomLayoutActionTypesJumpType CustomLayoutActionTypes = 1
	CustomLayoutActionTypesHint     CustomLayoutActionTypes = 2
	CustomLayoutActionTypesToast    CustomLayoutActionTypes = 3
)

type TutorialStatus int32

const (
	TutorialStatusStart            TutorialStatus = 0
	TutorialStatusTutorialDownLoad TutorialStatus = 1
	TutorialStatusMainScenario     TutorialStatus = 2
	TutorialStatusTheaterMovie     TutorialStatus = 3
	TutorialStatusHome             TutorialStatus = 4
	TutorialStatusInGame           TutorialStatus = 5
	TutorialStatusMiniTalk         TutorialStatus = 6
	TutorialStatusFinish           TutorialStatus = 99
)

type RouletteTypes int32

const (
	RouletteTypesPoint RouletteTypes = 1
	RouletteTypesItem  RouletteTypes = 2
)

type TripleCastGroupOrder int32

const (
	TripleCastGroupOrderFirst  TripleCastGroupOrder = 0
	TripleCastGroupOrderSecond TripleCastGroupOrder = 1
	TripleCastGroupOrderThird  TripleCastGroupOrder = 2
)

type TrialPartyCharacterLockTypes int32

const (
	TrialPartyCharacterLockTypesNone          TrialPartyCharacterLockTypes = 0
	TrialPartyCharacterLockTypesCharacterBase TrialPartyCharacterLockTypes = 1
	TrialPartyCharacterLockTypesCharacter     TrialPartyCharacterLockTypes = 2
)

type TrialPartyEquipmentLockTypes int32

const (
	TrialPartyEquipmentLockTypesNone          TrialPartyEquipmentLockTypes = 0
	TrialPartyEquipmentLockTypesCharacterBase TrialPartyEquipmentLockTypes = 1
	TrialPartyEquipmentLockTypesPosition      TrialPartyEquipmentLockTypes = 2
)

type ProcessPaymentTransactionResult int32

const (
	ProcessPaymentTransactionResultSuccess                 ProcessPaymentTransactionResult = 1
	ProcessPaymentTransactionResultTemporaryIssuesTryAgain ProcessPaymentTransactionResult = 2
	ProcessPaymentTransactionResultFailed                  ProcessPaymentTransactionResult = 3
	ProcessPaymentTransactionResultCouldNotConfirm         ProcessPaymentTransactionResult = 4
	ProcessPaymentTransactionResultCouldNotAcknowledge     ProcessPaymentTransactionResult = 5
	ProcessPaymentTransactionResultPending                 ProcessPaymentTransactionResult = 6
)

type UseSerialCodeResult int32

const (
	UseSerialCodeResultSentToInbox                     UseSerialCodeResult = 1
	UseSerialCodeResultInvalidUserIdOrConfirmationCode UseSerialCodeResult = 2
	UseSerialCodeResultInvalidSerialCode               UseSerialCodeResult = 3
	UseSerialCodeResultAlreadyUsed                     UseSerialCodeResult = 4
	UseSerialCodeResultAlreadyUsedAnotherCodeInGroup   UseSerialCodeResult = 5
)

type FeatureMaintenanceScopes int32

const (
	FeatureMaintenanceScopesAudition   FeatureMaintenanceScopes = 1
	FeatureMaintenanceScopesLeague     FeatureMaintenanceScopes = 2
	FeatureMaintenanceScopesMultiLive  FeatureMaintenanceScopes = 3
	FeatureMaintenanceScopesInbox      FeatureMaintenanceScopes = 4
	FeatureMaintenanceScopesGacha      FeatureMaintenanceScopes = 5
	FeatureMaintenanceScopesShop       FeatureMaintenanceScopes = 6
	FeatureMaintenanceScopesFriends    FeatureMaintenanceScopes = 7
	FeatureMaintenanceScopesCircle     FeatureMaintenanceScopes = 8
	FeatureMaintenanceScopesCommonHub  FeatureMaintenanceScopes = 9
	FeatureMaintenanceScopesTripleCast FeatureMaintenanceScopes = 10
	FeatureMaintenanceScopesWebShop    FeatureMaintenanceScopes = 11
	FeatureMaintenanceScopesMultiRoom  FeatureMaintenanceScopes = 12
)

type GenericContentBodyTypes int32

const (
	GenericContentBodyTypesIntroduction        GenericContentBodyTypes = 1
	GenericContentBodyTypesMediumHeader        GenericContentBodyTypes = 2
	GenericContentBodyTypesSmallHeader         GenericContentBodyTypes = 3
	GenericContentBodyTypesBodyTextLeftAligned GenericContentBodyTypes = 4
	GenericContentBodyTypesBodyTextCentering   GenericContentBodyTypes = 5
	GenericContentBodyTypesBodyImage           GenericContentBodyTypes = 6
	GenericContentBodyTypesLinkText            GenericContentBodyTypes = 7
)

type InvitationCodeResultStatuses int32

const (
	InvitationCodeResultStatusesSuccess               InvitationCodeResultStatuses = 0
	InvitationCodeResultStatusesInvalidInvitationCode InvitationCodeResultStatuses = 1
	InvitationCodeResultStatusesOwnInvitationCode     InvitationCodeResultStatuses = 2
	InvitationCodeResultStatusesInviteUserNotFound    InvitationCodeResultStatuses = 3
)

type AlbumResetTypes int32

const (
	AlbumResetTypesRemovePhotosFromCurrentPage AlbumResetTypes = 0
	AlbumResetTypesClearCurrentPage            AlbumResetTypes = 1
	AlbumResetTypesRemovePhotosFromAllPages    AlbumResetTypes = 2
	AlbumResetTypesClearAllPages               AlbumResetTypes = 3
)

type MultiRoomPlayModes int32

const (
	MultiRoomPlayModesAchievementRate MultiRoomPlayModes = 0
	MultiRoomPlayModesScore           MultiRoomPlayModes = 1
)

type Fault struct {
	ErrorCode  string `json:"errorCode" msgpack:"errorCode"`   // Key: 0
	Message    string `json:"message" msgpack:"message"`       // Key: 1
	StackTrace string `json:"stackTrace" msgpack:"stackTrace"` // Key: 2
}

type MemoryDatabase struct {
	AccessoryEffectFilterMasterTable                                *AccessoryEffectFilterMasterTable                                `json:"accessoryEffectFilterMasterTable" msgpack:"accessoryEffectFilterMasterTable"`
	AccessoryEffectMasterTable                                      *AccessoryEffectMasterTable                                      `json:"accessoryEffectMasterTable" msgpack:"accessoryEffectMasterTable"`
	AccessoryLevelPatternGroupMasterTable                           *AccessoryLevelPatternGroupMasterTable                           `json:"accessoryLevelPatternGroupMasterTable" msgpack:"accessoryLevelPatternGroupMasterTable"`
	AccessoryMasterTable                                            *AccessoryMasterTable                                            `json:"accessoryMasterTable" msgpack:"accessoryMasterTable"`
	ActivityLogMessageTemplateMasterTable                           *ActivityLogMessageTemplateMasterTable                           `json:"activityLogMessageTemplateMasterTable" msgpack:"activityLogMessageTemplateMasterTable"`
	AdditionalRewardPackageMasterTable                              *AdditionalRewardPackageMasterTable                              `json:"additionalRewardPackageMasterTable" msgpack:"additionalRewardPackageMasterTable"`
	AlbumEffectMasterTable                                          *AlbumEffectMasterTable                                          `json:"albumEffectMasterTable" msgpack:"albumEffectMasterTable"`
	AlbumThemeMasterTable                                           *AlbumThemeMasterTable                                           `json:"albumThemeMasterTable" msgpack:"albumThemeMasterTable"`
	AnotherNotationMasterTable                                      *AnotherNotationMasterTable                                      `json:"anotherNotationMasterTable" msgpack:"anotherNotationMasterTable"`
	AuditionMasterTable                                             *AuditionMasterTable                                             `json:"auditionMasterTable" msgpack:"auditionMasterTable"`
	AuditionPhaseMasterTable                                        *AuditionPhaseMasterTable                                        `json:"auditionPhaseMasterTable" msgpack:"auditionPhaseMasterTable"`
	AuditionRewardPackageMasterTable                                *AuditionRewardPackageMasterTable                                `json:"auditionRewardPackageMasterTable" msgpack:"auditionRewardPackageMasterTable"`
	BannerMasterTable                                               *BannerMasterTable                                               `json:"bannerMasterTable" msgpack:"bannerMasterTable"`
	BodyMotionMasterTable                                           *BodyMotionMasterTable                                           `json:"bodyMotionMasterTable" msgpack:"bodyMotionMasterTable"`
	BombMasterTable                                                 *BombMasterTable                                                 `json:"bombMasterTable" msgpack:"bombMasterTable"`
	BonusLiveMasterTable                                            *BonusLiveMasterTable                                            `json:"bonusLiveMasterTable" msgpack:"bonusLiveMasterTable"`
	BonusLiveStageMasterTable                                       *BonusLiveStageMasterTable                                       `json:"bonusLiveStageMasterTable" msgpack:"bonusLiveStageMasterTable"`
	BuffItemMasterTable                                             *BuffItemMasterTable                                             `json:"buffItemMasterTable" msgpack:"buffItemMasterTable"`
	CampaignMasterTable                                             *CampaignMasterTable                                             `json:"campaignMasterTable" msgpack:"campaignMasterTable"`
	CategoryGroupMasterTable                                        *CategoryGroupMasterTable                                        `json:"categoryGroupMasterTable" msgpack:"categoryGroupMasterTable"`
	CategoryMasterTable                                             *CategoryMasterTable                                             `json:"categoryMasterTable" msgpack:"categoryMasterTable"`
	ChangeBodyMotionMasterTable                                     *ChangeBodyMotionMasterTable                                     `json:"changeBodyMotionMasterTable" msgpack:"changeBodyMotionMasterTable"`
	CharacterAwakeningItemGroupMasterTable                          *CharacterAwakeningItemGroupMasterTable                          `json:"characterAwakeningItemGroupMasterTable" msgpack:"characterAwakeningItemGroupMasterTable"`
	CharacterAwakeningItemMasterTable                               *CharacterAwakeningItemMasterTable                               `json:"characterAwakeningItemMasterTable" msgpack:"characterAwakeningItemMasterTable"`
	CharacterBaseBloomGenericItemMasterTable                        *CharacterBaseBloomGenericItemMasterTable                        `json:"characterBaseBloomGenericItemMasterTable" msgpack:"characterBaseBloomGenericItemMasterTable"`
	CharacterBaseMasterTable                                        *CharacterBaseMasterTable                                        `json:"characterBaseMasterTable" msgpack:"characterBaseMasterTable"`
	CharacterBloomBonusGroupMasterTable                             *CharacterBloomBonusGroupMasterTable                             `json:"characterBloomBonusGroupMasterTable" msgpack:"characterBloomBonusGroupMasterTable"`
	CharacterBloomDetailMasterTable                                 *CharacterBloomDetailMasterTable                                 `json:"characterBloomDetailMasterTable" msgpack:"characterBloomDetailMasterTable"`
	CharacterBloomItemMasterTable                                   *CharacterBloomItemMasterTable                                   `json:"characterBloomItemMasterTable" msgpack:"characterBloomItemMasterTable"`
	CharacterEpisodeMasterTable                                     *CharacterEpisodeMasterTable                                     `json:"characterEpisodeMasterTable" msgpack:"characterEpisodeMasterTable"`
	CharacterEpisodeRelationMasterTable                             *CharacterEpisodeRelationMasterTable                             `json:"characterEpisodeRelationMasterTable" msgpack:"characterEpisodeRelationMasterTable"`
	CharacterEpisodeReleaseItemGroupMasterTable                     *CharacterEpisodeReleaseItemGroupMasterTable                     `json:"characterEpisodeReleaseItemGroupMasterTable" msgpack:"characterEpisodeReleaseItemGroupMasterTable"`
	CharacterExperienceItemMasterTable                              *CharacterExperienceItemMasterTable                              `json:"characterExperienceItemMasterTable" msgpack:"characterExperienceItemMasterTable"`
	CharacterKeyMissionMasterTable                                  *CharacterKeyMissionMasterTable                                  `json:"characterKeyMissionMasterTable" msgpack:"characterKeyMissionMasterTable"`
	CharacterLessonScoreRewardMasterTable                           *CharacterLessonScoreRewardMasterTable                           `json:"characterLessonScoreRewardMasterTable" msgpack:"characterLessonScoreRewardMasterTable"`
	CharacterLevelMasterTable                                       *CharacterLevelMasterTable                                       `json:"characterLevelMasterTable" msgpack:"characterLevelMasterTable"`
	CharacterMasterTable                                            *CharacterMasterTable                                            `json:"characterMasterTable" msgpack:"characterMasterTable"`
	CharacterMissionCategoryLevelMasterTable                        *CharacterMissionCategoryLevelMasterTable                        `json:"characterMissionCategoryLevelMasterTable" msgpack:"characterMissionCategoryLevelMasterTable"`
	CharacterMissionItemMasterTable                                 *CharacterMissionItemMasterTable                                 `json:"characterMissionItemMasterTable" msgpack:"characterMissionItemMasterTable"`
	CharacterMissionMasterTable                                     *CharacterMissionMasterTable                                     `json:"characterMissionMasterTable" msgpack:"characterMissionMasterTable"`
	CharacterMissionStageMasterTable                                *CharacterMissionStageMasterTable                                `json:"characterMissionStageMasterTable" msgpack:"characterMissionStageMasterTable"`
	CharacterPieceMasterTable                                       *CharacterPieceMasterTable                                       `json:"characterPieceMasterTable" msgpack:"characterPieceMasterTable"`
	CharacterPointEventCharacterRankingRewardItemPackageMasterTable *CharacterPointEventCharacterRankingRewardItemPackageMasterTable `json:"characterPointEventCharacterRankingRewardItemPackageMasterTable" msgpack:"characterPointEventCharacterRankingRewardItemPackageMasterTable"`
	CharacterPointEventCharacterRankingRewardMasterTable            *CharacterPointEventCharacterRankingRewardMasterTable            `json:"characterPointEventCharacterRankingRewardMasterTable" msgpack:"characterPointEventCharacterRankingRewardMasterTable"`
	CharacterPointEventMasterTable                                  *CharacterPointEventMasterTable                                  `json:"characterPointEventMasterTable" msgpack:"characterPointEventMasterTable"`
	CharacterProfileRestrictionMasterTable                          *CharacterProfileRestrictionMasterTable                          `json:"characterProfileRestrictionMasterTable" msgpack:"characterProfileRestrictionMasterTable"`
	CharacterSenseEnhanceItemGroupMasterTable                       *CharacterSenseEnhanceItemGroupMasterTable                       `json:"characterSenseEnhanceItemGroupMasterTable" msgpack:"characterSenseEnhanceItemGroupMasterTable"`
	CharacterStarRankMasterTable                                    *CharacterStarRankMasterTable                                    `json:"characterStarRankMasterTable" msgpack:"characterStarRankMasterTable"`
	CharacterStarRankRewardGroupMasterTable                         *CharacterStarRankRewardGroupMasterTable                         `json:"characterStarRankRewardGroupMasterTable" msgpack:"characterStarRankRewardGroupMasterTable"`
	CircleEventCirclePointRewardMasterTable                         *CircleEventCirclePointRewardMasterTable                         `json:"circleEventCirclePointRewardMasterTable" msgpack:"circleEventCirclePointRewardMasterTable"`
	CircleEventCircleRankingRewardItemPackageMasterTable            *CircleEventCircleRankingRewardItemPackageMasterTable            `json:"circleEventCircleRankingRewardItemPackageMasterTable" msgpack:"circleEventCircleRankingRewardItemPackageMasterTable"`
	CircleEventCircleRankingRewardMasterTable                       *CircleEventCircleRankingRewardMasterTable                       `json:"circleEventCircleRankingRewardMasterTable" msgpack:"circleEventCircleRankingRewardMasterTable"`
	CircleEventMasterTable                                          *CircleEventMasterTable                                          `json:"circleEventMasterTable" msgpack:"circleEventMasterTable"`
	CircleEventMissionMasterTable                                   *CircleEventMissionMasterTable                                   `json:"circleEventMissionMasterTable" msgpack:"circleEventMissionMasterTable"`
	CircleEventMissionRefreshSettingGroupMasterTable                *CircleEventMissionRefreshSettingGroupMasterTable                `json:"circleEventMissionRefreshSettingGroupMasterTable" msgpack:"circleEventMissionRefreshSettingGroupMasterTable"`
	CircleSupportCompanyLevelDetailMasterTable                      *CircleSupportCompanyLevelDetailMasterTable                      `json:"circleSupportCompanyLevelDetailMasterTable" msgpack:"circleSupportCompanyLevelDetailMasterTable"`
	CircleSupportCompanyLevelLimitMasterTable                       *CircleSupportCompanyLevelLimitMasterTable                       `json:"circleSupportCompanyLevelLimitMasterTable" msgpack:"circleSupportCompanyLevelLimitMasterTable"`
	CircleTheaterLevelMasterTable                                   *CircleTheaterLevelMasterTable                                   `json:"circleTheaterLevelMasterTable" msgpack:"circleTheaterLevelMasterTable"`
	ComebackCampaignMasterTable                                     *ComebackCampaignMasterTable                                     `json:"comebackCampaignMasterTable" msgpack:"comebackCampaignMasterTable"`
	ComicMasterTable                                                *ComicMasterTable                                                `json:"comicMasterTable" msgpack:"comicMasterTable"`
	CompanyMasterTable                                              *CompanyMasterTable                                              `json:"companyMasterTable" msgpack:"companyMasterTable"`
	ConcertMasterTable                                              *ConcertMasterTable                                              `json:"concertMasterTable" msgpack:"concertMasterTable"`
	ConcertStageMasterTable                                         *ConcertStageMasterTable                                         `json:"concertStageMasterTable" msgpack:"concertStageMasterTable"`
	ConcoursMasterTable                                             *ConcoursMasterTable                                             `json:"concoursMasterTable" msgpack:"concoursMasterTable"`
	ConcoursPointRewardMasterTable                                  *ConcoursPointRewardMasterTable                                  `json:"concoursPointRewardMasterTable" msgpack:"concoursPointRewardMasterTable"`
	CostumeGroupMasterTable                                         *CostumeGroupMasterTable                                         `json:"costumeGroupMasterTable" msgpack:"costumeGroupMasterTable"`
	CostumeMasterTable                                              *CostumeMasterTable                                              `json:"costumeMasterTable" msgpack:"costumeMasterTable"`
	CostumeWearableCharacterGroupMasterTable                        *CostumeWearableCharacterGroupMasterTable                        `json:"costumeWearableCharacterGroupMasterTable" msgpack:"costumeWearableCharacterGroupMasterTable"`
	CourseRankingRewardMasterTable                                  *CourseRankingRewardMasterTable                                  `json:"courseRankingRewardMasterTable" msgpack:"courseRankingRewardMasterTable"`
	DecorationMasterTable                                           *DecorationMasterTable                                           `json:"decorationMasterTable" msgpack:"decorationMasterTable"`
	DugongRunCourseMasterTable                                      *DugongRunCourseMasterTable                                      `json:"dugongRunCourseMasterTable" msgpack:"dugongRunCourseMasterTable"`
	DugongRunRewardMasterTable                                      *DugongRunRewardMasterTable                                      `json:"dugongRunRewardMasterTable" msgpack:"dugongRunRewardMasterTable"`
	EffectDurationGroupMasterTable                                  *EffectDurationGroupMasterTable                                  `json:"effectDurationGroupMasterTable" msgpack:"effectDurationGroupMasterTable"`
	EffectMasterTable                                               *EffectMasterTable                                               `json:"effectMasterTable" msgpack:"effectMasterTable"`
	EffectTriggerCharacterBaseGroupMasterTable                      *EffectTriggerCharacterBaseGroupMasterTable                      `json:"effectTriggerCharacterBaseGroupMasterTable" msgpack:"effectTriggerCharacterBaseGroupMasterTable"`
	EpisodeMasterTable                                              *EpisodeMasterTable                                              `json:"episodeMasterTable" msgpack:"episodeMasterTable"`
	EpisodeRewardPackageMasterTable                                 *EpisodeRewardPackageMasterTable                                 `json:"episodeRewardPackageMasterTable" msgpack:"episodeRewardPackageMasterTable"`
	EventBoxGachaBoxMasterTable                                     *EventBoxGachaBoxMasterTable                                     `json:"eventBoxGachaBoxMasterTable" msgpack:"eventBoxGachaBoxMasterTable"`
	EventBoxGachaMasterTable                                        *EventBoxGachaMasterTable                                        `json:"eventBoxGachaMasterTable" msgpack:"eventBoxGachaMasterTable"`
	EventBoxGachaTextTemplateMasterTable                            *EventBoxGachaTextTemplateMasterTable                            `json:"eventBoxGachaTextTemplateMasterTable" msgpack:"eventBoxGachaTextTemplateMasterTable"`
	EventCampClassMasterTable                                       *EventCampClassMasterTable                                       `json:"eventCampClassMasterTable" msgpack:"eventCampClassMasterTable"`
	EventCampMasterTable                                            *EventCampMasterTable                                            `json:"eventCampMasterTable" msgpack:"eventCampMasterTable"`
	EventCampSupportPointRewardMasterTable                          *EventCampSupportPointRewardMasterTable                          `json:"eventCampSupportPointRewardMasterTable" msgpack:"eventCampSupportPointRewardMasterTable"`
	EventMasterTable                                                *EventMasterTable                                                `json:"eventMasterTable" msgpack:"eventMasterTable"`
	ExchangeShopMasterTable                                         *ExchangeShopMasterTable                                         `json:"exchangeShopMasterTable" msgpack:"exchangeShopMasterTable"`
	FacialExpressionMasterTable                                     *FacialExpressionMasterTable                                     `json:"facialExpressionMasterTable" msgpack:"facialExpressionMasterTable"`
	FilmItemMasterTable                                             *FilmItemMasterTable                                             `json:"filmItemMasterTable" msgpack:"filmItemMasterTable"`
	FriendInvitationMissionMasterTable                              *FriendInvitationMissionMasterTable                              `json:"friendInvitationMissionMasterTable" msgpack:"friendInvitationMissionMasterTable"`
	GachaMasterTable                                                *GachaMasterTable                                                `json:"gachaMasterTable" msgpack:"gachaMasterTable"`
	GachaTextTemplateMasterTable                                    *GachaTextTemplateMasterTable                                    `json:"gachaTextTemplateMasterTable" msgpack:"gachaTextTemplateMasterTable"`
	GameHintMasterTable                                             *GameHintMasterTable                                             `json:"gameHintMasterTable" msgpack:"gameHintMasterTable"`
	GhostLiveMasterTable                                            *GhostLiveMasterTable                                            `json:"ghostLiveMasterTable" msgpack:"ghostLiveMasterTable"`
	GradualMissionGroupMasterTable                                  *GradualMissionGroupMasterTable                                  `json:"gradualMissionGroupMasterTable" msgpack:"gradualMissionGroupMasterTable"`
	HeadDirectionMasterTable                                        *HeadDirectionMasterTable                                        `json:"headDirectionMasterTable" msgpack:"headDirectionMasterTable"`
	HeadMotionMasterTable                                           *HeadMotionMasterTable                                           `json:"headMotionMasterTable" msgpack:"headMotionMasterTable"`
	HomeBGMDetailMasterTable                                        *HomeBGMDetailMasterTable                                        `json:"homeBGMDetailMasterTable" msgpack:"homeBGMDetailMasterTable"`
	HomeBGMMasterTable                                              *HomeBGMMasterTable                                              `json:"homeBGMMasterTable" msgpack:"homeBGMMasterTable"`
	HomeBackgroundMasterTable                                       *HomeBackgroundMasterTable                                       `json:"homeBackgroundMasterTable" msgpack:"homeBackgroundMasterTable"`
	HomeCharacterMasterTable                                        *HomeCharacterMasterTable                                        `json:"homeCharacterMasterTable" msgpack:"homeCharacterMasterTable"`
	HomeCharacterVoiceMasterTable                                   *HomeCharacterVoiceMasterTable                                   `json:"homeCharacterVoiceMasterTable" msgpack:"homeCharacterVoiceMasterTable"`
	HomeCharacterVoicePeriodMasterTable                             *HomeCharacterVoicePeriodMasterTable                             `json:"homeCharacterVoicePeriodMasterTable" msgpack:"homeCharacterVoicePeriodMasterTable"`
	HomePosterMasterTable                                           *HomePosterMasterTable                                           `json:"homePosterMasterTable" msgpack:"homePosterMasterTable"`
	HomeSkinMasterTable                                             *HomeSkinMasterTable                                             `json:"homeSkinMasterTable" msgpack:"homeSkinMasterTable"`
	IconFrameMasterTable                                            *IconFrameMasterTable                                            `json:"iconFrameMasterTable" msgpack:"iconFrameMasterTable"`
	ItemMasterTable                                                 *ItemMasterTable                                                 `json:"itemMasterTable" msgpack:"itemMasterTable"`
	JewelShopCategoryMasterTable                                    *JewelShopCategoryMasterTable                                    `json:"jewelShopCategoryMasterTable" msgpack:"jewelShopCategoryMasterTable"`
	JewelShopItemMasterTable                                        *JewelShopItemMasterTable                                        `json:"jewelShopItemMasterTable" msgpack:"jewelShopItemMasterTable"`
	LeaderSenseMasterTable                                          *LeaderSenseMasterTable                                          `json:"leaderSenseMasterTable" msgpack:"leaderSenseMasterTable"`
	LeagueAllClassGlobalRankingRewardPackageMasterTable             *LeagueAllClassGlobalRankingRewardPackageMasterTable             `json:"leagueAllClassGlobalRankingRewardPackageMasterTable" msgpack:"leagueAllClassGlobalRankingRewardPackageMasterTable"`
	LeagueClassGroupMasterTable                                     *LeagueClassGroupMasterTable                                     `json:"leagueClassGroupMasterTable" msgpack:"leagueClassGroupMasterTable"`
	LeagueGroupRankingRewardPackageMasterTable                      *LeagueGroupRankingRewardPackageMasterTable                      `json:"leagueGroupRankingRewardPackageMasterTable" msgpack:"leagueGroupRankingRewardPackageMasterTable"`
	LeagueMasterTable                                               *LeagueMasterTable                                               `json:"leagueMasterTable" msgpack:"leagueMasterTable"`
	LeaguePlayRewardPackageMasterTable                              *LeaguePlayRewardPackageMasterTable                              `json:"leaguePlayRewardPackageMasterTable" msgpack:"leaguePlayRewardPackageMasterTable"`
	LeagueRewardPackageMasterTable                                  *LeagueRewardPackageMasterTable                                  `json:"leagueRewardPackageMasterTable" msgpack:"leagueRewardPackageMasterTable"`
	LessonScoreRewardGroupMasterTable                               *LessonScoreRewardGroupMasterTable                               `json:"lessonScoreRewardGroupMasterTable" msgpack:"lessonScoreRewardGroupMasterTable"`
	LessonScoreRewardMasterTable                                    *LessonScoreRewardMasterTable                                    `json:"lessonScoreRewardMasterTable" msgpack:"lessonScoreRewardMasterTable"`
	LightLoadSplitEffectMasterTable                                 *LightLoadSplitEffectMasterTable                                 `json:"lightLoadSplitEffectMasterTable" msgpack:"lightLoadSplitEffectMasterTable"`
	LipSyncMasterTable                                              *LipSyncMasterTable                                              `json:"lipSyncMasterTable" msgpack:"lipSyncMasterTable"`
	LiveDropFrameGroupMasterTable                                   *LiveDropFrameGroupMasterTable                                   `json:"liveDropFrameGroupMasterTable" msgpack:"liveDropFrameGroupMasterTable"`
	LiveMasterTable                                                 *LiveMasterTable                                                 `json:"liveMasterTable" msgpack:"liveMasterTable"`
	LiveSettingMasterTable                                          *LiveSettingMasterTable                                          `json:"liveSettingMasterTable" msgpack:"liveSettingMasterTable"`
	LoginBonusSpineCostumeMasterTable                               *LoginBonusSpineCostumeMasterTable                               `json:"loginBonusSpineCostumeMasterTable" msgpack:"loginBonusSpineCostumeMasterTable"`
	LoopMotionMasterTable                                           *LoopMotionMasterTable                                           `json:"loopMotionMasterTable" msgpack:"loopMotionMasterTable"`
	MarketFrameThingMasterTable                                     *MarketFrameThingMasterTable                                     `json:"marketFrameThingMasterTable" msgpack:"marketFrameThingMasterTable"`
	MissionMasterTable                                              *MissionMasterTable                                              `json:"missionMasterTable" msgpack:"missionMasterTable"`
	MissionPassDetailMasterTable                                    *MissionPassDetailMasterTable                                    `json:"missionPassDetailMasterTable" msgpack:"missionPassDetailMasterTable"`
	MissionPassLoopRewardMasterTable                                *MissionPassLoopRewardMasterTable                                `json:"missionPassLoopRewardMasterTable" msgpack:"missionPassLoopRewardMasterTable"`
	MissionPassLoopRewardThingMasterTable                           *MissionPassLoopRewardThingMasterTable                           `json:"missionPassLoopRewardThingMasterTable" msgpack:"missionPassLoopRewardThingMasterTable"`
	MissionPassMasterTable                                          *MissionPassMasterTable                                          `json:"missionPassMasterTable" msgpack:"missionPassMasterTable"`
	MultiLiveScheduleMasterTable                                    *MultiLiveScheduleMasterTable                                    `json:"multiLiveScheduleMasterTable" msgpack:"multiLiveScheduleMasterTable"`
	MusicCourseMasterTable                                          *MusicCourseMasterTable                                          `json:"musicCourseMasterTable" msgpack:"musicCourseMasterTable"`
	MusicCourseRewardGroupMasterTable                               *MusicCourseRewardGroupMasterTable                               `json:"musicCourseRewardGroupMasterTable" msgpack:"musicCourseRewardGroupMasterTable"`
	MusicGroupDetailMasterTable                                     *MusicGroupDetailMasterTable                                     `json:"musicGroupDetailMasterTable" msgpack:"musicGroupDetailMasterTable"`
	MusicGroupMasterTable                                           *MusicGroupMasterTable                                           `json:"musicGroupMasterTable" msgpack:"musicGroupMasterTable"`
	MusicMasterTable                                                *MusicMasterTable                                                `json:"musicMasterTable" msgpack:"musicMasterTable"`
	MusicVideoDefaultCostumeGroupMasterTable                        *MusicVideoDefaultCostumeGroupMasterTable                        `json:"musicVideoDefaultCostumeGroupMasterTable" msgpack:"musicVideoDefaultCostumeGroupMasterTable"`
	MusicVideoMasterTable                                           *MusicVideoMasterTable                                           `json:"musicVideoMasterTable" msgpack:"musicVideoMasterTable"`
	MusicVocalVersionMasterTable                                    *MusicVocalVersionMasterTable                                    `json:"musicVocalVersionMasterTable" msgpack:"musicVocalVersionMasterTable"`
	NameBaseColorMasterTable                                        *NameBaseColorMasterTable                                        `json:"nameBaseColorMasterTable" msgpack:"nameBaseColorMasterTable"`
	NameColorMasterTable                                            *NameColorMasterTable                                            `json:"nameColorMasterTable" msgpack:"nameColorMasterTable"`
	NameplateMasterTable                                            *NameplateMasterTable                                            `json:"nameplateMasterTable" msgpack:"nameplateMasterTable"`
	NgWordMasterTable                                               *NgWordMasterTable                                               `json:"ngWordMasterTable" msgpack:"ngWordMasterTable"`
	NoteMasterTable                                                 *NoteMasterTable                                                 `json:"noteMasterTable" msgpack:"noteMasterTable"`
	PermanentMarketThingMasterTable                                 *PermanentMarketThingMasterTable                                 `json:"permanentMarketThingMasterTable" msgpack:"permanentMarketThingMasterTable"`
	PhotoEffectChangeItemMasterTable                                *PhotoEffectChangeItemMasterTable                                `json:"photoEffectChangeItemMasterTable" msgpack:"photoEffectChangeItemMasterTable"`
	PhotoEffectMasterTable                                          *PhotoEffectMasterTable                                          `json:"photoEffectMasterTable" msgpack:"photoEffectMasterTable"`
	PhotoEffectTypeGroupMasterTable                                 *PhotoEffectTypeGroupMasterTable                                 `json:"photoEffectTypeGroupMasterTable" msgpack:"photoEffectTypeGroupMasterTable"`
	PhotoEffectVarietyChangeDetailMasterTable                       *PhotoEffectVarietyChangeDetailMasterTable                       `json:"photoEffectVarietyChangeDetailMasterTable" msgpack:"photoEffectVarietyChangeDetailMasterTable"`
	PhotoEffectVarietyUpDetailMasterTable                           *PhotoEffectVarietyUpDetailMasterTable                           `json:"photoEffectVarietyUpDetailMasterTable" msgpack:"photoEffectVarietyUpDetailMasterTable"`
	PhotoLevelUpItemGroupMasterTable                                *PhotoLevelUpItemGroupMasterTable                                `json:"photoLevelUpItemGroupMasterTable" msgpack:"photoLevelUpItemGroupMasterTable"`
	PhotoSpotMasterTable                                            *PhotoSpotMasterTable                                            `json:"photoSpotMasterTable" msgpack:"photoSpotMasterTable"`
	PickupCharacterMissionDetailGroupMasterTable                    *PickupCharacterMissionDetailGroupMasterTable                    `json:"pickupCharacterMissionDetailGroupMasterTable" msgpack:"pickupCharacterMissionDetailGroupMasterTable"`
	PickupCharacterMissionDetailMasterTable                         *PickupCharacterMissionDetailMasterTable                         `json:"pickupCharacterMissionDetailMasterTable" msgpack:"pickupCharacterMissionDetailMasterTable"`
	PickupCharacterMissionMasterTable                               *PickupCharacterMissionMasterTable                               `json:"pickupCharacterMissionMasterTable" msgpack:"pickupCharacterMissionMasterTable"`
	PickupSelectionGachaMasterTable                                 *PickupSelectionGachaMasterTable                                 `json:"pickupSelectionGachaMasterTable" msgpack:"pickupSelectionGachaMasterTable"`
	PlayerRankCapMasterTable                                        *PlayerRankCapMasterTable                                        `json:"playerRankCapMasterTable" msgpack:"playerRankCapMasterTable"`
	PlayerRankMasterTable                                           *PlayerRankMasterTable                                           `json:"playerRankMasterTable" msgpack:"playerRankMasterTable"`
	PosterAbilityMasterTable                                        *PosterAbilityMasterTable                                        `json:"posterAbilityMasterTable" msgpack:"posterAbilityMasterTable"`
	PosterLevelPatternGroupMasterTable                              *PosterLevelPatternGroupMasterTable                              `json:"posterLevelPatternGroupMasterTable" msgpack:"posterLevelPatternGroupMasterTable"`
	PosterMasterTable                                               *PosterMasterTable                                               `json:"posterMasterTable" msgpack:"posterMasterTable"`
	PosterReleaseItemGroupMasterTable                               *PosterReleaseItemGroupMasterTable                               `json:"posterReleaseItemGroupMasterTable" msgpack:"posterReleaseItemGroupMasterTable"`
	PosterStoryMasterTable                                          *PosterStoryMasterTable                                          `json:"posterStoryMasterTable" msgpack:"posterStoryMasterTable"`
	RandomEffectGroupMasterTable                                    *RandomEffectGroupMasterTable                                    `json:"randomEffectGroupMasterTable" msgpack:"randomEffectGroupMasterTable"`
	ResultVoiceMasterTable                                          *ResultVoiceMasterTable                                          `json:"resultVoiceMasterTable" msgpack:"resultVoiceMasterTable"`
	RewardRuleMasterTable                                           *RewardRuleMasterTable                                           `json:"rewardRuleMasterTable" msgpack:"rewardRuleMasterTable"`
	RouletteEventMasterTable                                        *RouletteEventMasterTable                                        `json:"rouletteEventMasterTable" msgpack:"rouletteEventMasterTable"`
	RouletteMasterTable                                             *RouletteMasterTable                                             `json:"rouletteMasterTable" msgpack:"rouletteMasterTable"`
	RoulettePrizeMasterTable                                        *RoulettePrizeMasterTable                                        `json:"roulettePrizeMasterTable" msgpack:"roulettePrizeMasterTable"`
	SceneCameraMasterTable                                          *SceneCameraMasterTable                                          `json:"sceneCameraMasterTable" msgpack:"sceneCameraMasterTable"`
	SenseBranchMasterTable                                          *SenseBranchMasterTable                                          `json:"senseBranchMasterTable" msgpack:"senseBranchMasterTable"`
	SenseMasterTable                                                *SenseMasterTable                                                `json:"senseMasterTable" msgpack:"senseMasterTable"`
	SenseNotationMasterTable                                        *SenseNotationMasterTable                                        `json:"senseNotationMasterTable" msgpack:"senseNotationMasterTable"`
	SensePerformanceMasterTable                                     *SensePerformanceMasterTable                                     `json:"sensePerformanceMasterTable" msgpack:"sensePerformanceMasterTable"`
	SignMasterTable                                                 *SignMasterTable                                                 `json:"signMasterTable" msgpack:"signMasterTable"`
	SpecialEpisodeMasterTable                                       *SpecialEpisodeMasterTable                                       `json:"specialEpisodeMasterTable" msgpack:"specialEpisodeMasterTable"`
	SpecialEventMasterTable                                         *SpecialEventMasterTable                                         `json:"specialEventMasterTable" msgpack:"specialEventMasterTable"`
	SpecialStoryMasterTable                                         *SpecialStoryMasterTable                                         `json:"specialStoryMasterTable" msgpack:"specialStoryMasterTable"`
	SplashMasterTable                                               *SplashMasterTable                                               `json:"splashMasterTable" msgpack:"splashMasterTable"`
	SpotConversationMasterTable                                     *SpotConversationMasterTable                                     `json:"spotConversationMasterTable" msgpack:"spotConversationMasterTable"`
	StaminaRecoveryItemMasterTable                                  *StaminaRecoveryItemMasterTable                                  `json:"staminaRecoveryItemMasterTable" msgpack:"staminaRecoveryItemMasterTable"`
	StampMasterTable                                                *StampMasterTable                                                `json:"stampMasterTable" msgpack:"stampMasterTable"`
	StarActBranchMasterTable                                        *StarActBranchMasterTable                                        `json:"starActBranchMasterTable" msgpack:"starActBranchMasterTable"`
	StarActConditionMasterTable                                     *StarActConditionMasterTable                                     `json:"starActConditionMasterTable" msgpack:"starActConditionMasterTable"`
	StarActMasterTable                                              *StarActMasterTable                                              `json:"starActMasterTable" msgpack:"starActMasterTable"`
	StarRankRewardMasterTable                                       *StarRankRewardMasterTable                                       `json:"starRankRewardMasterTable" msgpack:"starRankRewardMasterTable"`
	StepupGachaGroupMasterTable                                     *StepupGachaGroupMasterTable                                     `json:"stepupGachaGroupMasterTable" msgpack:"stepupGachaGroupMasterTable"`
	StoryEventBonusCharacterBaseMasterTable                         *StoryEventBonusCharacterBaseMasterTable                         `json:"storyEventBonusCharacterBaseMasterTable" msgpack:"storyEventBonusCharacterBaseMasterTable"`
	StoryEventCircleHighScoreRewardMasterTable                      *StoryEventCircleHighScoreRewardMasterTable                      `json:"storyEventCircleHighScoreRewardMasterTable" msgpack:"storyEventCircleHighScoreRewardMasterTable"`
	StoryEventCircleMasterTable                                     *StoryEventCircleMasterTable                                     `json:"storyEventCircleMasterTable" msgpack:"storyEventCircleMasterTable"`
	StoryEventCircleMissionMasterTable                              *StoryEventCircleMissionMasterTable                              `json:"storyEventCircleMissionMasterTable" msgpack:"storyEventCircleMissionMasterTable"`
	StoryEventCircleMissionRewardMasterTable                        *StoryEventCircleMissionRewardMasterTable                        `json:"storyEventCircleMissionRewardMasterTable" msgpack:"storyEventCircleMissionRewardMasterTable"`
	StoryEventEpisodeMasterTable                                    *StoryEventEpisodeMasterTable                                    `json:"storyEventEpisodeMasterTable" msgpack:"storyEventEpisodeMasterTable"`
	StoryEventHighScoreBuffMasterTable                              *StoryEventHighScoreBuffMasterTable                              `json:"storyEventHighScoreBuffMasterTable" msgpack:"storyEventHighScoreBuffMasterTable"`
	StoryEventHighScoreBuffPatternGroupMasterTable                  *StoryEventHighScoreBuffPatternGroupMasterTable                  `json:"storyEventHighScoreBuffPatternGroupMasterTable" msgpack:"storyEventHighScoreBuffPatternGroupMasterTable"`
	StoryEventHighScoreBuffSettingMasterTable                       *StoryEventHighScoreBuffSettingMasterTable                       `json:"storyEventHighScoreBuffSettingMasterTable" msgpack:"storyEventHighScoreBuffSettingMasterTable"`
	StoryEventHighScoreMasterTable                                  *StoryEventHighScoreMasterTable                                  `json:"storyEventHighScoreMasterTable" msgpack:"storyEventHighScoreMasterTable"`
	StoryEventHighScoreRewardMasterTable                            *StoryEventHighScoreRewardMasterTable                            `json:"storyEventHighScoreRewardMasterTable" msgpack:"storyEventHighScoreRewardMasterTable"`
	StoryEventMasterTable                                           *StoryEventMasterTable                                           `json:"storyEventMasterTable" msgpack:"storyEventMasterTable"`
	StoryEventRewardItemPackageTable                                *StoryEventRewardItemPackageTable                                `json:"storyEventRewardItemPackageTable" msgpack:"storyEventRewardItemPackageTable"`
	StoryEventRewardMasterTable                                     *StoryEventRewardMasterTable                                     `json:"storyEventRewardMasterTable" msgpack:"storyEventRewardMasterTable"`
	StoryEventStoryBgmGroupMasterTable                              *StoryEventStoryBgmGroupMasterTable                              `json:"storyEventStoryBgmGroupMasterTable" msgpack:"storyEventStoryBgmGroupMasterTable"`
	StoryEventStoryMasterTable                                      *StoryEventStoryMasterTable                                      `json:"storyEventStoryMasterTable" msgpack:"storyEventStoryMasterTable"`
	StoryEventTotalPointRewardMasterTable                           *StoryEventTotalPointRewardMasterTable                           `json:"storyEventTotalPointRewardMasterTable" msgpack:"storyEventTotalPointRewardMasterTable"`
	StoryMasterTable                                                *StoryMasterTable                                                `json:"storyMasterTable" msgpack:"storyMasterTable"`
	StoryRelationMasterTable                                        *StoryRelationMasterTable                                        `json:"storyRelationMasterTable" msgpack:"storyRelationMasterTable"`
	TeamChallengeMasterTable                                        *TeamChallengeMasterTable                                        `json:"teamChallengeMasterTable" msgpack:"teamChallengeMasterTable"`
	TheaterChapterMasterTable                                       *TheaterChapterMasterTable                                       `json:"theaterChapterMasterTable" msgpack:"theaterChapterMasterTable"`
	TheaterStoryMasterTable                                         *TheaterStoryMasterTable                                         `json:"theaterStoryMasterTable" msgpack:"theaterStoryMasterTable"`
	TimeLimitedControlMasterTable                                   *TimeLimitedControlMasterTable                                   `json:"timeLimitedControlMasterTable" msgpack:"timeLimitedControlMasterTable"`
	TipMasterTable                                                  *TipMasterTable                                                  `json:"tipMasterTable" msgpack:"tipMasterTable"`
	TitleBackgroundMasterTable                                      *TitleBackgroundMasterTable                                      `json:"titleBackgroundMasterTable" msgpack:"titleBackgroundMasterTable"`
	TitleCallVoiceMasterTable                                       *TitleCallVoiceMasterTable                                       `json:"titleCallVoiceMasterTable" msgpack:"titleCallVoiceMasterTable"`
	TitleDecorationMasterTable                                      *TitleDecorationMasterTable                                      `json:"titleDecorationMasterTable" msgpack:"titleDecorationMasterTable"`
	TotalPointEventMasterTable                                      *TotalPointEventMasterTable                                      `json:"totalPointEventMasterTable" msgpack:"totalPointEventMasterTable"`
	TotalPointEventRewardMasterTable                                *TotalPointEventRewardMasterTable                                `json:"totalPointEventRewardMasterTable" msgpack:"totalPointEventRewardMasterTable"`
	TournamentMasterTable                                           *TournamentMasterTable                                           `json:"tournamentMasterTable" msgpack:"tournamentMasterTable"`
	TournamentQualifyingMasterTable                                 *TournamentQualifyingMasterTable                                 `json:"tournamentQualifyingMasterTable" msgpack:"tournamentQualifyingMasterTable"`
	TrialPartyAccessoryMasterTable                                  *TrialPartyAccessoryMasterTable                                  `json:"trialPartyAccessoryMasterTable" msgpack:"trialPartyAccessoryMasterTable"`
	TrialPartyCharacterMasterTable                                  *TrialPartyCharacterMasterTable                                  `json:"trialPartyCharacterMasterTable" msgpack:"trialPartyCharacterMasterTable"`
	TrialPartyEventMasterTable                                      *TrialPartyEventMasterTable                                      `json:"trialPartyEventMasterTable" msgpack:"trialPartyEventMasterTable"`
	TrialPartyEventStageMasterTable                                 *TrialPartyEventStageMasterTable                                 `json:"trialPartyEventStageMasterTable" msgpack:"trialPartyEventStageMasterTable"`
	TrialPartyMasterTable                                           *TrialPartyMasterTable                                           `json:"trialPartyMasterTable" msgpack:"trialPartyMasterTable"`
	TrialPartyPosterMasterTable                                     *TrialPartyPosterMasterTable                                     `json:"trialPartyPosterMasterTable" msgpack:"trialPartyPosterMasterTable"`
	TripleCastAllClassGlobalRankingRewardPackageMasterTable         *TripleCastAllClassGlobalRankingRewardPackageMasterTable         `json:"tripleCastAllClassGlobalRankingRewardPackageMasterTable" msgpack:"tripleCastAllClassGlobalRankingRewardPackageMasterTable"`
	TripleCastGroupRankingRewardPackageMasterTable                  *TripleCastGroupRankingRewardPackageMasterTable                  `json:"tripleCastGroupRankingRewardPackageMasterTable" msgpack:"tripleCastGroupRankingRewardPackageMasterTable"`
	TripleCastMasterTable                                           *TripleCastMasterTable                                           `json:"tripleCastMasterTable" msgpack:"tripleCastMasterTable"`
	TrophyGroupMasterTable                                          *TrophyGroupMasterTable                                          `json:"trophyGroupMasterTable" msgpack:"trophyGroupMasterTable"`
	TrophyMasterTable                                               *TrophyMasterTable                                               `json:"trophyMasterTable" msgpack:"trophyMasterTable"`
	UnlockConditionMasterTable                                      *UnlockConditionMasterTable                                      `json:"unlockConditionMasterTable" msgpack:"unlockConditionMasterTable"`
}

type MemoryDatabaseLite struct {
	PosterMasterTable          *PosterMasterTable          `json:"posterMasterTable" msgpack:"posterMasterTable"`
	TitleBackgroundMasterTable *TitleBackgroundMasterTable `json:"titleBackgroundMasterTable" msgpack:"titleBackgroundMasterTable"`
	TitleDecorationMasterTable *TitleDecorationMasterTable `json:"titleDecorationMasterTable" msgpack:"titleDecorationMasterTable"`
	TitleCallVoiceMasterTable  *TitleCallVoiceMasterTable  `json:"titleCallVoiceMasterTable" msgpack:"titleCallVoiceMasterTable"`
}

type TimedConfirmationCode struct {
	ConfirmationCode string `json:"confirmationCode" msgpack:"confirmationCode"` // Key: 0
	RemainingSeconds int    `json:"remainingSeconds" msgpack:"remainingSeconds"` // Key: 1
}

type UserResult struct {
	User *User `json:"user" msgpack:"user"` // Key: 0
}

type Status struct {
	Vocal         int `json:"vocal" msgpack:"vocal"`                 // Key: 0
	Expression    int `json:"expression" msgpack:"expression"`       // Key: 1
	Concentration int `json:"concentration" msgpack:"concentration"` // Key: 2
}

type SellAccessoryPayload struct {
	UAccessoryIds []int64 `json:"uAccessoryIds" msgpack:"uAccessoryIds"` // Key: 0
}

type AccessoryFavoritePayload struct {
	AccessoryId int64 `json:"accessoryId" msgpack:"accessoryId"` // Key: 0
	SetFavorite bool  `json:"setFavorite" msgpack:"setFavorite"` // Key: 1
}

type RegisterPayload struct {
	Name string `json:"name" msgpack:"name"` // Key: 0
}

type AuthenticatePayload struct {
	LoginToken              string       `json:"loginToken" msgpack:"loginToken"`                           // Key: 0
	GameVersion             GameVersions `json:"gameVersion" msgpack:"gameVersion"`                         // Key: 1
	ApkHash                 string       `json:"apkHash" msgpack:"apkHash"`                                 // Key: 2
	ApkApplicationSignature string       `json:"apkApplicationSignature" msgpack:"apkApplicationSignature"` // Key: 3
	ApplicationVersion      string       `json:"applicationVersion" msgpack:"applicationVersion"`           // Key: 4
}

type RegisterTakeOverPasswordPayload struct {
	Password string `json:"password" msgpack:"password"` // Key: 0
}

type AccountConnectPayload struct {
	Provider AuthenticationProviders `json:"provider" msgpack:"provider"` // Key: 0
	Token    string                  `json:"token" msgpack:"token"`       // Key: 1
}

type TakeOverAccountPayload struct {
	LinkageCode string `json:"linkageCode" msgpack:"linkageCode"` // Key: 0
	Password    string `json:"password" msgpack:"password"`       // Key: 1
}

type RegisterBirthDayPayload struct {
	BirthDate time.Time `json:"birthDate" msgpack:"birthDate"` // Key: 0
}

type GetPushNotificationTokenPayload struct {
	HashUserId string `json:"hashUserId" msgpack:"hashUserId"` // Key: 0
}

type UseExperienceItemsPayload struct {
	ItemMasterId int64 `json:"itemMasterId" msgpack:"itemMasterId"` // Key: 0
	Quantity     int   `json:"quantity" msgpack:"quantity"`         // Key: 1
}

type BulkLevelUpPayload struct {
	CharacterId int64 `json:"characterId" msgpack:"characterId"` // Key: 0
	Order       int   `json:"order" msgpack:"order"`             // Key: 1
}

type ActorPortalCharacterPayload struct {
	CharacterBaseId int64 `json:"characterBaseId" msgpack:"characterBaseId"` // Key: 0
	CharacterId     int64 `json:"characterId" msgpack:"characterId"`         // Key: 1
	IsAwakening     bool  `json:"isAwakening" msgpack:"isAwakening"`         // Key: 2
}

type CharacterFavoritePayload struct {
	CharacterId int64 `json:"characterId" msgpack:"characterId"` // Key: 0
	SetFavorite bool  `json:"setFavorite" msgpack:"setFavorite"` // Key: 1
}

type CirclePayload struct {
	HashedId              string                           `json:"hashedId" msgpack:"hashedId"`                           // Key: 0
	Name                  string                           `json:"name" msgpack:"name"`                                   // Key: 1
	Comment               string                           `json:"comment" msgpack:"comment"`                             // Key: 2
	PlayTimeStartType     PlayTimeTypes                    `json:"playTimeStartType" msgpack:"playTimeStartType"`         // Key: 3
	PlayTimeEndType       PlayTimeTypes                    `json:"playTimeEndType" msgpack:"playTimeEndType"`             // Key: 4
	EntryType             EntryTypes                       `json:"entryType" msgpack:"entryType"`                         // Key: 5
	MemberType            SearchCircleMemberConditionTypes `json:"memberType" msgpack:"memberType"`                       // Key: 6
	CompanyMasterId       *int64                           `json:"companyMasterId" msgpack:"companyMasterId"`             // Key: 7
	CharacterBaseMasterId *int64                           `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 8
}

type BannerPayload struct {
	CircleHashedId      string  `json:"circleHashedId" msgpack:"circleHashedId"`           // Key: 0
	PosterMasterId      int64   `json:"posterMasterId" msgpack:"posterMasterId"`           // Key: 1
	X1                  float32 `json:"x1" msgpack:"x1"`                                   // Key: 2
	Y1                  float32 `json:"y1" msgpack:"y1"`                                   // Key: 3
	X2                  float32 `json:"x2" msgpack:"x2"`                                   // Key: 4
	Y2                  float32 `json:"y2" msgpack:"y2"`                                   // Key: 5
	RotationAngle       int     `json:"rotationAngle" msgpack:"rotationAngle"`             // Key: 6
	DisplayPosterString bool    `json:"displayPosterString" msgpack:"displayPosterString"` // Key: 7
	BannerRatio         float32 `json:"bannerRatio" msgpack:"bannerRatio"`                 // Key: 8
}

type CircleAuthorityChangePayload struct {
	CircleHashedId string            `json:"circleHashedId" msgpack:"circleHashedId"` // Key: 0
	UserId         string            `json:"userId" msgpack:"userId"`                 // Key: 1
	Authority      CircleAuthorities `json:"authority" msgpack:"authority"`           // Key: 2
}

type DonateSupportLevelLimitPayload struct {
	Company        Companies                       `json:"company" msgpack:"company"`               // Key: 0
	NextLevelLimit int                             `json:"nextLevelLimit" msgpack:"nextLevelLimit"` // Key: 1
	CoinQuantity   int                             `json:"coinQuantity" msgpack:"coinQuantity"`     // Key: 2
	Details        []DonateSupportLevelLimitDetail `json:"details" msgpack:"details"`               // Key: 3
}

type DonateSupportLevelLimitDetail struct {
	Order    int `json:"order" msgpack:"order"`       // Key: 0
	Quantity int `json:"quantity" msgpack:"quantity"` // Key: 1
}

type ChatworkSendMessagePayload struct {
	Message string `json:"message" msgpack:"message"` // Key: 0
}

type AcquirableThingsPayload struct {
	Things []AcquirableThing `json:"things" msgpack:"things"` // Key: 0
}

type AcquirableThing struct {
	Id       int64 `json:"id" msgpack:"id"`             // Key: 0
	Quantity int   `json:"quantity" msgpack:"quantity"` // Key: 1
}

type DebugAlbumSimpleArrangingPayload struct {
	Publishing    bool                `json:"publishing" msgpack:"publishing"`       // Key: 0
	Page          int                 `json:"page" msgpack:"page"`                   // Key: 1
	Photos        []AlbumPhotoPayload `json:"photos" msgpack:"photos"`               // Key: 2
	MAlbumThemeId *int64              `json:"mAlbumThemeId" msgpack:"mAlbumThemeId"` // Key: 3
}

type DebugPrepareLeagueGroupUserPayload struct {
	LeagueMasterId int64            `json:"leagueMasterId" msgpack:"leagueMasterId"` // Key: 0
	ClassType      LeagueClassTypes `json:"classType" msgpack:"classType"`           // Key: 1
	UserIds        []int64          `json:"userIds" msgpack:"userIds"`               // Key: 2
}

type DebugEditLeagueBasicPayload struct {
	CurrentClassType LeagueClassTypes `json:"currentClassType" msgpack:"currentClassType"` // Key: 0
	BestClassType    LeagueClassTypes `json:"bestClassType" msgpack:"bestClassType"`       // Key: 1
}

type DebugModifyCharacterEnhanceInformationPayload struct {
	MCharacterId     int64 `json:"mCharacterId" msgpack:"mCharacterId"`         // Key: 0
	ActorLevel       int   `json:"actorLevel" msgpack:"actorLevel"`             // Key: 1
	IsAwaken         bool  `json:"isAwaken" msgpack:"isAwaken"`                 // Key: 2
	SenseLevel       int   `json:"senseLevel" msgpack:"senseLevel"`             // Key: 3
	TalentBloomStage int   `json:"talentBloomStage" msgpack:"talentBloomStage"` // Key: 4
}

type DebugModifyPosterEnhanceInformationPayload struct {
	MPosterId    int64 `json:"mPosterId" msgpack:"mPosterId"`       // Key: 0
	Level        int   `json:"level" msgpack:"level"`               // Key: 1
	ReleasePhase int   `json:"releasePhase" msgpack:"releasePhase"` // Key: 2
}

type GenerateLotteryResultPayload struct {
	Prize  int `json:"prize" msgpack:"prize"`   // Key: 0
	Amount int `json:"amount" msgpack:"amount"` // Key: 1
}

type PartyPayload struct {
	LeaderPosition   int                `json:"leaderPosition" msgpack:"leaderPosition"`     // Key: 0
	PartySlotPayload []PartySlotPayload `json:"partySlotPayload" msgpack:"partySlotPayload"` // Key: 1
}

type PartySlotPayload struct {
	Position                  int                        `json:"position" msgpack:"position"`                                   // Key: 0
	PartySlotCharacterPayload *PartySlotCharacterPayload `json:"partySlotCharacterPayload" msgpack:"partySlotCharacterPayload"` // Key: 1
	PartySlotPosterPayload    *PartySlotPosterPayload    `json:"partySlotPosterPayload" msgpack:"partySlotPosterPayload"`       // Key: 2
	PartySlotAccessoryPayload *PartySlotAccessoryPayload `json:"partySlotAccessoryPayload" msgpack:"partySlotAccessoryPayload"` // Key: 3
}

type PartySlotCharacterPayload struct {
	MCharacterId           int64                 `json:"mCharacterId" msgpack:"mCharacterId"`                     // Key: 0
	TalentStage            int                   `json:"talentStage" msgpack:"talentStage"`                       // Key: 1
	AwakeningPhase         int                   `json:"awakeningPhase" msgpack:"awakeningPhase"`                 // Key: 2
	Level                  int                   `json:"level" msgpack:"level"`                                   // Key: 3
	SenseLevel             int                   `json:"senseLevel" msgpack:"senseLevel"`                         // Key: 4
	CurrentExperience      int                   `json:"currentExperience" msgpack:"currentExperience"`           // Key: 5
	ReleasedEpisodeOrder   CharacterEpisodeOrder `json:"releasedEpisodeOrder" msgpack:"releasedEpisodeOrder"`     // Key: 6
	ReadEpisodeOrder       CharacterEpisodeOrder `json:"readEpisodeOrder" msgpack:"readEpisodeOrder"`             // Key: 7
	DisplayAwakeningStatus bool                  `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 8
}

type PartySlotPosterPayload struct {
	MPosterId       int64              `json:"mPosterId" msgpack:"mPosterId"`             // Key: 0
	Level           int                `json:"level" msgpack:"level"`                     // Key: 1
	Phase           int                `json:"phase" msgpack:"phase"`                     // Key: 2
	ReleasedEpisode PosterEpisodeTypes `json:"releasedEpisode" msgpack:"releasedEpisode"` // Key: 3
}

type PartySlotAccessoryPayload struct {
	MAccessoryId        int64  `json:"mAccessoryId" msgpack:"mAccessoryId"`               // Key: 0
	Level               int    `json:"level" msgpack:"level"`                             // Key: 1
	MAccessoryEffectId1 *int64 `json:"mAccessoryEffectId1" msgpack:"mAccessoryEffectId1"` // Key: 2
	MAccessoryEffectId2 *int64 `json:"mAccessoryEffectId2" msgpack:"mAccessoryEffectId2"` // Key: 3
	MAccessoryEffectId3 *int64 `json:"mAccessoryEffectId3" msgpack:"mAccessoryEffectId3"` // Key: 4
}

type MusicCourseRankingPayload struct {
	UserId                            int64      `json:"userId" msgpack:"userId"`                                                       // Key: 0
	CurrentChallengeCount             int        `json:"currentChallengeCount" msgpack:"currentChallengeCount"`                         // Key: 1
	PerfectStar                       int        `json:"perfectStar" msgpack:"perfectStar"`                                             // Key: 2
	Perfect                           int        `json:"perfect" msgpack:"perfect"`                                                     // Key: 3
	Great                             int        `json:"great" msgpack:"great"`                                                         // Key: 4
	Good                              int        `json:"good" msgpack:"good"`                                                           // Key: 5
	Bad                               int        `json:"bad" msgpack:"bad"`                                                             // Key: 6
	Miss                              int        `json:"miss" msgpack:"miss"`                                                           // Key: 7
	TotalAchievementRatePercentRecord float64    `json:"totalAchievementRatePercentRecord" msgpack:"totalAchievementRatePercentRecord"` // Key: 8
	BestRecordChallengeCount          int        `json:"bestRecordChallengeCount" msgpack:"bestRecordChallengeCount"`                   // Key: 9
	BestRecordDate                    *time.Time `json:"bestRecordDate" msgpack:"bestRecordDate"`                                       // Key: 10
}

type SimulateLiveDropPayload struct {
	LiveSettingMasterId     int64    `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"`         // Key: 0
	MusicMasterId           int64    `json:"musicMasterId" msgpack:"musicMasterId"`                     // Key: 1
	PartyId                 int64    `json:"partyId" msgpack:"partyId"`                                 // Key: 2
	ExecCount               int      `json:"execCount" msgpack:"execCount"`                             // Key: 3
	StaminaConsumptionRatio int      `json:"staminaConsumptionRatio" msgpack:"staminaConsumptionRatio"` // Key: 4
	LiveScore               *int64   `json:"liveScore" msgpack:"liveScore"`                             // Key: 5
	StarActFiredCount       *int     `json:"starActFiredCount" msgpack:"starActFiredCount"`             // Key: 6
	AchievementRatePercent  *float64 `json:"achievementRatePercent" msgpack:"achievementRatePercent"`   // Key: 7
	MemberCount             *int     `json:"memberCount" msgpack:"memberCount"`                         // Key: 8
	HasPurchasedAmulet      bool     `json:"hasPurchasedAmulet" msgpack:"hasPurchasedAmulet"`           // Key: 9
}

type SupportCompanyLevelLimitPayload struct {
	Company      Companies                        `json:"company" msgpack:"company"`           // Key: 0
	CoinQuantity int64                            `json:"coinQuantity" msgpack:"coinQuantity"` // Key: 1
	Details      []SupportCompanyLevelLimitDetail `json:"details" msgpack:"details"`           // Key: 2
}

type SupportCompanyLevelLimitDetail struct {
	Order    int   `json:"order" msgpack:"order"`       // Key: 0
	Quantity int64 `json:"quantity" msgpack:"quantity"` // Key: 1
}

type EditTrialPartyEventStagePartyPayload struct {
	TrialPartyEventStageMasterId int64                                      `json:"trialPartyEventStageMasterId" msgpack:"trialPartyEventStageMasterId"` // Key: 0
	Slots                        []EditTrialPartyEventStagePartyPayloadSlot `json:"slots" msgpack:"slots"`                                               // Key: 1
	LeaderPosition               int                                        `json:"leaderPosition" msgpack:"leaderPosition"`                             // Key: 2
}

type EditTrialPartyEventStagePartyPayloadSlot struct {
	Position                    int    `json:"position" msgpack:"position"`                                       // Key: 0
	TrialPartyCharacterMasterId int64  `json:"trialPartyCharacterMasterId" msgpack:"trialPartyCharacterMasterId"` // Key: 1
	TrialPartyPosterMasterId    *int64 `json:"trialPartyPosterMasterId" msgpack:"trialPartyPosterMasterId"`       // Key: 2
	TrialPartyAccessoryMasterId *int64 `json:"trialPartyAccessoryMasterId" msgpack:"trialPartyAccessoryMasterId"` // Key: 3
}

type FlashSaleReadStagePayload struct {
	ReadFlashSaleStageIds []int64 `json:"readFlashSaleStageIds" msgpack:"readFlashSaleStageIds"` // Key: 0
}

type FriendFavoritePayload struct {
	TargetUserId string `json:"targetUserId" msgpack:"targetUserId"` // Key: 0
	SetFavorite  bool   `json:"setFavorite" msgpack:"setFavorite"`   // Key: 1
}

type FriendInvitationMissionPayload struct {
	FriendInvitationMissionMasterIds []int64 `json:"friendInvitationMissionMasterIds" msgpack:"friendInvitationMissionMasterIds"` // Key: 0
}

type SetSelectedThingsPayload struct {
	GachaThingIds []int64 `json:"gachaThingIds" msgpack:"gachaThingIds"` // Key: 0
}

type ReadNotificationPayload struct {
	TabCategory NotificationTabCategory `json:"tabCategory" msgpack:"tabCategory"` // Key: 0
	ReadAt      time.Time               `json:"readAt" msgpack:"readAt"`           // Key: 1
}

type BulkReceivePayload struct {
	InboxIds []int64 `json:"inboxIds" msgpack:"inboxIds"` // Key: 0
}

type UseStaminaRecoveryItemsPayload struct {
	Items []UseStaminaRecoveryItem `json:"items" msgpack:"items"` // Key: 0
}

type UseStaminaRecoveryItem struct {
	ItemMasterId int64 `json:"itemMasterId" msgpack:"itemMasterId"` // Key: 0
	Quantity     int   `json:"quantity" msgpack:"quantity"`         // Key: 1
}

type SetLessonPartyPayload struct {
	Slots []SetLessonPartySlotPayload `json:"slots" msgpack:"slots"` // Key: 0
}

type SetLessonPartySlotPayload struct {
	Order       int   `json:"order" msgpack:"order"`             // Key: 0
	CharacterId int64 `json:"characterId" msgpack:"characterId"` // Key: 1
}

type FinishLessonPayload struct {
	Score    int64              `json:"score" msgpack:"score"`       // Key: 0
	MaxCombo int                `json:"maxCombo" msgpack:"maxCombo"` // Key: 1
	Judges   []NoteJudgePayload `json:"judges" msgpack:"judges"`     // Key: 2
}

type StartLivePayload struct {
	PartyId                      int64                `json:"partyId" msgpack:"partyId"`                                           // Key: 0
	LiveMasterId                 int64                `json:"liveMasterId" msgpack:"liveMasterId"`                                 // Key: 1
	AuditionMasterId             *int64               `json:"auditionMasterId" msgpack:"auditionMasterId"`                         // Key: 2
	LeagueMasterId               *int64               `json:"leagueMasterId" msgpack:"leagueMasterId"`                             // Key: 3
	UseStamina                   bool                 `json:"useStamina" msgpack:"useStamina"`                                     // Key: 4
	StaminaConsumptionRatio      int                  `json:"staminaConsumptionRatio" msgpack:"staminaConsumptionRatio"`           // Key: 5
	LiveSettingMasterId          int64                `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"`                   // Key: 6
	IsAutoPlay                   bool                 `json:"isAutoPlay" msgpack:"isAutoPlay"`                                     // Key: 7
	IsStoryEventChallenge        bool                 `json:"isStoryEventChallenge" msgpack:"isStoryEventChallenge"`               // Key: 8
	ConcertStageMasterId         *int64               `json:"concertStageMasterId" msgpack:"concertStageMasterId"`                 // Key: 9
	BonusLiveStageMasterId       *int64               `json:"bonusLiveStageMasterId" msgpack:"bonusLiveStageMasterId"`             // Key: 10
	MusicCourseDetailMasterId    *int64               `json:"musicCourseDetailMasterId" msgpack:"musicCourseDetailMasterId"`       // Key: 11
	MusicCourseGaugeType         MusicCourseGaugeType `json:"musicCourseGaugeType" msgpack:"musicCourseGaugeType"`                 // Key: 12
	GhostLiveMasterId            *int64               `json:"ghostLiveMasterId" msgpack:"ghostLiveMasterId"`                       // Key: 13
	TrialPartyEventStageMasterId *int64               `json:"trialPartyEventStageMasterId" msgpack:"trialPartyEventStageMasterId"` // Key: 14
}

type StartLessonPayload struct {
	CharacterBaseMasterId int   `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 0
	LiveMasterId          int64 `json:"liveMasterId" msgpack:"liveMasterId"`                   // Key: 1
}

type FinishLivePayload struct {
	Score                          int64                           `json:"score" msgpack:"score"`                                                   // Key: 0
	MaxCombo                       int                             `json:"maxCombo" msgpack:"maxCombo"`                                             // Key: 1
	Judges                         []NoteJudgePayload              `json:"judges" msgpack:"judges"`                                                 // Key: 2
	IsCleared                      bool                            `json:"isCleared" msgpack:"isCleared"`                                           // Key: 3
	BaseScoreBlocks                []BaseScoreBlock                `json:"baseScoreBlocks" msgpack:"baseScoreBlocks"`                               // Key: 4
	SenseScoreBlocks               []SenseScoreBlock               `json:"senseScoreBlocks" msgpack:"senseScoreBlocks"`                             // Key: 5
	StarActScoreBlocks             []StarActScoreBlock             `json:"starActScoreBlocks" msgpack:"starActScoreBlocks"`                         // Key: 6
	MultiLiveAdditionalScoreBlocks []MultiLiveAdditionalScoreBlock `json:"multiLiveAdditionalScoreBlocks" msgpack:"multiLiveAdditionalScoreBlocks"` // Key: 7
	TripleCastPartyScores          []TripleCastPartyScore          `json:"tripleCastPartyScores" msgpack:"tripleCastPartyScores"`                   // Key: 8
	TrialPartyEventStageResult     *TrialPartyEventStageResult     `json:"trialPartyEventStageResult" msgpack:"trialPartyEventStageResult"`         // Key: 9
}

type TripleCastPartyScore struct {
	Order int   `json:"order" msgpack:"order"` // Key: 0
	Score int64 `json:"score" msgpack:"score"` // Key: 1
}

type TrialPartyEventStageResult struct {
	PrincipalGaugeValue int `json:"principalGaugeValue" msgpack:"principalGaugeValue"` // Key: 0
}

type FinishAnotherNotationLivePayload struct {
	BaseScoreBlocks                []BaseScoreBlock                `json:"baseScoreBlocks" msgpack:"baseScoreBlocks"`                               // Key: 0
	SenseScoreBlocks               []SenseScoreBlock               `json:"senseScoreBlocks" msgpack:"senseScoreBlocks"`                             // Key: 1
	StarActScoreBlocks             []StarActScoreBlock             `json:"starActScoreBlocks" msgpack:"starActScoreBlocks"`                         // Key: 2
	MultiLiveAdditionalScoreBlocks []MultiLiveAdditionalScoreBlock `json:"multiLiveAdditionalScoreBlocks" msgpack:"multiLiveAdditionalScoreBlocks"` // Key: 3
	AnotherNotationMasterId        int64                           `json:"anotherNotationMasterId" msgpack:"anotherNotationMasterId"`               // Key: 4
	IsAutoPlay                     bool                            `json:"isAutoPlay" msgpack:"isAutoPlay"`                                         // Key: 5
}

type NoteJudgePayload struct {
	Timing TimingTypes `json:"timing" msgpack:"timing"` // Key: 0
	Count  int         `json:"count" msgpack:"count"`   // Key: 1
}

type ChangeVocalPayload struct {
	MusicMasterId int64 `json:"musicMasterId" msgpack:"musicMasterId"` // Key: 0
	VocalVersion  int   `json:"vocalVersion" msgpack:"vocalVersion"`   // Key: 1
}

type CalculateTimeEventPayload struct {
	MusicMasterId         int64  `json:"musicMasterId" msgpack:"musicMasterId"`                 // Key: 0
	SenseNotationMasterId *int64 `json:"senseNotationMasterId" msgpack:"senseNotationMasterId"` // Key: 1
	PartyId               int64  `json:"partyId" msgpack:"partyId"`                             // Key: 2
	VocalVersion          int    `json:"vocalVersion" msgpack:"vocalVersion"`                   // Key: 3
	IsStoryEventChallenge bool   `json:"isStoryEventChallenge" msgpack:"isStoryEventChallenge"` // Key: 4
}

type CalculateLessonTimeEventPayload struct {
	MusicMasterId         int64 `json:"musicMasterId" msgpack:"musicMasterId"`                 // Key: 0
	CharacterBaseMasterId int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
}

type StartMultiLiveSessionPayload struct {
	LiveSettingMasterId int64  `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"` // Key: 0
	RoomId              string `json:"roomId" msgpack:"roomId"`                           // Key: 1
	MusicMasterId       int64  `json:"musicMasterId" msgpack:"musicMasterId"`             // Key: 2
}

type StartMultiLivePayload struct {
	LiveMasterId            int64  `json:"liveMasterId" msgpack:"liveMasterId"`                       // Key: 0
	MultiLiveId             int64  `json:"multiLiveId" msgpack:"multiLiveId"`                         // Key: 1
	UseStamina              bool   `json:"useStamina" msgpack:"useStamina"`                           // Key: 2
	StaminaConsumptionRatio int    `json:"staminaConsumptionRatio" msgpack:"staminaConsumptionRatio"` // Key: 3
	PartyId                 *int64 `json:"partyId" msgpack:"partyId"`                                 // Key: 4
}

type StartTournamentPayload struct {
	TournamentDetailMasterId int64 `json:"tournamentDetailMasterId" msgpack:"tournamentDetailMasterId"` // Key: 0
	PartyId                  int64 `json:"partyId" msgpack:"partyId"`                                   // Key: 1
	StaminaConsumptionRatio  int   `json:"staminaConsumptionRatio" msgpack:"staminaConsumptionRatio"`   // Key: 2
}

type EditBookmarkPayload struct {
	MusicMasterId int64              `json:"musicMasterId" msgpack:"musicMasterId"` // Key: 0
	BookmarkFlag  MusicBookmarkFlags `json:"bookmarkFlag" msgpack:"bookmarkFlag"`   // Key: 1
}

type StartTripleCastLivePayload struct {
	MLiveId       int64 `json:"mLiveId" msgpack:"mLiveId"`             // Key: 0
	MTripleCastId int64 `json:"mTripleCastId" msgpack:"mTripleCastId"` // Key: 1
}

type StartTripleCastLiveResult struct {
	LiveUnits []any `json:"liveUnits" msgpack:"liveUnits"` // Key: 0
}

type SelectMusicCourseRandomMusicPayload struct {
	MusicCourseMasterId  int64                `json:"musicCourseMasterId" msgpack:"musicCourseMasterId"`   // Key: 0
	MusicCourseGaugeType MusicCourseGaugeType `json:"musicCourseGaugeType" msgpack:"musicCourseGaugeType"` // Key: 1
}

type StartMultiRoomLivePayload struct {
	HashedMultiRoomId string `json:"hashedMultiRoomId" msgpack:"hashedMultiRoomId"` // Key: 0
	LiveMasterId      int64  `json:"liveMasterId" msgpack:"liveMasterId"`           // Key: 1
	PartyId           *int64 `json:"partyId" msgpack:"partyId"`                     // Key: 2
}

type LoginPayload struct {
	PushNotificationToken string `json:"pushNotificationToken" msgpack:"pushNotificationToken"` // Key: 0
}

type SearchConditionMultiRoomPayload struct {
	HashedRoomId string `json:"hashedRoomId" msgpack:"hashedRoomId"` // Key: 0
	LiveMasterId *int64 `json:"liveMasterId" msgpack:"liveMasterId"` // Key: 1
}

type CreateMultiRoomPayload struct {
	RoomName         string             `json:"roomName" msgpack:"roomName"`                 // Key: 0
	Password         string             `json:"password" msgpack:"password"`                 // Key: 1
	CanJoinMaxMember int                `json:"canJoinMaxMember" msgpack:"canJoinMaxMember"` // Key: 2
	StartDate        time.Time          `json:"startDate" msgpack:"startDate"`               // Key: 3
	EndDate          time.Time          `json:"endDate" msgpack:"endDate"`                   // Key: 4
	PlayMode         MultiRoomPlayModes `json:"playMode" msgpack:"playMode"`                 // Key: 5
	LiveMasterId1    int64              `json:"liveMasterId1" msgpack:"liveMasterId1"`       // Key: 6
	LiveMasterId2    *int64             `json:"liveMasterId2" msgpack:"liveMasterId2"`       // Key: 7
	LiveMasterId3    *int64             `json:"liveMasterId3" msgpack:"liveMasterId3"`       // Key: 8
}

type JoinMultiRoomPayload struct {
	HashedRoomId string `json:"hashedRoomId" msgpack:"hashedRoomId"` // Key: 0
	Password     string `json:"password" msgpack:"password"`         // Key: 1
}

type InviteMultiRoomPayload struct {
	HashedUserIds []string `json:"hashedUserIds" msgpack:"hashedUserIds"` // Key: 0
}

type ChangeNamePayload struct {
	Name string `json:"name" msgpack:"name"` // Key: 0
}

type EditPositionPayload struct {
	Slots []EditPositionSlotPayload `json:"slots" msgpack:"slots"` // Key: 0
}

type EditPositionSlotPayload struct {
	UPartySlotId int64 `json:"uPartySlotId" msgpack:"uPartySlotId"` // Key: 0
	Position     int   `json:"position" msgpack:"position"`         // Key: 1
}

type EditPartyPayload struct {
	PartySlots            []EditPartySlotPayload `json:"partySlots" msgpack:"partySlots"`                       // Key: 0
	SenseNotationMasterId *int64                 `json:"senseNotationMasterId" msgpack:"senseNotationMasterId"` // Key: 1
	MusicMasterId         *int64                 `json:"musicMasterId" msgpack:"musicMasterId"`                 // Key: 2
	VocalVersion          *int                   `json:"vocalVersion" msgpack:"vocalVersion"`                   // Key: 3
	IsHighScoreChallenge  bool                   `json:"isHighScoreChallenge" msgpack:"isHighScoreChallenge"`   // Key: 4
	LeaderPosition        *int                   `json:"leaderPosition" msgpack:"leaderPosition"`               // Key: 5
}

type EditPartySlotPayload struct {
	UPartySlotId            int64                   `json:"uPartySlotId" msgpack:"uPartySlotId"`                       // Key: 0
	UCharacterId            int64                   `json:"uCharacterId" msgpack:"uCharacterId"`                       // Key: 1
	UAccessoryId            *int64                  `json:"uAccessoryId" msgpack:"uAccessoryId"`                       // Key: 2
	UPosterId               *int64                  `json:"uPosterId" msgpack:"uPosterId"`                             // Key: 3
	BonusAbilityEnableFlags BonusAbilityEnableFlags `json:"bonusAbilityEnableFlags" msgpack:"bonusAbilityEnableFlags"` // Key: 4
}

type RegisterAppStorePaymentPayload struct {
	ReceiptBase64 string `json:"receiptBase64" msgpack:"receiptBase64"` // Key: 0
}

type RegisterGooglePlayPaymentPayload struct {
	Json string `json:"json" msgpack:"json"` // Key: 0
}

type GeneratePhotosPayload struct {
	Payloads []GeneratePhotoPayload `json:"payloads" msgpack:"payloads"` // Key: 0
}

type GeneratePhotoPayload struct {
	CharacterBaseMasterId *int64                   `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 0
	ItemMasterId          int64                    `json:"itemMasterId" msgpack:"itemMasterId"`                   // Key: 1
	ImageData             []byte                   `json:"imageData" msgpack:"imageData"`                         // Key: 2
	ThumbnailImageData    []byte                   `json:"thumbnailImageData" msgpack:"thumbnailImageData"`       // Key: 3
	AppearedCharacters    []PhotoAppearedCharacter `json:"appearedCharacters" msgpack:"appearedCharacters"`       // Key: 4
}

type LevelUpPhotoPayload struct {
	PhotoId    int64 `json:"photoId" msgpack:"photoId"`       // Key: 0
	AfterLevel int   `json:"afterLevel" msgpack:"afterLevel"` // Key: 1
}

type AlbumArrangingPayload struct {
	Publishing    bool   `json:"publishing" msgpack:"publishing"`       // Key: 0
	Page          int    `json:"page" msgpack:"page"`                   // Key: 1
	Items         []byte `json:"items" msgpack:"items"`                 // Key: 2
	MAlbumThemeId *int64 `json:"mAlbumThemeId" msgpack:"mAlbumThemeId"` // Key: 3
}

type SetAlbumPublishingPayload struct {
	Publishing        bool `json:"publishing" msgpack:"publishing"`               // Key: 0
	PublishPageNumber int  `json:"publishPageNumber" msgpack:"publishPageNumber"` // Key: 1
}

type DetailEditItemPayload struct {
	ItemId        int64          `json:"itemId" msgpack:"itemId"`               // Key: 0
	AlbumItemType AlbumItemTypes `json:"albumItemType" msgpack:"albumItemType"` // Key: 1
	X             int            `json:"x" msgpack:"x"`                         // Key: 2
	Y             int            `json:"y" msgpack:"y"`                         // Key: 3
	Rotate        int            `json:"rotate" msgpack:"rotate"`               // Key: 4
	Scale         float32        `json:"scale" msgpack:"scale"`                 // Key: 5
	Layer         int            `json:"layer" msgpack:"layer"`                 // Key: 6
	Visible       bool           `json:"visible" msgpack:"visible"`             // Key: 7
	Lock          bool           `json:"lock" msgpack:"lock"`                   // Key: 8
}

type AlbumPhotoPayload struct {
	PhotoId int64 `json:"photoId" msgpack:"photoId"` // Key: 0
	Order   int   `json:"order" msgpack:"order"`     // Key: 1
}

type PhotoAppearedCharacter struct {
	CharacterBaseMasterId int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 0
	IsMainCharacter       bool  `json:"isMainCharacter" msgpack:"isMainCharacter"`             // Key: 1
}

type SetPhotoTagPayload struct {
	PhotoId          int64   `json:"photoId" msgpack:"photoId"`                   // Key: 0
	CharacterBaseIds []int64 `json:"characterBaseIds" msgpack:"characterBaseIds"` // Key: 1
}

type ChangePhotoAbilityPayload struct {
	PhotoId                      int64 `json:"photoId" msgpack:"photoId"`                                           // Key: 0
	ItemMasterId                 int64 `json:"itemMasterId" msgpack:"itemMasterId"`                                 // Key: 1
	PhotoEffectTypeGroupMasterId int64 `json:"photoEffectTypeGroupMasterId" msgpack:"photoEffectTypeGroupMasterId"` // Key: 2
}

type AbilityVarietyUpPayload struct {
	PhotoId      int64 `json:"photoId" msgpack:"photoId"`           // Key: 0
	VarietyTo    int   `json:"varietyTo" msgpack:"varietyTo"`       // Key: 1
	ItemMasterId int64 `json:"itemMasterId" msgpack:"itemMasterId"` // Key: 2
}

type UpdateHomeDisplayPreferencePayload struct {
	HomeCharacterBaseMasterId       *int64                    `json:"homeCharacterBaseMasterId" msgpack:"homeCharacterBaseMasterId"`             // Key: 0
	MemberCharacterBaseMasterId     *int64                    `json:"memberCharacterBaseMasterId" msgpack:"memberCharacterBaseMasterId"`         // Key: 2
	StoryCharacterBaseMasterId      *int64                    `json:"storyCharacterBaseMasterId" msgpack:"storyCharacterBaseMasterId"`           // Key: 3
	ShopCharacterBaseMasterId       *int64                    `json:"shopCharacterBaseMasterId" msgpack:"shopCharacterBaseMasterId"`             // Key: 4
	HomeCostumeMasterId             *int64                    `json:"homeCostumeMasterId" msgpack:"homeCostumeMasterId"`                         // Key: 5
	MemberCostumeMasterId           *int64                    `json:"memberCostumeMasterId" msgpack:"memberCostumeMasterId"`                     // Key: 7
	StoryCostumeMasterId            *int64                    `json:"storyCostumeMasterId" msgpack:"storyCostumeMasterId"`                       // Key: 8
	ShopCostumeMasterId             *int64                    `json:"shopCostumeMasterId" msgpack:"shopCostumeMasterId"`                         // Key: 9
	IllustCharacterMasterId         int64                     `json:"illustCharacterMasterId" msgpack:"illustCharacterMasterId"`                 // Key: 10
	DisplayAwakeningStatus          bool                      `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"`                   // Key: 11
	HomeCharacterDisplayType        HomeCharacterDisplayTypes `json:"homeCharacterDisplayType" msgpack:"homeCharacterDisplayType"`               // Key: 12
	LoginBonusCharacterBaseMasterId *int64                    `json:"loginBonusCharacterBaseMasterId" msgpack:"loginBonusCharacterBaseMasterId"` // Key: 13
	LoginBonusSpineCostumeMasterId  *int64                    `json:"loginBonusSpineCostumeMasterId" msgpack:"loginBonusSpineCostumeMasterId"`   // Key: 14
}

type UpdateTutorialPayload struct {
	TutorialStatus TutorialStatus `json:"tutorialStatus" msgpack:"tutorialStatus"` // Key: 0
}

type UpdateGameHintPayload struct {
	Categories []PageCategories `json:"categories" msgpack:"categories"` // Key: 0
}

type CostumeFavoritePayload struct {
	CostumeMasterId       int64 `json:"costumeMasterId" msgpack:"costumeMasterId"`             // Key: 0
	CharacterBaseMasterId int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	SetFavorite           bool  `json:"setFavorite" msgpack:"setFavorite"`                     // Key: 2
}

type FavoriteStampOrderPayload struct {
	StampMasterIds []int64 `json:"stampMasterIds" msgpack:"stampMasterIds"` // Key: 0
}

type PosterFavoritePayload struct {
	PosterId    int64 `json:"posterId" msgpack:"posterId"`       // Key: 0
	SetFavorite bool  `json:"setFavorite" msgpack:"setFavorite"` // Key: 1
}

type PosterAlternativeImagePayload struct {
	PosterId int64 `json:"posterId" msgpack:"posterId"` // Key: 0
	Pattern  int   `json:"pattern" msgpack:"pattern"`   // Key: 1
}

type UseSerialCodePayload struct {
	UserHashId       string `json:"userHashId" msgpack:"userHashId"`             // Key: 0
	ConfirmationCode string `json:"confirmationCode" msgpack:"confirmationCode"` // Key: 1
	SerialCode       string `json:"serialCode" msgpack:"serialCode"`             // Key: 2
}

type PurchaseItemPayload struct {
	MJewelShopItemId int64 `json:"mJewelShopItemId" msgpack:"mJewelShopItemId"` // Key: 0
}

type UpdateLastViewedAtPayload struct {
	ViewedShopCategoryTypes ViewedShopCategoryTypes `json:"viewedShopCategoryTypes" msgpack:"viewedShopCategoryTypes"` // Key: 0
	ExchangeShopMasterId    *int64                  `json:"exchangeShopMasterId" msgpack:"exchangeShopMasterId"`       // Key: 1
}

type ExchangeShopThingPayload struct {
	MExchangeShopThingId int64 `json:"mExchangeShopThingId" msgpack:"mExchangeShopThingId"` // Key: 0
	Quantity             int   `json:"quantity" msgpack:"quantity"`                         // Key: 1
}

type EditUserProfilePayload struct {
	Name                   string `json:"name" msgpack:"name"`                                     // Key: 0
	Introduction           string `json:"introduction" msgpack:"introduction"`                     // Key: 1
	MainUCharacterId       int64  `json:"mainUCharacterId" msgpack:"mainUCharacterId"`             // Key: 2
	MNameplateId           *int64 `json:"mNameplateId" msgpack:"mNameplateId"`                     // Key: 3
	MNameColorId           int64  `json:"mNameColorId" msgpack:"mNameColorId"`                     // Key: 4
	MTrophyId1             *int64 `json:"mTrophyId1" msgpack:"mTrophyId1"`                         // Key: 5
	MTrophyId2             *int64 `json:"mTrophyId2" msgpack:"mTrophyId2"`                         // Key: 6
	MTrophyId3             *int64 `json:"mTrophyId3" msgpack:"mTrophyId3"`                         // Key: 7
	IsPublicPlayerRate     bool   `json:"isPublicPlayerRate" msgpack:"isPublicPlayerRate"`         // Key: 8
	DisplayAwakeningStatus bool   `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 9
	MainCharacterMasterId  int64  `json:"mainCharacterMasterId" msgpack:"mainCharacterMasterId"`   // Key: 10
	NameBaseColorMasterid  int64  `json:"nameBaseColorMasterid" msgpack:"nameBaseColorMasterid"`   // Key: 11
	IconFrameMasterId      int64  `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`           // Key: 12
	HomeSkinMasterId       int64  `json:"homeSkinMasterId" msgpack:"homeSkinMasterId"`             // Key: 13
}

type AccountDeletionResult struct {
	IsSuccess bool                       `json:"isSuccess" msgpack:"isSuccess"` // Key: 0
	Error     *AccountDeletionErrorTypes `json:"error" msgpack:"error"`         // Key: 1
}

type AccountRegistResult struct {
	Token     string                    `json:"token" msgpack:"token"`         // Key: 0
	ErrorType AccountRegisterErrorTypes `json:"errorType" msgpack:"errorType"` // Key: 1
}

type AlbumPageSearchResult struct {
	AlbumPageResult *AlbumPageResult `json:"albumPageResult" msgpack:"albumPageResult"` // Key: 0
	IsSuccess       bool             `json:"isSuccess" msgpack:"isSuccess"`             // Key: 1
}

type AlbumPageResult struct {
	AlbumPageId        int64     `json:"albumPageId" msgpack:"albumPageId"`               // Key: 0
	AlbumId            int64     `json:"albumId" msgpack:"albumId"`                       // Key: 1
	Level              int       `json:"level" msgpack:"level"`                           // Key: 2
	Page               int       `json:"page" msgpack:"page"`                             // Key: 3
	EditType           EditTypes `json:"editType" msgpack:"editType"`                     // Key: 4
	Publishing         bool      `json:"publishing" msgpack:"publishing"`                 // Key: 5
	Items              []byte    `json:"items" msgpack:"items"`                           // Key: 6
	AlbumThemeMasterId *int64    `json:"albumThemeMasterId" msgpack:"albumThemeMasterId"` // Key: 7
}

type AuditionClearedInformationResult struct {
	PartiesPhases []AuditionClearedInformationParties `json:"partiesPhases" msgpack:"partiesPhases"` // Key: 0
}

type AuditionClearedInformationParties struct {
	Parties      []AuditionClearParty `json:"parties" msgpack:"parties"`           // Key: 0
	ClearedPhase byte                 `json:"clearedPhase" msgpack:"clearedPhase"` // Key: 1
}

type AuditionClearParty struct {
	Id             int64                    `json:"id" msgpack:"id"`                         // Key: 0
	Score          int64                    `json:"score" msgpack:"score"`                   // Key: 1
	Difficulty     MusicDifficulties        `json:"difficulty" msgpack:"difficulty"`         // Key: 2
	Slots          []AuditionClearPartySlot `json:"slots" msgpack:"slots"`                   // Key: 3
	UserName       string                   `json:"userName" msgpack:"userName"`             // Key: 4
	RateGrade      AchievementRateGrades    `json:"rateGrade" msgpack:"rateGrade"`           // Key: 5
	PlayerRank     string                   `json:"playerRank" msgpack:"playerRank"`         // Key: 9
	LeaderPosition int                      `json:"leaderPosition" msgpack:"leaderPosition"` // Key: 10
}

type AuditionClearPartySlot struct {
	Position                        int     `json:"position" msgpack:"position"`                                               // Key: 0
	CharacterMasterId               int64   `json:"characterMasterId" msgpack:"characterMasterId"`                             // Key: 1
	CharacterLevel                  int     `json:"characterLevel" msgpack:"characterLevel"`                                   // Key: 2
	CharacterTalentStage            int     `json:"characterTalentStage" msgpack:"characterTalentStage"`                       // Key: 3
	CharacterAwakeningPhase         int     `json:"characterAwakeningPhase" msgpack:"characterAwakeningPhase"`                 // Key: 4
	PosterMasterId                  *int64  `json:"posterMasterId" msgpack:"posterMasterId"`                                   // Key: 5
	PosterLevel                     *int    `json:"posterLevel" msgpack:"posterLevel"`                                         // Key: 6
	PosterBreakthroughPhase         *int    `json:"posterBreakthroughPhase" msgpack:"posterBreakthroughPhase"`                 // Key: 7
	AccessoryMasterId               *int64  `json:"accessoryMasterId" msgpack:"accessoryMasterId"`                             // Key: 8
	AccessoryLevel                  *int    `json:"accessoryLevel" msgpack:"accessoryLevel"`                                   // Key: 9
	StarRank                        int     `json:"starRank" msgpack:"starRank"`                                               // Key: 10
	TalentStage                     int     `json:"talentStage" msgpack:"talentStage"`                                         // Key: 11
	AwakeningPhase                  int64   `json:"awakeningPhase" msgpack:"awakeningPhase"`                                   // Key: 12
	CurrentStatus                   *Status `json:"currentStatus" msgpack:"currentStatus"`                                     // Key: 13
	CharacterDisplayAwakeningStatus bool    `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 14
}

type AuthenticateResult struct {
	Token       string     `json:"token" msgpack:"token"`             // Key: 0
	BanLevel    BanLevels  `json:"banLevel" msgpack:"banLevel"`       // Key: 1
	WarnedUntil *time.Time `json:"warnedUntil" msgpack:"warnedUntil"` // Key: 2
}

type BooleanResult struct {
	IsSuccess bool `json:"isSuccess" msgpack:"isSuccess"` // Key: 0
}

type CharacterLineupResult struct {
	NormalProbabilities []CharacterRarityProbability `json:"normalProbabilities" msgpack:"normalProbabilities"` // Key: 0
	FixedProbabilities  []CharacterRarityProbability `json:"fixedProbabilities" msgpack:"fixedProbabilities"`   // Key: 1
	NormalLineupItems   []GachaLineupItemProbability `json:"normalLineupItems" msgpack:"normalLineupItems"`     // Key: 2
	FixedLineupItems    []GachaLineupItemProbability `json:"fixedLineupItems" msgpack:"fixedLineupItems"`       // Key: 3
}

type CharacterRarityProbability struct {
	Rarity      CharacterRarities `json:"rarity" msgpack:"rarity"`           // Key: 0
	Probability float64           `json:"probability" msgpack:"probability"` // Key: 1
}

type CharacterPointEventInformationResult struct {
	CurrentTotalPoint int64 `json:"currentTotalPoint" msgpack:"currentTotalPoint"` // Key: 0
	OverallRank       *int  `json:"overallRank" msgpack:"overallRank"`             // Key: 1
	CharacterRank     *int  `json:"characterRank" msgpack:"characterRank"`         // Key: 2
}

type CircleEventInformationResult struct {
	CirclePoint             *int64                             `json:"circlePoint" msgpack:"circlePoint"`                         // Key: 0
	UserPoint               int64                              `json:"userPoint" msgpack:"userPoint"`                             // Key: 1
	LastReceivedCirclePoint int64                              `json:"lastReceivedCirclePoint" msgpack:"lastReceivedCirclePoint"` // Key: 2
	MissionRefreshCount     int                                `json:"missionRefreshCount" msgpack:"missionRefreshCount"`         // Key: 3
	Progresses              []CircleEventCircleMissionProgress `json:"progresses" msgpack:"progresses"`                           // Key: 4
}

type CircleEventCircleMissionProgress struct {
	CircleEventMissionMasterId int64 `json:"circleEventMissionMasterId" msgpack:"circleEventMissionMasterId"` // Key: 0
	CurrentCount               int64 `json:"currentCount" msgpack:"currentCount"`                             // Key: 1
}

type CircleInformationResult struct {
	UCircleHashedId        string             `json:"uCircleHashedId" msgpack:"uCircleHashedId"`               // Key: 0
	Name                   string             `json:"name" msgpack:"name"`                                     // Key: 1
	Comment                string             `json:"comment" msgpack:"comment"`                               // Key: 2
	PlayTimeStartType      PlayTimeTypes      `json:"playTimeStartType" msgpack:"playTimeStartType"`           // Key: 3
	PlayTimeEndType        PlayTimeTypes      `json:"playTimeEndType" msgpack:"playTimeEndType"`               // Key: 4
	EntryType              EntryTypes         `json:"entryType" msgpack:"entryType"`                           // Key: 5
	MemberCount            int                `json:"memberCount" msgpack:"memberCount"`                       // Key: 6
	InviteId               int64              `json:"inviteId" msgpack:"inviteId"`                             // Key: 7
	ResultStatus           CircleResultStatus `json:"resultStatus" msgpack:"resultStatus"`                     // Key: 9
	CircleBanner           *CircleBanner      `json:"circleBanner" msgpack:"circleBanner"`                     // Key: 10
	MainCharacterMasterId  int64              `json:"mainCharacterMasterId" msgpack:"mainCharacterMasterId"`   // Key: 11
	DisplayAwakeningStatus bool               `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 12
	CompanyMasterId        *int64             `json:"companyMasterId" msgpack:"companyMasterId"`               // Key: 13
	CharacterBaseMasterId  *int64             `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`   // Key: 14
	IconFameMasterId       int64              `json:"iconFameMasterId" msgpack:"iconFameMasterId"`             // Key: 15
}

type MyCircleInformationResult struct {
	UCircleHashedId           string                     `json:"uCircleHashedId" msgpack:"uCircleHashedId"`                     // Key: 0
	Name                      string                     `json:"name" msgpack:"name"`                                           // Key: 1
	Comment                   string                     `json:"comment" msgpack:"comment"`                                     // Key: 2
	PlayTimeStartType         PlayTimeTypes              `json:"playTimeStartType" msgpack:"playTimeStartType"`                 // Key: 3
	PlayTimeEndType           PlayTimeTypes              `json:"playTimeEndType" msgpack:"playTimeEndType"`                     // Key: 4
	EntryType                 EntryTypes                 `json:"entryType" msgpack:"entryType"`                                 // Key: 5
	MemberCount               int                        `json:"memberCount" msgpack:"memberCount"`                             // Key: 6
	MyAuthority               CircleAuthorities          `json:"myAuthority" msgpack:"myAuthority"`                             // Key: 7
	ResultStatus              CircleResultStatus         `json:"resultStatus" msgpack:"resultStatus"`                           // Key: 8
	CircleBanner              *CircleBanner              `json:"circleBanner" msgpack:"circleBanner"`                           // Key: 9
	CompanyMasterId           *int64                     `json:"companyMasterId" msgpack:"companyMasterId"`                     // Key: 10
	CharacterBaseMasterId     *int64                     `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`         // Key: 11
	SupportCompanyInformation *SupportCompanyInformation `json:"supportCompanyInformation" msgpack:"supportCompanyInformation"` // Key: 12
	SupportCompany            Companies                  `json:"supportCompany" msgpack:"supportCompany"`                       // Key: 13
	DailyAddedSupportPoint    int                        `json:"dailyAddedSupportPoint" msgpack:"dailyAddedSupportPoint"`       // Key: 14
	IsPublishRanking          bool                       `json:"isPublishRanking" msgpack:"isPublishRanking"`                   // Key: 15
	StaminaLastReceivedAt     time.Time                  `json:"staminaLastReceivedAt" msgpack:"staminaLastReceivedAt"`         // Key: 16
	CircleAuthorityUpdateType CircleAuthorityUpdateTypes `json:"circleAuthorityUpdateType" msgpack:"circleAuthorityUpdateType"` // Key: 17
}

type SupportCompanyInformation struct {
	Sirius  *SupportCompanyLevelStatus `json:"sirius" msgpack:"sirius"`   // Key: 0
	Eden    *SupportCompanyLevelStatus `json:"eden" msgpack:"eden"`       // Key: 1
	Gingaza *SupportCompanyLevelStatus `json:"gingaza" msgpack:"gingaza"` // Key: 2
	Denki   *SupportCompanyLevelStatus `json:"denki" msgpack:"denki"`     // Key: 3
}

type SupportCompanyLevelStatus struct {
	Level               int                             `json:"level" msgpack:"level"`                             // Key: 0
	CurrentSupportPoint int                             `json:"currentSupportPoint" msgpack:"currentSupportPoint"` // Key: 1
	LastLevelUppedAt    time.Time                       `json:"lastLevelUppedAt" msgpack:"lastLevelUppedAt"`       // Key: 2
	LevelLimit          int                             `json:"levelLimit" msgpack:"levelLimit"`                   // Key: 3
	LevelLimitStatus    *SupportCompanyLevelLimitStatus `json:"levelLimitStatus" msgpack:"levelLimitStatus"`       // Key: 4
}

type SupportCompanyLevelLimitStatus struct {
	CurrentCoinQuantity int64                                  `json:"currentCoinQuantity" msgpack:"currentCoinQuantity"` // Key: 0
	Details             []SupportCompanyLevelLimitStatusDetail `json:"details" msgpack:"details"`                         // Key: 1
}

type SupportCompanyLevelLimitStatusDetail struct {
	Order           int   `json:"order" msgpack:"order"`                     // Key: 0
	CurrentQuantity int64 `json:"currentQuantity" msgpack:"currentQuantity"` // Key: 1
}

type CircleSearchResult struct {
	Parameters   []CircleMemberParameters `json:"parameters" msgpack:"parameters"`     // Key: 0
	ResultStatus CircleResultStatus       `json:"resultStatus" msgpack:"resultStatus"` // Key: 1
	CircleBanner *CircleBanner            `json:"circleBanner" msgpack:"circleBanner"` // Key: 2
}

type CircleSearchIdResult struct {
	Parameters   *CircleMemberParameters `json:"parameters" msgpack:"parameters"`     // Key: 0
	ResultStatus CircleResultStatus      `json:"resultStatus" msgpack:"resultStatus"` // Key: 1
}

type CircleResult struct {
	ResultStatus CircleResultStatus `json:"resultStatus" msgpack:"resultStatus"` // Key: 0
}

type CircleInviteResult struct {
	ResultStatus CircleResultStatus `json:"resultStatus" msgpack:"resultStatus"` // Key: 0
	InviteId     *int64             `json:"inviteId" msgpack:"inviteId"`         // Key: 1
}

type CreateCircleResult struct {
	ResultStatus   CircleResultStatus `json:"resultStatus" msgpack:"resultStatus"`     // Key: 0
	CircleHashedId string             `json:"circleHashedId" msgpack:"circleHashedId"` // Key: 1
}

type CircleAuthorityResult struct {
	ResultStatus CircleAuthorityResultStatus `json:"resultStatus" msgpack:"resultStatus"` // Key: 0
}

type CircleMemberParameters struct {
	UserId                 string             `json:"userId" msgpack:"userId"`                                 // Key: 0
	UserName               string             `json:"userName" msgpack:"userName"`                             // Key: 1
	PlayerRank             int                `json:"playerRank" msgpack:"playerRank"`                         // Key: 2
	LastLoginAt            time.Time          `json:"lastLoginAt" msgpack:"lastLoginAt"`                       // Key: 3
	Authority              *CircleAuthorities `json:"authority" msgpack:"authority"`                           // Key: 4
	MainUCharacter         *MainCharacter     `json:"mainUCharacter" msgpack:"mainUCharacter"`                 // Key: 5
	InviteId               *int64             `json:"inviteId" msgpack:"inviteId"`                             // Key: 6
	RequestId              *int64             `json:"requestId" msgpack:"requestId"`                           // Key: 7
	TrophyMasterId1        int64              `json:"trophyMasterId1" msgpack:"trophyMasterId1"`               // Key: 8
	TrophyMasterId2        int64              `json:"trophyMasterId2" msgpack:"trophyMasterId2"`               // Key: 9
	TrophyMasterId3        int64              `json:"trophyMasterId3" msgpack:"trophyMasterId3"`               // Key: 10
	MainCharacterMasterId  int64              `json:"mainCharacterMasterId" msgpack:"mainCharacterMasterId"`   // Key: 11
	DisplayAwakeningStatus bool               `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 12
	IconFameMasterId       int64              `json:"iconFameMasterId" msgpack:"iconFameMasterId"`             // Key: 16
}

type CircleBanner struct {
	PosterMasterId      int64   `json:"posterMasterId" msgpack:"posterMasterId"`           // Key: 0
	X1                  float32 `json:"x1" msgpack:"x1"`                                   // Key: 1
	Y1                  float32 `json:"y1" msgpack:"y1"`                                   // Key: 2
	X2                  float32 `json:"x2" msgpack:"x2"`                                   // Key: 3
	Y2                  float32 `json:"y2" msgpack:"y2"`                                   // Key: 4
	RotationAngle       int     `json:"rotationAngle" msgpack:"rotationAngle"`             // Key: 5
	DisplayPosterString bool    `json:"displayPosterString" msgpack:"displayPosterString"` // Key: 6
	BannerRatio         float32 `json:"bannerRatio" msgpack:"bannerRatio"`                 // Key: 7
}

type CircleMemberInfoResult struct {
	Parameters   []CircleMemberInfoParameters `json:"parameters" msgpack:"parameters"`     // Key: 0
	ResultStatus CircleResultStatus           `json:"resultStatus" msgpack:"resultStatus"` // Key: 1
	CircleBanner *CircleBanner                `json:"circleBanner" msgpack:"circleBanner"` // Key: 2
}

type CircleMemberInfoParameters struct {
	PlayerRate         *float64         `json:"playerRate" msgpack:"playerRate"`                 // Key: 13
	IsPublicPlayerRate bool             `json:"isPublicPlayerRate" msgpack:"isPublicPlayerRate"` // Key: 14
	LeagueClass        LeagueClassTypes `json:"leagueClass" msgpack:"leagueClass"`               // Key: 15
}

type DonateSupportCompanyResult struct {
	MyCircleInformation *MyCircleInformationResult       `json:"myCircleInformation" msgpack:"myCircleInformation"` // Key: 0
	ResultStatus        CircleDonateSupportCompanyResult `json:"resultStatus" msgpack:"resultStatus"`               // Key: 1
}

type ConcoursInfomationResult struct {
	Details      []ConcoursDetailInformationResult `json:"details" msgpack:"details"`           // Key: 0
	CurrentPoint int                               `json:"currentPoint" msgpack:"currentPoint"` // Key: 1
}

type ConcoursDetailInformationResult struct {
	ConcoursDetailMasterId int64 `json:"concoursDetailMasterId" msgpack:"concoursDetailMasterId"` // Key: 0
	Rank                   *int  `json:"rank" msgpack:"rank"`                                     // Key: 1
}

type CourseRankingResult struct {
	Rank         int             `json:"rank" msgpack:"rank"`                 // Key: 0
	UserId       int64           `json:"userId" msgpack:"userId"`             // Key: 1
	UserProfile  *ApiUserProfile `json:"userProfile" msgpack:"userProfile"`   // Key: 2
	CourseResult *CourseResult   `json:"courseResult" msgpack:"courseResult"` // Key: 3
	NoteResult   *NoteResult     `json:"noteResult" msgpack:"noteResult"`     // Key: 4
}

type CourseResult struct {
	TotalAchievementRatePercentRecord float64   `json:"totalAchievementRatePercentRecord" msgpack:"totalAchievementRatePercentRecord"` // Key: 0
	BestRecordChallengeCount          int       `json:"bestRecordChallengeCount" msgpack:"bestRecordChallengeCount"`                   // Key: 1
	BestRecordDate                    time.Time `json:"bestRecordDate" msgpack:"bestRecordDate"`                                       // Key: 2
}

type NoteResult struct {
	PerfectStar int `json:"perfectStar" msgpack:"perfectStar"` // Key: 0
	Perfect     int `json:"perfect" msgpack:"perfect"`         // Key: 1
	Great       int `json:"great" msgpack:"great"`             // Key: 2
	Good        int `json:"good" msgpack:"good"`               // Key: 3
	Bad         int `json:"bad" msgpack:"bad"`                 // Key: 4
	Miss        int `json:"miss" msgpack:"miss"`               // Key: 5
}

type DebugLinkageCodeResult struct {
	Password           string              `json:"password" msgpack:"password"`                     // Key: 0
	TakeOverCodeResult *TakeOverCodeResult `json:"takeOverCodeResult" msgpack:"takeOverCodeResult"` // Key: 1
}

type DebugUserIdResult struct {
	UserId       string `json:"userId" msgpack:"userId"`             // Key: 0
	HashedUserId string `json:"hashedUserId" msgpack:"hashedUserId"` // Key: 1
}

type EexternalPaymentResult struct {
	ReceivedJewelShopItemMasterIds []int64 `json:"receivedJewelShopItemMasterIds" msgpack:"receivedJewelShopItemMasterIds"` // Key: 0
}

type EnvironmentResult struct {
	ApplicationVersion     string `json:"applicationVersion" msgpack:"applicationVersion"`         // Key: 0
	AssetVersion           string `json:"assetVersion" msgpack:"assetVersion"`                     // Key: 1
	ApiEndpoint            string `json:"apiEndpoint" msgpack:"apiEndpoint"`                       // Key: 2
	MaintenanceApiEndpoint string `json:"maintenanceApiEndpoint" msgpack:"maintenanceApiEndpoint"` // Key: 3
	NewsApiEndpoint        string `json:"newsApiEndpoint" msgpack:"newsApiEndpoint"`               // Key: 4
	IsMaintenance          bool   `json:"isMaintenance" msgpack:"isMaintenance"`                   // Key: 5
	MasterDataUrl          string `json:"masterDataUrl" msgpack:"masterDataUrl"`                   // Key: 6
	StaticContentUrl       string `json:"staticContentUrl" msgpack:"staticContentUrl"`             // Key: 7
	AssetUrl               string `json:"assetUrl" msgpack:"assetUrl"`                             // Key: 8
	IsAppReview            bool   `json:"isAppReview" msgpack:"isAppReview"`                       // Key: 9
	PhotoContentUrl        string `json:"photoContentUrl" msgpack:"photoContentUrl"`               // Key: 10
	MultiRealTimeServerUrl string `json:"multiRealTimeServerUrl" msgpack:"multiRealTimeServerUrl"` // Key: 11
	ExternalPaymentUrl     string `json:"externalPaymentUrl" msgpack:"externalPaymentUrl"`         // Key: 12
}

type EpisodeDetailResult struct {
	Id                               int64                                `json:"id" msgpack:"id"`                                                             // Key: 0
	EpisodeMasterId                  int64                                `json:"episodeMasterId" msgpack:"episodeMasterId"`                                   // Key: 1
	Order                            int                                  `json:"order" msgpack:"order"`                                                       // Key: 2
	GroupOrder                       int                                  `json:"groupOrder" msgpack:"groupOrder"`                                             // Key: 3
	Effect                           string                               `json:"effect" msgpack:"effect"`                                                     // Key: 4
	SpeakerName                      string                               `json:"speakerName" msgpack:"speakerName"`                                           // Key: 5
	FontSize                         FontSizes                            `json:"fontSize" msgpack:"fontSize"`                                                 // Key: 6
	Phrase                           string                               `json:"phrase" msgpack:"phrase"`                                                     // Key: 7
	Title                            string                               `json:"title" msgpack:"title"`                                                       // Key: 8
	BackgroundImageFileName          string                               `json:"backgroundImageFileName" msgpack:"backgroundImageFileName"`                   // Key: 9
	BackgroundCharacterImageFileName string                               `json:"backgroundCharacterImageFileName" msgpack:"backgroundCharacterImageFileName"` // Key: 10
	BackgroundImageFileFadeType      *FadeTypes                           `json:"backgroundImageFileFadeType" msgpack:"backgroundImageFileFadeType"`           // Key: 11
	FadeValue1                       *float32                             `json:"fadeValue1" msgpack:"fadeValue1"`                                             // Key: 21
	FadeValue2                       *float32                             `json:"fadeValue2" msgpack:"fadeValue2"`                                             // Key: 22
	FadeValue3                       *float32                             `json:"fadeValue3" msgpack:"fadeValue3"`                                             // Key: 23
	BgmFileName                      string                               `json:"bgmFileName" msgpack:"bgmFileName"`                                           // Key: 12
	SeFileName                       string                               `json:"seFileName" msgpack:"seFileName"`                                             // Key: 13
	StillPhotoFileName               string                               `json:"stillPhotoFileName" msgpack:"stillPhotoFileName"`                             // Key: 14
	MovieFileName                    string                               `json:"movieFileName" msgpack:"movieFileName"`                                       // Key: 15
	WindowEffect                     *WindowEffects                       `json:"windowEffect" msgpack:"windowEffect"`                                         // Key: 16
	SceneCameraMasterId              *int64                               `json:"sceneCameraMasterId" msgpack:"sceneCameraMasterId"`                           // Key: 17
	VoiceFileName                    string                               `json:"voiceFileName" msgpack:"voiceFileName"`                                       // Key: 18
	CharacterMotions                 []EpisodeDetailCharacterMotionResult `json:"characterMotions" msgpack:"characterMotions"`                                 // Key: 19
	SpeakerIconId                    string                               `json:"speakerIconId" msgpack:"speakerIconId"`                                       // Key: 20
}

type EpisodeDetailCharacterMotionResult struct {
	SlotNumber               int                       `json:"slotNumber" msgpack:"slotNumber"`                             // Key: 0
	FacialExpressionMasterId *int64                    `json:"facialExpressionMasterId" msgpack:"facialExpressionMasterId"` // Key: 1
	HeadMotionMasterId       *int64                    `json:"headMotionMasterId" msgpack:"headMotionMasterId"`             // Key: 2
	HeadDirectionMasterId    *int64                    `json:"headDirectionMasterId" msgpack:"headDirectionMasterId"`       // Key: 3
	BodyMotionMasterId       *int64                    `json:"bodyMotionMasterId" msgpack:"bodyMotionMasterId"`             // Key: 4
	LipSyncMasterId          *int64                    `json:"lipSyncMasterId" msgpack:"lipSyncMasterId"`                   // Key: 5
	SpineId                  int64                     `json:"spineId" msgpack:"spineId"`                                   // Key: 6
	CharacterAppearanceType  *CharacterAppearanceTypes `json:"characterAppearanceType" msgpack:"characterAppearanceType"`   // Key: 7
	CharacterPosition        CharacterPositions        `json:"characterPosition" msgpack:"characterPosition"`               // Key: 8
	CharacterLayerType       CharacterLayerTypes       `json:"characterLayerType" msgpack:"characterLayerType"`             // Key: 9
	SpineSize                SpineSizes                `json:"spineSize" msgpack:"spineSize"`                               // Key: 10
}

type EpisodeResult struct {
	EpisodeTitle             string     `json:"episodeTitle" msgpack:"episodeTitle"`                         // Key: 0
	StoryType                StoryTypes `json:"storyType" msgpack:"storyType"`                               // Key: 1
	EpisodeOrder             int        `json:"episodeOrder" msgpack:"episodeOrder"`                         // Key: 2
	EpisodeDetailAssetSource string     `json:"episodeDetailAssetSource" msgpack:"episodeDetailAssetSource"` // Key: 3
}

type EventBoxGachaRollResult struct {
	ReceivedThings                      []ReceivedThing `json:"receivedThings" msgpack:"receivedThings"`                                           // Key: 0
	HasReset                            bool            `json:"hasReset" msgpack:"hasReset"`                                                       // Key: 1
	DrawnEventBoxGachaBoxThingMasterIds []int64         `json:"drawnEventBoxGachaBoxThingMasterIds" msgpack:"drawnEventBoxGachaBoxThingMasterIds"` // Key: 2
}

type FinishLiveResult struct {
	LeagueRewards                  []ReceivedThing                 `json:"leagueRewards" msgpack:"leagueRewards"`                                   // Key: 1
	RateResult                     *RateResult                     `json:"rateResult" msgpack:"rateResult"`                                         // Key: 2
	PlayerRankPointResult          *PlayerRankPointResult          `json:"playerRankPointResult" msgpack:"playerRankPointResult"`                   // Key: 3
	ClearLamp                      ClearLamps                      `json:"clearLamp" msgpack:"clearLamp"`                                           // Key: 5
	RateGrade                      AchievementRateGrades           `json:"rateGrade" msgpack:"rateGrade"`                                           // Key: 6
	IsHighScore                    bool                            `json:"isHighScore" msgpack:"isHighScore"`                                       // Key: 8
	AuditionBeforePhase            byte                            `json:"auditionBeforePhase" msgpack:"auditionBeforePhase"`                       // Key: 9
	AuditionAfterPhase             byte                            `json:"auditionAfterPhase" msgpack:"auditionAfterPhase"`                         // Key: 10
	AuditionRewards                []ReceivedThing                 `json:"auditionRewards" msgpack:"auditionRewards"`                               // Key: 11
	StoryEventRewards              []ReceivedThing                 `json:"storyEventRewards" msgpack:"storyEventRewards"`                           // Key: 12
	AchievementRateAverage         float64                         `json:"achievementRateAverage" msgpack:"achievementRateAverage"`                 // Key: 13
	AuditionMasterId               *int64                          `json:"auditionMasterId" msgpack:"auditionMasterId"`                             // Key: 14
	SpRateUpdateResult             *SpRateUpdateResult             `json:"spRateUpdateResult" msgpack:"spRateUpdateResult"`                         // Key: 16
	TournamentResult               *TournamentResult               `json:"tournamentResult" msgpack:"tournamentResult"`                             // Key: 17
	CellingMaxCount                *int                            `json:"cellingMaxCount" msgpack:"cellingMaxCount"`                               // Key: 18
	CellingUserCount               *int                            `json:"cellingUserCount" msgpack:"cellingUserCount"`                             // Key: 19
	BeforeClearLamp                ClearLamps                      `json:"beforeClearLamp" msgpack:"beforeClearLamp"`                               // Key: 20
	PosterDropUpPercent            int                             `json:"posterDropUpPercent" msgpack:"posterDropUpPercent"`                       // Key: 21
	EventResult                    *EventResult                    `json:"eventResult" msgpack:"eventResult"`                                       // Key: 22
	LiveDropThings                 []LiveDropThing                 `json:"liveDropThings" msgpack:"liveDropThings"`                                 // Key: 23
	LessonResult                   *LessonResult                   `json:"lessonResult" msgpack:"lessonResult"`                                     // Key: 24
	LeagueResult                   *LeagueResult                   `json:"leagueResult" msgpack:"leagueResult"`                                     // Key: 25
	ConcertResult                  *ConcertResult                  `json:"concertResult" msgpack:"concertResult"`                                   // Key: 26
	BonusLiveResult                *BonusLiveResult                `json:"bonusLiveResult" msgpack:"bonusLiveResult"`                               // Key: 27
	GhostLiveResult                *GhostLiveResult                `json:"ghostLiveResult" msgpack:"ghostLiveResult"`                               // Key: 28
	TrialPartyEventResult          *TrialPartyEventResult          `json:"trialPartyEventResult" msgpack:"trialPartyEventResult"`                   // Key: 29
	AccessoryAutoSellConvertThings []AccessoryAutoSellConvertThing `json:"accessoryAutoSellConvertThings" msgpack:"accessoryAutoSellConvertThings"` // Key: 30
}

type PlayerRankPointResult struct {
	RankBefore        int `json:"rankBefore" msgpack:"rankBefore"`               // Key: 0
	RankAfter         int `json:"rankAfter" msgpack:"rankAfter"`                 // Key: 1
	RankPointBefore   int `json:"rankPointBefore" msgpack:"rankPointBefore"`     // Key: 2
	RankPointAfter    int `json:"rankPointAfter" msgpack:"rankPointAfter"`       // Key: 3
	RankPointAcquired int `json:"rankPointAcquired" msgpack:"rankPointAcquired"` // Key: 4
	StaminaBefore     int `json:"staminaBefore" msgpack:"staminaBefore"`         // Key: 5
}

type StarPointResult struct {
	RankBefore        int             `json:"rankBefore" msgpack:"rankBefore"`               // Key: 0
	RankAfter         int             `json:"rankAfter" msgpack:"rankAfter"`                 // Key: 1
	StarPointBefore   int             `json:"starPointBefore" msgpack:"starPointBefore"`     // Key: 2
	StarPointAfter    int             `json:"starPointAfter" msgpack:"starPointAfter"`       // Key: 3
	StarPointAcquired int             `json:"starPointAcquired" msgpack:"starPointAcquired"` // Key: 4
	ReceivedReward    []ReceivedThing `json:"receivedReward" msgpack:"receivedReward"`       // Key: 5
}

type CharacterBaseStarPointResult struct {
	CharacterBaseMasterId       int64            `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`             // Key: 0
	StarPointResult             *StarPointResult `json:"starPointResult" msgpack:"starPointResult"`                         // Key: 1
	LinkCharacterReceivedReward []ReceivedThing  `json:"linkCharacterReceivedReward" msgpack:"linkCharacterReceivedReward"` // Key: 2
}

type RateResult struct {
	AchievementRateResult *RateUpdateResult `json:"achievementRateResult" msgpack:"achievementRateResult"` // Key: 0
	LiveRateResult        *RateUpdateResult `json:"liveRateResult" msgpack:"liveRateResult"`               // Key: 1
	TotalRateBefore       float64           `json:"totalRateBefore" msgpack:"totalRateBefore"`             // Key: 2
	TotalRateAfter        float64           `json:"totalRateAfter" msgpack:"totalRateAfter"`               // Key: 3
}

type RateUpdateResult struct {
	BestEver float64 `json:"bestEver" msgpack:"bestEver"` // Key: 0
	ThisTime float64 `json:"thisTime" msgpack:"thisTime"` // Key: 1
}

type SpRateUpdateResult struct {
	BestEver      int `json:"bestEver" msgpack:"bestEver"`           // Key: 0
	ThisTime      int `json:"thisTime" msgpack:"thisTime"`           // Key: 1
	BestEverTotal int `json:"bestEverTotal" msgpack:"bestEverTotal"` // Key: 2
	ThisTimeTotal int `json:"thisTimeTotal" msgpack:"thisTimeTotal"` // Key: 3
}

type TournamentResult struct {
	BestEverUniqueScore      int64 `json:"bestEverUniqueScore" msgpack:"bestEverUniqueScore"`           // Key: 0
	ThisTimeUniqueScore      int64 `json:"thisTimeUniqueScore" msgpack:"thisTimeUniqueScore"`           // Key: 1
	TournamentDetailMasterId int64 `json:"tournamentDetailMasterId" msgpack:"tournamentDetailMasterId"` // Key: 2
}

type EventResult struct {
	AcquiredEventPoint                    *int   `json:"acquiredEventPoint" msgpack:"acquiredEventPoint"`                                       // Key: 0
	AttachedStoryEventMasterId            *int64 `json:"attachedStoryEventMasterId" msgpack:"attachedStoryEventMasterId"`                       // Key: 1
	CircleEventHighScore                  *int64 `json:"circleEventHighScore" msgpack:"circleEventHighScore"`                                   // Key: 2
	CircleEventHighScoreMultiplierPercent *int   `json:"circleEventHighScoreMultiplierPercent" msgpack:"circleEventHighScoreMultiplierPercent"` // Key: 3
	CircleEventBeforeHighScore            *int64 `json:"circleEventBeforeHighScore" msgpack:"circleEventBeforeHighScore"`                       // Key: 4
	HighScoreBefore                       *int64 `json:"highScoreBefore" msgpack:"highScoreBefore"`                                             // Key: 5
	HighScoreAfter                        *int64 `json:"highScoreAfter" msgpack:"highScoreAfter"`                                               // Key: 6
	BeforeRank                            int64  `json:"beforeRank" msgpack:"beforeRank"`                                                       // Key: 7
	AfterRank                             int64  `json:"afterRank" msgpack:"afterRank"`                                                         // Key: 8
	StoryEventMasterId                    *int64 `json:"storyEventMasterId" msgpack:"storyEventMasterId"`                                       // Key: 9
	EventMasterId                         *int64 `json:"eventMasterId" msgpack:"eventMasterId"`                                                 // Key: 10
	ThisTimeSupportPoint                  *int   `json:"thisTimeSupportPoint" msgpack:"thisTimeSupportPoint"`                                   // Key: 11
	CampBeforeRank                        *int64 `json:"campBeforeRank" msgpack:"campBeforeRank"`                                               // Key: 12
	CampAfterRank                         *int64 `json:"campAfterRank" msgpack:"campAfterRank"`                                                 // Key: 13
}

type LiveDropThing struct {
	ReceivedThing *ReceivedThing `json:"receivedThing" msgpack:"receivedThing"` // Key: 0
	Order         int64          `json:"order" msgpack:"order"`                 // Key: 1
	LiveDropType  LiveDropTypes  `json:"liveDropType" msgpack:"liveDropType"`   // Key: 2
	IsFixedDrop   bool           `json:"isFixedDrop" msgpack:"isFixedDrop"`     // Key: 3
}

type LessonResult struct {
	CharacterBaseMasterId int64            `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 0
	StarPointResult       *StarPointResult `json:"starPointResult" msgpack:"starPointResult"`             // Key: 1
	HighScoreRewards      []ReceivedThing  `json:"highScoreRewards" msgpack:"highScoreRewards"`           // Key: 2
	HighScoreBefore       int64            `json:"highScoreBefore" msgpack:"highScoreBefore"`             // Key: 3
	HighScoreAfter        int64            `json:"highScoreAfter" msgpack:"highScoreAfter"`               // Key: 4
}

type LeagueResult struct {
	BeforeGroupRanking    *int   `json:"beforeGroupRanking" msgpack:"beforeGroupRanking"`       // Key: 0
	AfterGroupRanking     int    `json:"afterGroupRanking" msgpack:"afterGroupRanking"`         // Key: 1
	LowestUpBorderScore   *int64 `json:"lowestUpBorderScore" msgpack:"lowestUpBorderScore"`     // Key: 2
	LowestKeepBorderScore *int64 `json:"lowestKeepBorderScore" msgpack:"lowestKeepBorderScore"` // Key: 3
	EnrolledGroupNumber   int    `json:"enrolledGroupNumber" msgpack:"enrolledGroupNumber"`     // Key: 4
	IsInCounting          bool   `json:"isInCounting" msgpack:"isInCounting"`                   // Key: 5
}

type ConcertResult struct {
	Rewards []ReceivedThing `json:"rewards" msgpack:"rewards"` // Key: 0
}

type BonusLiveResult struct {
	Rewards []ReceivedThing `json:"rewards" msgpack:"rewards"` // Key: 0
}

type GhostLiveResult struct {
	GhostUserProfile                *UserProfileDetail `json:"ghostUserProfile" msgpack:"ghostUserProfile"`                               // Key: 0
	GhostUserScore                  int64              `json:"ghostUserScore" msgpack:"ghostUserScore"`                                   // Key: 1
	TotalBattleWinCount             int64              `json:"totalBattleWinCount" msgpack:"totalBattleWinCount"`                         // Key: 2
	MatchingResult                  MatchingResult     `json:"matchingResult" msgpack:"matchingResult"`                                   // Key: 3
	DisplayCharacterMasterId        int64              `json:"displayCharacterMasterId" msgpack:"displayCharacterMasterId"`               // Key: 4
	DisplayCharacterAwakeningStatus bool               `json:"displayCharacterAwakeningStatus" msgpack:"displayCharacterAwakeningStatus"` // Key: 5
}

type TrialPartyEventResult struct {
	Rewards []ReceivedThing `json:"rewards" msgpack:"rewards"` // Key: 0
}

type AccessoryAutoSellConvertThing struct {
	AccessoryMasterId int64          `json:"accessoryMasterId" msgpack:"accessoryMasterId"` // Key: 0
	ConvertThing      *ReceivedThing `json:"convertThing" msgpack:"convertThing"`           // Key: 1
}

type FriendInvitationUserInfoResult struct {
	HashedUserId string                       `json:"hashedUserId" msgpack:"hashedUserId"` // Key: 0
	Name         string                       `json:"name" msgpack:"name"`                 // Key: 1
	PlayerRank   int64                        `json:"playerRank" msgpack:"playerRank"`     // Key: 2
	ResultStatus InvitationCodeResultStatuses `json:"resultStatus" msgpack:"resultStatus"` // Key: 3
}

type FriendListResult struct {
	Results            []FriendResult `json:"results" msgpack:"results"`                       // Key: 0
	CurrentFriendCount int            `json:"currentFriendCount" msgpack:"currentFriendCount"` // Key: 1
}

type FriendResult struct {
	UserId                 string           `json:"userId" msgpack:"userId"`                                 // Key: 0
	PlayerRank             int              `json:"playerRank" msgpack:"playerRank"`                         // Key: 2
	TrophyMasterId1        int64            `json:"trophyMasterId1" msgpack:"trophyMasterId1"`               // Key: 3
	TrophyMasterId2        int64            `json:"trophyMasterId2" msgpack:"trophyMasterId2"`               // Key: 4
	TrophyMasterId3        int64            `json:"trophyMasterId3" msgpack:"trophyMasterId3"`               // Key: 5
	Introduction           string           `json:"introduction" msgpack:"introduction"`                     // Key: 6
	LastLoggedInAt         time.Time        `json:"lastLoggedInAt" msgpack:"lastLoggedInAt"`                 // Key: 7
	Name                   string           `json:"name" msgpack:"name"`                                     // Key: 8
	PlayerRate             *float64         `json:"playerRate" msgpack:"playerRate"`                         // Key: 9
	IsPublicPlayerRate     bool             `json:"isPublicPlayerRate" msgpack:"isPublicPlayerRate"`         // Key: 10
	LeagueClass            LeagueClassTypes `json:"leagueClass" msgpack:"leagueClass"`                       // Key: 11
	CharacterRanks         []CharacterRank  `json:"characterRanks" msgpack:"characterRanks"`                 // Key: 12
	IsPublicAlbumMainPage  bool             `json:"isPublicAlbumMainPage" msgpack:"isPublicAlbumMainPage"`   // Key: 13
	MainCharacterMasterId  int64            `json:"mainCharacterMasterId" msgpack:"mainCharacterMasterId"`   // Key: 15
	DisplayAwakeningStatus bool             `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 16
	IconFrameMasterId      int64            `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`           // Key: 17
	IsFavorite             bool             `json:"isFavorite" msgpack:"isFavorite"`                         // Key: 18
}

type CharacterRank struct {
	MCharacterBaseId int64 `json:"mCharacterBaseId" msgpack:"mCharacterBaseId"` // Key: 0
	Rank             int   `json:"rank" msgpack:"rank"`                         // Key: 1
}

type FriendSearchResult struct {
	FriendResult *FriendResult            `json:"friendResult" msgpack:"friendResult"` // Key: 0
	ResultStatus FriendSearchResultStatus `json:"resultStatus" msgpack:"resultStatus"` // Key: 1
}

type FriendAcceptResult struct {
	ResultStatus FriendAcceptResultStatus `json:"resultStatus" msgpack:"resultStatus"` // Key: 0
}

type FriendRequestResult struct {
	ResultStatus FriendRequestResultStatus `json:"resultStatus" msgpack:"resultStatus"` // Key: 0
}

type MainCharacter struct {
	CharacterMasterId      int64 `json:"characterMasterId" msgpack:"characterMasterId"`           // Key: 0
	AwakeningPhase         int   `json:"awakeningPhase" msgpack:"awakeningPhase"`                 // Key: 1
	TalentStage            int   `json:"talentStage" msgpack:"talentStage"`                       // Key: 2
	Level                  int   `json:"level" msgpack:"level"`                                   // Key: 3
	DisplayAwakeningStatus bool  `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 4
}

type BlockListResult struct {
	Results []FriendResult `json:"results" msgpack:"results"` // Key: 0
}

type GachaHistoryResult struct {
	MasterId  int64     `json:"masterId" msgpack:"masterId"`   // Key: 0
	CreatedAt time.Time `json:"createdAt" msgpack:"createdAt"` // Key: 1
}

type GachaInfoResult struct {
	Id                  int64              `json:"id" msgpack:"id"`                                   // Key: 0
	RollLimits          []GachaRollLimit   `json:"rollLimits" msgpack:"rollLimits"`                   // Key: 1
	NormalEmissionFlags GachaEmissionFlags `json:"normalEmissionFlags" msgpack:"normalEmissionFlags"` // Key: 2
	FixedEmissionFlags  GachaEmissionFlags `json:"fixedEmissionFlags" msgpack:"fixedEmissionFlags"`   // Key: 3
}

type GachaRollLimit struct {
	GachaDetailMasterId int64 `json:"gachaDetailMasterId" msgpack:"gachaDetailMasterId"` // Key: 0
	RollLeft            int   `json:"rollLeft" msgpack:"rollLeft"`                       // Key: 1
}

type GachaLineupItemProbability struct {
	Id          int64   `json:"id" msgpack:"id"`                   // Key: 0
	Probability float64 `json:"probability" msgpack:"probability"` // Key: 2
}

type GachaRollResult struct {
	MGachaMasterId            int64              `json:"mGachaMasterId" msgpack:"mGachaMasterId"`                       // Key: 0
	ReceivedThings            []GachaThingResult `json:"receivedThings" msgpack:"receivedThings"`                       // Key: 1
	ReceivedBonusThings       []ReceivedThing    `json:"receivedBonusThings" msgpack:"receivedBonusThings"`             // Key: 2
	ReceivedDetailBonusThings []ReceivedThing    `json:"receivedDetailBonusThings" msgpack:"receivedDetailBonusThings"` // Key: 3
	ReceivedRollBonusThings   []ReceivedThing    `json:"receivedRollBonusThings" msgpack:"receivedRollBonusThings"`     // Key: 4
}

type GachaThingResult struct {
	ReceivedThings           []ReceivedThing `json:"receivedThings" msgpack:"receivedThings"`                     // Key: 0
	AdditionalReceivedThings []ReceivedThing `json:"additionalReceivedThings" msgpack:"additionalReceivedThings"` // Key: 1
	MovieCostumeMasterId     *int64          `json:"movieCostumeMasterId" msgpack:"movieCostumeMasterId"`         // Key: 2
}

type GachaSelectedThingsResult struct {
	SelectedGachaThingIds []int64 `json:"selectedGachaThingIds" msgpack:"selectedGachaThingIds"` // Key: 0
}

type GetPushNotificationTokenResult struct {
	PushNotificationToken string `json:"pushNotificationToken" msgpack:"pushNotificationToken"` // Key: 0
}

type InboxReceiveResult struct {
	ReceivedThings      []ReceivedThing `json:"receivedThings" msgpack:"receivedThings"`           // Key: 0
	HasNotReceiveThings bool            `json:"hasNotReceiveThings" msgpack:"hasNotReceiveThings"` // Key: 1
}

type LeaguePartyAndRankingResult struct {
	Rank           *int                    `json:"rank" msgpack:"rank"`                     // Key: 0
	UserId         string                  `json:"userId" msgpack:"userId"`                 // Key: 1
	BestScore      int64                   `json:"bestScore" msgpack:"bestScore"`           // Key: 2
	LeagueMasterId int64                   `json:"leagueMasterId" msgpack:"leagueMasterId"` // Key: 3
	ClassType      LeagueClassTypes        `json:"classType" msgpack:"classType"`           // Key: 4
	UserName       string                  `json:"userName" msgpack:"userName"`             // Key: 5
	ActingAbility  int                     `json:"actingAbility" msgpack:"actingAbility"`   // Key: 6
	Slots          []LeaguePartySlotResult `json:"slots" msgpack:"slots"`                   // Key: 7
	Difficulty     MusicDifficulties       `json:"difficulty" msgpack:"difficulty"`         // Key: 8
}

type LeaguePartySlotResult struct {
	SlotPosition                    int                 `json:"slotPosition" msgpack:"slotPosition"`                                       // Key: 0
	CharacterMasterId               int64               `json:"characterMasterId" msgpack:"characterMasterId"`                             // Key: 1
	Attribute                       Attributes          `json:"attribute" msgpack:"attribute"`                                             // Key: 2
	CharacterRarity                 CharacterRarities   `json:"characterRarity" msgpack:"characterRarity"`                                 // Key: 4
	CharacterLevel                  int                 `json:"characterLevel" msgpack:"characterLevel"`                                   // Key: 6
	PosterMasterId                  *int64              `json:"posterMasterId" msgpack:"posterMasterId"`                                   // Key: 7
	PosterRarity                    *PossessionRarities `json:"posterRarity" msgpack:"posterRarity"`                                       // Key: 8
	PosterLevel                     *int                `json:"posterLevel" msgpack:"posterLevel"`                                         // Key: 9
	PosterBreakthroughPhase         *int                `json:"posterBreakthroughPhase" msgpack:"posterBreakthroughPhase"`                 // Key: 10
	AccessoryMasterId               *int64              `json:"accessoryMasterId" msgpack:"accessoryMasterId"`                             // Key: 11
	AccessoryRarity                 *PossessionRarities `json:"accessoryRarity" msgpack:"accessoryRarity"`                                 // Key: 12
	AccessoryLevel                  *int                `json:"accessoryLevel" msgpack:"accessoryLevel"`                                   // Key: 13
	CurrentStatus                   *Status             `json:"currentStatus" msgpack:"currentStatus"`                                     // Key: 14
	Position                        int                 `json:"position" msgpack:"position"`                                               // Key: 15
	CharacterTalentStage            int                 `json:"characterTalentStage" msgpack:"characterTalentStage"`                       // Key: 16
	CharacterAwakeningPhase         int                 `json:"characterAwakeningPhase" msgpack:"characterAwakeningPhase"`                 // Key: 17
	CharacterDisplayAwakeningStatus bool                `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 18
}

type LeagueReceiveResults struct {
	ClassRewardReceivedThings        []ReceivedThing        `json:"classRewardReceivedThings" msgpack:"classRewardReceivedThings"`               // Key: 0
	ClassRewardType                  LeagueClassChangeTypes `json:"classRewardType" msgpack:"classRewardType"`                                   // Key: 1
	FirstAchieveReceivedThings       []ReceivedThing        `json:"firstAchieveReceivedThings" msgpack:"firstAchieveReceivedThings"`             // Key: 2
	GroupRankReceivedThings          []ReceivedThing        `json:"groupRankReceivedThings" msgpack:"groupRankReceivedThings"`                   // Key: 3
	IsFirstTimeEnrolledClass         bool                   `json:"isFirstTimeEnrolledClass" msgpack:"isFirstTimeEnrolledClass"`                 // Key: 4
	AllClassGlobalRankReceivedThings []ReceivedThing        `json:"allClassGlobalRankReceivedThings" msgpack:"allClassGlobalRankReceivedThings"` // Key: 5
}

type LeagueTopMenuInformationResult struct {
	DisplayStartAt         time.Time `json:"displayStartAt" msgpack:"displayStartAt"`                 // Key: 0
	CountingStartAt        time.Time `json:"countingStartAt" msgpack:"countingStartAt"`               // Key: 1
	DisplayEndAt           time.Time `json:"displayEndAt" msgpack:"displayEndAt"`                     // Key: 2
	MusicMasterId          int64     `json:"musicMasterId" msgpack:"musicMasterId"`                   // Key: 3
	LeagueGroupRanking     int       `json:"leagueGroupRanking" msgpack:"leagueGroupRanking"`         // Key: 6
	UpClassBorderScore     *int64    `json:"upClassBorderScore" msgpack:"upClassBorderScore"`         // Key: 8
	UpClassBorderRanking   *int      `json:"upClassBorderRanking" msgpack:"upClassBorderRanking"`     // Key: 9
	KeepClassBorderScore   *int64    `json:"keepClassBorderScore" msgpack:"keepClassBorderScore"`     // Key: 10
	KeepClassBorderRanking *int      `json:"keepClassBorderRanking" msgpack:"keepClassBorderRanking"` // Key: 11
	GroupLowestRank        *int      `json:"groupLowestRank" msgpack:"groupLowestRank"`               // Key: 12
	IsParticipated         bool      `json:"isParticipated" msgpack:"isParticipated"`                 // Key: 13
	EnrolledGroupNumber    *int      `json:"enrolledGroupNumber" msgpack:"enrolledGroupNumber"`       // Key: 14
}

type LoginBonusResult struct {
	LoginBonusId             int64                 `json:"loginBonusId" msgpack:"loginBonusId"`                         // Key: 0
	CurrentLoginCount        int                   `json:"currentLoginCount" msgpack:"currentLoginCount"`               // Key: 1
	ReceivedThings           []ReceivedThing       `json:"receivedThings" msgpack:"receivedThings"`                     // Key: 2
	Title                    string                `json:"title" msgpack:"title"`                                       // Key: 3
	Type                     LoginBonusTypes       `json:"type" msgpack:"type"`                                         // Key: 4
	StartDate                time.Time             `json:"startDate" msgpack:"startDate"`                               // Key: 5
	EndDate                  *time.Time            `json:"endDate" msgpack:"endDate"`                                   // Key: 6
	MessageTemplate          string                `json:"messageTemplate" msgpack:"messageTemplate"`                   // Key: 10
	BackgroundImagePath      string                `json:"backgroundImagePath" msgpack:"backgroundImagePath"`           // Key: 11
	CardboardImagePath       string                `json:"cardboardImagePath" msgpack:"cardboardImagePath"`             // Key: 12
	LogoImagePath            string                `json:"logoImagePath" msgpack:"logoImagePath"`                       // Key: 13
	TotalDays                int                   `json:"totalDays" msgpack:"totalDays"`                               // Key: 14
	IsLoop                   bool                  `json:"isLoop" msgpack:"isLoop"`                                     // Key: 15
	Order                    int                   `json:"order" msgpack:"order"`                                       // Key: 16
	Details                  []LoginBonusDetail    `json:"details" msgpack:"details"`                                   // Key: 17
	NavigationSpineId        string                `json:"navigationSpineId" msgpack:"navigationSpineId"`               // Key: 18
	VoiceAssetId             string                `json:"voiceAssetId" msgpack:"voiceAssetId"`                         // Key: 19
	HeadMotionMasterId       *int64                `json:"headMotionMasterId" msgpack:"headMotionMasterId"`             // Key: 20
	FacialExpressionMasterId *int64                `json:"facialExpressionMasterId" msgpack:"facialExpressionMasterId"` // Key: 21
	HeadDirectionMasterId    *int64                `json:"headDirectionMasterId" msgpack:"headDirectionMasterId"`       // Key: 22
	BodyMotionMasterId       *int64                `json:"bodyMotionMasterId" msgpack:"bodyMotionMasterId"`             // Key: 23
	LipSyncMasterId          *int64                `json:"lipSyncMasterId" msgpack:"lipSyncMasterId"`                   // Key: 24
	LayoutType               LoginBonusLayoutTypes `json:"layoutType" msgpack:"layoutType"`                             // Key: 25
	ComebackCampaignMasterId *int64                `json:"comebackCampaignMasterId" msgpack:"comebackCampaignMasterId"` // Key: 26
	LoginBonusSpineGroup     *LoginBonusSpineGroup `json:"loginBonusSpineGroup" msgpack:"loginBonusSpineGroup"`         // Key: 27
}

type LoginBonusDetail struct {
	ThingType     ThingTypes `json:"thingType" msgpack:"thingType"`         // Key: 0
	ThingId       int64      `json:"thingId" msgpack:"thingId"`             // Key: 1
	ThingQuantity int        `json:"thingQuantity" msgpack:"thingQuantity"` // Key: 2
	LoginCount    int        `json:"loginCount" msgpack:"loginCount"`       // Key: 3
}

type LoginBonusSpineGroup struct {
	SpineSelectType LoginBonusSpineSelectType `json:"spineSelectType" msgpack:"spineSelectType"` // Key: 0
	Details         []LoginBonusSpineDetail   `json:"details" msgpack:"details"`                 // Key: 1
}

type LoginBonusSpineDetail struct {
	TargetValue1             *int64 `json:"targetValue1" msgpack:"targetValue1"`                         // Key: 0
	TargetValue2             *int64 `json:"targetValue2" msgpack:"targetValue2"`                         // Key: 1
	MessageTemplate          string `json:"messageTemplate" msgpack:"messageTemplate"`                   // Key: 2
	NavigationSpineId        string `json:"navigationSpineId" msgpack:"navigationSpineId"`               // Key: 3
	VoiceAssetId             string `json:"voiceAssetId" msgpack:"voiceAssetId"`                         // Key: 4
	HeadMotionMasterId       *int64 `json:"headMotionMasterId" msgpack:"headMotionMasterId"`             // Key: 5
	FacialExpressionMasterId *int64 `json:"facialExpressionMasterId" msgpack:"facialExpressionMasterId"` // Key: 6
	HeadDirectionMasterId    *int64 `json:"headDirectionMasterId" msgpack:"headDirectionMasterId"`       // Key: 7
	BodyMotionMasterId       *int64 `json:"bodyMotionMasterId" msgpack:"bodyMotionMasterId"`             // Key: 8
	LipSyncMasterId          *int64 `json:"lipSyncMasterId" msgpack:"lipSyncMasterId"`                   // Key: 9
}

type LoginResult struct {
	InvalidedStarPasses             []StarPassTypes                 `json:"invalidedStarPasses" msgpack:"invalidedStarPasses"`                         // Key: 0
	LoginPassNotification           LoginPassNotificationTypes      `json:"loginPassNotification" msgpack:"loginPassNotification"`                     // Key: 1
	IsApproachingLoginPassInvalided bool                            `json:"isApproachingLoginPassInvalided" msgpack:"isApproachingLoginPassInvalided"` // Key: 2
	InvalidedItemMasterIds          []int64                         `json:"invalidedItemMasterIds" msgpack:"invalidedItemMasterIds"`                   // Key: 3
	ApproachingItemMasterIds        []int64                         `json:"approachingItemMasterIds" msgpack:"approachingItemMasterIds"`               // Key: 4
	StoryEventPointExchangeResult   []StoryEventPointExchangeResult `json:"storyEventPointExchangeResult" msgpack:"storyEventPointExchangeResult"`     // Key: 5
	InvalidedBuffItemMasterIds      []int64                         `json:"invalidedBuffItemMasterIds" msgpack:"invalidedBuffItemMasterIds"`           // Key: 6
}

type StoryEventPointExchangeResult struct {
	StoryEventMasterId          int64 `json:"storyEventMasterId" msgpack:"storyEventMasterId"`                   // Key: 0
	BeforeStoryEventPointAmount int   `json:"beforeStoryEventPointAmount" msgpack:"beforeStoryEventPointAmount"` // Key: 1
	AfterCoinAmount             int   `json:"afterCoinAmount" msgpack:"afterCoinAmount"`                         // Key: 2
}

type MarketThing struct {
	FrameNumber              int64 `json:"frameNumber" msgpack:"frameNumber"`                           // Key: 0
	MarketFrameThingMasterId int64 `json:"marketFrameThingMasterId" msgpack:"marketFrameThingMasterId"` // Key: 1
	HasPurchased             bool  `json:"hasPurchased" msgpack:"hasPurchased"`                         // Key: 4
	DiscountPercent          *int  `json:"discountPercent" msgpack:"discountPercent"`                   // Key: 5
}

type MarketResult struct {
	Things                  []MarketThing `json:"things" msgpack:"things"`                                   // Key: 0
	RequiredJewelForRefresh *int          `json:"requiredJewelForRefresh" msgpack:"requiredJewelForRefresh"` // Key: 1
}

type MasterDataManifest struct {
	Uri              string `json:"uri" msgpack:"uri"`                           // Key: 0
	SasToken         string `json:"sasToken" msgpack:"sasToken"`                 // Key: 1
	Version          string `json:"version" msgpack:"version"`                   // Key: 2
	PublishTimestamp int64  `json:"publishTimestamp" msgpack:"publishTimestamp"` // Key: 3
}

type MatchingGhostLiveResult struct {
	GhostLiveMasterId int64              `json:"ghostLiveMasterId" msgpack:"ghostLiveMasterId"` // Key: 0
	GhostUserProfile  *UserProfileDetail `json:"ghostUserProfile" msgpack:"ghostUserProfile"`   // Key: 1
	GhostLiveInfo     *GhostLiveInfo     `json:"ghostLiveInfo" msgpack:"ghostLiveInfo"`         // Key: 2
}

type UserProfileDetail struct {
	NameColorMasterId       int64  `json:"nameColorMasterId" msgpack:"nameColorMasterId"`             // Key: 9
	NameBaseColorMasterId   int64  `json:"nameBaseColorMasterId" msgpack:"nameBaseColorMasterId"`     // Key: 10
	NameplateMasterId       *int64 `json:"nameplateMasterId" msgpack:"nameplateMasterId"`             // Key: 11
	NameplateDetailMasterId *int64 `json:"nameplateDetailMasterId" msgpack:"nameplateDetailMasterId"` // Key: 12
}

type GhostLiveInfo struct {
	LeaderPosition int               `json:"leaderPosition" msgpack:"leaderPosition"` // Key: 0
	PartySlot      []PartySlotDetail `json:"partySlot" msgpack:"partySlot"`           // Key: 1
}

type PartySlotDetail struct {
	Position                        int    `json:"position" msgpack:"position"`                                               // Key: 0
	CharacterMasterId               int64  `json:"characterMasterId" msgpack:"characterMasterId"`                             // Key: 1
	CharacterLevel                  int    `json:"characterLevel" msgpack:"characterLevel"`                                   // Key: 2
	CharacterTalentStage            int    `json:"characterTalentStage" msgpack:"characterTalentStage"`                       // Key: 3
	CharacterAwakeningPhase         int    `json:"characterAwakeningPhase" msgpack:"characterAwakeningPhase"`                 // Key: 4
	CharacterDisplayAwakeningStatus bool   `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 5
	PosterMasterId                  *int64 `json:"posterMasterId" msgpack:"posterMasterId"`                                   // Key: 6
	PosterLevel                     *int   `json:"posterLevel" msgpack:"posterLevel"`                                         // Key: 7
	PosterBreakthroughPhase         *int   `json:"posterBreakthroughPhase" msgpack:"posterBreakthroughPhase"`                 // Key: 8
	AccessoryMasterId               *int64 `json:"accessoryMasterId" msgpack:"accessoryMasterId"`                             // Key: 9
	AccessoryLevel                  *int   `json:"accessoryLevel" msgpack:"accessoryLevel"`                                   // Key: 10
	CurrentStatusVocal              int    `json:"currentStatusVocal" msgpack:"currentStatusVocal"`                           // Key: 11
	CurrentStatusExpression         int    `json:"currentStatusExpression" msgpack:"currentStatusExpression"`                 // Key: 12
	CurrentStatusConcentration      int    `json:"currentStatusConcentration" msgpack:"currentStatusConcentration"`           // Key: 13
	UAccessoryId                    *int64 `json:"uAccessoryId" msgpack:"uAccessoryId"`                                       // Key: 14
}

type MissionPassRewardsResult struct {
	MissionPassRewards []ReceivedThing         `json:"missionPassRewards" msgpack:"missionPassRewards"` // Key: 0
	RewardResult       MissionPassRewardStatus `json:"rewardResult" msgpack:"rewardResult"`             // Key: 1
}

type MultiLiveRestrictionResult struct {
	RestrictionFinishedAt *time.Time `json:"restrictionFinishedAt" msgpack:"restrictionFinishedAt"` // Key: 0
}

type MultiRoomInformationResult struct {
	HashedMultiRoomId string             `json:"hashedMultiRoomId" msgpack:"hashedMultiRoomId"` // Key: 0
	RoomName          string             `json:"roomName" msgpack:"roomName"`                   // Key: 1
	PlayMode          MultiRoomPlayModes `json:"playMode" msgpack:"playMode"`                   // Key: 2
	JoinedMemberCount int                `json:"joinedMemberCount" msgpack:"joinedMemberCount"` // Key: 3
	CanJoinMaxMember  int                `json:"canJoinMaxMember" msgpack:"canJoinMaxMember"`   // Key: 4
	StartDate         time.Time          `json:"startDate" msgpack:"startDate"`                 // Key: 5
	EndDate           time.Time          `json:"endDate" msgpack:"endDate"`                     // Key: 6
	LiveMasterId1     int64              `json:"liveMasterId1" msgpack:"liveMasterId1"`         // Key: 7
	LiveMasterId2     *int64             `json:"liveMasterId2" msgpack:"liveMasterId2"`         // Key: 8
	LiveMasterId3     *int64             `json:"liveMasterId3" msgpack:"liveMasterId3"`         // Key: 9
	HasPassword       bool               `json:"hasPassword" msgpack:"hasPassword"`             // Key: 10
}

type MultiRoomJoinnedResult struct {
	MyRank                              int                         `json:"myRank" msgpack:"myRank"`                                                           // Key: 0
	MyScore                             *int64                      `json:"myScore" msgpack:"myScore"`                                                         // Key: 1
	MyTotalAchievementRatePercentRecord *float64                    `json:"myTotalAchievementRatePercentRecord" msgpack:"myTotalAchievementRatePercentRecord"` // Key: 2
	MultiRoomInformationResult          *MultiRoomInformationResult `json:"multiRoomInformationResult" msgpack:"multiRoomInformationResult"`                   // Key: 3
}

type MultiRoomInvitedResult struct {
	OwnerUserName              string                      `json:"ownerUserName" msgpack:"ownerUserName"`                           // Key: 0
	MultiRoomInformationResult *MultiRoomInformationResult `json:"multiRoomInformationResult" msgpack:"multiRoomInformationResult"` // Key: 1
}

type MultiRoomDetailResult struct {
	MultiRoomInformationResult *MultiRoomInformationResult `json:"multiRoomInformationResult" msgpack:"multiRoomInformationResult"` // Key: 0
	MultiRoomRankings          []MultiRoomRanking          `json:"multiRoomRankings" msgpack:"multiRoomRankings"`                   // Key: 1
	IsFinished                 bool                        `json:"isFinished" msgpack:"isFinished"`                                 // Key: 2
	Password                   string                      `json:"password" msgpack:"password"`                                     // Key: 3
}

type MultiRoomRanking struct {
	UserId                            string          `json:"userId" msgpack:"userId"`                                                       // Key: 0
	Rank                              int             `json:"rank" msgpack:"rank"`                                                           // Key: 1
	IsOwner                           bool            `json:"isOwner" msgpack:"isOwner"`                                                     // Key: 2
	Score                             *int64          `json:"score" msgpack:"score"`                                                         // Key: 3
	TotalAchievementRatePercentRecord *float64        `json:"totalAchievementRatePercentRecord" msgpack:"totalAchievementRatePercentRecord"` // Key: 4
	BestRecordChallengeCount          int             `json:"bestRecordChallengeCount" msgpack:"bestRecordChallengeCount"`                   // Key: 5
	BestRecordDate                    *time.Time      `json:"bestRecordDate" msgpack:"bestRecordDate"`                                       // Key: 6
	UserProfile                       *ApiUserProfile `json:"userProfile" msgpack:"userProfile"`                                             // Key: 7
	PartyInfo                         *PartyInfo      `json:"partyInfo" msgpack:"partyInfo"`                                                 // Key: 8
	NoteResult                        *NoteResult     `json:"noteResult" msgpack:"noteResult"`                                               // Key: 9
}

type PartyInfo struct {
	LeaderPosition      int                  `json:"leaderPosition" msgpack:"leaderPosition"`           // Key: 0
	MultiRoomPartySlots []MultiRoomPartySlot `json:"multiRoomPartySlots" msgpack:"multiRoomPartySlots"` // Key: 1
}

type MultiRoomPartySlot struct {
	Position                        int     `json:"position" msgpack:"position"`                                               // Key: 0
	CharacterMasterId               int64   `json:"characterMasterId" msgpack:"characterMasterId"`                             // Key: 1
	CharacterLevel                  int     `json:"characterLevel" msgpack:"characterLevel"`                                   // Key: 2
	CharacterTalentStage            int     `json:"characterTalentStage" msgpack:"characterTalentStage"`                       // Key: 3
	CharacterAwakeningPhase         int     `json:"characterAwakeningPhase" msgpack:"characterAwakeningPhase"`                 // Key: 4
	CharacterDisplayAwakeningStatus bool    `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 5
	PosterMasterId                  *int64  `json:"posterMasterId" msgpack:"posterMasterId"`                                   // Key: 6
	PosterLevel                     *int    `json:"posterLevel" msgpack:"posterLevel"`                                         // Key: 7
	PosterBreakthroughPhase         *int    `json:"posterBreakthroughPhase" msgpack:"posterBreakthroughPhase"`                 // Key: 8
	AccessoryMasterId               *int64  `json:"accessoryMasterId" msgpack:"accessoryMasterId"`                             // Key: 9
	AccessoryLevel                  *int    `json:"accessoryLevel" msgpack:"accessoryLevel"`                                   // Key: 10
	UAccessoryId                    *int64  `json:"uAccessoryId" msgpack:"uAccessoryId"`                                       // Key: 11
	CurrentStatus                   *Status `json:"currentStatus" msgpack:"currentStatus"`                                     // Key: 12
}

type MultiRoomCreateResult struct {
	HashedMultiRoomId string                 `json:"hashedMultiRoomId" msgpack:"hashedMultiRoomId"` // Key: 0
	RoomDetail        *MultiRoomDetailResult `json:"roomDetail" msgpack:"roomDetail"`               // Key: 1
}

type MusicCourseRandomSelectResult struct {
	Details []MusicCourseRandomSelectResultDetail `json:"details" msgpack:"details"` // Key: 0
}

type MusicCourseRandomSelectResultDetail struct {
	SetListNumber int   `json:"setListNumber" msgpack:"setListNumber"` // Key: 0
	LiveMasterId  int64 `json:"liveMasterId" msgpack:"liveMasterId"`   // Key: 1
}

type NotificationContentResult struct {
	Id                      int64                   `json:"id" msgpack:"id"`                                           // Key: 0
	Body                    string                  `json:"body" msgpack:"body"`                                       // Key: 1
	BannerPath              string                  `json:"bannerPath" msgpack:"bannerPath"`                           // Key: 7
	Title                   string                  `json:"title" msgpack:"title"`                                     // Key: 2
	PostingAt               time.Time               `json:"postingAt" msgpack:"postingAt"`                             // Key: 3
	LastUpdatedAt           time.Time               `json:"lastUpdatedAt" msgpack:"lastUpdatedAt"`                     // Key: 4
	NotificationTabCategory NotificationTabCategory `json:"notificationTabCategory" msgpack:"notificationTabCategory"` // Key: 5
	NotificationCategory    NotificationCategory    `json:"notificationCategory" msgpack:"notificationCategory"`       // Key: 6
}

type NotificationResult struct {
	Id                      int64                   `json:"id" msgpack:"id"`                                           // Key: 0
	BannerPath              string                  `json:"bannerPath" msgpack:"bannerPath"`                           // Key: 1
	Title                   string                  `json:"title" msgpack:"title"`                                     // Key: 2
	PostingAt               time.Time               `json:"postingAt" msgpack:"postingAt"`                             // Key: 3
	LastUpdatedAt           time.Time               `json:"lastUpdatedAt" msgpack:"lastUpdatedAt"`                     // Key: 4
	NotificationTabCategory NotificationTabCategory `json:"notificationTabCategory" msgpack:"notificationTabCategory"` // Key: 5
	NotificationCategory    NotificationCategory    `json:"notificationCategory" msgpack:"notificationCategory"`       // Key: 6
	IsConfirmation          bool                    `json:"isConfirmation" msgpack:"isConfirmation"`                   // Key: 7
	Order                   int                     `json:"order" msgpack:"order"`                                     // Key: 8
}

type GeneratePhotoResult struct {
	PhotoId             int64         `json:"photoId" msgpack:"photoId"`                         // Key: 0
	FileName            string        `json:"fileName" msgpack:"fileName"`                       // Key: 1
	SasToken            string        `json:"sasToken" msgpack:"sasToken"`                       // Key: 2
	Rarity              PhotoRarities `json:"rarity" msgpack:"rarity"`                           // Key: 3
	SignMasterId        *int64        `json:"signMasterId" msgpack:"signMasterId"`               // Key: 4
	PhotoEffectMasterId *int64        `json:"photoEffectMasterId" msgpack:"photoEffectMasterId"` // Key: 5
}

type LevelUpPhotoResult struct {
	PhotoId     int64 `json:"photoId" msgpack:"photoId"`         // Key: 0
	BeforeLevel int   `json:"beforeLevel" msgpack:"beforeLevel"` // Key: 1
	AfterLevel  int   `json:"afterLevel" msgpack:"afterLevel"`   // Key: 2
}

type PosterLineupResult struct {
	NormalProbabilities []PosterRarityProbability    `json:"normalProbabilities" msgpack:"normalProbabilities"` // Key: 0
	FixedProbabilities  []PosterRarityProbability    `json:"fixedProbabilities" msgpack:"fixedProbabilities"`   // Key: 1
	NormalLineupItems   []GachaLineupItemProbability `json:"normalLineupItems" msgpack:"normalLineupItems"`     // Key: 2
	FixedLineupItems    []GachaLineupItemProbability `json:"fixedLineupItems" msgpack:"fixedLineupItems"`       // Key: 3
}

type PosterRarityProbability struct {
	Rarity      PossessionRarities `json:"rarity" msgpack:"rarity"`           // Key: 0
	Probability float64            `json:"probability" msgpack:"probability"` // Key: 1
}

type ProcessPaymentResult struct {
	Result ProcessPaymentTransactionResult `json:"result" msgpack:"result"` // Key: 0
}

type RawRankingResult struct {
	RawRanking   []RawRanking     `json:"rawRanking" msgpack:"rawRanking"`     // Key: 0
	UserProfiles []ApiUserProfile `json:"userProfiles" msgpack:"userProfiles"` // Key: 1
}

type RawHighScoreRankingResult struct {
	RawRanking       []RawRankingWithLongPoint `json:"rawRanking" msgpack:"rawRanking"`             // Key: 0
	HighScoreParties []ApiHighScoreParty       `json:"highScoreParties" msgpack:"highScoreParties"` // Key: 1
	HighScoreBuffs   []HighScoreBuff           `json:"highScoreBuffs" msgpack:"highScoreBuffs"`     // Key: 2
	UserProfiles     []ApiUserProfile          `json:"userProfiles" msgpack:"userProfiles"`         // Key: 3
}

type TotalPointEventRankingResult struct {
	RawRanking   []RawRankingWithLongPoint `json:"rawRanking" msgpack:"rawRanking"`     // Key: 0
	UserProfiles []ApiUserProfile          `json:"userProfiles" msgpack:"userProfiles"` // Key: 1
}

type CharacterPointEventRankingResult struct {
	RawRanking   []RawRankingWithLongPoint `json:"rawRanking" msgpack:"rawRanking"`     // Key: 0
	UserProfiles []ApiUserProfile          `json:"userProfiles" msgpack:"userProfiles"` // Key: 1
}

type RawRanking struct {
	Rank       int    `json:"rank" msgpack:"rank"`             // Key: 0
	Point      int    `json:"point" msgpack:"point"`           // Key: 1
	UserId     string `json:"userId" msgpack:"userId"`         // Key: 2
	GroupValue int    `json:"groupValue" msgpack:"groupValue"` // Key: 3
}

type RawRankingWithLongPoint struct {
	Rank   int    `json:"rank" msgpack:"rank"`     // Key: 0
	Point  int64  `json:"point" msgpack:"point"`   // Key: 1
	UserId string `json:"userId" msgpack:"userId"` // Key: 2
}

type ApiUserProfile struct {
	UserId                          string `json:"userId" msgpack:"userId"`                                                   // Key: 0
	UserName                        string `json:"userName" msgpack:"userName"`                                               // Key: 1
	TrophyMasterId1                 *int64 `json:"trophyMasterId1" msgpack:"trophyMasterId1"`                                 // Key: 2
	TrophyMasterId2                 *int64 `json:"trophyMasterId2" msgpack:"trophyMasterId2"`                                 // Key: 3
	TrophyMasterId3                 *int64 `json:"trophyMasterId3" msgpack:"trophyMasterId3"`                                 // Key: 4
	MainMCharacterId                int64  `json:"mainMCharacterId" msgpack:"mainMCharacterId"`                               // Key: 5
	MainCharacterLevel              int    `json:"mainCharacterLevel" msgpack:"mainCharacterLevel"`                           // Key: 6
	CharacterDisplayAwakeningStatus bool   `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 7
	IconFrameMasterId               int64  `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`                             // Key: 8
}

type ApiHighScoreParty struct {
	UserId         string                  `json:"userId" msgpack:"userId"`                 // Key: 0
	UserName       string                  `json:"userName" msgpack:"userName"`             // Key: 1
	PartyName      string                  `json:"partyName" msgpack:"partyName"`           // Key: 2
	Slots          []ApiHighScorePartySlot `json:"slots" msgpack:"slots"`                   // Key: 3
	Difficulty     MusicDifficulties       `json:"difficulty" msgpack:"difficulty"`         // Key: 4
	LeaderPosition int                     `json:"leaderPosition" msgpack:"leaderPosition"` // Key: 5
}

type ApiHighScorePartySlot struct {
	Position                        int     `json:"position" msgpack:"position"`                                               // Key: 0
	CharacterMasterId               int64   `json:"characterMasterId" msgpack:"characterMasterId"`                             // Key: 1
	CharacterLevel                  int     `json:"characterLevel" msgpack:"characterLevel"`                                   // Key: 2
	CharacterTalentStage            int     `json:"characterTalentStage" msgpack:"characterTalentStage"`                       // Key: 3
	CharacterAwakeningPhase         int     `json:"characterAwakeningPhase" msgpack:"characterAwakeningPhase"`                 // Key: 4
	PosterMasterId                  *int64  `json:"posterMasterId" msgpack:"posterMasterId"`                                   // Key: 5
	PosterLevel                     *int    `json:"posterLevel" msgpack:"posterLevel"`                                         // Key: 6
	PosterBreakthroughPhase         *int    `json:"posterBreakthroughPhase" msgpack:"posterBreakthroughPhase"`                 // Key: 7
	AccessoryMasterId               *int64  `json:"accessoryMasterId" msgpack:"accessoryMasterId"`                             // Key: 8
	AccessoryLevel                  *int    `json:"accessoryLevel" msgpack:"accessoryLevel"`                                   // Key: 9
	CurrentStatus                   *Status `json:"currentStatus" msgpack:"currentStatus"`                                     // Key: 10
	CharacterDisplayAwakeningStatus bool    `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 11
}

type HighScoreBuff struct {
	UserId string                 `json:"userId" msgpack:"userId"` // Key: 0
	Buffs  []HighScoreBuffSetting `json:"buffs" msgpack:"buffs"`   // Key: 1
}

type HighScoreBuffSetting struct {
	StoryEventHighScoreBuffSettingId int64 `json:"storyEventHighScoreBuffSettingId" msgpack:"storyEventHighScoreBuffSettingId"` // Key: 0
	CurrentLevel                     int   `json:"currentLevel" msgpack:"currentLevel"`                                         // Key: 1
}

type CircleRawRanking struct {
	Rank     int    `json:"rank" msgpack:"rank"`         // Key: 0
	Point    int    `json:"point" msgpack:"point"`       // Key: 1
	CircleId string `json:"circleId" msgpack:"circleId"` // Key: 2
}

type CircleProfile struct {
	CircleId               string `json:"circleId" msgpack:"circleId"`                             // Key: 0
	Name                   string `json:"name" msgpack:"name"`                                     // Key: 1
	MemberCount            int    `json:"memberCount" msgpack:"memberCount"`                       // Key: 2
	MainCharacterMasterId  int64  `json:"mainCharacterMasterId" msgpack:"mainCharacterMasterId"`   // Key: 3
	DisplayAwakeningStatus bool   `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 4
	IconFrameMasterId      int64  `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`           // Key: 5
}

type CircleEventRanking struct {
	CircleRawRankings []CircleRawRanking `json:"circleRawRankings" msgpack:"circleRawRankings"` // Key: 0
	CircleProfiles    []CircleProfile    `json:"circleProfiles" msgpack:"circleProfiles"`       // Key: 1
}

type StoryEventCampInfo struct {
	RawRanking      *RawRanking `json:"rawRanking" msgpack:"rawRanking"`           // Key: 0
	Camp1TotalPoint int64       `json:"camp1TotalPoint" msgpack:"camp1TotalPoint"` // Key: 1
	Camp2TotalPoint int64       `json:"camp2TotalPoint" msgpack:"camp2TotalPoint"` // Key: 2
}

type ReceivedThing struct {
	Type         ThingTypes  `json:"type" msgpack:"type"`                 // Key: 0
	Id           *int64      `json:"id" msgpack:"id"`                     // Key: 1
	Quantity     int         `json:"quantity" msgpack:"quantity"`         // Key: 2
	OriginalType *ThingTypes `json:"originalType" msgpack:"originalType"` // Key: 3
	OriginalId   *int64      `json:"originalId" msgpack:"originalId"`     // Key: 4
	AfterPhase   *int        `json:"afterPhase" msgpack:"afterPhase"`     // Key: 5
	SentInbox    bool        `json:"sentInbox" msgpack:"sentInbox"`       // Key: 6
}

type RouletteRollResult struct {
	RollResults []RollResult `json:"rollResults" msgpack:"rollResults"` // Key: 0
}

type RollResult struct {
	RollOrder             int   `json:"rollOrder" msgpack:"rollOrder"`                         // Key: 0
	SelectedSlot          int   `json:"selectedSlot" msgpack:"selectedSlot"`                   // Key: 1
	RoulettePrizeMasterId int64 `json:"roulettePrizeMasterId" msgpack:"roulettePrizeMasterId"` // Key: 2
}

type ScoreWithDateResult struct {
	Score int64     `json:"score" msgpack:"score"` // Key: 0
	Date  time.Time `json:"date" msgpack:"date"`   // Key: 1
}

type SerialCodeResult struct {
	Result UseSerialCodeResult `json:"result" msgpack:"result"` // Key: 0
}

type ShopResult struct {
	Id               int64             `json:"id" msgpack:"id"`                             // Key: 0
	MExchangeShopId  int64             `json:"mExchangeShopId" msgpack:"mExchangeShopId"`   // Key: 1
	RequiredMItemId  int64             `json:"requiredMItemId" msgpack:"requiredMItemId"`   // Key: 2
	RequiredQuantity int               `json:"requiredQuantity" msgpack:"requiredQuantity"` // Key: 3
	Order            int               `json:"order" msgpack:"order"`                       // Key: 4
	ExchangeLimit    int               `json:"exchangeLimit" msgpack:"exchangeLimit"`       // Key: 5
	ExchangeTimes    int               `json:"exchangeTimes" msgpack:"exchangeTimes"`       // Key: 6
	Stock            int               `json:"stock" msgpack:"stock"`                       // Key: 7
	ReplaceType      *ShopReplaceTypes `json:"replaceType" msgpack:"replaceType"`           // Key: 8
	ReplaceValue     *int              `json:"replaceValue" msgpack:"replaceValue"`         // Key: 9
	IsDisplayLocking bool              `json:"isDisplayLocking" msgpack:"isDisplayLocking"` // Key: 10
	UnlockType       *ShopUnlockTypes  `json:"unlockType" msgpack:"unlockType"`             // Key: 11
	UnlockValue      *int64            `json:"unlockValue" msgpack:"unlockValue"`           // Key: 12
	StartDate        *time.Time        `json:"startDate" msgpack:"startDate"`               // Key: 13
	EndDate          *time.Time        `json:"endDate" msgpack:"endDate"`                   // Key: 14
	MMusicId         *int64            `json:"mMusicId" msgpack:"mMusicId"`                 // Key: 15
}

type StartMultiLiveResult struct {
	MultiLiveId int64 `json:"multiLiveId" msgpack:"multiLiveId"` // Key: 0
}

type StoryEventMissionCircleProgressResult struct {
	IsSuccess  bool                              `json:"isSuccess" msgpack:"isSuccess"`   // Key: 0
	Progresses []StoryEventMissionCircleProgress `json:"progresses" msgpack:"progresses"` // Key: 1
}

type StoryEventMissionCircleProgress struct {
	StoryEventCircleMissionMasterId int64 `json:"storyEventCircleMissionMasterId" msgpack:"storyEventCircleMissionMasterId"` // Key: 0
	CurrentCount                    int   `json:"currentCount" msgpack:"currentCount"`                                       // Key: 1
}

type TakeOverAccountResult struct {
	IsSuccess  bool   `json:"isSuccess" msgpack:"isSuccess"`   // Key: 0
	UserId     string `json:"userId" msgpack:"userId"`         // Key: 1
	Name       string `json:"name" msgpack:"name"`             // Key: 2
	Rank       int    `json:"rank" msgpack:"rank"`             // Key: 3
	LoginToken string `json:"loginToken" msgpack:"loginToken"` // Key: 4
}

type TakeOverCodeResult struct {
	IsSuccess   bool   `json:"isSuccess" msgpack:"isSuccess"`     // Key: 0
	LinkageCode string `json:"linkageCode" msgpack:"linkageCode"` // Key: 1
}

type TeamChallengeGoalLotResult struct {
	TeamChallengeMasterId int64 `json:"teamChallengeMasterId" msgpack:"teamChallengeMasterId"` // Key: 0
	GoalValue             int64 `json:"goalValue" msgpack:"goalValue"`                         // Key: 2
}

type TeamChallengeResult struct {
	ThisTimeGoal    *TeamChallengeGoalLotResult `json:"thisTimeGoal" msgpack:"thisTimeGoal"`       // Key: 0
	IsCleared       bool                        `json:"isCleared" msgpack:"isCleared"`             // Key: 1
	ActualGoalValue int64                       `json:"actualGoalValue" msgpack:"actualGoalValue"` // Key: 2
}

type StartTeamChallengeSessionPayload struct {
	Difficulty TeamChallengeDifficultyTypes `json:"difficulty" msgpack:"difficulty"` // Key: 0
	RoomId     string                       `json:"roomId" msgpack:"roomId"`         // Key: 1
}

type StartTeamChallengeSessionResult struct {
	MultiLiveId int64                       `json:"multiLiveId" msgpack:"multiLiveId"` // Key: 0
	LottedGoal  *TeamChallengeGoalLotResult `json:"lottedGoal" msgpack:"lottedGoal"`   // Key: 1
}

type FinishTeamChallengeSessionPayload struct {
	RoomId string `json:"roomId" msgpack:"roomId"` // Key: 0
}

type TotalPaymentResult struct {
	TotalPaymentThisMonth int `json:"totalPaymentThisMonth" msgpack:"totalPaymentThisMonth"` // Key: 0
}

type TotalPointEventInformationResult struct {
	CurrentTotalPoint int64 `json:"currentTotalPoint" msgpack:"currentTotalPoint"` // Key: 0
}

type TournamentQualifyingInformationResult struct {
	EntryCode string `json:"entryCode" msgpack:"entryCode"` // Key: 0
}

type TransitionTokenResult struct {
	Token string `json:"token" msgpack:"token"` // Key: 0
}

type TripleCastPartyAndRankingResult struct {
	Rank               *int                    `json:"rank" msgpack:"rank"`                             // Key: 0
	UserId             string                  `json:"userId" msgpack:"userId"`                         // Key: 1
	BestScore          int64                   `json:"bestScore" msgpack:"bestScore"`                   // Key: 2
	TripleCastMasterId int64                   `json:"tripleCastMasterId" msgpack:"tripleCastMasterId"` // Key: 3
	ClassType          LeagueClassTypes        `json:"classType" msgpack:"classType"`                   // Key: 4
	Difficulty         MusicDifficulties       `json:"difficulty" msgpack:"difficulty"`                 // Key: 5
	UserName           string                  `json:"userName" msgpack:"userName"`                     // Key: 6
	Parties            []TripleCastPartyResult `json:"parties" msgpack:"parties"`                       // Key: 7
}

type TripleCastPartyResult struct {
	Order int                     `json:"order" msgpack:"order"` // Key: 0
	Score int64                   `json:"score" msgpack:"score"` // Key: 1
	Slots []LeaguePartySlotResult `json:"slots" msgpack:"slots"` // Key: 2
}

type UpdateClearLampResult struct {
	ClearLamp ClearLamps `json:"clearLamp" msgpack:"clearLamp"` // Key: 0
}

type UrlResult struct {
	Url string `json:"url" msgpack:"url"` // Key: 0
}

type UserBlockIdsResult struct {
	BlockUserIds string `json:"blockUserIds" msgpack:"blockUserIds"` // Key: 0
}

type ViewShopResult struct {
	ConvertedThings []ConvertedThingResult `json:"convertedThings" msgpack:"convertedThings"` // Key: 0
}

type ConvertedThingResult struct {
	ExchangeShopMasterId int64          `json:"exchangeShopMasterId" msgpack:"exchangeShopMasterId"` // Key: 0
	OriginalQuantity     int            `json:"originalQuantity" msgpack:"originalQuantity"`         // Key: 1
	ReceivedThing        *ReceivedThing `json:"receivedThing" msgpack:"receivedThing"`               // Key: 2
}
