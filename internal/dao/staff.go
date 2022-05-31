package dao

import (
	"context"

	"myapp/internal/constant"
	"myapp/internal/logger"
	"myapp/internal/model"
	"myapp/internal/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Staff ...
type Staff struct{}

// InsertOne ...
func (d Staff) InsertOne(ctx context.Context, doc model.Staff) error {
	col := d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		logger.Error("dao.Staff - InsertOne", logger.LogData{
			"doc": doc,
			"err": err.Error(),
		})
	}
	return err
}

// UpdateOne ...
func (d Staff) UpdateOne(ctx context.Context, cond, payload interface{}) error {
	col := d.getCollection()
	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		logger.Error("dao.Staff - UpdateOne", logger.LogData{
			"cond":    cond,
			"payload": payload,
			"err":     err.Error(),
		})
	}
	return err
}

// UpdateByID ...
func (d Staff) UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error {
	return d.UpdateOne(ctx, bson.M{"_id": id}, payload)
}

// FindByID ...
func (d Staff) FindByID(ctx context.Context, id primitive.ObjectID) (doc model.Staff) {
	return d.FindOne(ctx, bson.M{"_id": id})
}

// FindOne ...
func (d Staff) FindOne(ctx context.Context, cond interface{}) (doc model.Staff) {
	col := d.getCollection()
	if err := col.FindOne(ctx, cond).Decode(&doc); err != nil {
		logger.Error("dao.Staff - FindOne", logger.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
	}
	return doc
}

// FindByCondition ...
func (d Staff) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.Staff) {
	col := d.getCollection()
	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		logger.Error("dao.Staff - FindByCondition", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return nil
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		logger.Error("dao.Staff - FindByCondition - decode", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return nil
	}
	return docs
}

// FindByIDs ...
func (d Staff) FindByIDs(ctx context.Context, ids []primitive.ObjectID) []model.Staff {
	cond := bson.M{"_id": bson.M{"$in": ids}}
	return d.FindByCondition(ctx, cond)
}

// CountByCondition ...
func (d Staff) CountByCondition(ctx context.Context, cond interface{}) int64 {
	col := d.getCollection()
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("dao.Staff - CountByCondition", logger.LogData{
			"err":  err.Error(),
			"cond": cond,
		})
	}
	return total
}

func (d Staff) getCollection() *mongo.Collection {
	db := mongodb.GetDB()
	return db.Collection(constant.ColStaff)
}
