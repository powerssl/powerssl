package server

import (
	"bytes"
	"context"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"powerssl.dev/webapp/internal/asset"
)

type Server struct {
	cfg    *Config
	logger *zap.SugaredLogger
}

func ServeHTTP(ctx context.Context, cfg *Config, logger *zap.SugaredLogger) error {
	return New(cfg, logger).ServeHTTP(ctx)
}

func New(cfg *Config, logger *zap.SugaredLogger) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) ServeHTTP(ctx context.Context) error {
	var buffer []byte
	{
		tmpl, err := template.ParseFS(asset.Template, "index.html")
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err = tmpl.Execute(&buf, map[string]string{
			"APIAddr":    s.cfg.APIAddr,
			"AuthURI":    s.cfg.AuthURI,
			"GRPCWebURI": s.cfg.GRPCWebURI,
		}); err != nil {
			return err
		}
		buffer = buf.Bytes()
	}
	mux := http.NewServeMux()
	mux.Handle("/css/", http.FileServer(http.FS(asset.CSS)))
	mux.Handle("/js/", http.FileServer(http.FS(asset.JS)))
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/favicon.ico" {
			writer.Header().Set("Content-Type", "image/x-ico")
			writer.Header().Set("Content-Length", strconv.Itoa(len(asset.Favicon)))
			_, err := writer.Write(asset.Favicon)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		} else {
			writer.Header().Set("Content-Type", "text/html")
			writer.Header().Set("Content-Length", strconv.Itoa(len(buffer)))
			_, err := writer.Write(buffer)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		}
	})
	srv := http.Server{
		Addr:    s.cfg.Addr,
		Handler: mux,
	}

	c := make(chan error)
	go func() {
		if s.cfg.Insecure {
			c <- srv.ListenAndServe()
		} else {
			c <- srv.ListenAndServeTLS(s.cfg.CertFile, s.cfg.KeyFile)
		}
		close(c)
	}()
	s.logger.Infof("listening on %s", s.cfg.Addr)
	select {
	case err := <-c:
		s.logger.Error(err)
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	case <-ctx.Done():
		s.logger.Error(ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}
