// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Lives
package dto

import (
	"time"
)

type AnotherNotation struct {
	Id                           int64                 `json:"id" msgpack:"id"`                                                     // Key: 0
	AnotherNotationMasterId      int64                 `json:"anotherNotationMasterId" msgpack:"anotherNotationMasterId"`           // Key: 1
	ClearLamp                    ClearLamps            `json:"clearLamp" msgpack:"clearLamp"`                                       // Key: 2
	RateGrade                    AchievementRateGrades `json:"rateGrade" msgpack:"rateGrade"`                                       // Key: 3
	AchievementRatePercentRecord float64               `json:"achievementRatePercentRecord" msgpack:"achievementRatePercentRecord"` // Key: 4
}

type AuditionClear struct {
	Id                   int64  `json:"id" msgpack:"id"`                                     // Key: 0
	AuditionMasterId     int64  `json:"auditionMasterId" msgpack:"auditionMasterId"`         // Key: 1
	ClearPhase           byte   `json:"clearPhase" msgpack:"clearPhase"`                     // Key: 2
	AuditionClearPartyId int64  `json:"auditionClearPartyId" msgpack:"auditionClearPartyId"` // Key: 3
	UserId               string `json:"userId" msgpack:"userId"`                             // Key: 4
	SkipClearPhase       byte   `json:"skipClearPhase" msgpack:"skipClearPhase"`             // Key: 5
}

type BonusLive struct {
	Id                int64 `json:"id" msgpack:"id"`                               // Key: 0
	BonusLiveMasterId int64 `json:"bonusLiveMasterId" msgpack:"bonusLiveMasterId"` // Key: 1
	ClearedStageOrder int   `json:"clearedStageOrder" msgpack:"clearedStageOrder"` // Key: 2
	ReadTips          bool  `json:"readTips" msgpack:"readTips"`                   // Key: 3
	DailyClearTimes   int   `json:"dailyClearTimes" msgpack:"dailyClearTimes"`     // Key: 4
}

type BonusLiveStage struct {
	Id                     int64 `json:"id" msgpack:"id"`                                         // Key: 0
	BonusLiveMasterStageId int64 `json:"bonusLiveMasterStageId" msgpack:"bonusLiveMasterStageId"` // Key: 1
	ClearTimes             int   `json:"clearTimes" msgpack:"clearTimes"`                         // Key: 2
}

type ConcertStage struct {
	ConcertStageMasterId int64 `json:"concertStageMasterId" msgpack:"concertStageMasterId"` // Key: 0
}

type DailyLimit struct {
	Id                            int64      `json:"id" msgpack:"id"`                                                       // Key: 0
	AutoPlayTimes                 int        `json:"autoPlayTimes" msgpack:"autoPlayTimes"`                                 // Key: 1
	DailyLessonTimes              int        `json:"dailyLessonTimes" msgpack:"dailyLessonTimes"`                           // Key: 2
	LastRefreshedAt               *time.Time `json:"lastRefreshedAt" msgpack:"lastRefreshedAt"`                             // Key: 3
	MusicCourseFreeChallengeTimes int        `json:"musicCourseFreeChallengeTimes" msgpack:"musicCourseFreeChallengeTimes"` // Key: 4
}

type Live struct {
	Id              int64                 `json:"id" msgpack:"id"`                           // Key: 0
	LiveMasterId    int64                 `json:"liveMasterId" msgpack:"liveMasterId"`       // Key: 1
	TimesCompleted  int                   `json:"timesCompleted" msgpack:"timesCompleted"`   // Key: 2
	AchievementRate float32               `json:"achievementRate" msgpack:"achievementRate"` // Key: 4
	NotationRate    float32               `json:"notationRate" msgpack:"notationRate"`       // Key: 5
	ClearLamp       ClearLamps            `json:"clearLamp" msgpack:"clearLamp"`             // Key: 6
	Status          LiveReleaseStatus     `json:"status" msgpack:"status"`                   // Key: 7
	RateGrade       AchievementRateGrades `json:"rateGrade" msgpack:"rateGrade"`             // Key: 8
}

type LiveAchievement struct {
	Id                   int64 `json:"id" msgpack:"id"`                                     // Key: 0
	OlivierReleasedCount int   `json:"olivierReleasedCount" msgpack:"olivierReleasedCount"` // Key: 1
	OlivierClearedLevel  int   `json:"olivierClearedLevel" msgpack:"olivierClearedLevel"`   // Key: 2
}

type LiveDropCelling struct {
	Id                        int64 `json:"id" msgpack:"id"`                                               // Key: 0
	MultiLiveScheduleMasterId int64 `json:"multiLiveScheduleMasterId" msgpack:"multiLiveScheduleMasterId"` // Key: 1
	Count                     int   `json:"count" msgpack:"count"`                                         // Key: 2
	TotalCellingCount         int   `json:"totalCellingCount" msgpack:"totalCellingCount"`                 // Key: 3
}

