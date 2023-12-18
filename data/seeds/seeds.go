package data

import (
	"base/internal/chess"
	"base/internal/db"
	"context"
)

func SeedImportedFens(db *db.Database) {
	ctx := context.Background()
	chess.ProcessFenPositions(ctx, db, "tmp/lichess_db_eval.json")

}
