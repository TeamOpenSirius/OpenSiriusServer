// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Leagues
package dto

type LeagueBasic struct {
	MyProperty                     int64            `json:"myProperty" msgpack:"myProperty"`                                         // Key: 0
	BestClassType                  LeagueClassTypes `json:"bestClassType" msgpack:"bestClassType"`                                   // Key: 4
	StarEnrollCount                int              `json:"starEnrollCount" msgpack:"starEnrollCount"`                               // Key: 1
	DaiStarEnrollCount             int              `json:"daiStarEnrollCount" msgpack:"daiStarEnrollCount"`                         // Key: 2
	CurrentClassType               LeagueClassTypes `json:"currentClassType" msgpack:"currentClassType"`                             // Key: 3
	LastJoinedLeagueSeasonMasterId *int64           `json:"lastJoinedLeagueSeasonMasterId" msgpack:"lastJoinedLeagueSeasonMasterId"` // Key: 5
}

type LeagueGroup struct {
	LeagueMasterId int64            `json:"leagueMasterId" msgpack:"leagueMasterId"` // Key: 0
	ClassType      LeagueClassTypes `json:"classType" msgpack:"classType"`           // Key: 1
	ClassOrder     int              `json:"classOrder" msgpack:"classOrder"`         // Key: 2
}

type LeagueGroupMember struct {
	LeagueGroupId  int64  `json:"leagueGroupId" msgpack:"leagueGroupId"`   // Key: 0
	LeagueMasterId int64  `json:"leagueMasterId" msgpack:"leagueMasterId"` // Key: 1
	BestScore      *int64 `json:"bestScore" msgpack:"bestScore"`           // Key: 2
}

type LeagueHighScoreParty struct {
	Id             int64                      `json:"id" msgpack:"id"`                         // Key: 0
	LeagueMasterId int64                      `json:"leagueMasterId" msgpack:"leagueMasterId"` // Key: 1
	HighScore      int64                      `json:"highScore" msgpack:"highScore"`           // Key: 2
	ClassType      LeagueClassTypes           `json:"classType" msgpack:"classType"`           // Key: 3
	Difficulty     MusicDifficulties          `json:"difficulty" msgpack:"difficulty"`         // Key: 4
	MusicMasterId  int64                      `json:"musicMasterId" msgpack:"musicMasterId"`   // Key: 5
	LeagueGroupId  int64                      `json:"leagueGroupId" msgpack:"leagueGroupId"`   // Key: 6
	Slots          []LeagueHighScorePartySlot `json:"slots" msgpack:"slots"`                   // Key: 7
	UserName       string                     `json:"userName" msgpack:"userName"`             // Key: 8
	ActingAbility  int                        `json:"actingAbility" msgpack:"actingAbility"`   // Key: 9
	LeaderPosition int                        `json:"leaderPosition" msgpack:"leaderPosition"` // Key: 10
}

