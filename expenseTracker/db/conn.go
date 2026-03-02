package db

import (
	"context"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() (*pgxpool.Pool, error) {
	dbStr := os.Getenv("DATABASE_URL")

	// Try to open DB
	pool, err := pgxpool.New(context.Background(), dbStr)
	if err != nil {
		return nil, err
	}

	// Ping DB
	if err = pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return pool, nil
}
