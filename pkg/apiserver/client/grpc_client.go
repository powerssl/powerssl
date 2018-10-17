package client

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

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

func NewGRPCClient(grpcAddr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool) *GRPCClient {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var conn *grpc.ClientConn
	{
		var err error
		opts := []grpc.DialOption{
			grpc.WithTimeout(time.Second),
		}
		if insecure {
			opts = append(opts, grpc.WithInsecure())
		} else if insecureSkipTLSVerify {
			creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			creds, err := credentials.NewClientTLSFromFile(certFile, serverNameOverride)
			if err != nil {
				logger.Log("transport", "gRPC", "err", fmt.Errorf("Failed to create TLS credentials %v", err))
			}
			opts = append(opts, grpc.WithTransportCredentials(creds))
		}
		conn, err = grpc.Dial(grpcAddr, opts...)
		if err != nil {
			logger.Log("transport", "gRPC", "err", err)
			os.Exit(1)
		}
	}

	return &GRPCClient{
		Certificate:          certificatetransport.NewGRPCClient(conn, logger),
		CertificateAuthority: certificateauthoritytransport.NewGRPCClient(conn, logger),
		CertificateIssue:     certificateissuetransport.NewGRPCClient(conn, logger),
	}
}
