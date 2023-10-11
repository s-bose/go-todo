package models

import (
	"go-todo/app/enums"
	"time"
)

type Todo struct {
	ID          uint             `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      enums.StatusType `json:"status"`
	UserID      uint             `json:"user_id"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}
