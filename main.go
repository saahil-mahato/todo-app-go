package main

import (
	"context"
	"net/http"
	"text/template"
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

	tmpl := template.Must(template.ParseFiles("ui/layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := models.TodosData{
			PageTitle: "My todo list",
			Todos:     services.GetAllTodos(collection),
			AddTodo:   services.AddTodo,
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":3000", nil)
}
