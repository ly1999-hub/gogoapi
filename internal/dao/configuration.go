package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"myapp/internal/constant"
	"myapp/internal/model"
	"myapp/internal/module/logger"
	"myapp/internal/module/mongodb"
)

// Configuration ...
type Configuration struct{}

// UpdateOne ...
func (d Configuration) UpdateOne(ctx context.Context, cond, payload interface{}) error {
	col := d.getCollection()
	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		logger.Error("dao.Configuration - UpdateOne", logger.LogData{
			"cond":    cond,
			"payload": payload,
			"err":     err.Error(),
		})
	}
	return err
}

// FindOne ...
func (d Configuration) FindOne(ctx context.Context, cond interface{}) (model.Configuration, error) {
	var (
		doc model.Configuration
		col = d.getCollection()
	)

	err := col.FindOne(ctx, cond).Decode(&doc)
	if err != nil {
		logger.Error("dao.Configuration - FindOne", logger.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
	}
	return doc, err
}

func (d Configuration) getCollection() *mongo.Collection {
	db := mongodb.GetDB()
	return db.Collection(constant.ColConfiguration)
}
