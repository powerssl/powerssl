package domain

import (
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"

	"github.com/jinzhu/gorm"

	pb "powerssl.io/pkg/api/v1"
	"powerssl.io/pkg/util/resource"

	endpoint "powerssl.io/pkg/domain/endpoint"
	service "powerssl.io/pkg/domain/service"
	transport "powerssl.io/pkg/domain/transport"
)

type Domain struct {
	db     *gorm.DB
	logger log.Logger

	endpoints endpoint.Set
	service   *service.Service
}

func New(db *gorm.DB, logger log.Logger) resource.APIResource {
	svc := service.New(db, logger)
	endpoints := endpoint.New(svc, logger)

	return &Domain{
		db:     db,
		logger: logger,

		endpoints: endpoints,
		service:   &svc,
	}
}

func (d Domain) RegisterGRPCServer(baseServer *grpc.Server) {
	pb.RegisterDomainServiceServer(baseServer, d.grpcServer())
}

func (d Domain) grpcServer() pb.DomainServiceServer {
	return transport.NewGRPCServer(d.endpoints, d.logger)
}
