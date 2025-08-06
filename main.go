package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevOps-Group-D/YouToFy-API/configs"
	controllersAcc "github.com/DevOps-Group-D/YouToFy-API/controllers"
	"github.com/go-chi/chi"
	"github.com/golang-migrate/migrate/v4"
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
	migrationsPath := "file://./migrations"
	runMigrations(migrationsPath, cfg.DBConfig.URL)

	// Listening and serving service
	router := chi.NewRouter()

	// Registering controllers
	router.Put("/register", controllersAcc.Register)
	router.Post("/login", controllersAcc.Login)
	router.Get("/authorize", controllersAcc.Authorize)

	fmt.Println("Listening and serving on port", cfg.ApiConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.ApiConfig.Port), router)
}

func runMigrations(migrationsPath, databaseURL string) {
	log.Println("Running database migrations...")

	m, err := migrate.New(
		migrationsPath,
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database is up to date!")
}
