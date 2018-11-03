package controller

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	apiserverclient "powerssl.io/pkg/apiserver/client"
	"powerssl.io/pkg/controller/acme"
	"powerssl.io/pkg/controller/integration"
	"powerssl.io/pkg/controller/workflow"
)

type service interface {
	RegisterGRPCServer(baseServer *grpc.Server)
}

func makeServices(logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *apiserverclient.GRPCClient) []service {
	return []service{
		acme.New(logger, tracer, duration),
		integration.New(logger, duration), // TODO: tracing
		workflow.New(logger, tracer, duration),
	}
}
