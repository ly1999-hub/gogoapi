package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myapp/internal/constant"
	"myapp/internal/util/echoutil"
)

func JWT(authSecret string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(authSecret),
		Skipper: func(c echo.Context) bool {
			authHeader := echoutil.GetHeader(c, constant.HeaderAuthorization)

			// return true to skip middleware
			return authHeader == "" || authHeader == "Bearer" || authHeader == "Bearer "
		},
	})
}
