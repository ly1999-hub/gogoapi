package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Configuration ...
type Configuration struct {
	ID        primitive.ObjectID `bson:"_id"`
	Order     ConfigurationOrder `bson:"order"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

// ConfigurationOrder ...
type ConfigurationOrder struct {
	CancelledReasons []string `bson:"cancelledReasons" json:"cancelledReasons"`
}
