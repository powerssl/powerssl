package transport // import "powerssl.io/pkg/controller/workflow/transport"

import (
	"fmt"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/workflow/endpoint"
	service "powerssl.io/pkg/controller/workflow/service"
	"powerssl.io/pkg/util/auth"
)

const serviceName = "powerssl.controller.v1.WorkflowService"

func NewGRPCClient(conn *grpc.ClientConn, key []byte, logger log.Logger, tracer stdopentracing.Tracer) service.Service {
	jwtSigner := jwt.NewSigner("TODO", key, auth.Method, stdjwt.StandardClaims{})

	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(jwt.ContextToGRPC()),
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	}

	var createEndpoint kitendpoint.Endpoint
	{
		createEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"Create",
			encodeGRPCCreateRequest,
			decodeGRPCCreateResponse,
			apiv1.Workflow{},
			options...,
		).Endpoint()
		createEndpoint = jwtSigner(createEndpoint)
		createEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/Create", serviceName))(createEndpoint)
	}

	return endpoint.Endpoints{
		CreateEndpoint: createEndpoint,
	}
}
