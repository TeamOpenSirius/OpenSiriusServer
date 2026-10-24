// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusRealtime.Shared.Entity.Notification
package dto

import (
	"time"
)

type CircleChatSystemNotification struct {
	ChatId                 int64           `json:"chatId" msgpack:"chatId"`                                 // Key: 0
	CircleId               string          `json:"circleId" msgpack:"circleId"`                             // Key: 1
	ChatType               CircleChatTypes `json:"chatType" msgpack:"chatType"`                             // Key: 2
	SendTime               time.Time       `json:"sendTime" msgpack:"sendTime"`                             // Key: 3
	Message                string          `json:"message" msgpack:"message"`                               // Key: 4
	UserName               string          `json:"userName" msgpack:"userName"`                             // Key: 5
	MainMCharacterId       *int64          `json:"mainMCharacterId" msgpack:"mainMCharacterId"`             // Key: 6
	UserId                 string          `json:"userId" msgpack:"userId"`                                 // Key: 7
	DisplayAwakeningStatus *bool           `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 8
}

type MultiLiveRequestNotification struct {
	CircleId         string            `json:"circleId" msgpack:"circleId"`                 // Key: 0
	HallId           string            `json:"hallId" msgpack:"hallId"`                     // Key: 1
	UserId           string            `json:"userId" msgpack:"userId"`                     // Key: 2
	RequestUserName  string            `json:"requestUserName" msgpack:"requestUserName"`   // Key: 3
	MemberCount      int               `json:"memberCount" msgpack:"memberCount"`           // Key: 4
	CircleChatResult *CircleChatResult `json:"circleChatResult" msgpack:"circleChatResult"` // Key: 5
	MultiLiveType    MultiLiveTypes    `json:"multiLiveType" msgpack:"multiLiveType"`       // Key: 6
}

type MultiLiveRequestForFriendNotification struct {
	FriendUserId    string         `json:"friendUserId" msgpack:"friendUserId"`       // Key: 0
	HallId          string         `json:"hallId" msgpack:"hallId"`                   // Key: 1
	UserId          string         `json:"userId" msgpack:"userId"`                   // Key: 2
	RequestUserName string         `json:"requestUserName" msgpack:"requestUserName"` // Key: 3
	MultiLiveType   MultiLiveTypes `json:"multiLiveType" msgpack:"multiLiveType"`     // Key: 4
}
