package telemetry

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/export/metric/aggregation"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"

	"powerssl.dev/common/log"
)

type Meter struct {
	logger        log.Logger
	meterProvider metric.MeterProvider
	server        *http.Server
}

func NewMeter(cfg MeterConfig, logger log.Logger) (*Meter, error) {
	config := prometheus.Config{}
	c := controller.New(
		processor.NewFactory(
			selector.NewWithHistogramDistribution(
				histogram.WithExplicitBoundaries(config.DefaultHistogramBoundaries),
			),
			aggregation.CumulativeTemporalitySelector(),
			processor.WithMemory(true),
		),
	)
	exporter, err := prometheus.New(config, c)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize prometheus exporter: %w", err)
	}

	return &Meter{
		logger:        logger,
		meterProvider: exporter.MeterProvider(),
		server:        newServer(cfg.Addr, exporter),
	}, nil
}

func (m *Meter) Meter(instrumentationName string, opts ...metric.MeterOption) metric.Meter {
	return m.meterProvider.Meter(instrumentationName, opts...)
}

func (m *Meter) Serve(ctx context.Context) error {
	c := make(chan error)
	go func() {
		c <- m.server.ListenAndServe()
		close(c)
	}()
	m.logger.Infof("listening on %s", m.server.Addr)
	select {
	case err := <-c:
		m.logger.Error(err)
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	case <-ctx.Done():
		m.logger.Error(ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := m.server.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}

func newServer(addr string, exporter *prometheus.Exporter) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/metrics", exporter.ServeHTTP)
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
