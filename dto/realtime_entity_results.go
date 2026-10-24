// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusRealtime.Shared.Entity.Results
package dto

import (
	"time"
)

type CircleChatSendResult struct {
	ErrorCode CircleChatSendErrorCodes `json:"errorCode" msgpack:"errorCode"` // Key: 0
}

type MultiLiveCreatePrivateHallResult struct {
	IsSucceeded           bool       `json:"isSucceeded" msgpack:"isSucceeded"`                     // Key: 0
	HallId                string     `json:"hallId" msgpack:"hallId"`                               // Key: 1
	MemberId              int        `json:"memberId" msgpack:"memberId"`                           // Key: 2
	PrivateHallKeyCode    int        `json:"privateHallKeyCode" msgpack:"privateHallKeyCode"`       // Key: 3
	RestrictionFinishedAt *time.Time `json:"restrictionFinishedAt" msgpack:"restrictionFinishedAt"` // Key: 4
}

type MultiLiveFetchUsersResult struct {
	HostMemberId int             `json:"hostMemberId" msgpack:"hostMemberId"` // Key: 0
	Users        []MultiLiveUser `json:"users" msgpack:"users"`               // Key: 1
	CanOpenHall  bool            `json:"canOpenHall" msgpack:"canOpenHall"`   // Key: 2
}

type MultiLiveJoinResult struct {
	ErrorCode             MultiLiveJoinErrorCodes `json:"errorCode" msgpack:"errorCode"`                         // Key: 0
	HallId                string                  `json:"hallId" msgpack:"hallId"`                               // Key: 1
	MemberId              int                     `json:"memberId" msgpack:"memberId"`                           // Key: 2
	KeyCode               int                     `json:"keyCode" msgpack:"keyCode"`                             // Key: 3
	LiveSettingMasterId   int64                   `json:"liveSettingMasterId" msgpack:"liveSettingMasterId"`     // Key: 4
	RestrictionFinishedAt *time.Time              `json:"restrictionFinishedAt" msgpack:"restrictionFinishedAt"` // Key: 5
}

type MultiLiveMatchingAssignmentResult struct {
	MatchingAssignments []MultiLiveMatchingAssignment `json:"matchingAssignments" msgpack:"matchingAssignments"` // Key: 0
}
