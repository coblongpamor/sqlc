// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package querytest

import (
	"context"
)

type Querier interface {
	User(ctx context.Context) ([]int64, error)
	UsersB(ctx context.Context, id []int64) *UsersBBatchResults
	UsersC(ctx context.Context, id []int64) (int64, error)
}

var _ Querier = (*Queries)(nil)
