package signer

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Run(grpcAddr, grpcCertFile, grpcKeyFile string, grpcInsecure bool, httpAddr string) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var duration metrics.Histogram
	{
		duration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "powerssl_io",
			Subsystem: "controller",
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds.",
		}, []string{"method", "success"})
	}

	var _ = duration // TODO

	var g group.Group
	{
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", grpcAddr, "secure", !grpcInsecure)
			options := []grpc.ServerOption{
				grpc.UnaryInterceptor(kitgrpc.Interceptor),
			}
			if !grpcInsecure {
				creds, err := credentials.NewServerTLSFromFile(grpcCertFile, grpcKeyFile)
				if err != nil {
					logger.Log("transport", "gRPC", "err", fmt.Errorf("Failed to load TLS credentials %v", err))
					os.Exit(1)
				}
				options = append(options, grpc.Creds(creds))
			}
			baseServer := grpc.NewServer(options...)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		if httpAddr != "" {
			http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
			httpListener, err := net.Listen("tcp", httpAddr)
			if err != nil {
				logger.Log("transport", "HTTP", "during", "Listen", "err", err)
				os.Exit(1)
			}
			g.Add(func() error {
				logger.Log("transport", "HTTP", "addr", httpAddr)
				return http.Serve(httpListener, nil)
			}, func(error) {
				httpListener.Close()
			})
		}
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}
