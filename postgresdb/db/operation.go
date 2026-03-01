package db

import (
	"database/sql"
	"fmt"
	"math"

	"github.com/aprimr/goanddatabase/models"
)

func InsertTodo(db *sql.DB, title string, isCompleted bool) error {
	_, err := db.Exec("INSERT INTO todos (title, is_completed) VALUES($1, $2)", title, isCompleted)
	if err != nil {
		return err
	}
	return nil
}

func FetchAllTodos(db *sql.DB, page int, limit int) (models.PaginatedTodos, error) {
	// Calculate skip offset value
	offset := (page - 1) * limit

	// First get total count
	var totalCount int
	err := db.QueryRow("SELECT COUNT(*) FROM todos").Scan(&totalCount)
	if err != nil {
		return models.PaginatedTodos{}, err
	}

	rows, err := db.Query("SELECT id, title, is_completed FROM todos ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return models.PaginatedTodos{}, err
	}
	defer rows.Close()

	var todos []models.Todo

	// Loop through rows
	for rows.Next() {
		var todo models.Todo

		// Scan
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Is_completed)
		if err != nil {
			return models.PaginatedTodos{}, err
		}

		todos = append(todos, todo)
	}

	// Calculate page count
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return models.PaginatedTodos{
		Data:       todos,
		Page:       page,
		Limit:      limit,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}

func FetchTodosByID(db *sql.DB, id int) (*models.Todo, error) {
	var todo models.Todo
	row := db.QueryRow("SELECT id, title, is_completed FROM todos WHERE id=$1", id)
	err := row.Scan(&todo.Id, &todo.Title, &todo.Is_completed)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func DeleteTodosByID(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return err
	}

	// Check if something is deleted or not
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Todo not found")
	}

	return nil
}

func UpdateTodo(db *sql.DB, id int, newtodo models.Todo) error {
	result, err := db.Exec("UPDATE todos SET title=$1, is_completed=$2 WHERE id=$3", newtodo.Title, newtodo.Is_completed, id)
	if err != nil {
		return err
	}

	// Check if something is actually updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Todo not found")
	}

	return nil
}
