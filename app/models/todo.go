package models

import (
	"go-todo/app/enums"
	"time"

	"github.com/google/uuid"
)

// main DTO model and gorm ORM
type Todo struct {
	ID          uuid.UUID          `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      uuid.UUID          `json:"user_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Priority    enums.PriorityType `json:"priority"`
	Status      enums.StatusType   `json:"status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// auxiliary DTOs for CRUD
type TodoCreate struct {
	Title       string             `json:"title" validate:"required"`
	Description string             `json:"description"`
	Priority    enums.PriorityType `json:"priority" validate:"oneof='TODO' 'IN PROGRESS' 'DONE' 'CANCELLED'"`
	Status      enums.StatusType   `json:"status" validate:"oneof='HIGH' 'MEDIUM' 'LOW' 'VERY LOW'"`
}

type TodoUpdate struct {
	TodoCreate
}
