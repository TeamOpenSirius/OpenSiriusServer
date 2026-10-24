// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.FriendInvitation
package dto

type FriendInvitation struct {
	Id                          int64  `json:"id" msgpack:"id"`                                                   // Key: 0
	InvitationCode              string `json:"invitationCode" msgpack:"invitationCode"`                           // Key: 1
	HasInputOtherInvitationCode bool   `json:"hasInputOtherInvitationCode" msgpack:"hasInputOtherInvitationCode"` // Key: 2
}

type FriendInvitationMission struct {
	Id                                   int64 `json:"id" msgpack:"id"`                                                                     // Key: 0
	FriendInvitationMissionMasterId      int64 `json:"friendInvitationMissionMasterId" msgpack:"friendInvitationMissionMasterId"`           // Key: 1
	FriendInvitationMissionStageMasterId int64 `json:"friendInvitationMissionStageMasterId" msgpack:"friendInvitationMissionStageMasterId"` // Key: 2
	CurrentCount                         int   `json:"currentCount" msgpack:"currentCount"`                                                 // Key: 3
	IsCleared                            bool  `json:"isCleared" msgpack:"isCleared"`                                                       // Key: 4
	IsRewardReceived                     bool  `json:"isRewardReceived" msgpack:"isRewardReceived"`                                         // Key: 5
}
