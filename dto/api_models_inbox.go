// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Inbox
package dto

import (
	"time"
)

type Inbox struct {
	Id             int64      `json:"id" msgpack:"id"`                         // Key: 0
	ThingType      ThingTypes `json:"thingType" msgpack:"thingType"`           // Key: 1
	ThingId        int64      `json:"thingId" msgpack:"thingId"`               // Key: 2
	ThingQuantity  int        `json:"thingQuantity" msgpack:"thingQuantity"`   // Key: 3
	IsTimeLimited  bool       `json:"isTimeLimited" msgpack:"isTimeLimited"`   // Key: 4
	HasReceived    bool       `json:"hasReceived" msgpack:"hasReceived"`       // Key: 5
	Title          string     `json:"title" msgpack:"title"`                   // Key: 6
	Description    string     `json:"description" msgpack:"description"`       // Key: 7
	SentAt         time.Time  `json:"sentAt" msgpack:"sentAt"`                 // Key: 8
	ReceivedAt     *time.Time `json:"receivedAt" msgpack:"receivedAt"`         // Key: 9
	ReceiveLimitAt time.Time  `json:"receiveLimitAt" msgpack:"receiveLimitAt"` // Key: 10
}
