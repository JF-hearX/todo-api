package todolistsql

import (
	"context"
	"fmt"
	"time"

	"github.com/JF-hearX/todo-api/models"
	"github.com/jmoiron/sqlx"
)

type MysqlRepo struct {
	DBClient *sqlx.DB
}

func (m *MysqlRepo) Insert(ctx context.Context, todo models.TodoListCreate) (int64, error) {
	const sqlstr = `INSERT INTO todolist 
                     (title, description, due_date) VALUES 
                      (:title, :description, :due_date);`

	_, err := m.DBClient.NamedExecContext(ctx, sqlstr, todo)
	if err != nil {
		return 0, fmt.Errorf("failed to insert: %w", err)
	}

	const getIDQuery = `SELECT id FROM todolist ORDER BY id DESC LIMIT 1;`
	var id int64
	err = m.DBClient.GetContext(ctx, &id, getIDQuery)
	if err != nil {
		return 0, fmt.Errorf("failed to get last inserted ID: %w", err)
	}

	return id, nil
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

		cursor = todos[len(todos)-1].ID
	}

	res := &FindResult{
		Todos:  todos,
		Cursor: cursor,
	}

	return res, nil
}

func (m *MysqlRepo) Update(ctx context.Context, todo models.TodoList) error {
	if todo.ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	data := struct {
		ID          uint64     `db:"id"`
		Title       string     `db:"title"`
		Description string     `db:"description"`
		DueDate     *time.Time `db:"due_date"`
		Completed   bool       `db:"completed"`
	}{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.Due_Date,
		Completed:   todo.Completed,
	}

	const sqlstr = `UPDATE todolist SET title = :title, description = :description, due_date = :due_date, completed = :completed WHERE id = :id`

	result, err := m.DBClient.NamedExecContext(ctx, sqlstr, data)
	if err != nil {
		return fmt.Errorf("failed to update: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no changes made - affected rows: %d - please make an update to all objects in array", rowsAffected)
	}

	return nil
}
