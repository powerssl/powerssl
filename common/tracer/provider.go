package tracer // import "powerssl.dev/common/tracer"

import (
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

func Provide(cfg Config, logger log.Logger) (opentracing.Tracer, func(), error) {
	logger = logger.With("component", "tracing")
	tracer, closer, err := New(cfg, logger)
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