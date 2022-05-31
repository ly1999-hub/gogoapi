package router

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/middlewares"
	"myapp/pkg/admin/handler"
	"myapp/pkg/admin/router/checkexist"
	"myapp/pkg/admin/router/middleware"
	"myapp/pkg/admin/router/validation"
)

func Staff(e *echo.Echo) {
	g := e.Group("/staffs")
	e.Use(middlewares.JWT("AUTH_SECRET"))
	h := handler.Staff{}
	v := validation.Staff{}

	g.GET("", h.All, middleware.RequireLogin, middleware.CheckRootPermission, v.All)

	g.POST("", h.Create, middleware.RequireLogin, middleware.CheckRootPermission, v.Create)

	g.PUT("/:id", h.Update, middleware.RequireLogin, middleware.CheckRootPermission, v.Update, checkexist.Staff)

	g.POST("/login", h.Login, v.LoginWithEmail)

	g.PATCH("/:id/changeStatus", h.ChangeStatus, middleware.RequireLogin, middleware.CheckRootPermission, checkexist.Staff)

}
