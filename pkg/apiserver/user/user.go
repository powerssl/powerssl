package user // import "powerssl.io/pkg/apiserver/user"

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/apiserver/api/v1"
	"powerssl.io/pkg/apiserver/user/endpoint"
	"powerssl.io/pkg/apiserver/user/service"
	"powerssl.io/pkg/apiserver/user/transport"
	controllerclient "powerssl.io/pkg/controller/client"
)

type User struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(db *gorm.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *controllerclient.GRPCClient, auth kitendpoint.Middleware) *User {
	svc := service.New(db, logger, client)
	endpoints := endpoint.NewEndpoints(svc, logger, tracer, duration, auth)

	return &User{
		endpoints: endpoints,
		logger:    logger,
		tracer:    tracer,
	}
}

func (user *User) RegisterGRPCServer(baseServer *grpc.Server) {
	grpcServer := transport.NewGRPCServer(user.endpoints, user.logger, user.tracer)
	apiv1.RegisterUserServiceServer(baseServer, grpcServer)
}

func (*User) ServiceName() string {
	return "powerssl.apiserver.v1.UserService"
}
