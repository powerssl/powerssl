package transport // import "powerssl.io/pkg/apiserver/acmeserver/transport"

import (
	"fmt"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"powerssl.io/pkg/apiserver/acmeserver/endpoint"
	"powerssl.io/pkg/apiserver/acmeserver/service"
	apiv1 "powerssl.io/pkg/apiserver/api/v1"
	"powerssl.io/pkg/util/auth"
)

const serviceName = "powerssl.apiserver.v1.ACMEServerService"

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
			apiv1.ACMEServer{},
			options...,
		).Endpoint()
		createEndpoint = jwtSigner(createEndpoint)
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
		deleteEndpoint = jwtSigner(deleteEndpoint)
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
			apiv1.ACMEServer{},
			options...,
		).Endpoint()
		getEndpoint = jwtSigner(getEndpoint)
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
			apiv1.ListACMEServersResponse{},
			options...,
		).Endpoint()
		listEndpoint = jwtSigner(listEndpoint)
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
			apiv1.ACMEServer{},
			options...,
		).Endpoint()
		updateEndpoint = jwtSigner(updateEndpoint)
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
