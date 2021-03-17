package service

import (
	"context"
	"crypto/x509"

	"github.com/go-kit/kit/log"
	temporalclient "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	integrationactivity "powerssl.dev/controller/internal/integration/activity"
	"powerssl.dev/controller/internal/integration/activity/acme"
	service "powerssl.dev/sdk/controller/acme"
	"powerssl.dev/sdk/controller/api"
	"powerssl.dev/workflow/activity"
)

func New(logger log.Logger, temporalClient temporalclient.Client) service.Service {
	var svc service.Service
	{
		svc = NewBasicService(logger, temporalClient)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	logger   log.Logger
	temporal temporalclient.Client
}

func NewBasicService(logger log.Logger, temporalClient temporalclient.Client) service.Service {
	return basicService{
		logger:   logger,
		temporal: temporalClient,
	}
}

func appErr(err *api.Error) error {
	var activityError error
	if err != nil && err.Message != "" {
		activityError = temporal.NewApplicationError(err.Message, "")
	}
	return activityError
}

func (s basicService) integrationErr(err error) error {
	s.logger.Log("err", err)
	return status.Error(codes.Internal, "")
}

func (s basicService) GetCreateAccountRequest(ctx context.Context, apiActivity *api.Activity) (*api.Activity, string, string, bool, []string, error) {
	var input *CreateACMEAccountParams
	if err := integrationactivity.GetInput(ctx, apiActivity, &input); err != nil {
		return nil, "", "", false, nil, s.integrationErr(err)
	}
	return apiActivity, input.KeyToken, input.DirectoryURL, input.TermsOfServiceAgreed, input.Contacts, nil
}

func (s basicService) SetCreateAccountResponse(ctx context.Context, apiActivity *api.Activity, account *api.Account, apiErr *api.Error) error {
	result := &activity.CreateACMEAccountResults{
		Account: account,
	}
	if err := integrationactivity.SetResult(ctx, apiActivity, s.temporal, result, appErr(apiErr)); err != nil {
		return s.integrationErr(err)
	}
	return nil
}

func (s basicService) GetDeactivateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, error))(activity)
}

func (s basicService) SetDeactivateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Account, *api.Error) error)(account, erro)
}

func (s basicService) GetRekeyAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (s basicService) SetRekeyAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Account, *api.Error) error)(account, erro)
}

func (s basicService) GetUpdateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, []string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, []string, error))(activity)
}

func (s basicService) SetUpdateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Account, *api.Error) error)(account, erro)
}

func (s basicService) GetCreateOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, []*api.Identifier, string, string, error) {
	var input acme.CreateOrderInput
	if err := integrationactivity.GetInput(ctx, activity, &input); err != nil {
		return nil, "", "", nil, "", "", err
	}
	return activity, input.DirectoryURL, input.AccountURL, input.Identifiers, input.NotBefore, input.NotAfter, nil
}

func (s basicService) SetCreateOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, err *api.Error) error {
	result := &acme.CreateOrderResult{
		Order: order,
	}
	return integrationactivity.SetResult(ctx, activity, s.temporal, result, appErr(err))
}

func (s basicService) GetFinalizeOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error))(activity)
}

func (s basicService) SetFinalizeOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Order, *api.Error) error)(order, erro)
}

func (s basicService) GetGetOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (s basicService) SetGetOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Order, *api.Error) error)(order, erro)
}

func (s basicService) GetCreateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *api.Identifier, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, *api.Identifier, error))(activity)
}

func (s basicService) SetCreateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Authorization, *api.Error) error)(authorization, erro)
}

func (s basicService) GetDeactivateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (s basicService) SetDeactivateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Authorization, *api.Error) error)(authorization, erro)
}

func (s basicService) GetGetAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (s basicService) SetGetAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Authorization, *api.Error) error)(authorization, erro)
}

func (s basicService) GetGetChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (s basicService) SetGetChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Challenge, *api.Error) error)(challenge, erro)
}

func (s basicService) GetValidateChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (s basicService) SetValidateChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Challenge, *api.Error) error)(challenge, erro)
}

func (s basicService) GetGetCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (s basicService) SetGetCertificateResponse(ctx context.Context, activity *api.Activity, certificates []*x509.Certificate, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func([]*x509.Certificate, *api.Error) error)(certificates, erro)
}

func (s basicService) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", nil, nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error))(activity)
}

func (s basicService) SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Error) error)(erro)
}
