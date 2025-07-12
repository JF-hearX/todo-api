package main

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/JF-hearX/todo-api/application"
)

func main() {

	fx.New(
		fx.Provide(
			application.DatabaseProvider,
			application.NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
