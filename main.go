package main

import (
	"go-todo/app/controllers"
	"go-todo/db"
	"log"
	"net/http"
)

func main() {

	database, err := db.InitDatabase()
	if err != nil {
		log.Fatal("something went wrong while setting up database", err)
	}

	userController := &controllers.UserController{
		DB: database,
	}

	todoController := &controllers.TodoController{
		DB: database,
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend service up and running"))
	})

	router.HandleFunc("POST /api/v1/user/register", userController.RegisterUser)
	router.HandleFunc("POST /api/v1/user/login", userController.LoginUser)
	router.HandleFunc("GET /api/v1/user/", userController.GetUser)

	router.HandleFunc("GET /api/v1/todos/", todoController.GetAllTodos)
	router.HandleFunc("POST /api/v1/todos/", todoController.CreateTodo)
	router.HandleFunc("GET /api/v1/todos/{id}", todoController.GetTodoById)
	router.HandleFunc("PUT /api/v1/todos/{id}", todoController.UpdateTodoById)
	router.HandleFunc("DELETE /api/v1/todos/{id}", todoController.DeleteTodoById)

	http.ListenAndServe(":8000", router)
}
