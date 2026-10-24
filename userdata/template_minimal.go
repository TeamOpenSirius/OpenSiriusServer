package userdata

import (
	"context"
	"strings"
	"time"

	"ymst_server/apiwire"
)

const ProfileTemplateMinimal = "template_minimal"

func BuildTemplateMinimalInitialState(ctx context.Context, opts BuildInitialStateOptions) ([]any, error) {
	_ = ctx
	if opts.TemplateRegistry == nil || opts.TemplateRegistry.TemplateCount() == 0 {
		return BuildGeneratedInitialState(ctx, opts)
	}

	now := opts.Now
	if now.IsZero() {
		now = time.Now().UTC()
	}
	gen := opts.Generated
	if gen.PlayerRank <= 0 {
		gen = DefaultGeneratedUserOptions()
	}
	name := strings.TrimSpace(opts.DisplayName)
	if name == "" {
		name = "Player"
	}
	userKey := strings.TrimSpace(opts.UserKey)
	if userKey == "" {
		userKey = "default"
	}

	t := opts.TemplateRegistry
	payload := make([]any, 0, 64)
	add := func(unionID int64, data []any) { payload = append(payload, apiwire.DataObject(unionID, data)) }

	userData := cloneTemplateOrNil(t, UnionUser)
	if userData != nil {
		patchArray(userData, map[int]any{
			1:  gen.PlayerRank,
			3:  gen.CurrentStamina,
			5:  gen.PaidJewel,
			6:  gen.FreeJewel,
			7:  gen.Coin,
			12: now,
			15: gen.TutorialStatus,
			17: now,
		})
		add(UnionUser, userData)
	}

	profileData := cloneTemplateOrNil(t, UnionUserProfile)
	if profileData != nil {
		patchArray(profileData, map[int]any{1: name})
		add(UnionUserProfile, profileData)
	}

	prefData := cloneTemplateOrNil(t, UnionUserPreference)
	if prefData != nil {
		add(UnionUserPreference, prefData)
	}

	homeData := cloneTemplateOrNil(t, UnionHomeDisplayPreference)
	if homeData != nil {
		add(UnionHomeDisplayPreference, homeData)
	}

	currencyData := cloneTemplateOrNil(t, UnionCurrency)
	if currencyData != nil {
		patchArray(currencyData, map[int]any{
			1: gen.Coin,
			2: gen.FreeJewel,
			3: gen.PaidJewel,
		})
		add(UnionCurrency, currencyData)
	}

	viewedShopData := cloneTemplateOrNil(t, UnionViewedShop)
	if viewedShopData != nil {
		add(UnionViewedShop, viewedShopData)
	}

	for _, template := range t.Templates[UnionGameHint] {
		add(UnionGameHint, cloneArray(template))
	}

	userBonusData := cloneTemplateOrNil(t, UnionUserBonus)
	if userBonusData != nil {
		add(UnionUserBonus, userBonusData)
	}

	restrictionData := cloneTemplateOrNil(t, UnionRestriction)
	if restrictionData != nil {
		add(UnionRestriction, restrictionData)
	}

	albumData := cloneTemplateOrNil(t, UnionAlbumPreset)
	if albumData != nil {
		add(UnionAlbumPreset, albumData)
	}

	invitationData := cloneTemplateOrNil(t, UnionFriendInvitation)
	if invitationData != nil {
		add(UnionFriendInvitation, invitationData)
	}

	charBases, chars, party, slots := extractCoherentParty(t)
	for _, cb := range charBases {
		add(UnionCharacterBase, cb)
	}
	for _, c := range chars {
		add(UnionCharacter, c)
	}
	if party != nil {
		add(UnionParty, party)
	}
	for _, slot := range slots {
		add(UnionPartySlot, slot)
	}

	_ = userKey
	return payload, nil
}

func extractCoherentParty(t *UserDataTemplateRegistry) (charBases [][]any, chars [][]any, party []any, slots [][]any) {
	partyTemplates := t.Templates[UnionParty]
	if len(partyTemplates) == 0 {
		return nil, nil, nil, nil
	}
	firstParty := cloneArray(partyTemplates[0])
	party = firstParty

	partyID := fieldAt(firstParty, 0)

	var pickedSlots [][]any
	for _, slotTemplate := range t.Templates[UnionPartySlot] {
		slot := cloneArray(slotTemplate)
		if fieldAt(slot, 1) == partyID {
			pickedSlots = append(pickedSlots, slot)
			if len(pickedSlots) >= 5 {
				break
			}
		}
	}

	neededCharacters := setOf()
	neededCharacterBases := setOf()
	for _, slot := range pickedSlots {
		charID := fieldAt(slot, 3)
		neededCharacters[charID] = true
	}

	for _, charTemplate := range t.Templates[UnionCharacter] {
		charID := fieldAt(charTemplate, 0)
		if neededCharacters[charID] {
			char := cloneArray(charTemplate)
			chars = append(chars, char)
			baseID := fieldAt(char, 6)
			neededCharacterBases[baseID] = true
		}
	}

	for _, baseTemplate := range t.Templates[UnionCharacterBase] {
		baseID := fieldAt(baseTemplate, 0)
		if neededCharacterBases[baseID] {
			charBases = append(charBases, cloneArray(baseTemplate))
		}
	}

	return charBases, chars, party, pickedSlots
}

func cloneTemplateOrNil(t *UserDataTemplateRegistry, unionID int64) []any {
	templates, ok := t.Templates[unionID]
	if !ok || len(templates) == 0 {
		return nil
	}
	return cloneArray(templates[0])
}

func patchArray(data []any, patches map[int]any) {
	for idx, val := range patches {
		if idx >= 0 && idx < len(data) {
			data[idx] = val
		}
	}
}

func fieldAt(data []any, idx int) int64 {
	if idx < 0 || idx >= len(data) {
		return 0
	}
	return Int64Value(data[idx])
}
