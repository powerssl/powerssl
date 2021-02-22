package acme

import (
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	temporalclient "go.temporal.io/sdk/client"
	"google.golang.org/grpc"

	apiv1 "powerssl.dev/api/controller/v1"
	"powerssl.dev/backend/middleware"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/controller/internal/acme/service"
	"powerssl.dev/sdk/controller/acme"
	"powerssl.dev/sdk/controller/acme/endpoint"
	"powerssl.dev/sdk/controller/acme/transport"
)

func NewService(logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram, temporalClient temporalclient.Client) backendtransport.Service {
	svc := service.New(logger, temporalClient)
	endpoints := makeEndpoints(svc, logger, tracer, duration)

	return backendtransport.Service{
		ServiceName: "powerssl.controller.v1.ACMEService",
		RegisterGRPCServer: func(baseServer *grpc.Server) {
			grpcServer := transport.NewGRPCServer(endpoints, logger, tracer)
			apiv1.RegisterACMEServiceServer(baseServer, grpcServer)
		},
	}
}

func makeEndpoints(svc acme.Service, logger log.Logger, tracer stdopentracing.Tracer, duration metrics.Histogram) endpoint.Endpoints {
	var getCreateAccountRequestEndpoint kitendpoint.Endpoint
	{
		getCreateAccountRequestEndpoint = endpoint.MakeGetCreateAccountRequestEndpoint(svc)
		getCreateAccountRequestEndpoint = opentracing.TraceServer(tracer, "GetCreateAccountRequest")(getCreateAccountRequestEndpoint)
		getCreateAccountRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetCreateAccountRequest"))(getCreateAccountRequestEndpoint)
		getCreateAccountRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetCreateAccountRequest"))(getCreateAccountRequestEndpoint)
	}

	var setCreateAccountResponseEndpoint kitendpoint.Endpoint
	{
		setCreateAccountResponseEndpoint = endpoint.MakeSetCreateAccountResponseEndpoint(svc)
		setCreateAccountResponseEndpoint = opentracing.TraceServer(tracer, "SetCreateAccountResponse")(setCreateAccountResponseEndpoint)
		setCreateAccountResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetCreateAccountResponse"))(setCreateAccountResponseEndpoint)
		setCreateAccountResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetCreateAccountResponse"))(setCreateAccountResponseEndpoint)
	}

	var getDeactivateAccountRequestEndpoint kitendpoint.Endpoint
	{
		getDeactivateAccountRequestEndpoint = endpoint.MakeGetDeactivateAccountRequestEndpoint(svc)
		getDeactivateAccountRequestEndpoint = opentracing.TraceServer(tracer, "GetDeactivateAccountRequest")(getDeactivateAccountRequestEndpoint)
		getDeactivateAccountRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetDeactivateAccountRequest"))(getDeactivateAccountRequestEndpoint)
		getDeactivateAccountRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetDeactivateAccountRequest"))(getDeactivateAccountRequestEndpoint)
	}

	var setDeactivateAccountResponseEndpoint kitendpoint.Endpoint
	{
		setDeactivateAccountResponseEndpoint = endpoint.MakeSetDeactivateAccountResponseEndpoint(svc)
		setDeactivateAccountResponseEndpoint = opentracing.TraceServer(tracer, "SetDeactivateAccountResponse")(setDeactivateAccountResponseEndpoint)
		setDeactivateAccountResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetDeactivateAccountResponse"))(setDeactivateAccountResponseEndpoint)
		setDeactivateAccountResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetDeactivateAccountResponse"))(setDeactivateAccountResponseEndpoint)
	}

	var getRekeyAccountRequestEndpoint kitendpoint.Endpoint
	{
		getRekeyAccountRequestEndpoint = endpoint.MakeGetRekeyAccountRequestEndpoint(svc)
		getRekeyAccountRequestEndpoint = opentracing.TraceServer(tracer, "GetRekeyAccountRequest")(getRekeyAccountRequestEndpoint)
		getRekeyAccountRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetRekeyAccountRequest"))(getRekeyAccountRequestEndpoint)
		getRekeyAccountRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetRekeyAccountRequest"))(getRekeyAccountRequestEndpoint)
	}

	var setRekeyAccountResponseEndpoint kitendpoint.Endpoint
	{
		setRekeyAccountResponseEndpoint = endpoint.MakeSetRekeyAccountResponseEndpoint(svc)
		setRekeyAccountResponseEndpoint = opentracing.TraceServer(tracer, "SetRekeyAccountResponse")(setRekeyAccountResponseEndpoint)
		setRekeyAccountResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetRekeyAccountResponse"))(setRekeyAccountResponseEndpoint)
		setRekeyAccountResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetRekeyAccountResponse"))(setRekeyAccountResponseEndpoint)
	}

	var getUpdateAccountRequestEndpoint kitendpoint.Endpoint
	{
		getUpdateAccountRequestEndpoint = endpoint.MakeGetUpdateAccountRequestEndpoint(svc)
		getUpdateAccountRequestEndpoint = opentracing.TraceServer(tracer, "GetUpdateAccountRequest")(getUpdateAccountRequestEndpoint)
		getUpdateAccountRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetUpdateAccountRequest"))(getUpdateAccountRequestEndpoint)
		getUpdateAccountRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetUpdateAccountRequest"))(getUpdateAccountRequestEndpoint)
	}

	var setUpdateAccountResponseEndpoint kitendpoint.Endpoint
	{
		setUpdateAccountResponseEndpoint = endpoint.MakeSetUpdateAccountResponseEndpoint(svc)
		setUpdateAccountResponseEndpoint = opentracing.TraceServer(tracer, "SetUpdateAccountResponse")(setUpdateAccountResponseEndpoint)
		setUpdateAccountResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetUpdateAccountResponse"))(setUpdateAccountResponseEndpoint)
		setUpdateAccountResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetUpdateAccountResponse"))(setUpdateAccountResponseEndpoint)
	}

	var getCreateOrderRequestEndpoint kitendpoint.Endpoint
	{
		getCreateOrderRequestEndpoint = endpoint.MakeGetCreateOrderRequestEndpoint(svc)
		getCreateOrderRequestEndpoint = opentracing.TraceServer(tracer, "GetCreateOrderRequest")(getCreateOrderRequestEndpoint)
		getCreateOrderRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetCreateOrderRequest"))(getCreateOrderRequestEndpoint)
		getCreateOrderRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetCreateOrderRequest"))(getCreateOrderRequestEndpoint)
	}

	var setCreateOrderResponseEndpoint kitendpoint.Endpoint
	{
		setCreateOrderResponseEndpoint = endpoint.MakeSetCreateOrderResponseEndpoint(svc)
		setCreateOrderResponseEndpoint = opentracing.TraceServer(tracer, "SetCreateOrderResponse")(setCreateOrderResponseEndpoint)
		setCreateOrderResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetCreateOrderResponse"))(setCreateOrderResponseEndpoint)
		setCreateOrderResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetCreateOrderResponse"))(setCreateOrderResponseEndpoint)
	}

	var getFinalizeOrderRequestEndpoint kitendpoint.Endpoint
	{
		getFinalizeOrderRequestEndpoint = endpoint.MakeGetFinalizeOrderRequestEndpoint(svc)
		getFinalizeOrderRequestEndpoint = opentracing.TraceServer(tracer, "GetFinalizeOrderRequest")(getFinalizeOrderRequestEndpoint)
		getFinalizeOrderRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetFinalizeOrderRequest"))(getFinalizeOrderRequestEndpoint)
		getFinalizeOrderRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetFinalizeOrderRequest"))(getFinalizeOrderRequestEndpoint)
	}

	var setFinalizeOrderResponseEndpoint kitendpoint.Endpoint
	{
		setFinalizeOrderResponseEndpoint = endpoint.MakeSetFinalizeOrderResponseEndpoint(svc)
		setFinalizeOrderResponseEndpoint = opentracing.TraceServer(tracer, "SetFinalizeOrderResponse")(setFinalizeOrderResponseEndpoint)
		setFinalizeOrderResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetFinalizeOrderResponse"))(setFinalizeOrderResponseEndpoint)
		setFinalizeOrderResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetFinalizeOrderResponse"))(setFinalizeOrderResponseEndpoint)
	}

	var getGetOrderRequestEndpoint kitendpoint.Endpoint
	{
		getGetOrderRequestEndpoint = endpoint.MakeGetGetOrderRequestEndpoint(svc)
		getGetOrderRequestEndpoint = opentracing.TraceServer(tracer, "GetGetOrderRequest")(getGetOrderRequestEndpoint)
		getGetOrderRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetGetOrderRequest"))(getGetOrderRequestEndpoint)
		getGetOrderRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetGetOrderRequest"))(getGetOrderRequestEndpoint)
	}

	var setGetOrderResponseEndpoint kitendpoint.Endpoint
	{
		setGetOrderResponseEndpoint = endpoint.MakeSetGetOrderResponseEndpoint(svc)
		setGetOrderResponseEndpoint = opentracing.TraceServer(tracer, "SetGetOrderResponse")(setGetOrderResponseEndpoint)
		setGetOrderResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetGetOrderResponse"))(setGetOrderResponseEndpoint)
		setGetOrderResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetGetOrderResponse"))(setGetOrderResponseEndpoint)
	}

	var getCreateAuthorizationRequestEndpoint kitendpoint.Endpoint
	{
		getCreateAuthorizationRequestEndpoint = endpoint.MakeGetCreateAuthorizationRequestEndpoint(svc)
		getCreateAuthorizationRequestEndpoint = opentracing.TraceServer(tracer, "GetCreateAuthorizationRequest")(getCreateAuthorizationRequestEndpoint)
		getCreateAuthorizationRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetCreateAuthorizationRequest"))(getCreateAuthorizationRequestEndpoint)
		getCreateAuthorizationRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetCreateAuthorizationRequest"))(getCreateAuthorizationRequestEndpoint)
	}

	var setCreateAuthorizationResponseEndpoint kitendpoint.Endpoint
	{
		setCreateAuthorizationResponseEndpoint = endpoint.MakeSetCreateAuthorizationResponseEndpoint(svc)
		setCreateAuthorizationResponseEndpoint = opentracing.TraceServer(tracer, "SetCreateAuthorizationResponse")(setCreateAuthorizationResponseEndpoint)
		setCreateAuthorizationResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetCreateAuthorizationResponse"))(setCreateAuthorizationResponseEndpoint)
		setCreateAuthorizationResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetCreateAuthorizationResponse"))(setCreateAuthorizationResponseEndpoint)
	}

	var getDeactivateAuthorizationRequestEndpoint kitendpoint.Endpoint
	{
		getDeactivateAuthorizationRequestEndpoint = endpoint.MakeGetDeactivateAuthorizationRequestEndpoint(svc)
		getDeactivateAuthorizationRequestEndpoint = opentracing.TraceServer(tracer, "GetDeactivateAuthorizationRequest")(getDeactivateAuthorizationRequestEndpoint)
		getDeactivateAuthorizationRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetDeactivateAuthorizationRequest"))(getDeactivateAuthorizationRequestEndpoint)
		getDeactivateAuthorizationRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetDeactivateAuthorizationRequest"))(getDeactivateAuthorizationRequestEndpoint)
	}

	var setDeactivateAuthorizationResponseEndpoint kitendpoint.Endpoint
	{
		setDeactivateAuthorizationResponseEndpoint = endpoint.MakeSetDeactivateAuthorizationResponseEndpoint(svc)
		setDeactivateAuthorizationResponseEndpoint = opentracing.TraceServer(tracer, "SetDeactivateAuthorizationResponse")(setDeactivateAuthorizationResponseEndpoint)
		setDeactivateAuthorizationResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetDeactivateAuthorizationResponse"))(setDeactivateAuthorizationResponseEndpoint)
		setDeactivateAuthorizationResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetDeactivateAuthorizationResponse"))(setDeactivateAuthorizationResponseEndpoint)
	}

	var getGetAuthorizationRequestEndpoint kitendpoint.Endpoint
	{
		getGetAuthorizationRequestEndpoint = endpoint.MakeGetGetAuthorizationRequestEndpoint(svc)
		getGetAuthorizationRequestEndpoint = opentracing.TraceServer(tracer, "GetGetAuthorizationRequest")(getGetAuthorizationRequestEndpoint)
		getGetAuthorizationRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetGetAuthorizationRequest"))(getGetAuthorizationRequestEndpoint)
		getGetAuthorizationRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetGetAuthorizationRequest"))(getGetAuthorizationRequestEndpoint)
	}

	var setGetAuthorizationResponseEndpoint kitendpoint.Endpoint
	{
		setGetAuthorizationResponseEndpoint = endpoint.MakeSetGetAuthorizationResponseEndpoint(svc)
		setGetAuthorizationResponseEndpoint = opentracing.TraceServer(tracer, "SetGetAuthorizationResponse")(setGetAuthorizationResponseEndpoint)
		setGetAuthorizationResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetGetAuthorizationResponse"))(setGetAuthorizationResponseEndpoint)
		setGetAuthorizationResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetGetAuthorizationResponse"))(setGetAuthorizationResponseEndpoint)
	}

	var getGetChallengeRequestEndpoint kitendpoint.Endpoint
	{
		getGetChallengeRequestEndpoint = endpoint.MakeGetGetChallengeRequestEndpoint(svc)
		getGetChallengeRequestEndpoint = opentracing.TraceServer(tracer, "GetGetChallengeRequest")(getGetChallengeRequestEndpoint)
		getGetChallengeRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetGetChallengeRequest"))(getGetChallengeRequestEndpoint)
		getGetChallengeRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetGetChallengeRequest"))(getGetChallengeRequestEndpoint)
	}

	var setGetChallengeResponseEndpoint kitendpoint.Endpoint
	{
		setGetChallengeResponseEndpoint = endpoint.MakeSetGetChallengeResponseEndpoint(svc)
		setGetChallengeResponseEndpoint = opentracing.TraceServer(tracer, "SetGetChallengeResponse")(setGetChallengeResponseEndpoint)
		setGetChallengeResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetGetChallengeResponse"))(setGetChallengeResponseEndpoint)
		setGetChallengeResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetGetChallengeResponse"))(setGetChallengeResponseEndpoint)
	}

	var getValidateChallengeRequestEndpoint kitendpoint.Endpoint
	{
		getValidateChallengeRequestEndpoint = endpoint.MakeGetValidateChallengeRequestEndpoint(svc)
		getValidateChallengeRequestEndpoint = opentracing.TraceServer(tracer, "GetValidateChallengeRequest")(getValidateChallengeRequestEndpoint)
		getValidateChallengeRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetValidateChallengeRequest"))(getValidateChallengeRequestEndpoint)
		getValidateChallengeRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetValidateChallengeRequest"))(getValidateChallengeRequestEndpoint)
	}

	var setValidateChallengeResponseEndpoint kitendpoint.Endpoint
	{
		setValidateChallengeResponseEndpoint = endpoint.MakeSetValidateChallengeResponseEndpoint(svc)
		setValidateChallengeResponseEndpoint = opentracing.TraceServer(tracer, "SetValidateChallengeResponse")(setValidateChallengeResponseEndpoint)
		setValidateChallengeResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetValidateChallengeResponse"))(setValidateChallengeResponseEndpoint)
		setValidateChallengeResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetValidateChallengeResponse"))(setValidateChallengeResponseEndpoint)
	}

	var getGetCertificateRequestEndpoint kitendpoint.Endpoint
	{
		getGetCertificateRequestEndpoint = endpoint.MakeGetGetCertificateRequestEndpoint(svc)
		getGetCertificateRequestEndpoint = opentracing.TraceServer(tracer, "GetGetCertificateRequest")(getGetCertificateRequestEndpoint)
		getGetCertificateRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetGetCertificateRequest"))(getGetCertificateRequestEndpoint)
		getGetCertificateRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetGetCertificateRequest"))(getGetCertificateRequestEndpoint)
	}

	var setGetCertificateResponseEndpoint kitendpoint.Endpoint
	{
		setGetCertificateResponseEndpoint = endpoint.MakeSetGetCertificateResponseEndpoint(svc)
		setGetCertificateResponseEndpoint = opentracing.TraceServer(tracer, "SetGetCertificateResponse")(setGetCertificateResponseEndpoint)
		setGetCertificateResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetGetCertificateResponse"))(setGetCertificateResponseEndpoint)
		setGetCertificateResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetGetCertificateResponse"))(setGetCertificateResponseEndpoint)
	}

	var getRevokeCertificateRequestEndpoint kitendpoint.Endpoint
	{
		getRevokeCertificateRequestEndpoint = endpoint.MakeGetRevokeCertificateRequestEndpoint(svc)
		getRevokeCertificateRequestEndpoint = opentracing.TraceServer(tracer, "GetRevokeCertificateRequest")(getRevokeCertificateRequestEndpoint)
		getRevokeCertificateRequestEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "GetRevokeCertificateRequest"))(getRevokeCertificateRequestEndpoint)
		getRevokeCertificateRequestEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "GetRevokeCertificateRequest"))(getRevokeCertificateRequestEndpoint)
	}

	var setRevokeCertificateResponseEndpoint kitendpoint.Endpoint
	{
		setRevokeCertificateResponseEndpoint = endpoint.MakeSetRevokeCertificateResponseEndpoint(svc)
		setRevokeCertificateResponseEndpoint = opentracing.TraceServer(tracer, "SetRevokeCertificateResponse")(setRevokeCertificateResponseEndpoint)
		setRevokeCertificateResponseEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "SetRevokeCertificateResponse"))(setRevokeCertificateResponseEndpoint)
		setRevokeCertificateResponseEndpoint = middleware.InstrumentingMiddleware(duration.With("method", "SetRevokeCertificateResponse"))(setRevokeCertificateResponseEndpoint)
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
