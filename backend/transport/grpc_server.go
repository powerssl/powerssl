package transport // import "powerssl.dev/backend/transport"

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net"
	"net/url"
	"time"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-playground/validator/v10"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/vault"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"powerssl.dev/common"
	"powerssl.dev/common/log"
)

type ServerConfig struct {
	Addr       string
	CAFile     string `mapstructure:"ca-file"`
	CertFile   string `mapstructure:"cert-file"`
	CommonName string `mapstructure:"common-name"`
	Insecure   bool
	KeyFile    string `mapstructure:"key-file"`
	VaultRole  string
	VaultToken string
	VaultURL   string
}

func ServerConfigValidator(sl validator.StructLevel) {
	cfg := sl.Current().Interface().(ServerConfig)

	if !cfg.Insecure && cfg.CommonName == "" && (cfg.CertFile == "" || cfg.KeyFile == "") {
		sl.ReportError(cfg.CommonName, "CommonName", "CommonName", "required", "CommonName or CertFile and KeyFile required")
		sl.ReportError(cfg.CertFile, "CertFile", "CertFile", "required", "CommonName or CertFile and KeyFile required")
		sl.ReportError(cfg.KeyFile, "KeyFile", "KeyFile", "required", "CommonName or CertFile and KeyFile required")
	}
}

type keyGeneratorFunc func() (crypto.PrivateKey, error)

func (kgf keyGeneratorFunc) Generate() (crypto.PrivateKey, error) {
	return kgf()
}

type Service struct {
	ServiceName        string
	RegisterGRPCServer func(baseServer *grpc.Server)
}

type Services []Service

func ServeGRPC(ctx context.Context, cfg ServerConfig, logger log.Logger, services Services) (err error) {
	var listener net.Listener
	if listener, err = net.Listen("tcp", cfg.Addr); err != nil {
		return err
	}
	defer common.ErrWrapCloser(listener, &err)

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
				return fmt.Errorf("failed to load TLS credentials %v", err)
			}
		} else {
			certPool := x509.NewCertPool()
			var caData []byte
			if caData, err = ioutil.ReadFile(cfg.CAFile); err != nil {
				return err
			}
			certPool.AppendCertsFromPEM(caData)
			var vaultURL *url.URL
			if vaultURL, err = url.Parse(cfg.VaultURL); err != nil {
				return err
			}
			c := &certify.Certify{
				Cache:      certify.NewMemCache(),
				CommonName: cfg.CommonName,
				Issuer: &vault.Issuer{
					URL:   vaultURL,
					Role:  cfg.VaultRole,
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
	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(srv, healthSrv)
	for _, service := range services {
		service.RegisterGRPCServer(srv)
		healthSrv.SetServingStatus(service.ServiceName, healthpb.HealthCheckResponse_SERVING)
	}

	c := make(chan error)
	go func() {
		c <- srv.Serve(listener)
		close(c)
	}()
	logger.With("secure", !cfg.Insecure).Infof("listening on %s", cfg.Addr)
	select {
	case err = <-c:
		logger.Error(err)
		return err
	case <-ctx.Done():
		logger.Error(ctx.Err())
		srv.GracefulStop()
		return ctx.Err()
	}
}

func recoveryHandler(logger log.Logger) func(interface{}) error {
	return func(err interface{}) error {
		logger.With(zap.Stack("stack")).Errorf("%s", err)
		return errors.New("unknown error")

	}
}
