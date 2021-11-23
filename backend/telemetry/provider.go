package telemetry

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

func Provide(ctx context.Context, cfg Config, logger log.Logger) (*Client, func(), error) {
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
