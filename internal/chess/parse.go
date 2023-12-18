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

func ProcessFenPositions(ctx context.Context, db *db.Database, filepath string) []*types.ImportedFenParent {
	f, err := os.Open(filepath)

	if err != nil {
		slog.Error("Invalid filepath for import")
	}
	defer f.Close()

	fmt.Println("Starting processing...")
	count := 0
	start := time.Now()
	fenSlice := []*types.ImportedFenParent{}
	scanner := bufio.NewScanner(f)
	wg := sync.WaitGroup{}
	memstat := runtime.MemStats{}

	for scanner.Scan() {
		wg.Add(1)
		count++
		if count > 3500000 {
			fmt.Println("Done", count)
			fmt.Println(runtime.NumGoroutine(), "goroutines")
			break
		}
		if count%500000 == 0 {
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
		imported := &types.ImportedFenParent{}
		if err := json.Unmarshal(scanner.Bytes(), imported); err != nil {
			slog.Error("Error unmarshaling\n", err)
		}
		if db != nil {
			go func(imported *types.ImportedFenParent) {
				db.InsertEvalLines(ctx, imported)
				wg.Done()
			}(imported)
		} else {
			fenSlice = append(fenSlice, imported)
		}
	}
	elapsed := time.Since(start)
	fmt.Println(count, "fens processed.")
	fmt.Printf("To parse the file took [%v]\n", elapsed)
	wg.Wait()
	fmt.Println(runtime.NumGoroutine(), "goroutines")
	// Roughly 45 seconds
	time.Sleep(time.Second * 5)
	fmt.Println(runtime.NumGoroutine(), "goroutines")
	return fenSlice

}
