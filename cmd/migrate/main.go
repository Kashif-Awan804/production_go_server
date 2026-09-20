package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	// Create migration instance
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)

	if err != nil {
		log.Fatal("Failed to create migration:", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Use: migrate up | migrate down")
	}

	switch os.Args[1] {

	case "up":

		err := m.Up()

		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration failed:", err)
		}

		log.Println("Migrations applied successfully")

	case "down":

		err := m.Steps(-1)

		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration rollback failed:", err)
		}

		log.Println("Migration rolled back successfully")

	default:
		log.Fatal("Unknown command. Use: up | down")
	}
}
