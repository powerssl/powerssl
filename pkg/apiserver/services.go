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
)

func makeServices(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient) []util.Service {
	return []util.Service{
		acmeaccount.New(db, logger, tracer, duration, client),
		acmeserver.New(db, logger, tracer, duration, client),
		certificate.New(db, logger, tracer, duration, client),
	}
}
