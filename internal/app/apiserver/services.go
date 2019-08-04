package apiserver

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.dev/powerssl/internal/app/apiserver/acmeaccount"
	"powerssl.dev/powerssl/internal/app/apiserver/acmeserver"
	"powerssl.dev/powerssl/internal/app/apiserver/certificate"
	"powerssl.dev/powerssl/internal/app/apiserver/user"
	"powerssl.dev/powerssl/internal/pkg/auth"
	"powerssl.dev/powerssl/internal/pkg/transport"
	"powerssl.dev/powerssl/internal/pkg/vault"
	controllerclient "powerssl.dev/powerssl/pkg/controller/client"
)

func makeServices(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient, vaultClient *vault.Client, jwtPublicKeyFile string) ([]transport.Service, error) {
	auth, err := auth.NewParser(jwtPublicKeyFile)
	if err != nil {
		return nil, err
	}
	return []transport.Service{
		acmeaccount.New(db, logger, tracer, duration, client, vaultClient, auth),
		acmeserver.New(db, logger, tracer, duration, client, auth),
		certificate.New(db, logger, tracer, duration, client, auth),
		user.New(db, logger, tracer, duration, client, auth),
	}, nil
}
