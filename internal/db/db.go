package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DbConfig struct {
	MinConnections,
	MaxConnections,
	Username,
	Password,
	Host,
	Port,
	Database_name string
}

type Database struct {
	conn *pgxpool.Pool
}

func New(config DbConfig) *Database {

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_min_conns=%s pool_max_conns=%s",
		config.Username, config.Password, config.Host, config.Port, config.Database_name, config.MinConnections, config.MaxConnections)

	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	pgInstance := &Database{conn: dbPool}

	return pgInstance
}

func (db *Database) Close() {
	db.conn.Close()
}

func (db *Database) BatchRequests(ctx context.Context, query string, args []pgx.NamedArgs) {
	batch := &pgx.Batch{}
	for range args {
		batch.Queue(query, args)
	}
	tx, err := db.conn.Begin(ctx)
	if err != nil {
		// fmt.Errorf("Error retrieving transaction: %w\n", err)
		// slog.Error("Error inserting eval")
		fmt.Println("Error retriving transaction", err)
	}
	defer tx.Rollback(ctx)
	results := tx.SendBatch(ctx, batch)

	for range args {
		_, err := results.Exec()

		if err != nil {
			// fmt.Errorf("Error inserting eval: %w", err)
			// slog.Error("Error inserting eval")
			fmt.Println("Error inserting eval", err)
			log.Fatal("Error inserting eval", err)
		}
	}

}
