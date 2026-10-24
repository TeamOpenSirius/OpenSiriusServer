package api

import "ymst_server/core/database"

func UserLiteFromModel(user *database.User) UserLite {
	if user == nil {
		return UserLite{}
	}
	return UserLite{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
}

func UserSummaryFromModel(user *database.User) UserSummary {
	if user == nil {
		return UserSummary{}
	}
	return UserSummary{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Gender:   GenderName(user.Gender),
		Bio:      user.Bio,
		Website:  user.Website,
		Role:     UserRoleName(user.Role),
	}
}

func UserProfileFromModel(user *database.User) UserProfile {
	if user == nil {
		return UserProfile{}
	}
	return UserProfile{
		ID:            user.ID,
		State:         UserStateName(user.State),
		Role:          UserRoleName(user.Role),
		Username:      user.Username,
		Nickname:      user.Nickname,
		Avatar:        user.Avatar,
		Gender:        GenderName(user.Gender),
		Email:         user.Email,
		Bio:           user.Bio,
		Website:       user.Website,
		RegisterDate:  user.RegisterDate,
		LastLogin:     user.LastLogin,
		EmailVerified: user.EmailVerified,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
