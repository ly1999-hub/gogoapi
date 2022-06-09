package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myapp/internal/constant"
	"myapp/internal/dao"
	"myapp/internal/module/mongodb"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
)

// RequireLogin ...
func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		staffID := getStaffIDFromContext(c)
		if staffID.IsZero() {
			return response.R401(c, nil, "")
		}

		// Check staff status
		staff := dao.Staff{}.FindByID(echoutil.GetRequestContext(c), staffID)
		if staff.ID.IsZero() || !staff.Active {
			fmt.Println("hello")
			return response.R401(c, nil, "")
		}

		c.Set(constant.ContextKeyCurrentStaff, staff)
		return next(c)
	}
}

func getStaffIDFromContext(c echo.Context) (id primitive.ObjectID) {
	token := c.Get("user")
	if token == nil {
		return
	}

	data, ok := token.(*jwt.Token)
	if !ok {
		return
	}

	m, ok := data.Claims.(jwt.MapClaims)
	if ok && data.Valid && m["_id"] != "" {
		idString := m["_id"].(string)
		id, _ = mongodb.NewIDFromString(idString)
	}

	return
}

// CheckRootPermission ...
func CheckRootPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		staffID := getStaffIDFromContext(c)
		if staffID.IsZero() {
			return response.R401(c, nil, "")
		}

		//Check staff status

		staff := dao.Staff{}.FindByID(echoutil.GetRequestContext(c), staffID)
		if staff.ID.IsZero() || !staff.IsRoot {
			return response.R403(c, nil, "")
		}
		c.Set(constant.ContextKeyCurrentStaff, staff)

		return next(c)
	}
}
