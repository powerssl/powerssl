package powerctl

import (
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"

	certificateservice "powerssl.io/pkg/resource/generated/certificate/service"
	certificatetransport "powerssl.io/pkg/resource/generated/certificate/transport"
	certificateauthorityservice "powerssl.io/pkg/resource/generated/certificateauthority/service"
	certificateauthoritytransport "powerssl.io/pkg/resource/generated/certificateauthority/transport"
	certificateissueservice "powerssl.io/pkg/resource/generated/certificateissue/service"
	certificateissuetransport "powerssl.io/pkg/resource/generated/certificateissue/transport"
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
