package service

import (
	"context"
	"crypto/x509"
	"github.com/go-kit/kit/log"
	temporalclient "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"

	integrationactivity "powerssl.dev/powerssl/internal/app/controller/integration/activity"
	"powerssl.dev/powerssl/internal/app/controller/integration/activity/acme"
	"powerssl.dev/powerssl/internal/pkg/temporal/activity"
	service "powerssl.dev/powerssl/pkg/controller/acme"
	"powerssl.dev/powerssl/pkg/controller/api"
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

func (bs basicService) GetCreateAccountRequest(ctx context.Context, apiActivity *api.Activity) (*api.Activity, string, bool, []string, error) {
	var input activity.CreateACMEAccountParams
	if err := integrationactivity.GetInput(ctx, apiActivity, &input); err != nil {
		return nil, "", false, nil, err
	}
	return apiActivity, input.DirectoryURL, input.TermsOfServiceAgreed, input.Contacts, nil
}

func (bs basicService) SetCreateAccountResponse(ctx context.Context, apiActivity *api.Activity, account *api.Account, err *api.Error) error {
	result := &activity.CreateACMEAccountResults{
		Account: account,
	}
	return integrationactivity.SetResult(ctx, apiActivity, bs.temporal, result, appErr(err))
}

func (bs basicService) GetDeactivateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, error))(activity)
}

func (bs basicService) SetDeactivateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Account, *api.Error) error)(account, erro)
}

func (bs basicService) GetRekeyAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetRekeyAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Account, *api.Error) error)(account, erro)
}

func (bs basicService) GetUpdateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, []string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, []string, error))(activity)
}

func (bs basicService) SetUpdateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Account, *api.Error) error)(account, erro)
}

func (bs basicService) GetCreateOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, []*api.Identifier, string, string, error) {
	var input acme.CreateOrderInput
	if err := integrationactivity.GetInput(ctx, activity, &input); err != nil {
		return nil, "", "", nil, "", "", err
	}
	return activity, input.DirectoryURL, input.AccountURL, input.Identifiers, input.NotBefore, input.NotAfter, nil
}

func (bs basicService) SetCreateOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, err *api.Error) error {
	result := &acme.CreateOrderResult{
		Order: order,
	}
	return integrationactivity.SetResult(ctx, activity, bs.temporal, result, appErr(err))
}

func (bs basicService) GetFinalizeOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error))(activity)
}

func (bs basicService) SetFinalizeOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Order, *api.Error) error)(order, erro)
}

func (bs basicService) GetGetOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetGetOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Order, *api.Error) error)(order, erro)
}

func (bs basicService) GetCreateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *api.Identifier, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, *api.Identifier, error))(activity)
}

func (bs basicService) SetCreateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Authorization, *api.Error) error)(authorization, erro)
}

func (bs basicService) GetDeactivateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetDeactivateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Authorization, *api.Error) error)(authorization, erro)
}

func (bs basicService) GetGetAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetGetAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Authorization, *api.Error) error)(authorization, erro)
}

func (bs basicService) GetGetChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetGetChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Challenge, *api.Error) error)(challenge, erro)
}

func (bs basicService) GetValidateChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetValidateChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Challenge, *api.Error) error)(challenge, erro)
}

func (bs basicService) GetGetCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetGetCertificateResponse(ctx context.Context, activity *api.Activity, certificates []*x509.Certificate, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func([]*x509.Certificate, *api.Error) error)(certificates, erro)
}

func (bs basicService) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error) {
	f, err := integrationactivity.GetRequestDeprecated(activity)
	if err != nil {
		return nil, "", "", nil, nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error))(activity)
}

func (bs basicService) SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	f, err := integrationactivity.SetResponseDeprecated(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Error) error)(erro)
}