type LeagueHighScorePartySlot struct {
	Id                              int64                 `json:"id" msgpack:"id"`                                         // Key: 0
	LeagueHighScorePartyId          int64                 `json:"leagueHighScorePartyId" msgpack:"leagueHighScorePartyId"` // Key: 1
	LeagueHighScoreParty            *LeagueHighScoreParty `json:"leagueHighScoreParty" msgpack:"leagueHighScoreParty"`
	Position                        int                   `json:"position" msgpack:"position"`                                               // Key: 2
	CharacterMasterId               int64                 `json:"characterMasterId" msgpack:"characterMasterId"`                             // Key: 3
	CharacterLevel                  int                   `json:"characterLevel" msgpack:"characterLevel"`                                   // Key: 4
	PosterMasterId                  *int64                `json:"posterMasterId" msgpack:"posterMasterId"`                                   // Key: 7
	PosterLevel                     *int                  `json:"posterLevel" msgpack:"posterLevel"`                                         // Key: 8
	PosterBreakthroughPhase         *int                  `json:"posterBreakthroughPhase" msgpack:"posterBreakthroughPhase"`                 // Key: 9
	AccessoryMasterId               *int64                `json:"accessoryMasterId" msgpack:"accessoryMasterId"`                             // Key: 10
	AccessoryLevel                  *int                  `json:"accessoryLevel" msgpack:"accessoryLevel"`                                   // Key: 11
	CurrentStatus                   *Status               `json:"currentStatus" msgpack:"currentStatus"`                                     // Key: 12
	CharacterTalentStage            int                   `json:"characterTalentStage" msgpack:"characterTalentStage"`                       // Key: 13
	CharacterAwakeningPhase         int                   `json:"characterAwakeningPhase" msgpack:"characterAwakeningPhase"`                 // Key: 14
	CharacterDisplayAwakeningStatus bool                  `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 15
}

type LeagueHistory struct {
	LeagueMasterId     int64                  `json:"leagueMasterId" msgpack:"leagueMasterId"`         // Key: 0
	ClassType          LeagueClassTypes       `json:"classType" msgpack:"classType"`                   // Key: 1
	HistoryCount       int                    `json:"historyCount" msgpack:"historyCount"`             // Key: 2
	IsSendedReward     bool                   `json:"isSendedReward" msgpack:"isSendedReward"`         // Key: 3
	IsPlayed           bool                   `json:"isPlayed" msgpack:"isPlayed"`                     // Key: 5
	ClassChangeType    LeagueClassChangeTypes `json:"classChangeType" msgpack:"classChangeType"`       // Key: 6
	GroupRank          int                    `json:"groupRank" msgpack:"groupRank"`                   // Key: 8
	GlobalRank         int                    `json:"globalRank" msgpack:"globalRank"`                 // Key: 9
	AllClassGlobalRank int                    `json:"allClassGlobalRank" msgpack:"allClassGlobalRank"` // Key: 10
}

type LeagueSeasonResult struct {
	LeagueSeasonMasterId  int64 `json:"leagueSeasonMasterId" msgpack:"leagueSeasonMasterId"`   // Key: 0
	DaiStarMaxEnrollCount int   `json:"daiStarMaxEnrollCount" msgpack:"daiStarMaxEnrollCount"` // Key: 1
}

type TripleCastBasic struct {
	Id                                 int64            `json:"id" msgpack:"id"`                                                                 // Key: 0
	BestClassType                      LeagueClassTypes `json:"bestClassType" msgpack:"bestClassType"`                                           // Key: 4
	StarEnrollCount                    int              `json:"starEnrollCount" msgpack:"starEnrollCount"`                                       // Key: 1
	DaiStarEnrollCount                 int              `json:"daiStarEnrollCount" msgpack:"daiStarEnrollCount"`                                 // Key: 2
	CurrentClassType                   LeagueClassTypes `json:"currentClassType" msgpack:"currentClassType"`                                     // Key: 3
	LastJoinedTripleCastSeasonMasterId *int64           `json:"lastJoinedTripleCastSeasonMasterId" msgpack:"lastJoinedTripleCastSeasonMasterId"` // Key: 5
	PartyOrder1                        *int             `json:"partyOrder1" msgpack:"partyOrder1"`                                               // Key: 6
	PartyOrder2                        *int             `json:"partyOrder2" msgpack:"partyOrder2"`                                               // Key: 7
	PartyOrder3                        *int             `json:"partyOrder3" msgpack:"partyOrder3"`                                               // Key: 8
}

type TripleCastGroup struct {
	TripleCastMasterId int64            `json:"tripleCastMasterId" msgpack:"tripleCastMasterId"` // Key: 0
	ClassType          LeagueClassTypes `json:"classType" msgpack:"classType"`                   // Key: 1
	ClassOrder         int              `json:"classOrder" msgpack:"classOrder"`                 // Key: 2
}

type TripleCastGroupMember struct {
	TripleCastGroupId  int64  `json:"tripleCastGroupId" msgpack:"tripleCastGroupId"`   // Key: 0
	TripleCastMasterId int64  `json:"tripleCastMasterId" msgpack:"tripleCastMasterId"` // Key: 1
	BestScore          *int64 `json:"bestScore" msgpack:"bestScore"`                   // Key: 2
}

type TripleCastHighScoreParty struct {
	Id                 int64                          `json:"id" msgpack:"id"`                                 // Key: 0
	TripleCastMasterId int64                          `json:"tripleCastMasterId" msgpack:"tripleCastMasterId"` // Key: 1
	Order              int                            `json:"order" msgpack:"order"`                           // Key: 2
	HighScore          int64                          `json:"highScore" msgpack:"highScore"`                   // Key: 3
	Slots              []TripleCastHighScorePartySlot `json:"slots" msgpack:"slots"`                           // Key: 4
	ActingAbility      int                            `json:"actingAbility" msgpack:"actingAbility"`           // Key: 5
	LeaderPosition     int                            `json:"leaderPosition" msgpack:"leaderPosition"`         // Key: 6
	Difficulty         MusicDifficulties              `json:"difficulty" msgpack:"difficulty"`                 // Key: 7
}

type TripleCastHighScorePartySlot struct {
	Id                              int64                     `json:"id" msgpack:"id"`                                                 // Key: 0
	TripleCastHighScorePartyId      int64                     `json:"tripleCastHighScorePartyId" msgpack:"tripleCastHighScorePartyId"` // Key: 1
	TripleCastHighScoreParty        *TripleCastHighScoreParty `json:"tripleCastHighScoreParty" msgpack:"tripleCastHighScoreParty"`
	Position                        int                       `json:"position" msgpack:"position"`                                               // Key: 2
	CharacterMasterId               int64                     `json:"characterMasterId" msgpack:"characterMasterId"`                             // Key: 3
	CharacterLevel                  int                       `json:"characterLevel" msgpack:"characterLevel"`                                   // Key: 4
	PosterMasterId                  *int64                    `json:"posterMasterId" msgpack:"posterMasterId"`                                   // Key: 7
	PosterLevel                     *int                      `json:"posterLevel" msgpack:"posterLevel"`                                         // Key: 8
	PosterBreakthroughPhase         *int                      `json:"posterBreakthroughPhase" msgpack:"posterBreakthroughPhase"`                 // Key: 9
	AccessoryMasterId               *int64                    `json:"accessoryMasterId" msgpack:"accessoryMasterId"`                             // Key: 10
	AccessoryLevel                  *int                      `json:"accessoryLevel" msgpack:"accessoryLevel"`                                   // Key: 11
	CurrentStatus                   *Status                   `json:"currentStatus" msgpack:"currentStatus"`                                     // Key: 12
	CharacterTalentStage            int                       `json:"characterTalentStage" msgpack:"characterTalentStage"`                       // Key: 13
	CharacterAwakeningPhase         int                       `json:"characterAwakeningPhase" msgpack:"characterAwakeningPhase"`                 // Key: 14
	CharacterDisplayAwakeningStatus bool                      `json:"characterDisplayAwakeningStatus" msgpack:"characterDisplayAwakeningStatus"` // Key: 15
}

type TripleCastHistory struct {
	TripleCastMasterId int64                  `json:"tripleCastMasterId" msgpack:"tripleCastMasterId"` // Key: 0
	ClassType          LeagueClassTypes       `json:"classType" msgpack:"classType"`                   // Key: 1
	HistoryCount       int                    `json:"historyCount" msgpack:"historyCount"`             // Key: 2
	IsSendedReward     bool                   `json:"isSendedReward" msgpack:"isSendedReward"`         // Key: 3
	IsPlayed           bool                   `json:"isPlayed" msgpack:"isPlayed"`                     // Key: 5
	ClassChangeType    LeagueClassChangeTypes `json:"classChangeType" msgpack:"classChangeType"`       // Key: 6
	GroupRank          int                    `json:"groupRank" msgpack:"groupRank"`                   // Key: 8
	GlobalRank         int                    `json:"globalRank" msgpack:"globalRank"`                 // Key: 9
	AllClassGlobalRank int                    `json:"allClassGlobalRank" msgpack:"allClassGlobalRank"` // Key: 10
}

type TripleCastSeasonResult struct {
	TripleCastSeasonMasterId int64 `json:"tripleCastSeasonMasterId" msgpack:"tripleCastSeasonMasterId"` // Key: 0
	DaiStarMaxEnrollCount    int   `json:"daiStarMaxEnrollCount" msgpack:"daiStarMaxEnrollCount"`       // Key: 1
}
