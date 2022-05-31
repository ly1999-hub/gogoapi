package echoutil

import (
	"github.com/labstack/echo/v4"
)

// GetHeader ...
func GetHeader(c echo.Context, key string) string {
	return c.Request().Header.Get(key)
}
