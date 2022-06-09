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
	"myapp/pkg/admin/service"
)

type Role struct{}

// All godoc
// @tags Role
// @summary ALl
// @id role-all
// @security ApikeyAuth
// @accept json
// @produce json
// @param payload query apimodel.RoleAll false "Query"
// @success 200 {object} responsemodel.ResponseList
// @router /roles [get]
func (h Role) All(c echo.Context) error {
	var (
		queryParams = echoutil.GetQuery(c).(apimodel.RoleAll)
		s           = service.Role{}
		ctx         = echoutil.GetRequestContext(c)
		q           = query.Query{
			Page:     queryParams.Page,
			Limit:    queryParams.Limit,
			Keyword:  queryParams.Keyword,
			SortBSON: bson.M{"createdAt": -1},
		}
	)

	res := s.All(ctx, q)
	return response.R200(c, res, "")
}

// Create godoc
// @tags Role
// @summary Create
// @id role-create
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body apimodel.RoleCreate true "Payload"
// @success 200 {object} responsemodel.ResponseCreate
// @router /roles [post]
func (h Role) Create(c echo.Context) error {
	var (
		ctx     = echoutil.GetRequestContext(c)
		payload = echoutil.GetPayload(c).(apimodel.RoleCreate)
		s       = service.Role{}
	)

	res, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

// Update godoc
// @tags Role
// @summary Update
// @id role-update
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "role id"
// @param payload body apimodel.RoleUpdate true "Payload"
// @success 200 {object} responsemodel.ResponseUpdate
// @router /roles/{id} [put]
func (h Role) Update(c echo.Context) error {
	var (
		ctx     = echoutil.GetRequestContext(c)
		payload = echoutil.GetPayload(c).(apimodel.RoleUpdate)
		s       = service.Role{}
		role    = c.Get(constant.ContextKeyRole).(model.Role)
	)
	res, err := s.Update(ctx, role.ID, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

// Detail godoc
// @tags Role
// @summary Detail
// @id role-detail
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "role id"
// @success 200 {object} responsemodel.RoleDetail
// @router /roles/{id} [get]
func (h Role) Detail(c echo.Context) error {
	var (
		ctx  = echoutil.GetRequestContext(c)
		s    = service.Role{}
		role = c.Get(constant.ContextKeyRole).(model.Role)
	)
	res := s.Detail(ctx, role)

	return response.R200(c, res, "")

}

// DeleteOne godoc
// @tags Role
// @summary DeleteOne
// @id role-delete-one
// @security ApiKeyAuth
// @accept json
// @product json
// @param id path string true "role id"
// @success 200
// @router /roles/{id} [delete]
func (h Role) DeleteOne(c echo.Context) error {
	var (
		ctx  = echoutil.GetRequestContext(c)
		s    = service.Role{}
		role = c.Get("roleDelete").(model.Role)
	)

	result, err := s.DeleteOne(ctx, role.ID)
	if err != nil {
		return err
	}

	return response.R200(c, result, "")
}
