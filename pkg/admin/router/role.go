package router

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/middlewares"
	"myapp/pkg/admin/handler"
	"myapp/pkg/admin/router/checkexist"
	"myapp/pkg/admin/router/middleware"

	//"myapp/pkg/admin/router/middleware"
	"myapp/pkg/admin/router/validation"
)

func Role(e *echo.Echo) {
	g := e.Group("/roles")
	e.Use(middlewares.JWT("AUTH_SECRET"))
	h := handler.Role{}
	v := validation.Role{}

	g.GET("", h.All, middleware.CheckRootPermission, v.All)

	g.POST("", h.Create, middleware.CheckRootPermission, v.Create)

	g.DELETE("", h.DeleteOne, middleware.CheckRootPermission, checkexist.DeleteOne)
}
