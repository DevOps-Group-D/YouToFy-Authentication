package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// Read .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	router := chi.NewRouter()

	fmt.Println("Listening and serving on localhost:3333")
	http.ListenAndServe(":3333", router)
}
