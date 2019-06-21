package controller

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.io/powerssl/internal/app/controller/acme"
	"powerssl.io/powerssl/internal/app/controller/integration"
	"powerssl.io/powerssl/internal/app/controller/workflow"
	"powerssl.io/powerssl/internal/pkg/auth"
	"powerssl.io/powerssl/internal/pkg/transport"
	apiserverclient "powerssl.io/powerssl/pkg/apiserver/client"
)

func makeServices(logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *apiserverclient.GRPCClient, jwtPublicKeyFile string) ([]transport.Service, error) {
	auth, err := auth.NewParser(jwtPublicKeyFile)
	if err != nil {
		return nil, err
	}
	return []transport.Service{
		acme.New(logger, tracer, duration),
		integration.New(logger, duration), // TODO: tracing
		workflow.New(logger, tracer, duration, client, auth),
	}, nil
}
