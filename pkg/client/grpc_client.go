package transport

import (
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"

	certificateauthorityservice "powerssl.io/pkg/resources/certificate_authority/service"
	certificateauthoritytransport "powerssl.io/pkg/resources/certificate_authority/transport"
)

type grpcClient struct {
	CertificateAuthority certificateauthorityservice.Service
}

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) *grpcClient {
	return &grpcClient{
		CertificateAuthority: certificateauthoritytransport.NewGRPCClient(conn, logger),
	}
}
