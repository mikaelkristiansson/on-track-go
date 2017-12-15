package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/mikaelkristiansson/on-track-go/exercise"
)

func main() {
	router := exercise.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// launch server
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
