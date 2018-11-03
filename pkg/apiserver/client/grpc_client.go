package client

import (
	"crypto/tls"
	"time"

	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
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

func NewGRPCClient(grpcAddr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool, logger log.Logger, tracer stdopentracing.Tracer) (*GRPCClient, error) {
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
				return nil, err
			}
			opts = append(opts, grpc.WithTransportCredentials(creds))
		}
		conn, err = grpc.Dial(grpcAddr, opts...)
		if err != nil {
			return nil, err
		}
	}

	return &GRPCClient{
		Certificate:          certificatetransport.NewGRPCClient(conn, logger, tracer),
		CertificateAuthority: certificateauthoritytransport.NewGRPCClient(conn, logger),
		CertificateIssue:     certificateissuetransport.NewGRPCClient(conn, logger),
	}, nil
}
