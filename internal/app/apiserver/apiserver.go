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

	"powerssl.dev/powerssl/internal/pkg/auth"
	temporalclient "powerssl.dev/powerssl/internal/pkg/temporal/client"
	"powerssl.dev/powerssl/internal/pkg/tracing"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
	"powerssl.dev/powerssl/internal/pkg/vault"
	controllerclient "powerssl.dev/powerssl/pkg/controller/client"
)

const component = "powerssl-apiserver"

func Run(cfg *Config) {
	logger := util.NewLogger(os.Stdout)

	cfg.ServerConfig.VaultRole = component
	util.ValidateConfig(cfg, logger)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	tracer, closer, err := tracing.Init(component, cfg.Tracer, log.With(logger, "component", "tracing"))
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
		if db, err = gorm.Open(cfg.DBDialect, cfg.DBConnection); err != nil {
			logger.Log("database", cfg.DBDialect, "during", "Open", "err", err)
			os.Exit(1)
		}
		defer db.Close()
		otgorm.AddGormCallbacks(db)
		validations.RegisterCallbacks(db)
	}

	var client *controllerclient.GRPCClient
	{
		token, err := auth.NewServiceToken(cfg.AuthToken)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if client, err = controllerclient.NewGRPCClient(ctx, cfg.ControllerClientConfig, token, logger, tracer); err != nil {
			logger.Log("transport", "gRPC", "during", "Connect", "err", err)
			os.Exit(1)
		}
	}

	var temporalClient temporalclient.Client
	{
		var err error
		if temporalClient, err = temporalclient.NewClient(temporalclient.Config{
			CAFile:    cfg.VaultClientConfig.CAFile, // TODO Wrong cfg path
			HostPort:  cfg.TemporalClientConfig.HostPort,
			Namespace: cfg.TemporalClientConfig.Namespace,
		}, nil, tracer); err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		defer temporalClient.Close()
	}

	var vaultClient *vault.Client
	{
		var err error
		if vaultClient, err = vault.New(cfg.VaultClientConfig.URL, cfg.VaultClientConfig.Token, cfg.VaultClientConfig.CAFile); err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
	}

	services, err := makeServices(db, logger, tracer, duration, client, temporalClient, vaultClient, cfg.JWKSURL)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	if cfg.MetricsAddr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.MetricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return transport.ServeGRPC(ctx, cfg.ServerConfig, log.With(logger, "transport", "gRPC"), services)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}
