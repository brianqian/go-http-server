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
)

const (
	LOG_RATE     = 500_000
	IMPORT_LIMIT = 3_500_000
	CHUNK_SIZE   = 10_000
)

// The base file that this is ingesting is ~5GB
func ProcessFenPositions(ctx context.Context, db *db.Database, filepath string) {
	f, err := os.Open(filepath)

	if err != nil {
		slog.Error("Invalid filepath for import")
	}
	defer f.Close()

	fmt.Println("Starting processing...")

	var (
		count   = 0
		start   = time.Now()
		scanner = bufio.NewScanner(f)
		wg      = &sync.WaitGroup{}
		q       = []*types.ImportedFenParent{}
	)

	for {
		count++
		status, err := processLine(ctx, wg, scanner, db, &count)
		if status == "done" && err == nil {
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Println(count, "fens processed.")
	fmt.Printf("To parse the file took [%v]\n", elapsed)
	wg.Wait()
	fmt.Println(runtime.NumGoroutine(), "goroutines")
	// Roughly 45 seconds

}

func processLine(ctx context.Context, wg *sync.WaitGroup, scanner *bufio.Scanner, db *db.Database, count *int, q []*types.ImportedFenParent) (string, error) {
	more := scanner.Scan()
	if !more {
		return "done", scanner.Err()
	}
	memoryLogger(count, LOG_RATE)
	wg.Add(1)
	imported := &types.ImportedFenParent{}
	if err := json.Unmarshal(scanner.Bytes(), imported); err != nil {
		slog.Error("Error unmarshaling\n", err)
	}

	go func(imported *types.ImportedFenParent) {
		defer wg.Done()
		db.InsertEvalLines(ctx, imported)
	}(imported)
	*count += CHUNK_SIZE
	return "continue", nil
}

func memoryLogger(count *int, logRate int) {
	if logRate == 0 {
		return
	}
	var memstat = runtime.MemStats{}

	if *count%logRate == 0 {
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
