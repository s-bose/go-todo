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

type UpdateTodoSchema CreateTodoSchema

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
	userID, ok := r.Context().Value("userID").(uuid.UUID)

	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Dict{"message": "unauthorized"})
		return
	}

	var newTodo CreateTodoSchema
	err := utils.ParseJSON(r, &newTodo)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Dict{"message": "error parsing request body", "error": err.Error()})
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

func (t *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(uuid.UUID)

	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Dict{"message": "unauthorized"})
		return
	}

	todos, err := t.todoService.GetAllTodosByUserID(userID)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Dict{"message": "error fetching todos", "error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, todos)

}

func (t *TodoController) GetTodoById(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(r.PathValue("id"))

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Dict{"message": "error parsing id", "error": err.Error()})
		return
	}

	todo, err := t.todoService.GetTodoById(uuid)
	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, utils.Dict{"message": "error fetching todo", "error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, todo)
}

func (t *TodoController) UpdateTodoById(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(r.PathValue("id"))

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Dict{"message": "error parsing id", "error": err.Error()})
		return
	}

	var updateTodo UpdateTodoSchema
	err = utils.ParseJSON(r, &updateTodo)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Dict{"message": "error parsing request body", "error": err.Error()})
		return
	}

	todo, err := t.todoService.UpdateTodoById(uuid, updateTodo.Title, updateTodo.Description, enums.PriorityType(updateTodo.Priority), enums.StatusType(updateTodo.Status))
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Dict{"message": "error updating todo", "error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, todo)
}

func (t *TodoController) DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(r.PathValue("id"))

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Dict{"message": "error parsing id", "error": err.Error()})
		return
	}

	err = t.todoService.DeleteTodoById(uuid)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Dict{"message": "error deleting todo", "error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, utils.Dict{})
}
