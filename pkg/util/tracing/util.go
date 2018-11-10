package tracing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
)

func Init(serviceName, implementation string, logger log.Logger) (opentracing.Tracer, io.Closer, error) {
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

func TextMapCarrierFromSpan(span opentracing.Span) opentracing.TextMapCarrier {
	textMapCarrier := opentracing.TextMapCarrier{}
	span.Tracer().Inject(span.Context(), opentracing.TextMap, textMapCarrier)
	return textMapCarrier
}

func JSONCarrierFromSpan(span opentracing.Span) (string, error) {
	textMapCarrier := TextMapCarrierFromSpan(span)
	bytes, err := json.Marshal(textMapCarrier)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
