package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type sqlxInterface struct {
	db *sqlx.DB
}

func (r *sqlxInterface) ExecerContext(ctx context.Context) sqlx.ExecerContext {
	return r.sqlInterface(ctx).(sqlx.ExecerContext)
}

func (r *sqlxInterface) QueryerContext(ctx context.Context) sqlx.QueryerContext {
	return r.sqlInterface(ctx).(sqlx.QueryerContext)
}

func (r *sqlxInterface) sqlInterface(ctx context.Context) interface{} {
	tx, ok := ctx.Value(transactionValue).(*sqlx.Tx)
	if ok {
		return tx
	}
	return r.db
}
