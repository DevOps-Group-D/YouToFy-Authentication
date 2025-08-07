package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevOps-Group-D/YouToFy-Authentication/configs"
	"github.com/DevOps-Group-D/YouToFy-Authentication/controllers"
	"github.com/DevOps-Group-D/YouToFy-Authentication/database"
	"github.com/go-chi/chi"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	// Running migrations
	database.RunMigrations(cfg.DBConfig.Url)

	// Listening and serving service
	router := chi.NewRouter()

	// Registering controllers
	router.Put("/register", controllers.Register)
	router.Post("/login", controllers.Login)
	router.Get("/authorize", controllers.Authorize)

	fmt.Println("Listening and serving on port", cfg.ApiConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.ApiConfig.Port), router)
}
