package tracing // import "powerssl.dev/common/tracing"

import (
	"io"

	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerzap "github.com/uber/jaeger-client-go/log/zap"
	jaegerprometheus "github.com/uber/jaeger-lib/metrics/prometheus"

	"powerssl.dev/common/log"
)

func NewJaegerTracer(serviceName string, logger log.Logger) (opentracing.Tracer, io.Closer, error) {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return nil, nil, err
	}

	jaegerLogger := jaegerzap.NewLogger(logger.Desugar())
	jeagerMetricsFactory := jaegerprometheus.New()

	closer, err := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(jaegerLogger),
		jaegercfg.Metrics(jeagerMetricsFactory),
	)
	if err != nil {
		return nil, nil, err
	}
	return opentracing.GlobalTracer(), closer, nil
}
