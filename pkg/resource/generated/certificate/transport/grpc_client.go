// Code generated by protoc-gen-gotemplate. DO NOT EDIT.

package transport // import "powerssl.io/pkg/resource/generated/certificate/transport"

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/apiserver/api/v1"
	service "powerssl.io/pkg/resource/certificate"
	"powerssl.io/pkg/resource/generated/certificate/endpoint"
)

const serviceName = "powerssl.apiserver.v1.CertificateService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	options := []grpctransport.ClientOption{}

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
	}

	return endpoint.Endpoints{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
		UpdateEndpoint: updateEndpoint,
	}
}
