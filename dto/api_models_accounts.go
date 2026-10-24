// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Accounts
package dto

import (
	"time"
)

type ConnectWithAccount struct {
	Id       int64                   `json:"id" msgpack:"id"`             // Key: 0
	Provider AuthenticationProviders `json:"provider" msgpack:"provider"` // Key: 1
}

type ConnectWithPassword struct {
	Id int64 `json:"id" msgpack:"id"` // Key: 0
}

type Restriction struct {
	Id                             int64      `json:"id" msgpack:"id"`                                                         // Key: 0
	MultiLiveRestrictionFinishedAt *time.Time `json:"multiLiveRestrictionFinishedAt" msgpack:"multiLiveRestrictionFinishedAt"` // Key: 1
	ReadMultiLiveRestrictionDialog *bool      `json:"readMultiLiveRestrictionDialog" msgpack:"readMultiLiveRestrictionDialog"` // Key: 2
}
