package main

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/JF-hearX/todo-api/application"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func main() {

	apiCommand := &cobra.Command{
		Use:   "api",
		Short: "Run http api service",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fx.New(
				fx.Provide(
					application.DatabaseProvider,
					application.NewHTTPServer,
					application.NewHTTPRouter),
				fx.Invoke(func(*http.Server) {}),
			).Run()
		},
	}

	rootCommand := &cobra.Command{}
	rootCommand.AddCommand(apiCommand)
	if err := rootCommand.Execute(); err != nil {
		log.Error().Err(err).Msg("root command failed to execute")
	}

}
