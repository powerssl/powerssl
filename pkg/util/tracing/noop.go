package tracing

import (
	"io"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
)

func NewNoopTracer(serviceName string, logger log.Logger) (opentracing.Tracer, io.Closer, error) {
	return opentracing.NoopTracer{}, nil, nil
}
