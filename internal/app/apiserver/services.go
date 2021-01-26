package apiserver

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-pg/pg/v10"
	stdopentracing "github.com/opentracing/opentracing-go"
	temporalclient "go.temporal.io/sdk/client"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeserver"
	"powerssl.dev/powerssl/internal/app/apiserver/certificate"
	"powerssl.dev/powerssl/internal/app/apiserver/user"
	"powerssl.dev/powerssl/internal/pkg/auth"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/vault"
)

func makeServices(db *pg.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, temporalClient temporalclient.Client, vaultClient *vault.Client, jwtPublicKeyFile string) ([]transport.Service, error) {
	auth, err := auth.NewParser(jwtPublicKeyFile)
	if err != nil {
		return nil, err
	}
	return []transport.Service{
		acmeaccount.New(db, logger, tracer, duration, temporalClient, vaultClient, auth),
		acmeserver.New(db, logger, tracer, duration, auth),
		certificate.New(db, logger, tracer, duration, auth),
		user.New(db, logger, tracer, duration, auth),
	}, nil
}
