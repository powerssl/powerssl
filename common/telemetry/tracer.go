package telemetry

import (
	"context"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"powerssl.dev/common/log"
)

type Tracer struct {
	tracerProvider trace.TracerProvider
}

func NewTracer(cfg TracerConfig, logger log.Logger) *Tracer {
	opts := []tracesdk.TracerProviderOption{
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
	}
	logger.Debugw("initializing tracer provider")
	var tracerProvider trace.TracerProvider
	if cfg.Disabled {
		tracerProvider = trace.NewNoopTracerProvider()
	} else {
		tracerProvider = tracesdk.NewTracerProvider(opts...)
	}
	return &Tracer{
		tracerProvider: tracerProvider,
	}
}

func (c *Tracer) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return c.tracerProvider.Tracer(name, opts...)
}

func (c *Tracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return c.Tracer("").Start(ctx, spanName, opts...)
}
