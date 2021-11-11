package transport

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"time"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/vault"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	error2 "powerssl.dev/common/error"
)

type Server struct {
	cfg    Config
	health *health.Server
	logger *zap.SugaredLogger
	grpc   *grpc.Server
}

func New(cfg Config, logger *zap.SugaredLogger) (_ *Server, err error) {
	logger = logger.With("transport", "gRPC")
	recoveryOptions := []recovery.Option{
		recovery.WithRecoveryHandler(recoveryHandler(logger)),
	}
	options := []grpc.ServerOption{
		middleware.WithUnaryServerChain(
			kitgrpc.Interceptor,
			recovery.UnaryServerInterceptor(recoveryOptions...),
		),
		middleware.WithStreamServerChain(
			recovery.StreamServerInterceptor(recoveryOptions...),
		),
	}

	if !cfg.Insecure {
		var creds credentials.TransportCredentials
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			if creds, err = credentials.NewServerTLSFromFile(cfg.CertFile, cfg.KeyFile); err != nil {
				return nil, fmt.Errorf("failed to load TLS credentials %v", err)
			}
		} else {
			certPool := x509.NewCertPool()
			var caData []byte
			if caData, err = ioutil.ReadFile(cfg.CAFile); err != nil {
				return nil, err
			}
			certPool.AppendCertsFromPEM(caData)
			var vaultURL *url.URL
			if vaultURL, err = url.Parse(cfg.VaultURL); err != nil {
				return nil, err
			}
			c := &certify.Certify{
				Cache:      certify.NewMemCache(),
				CommonName: cfg.CommonName,
				Issuer: &vault.Issuer{
					URL:        vaultURL,
					Role:       cfg.VaultRole,
					AuthMethod: vault.ConstantToken(cfg.VaultToken),
					TLSConfig: &tls.Config{
						RootCAs: certPool,
					},
				},
				RenewBefore: time.Hour,
				CertConfig: &certify.CertConfig{
					KeyGenerator: keyGeneratorFunc(func() (crypto.PrivateKey, error) {
						return rsa.GenerateKey(rand.Reader, 2048)
					}),
				},
			}
			getCertificate := func(hello *tls.ClientHelloInfo) (cert *tls.Certificate, err error) {
				// TODO: ???
				hello.ServerName = cfg.CommonName
				if cert, err = c.GetCertificate(hello); err != nil {
					logger.Error(err)
				}
				return cert, err
			}
			// TODO: This was priming the server before the first request. Certify is now failing with this.
			// if _, err := getCertificate(&tls.ClientHelloInfo{ServerName: cfg.CommonName}); err != nil {
			// 	return err
			// }
			creds = credentials.NewTLS(&tls.Config{GetCertificate: getCertificate})
		}
		options = append(options, grpc.Creds(creds))
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
	defer error2.ErrWrapCloser(listener, &err)
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

type keyGeneratorFunc func() (crypto.PrivateKey, error)

func (kgf keyGeneratorFunc) Generate() (crypto.PrivateKey, error) {
	return kgf()
}

func recoveryHandler(logger *zap.SugaredLogger) func(interface{}) error {
	return func(err interface{}) error {
		logger.With(zap.Stack("stack")).Errorf("%s", err)
		return errors.New("unknown error")

	}
}
