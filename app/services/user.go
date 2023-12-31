package services

import (
	"go-todo/app/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (r *UserService) CreateUser(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserService) GetUserById(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.Where(&models.User{ID: id}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserService) AuthenticateUser(email string, password string, user *models.User) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
}
