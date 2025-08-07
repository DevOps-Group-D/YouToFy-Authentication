package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
)

const MIGRATIONS_PATH = "file://./migrations"

func RunMigrations(databaseURL string) {
	log.Println("Running database migrations...")

	m, err := migrate.New(
		MIGRATIONS_PATH,
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
