package transport // import "powerssl.io/pkg/controller/ca/transport"

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/ca/endpoint"
	service "powerssl.io/pkg/controller/ca/service"
)

const serviceName = "powerssl.controller.v1.CAService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	options := []grpctransport.ClientOption{}

	var getAuthorizeDomainRequestEndpoint kitendpoint.Endpoint
	{
		getAuthorizeDomainRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetAuthorizeDomainRequest",
			encodeGRPCGetAuthorizeDomainRequestRequest,
			decodeGRPCGetAuthorizeDomainRequestResponse,
			apiv1.GetAuthorizeDomainRequestResponse{},
			options...,
		).Endpoint()
	}

	var setAuthorizeDomainResponseEndpoint kitendpoint.Endpoint
	{
		setAuthorizeDomainResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetAuthorizeDomainResponse",
			encodeGRPCSetAuthorizeDomainResponseRequest,
			decodeGRPCSetAuthorizeDomainResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getRequestCertificateRequestEndpoint kitendpoint.Endpoint
	{
		getRequestCertificateRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetRequestCertificateRequest",
			encodeGRPCGetRequestCertificateRequestRequest,
			decodeGRPCGetRequestCertificateRequestResponse,
			apiv1.GetRequestCertificateRequestResponse{},
			options...,
		).Endpoint()
	}

	var setRequestCertificateResponseEndpoint kitendpoint.Endpoint
	{
		setRequestCertificateResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetRequestCertificateResponse",
			encodeGRPCSetRequestCertificateResponseRequest,
			decodeGRPCSetRequestCertificateResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getRevokeCertificateRequestEndpoint kitendpoint.Endpoint
	{
		getRevokeCertificateRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetRevokeCertificateRequest",
			encodeGRPCGetRevokeCertificateRequestRequest,
			decodeGRPCGetRevokeCertificateRequestResponse,
			apiv1.GetRevokeCertificateRequestResponse{},
			options...,
		).Endpoint()
	}

	var setRevokeCertificateResponseEndpoint kitendpoint.Endpoint
	{
		setRevokeCertificateResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetRevokeCertificateResponse",
			encodeGRPCSetRevokeCertificateResponseRequest,
			decodeGRPCSetRevokeCertificateResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getVerifyDomainRequestEndpoint kitendpoint.Endpoint
	{
		getVerifyDomainRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetVerifyDomainRequest",
			encodeGRPCGetVerifyDomainRequestRequest,
			decodeGRPCGetVerifyDomainRequestResponse,
			apiv1.GetVerifyDomainRequestResponse{},
			options...,
		).Endpoint()
	}

	var setVerifyDomainResponseEndpoint kitendpoint.Endpoint
	{
		setVerifyDomainResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetVerifyDomainResponse",
			encodeGRPCSetVerifyDomainResponseRequest,
			decodeGRPCSetVerifyDomainResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	return endpoint.Endpoints{
		GetAuthorizeDomainRequestEndpoint:  getAuthorizeDomainRequestEndpoint,
		SetAuthorizeDomainResponseEndpoint: setAuthorizeDomainResponseEndpoint,

		GetRequestCertificateRequestEndpoint:  getRequestCertificateRequestEndpoint,
		SetRequestCertificateResponseEndpoint: setRequestCertificateResponseEndpoint,

		GetRevokeCertificateRequestEndpoint:  getRevokeCertificateRequestEndpoint,
		SetRevokeCertificateResponseEndpoint: setRevokeCertificateResponseEndpoint,

		GetVerifyDomainRequestEndpoint:  getVerifyDomainRequestEndpoint,
		SetVerifyDomainResponseEndpoint: setVerifyDomainResponseEndpoint,
	}
}
