package controllers

import (
	"github.com/jdpadillaac/go-mysql-api/services/implementations"
	"net/http"
)

var (
	postService implementations.UserService
)

type UserController struct {}

type UserControllerInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}



func NewUserController(service  implementations.UserService) *UserController{
	postService = service
	return &UserController{}
}