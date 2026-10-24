// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Missions
package dto

import (
	"time"
)

type FlashSaleStage struct {
	Id                     int64      `json:"id" msgpack:"id"`                                         // Key: 0
	FlashSaleStageMasterId int64      `json:"flashSaleStageMasterId" msgpack:"flashSaleStageMasterId"` // Key: 1
	PurchaseLimitedAt      *time.Time `json:"purchaseLimitedAt" msgpack:"purchaseLimitedAt"`           // Key: 2
	IsDefault              bool       `json:"isDefault" msgpack:"isDefault"`                           // Key: 3
	IsCompleted            bool       `json:"isCompleted" msgpack:"isCompleted"`                       // Key: 4
}

type GradualMissionGroup struct {
	Id                          int64     `json:"id" msgpack:"id"`                                                   // Key: 0
	GradualMissionGroupMasterId int64     `json:"gradualMissionGroupMasterId" msgpack:"gradualMissionGroupMasterId"` // Key: 1
	StartAt                     time.Time `json:"startAt" msgpack:"startAt"`                                         // Key: 2
}

type Mission struct {
	Id                          int64 `json:"id" msgpack:"id"`                                                   // Key: 0
	IsCleared                   bool  `json:"isCleared" msgpack:"isCleared"`                                     // Key: 1
	IsRewardReceived            bool  `json:"isRewardReceived" msgpack:"isRewardReceived"`                       // Key: 2
	MissionCurrentCount         int64 `json:"missionCurrentCount" msgpack:"missionCurrentCount"`                 // Key: 3
	MissionMasterId             int64 `json:"missionMasterId" msgpack:"missionMasterId"`                         // Key: 5
	CurrentMissionStageMasterId int64 `json:"currentMissionStageMasterId" msgpack:"currentMissionStageMasterId"` // Key: 6
}

type MissionPass struct {
	Id                          int64 `json:"id" msgpack:"id"`                                                   // Key: 0
	UserId                      int64 `json:"userId" msgpack:"userId"`                                           // Key: 1
	MissionPassMasterId         int64 `json:"missionPassMasterId" msgpack:"missionPassMasterId"`                 // Key: 2
	Paid                        bool  `json:"paid" msgpack:"paid"`                                               // Key: 3
	FreeRewardReceivedPhase     *int  `json:"freeRewardReceivedPhase" msgpack:"freeRewardReceivedPhase"`         // Key: 4
	SpRewardReceivedPhase       *int  `json:"spRewardReceivedPhase" msgpack:"spRewardReceivedPhase"`             // Key: 5
	Terminated                  bool  `json:"terminated" msgpack:"terminated"`                                   // Key: 6
	FreeRewardLoopCount         *int  `json:"freeRewardLoopCount" msgpack:"freeRewardLoopCount"`                 // Key: 7
	FreeRewardLoopReceivedPhase *int  `json:"freeRewardLoopReceivedPhase" msgpack:"freeRewardLoopReceivedPhase"` // Key: 8
	PaidRewardLoopCount         *int  `json:"paidRewardLoopCount" msgpack:"paidRewardLoopCount"`                 // Key: 9
	PaidRewardLoopReceivedPhase *int  `json:"paidRewardLoopReceivedPhase" msgpack:"paidRewardLoopReceivedPhase"` // Key: 10
}

type PickupCharacterMission struct {
	PickupCharacterMissionMasterId int64   `json:"pickupCharacterMissionMasterId" msgpack:"pickupCharacterMissionMasterId"` // Key: 0
	ReceivedDetailMasterIds        []int64 `json:"receivedDetailMasterIds" msgpack:"receivedDetailMasterIds"`               // Key: 1
}
