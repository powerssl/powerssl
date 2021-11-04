package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	New,
	ProvideDBTX,
)

type ConnString string

func ProvideDBTX(ctx context.Context, connString ConnString, logger *zap.SugaredLogger) (DBTX, func(), error) {
	logger = logger.With("component", "db")
	db, err := pgx.Connect(ctx, string(connString))
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
