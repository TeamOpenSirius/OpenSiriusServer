// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Possession
package dto

import (
	"time"
)

type Bomb struct {
	Id            int64   `json:"id" msgpack:"id"`                       // Key: 0
	BombMasterIds []int64 `json:"bombMasterIds" msgpack:"bombMasterIds"` // Key: 1
}

type BuffItemStatus struct {
	Id               int64               `json:"id" msgpack:"id"`                             // Key: 0
	EffectType       CampaignEffectTypes `json:"effectType" msgpack:"effectType"`             // Key: 1
	BuffItemMasterId int64               `json:"buffItemMasterId" msgpack:"buffItemMasterId"` // Key: 2
	ValidUntil       time.Time           `json:"validUntil" msgpack:"validUntil"`             // Key: 3
}

type Costume struct {
	Id              int64 `json:"id" msgpack:"id"`                           // Key: 0
	CostumeMasterId int64 `json:"costumeMasterId" msgpack:"costumeMasterId"` // Key: 1
}

type Decoration struct {
	Id                 int64 `json:"id" msgpack:"id"`                                 // Key: 0
	DecorationMasterId int64 `json:"decorationMasterId" msgpack:"decorationMasterId"` // Key: 1
}

type FavoriteCostume struct {
	Id                       int64   `json:"id" msgpack:"id"`                                             // Key: 0
	CharacterBaseMasterId    int64   `json:"characterBaseMasterId" msgpack:"characterBaseMasterId"`       // Key: 1
	FavoriteCostumeMasterIds []int64 `json:"favoriteCostumeMasterIds" msgpack:"favoriteCostumeMasterIds"` // Key: 2
}

type HomeSkin struct {
	Id                int64   `json:"id" msgpack:"id"`                               // Key: 0
	HomeSkinMasterIds []int64 `json:"homeSkinMasterIds" msgpack:"homeSkinMasterIds"` // Key: 1
}

type IconFrame struct {
	Id                 int64   `json:"id" msgpack:"id"`                                 // Key: 0
	IconFrameMasterIds []int64 `json:"iconFrameMasterIds" msgpack:"iconFrameMasterIds"` // Key: 1
}

type Item struct {
	Id           int64 `json:"id" msgpack:"id"`                     // Key: 0
	ItemMasterId int64 `json:"itemMasterId" msgpack:"itemMasterId"` // Key: 1
	Stock        int   `json:"stock" msgpack:"stock"`               // Key: 2
}

type NameBaseColor struct {
	Id                     int64   `json:"id" msgpack:"id"`                                         // Key: 0
	NameBaseColorMasterIds []int64 `json:"nameBaseColorMasterIds" msgpack:"nameBaseColorMasterIds"` // Key: 1
}

type NameColor struct {
	Id                 int64   `json:"id" msgpack:"id"`                                 // Key: 0
	NameColorMasterIds []int64 `json:"nameColorMasterIds" msgpack:"nameColorMasterIds"` // Key: 1
}

type Nameplate struct {
	Id                int64 `json:"id" msgpack:"id"`                               // Key: 0
	NamePlateMasterId int64 `json:"namePlateMasterId" msgpack:"namePlateMasterId"` // Key: 1
}

type Note struct {
	Id            int64   `json:"id" msgpack:"id"`                       // Key: 0
	NoteMasterIds []int64 `json:"noteMasterIds" msgpack:"noteMasterIds"` // Key: 1
}

type Stamp struct {
	Id                     int64   `json:"id" msgpack:"id"`                                         // Key: 0
	StampMasterIds         []int64 `json:"stampMasterIds" msgpack:"stampMasterIds"`                 // Key: 1
	FavoriteStampMasterIds []int64 `json:"favoriteStampMasterIds" msgpack:"favoriteStampMasterIds"` // Key: 2
}

type Trophy struct {
	Id                  int64 `json:"id" msgpack:"id"`                                   // Key: 0
	TrophyMasterId      int64 `json:"trophyMasterId" msgpack:"trophyMasterId"`           // Key: 1
	TrophyGroupMasterId int64 `json:"trophyGroupMasterId" msgpack:"trophyGroupMasterId"` // Key: 2
	CurrentOrder        int   `json:"currentOrder" msgpack:"currentOrder"`               // Key: 3
}
