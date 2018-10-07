package transport

import (
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"

	certificateservice "powerssl.io/pkg/gen/certificate/service"
	certificatetransport "powerssl.io/pkg/gen/certificate/transport"
	certificateauthorityservice "powerssl.io/pkg/gen/certificateauthority/service"
	certificateauthoritytransport "powerssl.io/pkg/gen/certificateauthority/transport"
	certificateissueservice "powerssl.io/pkg/gen/certificateissue/service"
	certificateissuetransport "powerssl.io/pkg/gen/certificateissue/transport"
)

type grpcClient struct {
	Certificate          certificateservice.Service
	CertificateAuthority certificateauthorityservice.Service
	CertificateIssue     certificateissueservice.Service
}

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) *grpcClient {
	return &grpcClient{
		Certificate:          certificatetransport.NewGRPCClient(conn, logger),
		CertificateAuthority: certificateauthoritytransport.NewGRPCClient(conn, logger),
		CertificateIssue:     certificateissuetransport.NewGRPCClient(conn, logger),
	}
}
