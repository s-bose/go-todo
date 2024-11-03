package controllers

import (
	"go-todo/db"
	"net/http"
)

func RegisterTodoRoutes(DB *db.Database) *http.ServeMux {
	todoController := &TodoController{
		DB: DB,
	}

	todoMux := http.NewServeMux()
	todoMux.HandleFunc("POST /{$}", todoController.CreateTodo)
	todoMux.HandleFunc("GET /{id}", todoController.GetTodoById)
	todoMux.HandleFunc("PUT /{id}", todoController.UpdateTodoById)
	todoMux.HandleFunc("DELETE /{id}", todoController.DeleteTodoById)
	todoMux.HandleFunc("GET /{$}", todoController.GetAllTodos)

	return todoMux
}

type TodoController struct {
	DB *db.Database
}

func (t *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Todo"))
}

func (t *TodoController) GetTodoById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Todo By Id"))
}

func (t *TodoController) UpdateTodoById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Todo"))
}

func (t *TodoController) DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Todo"))
}

func (t *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get All Todos"))
}
