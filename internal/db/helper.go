package db

import (
	"bufio"
	"strings"

	"github.com/jackc/pgx/v5"
)

func createBatchedQueryString(table pgx.Identifier, target string, items [][]string, errorOnConflict bool) string {
	var query = &strings.Builder{}
	query.WriteString("INSERT INTO ")
	query.WriteString(table.Sanitize())
	query.WriteString(target)
	query.WriteString(" VALUES ")
	for idx, item := range items {
		query.WriteRune('(')
		query.WriteString(strings.Join(item, ", "))
		if idx == len(items)-1 {
			query.WriteString(")")
		} else {
			query.WriteString("), ")
		}

	}
	if !errorOnConflict {
		query.WriteString(" ON CONFLICT DO NOTHING")
	}
	query.WriteRune(';')

	return query.String()
}

func copyFromFileSource(sc *bufio.Scanner, fn parsingFn) pgx.CopyFromSource {
	return &copyFromFile{scanner: sc, parsingFn: fn}
}

type parsingFn func([]byte) ([]any, error)

type copyFromFile struct {
	idx       int
	rowStore  [][]any
	scanner   *bufio.Scanner
	parsingFn parsingFn
}

func (cff *copyFromFile) Next() bool {
	if cff.idx == len(cff.rowStore)-1 {
		cff.idx = 0
		clear(cff.rowStore)
		return cff.scanner.Scan()
	} else {
		cff.idx++
		return true
	}

}
func (cff *copyFromFile) Values() ([]any, error) {
	return cff.parsingFn(cff.scanner.Bytes())
}
func (cff *copyFromFile) Err() error {
	return cff.scanner.Err()
}

func (cff *copyFromFile) GetRows() bool {
	return true
}
