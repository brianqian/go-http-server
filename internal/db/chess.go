package db

import (
	"base/internal/util"
	"base/types"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

type FenEvaluation struct {
	Id        pgtype.UUID        `json:"id" db:"id"`
	Fen       pgtype.Text        `json:"fen" db:"fen"`
	CreatedAt pgtype.Timestamptz `json:"createdAt" db:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt" db:"updated_at"`
	Line      pgtype.Text        `json:"line" db:"line"`
	Eval      pgtype.Int4        `json:"eval" db:"eval"`
	Knodes    pgtype.Int4        `json:"knodes" db:"knodes"`
	Depth     pgtype.Int4        `json:"depth" db:"depth"`
	Mate      pgtype.Int4        `json:"mate" db:"mate"`
}
type InsertFenEvaluation struct {
	Fen    string `json:"fen" db:"fen"`
	Line   string `json:"line" db:"line"`
	Eval   int    `json:"eval" db:"eval"`
	Knodes int    `json:"knodes" db:"knodes"`
	Depth  int    `json:"depth" db:"depth"`
	Mate   int    `json:"mate" db:"mate"`
}

func (db *Database) InsertEvalLines(ctx context.Context, data []*types.ImportedFenParent) (errors []error) {

	table := pgx.Identifier{"fen_pv"}
	cols := "(id, created_at, updated_at, fen, line, eval, knodes, depth, mate)"
	fenEvals := convertImportLineToEvals(data)

	if len(fenEvals) > 5 {
		conn, err := db.conn.Acquire(ctx)
		if err != nil {
			slog.Error("Error getting db connection")
		}
		defer conn.Release()
		asAnyArr := lo.Map[InsertFenEvaluation, []any](fenEvals, func(item InsertFenEvaluation, index int) []any {
			return []any{
				item.Fen, item.Line, item.Eval, item.Knodes, item.Depth, item.Mate,
			}
		})
		_, err = conn.CopyFrom(ctx, table, []string{"fen", "line", "eval", "knodes", "depth", "mate"}, pgx.CopyFromRows(asAnyArr))
		if err != nil {
			errors = []error{err}
		}
	} else {
		asStringArr := lo.Map[InsertFenEvaluation, []string](fenEvals, func(item InsertFenEvaluation, index int) []string {
			fn := util.WrapInSingleQuotes
			return []string{
				"default", "default", "default",
				fn(item.Fen), fn(item.Line), fn(item.Eval), fn(item.Knodes), fn(item.Depth), fn(item.Mate),
			}
		})
		errors = db.insertMany(ctx, table, cols, asStringArr, false)
	}
	fmt.Println(len(data), "items converted into lines: ", len(fenEvals))

	return errors
}

func (db *Database) ImportEvalFromFile(ctx context.Context, filepath string) {
	shape := &FenEvaluation{}
	db.copyFromFile(copyFromFileArgs{
		ctx:      ctx,
		table:    "fen_pv",
		filePath: filepath,
		columns:  []string{"fen", "line", "eval", "knodes", "depth", "mate"},
		parsingFn: func(data []byte) ([]any, error) {
			err := json.Unmarshal(data, shape)
			parsed := []any{shape.Fen, shape}
			return parsed, err
		},
	})
}

// func (db *Database) InsertFenPosition(ctx context.Context, data *types.ImportedFenParent) {
// 	query := `INSERT INTO fen_positions(
// 		fen, created_at, updated_at, classic_eval, nnue_eval, final_eval, depth)
// 	VALUES(@fen, default, default, @classic_eval, @nnue_eval, @final_eval, @depth);`

// 	args := []pgx.NamedArgs{{
// 		"fen":          data.Fen,
// 		"classic_eval": nil,
// 		"nnue_eval":    nil,
// 		"final_eval":   nil,
// 		"depth":        nil,
// 	}}

// 	db.batchRequests(ctx, query, args)

// }

func convertImportLineToEvals(data []*types.ImportedFenParent) []InsertFenEvaluation {
	var values []InsertFenEvaluation
	for _, val := range data {
		for _, eval := range val.Evals {
			for _, pv := range eval.Pvs {
				values = append(values, InsertFenEvaluation{
					Fen:    val.Fen,
					Line:   pv.Line,
					Eval:   pv.Eval,
					Knodes: eval.Knodes,
					Depth:  eval.Depth,
					Mate:   pv.Mate,
				})
			}
		}
	}
	return values
}

/*DEPRECATED*/
// func (db *Database) Old_InsertEvalLines(ctx context.Context, data *types.ImportedFenParent) {
// 	eval := data.Evals[0]
// 	query := `INSERT INTO fen_pv
// 		(id, fen, created_at, updated_at, line, eval, knodes, depth)
// 		VALUES
// 		(default, @fen, default, default, @line, @eval, @knodes, @depth);`
// 	var args []pgx.NamedArgs
// 	for _, pv := range eval.Pvs {
// 		args = append(args, pgx.NamedArgs{
// 			"fen":    data.Fen,
// 			"line":   pv.Line,
// 			"eval":   pv.Eval,
// 			"knodes": eval.Knodes,
// 			"depth":  eval.Depth,
// 		})
// 	}
// 	db.batchRequests(ctx, query, args)
// }

// func (db *Database) Tester(ctx context.Context) {
// 	query := `INSERT INTO fen_pv (id, fen, created_at, updated_at, line, eval, knodes, depth) VALUES (default, @fen, default, default, @line, @eval, @knodes, @depth);`
// 	args := pgx.NamedArgs{
// 		"fen":    "fdfdfdf",
// 		"line":   "fdfdfdf",
// 		"eval":   1,
// 		"knodes": 2,
// 		"depth":  3,
// 	}
// 	db.batchRequests(ctx, query, []pgx.NamedArgs{args})
// }
