package application

import (
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func DatabaseProvider(lc fx.Lifecycle) (*sqlx.DB, error) {
	url, ok := os.LookupEnv("MYSQL_DSN")
	if !ok {
		return nil, fmt.Errorf("MYSQL_DSN environment variable not set")
	}

	var err error
	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}
