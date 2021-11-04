package client

import (
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	ProvideTemporalClient,
)

type TemporalClientComponent string

func ProvideTemporalClient(cfg Config, logger *zap.SugaredLogger, tracer opentracing.Tracer, component TemporalClientComponent) (client.Client, func(), error) {
	logger = logger.With("component", "temporal")
	c, closer, err := NewClient(cfg, logger, tracer, string(component))
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
