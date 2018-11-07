package apiserver

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	certificate "powerssl.io/pkg/apiserver/certificate/generated"
	certificateauthority "powerssl.io/pkg/apiserver/certificateauthority/generated"
	certificateissue "powerssl.io/pkg/apiserver/certificateissue/generated"
	controllerclient "powerssl.io/pkg/controller/client"
	"powerssl.io/pkg/util/health"
)

type service interface {
	RegisterGRPCServer(baseServer *grpc.Server)
}

func makeServices(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient) []service {
	return []service{
		certificate.New(db, logger, tracer, duration, client),
		certificateauthority.New(db, logger, duration, client), // TODO: tracing
		certificateissue.New(db, logger, duration, client),     // TODO: tracing
		health.New(),
	}
}
