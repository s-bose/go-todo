package services

import (
	"go-todo/app/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(&todo).Error
}

func (r *TodoRepository) GetTodo() error {
	return nil
}
