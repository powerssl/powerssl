package transport

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ClientConfig struct {
	Addr                  string `validate:"required,hostname_port"`
	CAFile                string `mapstructure:"ca-file"`
	Insecure              bool
	InsecureSkipTLSVerify bool   `mapstructure:"insecure-skip-tls-verify"`
	ServerNameOverride    string `mapstructure:"server-name-override"`
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
