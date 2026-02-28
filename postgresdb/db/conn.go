package db

import (
	"database/sql"
	"fmt"
)

func ConnectDB() (db *sql.DB, err error) {
	connStr := "host=localhost port=5432 user=postgres password=admin123 dbname=tododb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")
	return
}
