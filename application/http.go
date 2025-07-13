package application

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/JF-hearX/todo-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func NewHTTPRouter(db *sqlx.DB) (chi.Router, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(5 * time.Minute))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	// Serve API endpoints from /api
	r.Route("/todo", func(r chi.Router) {
		handlers.RouteAPI(r, db)
	})

	return r, nil
}

func NewHTTPServer(lc fx.Lifecycle, r chi.Router) *http.Server {
	serverport, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		serverport = ":3000"
	}

	srv := &http.Server{
		Addr:         serverport,
		Handler:      r,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  60 * time.Second,
	}
	println("server running on port: ", serverport)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Err(err).Msg("Server error")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Shutting down server")
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
