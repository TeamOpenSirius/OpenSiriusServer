// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Photo
package dto

import (
	"time"
)

type Album struct {
	Id                 int64 `json:"id" msgpack:"id"`                                 // Key: 0
	Level              int   `json:"level" msgpack:"level"`                           // Key: 1
	PublishPageNumber  int   `json:"publishPageNumber" msgpack:"publishPageNumber"`   // Key: 2
	CurrentPresetOrder int   `json:"currentPresetOrder" msgpack:"currentPresetOrder"` // Key: 3
}

type AlbumPage struct {
	Id                 int64     `json:"id" msgpack:"id"`                                 // Key: 0
	Page               int       `json:"page" msgpack:"page"`                             // Key: 1
	EditType           EditTypes `json:"editType" msgpack:"editType"`                     // Key: 2
	Publishing         bool      `json:"publishing" msgpack:"publishing"`                 // Key: 3
	Items              []byte    `json:"items" msgpack:"items"`                           // Key: 4
	AlbumThemeMasterId *int64    `json:"albumThemeMasterId" msgpack:"albumThemeMasterId"` // Key: 5
}

type AlbumPreset struct {
	Id    int64  `json:"id" msgpack:"id"`       // Key: 0
	Name  string `json:"name" msgpack:"name"`   // Key: 1
	Order int    `json:"order" msgpack:"order"` // Key: 2
}

type AlbumTheme struct {
	Id                 int64 `json:"id" msgpack:"id"`                                 // Key: 0
	AlbumThemeMasterId int64 `json:"albumThemeMasterId" msgpack:"albumThemeMasterId"` // Key: 1
}

type MusicVideo struct {
	Id                 int64 `json:"id" msgpack:"id"`                                 // Key: 0
	MusicVideoMasterId int64 `json:"musicVideoMasterId" msgpack:"musicVideoMasterId"` // Key: 1
}

type Photo struct {
	Id                             int64           `json:"id" msgpack:"id"`                                                         // Key: 0
	FileName                       string          `json:"fileName" msgpack:"fileName"`                                             // Key: 1
	SasToken                       string          `json:"sasToken" msgpack:"sasToken"`                                             // Key: 2
	PhotoEffectMasterId            *int64          `json:"photoEffectMasterId" msgpack:"photoEffectMasterId"`                       // Key: 3
	Lock                           bool            `json:"lock" msgpack:"lock"`                                                     // Key: 4
	UseAlbumPage                   *int            `json:"useAlbumPage" msgpack:"useAlbumPage"`                                     // Key: 5
	Level                          int             `json:"level" msgpack:"level"`                                                   // Key: 6
	Rarity                         PhotoRarities   `json:"rarity" msgpack:"rarity"`                                                 // Key: 7
	SignMasterId                   *int64          `json:"signMasterId" msgpack:"signMasterId"`                                     // Key: 8
	GeneratedAt                    time.Time       `json:"generatedAt" msgpack:"generatedAt"`                                       // Key: 9
	ThumbnailSasToken              string          `json:"thumbnailSasToken" msgpack:"thumbnailSasToken"`                           // Key: 10
	AppearedCharacterBaseMasterIds []int64         `json:"appearedCharacterBaseMasterIds" msgpack:"appearedCharacterBaseMasterIds"` // Key: 11
	TaggedCharacterBaseMasterIds   []int64         `json:"taggedCharacterBaseMasterIds" msgpack:"taggedCharacterBaseMasterIds"`     // Key: 12
	UseDecoPage                    UseDecoPageFlag `json:"useDecoPage" msgpack:"useDecoPage"`                                       // Key: 13
}

type TheaterStory struct {
	Id                   int64 `json:"id" msgpack:"id"`                                     // Key: 0
	TheaterStoryMasterId int64 `json:"theaterStoryMasterId" msgpack:"theaterStoryMasterId"` // Key: 1
}
