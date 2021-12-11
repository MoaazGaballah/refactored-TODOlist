package main

import (
	"database/sql"
	"errors"
)

type todo struct {
	ID   int    `json:"id"`
	name string `json:"name"`
}

func (p *todo) createTodo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getTodos(db *sql.DB, start, count int) ([]todo, error) {
	return nil, errors.New("Not implemented")
}
