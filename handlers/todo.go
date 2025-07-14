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
	}

	r.Route("/", func(r chi.Router) {
		// Add routes
		r.Post("/", todoRepo.Create)
		r.Get("/", todoRepo.GetAll)
		r.Patch("/", todoRepo.Update)
	})
}

func (t TodoRepo) Create(w http.ResponseWriter, r *http.Request) {
	var body []struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Due_Date    *time.Time `json:"due_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, fmt.Sprintf("Format Error: %v", err), http.StatusBadRequest)
		return
	}

	var ids []int64
	for _, createData := range body {
		todoCreate := models.TodoListCreate{
			Title:       createData.Title,
			Description: createData.Description,
			Due_Date:    createData.Due_Date,
		}

		id, err := t.Repo.Insert(r.Context(), todoCreate)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to insert: %v", err), http.StatusInternalServerError)
			return
		}

		ids = append(ids, id)
	}

	res, err := json.Marshal(&ids) // Return the IDs in an array
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal: %v", err), http.StatusInternalServerError)
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

	pageSizeStr := r.URL.Query().Get("pagesize")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}

	pageSize, err := strconv.ParseUint(pageSizeStr, 0, 64)
	if err != nil {
		fmt.Println("pagesize Error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
	if err != nil {
		fmt.Println("Cursor Error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	// Define the page size and use the cursor as offset for pagination

	findAllPage := todolistsql.FindAllPage{
		Size:   pageSize,
		Offset: cursor/pageSize + 1,
	}

	res, err := t.Repo.GetAll(r.Context(), findAllPage)
	if err != nil {
		fmt.Println("failed to all:", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, fmt.Sprintf("failed to all: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		http.Error(w, fmt.Sprintf("failed to marshal: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (t *TodoRepo) Update(w http.ResponseWriter, r *http.Request) {
	var body []struct {
		ID          uint64     `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Due_Date    *time.Time `json:"due_date"`
		Completed   bool       `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("Invalid request format:", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	for _, updateData := range body {
		todoUpdate := models.TodoList{
			ID:          updateData.ID,
			Title:       updateData.Title,
			Description: updateData.Description,
			Due_Date:    updateData.Due_Date,
			Completed:   updateData.Completed,
		}

		err := t.Repo.Update(r.Context(), todoUpdate)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to update: %v", err), http.StatusInternalServerError)
			return
		}
	}

	res, err := json.Marshal(&body)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
