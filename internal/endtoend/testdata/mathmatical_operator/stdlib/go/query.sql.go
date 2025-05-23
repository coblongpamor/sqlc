// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const math = `-- name: Math :many
SELECT num, num / 1024 as division FROM foo
`

type MathRow struct {
	Num      int32
	Division int32
}

func (q *Queries) Math(ctx context.Context) ([]MathRow, error) {
	rows, err := q.db.QueryContext(ctx, math)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MathRow
	for rows.Next() {
		var i MathRow
		if err := rows.Scan(&i.Num, &i.Division); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
