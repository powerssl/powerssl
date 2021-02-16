package transport // import "powerssl.dev/sdk/apiserver/user/transport"

import (
	"fmt"

	"github.com/go-kit/kit/auth/jwt"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	apiv1 "powerssl.dev/sdk/apiserver/api/v1"
	"powerssl.dev/sdk/apiserver/user"
	"powerssl.dev/sdk/apiserver/user/endpoint"
)

const serviceName = "powerssl.apiserver.v1.UserService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger, tracer stdopentracing.Tracer, authSigner kitendpoint.Middleware) user.Service {
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
			apiv1.User{},
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
			apiv1.User{},
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
			apiv1.ListUsersResponse{},
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
			apiv1.User{},
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
