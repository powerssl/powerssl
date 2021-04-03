package transport // import "powerssl.dev/common/transport"

import (
	"context"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ServeMetrics(ctx context.Context, addr string, logger log.Logger) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	c := make(chan error)
	go func() {
		c <- srv.ListenAndServe()
		close(c)
	}()
	_ = logger.Log("listening", addr)
	select {
	case err := <-c:
		_ = logger.Log("err", err)
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	case <-ctx.Done():
		_ = logger.Log("err", ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}
