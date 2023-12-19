package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/lo"
)

type DbConfig struct {
	MinConnections,
	MaxConnections string
}

type Database struct {
	conn *pgxpool.Pool
}

type DbHelper[T any] interface {
	FindById(id string) (T, error)
	DeleteById(id string) (bool, error)
	UpdateById(id string)
	Insert(entity T) (int, error)
	InsertMany(entity []T) (int, error)
}

func New(config DbConfig) *Database {
	var (
		Username      = os.Getenv("DB_USERNAME")
		Password      = os.Getenv("DB_PASSWORD")
		Host          = os.Getenv("DB_HOST")
		Port          = os.Getenv("DB_PORT")
		Database_name = os.Getenv("DB_NAME")
	)
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_min_conns=%s pool_max_conns=%s",
		Username, Password, Host, Port, Database_name, config.MinConnections, config.MaxConnections)

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

func (db *Database) batchRequests(ctx context.Context, query string, args []pgx.NamedArgs) {
	batch := &pgx.Batch{}
	for _, arg := range args {
		batch.Queue(query, arg)
	}

	results := db.conn.SendBatch(ctx, batch)
	defer results.Close()
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

// Can adjust later but for now, we'll default to 500 items per query and 10 queries per batch
func (db *Database) insertMany(ctx context.Context, target string, values []string, errorOnConflict bool) {

	chunked := lo.Chunk(values, 500)
	query := strings.Builder{}

	query.WriteString("INSERT INTO ")
	query.WriteString(target)
	query.WriteString(" VALUES ")
	for idx, chunk := range chunked {
		query.WriteRune('(')
		query.WriteString(strings.Join(chunk, ", "))
		if idx == len(chunked)-1 {
			query.WriteString(")")
		} else {
			query.WriteString("), ")
		}
	}

	// query.
	if !errorOnConflict {
		query.WriteString(" ON CONFLICT DO NOTHING")
	}
	query.WriteRune(';')

	// s.
	// query := fmt.Sprintf("INSERT INTO %s VALUES ")
}
