package controllers

import (
	"go-todo/db"
	"net/http"
)

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
