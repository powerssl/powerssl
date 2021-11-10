package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	New,
	Provide,
)

func Provide(ctx context.Context, cfg Config, logger *zap.SugaredLogger) (DBTX, func(), error) {
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
