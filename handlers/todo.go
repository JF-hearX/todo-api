package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// type Todo struct {
// 	Repo *todolist.mysqlRepo
// }

// func (td *Todo) Insert(ctx context.Context, todo models.TodoList) error {

// 	if err != nil {
// 		return fmt.Errorf("failed to encode order: %w", err)
// 	}

// }

func RouteAPI(r chi.Router, db *sqlx.DB) {
	r.Route("/hello", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
	})

}
