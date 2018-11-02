package apiserver

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"

	controllerclient "powerssl.io/pkg/controller/client"
	"powerssl.io/pkg/resource/generated/certificate"
	"powerssl.io/pkg/resource/generated/certificateauthority"
	"powerssl.io/pkg/resource/generated/certificateissue"
)

type resource interface {
	RegisterGRPCServer(baseServer *grpc.Server)
}

func makeResources(db *gorm.DB, logger log.Logger, duration metrics.Histogram, client *controllerclient.GRPCClient) []resource {
	return []resource{
		certificate.New(db, logger, duration, client),
		certificateauthority.New(db, logger, duration, client),
		certificateissue.New(db, logger, duration, client),
	}
}
