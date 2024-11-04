package controllers

import (
	"encoding/json"
	"go-todo/app/auth"
	"go-todo/app/db"
	"go-todo/app/models"
	"go-todo/app/services"
	"go-todo/app/utils"
	"net/http"

	"github.com/google/uuid"
)

func RegisterUserRoutes(DB *db.Database) *http.ServeMux {

	userController := &UserController{
		DB: DB,
		userService: &services.UserService{
			Db: DB.Db,
		},
	}

	userMux := http.NewServeMux()
	userMux.HandleFunc("GET /{$}", auth.JWTMiddleware(userController.GetUser, DB))
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

	accessToken, err := auth.CreateJWT(&userDb.ID)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Dict{"message": "error creating access token", "error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, utils.Dict{"access_token": accessToken})
}

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Dict{"message": "unauthorized"})
		return
	}

	user, err := u.userService.GetUserById(userID)
	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, utils.Dict{"message": "user not found"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}
