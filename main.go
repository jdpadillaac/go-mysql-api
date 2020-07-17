package main

import (
	"github.com/jdpadillaac/go-mysql-api/handlers"
	"github.com/jdpadillaac/go-mysql-api/repository"
)


func main() {

	repository.NewMySQLDB()
	handlers.Handlers()
}
