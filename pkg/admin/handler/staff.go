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

// LoginWithEmail godoc
// @tags Staff
// @summary LoginWithEmail
// @id login-with-email
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body apimodel.StaffLoginWithEmail true "Payload"
// @success 200 {object} responsemodel.StaffLogin
// @router /staffs/login-with-email [post]
func (h Staff) LoginWithEmail(c echo.Context) error {
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

// All godoc
// @tags Staff
// @summary All
// @id staff-all
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query apimodel.StaffAll false "Query"
// @success 200 {object} responsemodel.ResponseList
// @router /staff [get]
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

// Update godoc
// @tags Staff
// @summary Update
// @id update-staff
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body apimodel.StaffUpdate true "Payload"
// @param id path string true "staff id"
// @success 200 {object} responsemodel.ResponseUpdate
// @router /staffs/{id} [put]
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

// Create godoc
// @tags Staff
// @summary All
// @id staff-all
// @security ApikeyAuth
// @accept json
// @produce json
// @param payload body apimodel.StaffCreate true "Payload"
// @success 200 {object} responsemodel.ResponseCreate
// @router /staffs [post]
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

// ChangeStatus godoc
// @tags Staff
// @summary ChangeStatus
// @id change-status
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "staff id"
// @success 200 {object} responsemodel.ResponseUpdate
// @router /staff/{id}/change-status [patch]
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
