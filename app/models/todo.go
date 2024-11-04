package models

import (
	"go-todo/app/enums"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uint               `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uuid.UUID          `json:"user_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Priority    enums.PriorityType `json:"priority"`
	Status      enums.StatusType   `json:"status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
