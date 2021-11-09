package tracer // import "powerssl.dev/common/tracer"

import (
	"io"

	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerzap "github.com/uber/jaeger-client-go/log/zap"
	jaegerprometheus "github.com/uber/jaeger-lib/metrics/prometheus"
	"go.uber.org/zap"
)

func NewJaegerTracer(cfg Config, logger *zap.SugaredLogger) (opentracing.Tracer, io.Closer, error) {
	config, err := jaegercfg.FromEnv()
	if err != nil {
		return nil, nil, err
	}

	jaegerLogger := jaegerzap.NewLogger(logger.Desugar())
	jeagerMetricsFactory := jaegerprometheus.New()

	closer, err := config.InitGlobalTracer(
		cfg.Component,
		jaegercfg.Logger(jaegerLogger),
		jaegercfg.Metrics(jeagerMetricsFactory),
	)
	if err != nil {
		return nil, nil, err
	}
	return opentracing.GlobalTracer(), closer, nil
}
