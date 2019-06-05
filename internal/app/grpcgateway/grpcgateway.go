package grpcgateway

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	apiv1 "powerssl.io/powerssl/internal/pkg/apiserver/api/v1"
	"powerssl.io/powerssl/internal/pkg/util"
)

// Endpoint describes a gRPC endpoint
type Endpoint struct {
	Addr                  string
	CertFile              string
	Insecure              bool
	InsecureSkipTLSVerify bool
	ServerNameOverride    string
}

// Options is a set of options to be passed to Run
type Options struct {
	// Addr is the address to listen
	Addr string

	// GRPCServer defines an endpoint of a gRPC service
	GRPCServer Endpoint

	// OpenapiDir is a path to a directory from which the server
	// serves openapi specs.
	OpenapiDir string

	// Mux is a list of options to be passed to the grpc-gateway multiplexer
	Mux []runtime.ServeMuxOption
}

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func Run(ctx context.Context, opts Options) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := util.NewClientConn(opts.GRPCServer.Addr, opts.GRPCServer.CertFile, opts.GRPCServer.ServerNameOverride, opts.GRPCServer.Insecure, opts.GRPCServer.InsecureSkipTLSVerify)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			glog.Errorf("Failed to close a client connection to the gRPC server: %v", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/openapi/", openapiServer(opts.OpenapiDir))
	mux.HandleFunc("/healthz", healthzServer(conn))

	gw, err := newGateway(ctx, conn, opts.Mux)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	s := &http.Server{
		Addr:    opts.Addr,
		Handler: allowCORS(mux),
	}
	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http server: %v", err)
		}
	}()

	glog.Infof("Starting listening at %s", opts.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
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
func newGateway(ctx context.Context, conn *grpc.ClientConn, opts []runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)

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

// openapiServer returns openapi specification files located under "/openapi/"
func openapiServer(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".yaml") {
			glog.Errorf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		glog.Infof("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/openapi/")
		p = path.Join(dir, p)
		http.ServeFile(w, r, p)
	}
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
}
