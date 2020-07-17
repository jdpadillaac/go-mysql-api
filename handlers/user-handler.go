package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jdpadillaac/go-mysql-api/controllers"
	"github.com/jdpadillaac/go-mysql-api/repository"
	"github.com/jdpadillaac/go-mysql-api/services/implementations"
)


func userRoutes(r *mux.Router) {

	userStorage := repository.NewUserRepository(repository.Pool())
	userSrv := implementations.NewUserService(userStorage)
	userCtr := controllers.NewUserController(userSrv)

	r.HandleFunc("/api/user", userCtr.CreateNewUser ).Methods("POST")

}
//