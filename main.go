package main

import (
	"fmt"
	"github.com/jdpadillaac/go-mysql-api/handlers"
	"github.com/jdpadillaac/go-mysql-api/repository"
)

func main() {
	fmt.Println("Welcome")

	repository.NewMySQLDB()
	handlers.Handlers()
}
