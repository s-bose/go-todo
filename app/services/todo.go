package services

import (
	"go-todo/app/enums"
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

func (r *TodoService) GetTodoById(id uuid.UUID) (*models.Todo, error) {
	var todo models.Todo

	if err := r.Db.Where(&models.Todo{ID: id}).First(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoService) GetAllTodosByUserID(userID uuid.UUID) ([]models.Todo, error) {
	var todos []models.Todo

	if err := r.Db.Where(&models.Todo{UserID: userID}).Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoService) UpdateTodoById(id uuid.UUID, title, description string, priority enums.PriorityType, status enums.StatusType) (*models.Todo, error) {
	todo, err := r.GetTodoById(id)
	if err != nil {
		return nil, err
	}

	if title != "" {
		todo.Title = title
	}

	if description != "" {
		todo.Description = description
	}

	if priority != "" {
		todo.Priority = priority
	}

	if status != "" {
		todo.Status = status
	}

	err = r.Db.Save(todo).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoService) DeleteTodoById(id uuid.UUID) error {
	return r.Db.Delete(&models.Todo{}, id).Error
}
