package tracing // import "powerssl.dev/common/tracing"

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"

	"powerssl.dev/common/log"
)

func Init(serviceName string, implementation string, logger log.Logger) (opentracing.Tracer, io.Closer, error) {
	switch implementation {
	case "":
		return NewNoopTracer(serviceName, logger)
	case "jaeger":
		return NewJaegerTracer(serviceName, logger)
	default:
		return nil, nil, fmt.Errorf("tracing implementation does not exist: %s", implementation)
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
