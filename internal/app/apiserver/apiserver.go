package apiserver

import (
	"context"
	"io"
	"os"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeserver"
	"powerssl.dev/powerssl/internal/app/apiserver/certificate"
	"powerssl.dev/powerssl/internal/app/apiserver/repository"
	"powerssl.dev/powerssl/internal/app/apiserver/user"
	"powerssl.dev/powerssl/internal/pkg/auth"
	temporalclient "powerssl.dev/powerssl/internal/pkg/temporal/client"
	"powerssl.dev/powerssl/internal/pkg/tracing"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/util"
	"powerssl.dev/powerssl/internal/pkg/vault"
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

	var tracer opentracing.Tracer
	{
		var err error
		var closer io.Closer
		if tracer, closer, err = tracing.Init(component, cfg.Tracer, log.With(logger, "component", "tracing")); err != nil {
			_ = logger.Log("component", "tracing", "err", err)
			os.Exit(1)
		}
		defer func() {
			_ = closer.Close()
		}()
	}

	var repositories *repository.Repositories
	{
		var db *sqlx.DB
		var err error
		if db, err = sqlx.Connect(cfg.DBDialect, cfg.DBConnection); err != nil {
			_ = logger.Log("database", cfg.DBDialect, "err", err)
			os.Exit(1)
		}
		defer func() {
			_ = db.Close()
		}()
		zapLogger, _ := zap.NewDevelopment() // TODO
		repositories = repository.NewRepositories(db, zapLogger)
	}

	var temporalClient temporalclient.Client
	{
		var err error
		if temporalClient, err = temporalclient.NewClient(temporalclient.Config{
			CAFile:    cfg.VaultClientConfig.CAFile, // TODO Wrong cfg path
			HostPort:  cfg.TemporalClientConfig.HostPort,
			Namespace: cfg.TemporalClientConfig.Namespace,
		}, nil, tracer); err != nil {
			_ = logger.Log("err", err)
			os.Exit(1)
		}
		defer temporalClient.Close()
	}

	var vaultClient *vault.Client
	{
		var err error
		if vaultClient, err = vault.New(cfg.VaultClientConfig.URL, cfg.VaultClientConfig.Token, cfg.VaultClientConfig.CAFile); err != nil {
			_ = logger.Log("err", err)
			os.Exit(1)
		}
	}

	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "powerssl_io",
		Subsystem: "api",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

	var authMiddleware kitendpoint.Middleware
	{
		var err error
		if authMiddleware, err = auth.NewParser(cfg.JWKSURL); err != nil {
			_ = logger.Log("err", err)
			os.Exit(1)
		}
	}

	if cfg.MetricsAddr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.MetricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return transport.ServeGRPC(ctx, cfg.ServerConfig, log.With(logger, "transport", "gRPC"), []transport.Service{
			acmeaccount.New(repositories, logger, tracer, duration, temporalClient, vaultClient, authMiddleware),
			acmeserver.New(repositories, logger, tracer, duration, authMiddleware),
			certificate.New(repositories, logger, tracer, duration, authMiddleware),
			user.New(repositories, logger, tracer, duration, authMiddleware),
		})
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			_ = logger.Log("err", err)
		}
	}
}
