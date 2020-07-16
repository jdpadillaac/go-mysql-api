package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handlers() {

	router := mux.NewRouter()


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "45782"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}