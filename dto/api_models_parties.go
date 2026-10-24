// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Parties
package dto

type Party struct {
	Id             int64       `json:"id" msgpack:"id"`                         // Key: 0
	Order          int         `json:"order" msgpack:"order"`                   // Key: 1
	Name           string      `json:"name" msgpack:"name"`                     // Key: 2
	LeaderPosition int         `json:"leaderPosition" msgpack:"leaderPosition"` // Key: 3
	Slots          []PartySlot `json:"slots" msgpack:"slots"`
}

type PartySlot struct {
	Id                      int64                   `json:"id" msgpack:"id"`           // Key: 0
	PartyId                 int64                   `json:"partyId" msgpack:"partyId"` // Key: 1
	Party                   *Party                  `json:"party" msgpack:"party"`
	Position                int                     `json:"position" msgpack:"position"`       // Key: 2
	CharacterId             int64                   `json:"characterId" msgpack:"characterId"` // Key: 3
	Character               *Character              `json:"character" msgpack:"character"`
	PosterId                *int64                  `json:"posterId" msgpack:"posterId"` // Key: 4
	Poster                  *Poster                 `json:"poster" msgpack:"poster"`
	AccessoryId             *int64                  `json:"accessoryId" msgpack:"accessoryId"` // Key: 5
	Accessory               *Accessory              `json:"accessory" msgpack:"accessory"`
	BonusAbilityEnableFlags BonusAbilityEnableFlags `json:"bonusAbilityEnableFlags" msgpack:"bonusAbilityEnableFlags"` // Key: 6
}

type TripleCastParty struct {
	Id             int64                 `json:"id" msgpack:"id"`                         // Key: 0
	Order          int                   `json:"order" msgpack:"order"`                   // Key: 1
	Name           string                `json:"name" msgpack:"name"`                     // Key: 2
	LeaderPosition int                   `json:"leaderPosition" msgpack:"leaderPosition"` // Key: 3
	Slots          []TripleCastPartySlot `json:"slots" msgpack:"slots"`
}

type TripleCastPartySlot struct {
	Id                      int64                   `json:"id" msgpack:"id"`                   // Key: 0
	PartyId                 int64                   `json:"partyId" msgpack:"partyId"`         // Key: 1
	Position                int                     `json:"position" msgpack:"position"`       // Key: 2
	CharacterId             int64                   `json:"characterId" msgpack:"characterId"` // Key: 3
	Character               *Character              `json:"character" msgpack:"character"`
	PosterId                *int64                  `json:"posterId" msgpack:"posterId"` // Key: 4
	Poster                  *Poster                 `json:"poster" msgpack:"poster"`
	AccessoryId             *int64                  `json:"accessoryId" msgpack:"accessoryId"` // Key: 5
	Accessory               *Accessory              `json:"accessory" msgpack:"accessory"`
	BonusAbilityEnableFlags BonusAbilityEnableFlags `json:"bonusAbilityEnableFlags" msgpack:"bonusAbilityEnableFlags"` // Key: 6
}
