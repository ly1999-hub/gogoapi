package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myapp/internal/constant"
	"myapp/internal/model"
	"myapp/internal/module/logger"
	"myapp/internal/module/mongodb"
)

type Role struct {
}

// InsetOne ...
func (d Role) InsetOne(ctx context.Context, doc model.Role) error {

	col := d.getCollection()

	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		logger.Error("insertone- dao Role", logger.LogData{
			"doc": doc,
			"err": err.Error(),
		})
	}

	return err
}

// UpdateOne...
func (d Role) UpdateOne(ctx context.Context, cond, payload interface{}) error {
	col := d.getCollection()
	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		logger.Error("updateone dao role", logger.LogData{
			"cond":    cond,
			"payload": payload,
			"err":     err.Error(),
		})
	}
	return err
}

//UpdateByID...
func (d Role) UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error {
	return d.UpdateOne(ctx, bson.M{"_id": id}, payload)
}

//FindOne..
func (d Role) FindOne(ctx context.Context, cond interface{}) (doc model.Role) {
	col := d.getCollection()

	err := col.FindOne(ctx, cond).Decode(&doc)
	if err != nil {
		logger.Error("FindOne Role dao", logger.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
	}
	return doc
}

//FindByID..
func (d Role) FindByID(ctx context.Context, id primitive.ObjectID) (doc model.Role) {
	return d.FindOne(ctx, bson.M{"_id": id})
}

//FindByCondition...

func (d Role) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.Role) {
	col := d.getCollection()

	cursor, err := col.Find(ctx, cond, opts...)

	if err != nil {
		logger.Error("FindByCondition Role dao", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return nil
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &docs); err != nil {
		logger.Error("dao.Role-FindByCondition", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return nil
	}
	return docs
}

//FindByIds..
func (d Role) FindByIDs(ctx context.Context, ids []primitive.ObjectID) []model.Role {
	cond := bson.M{"_id": bson.M{"$in": ids}}
	return d.FindByCondition(ctx, cond)
}

// CountByCondition ...
func (d Role) CountByCondition(ctx context.Context, cond interface{}) int64 {
	col := d.getCollection()
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("dao.Role-CountByCondition", logger.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
	}
	return total

}

//DeleteOne...
func (d Role) DeleteOne(ctx context.Context, id primitive.ObjectID) (result int64, err error) {
	var col = d.getCollection()

	results, err := col.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		logger.Error("DeleteOne-Dao Role", logger.LogData{
			"id":  id,
			"err": err.Error(),
		})
	}
	return results.DeletedCount, err
}

func (d Role) getCollection() *mongo.Collection {
	db := mongodb.GetDB()
	return db.Collection(constant.ColRole)
}
