package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/JF-hearX/todo-api/models"
	"github.com/JF-hearX/todo-api/repository/todolistsql"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type TodoRepo struct {
	Repo *todolistsql.MysqlRepo
}

func RouteAPI(r chi.Router, db *sqlx.DB) {
	todoRepo := &TodoRepo{
		Repo: &todolistsql.MysqlRepo{
			DBClient: db,
		},
	} // Assuming "db" is your sqlx.DB instance

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})

		// Add POST route for creating a new Todo item
		r.Post("/", todoRepo.Create)
	})
}

func (t *TodoRepo) Create(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Due_Date    *time.Time `json:"due_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("Format Error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todoinsert := models.TodoListCreate{
		Title:       body.Title,
		Description: body.Description,
		Due_Date:    body.Due_Date,
	}

	err := t.Repo.Insert(r.Context(), todoinsert)
	if err != nil {
		fmt.Println("failed to insert: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	res, err := json.Marshal(todoinsert)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusCreated)
}
