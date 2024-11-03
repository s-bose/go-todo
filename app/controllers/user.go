package controllers

import (
	"go-todo/db"
	"net/http"
)

func CreateUserRouter(DB *db.Database) *http.ServeMux {
	userRouter := http.NewServeMux()
	userController := UserController{DB: DB}
	userRouter.HandleFunc(
		"POST /register", userController.RegisterUser,
	)
	userRouter.HandleFunc(
		"POST /login", userController.LoginUser,
	)
	userRouter.HandleFunc(
		"GET /", userController.GetUser,
	)

	return userRouter
}

type UserController struct {
	DB *db.Database
}

func (u *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register User"))
}

func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login User"))

}

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))

}
