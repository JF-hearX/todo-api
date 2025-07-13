package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
		// Add routes for creating a new Todo item
		r.Post("/", todoRepo.Create)
		r.Get("/", todoRepo.GetAll)
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.WriteHeader(http.StatusCreated)
}

func (t *TodoRepo) GetAll(w http.ResponseWriter, r *http.Request) {
	// Parse cursor from URL Query
	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}

	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Define the page size and use the cursor as offset for pagination
	const pageSize = 10 // you can adjust this value based on your requirement
	findAllPage := todolistsql.FindAllPage{
		Size:   pageSize,
		Offset: cursor/pageSize + 1,
	}

	res, err := t.Repo.GetAll(r.Context(), findAllPage)
	if err != nil {
		fmt.Println("Failed to all:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
