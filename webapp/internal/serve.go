package internal

import (
	"bytes"
	"context"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"powerssl.dev/common/log"

	"powerssl.dev/webapp/internal/asset"
)

func ServeHTTP(ctx context.Context, addr string, insecure bool, certFile, keyFile string, logger log.Logger, authURI, apiAddr, grpcWebURI string) error {
	var buffer []byte
	{
		tmpl, err := template.ParseFS(asset.Template, "index.html")
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err = tmpl.Execute(&buf, map[string]string{
			"APIAddr":    apiAddr,
			"AuthURI":    authURI,
			"GRPCWebURI": grpcWebURI,
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
		Addr:    addr,
		Handler: mux,
	}

	c := make(chan error)
	go func() {
		if insecure {
			c <- srv.ListenAndServe()
		} else {
			c <- srv.ListenAndServeTLS(certFile, keyFile)
		}
		close(c)
	}()
	logger.Infof("listening on %s", addr)
	select {
	case err := <-c:
		logger.Error(err)
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	case <-ctx.Done():
		logger.Error(ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}
