package services

import (
	"context"
	"todo-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	for index, result := range results {
		results[index].ObjectId = result.ID.Hex()
	}

	return results
}

func DeleteTodo(collection *mongo.Collection, objectId string) *mongo.DeleteResult {
	primitiveObjectId, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		panic(nil)
	}

	filter := bson.D{{Key: "_id", Value: primitiveObjectId}}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return result
}

func CompleteTodo(collection *mongo.Collection, objectId string) *mongo.UpdateResult {
	primitiveObjectId, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		panic(nil)
	}

	filter := bson.D{{Key: "_id", Value: primitiveObjectId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isCompleted", Value: true}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	return result
}

func CompleteAllTodos(collection *mongo.Collection) *mongo.UpdateResult {
	filter := bson.D{}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isCompleted", Value: true}}}}

	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	return result
}
