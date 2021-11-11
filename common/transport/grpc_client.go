package transport // import "powerssl.dev/common/transport"

import (
	"context"
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func New(ctx context.Context, cfg Config) (*grpc.ClientConn, error) {
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
