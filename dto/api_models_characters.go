// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Characters
package dto

type Character struct {
	Id                       int64                   `json:"id" msgpack:"id"`                               // Key: 0
	CharacterMasterId        int64                   `json:"characterMasterId" msgpack:"characterMasterId"` // Key: 1
	CharacterMaster          *CharacterMaster        `json:"characterMaster" msgpack:"characterMaster"`
	Level                    int                     `json:"level" msgpack:"level"`                         // Key: 2
	CurrentExperience        int                     `json:"currentExperience" msgpack:"currentExperience"` // Key: 3
	TalentStage              int                     `json:"talentStage" msgpack:"talentStage"`             // Key: 4
	AwakeningPhase           int                     `json:"awakeningPhase" msgpack:"awakeningPhase"`       // Key: 5
	CharacterBaseId          int64                   `json:"characterBaseId" msgpack:"characterBaseId"`     // Key: 6
	CharacterBase            *CharacterBase          `json:"characterBase" msgpack:"characterBase"`
	SenseLevel               int                     `json:"senseLevel" msgpack:"senseLevel"`                             // Key: 7
	ReadEpisodeOrder         CharacterEpisodeOrder   `json:"readEpisodeOrder" msgpack:"readEpisodeOrder"`                 // Key: 8
	ReleasedEpisodeOrder     CharacterEpisodeOrder   `json:"releasedEpisodeOrder" msgpack:"releasedEpisodeOrder"`         // Key: 9
	DisplayAwakeningStatus   bool                    `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"`     // Key: 10
	SecondaryCharacterBaseId *int64                  `json:"secondaryCharacterBaseId" msgpack:"secondaryCharacterBaseId"` // Key: 11
	SecondaryCharacterBase   *CharacterBase          `json:"secondaryCharacterBase" msgpack:"secondaryCharacterBase"`
	SecondarySenseLevel      int                     `json:"secondarySenseLevel" msgpack:"secondarySenseLevel"` // Key: 12
	SelectionType            CharacterSelectionTypes `json:"selectionType" msgpack:"selectionType"`             // Key: 13
	IsFavorite               bool                    `json:"isFavorite" msgpack:"isFavorite"`                   // Key: 14
}

type CharacterBase struct {
	Id                           int64                `json:"id" msgpack:"id"`                                       // Key: 0
	CharacterBaseMasterId        int64                `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"` // Key: 1
	CharacterBaseMaster          *CharacterBaseMaster `json:"characterBaseMaster" msgpack:"characterBaseMaster"`
	StarRank                     int                  `json:"starRank" msgpack:"starRank"`                                         // Key: 2
	TotalStarPoint               int                  `json:"totalStarPoint" msgpack:"totalStarPoint"`                             // Key: 3
	CostumeMasterId              *int64               `json:"costumeMasterId" msgpack:"costumeMasterId"`                           // Key: 4
	KeyMissionLevel              int                  `json:"keyMissionLevel" msgpack:"keyMissionLevel"`                           // Key: 5
	PortalCharacterId            int64                `json:"portalCharacterId" msgpack:"portalCharacterId"`                       // Key: 6
	PortalDisplayAwakeningStatus bool                 `json:"portalDisplayAwakeningStatus" msgpack:"portalDisplayAwakeningStatus"` // Key: 7
}

type CharacterMission struct {
	Id                       int64 `json:"id" msgpack:"id"`                                             // Key: 0
	CharacterBaseMasterId    int64 `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`       // Key: 1
	CharacterMissionMasterId int64 `json:"characterMissionMasterId" msgpack:"characterMissionMasterId"` // Key: 2
	CurrentStageMasterId     int64 `json:"currentStageMasterId" msgpack:"currentStageMasterId"`         // Key: 3
	CurrentCount             int64 `json:"currentCount" msgpack:"currentCount"`                         // Key: 4
	ClearedStageOrder        int   `json:"clearedStageOrder" msgpack:"clearedStageOrder"`               // Key: 5
	RewardReceivedStageOrder int   `json:"rewardReceivedStageOrder" msgpack:"rewardReceivedStageOrder"` // Key: 6
	CompletedLevel           int   `json:"completedLevel" msgpack:"completedLevel"`                     // Key: 7
}

type LinkCharacter struct {
	Id                          int64  `json:"id" msgpack:"id"`                                                   // Key: 0
	CharacterBaseMasterId       int64  `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`             // Key: 1
	CompanyMasterId             int64  `json:"companyMasterId" msgpack:"companyMasterId"`                         // Key: 2
	LinkedCharacterBaseMasterId *int64 `json:"linkedCharacterBaseMasterId" msgpack:"linkedCharacterBaseMasterId"` // Key: 3
	RewardReceivedMaxRank       int    `json:"rewardReceivedMaxRank" msgpack:"rewardReceivedMaxRank"`             // Key: 4
}
