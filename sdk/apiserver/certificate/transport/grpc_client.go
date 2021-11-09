package transport // import "powerssl.dev/sdk/apiserver/certificate/transport"

import (
	"fmt"

	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"
	stdopentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/common/log"

	"powerssl.dev/sdk/apiserver/certificate"
	"powerssl.dev/sdk/apiserver/certificate/endpoint"
)

const serviceName = "powerssl.apiserver.v1.CertificateService"

func NewGRPCClient(conn *grpc.ClientConn, logger *zap.SugaredLogger, tracer stdopentracing.Tracer, authSigner kitendpoint.Middleware) certificate.Service {
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(jwt.ContextToGRPC()),
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, log.KitLogger(logger))),
	}

	var createEndpoint kitendpoint.Endpoint
	{
		createEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"Create",
			encodeGRPCCreateRequest,
			decodeGRPCCreateResponse,
			apiv1.Certificate{},
			options...,
		).Endpoint()
		createEndpoint = authSigner(createEndpoint)
		createEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/Create", serviceName))(createEndpoint)
	}

	var deleteEndpoint kitendpoint.Endpoint
	{
		deleteEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"Delete",
			encodeGRPCDeleteRequest,
			decodeGRPCDeleteResponse,
			types.Empty{},
			options...,
		).Endpoint()
		deleteEndpoint = authSigner(deleteEndpoint)
		deleteEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/Delete", serviceName))(deleteEndpoint)
	}

	var getEndpoint kitendpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"Get",
			encodeGRPCGetRequest,
			decodeGRPCGetResponse,
			apiv1.Certificate{},
			options...,
		).Endpoint()
		getEndpoint = authSigner(getEndpoint)
		getEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/Get", serviceName))(getEndpoint)
	}

	var listEndpoint kitendpoint.Endpoint
	{
		listEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"List",
			encodeGRPCListRequest,
			decodeGRPCListResponse,
			apiv1.ListCertificatesResponse{},
			options...,
		).Endpoint()
		listEndpoint = authSigner(listEndpoint)
		listEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/List", serviceName))(listEndpoint)
	}

	var updateEndpoint kitendpoint.Endpoint
	{
		updateEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"Update",
			encodeGRPCUpdateRequest,
			decodeGRPCUpdateResponse,
			apiv1.Certificate{},
			options...,
		).Endpoint()
		updateEndpoint = authSigner(updateEndpoint)
		updateEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/Update", serviceName))(updateEndpoint)
	}

	return endpoint.Endpoints{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
		UpdateEndpoint: updateEndpoint,
	}
}
