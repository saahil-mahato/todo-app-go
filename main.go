package main

import (
	"context"
	"fmt"
	"todo-app/models"
	"todo-app/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb+srv://saahil:xLMjpTszF0cU5adN@cluster0.rruqn79.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("test").Collection("todos")

	var todoPayload models.Todo

	todoPayload.Title = "Task 1"
	todoPayload.Description = "Description 1"
	todoPayload.IsCompleted = true

	fmt.Println(services.AddTodo(collection, todoPayload))
}
