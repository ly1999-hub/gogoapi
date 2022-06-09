package apimodel

import (
	"myapp/internal/model"
	"myapp/internal/util/ptime"
	mongodb2 "myapp/module/mongodb"
)

type RoleAll struct {
	Page    int64  `query:"page"`
	Limit   int64  `query:"limit"`
	Keyword string `query:"keyword"`
}

// RoleCreate ...
type RoleCreate struct {
	Name       string
	Permission []string
}

// RoleUpdate ...
type RoleUpdate struct {
	Name       string
	Permission []string
}

// ConvertToRaw ...
func (a RoleCreate) ConvertToRaw() model.Role {
	now := ptime.Now()
	return model.Role{
		ID:          mongodb2.NewObjectID(),
		Name:        a.Name,
		Code:        mongodb2.NonAccentVietnamese(a.Name),
		Permissions: a.Permission,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
