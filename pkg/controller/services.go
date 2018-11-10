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
)

func makeServices(logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, client *apiserverclient.GRPCClient) []util.Service {
	return []util.Service{
		acme.New(logger, tracer, duration),
		integration.New(logger, duration), // TODO: tracing
		workflow.New(logger, tracer, duration),
	}
}
