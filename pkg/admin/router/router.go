package router

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/middlewares"
	"myapp/pkg/admin/config"
)

func Init(e *echo.Echo) {
	e.Use(middlewares.CORSConfig())
	e.Use(middlewares.JWT(config.GetENV().AuthSecret))

	Role(e)
	User(e)
	Staff(e)
}