package main

import (
	"database/sql"
)

type todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (t *todo) createTodo(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO todos(name) VALUES($1) RETURNING id",
		t.Name).Scan(&t.ID)

	if err != nil {
		return err
	}

	return nil
}

func getTodos(db *sql.DB, start, count int) ([]todo, error) {
	rows, err := db.Query(
		"SELECT id, name FROM todos LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []todo{}

	for rows.Next() {
		var t todo
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}
