package application

import (
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
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

func DatabaseMigrationProvider() (*sqlx.DB, error) {
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

	n, err := migrate.Exec(db.DB, "mysql", migrate.FileMigrationSource{Dir: "./database/migrations"}, migrate.Up)
	if err != nil {
		log.Error().Err(err).Msg("Failed to run migrations")
		return nil, err
	}

	log.Debug().
		Int("migrations", n).
		Msg("Migrations completed successfully")

	return db, nil
}

func DatabaseMigrationDownProvider() (*sqlx.DB, error) {
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

	n, err := migrate.Exec(db.DB, "mysql", migrate.FileMigrationSource{Dir: "./database/migrations"}, migrate.Down)
	if err != nil {
		log.Error().Err(err).Msg("Failed to run migrations")
		return nil, err
	}

	log.Debug().
		Int("migrations", n).
		Msg("Migrations completed successfully")

	return db, nil
}
