package responsemodel

import (
	"myapp/internal/model"
	"myapp/internal/util/ptime"
)

// User ...
type User struct {
	ID        string              `json:"_id"`
	Name      string              `json:"name"`
	Phone     string              `json:"phone"` // unique
	Email     string              `json:"email"` // unique
	Banned    bool                `json:"banned"`
	Avatar    string              `json:"avatar"`
	Statistic model.UserStatistic `json:"statistic"`
	CreatedAt *ptime.TimeResponse `json:"createdAt"`
	UpdatedAt *ptime.TimeResponse `json:"updatedAt"`
}

// UserInfo ...
type UserInfo struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Phone string `json:"phone"` // unique
	Email string `json:"email"` // unique
}

// UserDetail ...
type UserDetail struct {
	User User `json:"user"`
}
