package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (q *Queries) Db() *pgx.Conn {
	return q.db.(*pgx.Conn)
}

func (q *Queries) NewTx(ctx context.Context) (*Queries, func(wErr *error), error) {
	tx, err := q.Db().Begin(ctx)
	if err != nil {
		return nil, nil, err
	}
	return &Queries{
			db: tx,
		}, func(wErr *error) {
			if wErr == nil {
				return
			}
			if err = tx.Rollback(ctx); err != nil && *wErr != nil {
				*wErr = fmt.Errorf("%s: %w", err, *wErr)
			} else {
				*wErr = err
			}
		}, nil
}

func (q *Queries) TxF(ctx context.Context, f func(queries *Queries) error) (err error) {
	var queries *Queries
	var rollback func(*error)
	if queries, rollback, err = q.NewTx(ctx); err != nil {
		return err
	}
	defer rollback(&err)
	if err = f(queries); err != nil {
		return err
	}
	return queries.Tx().Commit(ctx)
}

func (q *Queries) Tx() pgx.Tx {
	return q.db.(pgx.Tx)
}
