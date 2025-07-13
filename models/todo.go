package models

import (
	"time"
)

type TodoList struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Due_Date    *time.Time `json:"due_date"`
	Completed   bool       `json:"completed"`
}

type TodoListCreate struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Due_Date    *time.Time `json:"due_date"`
}

type TodoListCreateNoDate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
