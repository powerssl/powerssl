package resource

import (
	"google.golang.org/grpc"
)

type APIResource interface {
	RegisterGRPCServer(baseServer *grpc.Server)
}
