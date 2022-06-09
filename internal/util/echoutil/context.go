package echoutil

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"
	"myapp/internal/constant"
	"myapp/internal/model"
)

// GetRequestContext ...
func GetRequestContext(c echo.Context) context.Context {
	return c.Request().Context()
}

// GetPayload ...
func GetPayload(c echo.Context) interface{} {
	return c.Get(constant.ContextKeyPayload)
}

// SetPayload ...
func SetPayload(c echo.Context, value interface{}) {
	c.Set(constant.ContextKeyPayload, value)
}

// GetQuery ...
func GetQuery(c echo.Context) interface{} {
	return c.Get(constant.ContextKeyQuery)
}

// SetQuery ...
func SetQuery(c echo.Context, value interface{}) {
	c.Set(constant.ContextKeyQuery, value)
}

// GetCurrentStaff ...
func GetCurrentStaff(c echo.Context) model.Staff {
	return c.Get(constant.ContextKeyCurrentStaff).(model.Staff)
}

// GetAuthToken ...
func GetAuthToken(c echo.Context) string {
	return strings.Split(GetHeader(c, constant.HeaderAuthorization), " ")[1]
}
