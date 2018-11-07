package health

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Health struct {
}

func New() *Health {
	return &Health{}
}

func (*Health) RegisterGRPCServer(baseServer *grpc.Server) {
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(baseServer, healthServer)
}
