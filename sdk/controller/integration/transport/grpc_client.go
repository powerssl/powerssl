package transport // import "powerssl.dev/sdk/controller/api"

import (
	"google.golang.org/grpc"

	apiv1 "powerssl.dev/sdk/controller/api/v1"
)

func NewGRPCClient(conn *grpc.ClientConn) apiv1.IntegrationServiceClient {
	return apiv1.NewIntegrationServiceClient(conn)
}