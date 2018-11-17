package apiserver

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.io/pkg/apiserver/acmeaccount"
	"powerssl.io/pkg/apiserver/acmeserver"
	"powerssl.io/pkg/apiserver/certificate"
	controllerclient "powerssl.io/pkg/controller/client"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/auth"
)

func makeServices(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient, jwtPublicKeyFile string) ([]util.Service, error) {
	auth, err := auth.NewParser(jwtPublicKeyFile)
	if err != nil {
		return nil, err
	}
	return []util.Service{
		acmeaccount.New(db, logger, tracer, duration, client, auth),
		acmeserver.New(db, logger, tracer, duration, client, auth),
		certificate.New(db, logger, tracer, duration, client, auth),
	}, nil
}
