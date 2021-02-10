package grpcgateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	"powerssl.dev/powerssl/internal/app/grpcgateway/openapi"
	"powerssl.dev/powerssl/internal/app/grpcgateway/swaggerui"
	apiv1 "powerssl.dev/powerssl/internal/pkg/apiserver/api/v1"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
	utilhttp "powerssl.dev/powerssl/internal/pkg/util/http"
)

func Run(cfg *Config) (err error) {
	_, logger := util.NewZapAndKitLogger()

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	var conn *grpc.ClientConn
	{
		if conn, err = transport.NewClientConn(ctx, &cfg.APIServerClientConfig); err != nil {
			return err
		}
	}

	if cfg.Metrics.Addr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.Metrics.Addr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return ServeHTTP(ctx, cfg.Addr, log.With(logger, "component", "http"), conn)
	})

	if err = g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			return err
		}
	}
	return nil
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
	mux.Handle("/openapi/", http.StripPrefix("/openapi", http.FileServer(utilhttp.NewFileSystem(openapi.AssetFile()))))
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
		_, _ = fmt.Fprintln(w, "ok")
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
func preflightHandler(w http.ResponseWriter, _ *http.Request) {
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
		_, _ = fmt.Fprintln(w, string(res))
	}, nil
}
