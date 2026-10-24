// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Events
package dto

import (
	"time"
)

type CharacterPointEvent struct {
	Id                          int64  `json:"id" msgpack:"id"`                                                   // Key: 0
	CharacterPointEventMasterId int64  `json:"characterPointEventMasterId" msgpack:"characterPointEventMasterId"` // Key: 1
	CharacterBaseMasterId       *int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`             // Key: 2
	TotalAcquiredPoint          int    `json:"totalAcquiredPoint" msgpack:"totalAcquiredPoint"`                   // Key: 3
	LastRank                    int    `json:"lastRank" msgpack:"lastRank"`                                       // Key: 4
	ReadTips                    bool   `json:"readTips" msgpack:"readTips"`                                       // Key: 5
}

type CircleEventMission struct {
	CircleEventMissionMasterId int64 `json:"circleEventMissionMasterId" msgpack:"circleEventMissionMasterId"` // Key: 0
	CurrentCount               int64 `json:"currentCount" msgpack:"currentCount"`                             // Key: 1
	IsActive                   bool  `json:"isActive" msgpack:"isActive"`                                     // Key: 2
}

type DugongRun struct {
	Id                     int64   `json:"id" msgpack:"id"`                                         // Key: 0
	ClearedCourseIds       []int64 `json:"clearedCourseIds" msgpack:"clearedCourseIds"`             // Key: 1
	NoMistakeCourseIds     []int64 `json:"noMistakeCourseIds" msgpack:"noMistakeCourseIds"`         // Key: 2
	DugongRunCourseGroupId int64   `json:"dugongRunCourseGroupId" msgpack:"dugongRunCourseGroupId"` // Key: 3
}

type Event struct {
	Id                       int64     `json:"id" msgpack:"id"`                                             // Key: 0
	EventMasterId            int64     `json:"eventMasterId" msgpack:"eventMasterId"`                       // Key: 1
	TotalAcquiredPoint       int       `json:"totalAcquiredPoint" msgpack:"totalAcquiredPoint"`             // Key: 2
	AcquiredPointUpdatedDate time.Time `json:"acquiredPointUpdatedDate" msgpack:"acquiredPointUpdatedDate"` // Key: 3
	LastRank                 *int      `json:"lastRank" msgpack:"lastRank"`                                 // Key: 4
	ReadTips                 bool      `json:"readTips" msgpack:"readTips"`                                 // Key: 5
	LoginDays                int       `json:"loginDays" msgpack:"loginDays"`                               // Key: 6
}

type EventBoxGacha struct {
	Id                    int64 `json:"id" msgpack:"id"`                                       // Key: 0
	EventBoxGachaMasterId int64 `json:"eventBoxGachaMasterId" msgpack:"eventBoxGachaMasterId"` // Key: 1
	CurrentBoxCount       int   `json:"currentBoxCount" msgpack:"currentBoxCount"`             // Key: 2
}

type EventBoxGachaBoxThing struct {
	Id                            int64 `json:"id" msgpack:"id"`                                                       // Key: 0
	EventBoxGachaBoxThingMasterId int64 `json:"eventBoxGachaBoxThingMasterId" msgpack:"eventBoxGachaBoxThingMasterId"` // Key: 1
	HitCount                      int   `json:"hitCount" msgpack:"hitCount"`                                           // Key: 2
}

type EventCamp struct {
	Id                int64     `json:"id" msgpack:"id"`                               // Key: 0
	EventMasterId     int64     `json:"eventMasterId" msgpack:"eventMasterId"`         // Key: 1
	CampType          CampTypes `json:"campType" msgpack:"campType"`                   // Key: 2
	TotalSupportPoint int       `json:"totalSupportPoint" msgpack:"totalSupportPoint"` // Key: 3
}

type RouletteEvent struct {
	Id                    int64 `json:"id" msgpack:"id"`                                       // Key: 0
	RouletteEventMasterId int64 `json:"rouletteEventMasterId" msgpack:"rouletteEventMasterId"` // Key: 1
	TotalAcquiredPoint    int   `json:"totalAcquiredPoint" msgpack:"totalAcquiredPoint"`       // Key: 2
}

type SpecialEvent struct {
	Id                   int64 `json:"id" msgpack:"id"`                                     // Key: 0
	SpecialEventMasterId int64 `json:"specialEventMasterId" msgpack:"specialEventMasterId"` // Key: 1
	ReadTips             bool  `json:"readTips" msgpack:"readTips"`                         // Key: 2
}

type StoryEvent struct {
	Id                       int64     `json:"id" msgpack:"id"`                                             // Key: 0
	StoryEventMasterId       int64     `json:"storyEventMasterId" msgpack:"storyEventMasterId"`             // Key: 1
	TotalAcquiredPoint       int       `json:"totalAcquiredPoint" msgpack:"totalAcquiredPoint"`             // Key: 2
	AcquiredPointUpdatedDate time.Time `json:"acquiredPointUpdatedDate" msgpack:"acquiredPointUpdatedDate"` // Key: 3
	LastRank                 *int      `json:"lastRank" msgpack:"lastRank"`                                 // Key: 4
	ReadTips                 bool      `json:"readTips" msgpack:"readTips"`                                 // Key: 5
	LoginDays                int       `json:"loginDays" msgpack:"loginDays"`                               // Key: 6
}

type StoryEventCircle struct {
	Id                 int64  `json:"id" msgpack:"id"`                                 // Key: 0
	StoryEventMasterId int64  `json:"storyEventMasterId" msgpack:"storyEventMasterId"` // Key: 1
	CurrentPoint       int    `json:"currentPoint" msgpack:"currentPoint"`             // Key: 2
	HighScore          int64  `json:"highScore" msgpack:"highScore"`                   // Key: 3
	CircleId           *int64 `json:"circleId" msgpack:"circleId"`                     // Key: 4
}

type StoryEventCircleMission struct {
	Id                              int64 `json:"id" msgpack:"id"`                                                           // Key: 0
	StoryEventCircleMissionMasterId int64 `json:"storyEventCircleMissionMasterId" msgpack:"storyEventCircleMissionMasterId"` // Key: 1
	CurrentCount                    int64 `json:"currentCount" msgpack:"currentCount"`                                       // Key: 2
}

type StoryEventCircleMissionReward struct {
	Id                                    int64 `json:"id" msgpack:"id"`                                                                       // Key: 0
	StoryEventMasterId                    int64 `json:"storyEventMasterId" msgpack:"storyEventMasterId"`                                       // Key: 1
	StoryEventCircleMissionRewardMasterId int64 `json:"storyEventCircleMissionRewardMasterId" msgpack:"storyEventCircleMissionRewardMasterId"` // Key: 2
}

type StoryEventHighScore struct {
	Id                            int64 `json:"id" msgpack:"id"`                                                       // Key: 0
	StoryEventMasterId            int64 `json:"storyEventMasterId" msgpack:"storyEventMasterId"`                       // Key: 1
	CurrentEnhancementPoint       int   `json:"currentEnhancementPoint" msgpack:"currentEnhancementPoint"`             // Key: 2
	TotalAcquiredEnhancementPoint int   `json:"totalAcquiredEnhancementPoint" msgpack:"totalAcquiredEnhancementPoint"` // Key: 3
}

type StoryEventHighScoreBuffSetting struct {
	StoryEventHighScoreBuffSettingMasterId int64 `json:"storyEventHighScoreBuffSettingMasterId" msgpack:"storyEventHighScoreBuffSettingMasterId"` // Key: 0
	CurrentLevel                           int   `json:"currentLevel" msgpack:"currentLevel"`                                                     // Key: 1
}

type StoryEventHighScoreParty struct {
	Id                  int64                          `json:"id" msgpack:"id"`                                   // Key: 0
	StoryEventMasterId  int64                          `json:"storyEventMasterId" msgpack:"storyEventMasterId"`   // Key: 1
	HighScore           int64                          `json:"highScore" msgpack:"highScore"`                     // Key: 3
	RateGrade           AchievementRateGrades          `json:"rateGrade" msgpack:"rateGrade"`                     // Key: 5
	Difficulty          MusicDifficulties              `json:"difficulty" msgpack:"difficulty"`                   // Key: 6
	HighScoreType       HighScoreTypes                 `json:"highScoreType" msgpack:"highScoreType"`             // Key: 8
	LiveSettingMasterId int64                          `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"` // Key: 9
	Slots               []StoryEventHighScorePartySlot `json:"slots" msgpack:"slots"`                             // Key: 10
	LeaderPosition      int                            `json:"leaderPosition" msgpack:"leaderPosition"`           // Key: 11
}

