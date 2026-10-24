// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Posters
package dto

type Poster struct {
	Id                           int64              `json:"id" msgpack:"id"`                         // Key: 0
	PosterMasterId               int64              `json:"posterMasterId" msgpack:"posterMasterId"` // Key: 1
	PosterMaster                 *PosterMaster      `json:"posterMaster" msgpack:"posterMaster"`
	Level                        int                `json:"level" msgpack:"level"`                                               // Key: 2
	BreakthroughPhase            int                `json:"breakthroughPhase" msgpack:"breakthroughPhase"`                       // Key: 3
	ReleasedEpisode              PosterEpisodeTypes `json:"releasedEpisode" msgpack:"releasedEpisode"`                           // Key: 4
	ItemConsumeBreakThroughCount int                `json:"itemConsumeBreakThroughCount" msgpack:"itemConsumeBreakThroughCount"` // Key: 5
	IsFavorite                   bool               `json:"isFavorite" msgpack:"isFavorite"`                                     // Key: 6
	AlternativeImagePattern      int                `json:"alternativeImagePattern" msgpack:"alternativeImagePattern"`           // Key: 7
}
