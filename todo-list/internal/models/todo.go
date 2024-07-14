package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type TodoModel struct {
	DB *sql.DB
}

func (m *TodoModel) Insert(title, content string, expires int) (int, error) {
	stmt := `INSERT INTO todos (title, content, created, expires) VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TodoModel) Get(id int) (*Todo, error) {
	return nil, nil
}

func (m *TodoModel) Latest() ([]*Todo, error) {
	return nil, nil

}
