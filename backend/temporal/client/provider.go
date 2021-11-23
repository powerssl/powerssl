package client

import (
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"go.temporal.io/sdk/client"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

func Provide(cfg Config, logger log.Logger, tracer opentracing.Tracer) (client.Client, func(), error) {
	logger = logger.With("component", "temporal")
	c, closer, err := New(cfg, logger, tracer)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err = closer.Close(); err != nil {
			logger.Error(err)
		}
	}
	return c, cleanup, nil
}