type LiveDropLimit struct {
	MultiLiveScheduleMasterId int64 `json:"multiLiveScheduleMasterId" msgpack:"multiLiveScheduleMasterId"` // Key: 0
	CurrentCount              int   `json:"currentCount" msgpack:"currentCount"`                           // Key: 1
	CountLimit                int   `json:"countLimit" msgpack:"countLimit"`                               // Key: 2
}

type Music struct {
	Id                   int64                  `json:"id" msgpack:"id"`                                     // Key: 0
	MusicMasterId        int64                  `json:"musicMasterId" msgpack:"musicMasterId"`               // Key: 2
	StellaReleased       bool                   `json:"stellaReleased" msgpack:"stellaReleased"`             // Key: 5
	VocalVersion         int                    `json:"vocalVersion" msgpack:"vocalVersion"`                 // Key: 7
	OlivierReleaseStatus OlivierReleaseStatuses `json:"olivierReleaseStatus" msgpack:"olivierReleaseStatus"` // Key: 8
	IsPossession         bool                   `json:"isPossession" msgpack:"isPossession"`                 // Key: 9
}

type MusicBookmark struct {
	MusicMasterId     int64              `json:"musicMasterId" msgpack:"musicMasterId"`         // Key: 0
	MusicBookmarkFlag MusicBookmarkFlags `json:"musicBookmarkFlag" msgpack:"musicBookmarkFlag"` // Key: 1
}

type MusicCourse struct {
	Id                                int64                         `json:"id" msgpack:"id"`                                                               // Key: 0
	MusicCourseMasterId               int64                         `json:"musicCourseMasterId" msgpack:"musicCourseMasterId"`                             // Key: 1
	ClearLamp                         ClearLamps                    `json:"clearLamp" msgpack:"clearLamp"`                                                 // Key: 2
	CertificationGrade                MusicCourseCertificationGrade `json:"certificationGrade" msgpack:"certificationGrade"`                               // Key: 3
	TotalAchievementRatePercentRecord float64                       `json:"totalAchievementRatePercentRecord" msgpack:"totalAchievementRatePercentRecord"` // Key: 4
}

type MusicCourseRanking struct {
	Id                                int64      `json:"id" msgpack:"id"`                                                               // Key: 0
	MusicCourseMasterId               int64      `json:"musicCourseMasterId" msgpack:"musicCourseMasterId"`                             // Key: 1
	CurrentChallengeCount             int        `json:"currentChallengeCount" msgpack:"currentChallengeCount"`                         // Key: 2
	PerfectStar                       int        `json:"perfectStar" msgpack:"perfectStar"`                                             // Key: 3
	Perfect                           int        `json:"perfect" msgpack:"perfect"`                                                     // Key: 4
	Great                             int        `json:"great" msgpack:"great"`                                                         // Key: 5
	Good                              int        `json:"good" msgpack:"good"`                                                           // Key: 6
	Bad                               int        `json:"bad" msgpack:"bad"`                                                             // Key: 7
	Miss                              int        `json:"miss" msgpack:"miss"`                                                           // Key: 8
	TotalAchievementRatePercentRecord float64    `json:"totalAchievementRatePercentRecord" msgpack:"totalAchievementRatePercentRecord"` // Key: 9
	BestRecordChallengeCount          int        `json:"bestRecordChallengeCount" msgpack:"bestRecordChallengeCount"`                   // Key: 10
	BestRecordDate                    *time.Time `json:"bestRecordDate" msgpack:"bestRecordDate"`                                       // Key: 11
	HasReceivedReward                 bool       `json:"hasReceivedReward" msgpack:"hasReceivedReward"`                                 // Key: 12
}

type SpRate struct {
	Id           int64 `json:"id" msgpack:"id"`                     // Key: 0
	LiveMasterId int64 `json:"liveMasterId" msgpack:"liveMasterId"` // Key: 1
	Point        int   `json:"point" msgpack:"point"`               // Key: 2
}

type Concert struct {
	ConcertMasterId   int64 `json:"concertMasterId" msgpack:"concertMasterId"`     // Key: 0
	CurrentStageOrder int   `json:"currentStageOrder" msgpack:"currentStageOrder"` // Key: 1
	ReadTips          bool  `json:"readTips" msgpack:"readTips"`                   // Key: 2
}

