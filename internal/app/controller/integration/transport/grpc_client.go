package transport

import (
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/powerssl/internal/pkg/controller/api/v1"
)

func NewGRPCClient(conn *grpc.ClientConn, _ log.Logger) apiv1.IntegrationServiceClient {
	return apiv1.NewIntegrationServiceClient(conn)
}
