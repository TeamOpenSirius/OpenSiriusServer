package api

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type UserProfileUpdateRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Bio      string `json:"bio"`
	Website  string `json:"website"`
	Avatar   string `json:"avatar"`
}

type PasswordUpdateRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NextPassword    string `json:"next_password" binding:"required"`
}

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role"`
	Gender   string `json:"gender"`
	Bio      string `json:"bio"`
	Website  string `json:"website"`
	Avatar   string `json:"avatar"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Gender   string `json:"gender"`
	Bio      string `json:"bio"`
	Website  string `json:"website"`
	Avatar   string `json:"avatar"`
}

type UserStateRequest struct {
	State string `json:"state" binding:"required"`
}

type UserRoleRequest struct {
	Role string `json:"role" binding:"required"`
}
