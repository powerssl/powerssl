package util

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
	"runtime/debug"
	"time"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/johanbrandhorst/certify"
	"github.com/johanbrandhorst/certify/issuers/vault"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type keyGeneratorFunc func() (crypto.PrivateKey, error)

func (kgf keyGeneratorFunc) Generate() (crypto.PrivateKey, error) {
	return kgf()
}

type Service interface {
	RegisterGRPCServer(baseServer *grpc.Server)
	ServiceName() string
}

func NewClientConn(addr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTimeout(time.Second),
	}
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		var creds credentials.TransportCredentials
		if insecureSkipTLSVerify {
			creds = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		} else {
			var err error
			if creds, err = credentials.NewClientTLSFromFile(certFile, serverNameOverride); err != nil {
				return nil, err
			}
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	return grpc.Dial(addr, opts...)
}

func ServeGRPC(ctx context.Context, addr, certFile, keyFile, caFile, commonName, vaultURL, vaultToken, vaultRole string, insecure bool, logger log.Logger, services []Service) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

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

	if !insecure {
		var creds credentials.TransportCredentials
		if certFile != "" && keyFile != "" {
			if creds, err = credentials.NewServerTLSFromFile(certFile, keyFile); err != nil {
				return fmt.Errorf("Failed to load TLS credentials %v", err)
			}
		} else {
			certPool := x509.NewCertPool()
			caData, err := ioutil.ReadFile(caFile)
			if err != nil {
				return err
			}
			certPool.AppendCertsFromPEM(caData)
			url, err := url.Parse(vaultURL)
			if err != nil {
				return err
			}
			c := &certify.Certify{
				Cache:      certify.NewMemCache(),
				CommonName: commonName,
				Issuer: &vault.Issuer{
					Role:  vaultRole,
					Token: vaultToken,
					URL:   url,
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
				hello.ServerName = commonName
				if cert, err = c.GetCertificate(hello); err != nil {
					logger.Log("err", err)
				}
				return cert, err
			}
			if _, err := getCertificate(&tls.ClientHelloInfo{ServerName: commonName}); err != nil {
				return err
			}
			creds = credentials.NewTLS(&tls.Config{GetCertificate: getCertificate})
		}
		options = append(options, grpc.Creds(creds))
	}

	srv := grpc.NewServer(options...)
	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(srv, healthSrv)
	for _, service := range services {
		service.RegisterGRPCServer(srv)
		healthSrv.SetServingStatus(service.ServiceName(), healthpb.HealthCheckResponse_SERVING)
	}

	c := make(chan error)
	go func() {
		c <- srv.Serve(listener)
		close(c)
	}()
	logger.Log("listening", addr, "secure", !insecure)
	select {
	case err := <-c:
		logger.Log("err", err)
		return err
	case <-ctx.Done():
		logger.Log("err", ctx.Err())
		srv.GracefulStop()
		return ctx.Err()
	}
}

func recoveryHandler(logger log.Logger) func(interface{}) error {
	return func(err interface{}) error {
		logger.Log("err", err)
		debug.PrintStack()
		return errors.New("Unknown error")

	}
}
