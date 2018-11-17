package util

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"runtime/debug"
	"time"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var ErrUnkown = errors.New("Unknown error")

type Service interface {
	RegisterGRPCServer(baseServer *grpc.Server)
}

func ServeGRPC(ctx context.Context, addr, certFile, keyFile string, insecure bool, logger log.Logger, services []Service) error {
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
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			return fmt.Errorf("Failed to load TLS credentials %v", err)
		}
		options = append(options, grpc.Creds(creds))
	}

	srv := grpc.NewServer(options...)
	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(srv, healthSrv)
	for _, service := range services {
		service.RegisterGRPCServer(srv)
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
		return ErrUnkown
	}
}

func NewClientConn(addr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTimeout(time.Second),
	}
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else if insecureSkipTLSVerify {
		creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds, err := credentials.NewClientTLSFromFile(certFile, serverNameOverride)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
