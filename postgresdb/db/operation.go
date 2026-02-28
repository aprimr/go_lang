package db

import (
	"database/sql"
	"fmt"
)

func InsertTodo(db *sql.DB, title string, isCompleted bool) {
	_, err := db.Exec("INSERT INTO todos (title, is_completed) VALUES($1, $2)", title, isCompleted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data inserted success.")
}
