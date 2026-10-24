package userdata

import "fmt"

const (
	UnionUser                    int64 = 0
	UnionUserProfile             int64 = 1
	UnionUserPreference          int64 = 2
	UnionHomeDisplayPreference   int64 = 3
	UnionCharacter               int64 = 4
	UnionCharacterBase           int64 = 5
	UnionParty                   int64 = 6
	UnionPartySlot               int64 = 7
	UnionLive                    int64 = 24
	UnionMusic                   int64 = 25
	UnionViewedShop              int64 = 65
	UnionGameHint                int64 = 66
	UnionUserBonus               int64 = 67
	UnionAuditionClear           int64 = 90
	UnionSpRate                  int64 = 93
	UnionEpisode                 int64 = 95
	UnionCurrency                int64 = 128
	UnionLiveAchievement         int64 = 130
	UnionMusicBookmark           int64 = 146
	UnionLiveDropLimit           int64 = 147
	UnionRestriction             int64 = 148
	UnionTripleCastParty         int64 = 166
	UnionTripleCastPartySlot     int64 = 167
	UnionAlbumPreset             int64 = 174
	UnionFriendInvitation        int64 = 179
	UnionFriendInvitationMission int64 = 180
	UnionMalformed               int64 = -2
)

type UnionSpec struct {
	UnionID     int64
	Name        string
	MinArrayLen int
	IDIndex     int
}

var UserDataUnionRegistry = map[int64]UnionSpec{
	UnionUser:                    {UnionUser, "User", 20, 0},
	UnionUserProfile:             {UnionUserProfile, "UserProfile", 22, 0},
	UnionUserPreference:          {UnionUserPreference, "UserPreference", 3, 0},
	UnionHomeDisplayPreference:   {UnionHomeDisplayPreference, "HomeDisplayPreference", 16, 0},
	UnionCharacter:               {UnionCharacter, "Character", 15, 0},
	UnionCharacterBase:           {UnionCharacterBase, "CharacterBase", 8, 0},
	UnionParty:                   {UnionParty, "Party", 4, 0},
	UnionPartySlot:               {UnionPartySlot, "PartySlot", 7, 0},
	UnionLive:                    {UnionLive, "Live", 9, 0},
	UnionMusic:                   {UnionMusic, "Music", 10, 0},
	UnionViewedShop:              {UnionViewedShop, "ViewedShop", 4, 0},
	UnionGameHint:                {UnionGameHint, "GameHint", 3, 0},
	UnionUserBonus:               {UnionUserBonus, "UserBonus", 3, 0},
	UnionAuditionClear:           {UnionAuditionClear, "AuditionClear", 6, 0},
	UnionSpRate:                  {UnionSpRate, "SpRate", 4, 0},
	UnionEpisode:                 {UnionEpisode, "Episode", 2, 0},
	UnionCurrency:                {UnionCurrency, "Currency", 4, 0},
	UnionLiveAchievement:         {UnionLiveAchievement, "LiveAchievement", 3, 0},
	UnionMusicBookmark:           {UnionMusicBookmark, "MusicBookmark", 2, 0},
	UnionLiveDropLimit:           {UnionLiveDropLimit, "LiveDropLimit", 3, 0},
	UnionRestriction:             {UnionRestriction, "Restriction", 3, 0},
	UnionTripleCastParty:         {UnionTripleCastParty, "TripleCastParty", 4, 0},
	UnionTripleCastPartySlot:     {UnionTripleCastPartySlot, "TripleCastPartySlot", 7, 0},
	UnionAlbumPreset:             {UnionAlbumPreset, "AlbumPreset", 3, 0},
	UnionFriendInvitation:        {UnionFriendInvitation, "FriendInvitation", 3, 0},
	UnionFriendInvitationMission: {UnionFriendInvitationMission, "FriendInvitationMission", 6, 0},
}

func MakeRawObjectKey(unionID int64, data []any, index int) string {
	spec, ok := UserDataUnionRegistry[unionID]
	if ok && spec.IDIndex >= 0 && len(data) > spec.IDIndex {
		id := Int64Value(data[spec.IDIndex])
		if id != 0 {
			return fmt.Sprintf("%d:%d", unionID, id)
		}
	}
	return fmt.Sprintf("%d:idx:%d", unionID, index)
}

func SplitDataObject(raw any) (int64, []any, bool) {
	obj, ok := raw.([]any)
	if !ok || len(obj) < 2 {
		return 0, nil, false
	}
	data, ok := obj[1].([]any)
	if !ok {
		return 0, nil, false
	}
	return Int64Value(obj[0]), data, true
}

func Int64Value(v any) int64 {
	switch n := v.(type) {
	case int:
		return int64(n)
	case int8:
		return int64(n)
	case int16:
		return int64(n)
	case int32:
		return int64(n)
	case int64:
		return n
	case uint:
		return int64(n)
	case uint8:
		return int64(n)
	case uint16:
		return int64(n)
	case uint32:
		return int64(n)
	case uint64:
		return int64(n)
	case float64:
		return int64(n)
	default:
		return 0
	}
}
