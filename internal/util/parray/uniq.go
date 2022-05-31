package parray

import (
	"myapp/internal/mongodb"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UniqString ...
func UniqString(arr []string) []string {
	return funk.UniqString(arr)
}

// UniqObjectID ...
func UniqObjectID(arr []primitive.ObjectID) []primitive.ObjectID {
	str := make([]string, len(arr))
	for i, s := range arr {
		str[i] = s.Hex()
	}
	str = UniqString(str)
	res := make([]primitive.ObjectID, len(str))

	for i, s := range str {
		id, _ := mongodb.NewIDFromString(s)
		res[i] = id
	}
	return res
}
