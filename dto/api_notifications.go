// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Notifications
package dto

import (
	"time"
)

type AcceptFriendRequest struct {
	RequestUserId    string `json:"requestUserId" msgpack:"requestUserId"`       // Key: 0
	ReceivedUserId   string `json:"receivedUserId" msgpack:"receivedUserId"`     // Key: 1
	ReceivedUserName string `json:"receivedUserName" msgpack:"receivedUserName"` // Key: 2
}

type FriendRequest struct {
	RequestUserId   string `json:"requestUserId" msgpack:"requestUserId"`     // Key: 0
	ReceivedUserId  string `json:"receivedUserId" msgpack:"receivedUserId"`   // Key: 1
	RequestUserName string `json:"requestUserName" msgpack:"requestUserName"` // Key: 2
}

type MissionCleared struct {
	MissionMasterId      int64 `json:"missionMasterId" msgpack:"missionMasterId"`           // Key: 0
	MissionStageMasterId int64 `json:"missionStageMasterId" msgpack:"missionStageMasterId"` // Key: 1
}

type MultiLiveRestriction struct {
	RestrictionFinishedAt *time.Time `json:"restrictionFinishedAt" msgpack:"restrictionFinishedAt"` // Key: 0
}
