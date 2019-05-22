package apiserver

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.io/powerssl/internal/app/apiserver/acmeaccount"
	"powerssl.io/powerssl/internal/app/apiserver/acmeserver"
	"powerssl.io/powerssl/internal/app/apiserver/certificate"
	"powerssl.io/powerssl/internal/app/apiserver/user"
	"powerssl.io/powerssl/internal/pkg/auth"
	"powerssl.io/powerssl/internal/pkg/util"
	"powerssl.io/powerssl/internal/pkg/vault"
	controllerclient "powerssl.io/powerssl/pkg/controller/client"
)

func makeServices(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient, vaultClient *vault.Client, jwtPublicKeyFile string) ([]util.Service, error) {
	auth, err := auth.NewParser(jwtPublicKeyFile)
	if err != nil {
		return nil, err
	}
	return []util.Service{
		acmeaccount.New(db, logger, tracer, duration, client, vaultClient, auth),
		acmeserver.New(db, logger, tracer, duration, client, auth),
		certificate.New(db, logger, tracer, duration, client, auth),
		user.New(db, logger, tracer, duration, client, auth),
	}, nil
}
