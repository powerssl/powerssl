package _interface

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type SQLX interface {
	ExecerContext(ctx context.Context) sqlx.ExecerContext
	QueryerContext(ctx context.Context) sqlx.QueryerContext
}
