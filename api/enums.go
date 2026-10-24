package api

import "ymst_server/core/database"

func UserStateName(state int) string {
	switch state {
	case database.UserStateRegistered:
		return "registered"
	case database.UserStateVerified:
		return "verified"
	case database.UserStateBanned:
		return "banned"
	case database.UserStateDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}

func UserRoleName(role int) string {
	switch role {
	case database.UserRoleUser:
		return "user"
	case database.UserRoleAdmin:
		return "admin"
	case database.UserRoleSuperAdmin:
		return "super_admin"
	default:
		return "unknown"
	}
}

func GenderName(gender int) string {
	switch gender {
	case database.GenderMale:
		return "male"
	case database.GenderFemale:
		return "female"
	case database.GenderOther:
		return "other"
	default:
		return "unknown"
	}
}
