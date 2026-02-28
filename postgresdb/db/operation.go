package db

import (
	"database/sql"
	"fmt"

	"github.com/aprimr/goanddatabase/models"
)

func InsertTodo(db *sql.DB, title string, isCompleted bool) {
	_, err := db.Exec("INSERT INTO todos (title, is_completed) VALUES($1, $2)", title, isCompleted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data inserted success.")
}

func FetchAllTodos(db *sql.DB) ([]models.Todo, error) {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var todos []models.Todo

	// Loop through rows
	for rows.Next() {
		var todo models.Todo

		// Scan
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Is_completed)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}
	return todos, nil
}

func FetchTodosByID(db *sql.DB, id int) ([]models.Todo, error) {
	rows, err := db.Query("SELECT * FROM todos WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	// Loop through rows
	for rows.Next() {
		var todo models.Todo

		// Scan each data and store in todo struct
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Is_completed)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
