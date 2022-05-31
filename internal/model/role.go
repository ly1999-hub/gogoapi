package model

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Role
type Role struct{
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Code        string             `bson:"code"`
	Permissions []string           `bson:"permissions"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}