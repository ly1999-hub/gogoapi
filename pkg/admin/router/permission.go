package router
import(
	"github.com/labstack/echo/v4"
	"myapp/pkg/admin/handler"
)
func permission(e *echo.Echo){

	g:=e.Group("/permissions")
	h:=handler.Permission{}

	g.POST("",h.Migration)
}