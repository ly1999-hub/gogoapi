package handler

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"myapp/internal/constant"
	"myapp/internal/model"
	"myapp/internal/response"
	"myapp/internal/util/echoutil"
	"myapp/internal/util/ptime"
	"myapp/internal/util/query"
	apimodel "myapp/pkg/admin/model/api"
	"myapp/pkg/admin/service"
)

// User ...
type User struct{}

// All godoc
// @tags User
// @summary All
// @id user-all
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query apimodel.UserAll false "Query"
// @success 200 {object} responsemodel.User
// @router /users [get]
func (h User) All(c echo.Context) error {
	var (
		queryParams = echoutil.GetQuery(c).(apimodel.UserAll)
		s           = service.User{}
		ctx         = echoutil.GetRequestContext(c)
		q           = query.Query{
			Page:     queryParams.Page,
			Limit:    queryParams.Limit,
			Keyword:  queryParams.Keyword,
			Banned:   queryParams.Banned,
			FromAt:   ptime.ParseISODate(queryParams.FromAt),
			ToAt:     ptime.ParseISODate(queryParams.ToAt),
			SortBSON: bson.M{"createdAt": -1},
		}
	)

	res := s.All(ctx, q)

	return response.R200(c, res, "")
}

// Detail godoc
// @tags User
// @summary Detail
// @id user-detail
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "user id"
// @success 200 {object} responsemodel.UserDetail
// @router /users/{id} [get]
func (h User) Detail(c echo.Context) error {
	var (
		ctx  = echoutil.GetRequestContext(c)
		s    = service.User{}
		user = c.Get(constant.ContextKeyUser).(model.User)
	)

	res := s.Detail(ctx, user)

	return response.R200(c, res, "")

}

// Banned godoc
// @tags User
// @summary Banned
// @id user-banned
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "user id"
// @success 200 {object} responsemodel.ResponseUpdate
// @router /users/{id}/banned [patch]
func (h User) Banned(c echo.Context) error {
	var (
		ctx  = echoutil.GetRequestContext(c)
		s    = service.User{}
		user = c.Get(constant.ContextKeyUser).(model.User)
	)
	res, err := s.Banned(ctx, user)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res, "")
}
