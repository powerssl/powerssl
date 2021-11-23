package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"net"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"powerssl.dev/common/errutil"
	"powerssl.dev/common/log"
)

type Server struct {
	cfg    Config
	health *health.Server
	logger log.Logger
	grpc   *grpc.Server
}

func New(cfg Config, logger log.Logger) (*Server, error) {
	recoveryOptions := []recovery.Option{
		recovery.WithRecoveryHandler(recoveryHandler(logger)),
	}
	options := []grpc.ServerOption{
		middleware.WithUnaryServerChain(
			recovery.UnaryServerInterceptor(recoveryOptions...),
		),
		middleware.WithStreamServerChain(
			recovery.StreamServerInterceptor(recoveryOptions...),
		),
	}

	if !cfg.Insecure {
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			creds, err := credentials.NewServerTLSFromFile(cfg.CertFile, cfg.KeyFile)
			if err != nil {
				return nil, fmt.Errorf("failed to load TLS credentials %v", err)
			}
			options = append(options, grpc.Creds(creds))
		} else {
			c, err := makeCertify(cfg, logger)
			if err != nil {
				return nil, err
			}
			options = append(options, grpc.Creds(credentials.NewTLS(c)))
		}
	}

	srv := grpc.NewServer(options...)
	reflection.Register(srv)
	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(srv, healthSrv)
	return &Server{
		health: healthSrv,
		cfg:    cfg,
		logger: logger,
		grpc:   srv,
	}, nil
}

func (s *Server) Serve(ctx context.Context) (err error) {
	var listener net.Listener
	if listener, err = net.Listen("tcp", s.cfg.Addr); err != nil {
		return err
	}
	defer errutil.ErrWrapCloser(listener, &err)
	c := make(chan error)
	go func() {
		c <- s.grpc.Serve(listener)
		close(c)
	}()
	s.logger.With("secure", !s.cfg.Insecure).Infof("listening on %s", s.cfg.Addr)
	select {
	case err = <-c:
		s.logger.Error(err)
		return err
	case <-ctx.Done():
		s.logger.Error(ctx.Err())
		s.grpc.GracefulStop()
		return ctx.Err()
	}
}

func (s *Server) RegisterService(sd *grpc.ServiceDesc, ss interface{}) {
	s.grpc.RegisterService(sd, ss)
	s.health.SetServingStatus(sd.ServiceName, healthpb.HealthCheckResponse_SERVING)
}

func recoveryHandler(logger log.Logger) func(interface{}) error {
	return func(err interface{}) error {
		logger.With(zap.Stack("stack")).Errorf("%s", err)
		return errors.New("unknown error")
	}
}
