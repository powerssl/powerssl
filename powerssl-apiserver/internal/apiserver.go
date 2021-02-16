package internal

import (
	"context"
	"io"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/apiserver/internal/acmeaccount"
	"powerssl.dev/apiserver/internal/acmeserver"
	"powerssl.dev/apiserver/internal/certificate"
	"powerssl.dev/apiserver/internal/repository"
	"powerssl.dev/apiserver/internal/user"
	temporalclient "powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/auth"
	"powerssl.dev/common/tracing"
	"powerssl.dev/common/transport"
	"powerssl.dev/common/util"
)

const component = "powerssl-apiserver"

func Run(cfg *Config) (err error) {
	zapLogger, logger := util.NewZapAndKitLogger()

	cfg.ServerConfig.VaultToken = cfg.VaultClientConfig.Token
	cfg.ServerConfig.VaultURL = cfg.VaultClientConfig.URL
	cfg.ServerConfig.VaultRole = component

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	var tracer opentracing.Tracer
	{
		var closer io.Closer
		if tracer, closer, err = tracing.Init(component, cfg.Tracer, log.With(logger, "component", "tracing")); err != nil {
			return err
		}
		defer func() {
			err = closer.Close()
		}()
	}

	var repositories *repository.Repositories
	{
		var db *sqlx.DB
		if db, err = sqlx.Connect(cfg.DB.Dialect, cfg.DB.Connection); err != nil {
			return err
		}
		defer func() {
			_ = db.Close()
		}()
		repositories = repository.NewRepositories(db, zapLogger)
	}

	var temporalClient temporalclient.Client
	{
		var closer io.Closer
		if temporalClient, closer, err = temporalclient.NewClient(cfg.TemporalClientConfig, zapLogger, tracer, component); err != nil {
			return err
		}
		defer func() {
			temporalClient.Close()
			err = closer.Close()
		}()
	}

	var vaultClient *vault.Client
	{
		if vaultClient, err = vault.New(cfg.VaultClientConfig); err != nil {
			return err
		}
	}
	var _ = vaultClient // TODO: Needed here?

	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "powerssl_io",
		Subsystem: "api",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

	var authMiddleware kitendpoint.Middleware
	{
		if authMiddleware, err = auth.NewParser(cfg.JWKS.URL); err != nil {
			return err
		}
	}

	if cfg.Metrics.Addr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.Metrics.Addr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return backendtransport.ServeGRPC(ctx, &cfg.ServerConfig, log.With(logger, "transport", "gRPC"), []backendtransport.Service{
			acmeaccount.New(repositories, logger, tracer, duration, temporalClient, authMiddleware),
			acmeserver.New(repositories, logger, tracer, duration, authMiddleware),
			certificate.New(repositories, logger, tracer, duration, authMiddleware),
			user.New(repositories, logger, tracer, duration, authMiddleware),
		})
	})

	if err = g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			return err
		}
	}
	return nil
}
