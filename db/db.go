package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	username      = os.Getenv("DB_USERNAME")
	password      = os.Getenv("DB_PASSWORD")
	host          = os.Getenv("DB_HOST")
	port          = os.Getenv("DB_PORT")
	database_name = os.Getenv("DB_NAME")
)

type DbConfig struct {
	MinConnections string
	MaxConnections string
}

func New(config DbConfig) *pgxpool.Pool {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_min_conns=%s pool_max_conns=%s",
		username, password, host, port, database_name, config.MinConnections, config.MaxConnections)
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return dbPool
}
