package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myapp/internal/constant"
	"net/http"
)

func CORSConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentLength, echo.HeaderContentType,
			echo.HeaderAuthorization, constant.HeaderAcceptLanguage, constant.HeaderAppVersion,
			constant.HeaderDeviceID, constant.HeaderUserAgent, constant.HeaderModel,
			constant.HeaderManufacturer, constant.HeaderFCMToken,
		},
		AllowCredentials: false,
		MaxAge:           600,
	})
}
