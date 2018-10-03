package resources

import "google.golang.org/grpc"

type Resource interface {
	RegisterGRPCServer(baseServer *grpc.Server)
}
