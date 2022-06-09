package middleware

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/constant"
	"myapp/internal/model"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
	"myapp/pkg/admin/service"
)

func CheckPermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := echoutil.GetRequestContext(c)
			staff := c.Get(constant.ContextKeyCurrentStaff).(model.Staff)
			if staff.ID.IsZero() {
				return response.R400(c, nil, "")
			}
			if staff.IsRoot {
				return next(c)
			}

			// Check Permission
			if staff.Role != nil {
				roleService := service.Role{}
				if !roleService.HasPermission(ctx, *staff.Role, permission) {
					return response.R403(c, nil, "")
				}
			}
			return next(c)
		}
	}
}
