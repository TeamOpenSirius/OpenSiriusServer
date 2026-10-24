package api

import (
	"ymst_server/dto"
)

func ArrayByKey(size int, values map[int]any) []any {
	out := make([]any, size)
	for key, value := range values {
		if key >= 0 && key < size {
			out[key] = value
		}
	}
	return out
}

func DataObject(unionID int64, arr []any) []any {
	return []any{unionID, arr}
}

func UserToArray(v dto.User) []any {
	return []any{
		v.Id,
		v.PlayerRank,
		v.CurrentRankPoint,
		v.CurrentStamina,
		v.MaxStaminaRestoredAt,
		v.PaidJewel,
		v.FreeJewel,
		v.Coin,
		v.PlayerRankLimit,
		v.StaminaRecoverTimesWithJewel,
		v.CircleUsageRestrictionsEndTime,
		v.CircleId,
		v.GameStartAt,
		v.HashUserId,
		int32(v.BanLevel),
		int32(v.TutorialStatus),
		v.MonthlyPayment,
		v.SplashLastDisplayedAt,
		v.IsCapedPlayerRank,
		v.RequireCapedPlayerRankAnnounce,
	}
}

func UserProfileToArray(v dto.ApiPlayerUserProfile) []any {
	return ArrayByKey(22, map[int]any{
		0:  v.Id,
		1:  v.Name,
		2:  v.Introduction,
		3:  v.MainUCharacterId,
		4:  v.MNameplateId,
		5:  v.MNameColorId,
		6:  v.MTrophyId1,
		7:  v.MTrophyId2,
		8:  v.MTrophyId3,
		9:  v.PlayerRate,
		11: v.IsPublicPlayerRate,
		12: int32(v.LeagueClass),
		13: v.TotalSpCount,
		14: v.IsPublicAlbumMainPage,
		15: v.MNameplateDetailId,
		16: v.MainCharacterMasterId,
		17: v.DisplayAwakeningStatus,
		18: v.IsPublicActivityLog,
		19: v.NameBaseColorMasterId,
		20: v.IconFrameMasterId,
		21: v.HomeSkinMasterId,
	})
}

func UserPreferenceToArray(v dto.UserPreference) []any {
	return []any{v.Id, v.MultiPartyId, v.BirthDate}
}

func HomeDisplayPreferenceToArray(v dto.HomeDisplayPreference) []any {
	return ArrayByKey(16, map[int]any{
		0:  v.Id,
		1:  v.HomeCharacterBaseMasterId,
		3:  v.MemberCharacterBaseMasterId,
		4:  v.StoryCharacterBaseMasterId,
		5:  v.ShopCharacterBaseMasterId,
		6:  v.HomeCostumeMasterId,
		8:  v.MemberCostumeMasterId,
		9:  v.StoryCostumeMasterId,
		10: v.ShopCostumeMasterId,
		11: v.IllustCharacterMasterId,
		12: v.DisplayAwakeningStatus,
		13: int32(v.HomeCharacterDisplayType),
		14: v.LoginBonusCharacterBaseMasterId,
		15: v.LoginBonusCostumeMasterId,
	})
}

func CurrencyToArray(v dto.Currency) []any {
	return []any{v.Id, v.Coin, v.FreeJewel, v.PaidJewel}
}

func RestrictionToArray(v dto.Restriction) []any {
	return []any{v.Id, v.MultiLiveRestrictionFinishedAt, v.ReadMultiLiveRestrictionDialog}
}

func UserBonusToArray(v dto.UserBonus) []any {
	return []any{v.Id, v.ExperienceBonus, v.LessonStarRankBonus}
}

func GameHintToArray(v dto.GameHint) []any {
	return []any{v.Id, int32(v.PageCategory), v.HasAlreadyRead}
}

func ViewedShopToArray(v dto.ViewedShop) []any {
	return []any{v.Id, v.ExchangeShopMasterId, v.LastViewedAt, int32(v.ViewedShopCategory)}
}

func CharacterBaseToArray(v dto.CharacterBase) []any {
	return []any{
		v.Id,
		v.CharacterBaseMasterId,
		v.StarRank,
		v.TotalStarPoint,
		v.CostumeMasterId,
		v.KeyMissionLevel,
		v.PortalCharacterId,
		v.PortalDisplayAwakeningStatus,
	}
}

func CharacterToArray(v dto.Character) []any {
	return []any{
		v.Id,
		v.CharacterMasterId,
		v.Level,
		v.CurrentExperience,
		v.TalentStage,
		v.AwakeningPhase,
		v.CharacterBaseId,
		v.SenseLevel,
		int32(v.ReadEpisodeOrder),
		int32(v.ReleasedEpisodeOrder),
		v.DisplayAwakeningStatus,
		v.SecondaryCharacterBaseId,
		v.SecondarySenseLevel,
		int32(v.SelectionType),
		v.IsFavorite,
	}
}

func PartyToArray(v dto.Party) []any {
	return []any{v.Id, v.Order, v.Name, v.LeaderPosition}
}

func PartySlotToArray(v dto.PartySlot) []any {
	return []any{
		v.Id,
		v.PartyId,
		v.Position,
		v.CharacterId,
		v.PosterId,
		v.AccessoryId,
		int32(v.BonusAbilityEnableFlags),
	}
}

func AlbumPresetToArray(v dto.AlbumPreset) []any {
	return []any{v.Id, v.Name, v.Order}
}

func FriendInvitationToArray(v dto.FriendInvitation) []any {
	return []any{v.Id, v.InvitationCode, v.HasInputOtherInvitationCode}
}

func FriendInvitationMissionToArray(v dto.FriendInvitationMission) []any {
	return []any{
		v.Id,
		v.FriendInvitationMissionMasterId,
		v.FriendInvitationMissionStageMasterId,
		v.CurrentCount,
		v.IsCleared,
		v.IsRewardReceived,
	}
}
