// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Accessories
package dto

type Accessory struct {
	Id                int64            `json:"id" msgpack:"id"`                               // Key: 0
	AccessoryMasterId int64            `json:"accessoryMasterId" msgpack:"accessoryMasterId"` // Key: 1
	AccessoryMaster   *AccessoryMaster `json:"accessoryMaster" msgpack:"accessoryMaster"`
	Level             int              `json:"level" msgpack:"level"`                         // Key: 2
	Locked            bool             `json:"locked" msgpack:"locked"`                       // Key: 3
	AccessoryEffects  []int64          `json:"accessoryEffects" msgpack:"accessoryEffects"`   // Key: 4
	ReferenceCounting int              `json:"referenceCounting" msgpack:"referenceCounting"` // Key: 5
	IsFavorite        bool             `json:"isFavorite" msgpack:"isFavorite"`               // Key: 6
}

type AccessoryAutoSell struct {
	Id             int64                `json:"id" msgpack:"id"`                         // Key: 0
	AutoSellRarity PossessionRarityFlag `json:"autoSellRarity" msgpack:"autoSellRarity"` // Key: 1
}
