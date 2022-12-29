package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/add-todo", func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var todoPayload models.Todo
		todoPayload.IsCompleted = false

		err = json.Unmarshal(reqBody, &todoPayload)
		if err != nil {
			panic(err)
		}

		services.AddTodo(collection, todoPayload)
	})

	http.HandleFunc("/delete-todo/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/delete-todo/")
		services.DeleteTodo(collection, id)
	})

	http.HandleFunc("/complete-todo/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/complete-todo/")
		services.CompleteTodo(collection, id)
	})

	http.HandleFunc("/complete-all-todos", func(w http.ResponseWriter, r *http.Request) {
		services.CompleteAllTodos(collection)
	})

	http.ListenAndServe(":3000", nil)
}
