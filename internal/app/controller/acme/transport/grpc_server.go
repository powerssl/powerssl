package transport

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gogo/protobuf/types"
	stdopentracing "github.com/opentracing/opentracing-go"

	"powerssl.io/internal/app/controller/acme/endpoint"
	apiv1 "powerssl.io/internal/pkg/controller/api/v1"
)

type grpcServer struct {
	getCreateAccountRequest  grpctransport.Handler
	setCreateAccountResponse grpctransport.Handler

	getDeactivateAccountRequest  grpctransport.Handler
	setDeactivateAccountResponse grpctransport.Handler

	getRekeyAccountRequest  grpctransport.Handler
	setRekeyAccountResponse grpctransport.Handler

	getUpdateAccountRequest  grpctransport.Handler
	setUpdateAccountResponse grpctransport.Handler

	getCreateOrderRequest  grpctransport.Handler
	setCreateOrderResponse grpctransport.Handler

	getFinalizeOrderRequest  grpctransport.Handler
	setFinalizeOrderResponse grpctransport.Handler

	getGetOrderRequest  grpctransport.Handler
	setGetOrderResponse grpctransport.Handler

	getCreateAuthorizationRequest  grpctransport.Handler
	setCreateAuthorizationResponse grpctransport.Handler

	getDeactivateAuthorizationRequest  grpctransport.Handler
	setDeactivateAuthorizationResponse grpctransport.Handler

	getGetAuthorizationRequest  grpctransport.Handler
	setGetAuthorizationResponse grpctransport.Handler

	getGetChallengeRequest  grpctransport.Handler
	setGetChallengeResponse grpctransport.Handler

	getValidateChallengeRequest  grpctransport.Handler
	setValidateChallengeResponse grpctransport.Handler

	getGetCertificateRequest  grpctransport.Handler
	setGetCertificateResponse grpctransport.Handler

	getRevokeCertificateRequest  grpctransport.Handler
	setRevokeCertificateResponse grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger, tracer stdopentracing.Tracer) apiv1.ACMEServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		getCreateAccountRequest: grpctransport.NewServer(
			endpoints.GetCreateAccountRequestEndpoint,
			decodeGRPCGetCreateAccountRequestRequest,
			encodeGRPCGetCreateAccountRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetCreateAccountRequest", serviceName), logger)))...,
		),
		setCreateAccountResponse: grpctransport.NewServer(
			endpoints.SetCreateAccountResponseEndpoint,
			decodeGRPCSetCreateAccountResponseRequest,
			encodeGRPCSetCreateAccountResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetCreateAccountResponse", serviceName), logger)))...,
		),

		getDeactivateAccountRequest: grpctransport.NewServer(
			endpoints.GetDeactivateAccountRequestEndpoint,
			decodeGRPCGetDeactivateAccountRequestRequest,
			encodeGRPCGetDeactivateAccountRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetDeactivateAccountRequest", serviceName), logger)))...,
		),
		setDeactivateAccountResponse: grpctransport.NewServer(
			endpoints.SetDeactivateAccountResponseEndpoint,
			decodeGRPCSetDeactivateAccountResponseRequest,
			encodeGRPCSetDeactivateAccountResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetDeactivateAccountResponse", serviceName), logger)))...,
		),

		getRekeyAccountRequest: grpctransport.NewServer(
			endpoints.GetRekeyAccountRequestEndpoint,
			decodeGRPCGetRekeyAccountRequestRequest,
			encodeGRPCGetRekeyAccountRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetRekeyAccountRequest", serviceName), logger)))...,
		),
		setRekeyAccountResponse: grpctransport.NewServer(
			endpoints.SetRekeyAccountResponseEndpoint,
			decodeGRPCSetRekeyAccountResponseRequest,
			encodeGRPCSetRekeyAccountResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetRekeyAccountResponse", serviceName), logger)))...,
		),

		getUpdateAccountRequest: grpctransport.NewServer(
			endpoints.GetUpdateAccountRequestEndpoint,
			decodeGRPCGetUpdateAccountRequestRequest,
			encodeGRPCGetUpdateAccountRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetUpdateAccountRequest", serviceName), logger)))...,
		),
		setUpdateAccountResponse: grpctransport.NewServer(
			endpoints.SetUpdateAccountResponseEndpoint,
			decodeGRPCSetUpdateAccountResponseRequest,
			encodeGRPCSetUpdateAccountResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetUpdateAccountResponse", serviceName), logger)))...,
		),

		getCreateOrderRequest: grpctransport.NewServer(
			endpoints.GetCreateOrderRequestEndpoint,
			decodeGRPCGetCreateOrderRequestRequest,
			encodeGRPCGetCreateOrderRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetCreateOrderRequest", serviceName), logger)))...,
		),
		setCreateOrderResponse: grpctransport.NewServer(
			endpoints.SetCreateOrderResponseEndpoint,
			decodeGRPCSetCreateOrderResponseRequest,
			encodeGRPCSetCreateOrderResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetCreateOrderResponse", serviceName), logger)))...,
		),

		getFinalizeOrderRequest: grpctransport.NewServer(
			endpoints.GetFinalizeOrderRequestEndpoint,
			decodeGRPCGetFinalizeOrderRequestRequest,
			encodeGRPCGetFinalizeOrderRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetFinalizeOrderRequest", serviceName), logger)))...,
		),
		setFinalizeOrderResponse: grpctransport.NewServer(
			endpoints.SetFinalizeOrderResponseEndpoint,
			decodeGRPCSetFinalizeOrderResponseRequest,
			encodeGRPCSetFinalizeOrderResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetFinalizeOrderResponse", serviceName), logger)))...,
		),

		getGetOrderRequest: grpctransport.NewServer(
			endpoints.GetGetOrderRequestEndpoint,
			decodeGRPCGetGetOrderRequestRequest,
			encodeGRPCGetGetOrderRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetGetOrderRequest", serviceName), logger)))...,
		),
		setGetOrderResponse: grpctransport.NewServer(
			endpoints.SetGetOrderResponseEndpoint,
			decodeGRPCSetGetOrderResponseRequest,
			encodeGRPCSetGetOrderResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetGetOrderResponse", serviceName), logger)))...,
		),

		getCreateAuthorizationRequest: grpctransport.NewServer(
			endpoints.GetCreateAuthorizationRequestEndpoint,
			decodeGRPCGetCreateAuthorizationRequestRequest,
			encodeGRPCGetCreateAuthorizationRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetCreateAuthorizationRequest", serviceName), logger)))...,
		),
		setCreateAuthorizationResponse: grpctransport.NewServer(
			endpoints.SetCreateAuthorizationResponseEndpoint,
			decodeGRPCSetCreateAuthorizationResponseRequest,
			encodeGRPCSetCreateAuthorizationResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetCreateAuthorizationResponse", serviceName), logger)))...,
		),

		getDeactivateAuthorizationRequest: grpctransport.NewServer(
			endpoints.GetDeactivateAuthorizationRequestEndpoint,
			decodeGRPCGetDeactivateAuthorizationRequestRequest,
			encodeGRPCGetDeactivateAuthorizationRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetDeactivateAuthorizationRequest", serviceName), logger)))...,
		),
		setDeactivateAuthorizationResponse: grpctransport.NewServer(
			endpoints.SetDeactivateAuthorizationResponseEndpoint,
			decodeGRPCSetDeactivateAuthorizationResponseRequest,
			encodeGRPCSetDeactivateAuthorizationResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetDeactivateAuthorizationResponse", serviceName), logger)))...,
		),

		getGetAuthorizationRequest: grpctransport.NewServer(
			endpoints.GetGetAuthorizationRequestEndpoint,
			decodeGRPCGetGetAuthorizationRequestRequest,
			encodeGRPCGetGetAuthorizationRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetGetAuthorizationRequest", serviceName), logger)))...,
		),
		setGetAuthorizationResponse: grpctransport.NewServer(
			endpoints.SetGetAuthorizationResponseEndpoint,
			decodeGRPCSetGetAuthorizationResponseRequest,
			encodeGRPCSetGetAuthorizationResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetGetAuthorizationResponse", serviceName), logger)))...,
		),

		getGetChallengeRequest: grpctransport.NewServer(
			endpoints.GetGetChallengeRequestEndpoint,
			decodeGRPCGetGetChallengeRequestRequest,
			encodeGRPCGetGetChallengeRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetGetChallengeRequest", serviceName), logger)))...,
		),
		setGetChallengeResponse: grpctransport.NewServer(
			endpoints.SetGetChallengeResponseEndpoint,
			decodeGRPCSetGetChallengeResponseRequest,
			encodeGRPCSetGetChallengeResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetGetChallengeResponse", serviceName), logger)))...,
		),

		getValidateChallengeRequest: grpctransport.NewServer(
			endpoints.GetValidateChallengeRequestEndpoint,
			decodeGRPCGetValidateChallengeRequestRequest,
			encodeGRPCGetValidateChallengeRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetValidateChallengeRequest", serviceName), logger)))...,
		),
		setValidateChallengeResponse: grpctransport.NewServer(
			endpoints.SetValidateChallengeResponseEndpoint,
			decodeGRPCSetValidateChallengeResponseRequest,
			encodeGRPCSetValidateChallengeResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetValidateChallengeResponse", serviceName), logger)))...,
		),

		getGetCertificateRequest: grpctransport.NewServer(
			endpoints.GetGetCertificateRequestEndpoint,
			decodeGRPCGetGetCertificateRequestRequest,
			encodeGRPCGetGetCertificateRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetGetCertificateRequest", serviceName), logger)))...,
		),
		setGetCertificateResponse: grpctransport.NewServer(
			endpoints.SetGetCertificateResponseEndpoint,
			decodeGRPCSetGetCertificateResponseRequest,
			encodeGRPCSetGetCertificateResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetGetCertificateResponse", serviceName), logger)))...,
		),

		getRevokeCertificateRequest: grpctransport.NewServer(
			endpoints.GetRevokeCertificateRequestEndpoint,
			decodeGRPCGetRevokeCertificateRequestRequest,
			encodeGRPCGetRevokeCertificateRequestResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/GetRevokeCertificateRequest", serviceName), logger)))...,
		),
		setRevokeCertificateResponse: grpctransport.NewServer(
			endpoints.SetRevokeCertificateResponseEndpoint,
			decodeGRPCSetRevokeCertificateResponseRequest,
			encodeGRPCSetRevokeCertificateResponseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, fmt.Sprintf("/%s/SetRevokeCertificateResponse", serviceName), logger)))...,
		),
	}
}

