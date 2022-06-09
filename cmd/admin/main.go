package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"myapp/docs/admin"
	"myapp/pkg/admin/config"
	"myapp/pkg/admin/server"
	"os"
)

// @title Admin API
// @version 1.0
// @description API for admin management.
// @basePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	// init server
	server.Init(e)

	// swagger
	admin.SwaggerInfo.Host = os.Getenv("DOMAIN_ADMIN")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(config.GetENV().Port))
}
