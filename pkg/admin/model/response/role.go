package responsemodel

import (
	"myapp/internal/util/ptime"
	"myapp/internal/constant"
)

type Role struct{
	ID string
	Name string
	Code string
	Permissions []constant.Permission
	CreatedAt *ptime.TimeResponse
	UpdatedAt *ptime.TimeResponse
}

//RoleShort...
type RoleShort struct{
	ID string `json:"_id"`
	Name string `json:"name"`
	IsAdmin bool `json:"isAdmin"`
}

//RoleDetail..
type RoleDetail struct{
	Role Role `json:"role"`
}


