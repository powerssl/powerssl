package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/sdk/trace"
	otrace "go.opentelemetry.io/otel/trace"

	"powerssl.dev/common/log"
)

type Client struct {
	DefaultTracer  otrace.Tracer
	cfg            Config
	log            log.Logger
	tracerProvider *trace.TracerProvider
}

func New(cfg Config, logger log.Logger) (*Client, error) {
	opts := []trace.TracerProviderOption{
		trace.WithSampler(trace.AlwaysSample()),
	}
	tracerProvider := trace.NewTracerProvider(opts...)
	defaultTracer := tracerProvider.Tracer("")
	return &Client{
		DefaultTracer:  defaultTracer,
		tracerProvider: tracerProvider,
		cfg:            cfg,
		log:            logger,
	}, nil
}

func (c *Client) Cleanup(ctx context.Context) error {
	err := c.tracerProvider.ForceFlush(ctx)
	if err != nil {
		return err
	}
	return c.tracerProvider.Shutdown(ctx)
}

func (c *Client) Tracer(name string, opts ...otrace.TracerOption) otrace.Tracer {
	return c.tracerProvider.Tracer(name, opts...)
}

func (c *Client) Start(ctx context.Context, spanName string, opts ...otrace.SpanStartOption) (context.Context, otrace.Span) {
	return c.DefaultTracer.Start(ctx, spanName, opts...)
}
