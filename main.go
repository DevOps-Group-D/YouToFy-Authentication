package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	controllersAcc "github.com/DevOps-Group-D/YouToFy-API/controllers"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// Reading .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initializing configs
	cfg := configs.LoadConfig()

	// Listening and serving service
	router := chi.NewRouter()

	// Registering controllers
	router.Put("/register", controllersAcc.Register)
	router.Post("/login", controllersAcc.Login)
	router.Get("/authorize", controllersAcc.Authorize)

	fmt.Println("Listening and serving on port", cfg.ApiConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.ApiConfig.Port), router)
}
