package telemetry

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	Provide,
)

func Provide(ctx context.Context, cfg Config, logger *zap.SugaredLogger) (*Client, func(), error) {
	logger = logger.With("component", "telemetry")
	telemetry, err := New(cfg, logger)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err = telemetry.Cleanup(ctx); err != nil {
			logger.Error(err)
		}
	}
	return telemetry, cleanup, nil
}
