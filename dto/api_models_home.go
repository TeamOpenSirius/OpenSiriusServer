// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Home
package dto

import (
	"time"
)

type HomeBGM struct {
	HomeBGMMasterId       int64                 `json:"homeBGMMasterId" msgpack:"homeBGMMasterId"`             // Key: 0
	SelectionType         HomeBGMSelectionTypes `json:"selectionType" msgpack:"selectionType"`                 // Key: 1
	HomeBGMDetailMasterId *int64                `json:"homeBGMDetailMasterId" msgpack:"homeBGMDetailMasterId"` // Key: 2
}

type Lottery struct {
	LotteryMasterId int64                `json:"lotteryMasterId" msgpack:"lotteryMasterId"` // Key: 0
	Results         []LotteryPriseResult `json:"results" msgpack:"results"`                 // Key: 1
}

type Notification struct {
	Id              int64     `json:"id" msgpack:"id"`                           // Key: 0
	ImportantReadAt time.Time `json:"importantReadAt" msgpack:"importantReadAt"` // Key: 1
	UpdateReadAt    time.Time `json:"updateReadAt" msgpack:"updateReadAt"`       // Key: 2
	BugReadAt       time.Time `json:"bugReadAt" msgpack:"bugReadAt"`             // Key: 3
}

type TimeLimitedControl struct {
	Id                         int64     `json:"id" msgpack:"id"`                                                 // Key: 0
	TimeLimitedControlMasterId int64     `json:"timeLimitedControlMasterId" msgpack:"timeLimitedControlMasterId"` // Key: 1
	ExpiredAt                  time.Time `json:"expiredAt" msgpack:"expiredAt"`                                   // Key: 2
}

type LotteryPriseResult struct {
	Prize  int `json:"prize" msgpack:"prize"`   // Key: 0
	Amount int `json:"amount" msgpack:"amount"` // Key: 1
}
