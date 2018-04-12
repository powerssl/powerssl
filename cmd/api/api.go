package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	pb "powerssl.io/api/v1"
	"powerssl.io/pkg/domain"
	domainendpoint "powerssl.io/pkg/domain/endpoint"
	domainservice "powerssl.io/pkg/domain/service"
	domaintransport "powerssl.io/pkg/domain/transport"
)

func main() {
	grpcAddr := ":8082"

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var db *gorm.DB
	{
		var err error
		if db, err = gorm.Open("sqlite3", "/tmp/gorm.db"); err != nil {
			logger.Log("database", "sqlite3", "during", "Open", "err", err)
			os.Exit(1)
		}
		defer db.Close()

		db.AutoMigrate(&domain.Domain{})
	}

	var (
		service    = domainservice.New(db, logger)
		endpoints  = domainendpoint.New(service, logger)
		grpcServer = domaintransport.NewGRPCServer(endpoints, logger)
	)

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
			pb.RegisterDomainServiceServer(baseServer, grpcServer)
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
