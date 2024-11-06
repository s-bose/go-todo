package controllers

import (
	"go-todo/app/auth"
	"go-todo/app/db"
	"go-todo/app/enums"
	"go-todo/app/models"
	"go-todo/app/services"
	"go-todo/app/utils"
	"net/http"

	"github.com/google/uuid"
)

type CreateTodoSchema struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}

func RegisterTodoRoutes(DB *db.Database) *http.ServeMux {
	todoController := &TodoController{
		DB: DB,
		todoService: &services.TodoService{
			Db: DB.Db,
		},
	}

	todoMux := http.NewServeMux()
	todoMux.HandleFunc("GET /{$}", auth.JWTMiddleware(todoController.GetAllTodos, DB))
	todoMux.HandleFunc("POST /{$}", auth.JWTMiddleware(todoController.CreateTodo, DB))
	todoMux.HandleFunc("GET /{id}", auth.JWTMiddleware(todoController.GetTodoById, DB))
	todoMux.HandleFunc("PUT /{id}", auth.JWTMiddleware(todoController.UpdateTodoById, DB))
	todoMux.HandleFunc("DELETE /{id}", auth.JWTMiddleware(todoController.DeleteTodoById, DB))

	return todoMux
}

type TodoController struct {
	DB          *db.Database
	todoService *services.TodoService
}

func (t *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo CreateTodoSchema
	err := utils.ParseJSON(r, &newTodo)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Dict{"message": "error parsing request body", "error": err.Error()})
		return
	}

	userID, ok := r.Context().Value("userID").(uuid.UUID)

	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Dict{"message": "unauthorized"})
		return
	}

	todo, err := t.todoService.Create(&models.Todo{
		UserID:      userID,
		Title:       newTodo.Title,
		Description: newTodo.Description,
		Priority:    enums.PriorityType(newTodo.Priority),
		Status:      enums.StatusType(newTodo.Status),
	})

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Dict{"message": "error creating todo", "error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, todo)
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
