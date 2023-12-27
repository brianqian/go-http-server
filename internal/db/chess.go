package db

import (
	"base/internal/util"
	"base/types"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type FenEvaluation struct {
	Id        pgtype.UUID        `json:"id" db:"id"`
	Fen       string             `json:"fen" db:"fen"`
	CreatedAt pgtype.Timestamptz `json:"createdAt" db:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt" db:"updated_at"`
	Line      string             `json:"line" db:"line"`
	Eval      int                `json:"eval" db:"eval"`
	Knodes    int                `json:"knodes" db:"knodes"`
	Depth     int                `json:"depth" db:"depth"`
}

func (db *Database) InsertEvalLines(ctx context.Context, data []*types.ImportedFenParent) []error {
	tableCols := "fen_pv (id, fen, created_at, updated_at, line, eval, knodes, depth, mate)"
	var values [][]string

	for _, val := range data {
		for _, eval := range val.Evals {
			for _, pv := range eval.Pvs {
				fn := util.WrapInSingleQuotes
				s := []string{"default", fn(val.Fen), "default", "default", fn(pv.Line), fn(pv.Eval), fn(eval.Knodes), fn(eval.Depth), fn(pv.Mate)}
				values = append(values, s)
			}
		}
	}
	fmt.Println(len(data), "items converted into lines: ", len(values))
	errors := db.insertMany(ctx, tableCols, values, false)

	return errors
}

func (db *Database) Old_InsertEvalLines(ctx context.Context, data *types.ImportedFenParent) {
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

	db.batchRequests(ctx, query, args)

}

func (db *Database) Tester(ctx context.Context) {

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
