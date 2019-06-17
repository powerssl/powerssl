package apiserver

import (
	"context"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/qor/validations"
	otgorm "github.com/smacker/opentracing-gorm"
	"golang.org/x/sync/errgroup"

	"powerssl.io/powerssl/internal/pkg/auth"
	"powerssl.io/powerssl/internal/pkg/tracing"
	"powerssl.io/powerssl/internal/pkg/util"
	"powerssl.io/powerssl/internal/pkg/vault"
	controllerclient "powerssl.io/powerssl/pkg/controller/client"
)

const component = "powerssl-apiserver"

func Run(grpcAddr, commonName, vaultURL, vaultToken, grpcCertFile, grpcKeyFile string, grpcInsecure bool, dbDialect, dbConnection, metricsAddr, tracerImpl, caFile, controllerAddr, controllerServerNameOverride string, controllerInsecure, controllerInsecureSkipTLSVerify bool, jwksURL, controllerAuthToken string, etcdEndpoints []string) {
	logger := util.NewLogger(os.Stdout)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	tracer, closer, err := tracing.Init(component, tracerImpl, log.With(logger, "component", "tracing"))
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

	var etcd *clientv3.Client
	{
		var err error
		// TODO: Add auth options here later as well
		etcd, err = clientv3.New(clientv3.Config{
			Endpoints:   etcdEndpoints,
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			logger.Log("etcd", "during", "Connect", "err", err)
			os.Exit(1)
		}
		defer etcd.Close()
	}

	var client *controllerclient.GRPCClient
	{
		token, err := auth.NewServiceToken(controllerAuthToken)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if client, err = controllerclient.NewGRPCClient(controllerAddr, caFile, controllerServerNameOverride, controllerInsecure, controllerInsecureSkipTLSVerify, token, logger, tracer); err != nil {
			logger.Log("transport", "gRPC", "during", "Connect", "err", err)
			os.Exit(1)
		}
	}

	var vaultClient *vault.Client
	{
		var err error
		if vaultClient, err = vault.New(vaultURL, vaultToken, caFile); err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
	}

	services, err := makeServices(db, etcd, logger, tracer, duration, client, vaultClient, jwksURL)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	if metricsAddr != "" {
		g.Go(func() error {
			return util.ServeMetrics(ctx, metricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return util.ServeGRPC(ctx, grpcAddr, grpcCertFile, grpcKeyFile, caFile, commonName, vaultURL, vaultToken, component, grpcInsecure, log.With(logger, "transport", "gRPC"), services)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}
