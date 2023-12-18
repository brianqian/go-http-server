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
	"sync"
	"time"
)

func ProcessFenPositions(ctx context.Context, db *db.Database, filepath string) []*types.ImportedFenParent {
	f, err := os.Open(filepath)

	if err != nil {
		slog.Error("Invalid filepath for import")
	}
	defer f.Close()

	start := time.Now()
	fenSlice := []*types.ImportedFenParent{}
	scanner := bufio.NewScanner(f)
	wg := sync.WaitGroup{}

	for scanner.Scan() {
		wg.Add(1)
		imported := &types.ImportedFenParent{}
		if err := json.Unmarshal(scanner.Bytes(), imported); err != nil {
			slog.Error("Error unmarshaling\n", err)
		}
		if db != nil {
			go func(imported *types.ImportedFenParent) {
				defer wg.Done()
				db.InsertEvalLines(ctx, imported)

			}(imported)

		} else {
			fenSlice = append(fenSlice, imported)
		}
	}
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println(len(fenSlice), "fens processed.")
	fmt.Printf("To parse the file took [%v]\n", elapsed)
	// Roughly 45 seconds

	return fenSlice

}