type ApiLivesHighScoreParty struct {
	Id                  int64                        `json:"id" msgpack:"id"`                                   // Key: 0
	HighScore           int64                        `json:"highScore" msgpack:"highScore"`                     // Key: 1
	RateGrade           AchievementRateGrades        `json:"rateGrade" msgpack:"rateGrade"`                     // Key: 3
	Difficulty          MusicDifficulties            `json:"difficulty" msgpack:"difficulty"`                   // Key: 4
	MusicMasterId       int64                        `json:"musicMasterId" msgpack:"musicMasterId"`             // Key: 5
	HighScoreType       HighScoreTypes               `json:"highScoreType" msgpack:"highScoreType"`             // Key: 6
	LiveSettingMasterId int64                        `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"` // Key: 7
	Slots               []ApiLivesHighScorePartySlot `json:"slots" msgpack:"slots"`                             // Key: 8
	UserName            string                       `json:"userName" msgpack:"userName"`                       // Key: 9
	ActingAbility       int                          `json:"actingAbility" msgpack:"actingAbility"`             // Key: 10
}

type ApiLivesHighScorePartySlot struct {
	Id                              int64                   `json:"id" msgpack:"id"`                             // Key: 0
	HighScorePartyId                int64                   `json:"highScorePartyId" msgpack:"highScorePartyId"` // Key: 1
	HighScoreParty                  *ApiLivesHighScoreParty `json:"highScoreParty" msgpack:"highScoreParty"`
	Position                        int                     `json:"position" msgpack:"position"`                   // Key: 2
	CharacterMasterId               int64                   `json:"characterMasterId" msgpack:"characterMasterId"` // Key: 3
	CharacterMaster                 *CharacterMaster        `json:"characterMaster" msgpack:"characterMaster"`
	CharacterLevel                  int                     `json:"characterLevel" msgpack:"characterLevel"`                   // Key: 4
	CharacterTalentStage            int                     `json:"characterTalentStage" msgpack:"characterTalentStage"`       // Key: 5
	CharacterAwakeningPhase         int                     `json:"characterAwakeningPhase" msgpack:"characterAwakeningPhase"` // Key: 6
	PosterMasterId                  *int64                  `json:"posterMasterId" msgpack:"posterMasterId"`                   // Key: 7
	PosterMaster                    *PosterMaster           `json:"posterMaster" msgpack:"posterMaster"`
	PosterLevel                     *int                    `json:"posterLevel" msgpack:"posterLevel"`                         // Key: 8
	PosterBreakthroughPhase         *int                    `json:"posterBreakthroughPhase" msgpack:"posterBreakthroughPhase"` // Key: 9
	AccessoryMasterId               *int64                  `json:"accessoryMasterId" msgpack:"accessoryMasterId"`             // Key: 10
	AccessoryMaster                 *AccessoryMaster        `json:"accessoryMaster" msgpack:"accessoryMaster"`
	AccessoryLevel                  *int                    `json:"accessoryLevel" msgpack:"accessoryLevel"`                                   // Key: 11
	CurrentStatus                   *Status                 `json:"currentStatus" msgpack:"currentStatus"`                                     // Key: 12
	CharacterDisplayAwakeningStatus bool                    `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 13
}

type BaseScoreBlock struct {
	Hash       int64       `json:"hash" msgpack:"hash"`             // Key: 0
	Score      int64       `json:"score" msgpack:"score"`           // Key: 1
	Life       int         `json:"life" msgpack:"life"`             // Key: 2
	NoteId     int64       `json:"noteId" msgpack:"noteId"`         // Key: 3
	TimingType TimingTypes `json:"timingType" msgpack:"timingType"` // Key: 4
	Combo      int         `json:"combo" msgpack:"combo"`           // Key: 5
}

type SenseScoreBlock struct {
	Hash            int64 `json:"hash" msgpack:"hash"`                       // Key: 0
	Score           int64 `json:"score" msgpack:"score"`                     // Key: 1
	Life            int   `json:"life" msgpack:"life"`                       // Key: 2
	TimeEventSecond int   `json:"timeEventSecond" msgpack:"timeEventSecond"` // Key: 3
	SenseId         int64 `json:"senseId" msgpack:"senseId"`                 // Key: 4
	Combo           int   `json:"combo" msgpack:"combo"`                     // Key: 5
}

type StarActScoreBlock struct {
	Hash            int64 `json:"hash" msgpack:"hash"`                       // Key: 0
	Score           int64 `json:"score" msgpack:"score"`                     // Key: 1
	Life            int   `json:"life" msgpack:"life"`                       // Key: 2
	TimeEventSecond int   `json:"timeEventSecond" msgpack:"timeEventSecond"` // Key: 3
	Combo           int   `json:"combo" msgpack:"combo"`                     // Key: 4
}

type MultiLiveAdditionalScoreBlock struct {
	Hash            int64 `json:"hash" msgpack:"hash"`                       // Key: 0
	Score           int64 `json:"score" msgpack:"score"`                     // Key: 1
	Life            int   `json:"life" msgpack:"life"`                       // Key: 2
	TimeEventSecond int   `json:"timeEventSecond" msgpack:"timeEventSecond"` // Key: 3
	Combo           int   `json:"combo" msgpack:"combo"`                     // Key: 4
}
