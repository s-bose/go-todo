package services

import (
	"fmt"
	"go-todo/app/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	err = u.Db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUserById(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := u.Db.Where(&models.User{ID: id}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.Db.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) GetAuthenticatedUser(email string, password string) (*models.User, error) {
	user, err := u.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid username / password")
	}

	return user, nil
}
