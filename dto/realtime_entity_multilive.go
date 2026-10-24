// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusRealtime.Shared.Entity.MultiLive
package dto

import (
	"time"
)

type MultiLiveCharacter struct {
	MCharacterId           int64  `json:"mCharacterId" msgpack:"mCharacterId"`                     // Key: 0
	TalentStage            int    `json:"talentStage" msgpack:"talentStage"`                       // Key: 1
	MNameColorId           *int64 `json:"mNameColorId" msgpack:"mNameColorId"`                     // Key: 2
	MNamePlateId           *int64 `json:"mNamePlateId" msgpack:"mNamePlateId"`                     // Key: 3
	MTrophyId1             *int64 `json:"mTrophyId1" msgpack:"mTrophyId1"`                         // Key: 4
	MTrophyId2             *int64 `json:"mTrophyId2" msgpack:"mTrophyId2"`                         // Key: 5
	MTrophyId3             *int64 `json:"mTrophyId3" msgpack:"mTrophyId3"`                         // Key: 6
	DisplayAwakeningStatus bool   `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 7
	MNameplateDetailId     *int64 `json:"mNameplateDetailId" msgpack:"mNameplateDetailId"`         // Key: 8
	TotalStatus            int    `json:"totalStatus" msgpack:"totalStatus"`                       // Key: 9
	MNameBaseColorid       int64  `json:"mNameBaseColorid" msgpack:"mNameBaseColorid"`             // Key: 10
}

type MultiLiveHallMember struct {
	UserId    string `json:"userId" msgpack:"userId"`       // Key: 0
	Timestamp int64  `json:"timestamp" msgpack:"timestamp"` // Key: 1
}

type MultiLiveMatchingAssignment struct {
	LiveSettingMasterId int64  `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"` // Key: 0
	GroupName           string `json:"groupName" msgpack:"groupName"`                     // Key: 1
	UserId              string `json:"userId" msgpack:"userId"`                           // Key: 2
}

type MultiLiveMatchMakingUserTicket struct {
	ServerId            string `json:"serverId" msgpack:"serverId"`                       // Key: 0
	UserId              string `json:"userId" msgpack:"userId"`                           // Key: 1
	LiveSettingMasterId int64  `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"` // Key: 2
}

type MultiLiveRecruitingRoomTicket struct {
	LiveSettingMasterId int64     `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"` // Key: 0
	GroupName           string    `json:"groupName" msgpack:"groupName"`                     // Key: 1
	UpdatedAt           time.Time `json:"updatedAt" msgpack:"updatedAt"`                     // Key: 2
}

type MultiLiveUser struct {
	MemberId             int                 `json:"memberId" msgpack:"memberId"`                         // Key: 0
	UserName             string              `json:"userName" msgpack:"userName"`                         // Key: 1
	IsRandomMusic        bool                `json:"isRandomMusic" msgpack:"isRandomMusic"`               // Key: 2
	SelectedMusicId      int64               `json:"selectedMusicId" msgpack:"selectedMusicId"`           // Key: 3
	Status               MultiLiveUserStatus `json:"status" msgpack:"status"`                             // Key: 4
	LeaderCharacter      *MultiLiveCharacter `json:"leaderCharacter" msgpack:"leaderCharacter"`           // Key: 5
	UserNamePlateColorId int64               `json:"userNamePlateColorId" msgpack:"userNamePlateColorId"` // Key: 6
	Difficulty           MusicDifficulties   `json:"difficulty" msgpack:"difficulty"`                     // Key: 7
	ClearLamp            ClearLamps          `json:"clearLamp" msgpack:"clearLamp"`                       // Key: 8
	Score                int64               `json:"score" msgpack:"score"`                               // Key: 9
	HashUserId           string              `json:"hashUserId" msgpack:"hashUserId"`                     // Key: 10
	IsExitGame           bool                `json:"isExitGame" msgpack:"isExitGame"`                     // Key: 11
	IsSelectedMusic      bool                `json:"isSelectedMusic" msgpack:"isSelectedMusic"`
	IsAFK                bool                `json:"isAFK" msgpack:"isAFK"`                             // Key: 12
	IsReadyDecideMember  bool                `json:"isReadyDecideMember" msgpack:"isReadyDecideMember"` // Key: 13
}

type TeamChallengeSelectDifficulyAndParty struct {
	MultiLiveCharacter *MultiLiveCharacter `json:"multiLiveCharacter" msgpack:"multiLiveCharacter"` // Key: 0
	Difficulty         MusicDifficulties   `json:"difficulty" msgpack:"difficulty"`                 // Key: 1
}
