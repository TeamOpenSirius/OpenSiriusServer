// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Gacha
package dto

type Gacha struct {
	Id            int64 `json:"id" msgpack:"id"`                       // Key: 0
	GachaMasterId int64 `json:"gachaMasterId" msgpack:"gachaMasterId"` // Key: 1
	RollCount     int   `json:"rollCount" msgpack:"rollCount"`         // Key: 2
}

type GachaReRoll struct {
	GachaMasterId int64 `json:"gachaMasterId" msgpack:"gachaMasterId"` // Key: 0
	RollCount     int   `json:"rollCount" msgpack:"rollCount"`         // Key: 1
	IsDecided     bool  `json:"isDecided" msgpack:"isDecided"`         // Key: 2
}

type GachaSelectedThing struct {
	GachaMasterId int64   `json:"gachaMasterId" msgpack:"gachaMasterId"` // Key: 0
	GachaThingIds []int64 `json:"gachaThingIds" msgpack:"gachaThingIds"` // Key: 1
}
