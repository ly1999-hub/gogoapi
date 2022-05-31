package validation

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
	apimodel "myapp/pkg/admin/model/api"
)

//Role...
type Role struct{}

//All...
func (Role) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload apimodel.RoleAll
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, err.Error())
		}
		echoutil.SetQuery(c, payload)
		return next(c)

	}
}

//Create...
func (Role) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload apimodel.RoleCreate

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, err.Error())
		}
		echoutil.SetPayload(c, payload)
		return next(c)
	}
}

//GetPermission...
func (Role) GetPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload apimodel.PermissionAll

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, err.Error())
		}
		echoutil.GetQuery(c)
		return next(c)
	}
}
