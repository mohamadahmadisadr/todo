package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID	primitive.ObjectID `bson:"_id.omitempty"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	UserType string `bson:"userType"`

}