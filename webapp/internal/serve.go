package internal

import (
	"bytes"
	"context"
	"net/http"
	"time"

	bindatahtmltemplate "github.com/arschles/go-bindata-html-template"

	"powerssl.dev/backend/httpfs"
	"powerssl.dev/common/log"

	"powerssl.dev/webapp/internal/asset"
	"powerssl.dev/webapp/internal/template"
)

func ServeHTTP(ctx context.Context, addr string, insecure bool, certFile, keyFile string, logger log.Logger, authURI, apiAddr, grpcWebURI string) error {
	var buffer []byte
	{
		tmpl := bindatahtmltemplate.Must(bindatahtmltemplate.New("index", template.Asset).Parse("index.html"))
		data := map[string]interface{}{
			"APIAddr":    apiAddr,
			"AuthURI":    authURI,
			"GRPCWebURI": grpcWebURI,
		}
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return err
		}
		buffer = buf.Bytes()
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			_, _ = w.Write(buffer)
			return
		}
		if req.URL.Path == "/favicon.ico" {
			w.Header().Add("content-type", "image/x-icon")
			_, _ = w.Write(asset.MustAsset("favicon.ico"))
			return
		}
		http.NotFound(w, req)
	})
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(httpfs.NewFileSystem(asset.AssetFile()))))
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
