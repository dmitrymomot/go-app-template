// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package repository

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)