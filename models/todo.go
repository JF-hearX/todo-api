package models

import "time"

type TodoList struct {
	ID          int64      `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	Due_Date    *time.Time `db:"due_date" json:"due_date"`
	Completed   bool       `db:"completed" json:"completed"`
}
