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

type FindAllPage struct {
	Size   uint64
	Offset uint64
}

type FindResult struct {
	Todos  []models.TodoList
	Cursor uint64
}

func (m *MysqlRepo) GetAll(ctx context.Context, page FindAllPage) (*FindResult, error) {
	offset := (page.Offset - 1) * page.Size

	todos := []models.TodoList{}
	const sqlstr = `SELECT id, title, description, due_date, completed FROM todolist ORDER BY id LIMIT ? OFFSET ?`

	err := m.DBClient.SelectContext(ctx, &todos, sqlstr, page.Size, offset)
	if err != nil {
		fmt.Println("could not fetch todos: ", err)
		return nil, err
	}

	var cursor uint64
	if len(todos) > 0 {
		// assuming ID is the unique identifier for each TodoList item
		cursor = todos[len(todos)-1].ID
	}

	res := &FindResult{
		Todos:  todos,
		Cursor: cursor,
	}

	return res, nil
}
