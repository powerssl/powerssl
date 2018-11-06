package tracing

import (
	"context"
	"encoding/json"

	"github.com/opentracing/opentracing-go"
)

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
