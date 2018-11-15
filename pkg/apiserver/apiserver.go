package apiserver

import (
	"context"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/qor/validations"
	otgorm "github.com/smacker/opentracing-gorm"
	"golang.org/x/sync/errgroup"

	controllerclient "powerssl.io/pkg/controller/client"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/tracing"
)

func Run(grpcAddr, grpcCertFile, grpcKeyFile string, grpcInsecure bool, dbDialect, dbConnection, metricsAddr, tracerImpl, controllerAddr, controllerCertFile, controllerServerNameOverride string, controllerInsecure, controllerInsecureSkipTLSVerify bool, controllerAuthToken string) {
	logger := util.NewLogger(os.Stdout)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	tracer, closer, err := tracing.Init("powerssl-apiserver", tracerImpl, log.With(logger, "component", "tracing"))
	if err != nil {
		logger.Log("component", "tracing", "err", err)
		os.Exit(1)
	}
	defer closer.Close()

	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "powerssl_io",
		Subsystem: "api",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

	var db *gorm.DB
	{
		var err error
		if db, err = gorm.Open(dbDialect, dbConnection); err != nil {
			logger.Log("database", dbDialect, "during", "Open", "err", err)
			os.Exit(1)
		}
		defer db.Close()
		otgorm.AddGormCallbacks(db)
		validations.RegisterCallbacks(db)
	}

	var client *controllerclient.GRPCClient
	{
		var err error
		if client, err = controllerclient.NewGRPCClient(controllerAddr, controllerCertFile, controllerServerNameOverride, controllerInsecure, controllerInsecureSkipTLSVerify, controllerAuthToken, logger, tracer); err != nil {
			logger.Log("transport", "gRPC", "during", "Connect", "err", err)
			os.Exit(1)
		}
	}

	services := makeServices(db, logger, tracer, duration, client)

	if metricsAddr != "" {
		g.Go(func() error {
			return util.ServeMetrics(ctx, metricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return util.ServeGRPC(ctx, grpcAddr, grpcCertFile, grpcKeyFile, grpcInsecure, log.With(logger, "transport", "gRPC"), services)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}
