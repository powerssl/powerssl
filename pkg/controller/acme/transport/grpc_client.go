package transport // import "powerssl.io/pkg/controller/acme/transport"

import (
	"fmt"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"powerssl.io/pkg/controller/acme/endpoint"
	service "powerssl.io/pkg/controller/acme/service"
	apiv1 "powerssl.io/pkg/controller/api/v1"
)

const serviceName = "powerssl.controller.v1.ACMEService"

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger, tracer stdopentracing.Tracer) service.Service {
	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	}

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
		getCreateAccountRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetCreateAccountRequest", serviceName))(getCreateAccountRequestEndpoint)
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
		setCreateAccountResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetCreateAccountResponse", serviceName))(setCreateAccountResponseEndpoint)
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
		getDeactivateAccountRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetDeactivateAccountRequest", serviceName))(getDeactivateAccountRequestEndpoint)
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
		setDeactivateAccountResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetDeactivateAccountResponse", serviceName))(setDeactivateAccountResponseEndpoint)
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
		getRekeyAccountRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetRekeyAccountRequest", serviceName))(getRekeyAccountRequestEndpoint)
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
		setRekeyAccountResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetRekeyAccountResponse", serviceName))(setRekeyAccountResponseEndpoint)
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
		getUpdateAccountRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetUpdateAccountRequest", serviceName))(getUpdateAccountRequestEndpoint)
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
		setUpdateAccountResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetUpdateAccountResponse", serviceName))(setUpdateAccountResponseEndpoint)
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
		getCreateOrderRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetCreateOrderRequest", serviceName))(getCreateOrderRequestEndpoint)
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
		setCreateOrderResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetCreateOrderResponse", serviceName))(setCreateOrderResponseEndpoint)
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
		getFinalizeOrderRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetFinalizeOrderRequest", serviceName))(getFinalizeOrderRequestEndpoint)
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
		setFinalizeOrderResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetFinalizeOrderResponse", serviceName))(setFinalizeOrderResponseEndpoint)
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
		getGetOrderRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetGetOrderRequest", serviceName))(getGetOrderRequestEndpoint)
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
		setGetOrderResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetGetOrderResponse", serviceName))(setGetOrderResponseEndpoint)
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
		getCreateAuthorizationRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetCreateAuthorizationRequest", serviceName))(getCreateAuthorizationRequestEndpoint)
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
		setCreateAuthorizationResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetCreateAuthorizationResponse", serviceName))(setCreateAuthorizationResponseEndpoint)
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
		getDeactivateAuthorizationRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetDeactivateAuthorizationRequest", serviceName))(getDeactivateAuthorizationRequestEndpoint)
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
		setDeactivateAuthorizationResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetDeactivateAuthorizationResponse", serviceName))(setDeactivateAuthorizationResponseEndpoint)
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
		getGetAuthorizationRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetGetAuthorizationRequest", serviceName))(getGetAuthorizationRequestEndpoint)
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
		setGetAuthorizationResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetGetAuthorizationResponse", serviceName))(setGetAuthorizationResponseEndpoint)
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
		getGetChallengeRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetGetChallengeRequest", serviceName))(getGetChallengeRequestEndpoint)
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
		setGetChallengeResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetGetChallengeResponse", serviceName))(setGetChallengeResponseEndpoint)
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
		getValidateChallengeRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetValidateChallengeRequest", serviceName))(getValidateChallengeRequestEndpoint)
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
		setValidateChallengeResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetValidateChallengeResponse", serviceName))(setValidateChallengeResponseEndpoint)
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
		getGetCertificateRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetGetCertificateRequest", serviceName))(getGetCertificateRequestEndpoint)
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
		setGetCertificateResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetGetCertificateResponse", serviceName))(setGetCertificateResponseEndpoint)
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
		getRevokeCertificateRequestEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/GetRevokeCertificateRequest", serviceName))(getRevokeCertificateRequestEndpoint)
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
		setRevokeCertificateResponseEndpoint = opentracing.TraceClient(tracer, fmt.Sprintf("/%s/SetRevokeCertificateResponse", serviceName))(setRevokeCertificateResponseEndpoint)
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
