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
	return 0, nil
}

func (m *TodoModel) Get(id int) (*Todo, error) {
	return nil, nil
}

func (m *TodoModel) Latest() ([]*Todo, error) {
	return nil, nil

}
