package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(databaseURL string) *pgxpool.Pool {

	db, err := pgxpool.New(context.Background(), databaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connected successfully")

	return db
}
