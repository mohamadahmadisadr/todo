package repository

import (
	"context"
	"errors"
	"log"

	"todo/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*primitive.ObjectID, error)
	FindUserByID(ctx context.Context, id string) (*model.User, error)
}

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &UserRepo{
		collection: collection,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) (*primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, user)

	if err != nil {
		log.Println("Insert One Failed")
		return nil, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, errors.New("failed to cased inserted id to objID")
	}
	return &id, nil
}

func (r *UserRepo) FindUserByID(ctx context.Context, id string) (*model.User, error) {

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	filter := bson.M{"_id": objID}

	var user model.User

	err = r.collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // User not found
		}
		log.Printf("FindOne failed: %v", err)
		return nil, err
	}

	return &user, nil
}
