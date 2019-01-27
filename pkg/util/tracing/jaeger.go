package tracing

import (
	"io"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-lib/client/log/go-kit"
	jaegerprometheus "github.com/uber/jaeger-lib/metrics/prometheus"
)

func NewJaegerTracer(serviceName string, logger log.Logger) (opentracing.Tracer, io.Closer, error) {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return nil, nil, err
	}

	jaegerLogger := jaegerlog.NewLogger(logger)
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
