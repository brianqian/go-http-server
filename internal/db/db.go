package db

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"os"

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

	return &Database{conn: dbPool}

}

func (db *Database) Close() {
	db.conn.Close()
}

func (db *Database) Conn(ctx context.Context) (*pgxpool.Conn, error) {
	return db.conn.Acquire(ctx)
}

/* DEPRECATED */
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

type NullValue *string

// Can adjust later but for now, we'll default to 500 items per query and 10 queries per batch
// As values this takes a slice of strings
func (db *Database) insertMany(ctx context.Context, table pgx.Identifier, target string, values [][]string, errorOnConflict bool) []error {
	const (
		ITEMS_PER_QUERY   = 1000
		QUERIES_PER_BATCH = 50
	)
	var (
		valLen       = len(values)
		currentBatch = 0
		numBatches   = int(math.Ceil(float64(valLen) / float64((ITEMS_PER_QUERY * QUERIES_PER_BATCH))))
		queryBatch   = lo.Chunk(values, ITEMS_PER_QUERY)
		batches      = make([]*pgx.Batch, numBatches)
		errors       []error
	)

	for i := range batches {
		batches[i] = &pgx.Batch{}
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
		query := createBatchedQueryString(table, target, batchItem, errorOnConflict)
		batch.Queue(query)
		batchSize := batch.Len()
		if batchSize > QUERIES_PER_BATCH {
			fmt.Println("Error -- batch overflow: ", batchSize, valLen)
		}

		if len(queryBatch) == 1 {
			if _, err := conn.Exec(ctx, query); err != nil {
				errors = append(errors, err)
			}
		} else {
			// When batch hits its max or is the last batch, send the batch
			if idx == len(queryBatch)-1 || batchSize%QUERIES_PER_BATCH == 0 {
				results := conn.SendBatch(ctx, batch)
				defer func() {
					if closeErr := results.Close(); closeErr != nil {
						fmt.Println("CLOSE ERROR", closeErr, "ID", valLen)
					} else {
						// fmt.Println(idx+1, "chunks of ", len(queryBatch), " sent. ID: ", valLen)
						// fmt.Println(currentBatch, "/", numBatches)
					}
					currentBatch++
				}()
				// fmt.Println("Sending Batch", batchSize, valLen)
				// Check for errors within batch
				for i := 0; i < batchSize; i++ {
					_, err := results.Exec()
					if err != nil {
						fmt.Println("Error in batch", err)
						errors = append(errors, err)
					}
				}
			}
		}
	}
	return errors
}

type copyFromFileArgs struct {
	ctx       context.Context
	table     string
	columns   []string
	parsingFn parsingFn
	filePath  string
}

// Doesn't work because each scanned line results in multiple chess lines and evaulations
func (db *Database) copyFromFile(c copyFromFileArgs) (int, error) {
	f, err := os.Open(c.filePath)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	sc := bufio.NewScanner(f)

	copyCount, err := db.conn.CopyFrom(
		c.ctx,
		pgx.Identifier{c.table},
		c.columns,
		copyFromFileSource(sc, c.parsingFn),
	)
	return int(copyCount), err
}
