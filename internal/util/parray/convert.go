package parray

import (
	"myapp/internal/mongodb"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ConvertStringsToObjectIDs ...
func ConvertStringsToObjectIDs(strValues []string) []primitive.ObjectID {
	return funk.Map(strValues, func(item string) primitive.ObjectID {
		id, _ := mongodb.NewIDFromString(item)
		return id
	}).([]primitive.ObjectID)
}
