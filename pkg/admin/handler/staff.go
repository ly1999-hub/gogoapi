package handler

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"myapp/internal/constant"
	"myapp/internal/model"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
	"myapp/internal/util/query"
	apimodel "myapp/pkg/admin/model/api"
	responsemodel "myapp/pkg/admin/model/response"
	"myapp/pkg/admin/service"
)

type Staff struct{}

func (h Staff) Login(c echo.Context) error {
	var (
		s       = service.Staff{}
		ctx     = echoutil.GetRequestContext(c)
		payload = echoutil.GetPayload(c).(apimodel.StaffLoginWithEmail)
	)
	res, err := s.LoginWithEmail(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res, "")
}

func (h Staff) All(c echo.Context) error {
	var (
		ctx         = echoutil.GetRequestContext(c)
		queryParams = echoutil.GetQuery(c).(apimodel.StaffAll)
		s           = service.Staff{}
		q           = query.Query{
			Page:     queryParams.Page,
			Limit:    queryParams.Limit,
			Active:   queryParams.Active,
			Keyword:  queryParams.Keyword,
			SortBSON: bson.M{"createdAt": -1},
		}
	)
	res := s.All(ctx, q)
	return response.R200(c, res, "")
}

//Update...
func (h Staff) Update(c echo.Context) error {
	var (
		ctx     = echoutil.GetRequestContext(c)
		s       = service.Staff{}
		payload = echoutil.GetPayload(c).(apimodel.StaffUpdate)
		staff   = c.Get(constant.ContextKeyStaff).(model.Staff)
	)
	res, err := s.Update(ctx, staff, payload)

	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res, "")
}

func (h Staff) Create(c echo.Context) error {
	var (
		ctx     = echoutil.GetRequestContext(c)
		payload = echoutil.GetPayload(c).(apimodel.StaffCreate)
		s       = service.Staff{}
	)
	res, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, "")
	}

	return response.R200(c, res, "")
}

func (h Staff) ChangeStatus(c echo.Context) error {
	var (
		s     = service.Staff{}
		ctx   = echoutil.GetRequestContext(c)
		_     = responsemodel.ResponseList{}
		staff = c.Get(constant.ContextKeyStaff).(model.Staff)
	)

	res, err := s.ChangeStatus(ctx, staff)
	if err != nil {
		return response.R400(c, err.Error(), "")
	}

	return response.R200(c, res, "")
}
