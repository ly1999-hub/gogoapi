package checkexist

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/internal/dao"
	"myapp/internal/mongodb"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
)

func DeleteOne(next echo.HandlerFunc) echo.HandlerFunc {
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
		fmt.Println("hello2")
		c.Set("roleDelete", doc)

		return next(c)
	}
}
