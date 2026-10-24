// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusRealtime.Shared
package dto

type MultiLiveHallType int32

const (
	MultiLiveHallTypeNone          MultiLiveHallType = 0
	MultiLiveHallTypeSirius        MultiLiveHallType = 11
	MultiLiveHallTypeEden          MultiLiveHallType = 12
	MultiLiveHallTypeGingaza       MultiLiveHallType = 13
	MultiLiveHallTypeDenki         MultiLiveHallType = 14
	MultiLiveHallTypeDenkiReprint  MultiLiveHallType = 15
	MultiLiveHallTypeTeamChallenge MultiLiveHallType = 21
)

type ComboTypes int32

const (
	ComboTypesNone       ComboTypes = 0
	ComboTypesFullCombo  ComboTypes = 1
	ComboTypesAllPerfect ComboTypes = 2
)

type MultiLiveJoinErrorCodes int32

const (
	MultiLiveJoinErrorCodesNone                          MultiLiveJoinErrorCodes = 0
	MultiLiveJoinErrorCodesReachedMaxMember              MultiLiveJoinErrorCodes = 1
	MultiLiveJoinErrorCodesRoomClosed                    MultiLiveJoinErrorCodes = 2
	MultiLiveJoinErrorCodesNotFoundHall                  MultiLiveJoinErrorCodes = 3
	MultiLiveJoinErrorCodesNotSubscribeMatchMakingServer MultiLiveJoinErrorCodes = 4
	MultiLiveJoinErrorCodesUserCanceled                  MultiLiveJoinErrorCodes = 5
	MultiLiveJoinErrorCodesRestricting                   MultiLiveJoinErrorCodes = 6
)

type MultiLiveMusicSelectionTypes int32

const (
	MultiLiveMusicSelectionTypesNone    MultiLiveMusicSelectionTypes = 0
	MultiLiveMusicSelectionTypesSpecify MultiLiveMusicSelectionTypes = 1
	MultiLiveMusicSelectionTypesRandom  MultiLiveMusicSelectionTypes = 2
)

type MultiLiveUserStatus int32

const (
	MultiLiveUserStatusJoined              MultiLiveUserStatus = 0
	MultiLiveUserStatusSelectedMusic       MultiLiveUserStatus = 1
	MultiLiveUserStatusSelectedDifficulty  MultiLiveUserStatus = 2
	MultiLiveUserStatusReadyGame           MultiLiveUserStatus = 3
	MultiLiveUserStatusBeforeGameCalculate MultiLiveUserStatus = 4
	MultiLiveUserStatusPlayingGame         MultiLiveUserStatus = 5
	MultiLiveUserStatusExitGame            MultiLiveUserStatus = 6
	MultiLiveUserStatusDisconnected        MultiLiveUserStatus = 7
)

type MultiLiveHallStatus int32

const (
	MultiLiveHallStatusNone                MultiLiveHallStatus = 0
	MultiLiveHallStatusRecruiting          MultiLiveHallStatus = 1
	MultiLiveHallStatusStandbyGame         MultiLiveHallStatus = 2
	MultiLiveHallStatusReadyGame           MultiLiveHallStatus = 3
	MultiLiveHallStatusBeforeGameCalculate MultiLiveHallStatus = 4
	MultiLiveHallStatusPlayingGame         MultiLiveHallStatus = 5
)

type UserNotificationTypes int32

const (
	UserNotificationTypesFriendRequest UserNotificationTypes = 0
	UserNotificationTypesMultiLive     UserNotificationTypes = 1
)

type CircleChatTypes int32

const (
	CircleChatTypesDefault            CircleChatTypes = 0
	CircleChatTypesJoinCircle         CircleChatTypes = 1
	CircleChatTypesLeaveCircle        CircleChatTypes = 2
	CircleChatTypesChangedPosition    CircleChatTypes = 3
	CircleChatTypesRecuitingMultiLive CircleChatTypes = 4
)

type CircleChatSendErrorCodes int32

const (
	CircleChatSendErrorCodesNone            CircleChatSendErrorCodes = 0
	CircleChatSendErrorCodesShortSpanSend   CircleChatSendErrorCodes = 1
	CircleChatSendErrorCodesLongWordLength  CircleChatSendErrorCodes = 2
	CircleChatSendErrorCodesNgWordContained CircleChatSendErrorCodes = 3
	CircleChatSendErrorCodesNotFinishedJoin CircleChatSendErrorCodes = 4
)

type CircleJoinStatus int32

const (
	CircleJoinStatusNone         CircleJoinStatus = 0
	CircleJoinStatusSuccess      CircleJoinStatus = 1
	CircleJoinStatusCircleIdNone CircleJoinStatus = 2
)

type ForceDisconnectReason int32

const (
	ForceDisconnectReasonNone           ForceDisconnectReason = 0
	ForceDisconnectReasonPlayedContinue ForceDisconnectReason = 1
)

type RealtimeActivityLogTypes int32

const (
	RealtimeActivityLogTypesNone                     RealtimeActivityLogTypes = 0
	RealtimeActivityLogTypesJoinCircle               RealtimeActivityLogTypes = 1
	RealtimeActivityLogTypesLeaveCircle              RealtimeActivityLogTypes = 2
	RealtimeActivityLogTypesChangedPosition          RealtimeActivityLogTypes = 3
	RealtimeActivityLogTypesGachaAcquiredCharacter   RealtimeActivityLogTypes = 4
	RealtimeActivityLogTypesGachaAcquiredPoster      RealtimeActivityLogTypes = 5
	RealtimeActivityLogTypesPlayerRate               RealtimeActivityLogTypes = 6
	RealtimeActivityLogTypesFirstClearedAllPerfect   RealtimeActivityLogTypes = 7
	RealtimeActivityLogTypesFirstClearedFullCombo    RealtimeActivityLogTypes = 8
	RealtimeActivityLogTypesDonateSupportLevelLimit  RealtimeActivityLogTypes = 9
	RealtimeActivityLogTypesReleaseSupportLevelLimit RealtimeActivityLogTypes = 10
)

type MultiLiveTypes int32

const (
	MultiLiveTypesDefault       MultiLiveTypes = 1
	MultiLiveTypesTeamChallenge MultiLiveTypes = 2
)
