package db

import (
	"context"
	"fmt"
	"log"
	"math"
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
			fmt.Println("Error inserting eval", err)
			log.Fatal("Error inserting eval", err)
		}
	}
}

// Can adjust later but for now, we'll default to 500 items per query and 10 queries per batch
// As values this takes a slice of strings
func (db *Database) insertMany(ctx context.Context, target string, values [][]string, errorOnConflict bool) []error {
	const (
		ITEMS_PER_QUERY   = 1000
		QUERIES_PER_BATCH = 50
	)
	var (
		valLen       = len(values)
		errors       []error
		numBatches   = int(math.Ceil(float64(valLen) / float64((ITEMS_PER_QUERY * QUERIES_PER_BATCH))))
		currentBatch = 0
		queryBatch   = lo.Chunk(values, ITEMS_PER_QUERY)
	)

	batches := []*pgx.Batch{}

	for i := 0; i < numBatches; i++ {
		batches = append(batches, &pgx.Batch{})
	}

	conn, err := db.conn.Acquire(ctx)
	// defer db.conn.Close()
	// defer conn.Conn().Close(ctx)
	defer conn.Release()
	if err != nil {
		fmt.Println("error getting conn", err)
	}

	for idx, batchItem := range queryBatch {
		batch := batches[currentBatch]
		query := createBatchedQueryString(target, batchItem, errorOnConflict)
		batch.Queue(query)
		batchSize := batch.Len()
		if batchSize > QUERIES_PER_BATCH {
			fmt.Println(batchSize, valLen)
		}

		if idx == len(queryBatch)-1 || batchSize%QUERIES_PER_BATCH == 0 {
			// Batch limit hit, sending batch
			results := conn.SendBatch(ctx, batch)
			fmt.Println("Sending Batch", batchSize, valLen)
			for i := 0; i < batchSize; i++ {
				_, err := results.Exec()
				if err != nil {
					fmt.Println(err)
					errors = append(errors, err)
				}
			}
			closeErr := results.Close()
			if closeErr != nil {
				fmt.Println("CLOSE ERROR", closeErr)
			} else {
				fmt.Println(idx+1, "chunks of ", len(queryBatch), " sent. ID: ", valLen)
				fmt.Println(currentBatch, "/", numBatches)
				currentBatch++
			}
		}
	}
	return errors
}

func (db *Database) copyInto(ctx context.Context, table string, rows [][]any, columns []string) (int, error) {
	copyCount, err := db.conn.CopyFrom(
		ctx,
		pgx.Identifier{table},
		columns,
		pgx.CopyFromRows(rows),
	)
	return int(copyCount), err
}

func createBatchedQueryString(target string, items [][]string, errorOnConflict bool) string {
	var query = &strings.Builder{}
	query.WriteString("INSERT INTO ")
	query.WriteString(target)
	query.WriteString(" VALUES ")
	for idx, item := range items {
		query.WriteRune('(')
		query.WriteString(strings.Join(item, ", "))
		if idx == len(items)-1 {
			query.WriteString(")")
		} else {
			query.WriteString("), ")
		}

	}
	if !errorOnConflict {
		query.WriteString(" ON CONFLICT DO NOTHING")
	}
	query.WriteRune(';')

	// fmt.Println(query.String())
	return query.String()
}
