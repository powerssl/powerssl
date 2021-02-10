package webapp

import (
	"bytes"
	"context"
	"net/http"
	"time"

	bindatahtmltemplate "github.com/arschles/go-bindata-html-template"
	"github.com/go-kit/kit/log"

	"powerssl.dev/powerssl/internal/app/webapp/asset"
	"powerssl.dev/powerssl/internal/app/webapp/template"
	utilhttp "powerssl.dev/powerssl/internal/pkg/util/http"
)

func ServeHTTP(ctx context.Context, addr string, logger log.Logger, authURI, apiAddr, grpcWebURI string) error {
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
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(utilhttp.NewFileSystem(asset.AssetFile()))))
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
