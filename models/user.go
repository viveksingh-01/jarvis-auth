package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username" validate:"required,min=6,max=12"`
	Password  string             `bson:"password" validate:"required,min=8,max=12"`
	CreatedAt time.Time          `bson:"createdAt"`
}
