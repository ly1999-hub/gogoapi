package router

import (
	"github.com/labstack/echo/v4"
	"myapp/pkg/admin/handler"
	"myapp/pkg/admin/router/checkexist"
	"myapp/pkg/admin/router/validation"
)

func user(e *echo.Echo) {
	g := e.Group("/users")

	h := handler.User{}
	v := validation.User{}

	g.GET("", h.All, v.All)

	g.GET("/:id", h.Detail, checkexist.User)

	g.PATCH("/banned/:id", h.Banned, checkexist.User)
}
