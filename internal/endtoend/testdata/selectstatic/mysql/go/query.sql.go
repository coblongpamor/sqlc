// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const selectStatic = `-- name: SelectStatic :one
SELECT 'a', 'b' AS b, 1 AS num, true AS truefield, 1.0 AS floater
`

type SelectStaticRow struct {
	Column1   string
	B         string
	Num       int32
	Truefield int32
	Floater   float64
}

func (q *Queries) SelectStatic(ctx context.Context) (SelectStaticRow, error) {
	row := q.db.QueryRowContext(ctx, selectStatic)
	var i SelectStaticRow
	err := row.Scan(
		&i.Column1,
		&i.B,
		&i.Num,
		&i.Truefield,
		&i.Floater,
	)
	return i, err
}
