package checkexist

import (
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"myapp/internal/util/echoutil"
	"myapp/internal/dao"
	"github.com/labstack/echo/v4"
)

func Login(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context)error{
		var (
			ctx =echoutil.GetRequestContext(c)
			email =c.FormValue("email")
			password =c.FormValue("password")
			d =dao.Staff{}
		)

		if email ==""{
			return c.JSON(http.StatusNotFound,"no email")
		}
		if password ==""{
			return c.JSON(http.StatusNotFound,"no password")
		}
		
		filter := bson.M{"email":email,"hashedPassword":password}
		doc := d.FindOne(ctx,filter)
		if doc.ID.IsZero(){
			return c.JSON(http.StatusNotFound,"err password")
		}

		c.Set("staffLogined",doc)
		return next(c)
	}
}