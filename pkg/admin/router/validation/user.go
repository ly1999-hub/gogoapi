package validation


import (
	"github.com/labstack/echo/v4"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
	apimodel "myapp/pkg/admin/model/api"
)

// User ...
type User struct{}

// All ...
func (User) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload apimodel.UserAll

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		echoutil.SetQuery(c, payload)
		return next(c)
	}
}
