package apiserver

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.io/pkg/apiserver/acmeaccount"
	"powerssl.io/pkg/apiserver/acmeserver"
	"powerssl.io/pkg/apiserver/certificate"
	"powerssl.io/pkg/apiserver/user"
	controllerclient "powerssl.io/pkg/controller/client"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/auth"
	"powerssl.io/pkg/util/vault"
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
