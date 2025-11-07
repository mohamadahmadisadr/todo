package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID	primitive.ObjectID `bson:"_id.omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Status string `bson:"status"`
	Priority string `bson:"Priority"`
	DueDate time.Time `bson:"due_date"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`

}