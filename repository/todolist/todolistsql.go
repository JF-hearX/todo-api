package todolist

import (
	"context"
	"fmt"

	"github.com/JF-hearX/todo-api/models"
	"github.com/jmoiron/sqlx"
)

type mysqlRepo struct {
	db *sqlx.DB
}

func (m *mysqlRepo) Create(ctx context.Context, todo models.TodoList) error {
	const sqlstr = `INSERT INTO todolist (` +
		`title, description, due_date, ` +
		`) VALUES (` +
		`:title, :description, :due_date, ` +
		`) RETURNING id`

	res, err := m.db.NamedQueryContext(ctx, sqlstr, todo)
	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	for res.Next() {
		err = res.StructScan(&todo)
		if err != nil {
			return fmt.Errorf("failed to insert: %w", err)
		}
	}

	return nil

}
