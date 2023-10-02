package models

import "github.com/google/uuid"

type Users struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type Board struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
