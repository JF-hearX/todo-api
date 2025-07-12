package application

import (
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

const driverName = "mysql"

func DatabaseProvider(lc fx.Lifecycle) (*sqlx.DB, error) {
	connectionString, ok := os.LookupEnv("MYSQL_DSN")
	if !ok {
		return nil, fmt.Errorf("MYSQL_DSN environment variable not set")
	}

	var err error
	db, err := sqlx.Connect(driverName, connectionString)
	if err != nil {
		fmt.Println("Database Error:", err)
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}
