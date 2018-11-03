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
)

type resource interface {
	RegisterGRPCServer(baseServer *grpc.Server)
}

func makeResources(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient) []resource {
	return []resource{
		certificate.New(db, logger, tracer, duration, client),
		certificateauthority.New(db, logger, duration, client), // TODO: tracing
		certificateissue.New(db, logger, duration, client),     // TODO: tracing
	}
}
