package controller

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	stdopentracing "github.com/opentracing/opentracing-go"

	apiserverclient "powerssl.io/pkg/apiserver/client"
	"powerssl.io/pkg/controller/acme"
	"powerssl.io/pkg/controller/integration"
	"powerssl.io/pkg/controller/workflow"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/auth"
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
