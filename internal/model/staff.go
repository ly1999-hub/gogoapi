package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Staff ...
type Staff struct {
	ID             primitive.ObjectID  `bson:"_id"`
	Name           string              `bson:"name"`
	SearchString   string              `bson:"searchString"`
	Email          string              `bson:"email"`
	Phone          string              `bson:"phone"`
	Active         bool                `bson:"active"`
	Role           *primitive.ObjectID `bson:"role,omitempty"`
	Avatar         string              `bson:"avatar"`
	IsRoot         bool                `bson:"isRoot"`
	HashedPassword string              `bson:"hashedPassword"`
	CreatedAt      time.Time           `bson:"createdAt"`
	UpdatedAt      time.Time           `bson:"updatedAt"`
}
