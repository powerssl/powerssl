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
			if err = tx.Rollback(ctx); err != nil && *wErr != nil {
				*wErr = fmt.Errorf("%s: %w", err, *wErr)
			} else {
				*wErr = err
			}
		}, nil
}

func (q *Queries) Tx() pgx.Tx {
	return q.db.(pgx.Tx)
}