type StoryEventHighScorePartySlot struct {
	Id                              int64                     `json:"id" msgpack:"id"`                                                           // Key: 0
	StoryEventHighScoreClearPartyId int64                     `json:"storyEventHighScoreClearPartyId" msgpack:"storyEventHighScoreClearPartyId"` // Key: 1
	StoryEventHighScoreParty        *StoryEventHighScoreParty `json:"storyEventHighScoreParty" msgpack:"storyEventHighScoreParty"`
	Position                        int                       `json:"position" msgpack:"position"`                   // Key: 2
	CharacterMasterId               int64                     `json:"characterMasterId" msgpack:"characterMasterId"` // Key: 3
	CharacterMaster                 *CharacterMaster          `json:"characterMaster" msgpack:"characterMaster"`
	CharacterLevel                  int                       `json:"characterLevel" msgpack:"characterLevel"`                   // Key: 4
	CharacterTalentStage            int                       `json:"characterTalentStage" msgpack:"characterTalentStage"`       // Key: 5
	CharacterAwakeningPhase         int                       `json:"characterAwakeningPhase" msgpack:"characterAwakeningPhase"` // Key: 6
	PosterMasterId                  *int64                    `json:"posterMasterId" msgpack:"posterMasterId"`                   // Key: 7
	PosterMaster                    *PosterMaster             `json:"posterMaster" msgpack:"posterMaster"`
	PosterLevel                     *int                      `json:"posterLevel" msgpack:"posterLevel"`                         // Key: 8
	PosterBreakthroughPhase         *int                      `json:"posterBreakthroughPhase" msgpack:"posterBreakthroughPhase"` // Key: 9
	AccessoryMasterId               *int64                    `json:"accessoryMasterId" msgpack:"accessoryMasterId"`             // Key: 10
	AccessoryMaster                 *AccessoryMaster          `json:"accessoryMaster" msgpack:"accessoryMaster"`
	AccessoryLevel                  *int                      `json:"accessoryLevel" msgpack:"accessoryLevel"`                                   // Key: 11
	CurrentStatus                   *Status                   `json:"currentStatus" msgpack:"currentStatus"`                                     // Key: 12
	CharacterDisplayAwakeningStatus bool                      `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 13
}

type TotalPointEvent struct {
	Id                      int64 `json:"id" msgpack:"id"`                                           // Key: 0
	TotalPointEventMasterId int64 `json:"totalPointEventMasterId" msgpack:"totalPointEventMasterId"` // Key: 1
	TotalAcquiredPoint      int64 `json:"totalAcquiredPoint" msgpack:"totalAcquiredPoint"`           // Key: 2
	ReceivedRewardOrder     int   `json:"receivedRewardOrder" msgpack:"receivedRewardOrder"`         // Key: 3
}

type TournamentDetail struct {
	Id                       int64     `json:"id" msgpack:"id"`                                             // Key: 0
	TournamentDetailMasterId int64     `json:"tournamentDetailMasterId" msgpack:"tournamentDetailMasterId"` // Key: 1
	BestUniqueScore          int64     `json:"bestUniqueScore" msgpack:"bestUniqueScore"`                   // Key: 2
	PerfectStar              int       `json:"perfectStar" msgpack:"perfectStar"`                           // Key: 3
	Perfect                  int       `json:"perfect" msgpack:"perfect"`                                   // Key: 4
	Great                    int       `json:"great" msgpack:"great"`                                       // Key: 5
	Good                     int       `json:"good" msgpack:"good"`                                         // Key: 6
	Bad                      int       `json:"bad" msgpack:"bad"`                                           // Key: 7
	Miss                     int       `json:"miss" msgpack:"miss"`                                         // Key: 8
	RecordedAt               time.Time `json:"recordedAt" msgpack:"recordedAt"`                             // Key: 9
}

type TournamentQualifying struct {
	Id                                int64      `json:"id" msgpack:"id"`                                                               // Key: 0
	TournamentQualifyingMasterId      int64      `json:"tournamentQualifyingMasterId" msgpack:"tournamentQualifyingMasterId"`           // Key: 1
	MusicCourseMasterId               int64      `json:"musicCourseMasterId" msgpack:"musicCourseMasterId"`                             // Key: 2
	CurrentChallengeCount             int        `json:"currentChallengeCount" msgpack:"currentChallengeCount"`                         // Key: 3
	PerfectStar                       int        `json:"perfectStar" msgpack:"perfectStar"`                                             // Key: 4
	Perfect                           int        `json:"perfect" msgpack:"perfect"`                                                     // Key: 5
	Great                             int        `json:"great" msgpack:"great"`                                                         // Key: 6
	Good                              int        `json:"good" msgpack:"good"`                                                           // Key: 7
	Bad                               int        `json:"bad" msgpack:"bad"`                                                             // Key: 8
	Miss                              int        `json:"miss" msgpack:"miss"`                                                           // Key: 9
	TotalAchievementRatePercentRecord float64    `json:"totalAchievementRatePercentRecord" msgpack:"totalAchievementRatePercentRecord"` // Key: 10
	BestRecordChallengeCount          int        `json:"bestRecordChallengeCount" msgpack:"bestRecordChallengeCount"`                   // Key: 11
	BestRecordDate                    *time.Time `json:"bestRecordDate" msgpack:"bestRecordDate"`                                       // Key: 12
}

type TrialPartyEvent struct {
	Id                      int64 `json:"id" msgpack:"id"`                                           // Key: 0
	TrialPartyEventMasterId int64 `json:"trialPartyEventMasterId" msgpack:"trialPartyEventMasterId"` // Key: 1
	CurrentStageOrder       int   `json:"currentStageOrder" msgpack:"currentStageOrder"`             // Key: 2
	IsCompleted             bool  `json:"isCompleted" msgpack:"isCompleted"`                         // Key: 3
}

type TrialPartyEventStage struct {
	Id                           int64 `json:"id" msgpack:"id"`                                                     // Key: 0
	TrialPartyEventStageMasterId int64 `json:"trialPartyEventStageMasterId" msgpack:"trialPartyEventStageMasterId"` // Key: 1
	IsCleared                    bool  `json:"isCleared" msgpack:"isCleared"`                                       // Key: 2
}

type TrialPartyEventStageParty struct {
	Id                           int64 `json:"id" msgpack:"id"`                                                     // Key: 0
	TrialPartyEventStageMasterId int64 `json:"trialPartyEventStageMasterId" msgpack:"trialPartyEventStageMasterId"` // Key: 1
	LeaderPosition               int   `json:"leaderPosition" msgpack:"leaderPosition"`                             // Key: 2
}

type TrialPartyEventStagePartySlot struct {
	Id                          int64  `json:"id" msgpack:"id"`                                                   // Key: 0
	TrialPartyEventStagePartyId int64  `json:"trialPartyEventStagePartyId" msgpack:"trialPartyEventStagePartyId"` // Key: 1
	Position                    int    `json:"position" msgpack:"position"`                                       // Key: 2
	TrialPartyCharacterMasterId int64  `json:"trialPartyCharacterMasterId" msgpack:"trialPartyCharacterMasterId"` // Key: 3
	TrialPartyPosterMasterId    *int64 `json:"trialPartyPosterMasterId" msgpack:"trialPartyPosterMasterId"`       // Key: 4
	TrialPartyAccessoryMasterId *int64 `json:"trialPartyAccessoryMasterId" msgpack:"trialPartyAccessoryMasterId"` // Key: 5
}
