package transport

import (
	"context"
	"crypto/tls"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/go-playground/validator.v9"
)

type ClientConfig struct {
	Addr                  string `validate:"required"`
	CAFile                string
	Insecure              bool
	InsecureSkipTLSVerify bool
	ServerNameOverride    string
}

func ClientConfigValidator(sl validator.StructLevel) {
	cfg := sl.Current().Interface().(ClientConfig)

	if !cfg.Insecure && !cfg.InsecureSkipTLSVerify && cfg.CAFile == "" {
		sl.ReportError(cfg.CAFile, "CAFile", "CAFile", "required", "")
	}
}

func NewClientConn(ctx context.Context, cfg *ClientConfig) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	var opts []grpc.DialOption
	if cfg.Insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		var creds credentials.TransportCredentials
		if cfg.InsecureSkipTLSVerify {
			creds = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		} else {
			var err error
			if creds, err = credentials.NewClientTLSFromFile(cfg.CAFile, cfg.ServerNameOverride); err != nil {
				return nil, err
			}
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	return grpc.DialContext(ctx, cfg.Addr, opts...)
}
