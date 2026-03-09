package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	databaseURI := os.Getenv("DATABASE_URI")
	pool, err := pgxpool.New(context.Background(), databaseURI)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	DB = pool
	log.Println("Database connected")
}
