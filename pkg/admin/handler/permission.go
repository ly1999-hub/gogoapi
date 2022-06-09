package handler

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
	"myapp/pkg/admin/service"
)

type Permission struct{}

func (h Permission) Migration(c echo.Context) error {
	var (
		ctx = echoutil.GetRequestContext(c)
		s   = service.Permission{}
	)

	if err := s.Migration(ctx); err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, nil, "")
}
