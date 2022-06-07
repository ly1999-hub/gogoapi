package server

import (
	"github.com/labstack/echo/v4"
	"myapp/pkg/admin/config"
	"myapp/pkg/admin/initialize"
	"myapp/pkg/admin/router"
	"myapp/pkg/admin/service"
)

func Init(e *echo.Echo) {
	// Init .env
	config.Init()

	// InitMongoDB connect DB
	initialize.InitMongoDB()

	// InitRootStaff check root staff
	service.Staff{}.InitRootStaff()

	// Init router
	router.Init(e)
}
