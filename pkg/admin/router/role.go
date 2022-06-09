package router

import (
	"github.com/labstack/echo/v4"
	"myapp/pkg/admin/handler"
	"myapp/pkg/admin/router/checkexist"
	"myapp/pkg/admin/router/middleware"

	//"myapp/pkg/admin/router/middleware"
	"myapp/pkg/admin/router/validation"
)

func role(e *echo.Echo) {
	g := e.Group("/roles")
	h := handler.Role{}
	v := validation.Role{}

	g.GET("", h.All, middleware.CheckRootPermission, v.All)

	g.POST("", h.Create, middleware.CheckRootPermission, v.Create)

	g.DELETE("/:id", h.DeleteOne, middleware.CheckRootPermission, checkexist.Role)
}
