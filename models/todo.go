package models

import (
	"time"
)

type TodoList struct {
	ID          uint64     `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Due_Date    *time.Time `json:"due_date" db:"due_date"`
	Completed   bool       `json:"completed" db:"completed"`
}

type TodoListCreate struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Due_Date    *time.Time `json:"due_date"`
}

type TodoListCreateId struct {
	ID          int64      `db:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Due_Date    *time.Time `json:"due_date"`
}

type TodoListCreateNoDate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
