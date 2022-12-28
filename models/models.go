package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	IsCompleted bool               `json:"isCompleted" bson:"isCompleted"`
}

type TodosData struct {
	PageTitle string
	Todos     []Todo
	AddTodo   func(*mongo.Collection, Todo) *mongo.InsertOneResult
}
