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

type TaskRepository interface {
	CreateTask(ctx context.Context, task *model.Task) (*primitive.ObjectID, error)
	FindTasks(ctx context.Context, id string) ([]*model.Task, error)
}

type TaskRepo struct {
	collection *mongo.Collection
}

func NewTaskRepository(collecton *mongo.Collection) TaskRepository {
	return &TaskRepo{
		collection: collecton,
	}
}

func (r *TaskRepo) CreateTask(ctx context.Context, task *model.Task) (*primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, task)

	if err != nil {
		log.Printf("failed to insert task : %v ", err)
		return nil, err
	}
	id, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, errors.New("failed to cased inserted id to objID")
	}
	return &id, nil

}

func (r *TaskRepo) FindTasks(ctx context.Context, id string) ([]*model.Task, error) {
	ObjID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	filter := bson.M{"user_id": ObjID}
	var tasks []*model.Task

	cursor, err := r.collection.Find(ctx, filter)

	if err != nil {
		return nil, errors.New("failed retrieving data")
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, errors.New("error decoding all documents")
	}

	return tasks, nil
}
