package services

import (
	"go-todo/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoService struct {
	Db *gorm.DB
}

func (r *TodoService) Create(todo *models.Todo) (*models.Todo, error) {
	return todo, r.Db.Create(todo).Error
}

func (r *TodoService) GetTodoById(id uint) (*models.Todo, error) {
	var todo models.Todo

	if err := r.Db.Where(&models.Todo{ID: id}).First(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoService) GetAllTodos(userID uuid.UUID) ([]models.Todo, error) {
	var todos []models.Todo
	if err := r.Db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// func (r *TodoService) UpdateTodoById(todo *models.Todo) (*models.Todo, error) {
// 	//
// }

func (r *TodoService) DeleteTodoById(id uint) error {
	return r.Db.Delete(&models.Todo{}, id).Error
}
