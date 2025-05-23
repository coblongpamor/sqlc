// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
	"time"
)

const listAuthors = `-- name: ListAuthors :many
SELECT   id, name as name, bio
FROM     authors
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
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

const listAuthorsIdenticalAlias = `-- name: ListAuthorsIdenticalAlias :many
SELECT   id, name as name, bio
FROM     authors
`

func (q *Queries) ListAuthorsIdenticalAlias(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsIdenticalAlias)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
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

const listMetrics = `-- name: ListMetrics :many
SELECT date_trunc('day', time) AS bucket, city_name, AVG(temp_c)
FROM weather_metrics
WHERE time > NOW() - (6 * INTERVAL '1 month')
GROUP BY bucket, city_name
ORDER BY bucket DESC
`

type ListMetricsRow struct {
	Bucket   time.Time
	CityName sql.NullString
	Avg      float64
}

func (q *Queries) ListMetrics(ctx context.Context) ([]ListMetricsRow, error) {
	rows, err := q.db.QueryContext(ctx, listMetrics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMetricsRow
	for rows.Next() {
		var i ListMetricsRow
		if err := rows.Scan(&i.Bucket, &i.CityName, &i.Avg); err != nil {
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
