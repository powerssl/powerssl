package acme

import (
	"context"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "powerssl.dev/api/controller/v1"
	"powerssl.dev/common/log"
	"powerssl.dev/workflow/activity"

	integrationactivity "powerssl.dev/controller/internal/service/integration/activity"
)

var ServiceDesc = &apiv1.ACMEService_ServiceDesc

type Service struct {
	apiv1.UnimplementedACMEServiceServer
	logger   log.Logger
	temporal client.Client
}

func New(logger log.Logger, temporalClient client.Client) *Service {
	return &Service{
		logger:   logger,
		temporal: temporalClient,
	}
}

func appErr(err *apiv1.Error) error {
	if err.GetMessage() != "" {
		return temporal.NewApplicationError(err.GetMessage(), "")
	}
	return nil
}

func (s Service) integrationErr(err error) error {
	s.logger.Error(err)
	return status.Error(codes.Internal, "")
}

func (s Service) GetCreateAccountRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetCreateAccountRequestResponse, error) {
	var input *CreateACMEAccountParams
	if err := integrationactivity.GetInput(ctx, activity, &input); err != nil {
		return nil, s.integrationErr(err)
	}
	return &apiv1.GetCreateAccountRequestResponse{
		Activity:             activity,
		KeyToken:             input.KeyToken,
		DirectoryUrl:         input.DirectoryURL,
		TermsOfServiceAgreed: input.TermsOfServiceAgreed,
		Contacts:             input.Contacts,
	}, nil
}

func (s Service) SetCreateAccountResponse(ctx context.Context, request *apiv1.SetCreateAccountResponseRequest) (*emptypb.Empty, error) {
	result := &activity.CreateACMEAccountResults{
		Account: request.GetAccount(),
	}
	if err := integrationactivity.SetResult(ctx, request.GetActivity(), s.temporal, result, appErr(request.GetError())); err != nil {
		return nil, s.integrationErr(err)
	}
	return &emptypb.Empty{}, nil
}

func (s Service) GetDeactivateAccountRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetDeactivateAccountRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeactivateAccountRequest not implemented")
}

func (s Service) SetDeactivateAccountResponse(ctx context.Context, request *apiv1.SetDeactivateAccountResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeactivateAccountResponse not implemented")
}

func (s Service) GetRekeyAccountRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetRekeyAccountRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRekeyAccountRequest not implemented")
}

func (s Service) SetRekeyAccountResponse(ctx context.Context, request *apiv1.SetRekeyAccountResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRekeyAccountResponse not implemented")
}

func (s Service) GetUpdateAccountRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetUpdateAccountRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpdateAccountRequest not implemented")
}

func (s Service) SetUpdateAccountResponse(ctx context.Context, request *apiv1.SetUpdateAccountResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUpdateAccountResponse not implemented")
}

func (s Service) GetCreateOrderRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetCreateOrderRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreateOrderRequest not implemented")
}

func (s Service) SetCreateOrderResponse(ctx context.Context, request *apiv1.SetCreateOrderResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCreateOrderResponse not implemented")
}

func (s Service) GetFinalizeOrderRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetFinalizeOrderRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinalizeOrderRequest not implemented")
}

func (s Service) SetFinalizeOrderResponse(ctx context.Context, request *apiv1.SetFinalizeOrderResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFinalizeOrderResponse not implemented")
}

func (s Service) GetGetOrderRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetGetOrderRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGetOrderRequest not implemented")
}

func (s Service) SetGetOrderResponse(ctx context.Context, request *apiv1.SetGetOrderResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGetOrderResponse not implemented")
}

func (s Service) GetCreateAuthorizationRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetCreateAuthorizationRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreateAuthorizationRequest not implemented")
}

func (s Service) SetCreateAuthorizationResponse(ctx context.Context, request *apiv1.SetCreateAuthorizationResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCreateAuthorizationResponse not implemented")
}

func (s Service) GetDeactivateAuthorizationRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetDeactivateAuthorizationRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeactivateAuthorizationRequest not implemented")
}

func (s Service) SetDeactivateAuthorizationResponse(ctx context.Context, request *apiv1.SetDeactivateAuthorizationResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeactivateAuthorizationResponse not implemented")
}

func (s Service) GetGetAuthorizationRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetGetAuthorizationRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGetAuthorizationRequest not implemented")
}

func (s Service) SetGetAuthorizationResponse(ctx context.Context, request *apiv1.SetGetAuthorizationResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGetAuthorizationResponse not implemented")
}

func (s Service) GetGetChallengeRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetGetChallengeRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGetChallengeRequest not implemented")
}

func (s Service) SetGetChallengeResponse(ctx context.Context, request *apiv1.SetGetChallengeResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGetChallengeResponse not implemented")
}

func (s Service) GetValidateChallengeRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetValidateChallengeRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidateChallengeRequest not implemented")
}

func (s Service) SetValidateChallengeResponse(ctx context.Context, request *apiv1.SetValidateChallengeResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValidateChallengeResponse not implemented")
}

func (s Service) GetGetCertificateRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetGetCertificateRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGetCertificateRequest not implemented")
}

func (s Service) SetGetCertificateResponse(ctx context.Context, request *apiv1.SetGetCertificateResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGetCertificateResponse not implemented")
}

func (s Service) GetRevokeCertificateRequest(ctx context.Context, activity *apiv1.Activity) (*apiv1.GetRevokeCertificateRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRevokeCertificateRequest not implemented")
}

func (s Service) SetRevokeCertificateResponse(ctx context.Context, request *apiv1.SetRevokeCertificateResponseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRevokeCertificateResponse not implemented")
}
