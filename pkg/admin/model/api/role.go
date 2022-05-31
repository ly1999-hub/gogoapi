package apimodel

import (
	"myapp/internal/mongodb"
	"myapp/internal/util/ptime"
	"myapp/internal/model"
)

type RoleAll struct{
	Page int64 `query:"page"`
	Litmit int64 `query:"litmit"`
	Keyword string `query:"keyword"`
}

//RoleCreate...
type RoleCreate struct{
	Name string
	Permission []string
}

//RoleUpdate..
type RoleUpdate struct {
	Name string
	Permission []string
}

func (a RoleCreate)ConvertToRaw()model.Role{
	now :=ptime.Now()
	return model.Role{
		ID: 			mongodb.NewObjectID(),
		Name:	 		a.Name,
		Code :			mongodb.NonAccentVietnamese(a.Name),
		Permissions: 	a.Permission,
		CreatedAt:		now,
		UpdatedAt:		now,
	}
}