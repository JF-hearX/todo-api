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
	"gopkg.in/yaml.v2"
)

const driverName = "mysql"

type Config struct {
	Db struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"db"`
	Api struct {
		ServerPort int `yaml:"server_port"`
	} `yaml:"api"`
}

func DatabaseProvider(lc fx.Lifecycle) (*sqlx.DB, error) {
	var config Config

	// Read the file 'config.yml' into a byte slice
	data, err := os.ReadFile("config.yml") // replace with the actual path to your config file
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println("DSN: ", config.Db.Dsn)

	// connectionString, ok := os.LookupEnv("MYSQL_DSN")
	// if !ok {
	// 	return nil, fmt.Errorf("MYSQL_DSN environment variable not set")
	// }

	db, err := sqlx.Connect(driverName, config.Db.Dsn)
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
	var config Config

	// Read the file 'config.yml' into a byte slice
	data, err := os.ReadFile("config.yml") // replace with the actual path to your config file
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println("DSN: ", config.Db.Dsn)

	// connectionString, ok := os.LookupEnv("MYSQL_DSN")
	// if !ok {
	// 	return nil, fmt.Errorf("MYSQL_DSN environment variable not set")
	// }

	db, err := sqlx.Connect(driverName, config.Db.Dsn)
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
	var config Config

	// Read the file 'config.yml' into a byte slice
	data, err := os.ReadFile("config.yml") // replace with the actual path to your config file
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println("DSN: ", config.Db.Dsn)

	// connectionString, ok := os.LookupEnv("MYSQL_DSN")
	// if !ok {
	// 	return nil, fmt.Errorf("MYSQL_DSN environment variable not set")
	// }

	db, err := sqlx.Connect(driverName, config.Db.Dsn)
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
