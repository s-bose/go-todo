package models

import (
	"time"

	"github.com/google/uuid"
)

// main DTO model and gorm ORM
type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Todos     []Todo    `json:"todos" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// auxiliary DTOs for CRUD
type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserCreate struct {
	Name string `json:"name" validate:"required"`
	UserLogin
}
