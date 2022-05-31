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

// GetAppVersion ...
func GetAppVersion(c echo.Context) string {
	return strings.ToLower(GetHeader(c, constant.HeaderAppVersion))
}

// GetDeviceID ...
func GetDeviceID(c echo.Context) string {
	return strings.ToLower(GetHeader(c, constant.HeaderDeviceID))
}

// GetUserAgent ...
func GetUserAgent(c echo.Context) string {
	return strings.ToLower(GetHeader(c, constant.HeaderUserAgent))
}

// GetFCMToken ...
func GetFCMToken(c echo.Context) string {
	return strings.ToLower(GetHeader(c, constant.HeaderFCMToken))
}

// GetModel ...
func GetModel(c echo.Context) string {
	return strings.ToLower(GetHeader(c, constant.HeaderModel))
}

// GetManufacturer ...
func GetManufacturer(c echo.Context) string {
	return strings.ToLower(GetHeader(c, constant.HeaderManufacturer))
}

// GetIP ...
func GetIP(c echo.Context) string {
	return strings.TrimSpace(c.RealIP())
}

// GetLanguage ...
func GetLanguage(c echo.Context) string {
	lang := GetHeader(c, constant.HeaderAcceptLanguage)
	if lang != constant.LangEn {
		lang = constant.LangVi
	}
	return lang
}

// GetCurrentUser ...
func GetCurrentUser(c echo.Context) model.User {
	return c.Get(constant.ContextKeyCurrentUser).(model.User)
}

// GetCurrentStaff ...
/*func GetCurrentStaff(c echo.Context) model.Staff {
	return c.Get(constant.ContextKeyCurrentStaff).(model.Staff)
}
*/

// GetAuthToken ...
func GetAuthToken(c echo.Context) string {
	return strings.Split(GetHeader(c, constant.HeaderAuthorization), " ")[1]
}
