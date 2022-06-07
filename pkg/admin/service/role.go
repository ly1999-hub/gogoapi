package service

import (
	"context"
	"myapp/internal/constant"
	"myapp/internal/dao"
	"myapp/internal/model"
	"myapp/internal/mongodb"
	"myapp/internal/util/parray"
	"myapp/internal/util/ptime"
	"myapp/internal/util/query"
	apimodel "myapp/pkg/admin/model/api"
	responsemodel "myapp/pkg/admin/model/response"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct{}

// Create ...
func (s Role) Create(ctx context.Context, payload apimodel.RoleCreate) (result string, err error) {
	var (
		d = dao.Role{}
	)
	//new role data from payload
	doc := payload.ConvertToRaw()
	//create role
	if err = d.InsetOne(ctx, doc); err != nil {
		return
	}

	result = doc.ID.Hex()
	return
}

// DeleteOne ...
func (s Role) DeleteOne(ctx context.Context, roleID primitive.ObjectID) (result int64, err error) {
	var (
		d = dao.Role{}
	)

	results, errs := d.DeleteOne(ctx, roleID)
	if errs != nil {
		return
	}

	return results, nil
}

// Update ...
func (s Role) Update(ctx context.Context, roleID primitive.ObjectID, payload apimodel.RoleUpdate) (result string, err error) {
	var (
		d = dao.Role{}
	)
	//Setup Data
	updateData := bson.M{
		"$set": bson.M{
			"name":        payload.Name,
			"code":        mongodb.NonAccentVietnamese(payload.Name),
			"permissions": payload.Permission,
		},
	}

	//create role

	if err = d.UpdateByID(ctx, roleID, updateData); err != nil {
		return
	}
	result = roleID.Hex()
	return
}

// All ..
func (s Role) All(ctx context.Context, query query.Query) (r responsemodel.ResponseList) {
	var (
		wg   sync.WaitGroup
		d    = dao.Role{}
		cond = bson.M{}
	)

	//Assign
	query.SetDefaultLimit()
	query.AssignKeyword(cond)
	r.Limit = query.Limit

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs := d.FindByCondition(ctx, cond, query.GetFindOptionsUsingPage())
		res := make([]responsemodel.Role, 0)
		for _, value := range docs {
			res = append(res, s.getResponse(value))
		}
		r.List = res
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		r.Total = d.CountByCondition(ctx, cond)
	}()
	wg.Wait()

	return r
}

// getResponse

func (s Role) getResponse(doc model.Role) responsemodel.Role {
	//get permission
	PermissionRes := make([]constant.Permission, 0)
	for _, p := range doc.Permissions {
		for _, pRes := range constant.Permissions {
			if p == pRes.Code {
				PermissionRes = append(PermissionRes, pRes)
			}
		}
	}
	return responsemodel.Role{
		ID:          doc.ID.Hex(),
		Name:        doc.Name,
		Code:        doc.Code,
		Permissions: PermissionRes,
		CreatedAt:   ptime.TimeResponseInit(doc.CreatedAt),
		UpdatedAt:   ptime.TimeResponseInit(doc.UpdatedAt),
	}
}

// Detail ...
func (s Role) Detail(ctx context.Context, role model.Role) (result responsemodel.RoleDetail) {
	result.Role = s.getResponse(role)
	return
}

// GetRoleByUser ...
func (s Role) GetRoleByUser(ctx context.Context, role model.Role) (result responsemodel.RoleDetail) {
	result.Role = s.getResponse(role)
	return
}

// HasPermission ...
func (s Role) HasPermission(ctx context.Context, roleID primitive.ObjectID, permission string) bool {
	var d = dao.Role{}
	role := d.FindByID(ctx, roleID)
	return parray.Contains(role.Permissions, permission)
}
