package checkexist

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/constant"
	"myapp/internal/dao"
	"myapp/internal/mongodb"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
)

// User ...
func User(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx = echoutil.GetRequestContext(c)
			d   = dao.User{}
			id  = c.Param("id")
		)
		dbID, valid := mongodb.NewIDFromString(id)
		if !valid {
			return response.R404(c, nil, response.CommonNotFoundUser)
		}

		doc := d.FindByID(ctx, dbID)
		if doc.ID.IsZero() {
			return response.R404(c, nil, response.CommonNotFoundUser)
		}
		c.Set(constant.ContextKeyUser, doc)

		return next(c)
	}
}
