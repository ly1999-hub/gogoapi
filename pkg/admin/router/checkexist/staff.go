package checkexist

import (
	"myapp/internal/constant"
	"myapp/internal/response"
	"myapp/internal/mongodb"
	"myapp/internal/dao"
	"myapp/internal/util/echoutil"
	"github.com/labstack/echo/v4"
)

func Staff(next echo.HandlerFunc)echo.HandlerFunc{
	return func(c echo.Context)error{
		var(
			ctx = echoutil.GetRequestContext(c)
			d =dao.Staff{}
			id =c.Param("id")
		)
		//NewIDFromString by id string
		dbID,valid :=mongodb.NewIDFromString(id)
		if !valid{
			return response.R404(c,nil,response.CommonNotFound)
		}

		//find staff by OBJ_ID dbID
		doc :=d.FindByID(ctx,dbID)
		if doc.ID.IsZero(){
			return response.R404(c,nil,response.CommonNotFound)
		}
		c.Set(constant.ContextKeyStaff,doc)

		return next(c)
	}
}

