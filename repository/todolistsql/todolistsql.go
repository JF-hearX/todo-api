package todolistsql

import (
	"context"
	"fmt"

	"github.com/JF-hearX/todo-api/models"
	"github.com/jmoiron/sqlx"
)

type MysqlRepo struct {
	DBClient *sqlx.DB
}

func (m *MysqlRepo) Insert(ctx context.Context, todo models.TodoListCreate) error {
	data := models.TodoListCreate{ // Convert the model to a struct that sqlx understands
		Title:       todo.Title,
		Description: todo.Description,
		Due_Date:    todo.Due_Date,
	}

	// data, err := json.Marshal(todo)
	// if err != nil {
	// 	return fmt.Errorf("failed to encode order: %w", err)
	// }

	const sqlstr = `INSERT INTO todolist (title, description, due_date) VALUES (:title, :description, :due_date)`

	res, err := m.DBClient.NamedQueryContext(ctx, sqlstr, data)
	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	defer res.Close()

	return nil

}
