package transport // import "powerssl.io/pkg/controller/acme/transport"

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"

	"powerssl.io/pkg/controller/acme/endpoint"
	service "powerssl.io/pkg/controller/acme/service"
	apiv1 "powerssl.io/pkg/controller/api/v1"
)

const serviceName = "powerssl.controller.v1.ACMEService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	options := []grpctransport.ClientOption{}

	var getCreateAccountRequestEndpoint kitendpoint.Endpoint
	{
		getCreateAccountRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetCreateAccountRequest",
			encodeGRPCGetCreateAccountRequestRequest,
			decodeGRPCGetCreateAccountRequestResponse,
			apiv1.GetCreateAccountRequestResponse{},
			options...,
		).Endpoint()
	}

	var setCreateAccountResponseEndpoint kitendpoint.Endpoint
	{
		setCreateAccountResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetCreateAccountResponse",
			encodeGRPCSetCreateAccountResponseRequest,
			decodeGRPCSetCreateAccountResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getDeactivateAccountRequestEndpoint kitendpoint.Endpoint
	{
		getDeactivateAccountRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetDeactivateAccountRequest",
			encodeGRPCGetDeactivateAccountRequestRequest,
			decodeGRPCGetDeactivateAccountRequestResponse,
			apiv1.GetDeactivateAccountRequestResponse{},
			options...,
		).Endpoint()
	}

	var setDeactivateAccountResponseEndpoint kitendpoint.Endpoint
	{
		setDeactivateAccountResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetDeactivateAccountResponse",
			encodeGRPCSetDeactivateAccountResponseRequest,
			decodeGRPCSetDeactivateAccountResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getRekeyAccountRequestEndpoint kitendpoint.Endpoint
	{
		getRekeyAccountRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetRekeyAccountRequest",
			encodeGRPCGetRekeyAccountRequestRequest,
			decodeGRPCGetRekeyAccountRequestResponse,
			apiv1.GetRekeyAccountRequestResponse{},
			options...,
		).Endpoint()
	}

	var setRekeyAccountResponseEndpoint kitendpoint.Endpoint
	{
		setRekeyAccountResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetRekeyAccountResponse",
			encodeGRPCSetRekeyAccountResponseRequest,
			decodeGRPCSetRekeyAccountResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getUpdateAccountRequestEndpoint kitendpoint.Endpoint
	{
		getUpdateAccountRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetUpdateAccountRequest",
			encodeGRPCGetUpdateAccountRequestRequest,
			decodeGRPCGetUpdateAccountRequestResponse,
			apiv1.GetUpdateAccountRequestResponse{},
			options...,
		).Endpoint()
	}

	var setUpdateAccountResponseEndpoint kitendpoint.Endpoint
	{
		setUpdateAccountResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetUpdateAccountResponse",
			encodeGRPCSetUpdateAccountResponseRequest,
			decodeGRPCSetUpdateAccountResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getCreateOrderRequestEndpoint kitendpoint.Endpoint
	{
		getCreateOrderRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetCreateOrderRequest",
			encodeGRPCGetCreateOrderRequestRequest,
			decodeGRPCGetCreateOrderRequestResponse,
			apiv1.GetCreateOrderRequestResponse{},
			options...,
		).Endpoint()
	}

	var setCreateOrderResponseEndpoint kitendpoint.Endpoint
	{
		setCreateOrderResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetCreateOrderResponse",
			encodeGRPCSetCreateOrderResponseRequest,
			decodeGRPCSetCreateOrderResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getFinalizeOrderRequestEndpoint kitendpoint.Endpoint
	{
		getFinalizeOrderRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetFinalizeOrderRequest",
			encodeGRPCGetFinalizeOrderRequestRequest,
			decodeGRPCGetFinalizeOrderRequestResponse,
			apiv1.GetFinalizeOrderRequestResponse{},
			options...,
		).Endpoint()
	}

	var setFinalizeOrderResponseEndpoint kitendpoint.Endpoint
	{
		setFinalizeOrderResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetFinalizeOrderResponse",
			encodeGRPCSetFinalizeOrderResponseRequest,
			decodeGRPCSetFinalizeOrderResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getGetOrderRequestEndpoint kitendpoint.Endpoint
	{
		getGetOrderRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetGetOrderRequest",
			encodeGRPCGetGetOrderRequestRequest,
			decodeGRPCGetGetOrderRequestResponse,
			apiv1.GetGetOrderRequestResponse{},
			options...,
		).Endpoint()
	}

	var setGetOrderResponseEndpoint kitendpoint.Endpoint
	{
		setGetOrderResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetGetOrderResponse",
			encodeGRPCSetGetOrderResponseRequest,
			decodeGRPCSetGetOrderResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getCreateAuthorizationRequestEndpoint kitendpoint.Endpoint
	{
		getCreateAuthorizationRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetCreateAuthorizationRequest",
			encodeGRPCGetCreateAuthorizationRequestRequest,
			decodeGRPCGetCreateAuthorizationRequestResponse,
			apiv1.GetCreateAuthorizationRequestResponse{},
			options...,
		).Endpoint()
	}

	var setCreateAuthorizationResponseEndpoint kitendpoint.Endpoint
	{
		setCreateAuthorizationResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetCreateAuthorizationResponse",
			encodeGRPCSetCreateAuthorizationResponseRequest,
			decodeGRPCSetCreateAuthorizationResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getDeactivateAuthorizationRequestEndpoint kitendpoint.Endpoint
	{
		getDeactivateAuthorizationRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetDeactivateAuthorizationRequest",
			encodeGRPCGetDeactivateAuthorizationRequestRequest,
			decodeGRPCGetDeactivateAuthorizationRequestResponse,
			apiv1.GetDeactivateAuthorizationRequestResponse{},
			options...,
		).Endpoint()
	}

	var setDeactivateAuthorizationResponseEndpoint kitendpoint.Endpoint
	{
		setDeactivateAuthorizationResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetDeactivateAuthorizationResponse",
			encodeGRPCSetDeactivateAuthorizationResponseRequest,
			decodeGRPCSetDeactivateAuthorizationResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getGetAuthorizationRequestEndpoint kitendpoint.Endpoint
	{
		getGetAuthorizationRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetGetAuthorizationRequest",
			encodeGRPCGetGetAuthorizationRequestRequest,
			decodeGRPCGetGetAuthorizationRequestResponse,
			apiv1.GetGetAuthorizationRequestResponse{},
			options...,
		).Endpoint()
	}

	var setGetAuthorizationResponseEndpoint kitendpoint.Endpoint
	{
		setGetAuthorizationResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetGetAuthorizationResponse",
			encodeGRPCSetGetAuthorizationResponseRequest,
			decodeGRPCSetGetAuthorizationResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getGetChallengeRequestEndpoint kitendpoint.Endpoint
	{
		getGetChallengeRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetGetChallengeRequest",
			encodeGRPCGetGetChallengeRequestRequest,
			decodeGRPCGetGetChallengeRequestResponse,
			apiv1.GetGetChallengeRequestResponse{},
			options...,
		).Endpoint()
	}

	var setGetChallengeResponseEndpoint kitendpoint.Endpoint
	{
		setGetChallengeResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetGetChallengeResponse",
			encodeGRPCSetGetChallengeResponseRequest,
			decodeGRPCSetGetChallengeResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getValidateChallengeRequestEndpoint kitendpoint.Endpoint
	{
		getValidateChallengeRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetValidateChallengeRequest",
			encodeGRPCGetValidateChallengeRequestRequest,
			decodeGRPCGetValidateChallengeRequestResponse,
			apiv1.GetValidateChallengeRequestResponse{},
			options...,
		).Endpoint()
	}

	var setValidateChallengeResponseEndpoint kitendpoint.Endpoint
	{
		setValidateChallengeResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetValidateChallengeResponse",
			encodeGRPCSetValidateChallengeResponseRequest,
			decodeGRPCSetValidateChallengeResponseResponse,
			types.Empty{},
			options...,
		).Endpoint()
	}

	var getGetCertificateRequestEndpoint kitendpoint.Endpoint
	{
		getGetCertificateRequestEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"GetGetCertificateRequest",
			encodeGRPCGetGetCertificateRequestRequest,
			decodeGRPCGetGetCertificateRequestResponse,
			apiv1.GetGetCertificateRequestResponse{},
			options...,
		).Endpoint()
	}

	var setGetCertificateResponseEndpoint kitendpoint.Endpoint
	{
		setGetCertificateResponseEndpoint = grpctransport.NewClient(
			conn,
			serviceName,
			"SetGetCertificateResponse",
			encodeGRPCSetGetCertificateResponseRequest,
			decodeGRPCSetGetCertificateResponseResponse,
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

	return endpoint.Endpoints{
		GetCreateAccountRequestEndpoint:  getCreateAccountRequestEndpoint,
		SetCreateAccountResponseEndpoint: setCreateAccountResponseEndpoint,

		GetDeactivateAccountRequestEndpoint:  getDeactivateAccountRequestEndpoint,
		SetDeactivateAccountResponseEndpoint: setDeactivateAccountResponseEndpoint,

		GetRekeyAccountRequestEndpoint:  getRekeyAccountRequestEndpoint,
		SetRekeyAccountResponseEndpoint: setRekeyAccountResponseEndpoint,

		GetUpdateAccountRequestEndpoint:  getUpdateAccountRequestEndpoint,
		SetUpdateAccountResponseEndpoint: setUpdateAccountResponseEndpoint,

		GetCreateOrderRequestEndpoint:  getCreateOrderRequestEndpoint,
		SetCreateOrderResponseEndpoint: setCreateOrderResponseEndpoint,

		GetFinalizeOrderRequestEndpoint:  getFinalizeOrderRequestEndpoint,
		SetFinalizeOrderResponseEndpoint: setFinalizeOrderResponseEndpoint,

		GetGetOrderRequestEndpoint:  getGetOrderRequestEndpoint,
		SetGetOrderResponseEndpoint: setGetOrderResponseEndpoint,

		GetCreateAuthorizationRequestEndpoint:  getCreateAuthorizationRequestEndpoint,
		SetCreateAuthorizationResponseEndpoint: setCreateAuthorizationResponseEndpoint,

		GetDeactivateAuthorizationRequestEndpoint:  getDeactivateAuthorizationRequestEndpoint,
		SetDeactivateAuthorizationResponseEndpoint: setDeactivateAuthorizationResponseEndpoint,

		GetGetAuthorizationRequestEndpoint:  getGetAuthorizationRequestEndpoint,
		SetGetAuthorizationResponseEndpoint: setGetAuthorizationResponseEndpoint,

		GetGetChallengeRequestEndpoint:  getGetChallengeRequestEndpoint,
		SetGetChallengeResponseEndpoint: setGetChallengeResponseEndpoint,

		GetValidateChallengeRequestEndpoint:  getValidateChallengeRequestEndpoint,
		SetValidateChallengeResponseEndpoint: setValidateChallengeResponseEndpoint,

		GetGetCertificateRequestEndpoint:  getGetCertificateRequestEndpoint,
		SetGetCertificateResponseEndpoint: setGetCertificateResponseEndpoint,

		GetRevokeCertificateRequestEndpoint:  getRevokeCertificateRequestEndpoint,
		SetRevokeCertificateResponseEndpoint: setRevokeCertificateResponseEndpoint,
	}
}
