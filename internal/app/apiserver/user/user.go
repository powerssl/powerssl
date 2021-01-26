package user

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-pg/pg/v10"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"powerssl.dev/powerssl/internal/app/apiserver/user/endpoint"
	"powerssl.dev/powerssl/internal/app/apiserver/user/service"
	"powerssl.dev/powerssl/internal/app/apiserver/user/transport"
	apiv1 "powerssl.dev/powerssl/internal/pkg/apiserver/api/v1"
)

type User struct {
	endpoints endpoint.Endpoints
	logger    log.Logger
	tracer    stdopentracing.Tracer
}

func New(db *pg.DB, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, auth kitendpoint.Middleware) *User {
	svc := service.New(db, logger)
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
