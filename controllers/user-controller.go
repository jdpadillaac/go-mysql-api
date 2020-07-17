package controllers

import (
	"encoding/json"
	"github.com/jdpadillaac/go-mysql-api/models"
	"github.com/jdpadillaac/go-mysql-api/services/implementations"
	"net/http"
)

var (
	postService implementations.UserService
)

type userController struct {}

type UserController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	CreateNewUser(w http.ResponseWriter, r *http.Request)
}


func NewUserController(service  implementations.UserService) UserController{
	postService = service
	return &userController{}
}

func (u *userController) GetAll(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (u *userController) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error en los datos del body", http.StatusBadRequest)
		return
	}


	userCreated, _ := postService.Create(user)
	resp := newResponse(true, "Usuario registrado", userCreated, nil)
	newResponseJSON(w, http.StatusCreated, resp)
}

