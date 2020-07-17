package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	once sync.Once
)

func NewMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:12345678@tcp(localhost:3306)/go-test?parseTime=true")
		if err != nil {
			log.Fatalf("Error en conexion a la base de datos: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Imposible verificar estado de conexion: %v", err)
		}

		fmt.Println("Conexion establecida con MySQL")
	})
}

func Pool() *sql.DB {
	return db
}