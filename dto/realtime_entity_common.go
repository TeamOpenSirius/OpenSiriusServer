// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusRealtime.Shared.Entity.Common
package dto

import (
	"time"
)

type MultiLiveInvitation struct {
	InviteId               string         `json:"inviteId" msgpack:"inviteId"`                             // Key: 0
	UserName               string         `json:"userName" msgpack:"userName"`                             // Key: 1
	CharacterMasterId      int64          `json:"characterMasterId" msgpack:"characterMasterId"`           // Key: 2
	DisplayAwakeningStatus bool           `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 3
	PlayerRank             int            `json:"playerRank" msgpack:"playerRank"`                         // Key: 4
	TrophyMasterId1        *int64         `json:"trophyMasterId1" msgpack:"trophyMasterId1"`               // Key: 5
	TrophyMasterId2        *int64         `json:"trophyMasterId2" msgpack:"trophyMasterId2"`               // Key: 6
	TrophyMasterId3        *int64         `json:"trophyMasterId3" msgpack:"trophyMasterId3"`               // Key: 7
	MultiLiveType          MultiLiveTypes `json:"multiLiveType" msgpack:"multiLiveType"`                   // Key: 8
	HallId                 string         `json:"hallId" msgpack:"hallId"`                                 // Key: 9
	MemberCount            int            `json:"memberCount" msgpack:"memberCount"`                       // Key: 10
	InvitedAt              time.Time      `json:"invitedAt" msgpack:"invitedAt"`                           // Key: 11
	IconFrameMasterId      int64          `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`           // Key: 12
}

type MultiLiveInvitationCache struct {
	UserId                 string         `json:"userId" msgpack:"userId"`                                 // Key: 0
	UserName               string         `json:"userName" msgpack:"userName"`                             // Key: 1
	CharacterMasterId      int64          `json:"characterMasterId" msgpack:"characterMasterId"`           // Key: 2
	DisplayAwakeningStatus bool           `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 3
	PlayerRank             int            `json:"playerRank" msgpack:"playerRank"`                         // Key: 4
	TrophyMasterId1        *int64         `json:"trophyMasterId1" msgpack:"trophyMasterId1"`               // Key: 5
	TrophyMasterId2        *int64         `json:"trophyMasterId2" msgpack:"trophyMasterId2"`               // Key: 6
	TrophyMasterId3        *int64         `json:"trophyMasterId3" msgpack:"trophyMasterId3"`               // Key: 7
	MultiLiveType          MultiLiveTypes `json:"multiLiveType" msgpack:"multiLiveType"`                   // Key: 8
	HallId                 string         `json:"hallId" msgpack:"hallId"`                                 // Key: 9
	IconFrameMasterId      int64          `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`           // Key: 10
}
