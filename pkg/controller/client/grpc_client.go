package client

import (
	"crypto/tls"
	"time"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	caservice "powerssl.io/pkg/controller/ca/service"
	catransport "powerssl.io/pkg/controller/ca/transport"
	integrationservice "powerssl.io/pkg/controller/integration/service"
	intregrationtransport "powerssl.io/pkg/controller/integration/transport"
	workflowservice "powerssl.io/pkg/controller/workflow/service"
	workflowtransport "powerssl.io/pkg/controller/workflow/transport"
)

type GRPCClient struct {
	CA          caservice.Service
	Integration integrationservice.Service
	Workflow    workflowservice.Service
}

func NewGRPCClient(grpcAddr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool, logger log.Logger) (*GRPCClient, error) {
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
	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, err
	}

	return &GRPCClient{
		CA:          catransport.NewGRPCClient(conn, logger),
		Integration: intregrationtransport.NewGRPCClient(conn, logger),
		Workflow:    workflowtransport.NewGRPCClient(conn, logger),
	}, nil
}
