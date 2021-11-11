package metrics // import "powerssl.dev/common/transport"

import (
	"context"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func ServeMetrics(ctx context.Context, cfg Config, logger *zap.SugaredLogger) error {
	return NewMetrics(cfg, logger).Serve(ctx)
}

type Config struct {
	Addr string `flag:"addr;;;metrics addr"`
}

type Metrics struct {
	*http.Server
	logger *zap.SugaredLogger
}

func NewMetrics(cfg Config, logger *zap.SugaredLogger) *Metrics {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return &Metrics{
		Server: &http.Server{
			Addr:    cfg.Addr,
			Handler: mux,
		},
		logger: logger,
	}
}

func (m *Metrics) Serve(ctx context.Context) error {
	c := make(chan error)
	go func() {
		c <- m.ListenAndServe()
		close(c)
	}()
	m.logger.Infof("listening on %s", m.Addr)
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
		if err := m.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}