// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusRealtime.Shared.Entity.Circle
package dto

import (
	"time"
)

type ActivityLogCache struct {
	UserId                 string                   `json:"userId" msgpack:"userId"`                                 // Key: 0
	UserName               string                   `json:"userName" msgpack:"userName"`                             // Key: 1
	MainMCharacterId       int64                    `json:"mainMCharacterId" msgpack:"mainMCharacterId"`             // Key: 2
	DisplayAwakeningStatus bool                     `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 3
	LogType                RealtimeActivityLogTypes `json:"logType" msgpack:"logType"`                               // Key: 4
	LogValue               string                   `json:"logValue" msgpack:"logValue"`                             // Key: 5
}

type ActivityLogResult struct {
	ActivityLogId int64     `json:"activityLogId" msgpack:"activityLogId"` // Key: 6
	LoggedAt      time.Time `json:"loggedAt" msgpack:"loggedAt"`           // Key: 7
}

type ActivityLogNotification struct {
	CircleId          string             `json:"circleId" msgpack:"circleId"`                   // Key: 0
	ActivityLogResult *ActivityLogResult `json:"activityLogResult" msgpack:"activityLogResult"` // Key: 1
}

type CircleChatPayload struct {
	Message                string `json:"message" msgpack:"message"`                               // Key: 0
	StampId                *int64 `json:"stampId" msgpack:"stampId"`                               // Key: 1
	UserName               string `json:"userName" msgpack:"userName"`                             // Key: 2
	MainMCharacterId       int64  `json:"mainMCharacterId" msgpack:"mainMCharacterId"`             // Key: 3
	DisplayAwakeningStatus bool   `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 4
	MentionHashUserId      string `json:"mentionHashUserId" msgpack:"mentionHashUserId"`           // Key: 5
	IconFrameMasterId      int64  `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`           // Key: 6
}

type CircleChatCache struct {
	UserId                 string          `json:"userId" msgpack:"userId"`                                 // Key: 0
	Message                string          `json:"message" msgpack:"message"`                               // Key: 1
	StampId                *int64          `json:"stampId" msgpack:"stampId"`                               // Key: 2
	ChatType               CircleChatTypes `json:"chatType" msgpack:"chatType"`                             // Key: 3
	UserName               string          `json:"userName" msgpack:"userName"`                             // Key: 4
	MainMCharacterId       *int64          `json:"mainMCharacterId" msgpack:"mainMCharacterId"`             // Key: 5
	DisplayAwakeningStatus *bool           `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 6
	MentionHashUserId      string          `json:"mentionHashUserId" msgpack:"mentionHashUserId"`           // Key: 7
	IconFrameMasterId      int64           `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`           // Key: 8
}

type CircleChatResult struct {
	ChatId                 int64           `json:"chatId" msgpack:"chatId"`                                 // Key: 0
	UserId                 string          `json:"userId" msgpack:"userId"`                                 // Key: 1
	Message                string          `json:"message" msgpack:"message"`                               // Key: 2
	StampId                *int64          `json:"stampId" msgpack:"stampId"`                               // Key: 3
	SendTime               time.Time       `json:"sendTime" msgpack:"sendTime"`                             // Key: 4
	ChatType               CircleChatTypes `json:"chatType" msgpack:"chatType"`                             // Key: 5
	UserName               string          `json:"userName" msgpack:"userName"`                             // Key: 6
	MainMCharacterId       *int64          `json:"mainMCharacterId" msgpack:"mainMCharacterId"`             // Key: 7
	DisplayAwakeningStatus *bool           `json:"displayAwakeningStatus" msgpack:"displayAwakeningStatus"` // Key: 8
	MentionHashUserId      string          `json:"mentionHashUserId" msgpack:"mentionHashUserId"`           // Key: 9
	IconFrameMasterId      int64           `json:"iconFrameMasterId" msgpack:"iconFrameMasterId"`           // Key: 10
}

type CircleReadChat struct {
	LastReadChatId int64 `json:"lastReadChatId" msgpack:"lastReadChatId"` // Key: 0
	LatestChatId   int64 `json:"latestChatId" msgpack:"latestChatId"`     // Key: 1
}

type CircleReadChatCache struct {
	CircleId       string `json:"circleId" msgpack:"circleId"`             // Key: 0
	LastReadChatId int64  `json:"lastReadChatId" msgpack:"lastReadChatId"` // Key: 1
}

type CircleStatus struct {
	JoinStatus CircleJoinStatus `json:"joinStatus" msgpack:"joinStatus"` // Key: 0
}
