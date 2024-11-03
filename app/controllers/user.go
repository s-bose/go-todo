package controllers

import (
	"encoding/json"
	"go-todo/app/models"
	"go-todo/app/services"
	"go-todo/app/utils"
	"go-todo/db"
	"net/http"
)

func RegisterUserRoutes(DB *db.Database) *http.ServeMux {

	userController := &UserController{
		DB: DB,
		userService: &services.UserService{
			Db: DB.Db,
		},
	}

	userMux := http.NewServeMux()
	userMux.HandleFunc("GET /{$}", userController.GetUser)
	userMux.HandleFunc("POST /register", userController.RegisterUser)
	userMux.HandleFunc("POST /login", userController.LoginUser)

	return userMux
}

type UserController struct {
	DB          *db.Database
	userService *services.UserService
}

type RegisterUserSchema struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user RegisterUserSchema

	err := utils.ParseJSON(r, &user)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	createdUser, e := u.userService.CreateUser(&models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if e != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, e)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user LoginUserSchema

	err := utils.ParseJSON(r, &user)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	var userDb *models.User
	userDb, err = u.userService.GetAuthenticatedUser(user.Email, user.Password)

	if err != nil {
		utils.WriteJSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, userDb)
}

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}
