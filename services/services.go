package services

import (
	"context"
	"todo-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTodo(collection *mongo.Collection, todoPayload models.Todo) *mongo.InsertOneResult {
	result, err := collection.InsertOne(context.TODO(), todoPayload)
	if err != nil {
		panic(err)
	}
	return result
}

func GetAllTodos(collection *mongo.Collection) []models.Todo {
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []models.Todo
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	cursor.Decode(&results)

	return results
}
