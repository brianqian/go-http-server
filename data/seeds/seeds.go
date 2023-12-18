package data

import (
	"base/internal/db"
	"context"
)

func SeedImportedFens(db *db.Database) {
	ctx := context.Background()

	db.Tester(ctx)
	// chess.ProcessFenPositions(ctx, db, "tmp/lichess_db_eval.json")

}
