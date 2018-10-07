package apiserver

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	resource "powerssl.io/pkg/resource"
	"powerssl.io/pkg/resource/generated/certificate"
	"powerssl.io/pkg/resource/generated/certificateauthority"
	"powerssl.io/pkg/resource/generated/certificateissue"
)

func Run(grpcAddr string) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var duration metrics.Histogram
	{
		// Endpoint-level metrics.
		duration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "powerssl_io",
			Subsystem: "api",
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds.",
		}, []string{"method", "success"})
	}

	var db *gorm.DB
	{
		var err error
		if db, err = gorm.Open("sqlite3", "/tmp/gorm.db"); err != nil {
			logger.Log("database", "sqlite3", "during", "Open", "err", err)
			os.Exit(1)
		}
		defer db.Close()
	}

	resources := []resource.Resource{
		certificate.New(db, logger, duration),
		certificateauthority.New(db, logger, duration),
		certificateissue.New(db, logger, duration),
	}

	var g group.Group
	{
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", grpcAddr)
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

			for _, r := range resources {
				r.RegisterGRPCServer(baseServer)
			}

			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
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
