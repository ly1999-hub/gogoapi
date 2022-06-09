package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Role ...
type Role struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Code        string             `bson:"code"`
	Permissions []string           `bson:"permissions"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}
