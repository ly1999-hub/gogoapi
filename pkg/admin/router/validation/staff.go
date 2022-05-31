package validation

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
	"myapp/pkg/admin/model/api"
)

type Staff struct{}

//Create...
func (Staff) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload apimodel.StaffCreate

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}
		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}
		echoutil.SetPayload(c, payload)
		return next(c)
	}
}

//Update...
func (Staff) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload apimodel.StaffUpdate

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}
		echoutil.SetPayload(c, payload)
		return next(c)
	}
}

//All...
func (Staff) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload apimodel.StaffAll

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		echoutil.SetQuery(c, payload)
		return next(c)
	}
}

//ChangPassword...
func (Staff) ChangePassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload apimodel.StaffChangePassword

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echoutil.SetPayload(c, payload)
		return next(c)
	}
}

//LoginWithEmail...
func (Staff) LoginWithEmail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload apimodel.StaffLoginWithEmail

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echoutil.SetPayload(c, payload)
		return next(c)
	}
}
