package powerctl

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"

	certificateservice "powerssl.io/pkg/resource/certificate"
	certificateauthorityservice "powerssl.io/pkg/resource/certificateauthority"
	certificateissueservice "powerssl.io/pkg/resource/certificateissue"
	certificatetransport "powerssl.io/pkg/resource/generated/certificate/transport"
	certificateauthoritytransport "powerssl.io/pkg/resource/generated/certificateauthority/transport"
	certificateissuetransport "powerssl.io/pkg/resource/generated/certificateissue/transport"
)

type GRPCClient struct {
	Certificate          certificateservice.Service
	CertificateAuthority certificateauthorityservice.Service
	CertificateIssue     certificateissueservice.Service
}

func NewGRPCClient(grpcAddr string) *GRPCClient {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var conn *grpc.ClientConn
	{
		var err error
		conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			logger.Log("error: %v", err)
			os.Exit(1)
		}
	}

	return &GRPCClient{
		Certificate:          certificatetransport.NewGRPCClient(conn, logger),
		CertificateAuthority: certificateauthoritytransport.NewGRPCClient(conn, logger),
		CertificateIssue:     certificateissuetransport.NewGRPCClient(conn, logger),
	}
}
