// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusRealtime.Shared.Entity.Payloads
package dto

type MultiLiveMatchMakingIgnorePayload struct {
	GroupNames []string `json:"groupNames" msgpack:"groupNames"` // Key: 0
}

type MultiLiveMatchMakingRebornPayload struct {
	GroupNames []string `json:"groupNames" msgpack:"groupNames"` // Key: 0
}

type MultiLiveMatchMakingUserTicketsPayload struct {
	UserTickets []MultiLiveMatchMakingUserTicket `json:"userTickets" msgpack:"userTickets"` // Key: 0
}

type MultiLiveRecruitingRoomTicketsPayload struct {
	RoomTickets []MultiLiveRecruitingRoomTicket `json:"roomTickets" msgpack:"roomTickets"` // Key: 0
}

type UserNotificationPayload struct {
	MultiLiveRequestNotifications          []MultiLiveRequestNotification          `json:"multiLiveRequestNotifications" msgpack:"multiLiveRequestNotifications"`                   // Key: 0
	CircleChatSystemNotifications          []CircleChatSystemNotification          `json:"circleChatSystemNotifications" msgpack:"circleChatSystemNotifications"`                   // Key: 1
	ActivityLogNotifications               []ActivityLogNotification               `json:"activityLogNotifications" msgpack:"activityLogNotifications"`                             // Key: 2
	MultiLiveRequestForFriendNotifications []MultiLiveRequestForFriendNotification `json:"multiLiveRequestForFriendNotifications" msgpack:"multiLiveRequestForFriendNotifications"` // Key: 3
}
