package apiwire

import "time"

func TimedConfirmationCode(code string, remainingSeconds int64) []any {
	return []any{code, remainingSeconds}
}

func AccountDeletionResult(success bool, errorType any) []any { return []any{success, errorType} }

func ViewShopResult(convertedThings []any) []any { return []any{convertedThings} }

func InboxReceiveResult(receivedThings []any, hasNotReceiveThings bool) []any {
	return []any{receivedThings, hasNotReceiveThings}
}

func FriendListResult(results []any, currentFriendCount int64) []any {
	return []any{results, currentFriendCount}
}

func FriendAcceptResult(resultStatus int64) []any { return []any{resultStatus} }

func FriendRequestResult(resultStatus int64) []any { return []any{resultStatus} }

func CreateCircleResult(resultStatus int64, circleHashedID string) []any {
	return []any{resultStatus, circleHashedID}
}

func CircleResult(resultStatus int64) []any { return []any{resultStatus} }

func SupportCompanyLevelLimitStatus(currentCoinQuantity int64, details any) []any {
	return []any{currentCoinQuantity, details}
}

func SupportCompanyLevelStatus(level int64, currentSupportPoint int64, lastLevelUppedAt any, levelLimit int64, levelLimitStatus any) []any {
	return []any{level, currentSupportPoint, lastLevelUppedAt, levelLimit, levelLimitStatus}
}

func SupportCompanyInformation(sirius, eden, gingaza, denki []any) []any {
	return []any{sirius, eden, gingaza, denki}
}

func NotificationResult(id int64, bannerPath string, title string, postingAt time.Time, lastUpdatedAt time.Time, tabCategory int64, category int64, isConfirmation bool, order int64) []any {
	return []any{id, bannerPath, title, postingAt, lastUpdatedAt, tabCategory, category, isConfirmation, order}
}

func NotificationContentResult(id int64, body string, title string, postingAt time.Time, lastUpdatedAt time.Time, tabCategory int64, category int64, bannerPath string) []any {
	return []any{id, body, title, postingAt, lastUpdatedAt, tabCategory, category, bannerPath}
}