func (s *grpcServer) GetCreateAccountRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetCreateAccountRequestResponse, error) {
	_, rep, err := s.getCreateAccountRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetCreateAccountRequestResponse), nil
}

func (s *grpcServer) SetCreateAccountResponse(ctx context.Context, req *apiv1.SetCreateAccountResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setCreateAccountResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetDeactivateAccountRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetDeactivateAccountRequestResponse, error) {
	_, rep, err := s.getDeactivateAccountRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetDeactivateAccountRequestResponse), nil
}

func (s *grpcServer) SetDeactivateAccountResponse(ctx context.Context, req *apiv1.SetDeactivateAccountResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setDeactivateAccountResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetRekeyAccountRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetRekeyAccountRequestResponse, error) {
	_, rep, err := s.getRekeyAccountRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetRekeyAccountRequestResponse), nil
}

func (s *grpcServer) SetRekeyAccountResponse(ctx context.Context, req *apiv1.SetRekeyAccountResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setRekeyAccountResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetUpdateAccountRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetUpdateAccountRequestResponse, error) {
	_, rep, err := s.getUpdateAccountRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetUpdateAccountRequestResponse), nil
}

func (s *grpcServer) SetUpdateAccountResponse(ctx context.Context, req *apiv1.SetUpdateAccountResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setUpdateAccountResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetCreateOrderRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetCreateOrderRequestResponse, error) {
	_, rep, err := s.getCreateOrderRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetCreateOrderRequestResponse), nil
}

