package checkexist

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/internal/constant"
	"myapp/internal/dao"
	"myapp/internal/module/mongodb"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
)

// Role ...
func Role(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx = echoutil.GetRequestContext(c)
			id  = c.Param("id")

			d = dao.Role{}
		)
		dbID, valid := mongodb.NewIDFromString(id)
		if !valid {
			return response.R404(c, nil, response.CommonNotFound)
		}

		//find staff by OBJ_ID dbID
		doc := d.FindByID(ctx, dbID)
		if doc.ID.IsZero() {
			fmt.Println("hello1")
			return response.R404(c, nil, response.CommonNotFound)
		}
		c.Set(constant.ContextKeyRole, doc)

		return next(c)
	}
}
