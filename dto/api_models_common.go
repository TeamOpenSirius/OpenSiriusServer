// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Common
package dto

type AlbumDetailEditParameters struct {
	Items []AlbumDetailEditItem `json:"items" msgpack:"items"` // Key: 0
}

type AlbumDetailEditItem struct {
	ItemId        int64          `json:"itemId" msgpack:"itemId"`               // Key: 0
	AlbumItemType AlbumItemTypes `json:"albumItemType" msgpack:"albumItemType"` // Key: 1
	FileName      string         `json:"fileName" msgpack:"fileName"`           // Key: 2
	SasToken      string         `json:"sasToken" msgpack:"sasToken"`           // Key: 3
	X             float32        `json:"x" msgpack:"x"`                         // Key: 4
	Y             float32        `json:"y" msgpack:"y"`                         // Key: 5
	Rotate        float32        `json:"rotate" msgpack:"rotate"`               // Key: 6
	Scale         float32        `json:"scale" msgpack:"scale"`                 // Key: 7
	Layer         int            `json:"layer" msgpack:"layer"`                 // Key: 8
	Visible       bool           `json:"visible" msgpack:"visible"`             // Key: 9
	Lock          bool           `json:"lock" msgpack:"lock"`                   // Key: 10
	SignMasterId  *int64         `json:"signMasterId" msgpack:"signMasterId"`   // Key: 11
}

type AlbumSimpleEditParameters struct {
	Photos []AlbumSimpleEditPhoto `json:"photos" msgpack:"photos"` // Key: 0
}

type AlbumSimpleEditPhoto struct {
	PhotoId      int64  `json:"photoId" msgpack:"photoId"`           // Key: 0
	FileName     string `json:"fileName" msgpack:"fileName"`         // Key: 1
	SasToken     string `json:"sasToken" msgpack:"sasToken"`         // Key: 2
	Order        int    `json:"order" msgpack:"order"`               // Key: 3
	SignMasterId *int64 `json:"signMasterId" msgpack:"signMasterId"` // Key: 4
}

type HelpDialog struct {
	Categories []HelpCategory `json:"categories" msgpack:"categories"` // Key: 0
}

type HelpCategory struct {
	CategoryName string     `json:"categoryName" msgpack:"categoryName"` // Key: 0
	HelpTexts    []HelpText `json:"helpTexts" msgpack:"helpTexts"`       // Key: 1
}

type HelpText struct {
	Title string `json:"title" msgpack:"title"` // Key: 0
	Order int    `json:"order" msgpack:"order"` // Key: 1
	Body  string `json:"body" msgpack:"body"`   // Key: 2
}