func (s *grpcServer) SetCreateOrderResponse(ctx context.Context, req *apiv1.SetCreateOrderResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setCreateOrderResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetFinalizeOrderRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetFinalizeOrderRequestResponse, error) {
	_, rep, err := s.getFinalizeOrderRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetFinalizeOrderRequestResponse), nil
}

func (s *grpcServer) SetFinalizeOrderResponse(ctx context.Context, req *apiv1.SetFinalizeOrderResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setFinalizeOrderResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetGetOrderRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetGetOrderRequestResponse, error) {
	_, rep, err := s.getGetOrderRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetGetOrderRequestResponse), nil
}

func (s *grpcServer) SetGetOrderResponse(ctx context.Context, req *apiv1.SetGetOrderResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setGetOrderResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetCreateAuthorizationRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetCreateAuthorizationRequestResponse, error) {
	_, rep, err := s.getCreateAuthorizationRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetCreateAuthorizationRequestResponse), nil
}

func (s *grpcServer) SetCreateAuthorizationResponse(ctx context.Context, req *apiv1.SetCreateAuthorizationResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setCreateAuthorizationResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetDeactivateAuthorizationRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetDeactivateAuthorizationRequestResponse, error) {
	_, rep, err := s.getDeactivateAuthorizationRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetDeactivateAuthorizationRequestResponse), nil
}

