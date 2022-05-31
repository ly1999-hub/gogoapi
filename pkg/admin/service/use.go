package service

import (
	"context"
	"errors"
	"sync"

	"myapp/internal/dao"
	"myapp/internal/model"
	"myapp/internal/response"
	"myapp/internal/util/ptime"
	"myapp/internal/util/query"
	responsemodel "myapp/pkg/admin/model/response"
	"go.mongodb.org/mongo-driver/bson"
)

// User ...
type User struct {
}

// All ...
func (s User) All(ctx context.Context, q query.Query) (r responsemodel.ResponseList) {
	var (
		wg   sync.WaitGroup
		d    = dao.User{}
		cond = bson.M{}
	)

	// Assign
	q.SetDefaultLimit()
	q.AssignKeyword(cond)
	q.AssignFromAtToAt(cond)
	q.AssignBanned(cond)

	r.Limit = q.Limit

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs := d.FindByCondition(ctx, cond, q.GetFindOptionsUsingPage())

		res := make([]responsemodel.User, 0)
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

func (s User) getResponse(doc model.User) responsemodel.User {
	return responsemodel.User{
		ID:        doc.ID.Hex(),
		Name:      doc.Name,
		Phone:     doc.Phone,
		Email:     doc.Email,
		Banned:    doc.Banned,
		Avatar:    doc.Avatar,
		Statistic: doc.Statistic,
		CreatedAt: ptime.TimeResponseInit(doc.CreatedAt),
		UpdatedAt: ptime.TimeResponseInit(doc.UpdatedAt),
	}
}

// Detail ...
func (s User) Detail(ctx context.Context, user model.User) (result responsemodel.UserDetail) {
	result.User = s.getResponse(user)
	return
}

// Banned ...
func (s User) Banned(ctx context.Context, user model.User) (result responsemodel.ResponseUpdate, err error) {
	var (
		d = dao.User{}
	)

	// Setup data
	payload := bson.M{
		"$set": bson.M{
			"banned": !user.Banned,
		},
	}

	// Update
	if err = d.UpdateByID(ctx, user.ID, payload); err != nil {
		err = errors.New(response.CommonErrorWhenHandler)
		return
	}

	result.ID = user.ID.Hex()
	return
}
