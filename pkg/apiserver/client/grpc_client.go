package client

import (
	"crypto/tls"
	"time"

	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	acmeaccountservice "powerssl.io/pkg/apiserver/acmeaccount"
	acmeaccounttransport "powerssl.io/pkg/apiserver/acmeaccount/generated/transport"
	certificateservice "powerssl.io/pkg/apiserver/certificate"
	certificatetransport "powerssl.io/pkg/apiserver/certificate/generated/transport"
	certificateissueservice "powerssl.io/pkg/apiserver/certificateissue"
	certificateissuetransport "powerssl.io/pkg/apiserver/certificateissue/generated/transport"
)

type GRPCClient struct {
	ACMEAccount      acmeaccountservice.Service
	Certificate      certificateservice.Service
	CertificateIssue certificateissueservice.Service
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
		ACMEAccount:      acmeaccounttransport.NewGRPCClient(conn, logger, tracer),
		Certificate:      certificatetransport.NewGRPCClient(conn, logger, tracer),
		CertificateIssue: certificateissuetransport.NewGRPCClient(conn, logger),
	}, nil
}
