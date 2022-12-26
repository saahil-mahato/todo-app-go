package services

import (
	"context"
	"todo-app/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddTodo(collection *mongo.Collection, todoPayload models.Todo) *mongo.InsertOneResult {
	result, err := collection.InsertOne(context.TODO(), todoPayload)
	if err != nil {
		panic(err)
	}
	return result
}
