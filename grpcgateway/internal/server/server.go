package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/api/openapi"
	"powerssl.dev/backend/httpfs"
	"powerssl.dev/common/log"

	"powerssl.dev/grpcgateway/internal/swaggerui"
)

var serviceHandlers = []serviceHandler{
	apiv1.RegisterACMEAccountServiceHandler,
	apiv1.RegisterACMEServerServiceHandler,
	apiv1.RegisterCertificateIssueServiceHandler,
	apiv1.RegisterCertificateServiceHandler,
	apiv1.RegisterUserServiceHandler,
}

type serviceHandler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error

type Server struct {
	cfg    *Config
	logger *zap.SugaredLogger
	conn   *grpc.ClientConn
}

func ServeHTTP(ctx context.Context, cfg *Config, logger log.Logger, conn *grpc.ClientConn) error {
	return New(cfg, logger, conn).ServeHTTP(ctx)
}

func New(cfg *Config, logger log.Logger, conn *grpc.ClientConn) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		conn:   conn,
	}
}

func (s *Server) ServeHTTP(ctx context.Context) error {
	swaggerUIConfigHandler, err := swaggerUIConfigEndpoint()
	if err != nil {
		return err
	}

	gateway, err := s.newGateway(ctx)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gateway)
	mux.Handle("/openapi/", http.StripPrefix("/openapi", http.FileServer(httpfs.NewFileSystem(openapi.AssetFile()))))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", http.FileServer(swaggerui.AssetFile())))
	mux.HandleFunc("/healthz", s.healthzServer())
	mux.HandleFunc("/swagger-ui/config.json", swaggerUIConfigHandler)

	srv := &http.Server{
		Addr:    s.cfg.Addr,
		Handler: allowCORS(mux),
	}

	c := make(chan error)
	go func() {
		c <- srv.ListenAndServe()
		close(c)
	}()
	s.logger.Infof("listening on %s", s.cfg.Addr)
	select {
	case err = <-c:
		s.logger.Error(err)
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	case <-ctx.Done():
		s.logger.Error(ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err = srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}

// healthzServer returns a simple health handler which returns ok.
func (s *Server) healthzServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if s := s.conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
			return
		}
		_, _ = fmt.Fprintln(w, "ok")
	}
}

// newGateway returns a new gateway server which translates HTTP into gRPC.
func (s *Server) newGateway(ctx context.Context) (http.Handler, error) {
	mux := runtime.NewServeMux()

	for _, f := range serviceHandlers {
		if err := f(ctx, mux, s.conn); err != nil {
			return nil, err
		}
	}

	return mux, nil
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