func (s *grpcServer) SetDeactivateAuthorizationResponse(ctx context.Context, req *apiv1.SetDeactivateAuthorizationResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setDeactivateAuthorizationResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetGetAuthorizationRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetGetAuthorizationRequestResponse, error) {
	_, rep, err := s.getGetAuthorizationRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetGetAuthorizationRequestResponse), nil
}

func (s *grpcServer) SetGetAuthorizationResponse(ctx context.Context, req *apiv1.SetGetAuthorizationResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setGetAuthorizationResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetGetChallengeRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetGetChallengeRequestResponse, error) {
	_, rep, err := s.getGetChallengeRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetGetChallengeRequestResponse), nil
}

func (s *grpcServer) SetGetChallengeResponse(ctx context.Context, req *apiv1.SetGetChallengeResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setGetChallengeResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetValidateChallengeRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetValidateChallengeRequestResponse, error) {
	_, rep, err := s.getValidateChallengeRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetValidateChallengeRequestResponse), nil
}

func (s *grpcServer) SetValidateChallengeResponse(ctx context.Context, req *apiv1.SetValidateChallengeResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setValidateChallengeResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetGetCertificateRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetGetCertificateRequestResponse, error) {
	_, rep, err := s.getGetCertificateRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetGetCertificateRequestResponse), nil
}

func (s *grpcServer) SetGetCertificateResponse(ctx context.Context, req *apiv1.SetGetCertificateResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setGetCertificateResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetRevokeCertificateRequest(ctx context.Context, req *apiv1.Activity) (*apiv1.GetRevokeCertificateRequestResponse, error) {
	_, rep, err := s.getRevokeCertificateRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetRevokeCertificateRequestResponse), nil
}

func (s *grpcServer) SetRevokeCertificateResponse(ctx context.Context, req *apiv1.SetRevokeCertificateResponseRequest) (*types.Empty, error) {
	_, rep, err := s.setRevokeCertificateResponse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}
