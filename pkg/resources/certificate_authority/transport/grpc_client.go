package transport

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	types "github.com/gogo/protobuf/types"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/api/v1"
	"powerssl.io/pkg/resources/certificate_authority/endpoints"
	"powerssl.io/pkg/resources/certificate_authority/service"
)

const serviceName = "powerssl.api.v1.CertificateAuthorityService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	options := []grpctransport.ClientOption{}

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"Create",
			encodeGRPCCreateRequest,
			decodeGRPCCreateResponse,
			apiv1.CertificateAuthority{},
			options...,
		).Endpoint()
	}

	var deleteEndpoint endpoint.Endpoint
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
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"Get",
			encodeGRPCGetRequest,
			decodeGRPCGetResponse,
			apiv1.CertificateAuthority{},
			options...,
		).Endpoint()
	}

	var listEndpoint endpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"List",
			encodeGRPCListRequest,
			decodeGRPCListResponse,
			apiv1.ListCertificateAuthoritiesResponse{},
			options...,
		).Endpoint()
	}

	var updateEndpoint endpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"Update",
			encodeGRPCUpdateRequest,
			decodeGRPCUpdateResponse,
			apiv1.CertificateAuthority{},
			options...,
		).Endpoint()
	}

	return endpoints.Endpoints{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
		UpdateEndpoint: updateEndpoint,
	}
}
