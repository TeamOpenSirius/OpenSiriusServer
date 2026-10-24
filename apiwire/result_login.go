package apiwire

func LoginResult(invalidedStarPasses any, loginPassNotification any, isApproachingLoginPassInvalided bool, invalidedItemMasterIds any, approachingItemMasterIds any, storyEventPointExchangeResult any, invalidedBuffItemMasterIds any) []any {
	return []any{invalidedStarPasses, loginPassNotification, isApproachingLoginPassInvalided, invalidedItemMasterIds, approachingItemMasterIds, storyEventPointExchangeResult, invalidedBuffItemMasterIds}
}

func EmptyLoginResult() []any { return []any{[]any{}, nil, false, []any{}, []any{}, nil, []any{}} }
