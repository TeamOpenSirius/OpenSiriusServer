package api

import "time"

type UserLite struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar,omitempty"`
}

type UserSummary struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar,omitempty"`
	Gender   string `json:"gender"`
	Bio      string `json:"bio"`
	Website  string `json:"website"`
	Role     string `json:"role"`
}

type UserProfile struct {
	ID            uint       `json:"id"`
	State         string     `json:"state"`
	Role          string     `json:"role"`
	Username      string     `json:"username"`
	Nickname      string     `json:"nickname"`
	Avatar        string     `json:"avatar,omitempty"`
	Gender        string     `json:"gender"`
	Email         string     `json:"email"`
	Bio           string     `json:"bio"`
	Website       string     `json:"website"`
	RegisterDate  time.Time  `json:"register_date"`
	LastLogin     *time.Time `json:"last_login"`
	EmailVerified bool       `json:"email_verified"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type TrendPoint struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}
