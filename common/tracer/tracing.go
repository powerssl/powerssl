package tracer // import "powerssl.dev/common/tracer"

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func New(cfg Config, logger *zap.SugaredLogger) (opentracing.Tracer, io.Closer, error) {
	switch cfg.Implementation {
	case "":
		return NewNoopTracer(cfg, logger)
	case "jaeger":
		return NewJaegerTracer(cfg, logger)
	default:
		return nil, nil, fmt.Errorf("tracing implementation does not exist: %s", cfg.Implementation)
	}
}

func ContextWithSpanFromContext(ctx context.Context, spanCtx context.Context) context.Context {
	span := opentracing.SpanFromContext(spanCtx)
	return opentracing.ContextWithSpan(ctx, span)
}

func JSONCarrierFromSpan(span opentracing.Span) (string, error) {
	textMapCarrier, err := TextMapCarrierFromSpan(span)
	var bytes []byte
	if bytes, err = json.Marshal(textMapCarrier); err != nil {
		return "", err
	}
	return string(bytes), nil
}

func TextMapCarrierFromSpan(span opentracing.Span) (opentracing.TextMapCarrier, error) {
	textMapCarrier := opentracing.TextMapCarrier{}
	if err := span.Tracer().Inject(span.Context(), opentracing.TextMap, textMapCarrier); err != nil {
		return nil, err
	}
	return textMapCarrier, nil
}

func WireContextFromJSON(s string) (opentracing.SpanContext, error) {
	var tmc opentracing.TextMapCarrier
	if err := json.Unmarshal([]byte(s), &tmc); err != nil {
		return nil, err
	}
	wireContext, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, tmc)
	if err != nil {
		return nil, err
	}
	return wireContext, nil
}
