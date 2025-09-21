package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DevOps-Group-D/YouToFy-Authentication/configs"
	"github.com/DevOps-Group-D/YouToFy-Authentication/controllers"
	"github.com/DevOps-Group-D/YouToFy-Authentication/database"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	// Reading args from cli
	port := "0"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Reading .env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	// Initializing configs
	cfg := configs.LoadConfig()
	if port != "0" {
		cfg.ApiConfig.Port = port
	}

	// Running migrations
	database.RunMigrations(cfg.DBConfig.Url)

	// Listening and serving service
	router := chi.NewRouter()
	router.Use(chimiddleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Registering controllers
	router.Post("/register", controllers.Register)
	router.Post("/login", controllers.Login)
	router.Post("/authorize", controllers.Authorize)

	fmt.Println("Listening and serving on port", cfg.ApiConfig.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.ApiConfig.Port), router); err != nil {
		fmt.Println(err)
	}
}
