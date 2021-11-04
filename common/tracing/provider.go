package tracing

import (
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	ProvideTracer,
)

type TracerImplementation string
type TracerComponent string

func ProvideTracer(implementation TracerImplementation, component TracerComponent, logger *zap.SugaredLogger) (opentracing.Tracer, func(), error) {
	logger = logger.With("component", "tracing")
	tracer, closer, err := Init(string(component), string(implementation), logger)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err = closer.Close(); err != nil {
			logger.Error(err)
		}
	}
	return tracer, cleanup, nil
}