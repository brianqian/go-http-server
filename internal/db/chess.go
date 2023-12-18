package db

import (
	"base/types"
	"context"

	"github.com/jackc/pgx/v5"
)

var useFast = true

func (db *Database) InsertEvalLines(ctx context.Context, data *types.ImportedFenParent) {
	if useFast {

	} else {
		eval := data.Evals[0]
		query := `INSERT INTO fen_pv 
		(id, fen, created_at, updated_at, line, eval, knodes, depth) 
		VALUES
		(default, @fen, default, default, @line, @eval, @knodes, @depth);`
		var args []pgx.NamedArgs
		for _, pv := range eval.Pvs {
			args = append(args, pgx.NamedArgs{
				"fen":    data.Fen,
				"line":   pv.Line,
				"eval":   pv.Eval,
				"knodes": eval.Knodes,
				"depth":  eval.Depth,
			})
		}

		db.batchRequests(ctx, query, args)
	}

}

func (db *Database) InsertFenPosition(ctx context.Context, data *types.ImportedFenParent) {
	query := `INSERT INTO fen_positions(
		fen, created_at, updated_at, classic_eval, nnue_eval, final_eval, depth) 
	VALUES(@fen, default, default, @classic_eval, @nnue_eval, @final_eval, @depth);`

	args := []pgx.NamedArgs{{
		"fen":          data.Fen,
		"classic_eval": nil,
		"nnue_eval":    nil,
		"final_eval":   nil,
		"depth":        nil,
	}}

	db.BatchRequests(ctx, query, args)

}

func (db *Database) Tester(ctx context.Context) {
	// query := `SELECT * FROM fen_pv`

	query := `INSERT INTO fen_pv (id, fen, created_at, updated_at, line, eval, knodes, depth) VALUES (default, @fen, default, default, @line, @eval, @knodes, @depth);`

	args := pgx.NamedArgs{
		"fen":    "fdfdfdf",
		"line":   "fdfdfdf",
		"eval":   1,
		"knodes": 2,
		"depth":  3,
	}

	db.batchRequests(ctx, query, []pgx.NamedArgs{args})

}
