//go:generate gobin -m -run github.com/go-bindata/go-bindata/go-bindata -fs -modtime 726710400 -o openapi/bindata.go -pkg openapi -prefix ../../../api/openapi ../../../api/openapi/...
//go:generate gobin -m -run github.com/go-bindata/go-bindata/go-bindata -fs -modtime 726710400 -o swaggerui/bindata.go -pkg swaggerui -prefix ../../../web/swagger-ui ../../../web/swagger-ui

package grpcgateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	"powerssl.io/powerssl/internal/app/grpcgateway/openapi"
	"powerssl.io/powerssl/internal/app/grpcgateway/swaggerui"
	apiv1 "powerssl.io/powerssl/internal/pkg/apiserver/api/v1"
	"powerssl.io/powerssl/internal/pkg/transport"
	"powerssl.io/powerssl/internal/pkg/util"
)

type fileSystem struct {
	fs http.FileSystem
}

func (fs fileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, _ := f.Stat()
	if s.IsDir() {
		return nil, os.ErrNotExist
	}

	return f, nil
}

func Run(cfg *Config) {
	logger := util.NewLogger(os.Stdout)

	util.ValidateConfig(cfg, logger)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	var conn *grpc.ClientConn
	{
		var err error
		conn, err = transport.NewClientConn(ctx, cfg.APIServerClientConfig)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Connect", "err", err)
			os.Exit(1)
		}
	}

	if cfg.MetricsAddr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.MetricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return ServeHTTP(ctx, cfg.Addr, log.With(logger, "component", "http"), conn)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}

func ServeHTTP(ctx context.Context, addr string, logger log.Logger, conn *grpc.ClientConn) error {
	swaggerUIConfigHandler, err := swaggerUIConfigEndpoint()
	if err != nil {
		return err
	}

	gateway, err := newGateway(ctx, conn)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gateway)
	mux.Handle("/openapi/", http.StripPrefix("/openapi", http.FileServer(fileSystem{fs: openapi.AssetFile()})))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", http.FileServer(swaggerui.AssetFile())))
	mux.HandleFunc("/healthz", healthzServer(conn))
	mux.HandleFunc("/swagger-ui/config.json", swaggerUIConfigHandler)

	srv := &http.Server{
		Addr:    addr,
		Handler: allowCORS(mux),
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

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// TODO: Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// healthzServer returns a simple health handler which returns ok.
func healthzServer(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
			return
		}
		fmt.Fprintln(w, "ok")
	}
}

// newGateway returns a new gateway server which translates HTTP into gRPC.
func newGateway(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux := runtime.NewServeMux()

	for _, f := range []func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error{
		apiv1.RegisterACMEAccountServiceHandler,
		apiv1.RegisterACMEServerServiceHandler,
		apiv1.RegisterCertificateIssueServiceHandler,
		apiv1.RegisterCertificateServiceHandler,
		apiv1.RegisterUserServiceHandler,
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}

	return mux, nil
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// TODO: Don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

func swaggerUIConfigEndpoint() (func(w http.ResponseWriter, req *http.Request), error) {
	type url struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	}
	var config struct {
		URLs []url `json:"urls,omitempty"`
	}
	for _, assetName := range openapi.AssetNames() {
		config.URLs = append(config.URLs, url{Name: assetName, URL: fmt.Sprintf("/openapi/%s", assetName)})
	}
	res, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintln(w, string(res))
	}, nil
}