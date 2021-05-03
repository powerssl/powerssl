package internal

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/backend/auth"
	temporalclient "powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common"
	"powerssl.dev/common/log"
	"powerssl.dev/common/tracing"
	"powerssl.dev/common/transport"

	"powerssl.dev/apiserver/internal/acmeaccount"
	"powerssl.dev/apiserver/internal/acmeserver"
	"powerssl.dev/apiserver/internal/certificate"
	"powerssl.dev/apiserver/internal/repository"
	"powerssl.dev/apiserver/internal/user"
)

const component = "powerssl-apiserver"

func Run(cfg *Config) (err error) {
	var logger log.Logger
	if logger, err = log.NewLogger(false); err != nil {
		return err
	}
	defer common.ErrWrapSync(logger, &err)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return common.InterruptHandler(ctx, logger)
	})

	var tracer opentracing.Tracer
	{
		var closer io.Closer
		if tracer, closer, err = tracing.Init(component, cfg.Tracer, logger.With("component", "tracing")); err != nil {
			return err
		}
		defer common.ErrWrapCloser(closer, &err)
	}

	var repositories *repository.Repositories
	{
		var db *sqlx.DB
		if db, err = sqlx.Connect(cfg.DB.Dialect, cfg.DB.Connection); err != nil {
			return err
		}
		defer common.ErrWrapCloser(db, &err)
		repositories = repository.NewRepositories(db, logger)
	}

	var temporalClient temporalclient.Client
	{
		var closer io.Closer
		if temporalClient, closer, err = temporalclient.NewClient(cfg.TemporalClientConfig, logger, tracer, component); err != nil {
			return err
		}
		defer func() {
			temporalClient.Close()
			common.ErrWrapCloser(closer, &err)
		}()
	}

	var vaultClient *vault.Client
	{
		if vaultClient, err = vault.New(&cfg.VaultClientConfig); err != nil {
			return err
		}
		cfg.ServerConfig.VaultToken = cfg.VaultClientConfig.Token
		cfg.ServerConfig.VaultURL = cfg.VaultClientConfig.URL
		cfg.ServerConfig.VaultRole = component
	}
	var _ = vaultClient // TODO: Needed here?

	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "powerssl_io",
		Subsystem: "api",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

	var rootCAs *x509.CertPool
	{
		if rootCAs, _ = x509.SystemCertPool(); rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}
		var byt []byte
		if byt, err = ioutil.ReadFile(cfg.CAFile); err != nil {
			return fmt.Errorf("failed to append %q to RootCAs: %v", cfg.CAFile, err)
		}
		if ok := rootCAs.AppendCertsFromPEM(byt); !ok {
			logger.Infow("no certs appended, using system certs only")
		}
	}

	var authMiddleware kitendpoint.Middleware
	{
		if authMiddleware, err = auth.NewParser(cfg.JWKS.URL, &tls.Config{
			InsecureSkipVerify: cfg.JWKS.InsecureSkipTLSVerify,
			RootCAs:            rootCAs,
			ServerName:         cfg.JWKS.ServerNameOverride,
		}); err != nil {
			return err
		}
	}

	if cfg.Metrics.Addr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.Metrics.Addr, logger.With("component", "metrics"))
		})
	}

	g.Go(func() error {
		return backendtransport.ServeGRPC(ctx, cfg.ServerConfig, logger.With("transport", "gRPC"), backendtransport.Services{
			acmeaccount.NewService(repositories, logger, tracer, duration, temporalClient, authMiddleware),
			acmeserver.NewService(repositories, logger, tracer, duration, authMiddleware),
			certificate.NewService(repositories, logger, tracer, duration, authMiddleware),
			user.NewService(repositories, logger, tracer, duration, authMiddleware),
		})
	})

	if err = g.Wait(); err != nil {
		switch err.(type) {
		case common.InterruptError:
		default:
			return err
		}
	}
	return nil
}
