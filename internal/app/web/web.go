//go:generate go-bindata -pkg web -prefix ../../../web/app ../../../web/app

package web

import (
	"context"
	htmltemplate "html/template"
	"net/http"
	"os"
	"time"

	"github.com/arschles/go-bindata-html-template"
	"github.com/go-kit/kit/log"
	"golang.org/x/sync/errgroup"

	"powerssl.io/pkg/util"
)

func Run(httpAddr, metricsAddr, authURI string) {
	logger := util.NewLogger(os.Stdout)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	if metricsAddr != "" {
		g.Go(func() error {
			return util.ServeMetrics(ctx, metricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return ServeHTTP(ctx, httpAddr, log.With(logger, "component", "http"), authURI)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}

func ServeHTTP(ctx context.Context, addr string, logger log.Logger, authURI string) error {
	mux := http.NewServeMux()
	tmpl, err := template.New("index", Asset).Parse("index.html")
	if err != nil {
		return err
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if err := tmpl.Execute(w, map[string]interface{}{"AuthURI": htmltemplate.URL(authURI)}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	c := make(chan error)
	go func() {
		c <- srv.ListenAndServe()
		close(c)
	}()
	logger.Log("listening", addr)
	select {
	case err := <-c:
		logger.Log("err", err)
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	case <-ctx.Done():
		logger.Log("err", ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}