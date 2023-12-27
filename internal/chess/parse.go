package chess

import (
	"base/internal/db"
	"base/types"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/smallnest/exp/chanx"
)

const (
	LOG_RATE     = 500_000
	IMPORT_LIMIT = 1_000_000
	CHUNK_SIZE   = 10_000
	SCAN_TYPE    = "batch" // "sync" "go", "go_local" "batch" "copy"
)

// The base file that this is ingesting is ~5GB
func ProcessFenPositions(ctx context.Context, db *db.Database, filepath string) {
	fmt.Println("Starting processing...")
	f, err := os.Open(filepath)
	if err != nil {
		slog.Error("Invalid filepath for import")
	}
	defer f.Close()
	var (
		count     = 1
		start     = time.Now()
		scanner   = bufio.NewScanner(f)
		q         = []*types.ImportedFenParent{}
		dataStore = make(chan *types.ImportedFenParent, CHUNK_SIZE)
		wg        = &sync.WaitGroup{}
	)

	for scanner.Scan() {
		count++
		memoryLogger(count, LOG_RATE)
		if count > IMPORT_LIMIT {
			break
		}

		imported := &types.ImportedFenParent{}
		if err := json.Unmarshal(scanner.Bytes(), imported); err != nil {
			slog.Warn("Error unmarshalling")
		}
		wg.Add(1)
		switch SCAN_TYPE {
		case "sync":
			q = append(q, imported)
		case "go":
			// This is slower than sync. See https://appliedgo.net/concurrencyslower/
			go func() {
				dataStore <- imported
			}()
		case "go_local":
			// No difference, causes an error on writing to a closed channel?
			go func(imp *types.ImportedFenParent, store chan *types.ImportedFenParent) {
				store <- imp
			}(imported, dataStore)
		case "batch":
			// https://medium.com/@smallnest/how-to-efficiently-batch-read-data-from-go-channels-7fe70774a8a5
			go func() {
				dataStore <- imported
			}()

		}

	}
	elapsed := time.Since(start)
	fmt.Println(count, "fens processed.")
	fmt.Printf("Parsing the file took [%v]\n", elapsed)
	fmt.Println(runtime.NumGoroutine(), "goroutines")
	for range dataStore {
		go chanx.Batch[*types.ImportedFenParent](ctx, dataStore, CHUNK_SIZE, func(ifp []*types.ImportedFenParent) {
			fmt.Println("original batch size per fen:", len(ifp))
			errors := db.InsertEvalLines(ctx, ifp)
			if len(errors) == 0 {
				fmt.Println("finishing...", len(ifp))
				wg.Done()
			} else {
				fmt.Println("rturned erros", errors)
			}
		})
	}
	fmt.Println("************************")
	// close(dataStore)
	// wg.Wait()
	fmt.Printf("Whole import took [%v]\n", time.Since(start))
	fmt.Println("Finished importing")
}

func memoryLogger(count int, logRate int) {
	if logRate == 0 {
		return
	}
	var memstat = runtime.MemStats{}

	if count%logRate == 0 {
		fmt.Println(count, "items processed.")
		fmt.Println(runtime.NumGoroutine(), "goroutines")
		runtime.ReadMemStats(&memstat)
		fmt.Printf("memstat.HeapAlloc: %+v\n", memstat.HeapAlloc)
		fmt.Printf("memstat.HeapInuse: %+v\n", memstat.HeapInuse)
		fmt.Printf("memstat.HeapReleased: %+v\n", memstat.HeapReleased)
		fmt.Printf("memstat.Alloc: %+v\n", memstat.Alloc)
		fmt.Printf("memstat.TotalAlloc: %+v\n", memstat.TotalAlloc)
		fmt.Printf("memstat.Sys: %+v\n", memstat.Sys)
		fmt.Printf("memstat.NextGC: %+v\n", memstat.NextGC)
		fmt.Printf("memstat.NumGC: %+v\n", memstat.NumGC)
	}
}
