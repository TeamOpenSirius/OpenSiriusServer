// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Lessons
package dto

type CharacterLesson struct {
	CharacterBaseMasterId   int64                 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`     // Key: 0
	SetCharacters           []CharacterLessonSlot `json:"setCharacters" msgpack:"setCharacters"`                     // Key: 1
	BestScore               int64                 `json:"bestScore" msgpack:"bestScore"`                             // Key: 2
	LeaderPosition          int                   `json:"leaderPosition" msgpack:"leaderPosition"`                   // Key: 3
	RewardReceivedHighScore int64                 `json:"rewardReceivedHighScore" msgpack:"rewardReceivedHighScore"` // Key: 4
}

type CharacterLessonSlot struct {
	Position       int    `json:"position" msgpack:"position"`             // Key: 0
	SetCharacterId *int64 `json:"setCharacterId" msgpack:"setCharacterId"` // Key: 1
}

type DailyLesson struct {
	TimesLeft int `json:"timesLeft" msgpack:"timesLeft"` // Key: 0
}
