package data

import (
	"base/internal/chess"
	"base/internal/db"
	"context"
	"fmt"
	"strconv"
	"time"
)

func SeedImportedFens(db *db.Database) {
	ctx := context.Background()
	count := 0
	start := time.Now()

	for i := 0; i < 15; i++ {
		fmt.Println("Running ", i)
		chess.ImportFenSeed(ctx, db, "tmp/base-"+strconv.Itoa(count)+".json")
		count++
		time.Sleep(time.Second * 5)
	}

	elapsed := time.Since(start)

	fmt.Println("Total time: ", elapsed)
}
