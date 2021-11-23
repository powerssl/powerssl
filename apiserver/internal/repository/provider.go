package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/jackc/pgx/v4"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	New,
	Provide,
)

func Provide(ctx context.Context, cfg Config, logger log.Logger) (DBTX, func(), error) {
	logger = logger.With("component", "db")
	db, err := pgx.Connect(ctx, cfg.ConnString)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err = db.Close(ctx); err != nil {
			logger.Error(err)
		}
	}
	return db, cleanup, nil
}
