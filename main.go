package main

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/JF-hearX/todo-api/application"
	// // Autoload .env file
	// _ "github.com/joho/godotenv/autoload"
)

func main() {

	fx.New(
		fx.Provide(
			application.DatabaseProvider,
			application.NewHTTPServer,
			application.NewHTTPRouter),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
