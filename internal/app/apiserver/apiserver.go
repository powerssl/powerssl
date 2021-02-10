package apiserver

import (
	"context"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"
	"io"

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
		return transport.ServeGRPC(ctx, &cfg.ServerConfig, log.With(logger, "transport", "gRPC"), []transport.Service{
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
