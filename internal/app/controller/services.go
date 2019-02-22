package controller

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.io/internal/app/controller/acme"
	"powerssl.io/internal/app/controller/integration"
	"powerssl.io/internal/app/controller/workflow"
	"powerssl.io/internal/pkg/auth"
	"powerssl.io/internal/pkg/util"
	apiserverclient "powerssl.io/pkg/apiserver/client"
)

func makeServices(logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *apiserverclient.GRPCClient, jwtPublicKeyFile string) ([]util.Service, error) {
	auth, err := auth.NewParser(jwtPublicKeyFile)
	if err != nil {
		return nil, err
	}
	return []util.Service{
		acme.New(logger, tracer, duration),
		integration.New(logger, duration), // TODO: tracing
		workflow.New(logger, tracer, duration, client, auth),
	}, nil
}
