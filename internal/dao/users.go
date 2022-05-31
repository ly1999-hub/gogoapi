package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myapp/internal/mongodb"
	"myapp/internal/logger"
	"context"
	"myapp/internal/model"
	"myapp/internal/constant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type User struct{}

//InsetOne..
func (d User) InsetOne(ctx context.Context,doc model.User)error{
	col :=d.getCollection()
	_,err :=col.InsertOne(ctx,doc)

	if err!=nil {
		logger.Error("dao.User - InsertOne",logger.LogData{
			"doc":doc,
			"err":err.Error(),
		})
	}
	return err
}

//UpdateOne..
func (d User)UpdateOne(ctx context.Context,cond,payload interface{})error{
	col :=d.getCollection()
	_,err :=col.UpdateOne(ctx,cond,payload)
	if err !=nil {
		logger.Error("dao.User - UpdateOne", logger.LogData{
			"cond":    cond,
			"payload": payload,
			"err":     err.Error(),
		})
	}
	return err
}

//UpdateByID...
func (d User)UpdateByID(ctx context.Context,id primitive.ObjectID,payload interface{})error{
	return d.UpdateOne(ctx,bson.M{"_id":id},payload)
}

//FindByID..
func (d User)FindByID(ctx context.Context,id primitive.ObjectID)(doc model.User){
	return d.FindOne(ctx,bson.M{"_id":id})
}

//FindOne..
func (d User)FindOne(ctx context.Context,cond interface{})(doc model.User){
	col :=d.getCollection()
	if err :=col.FindOne(ctx,cond).Decode(&doc);err!=nil{
		logger.Error("dao.User - FindOne", logger.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
	}
	return doc
}

// FindByCondition ...
func (d User) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.User) {
	col := d.getCollection()
	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		logger.Error("dao.User - FindByCondition", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return nil
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		logger.Error("dao.User - FindByCondition - decode", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return nil
	}
	return docs
}

// FindByIDs ...
func (d User) FindByIDs(ctx context.Context, ids []primitive.ObjectID) []model.User {
	cond := bson.M{"_id": bson.M{"$in": ids}}
	return d.FindByCondition(ctx, cond)
}

// CountByCondition ...
func (d User) CountByCondition(ctx context.Context, cond interface{}) int64 {
	col := d.getCollection()
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("dao.User - CountByCondition", logger.LogData{
			"err":  err.Error(),
			"cond": cond,
		})
	}
	return total
}

func (d User) getCollection() *mongo.Collection {
	db := mongodb.GetDB()
	return db.Collection(constant.ColUser)
}
