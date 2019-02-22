package service

import (
	"context"
	"crypto/x509"

	"github.com/go-kit/kit/log"

	engineactivity "powerssl.io/internal/app/controller/workflow/engine/activity"
	"powerssl.io/internal/app/controller/workflow/engine/activity/acme"
	service "powerssl.io/pkg/controller/acme"
	"powerssl.io/pkg/controller/api"
)

func New(logger log.Logger) service.Service {
	var svc service.Service
	{
		svc = NewBasicService(logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	logger log.Logger
}

func NewBasicService(logger log.Logger) service.Service {
	return basicService{
		logger: logger,
	}
}

func (bs basicService) GetCreateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, bool, []string, error) {
	var input acme.CreateAccountInput
	if err := engineactivity.GetInput(activity, &input); err != nil {
		return nil, "", false, nil, err
	}
	return activity, input.DirectoryURL, input.TermsOfServiceAgreed, input.Contacts, nil
}

func (bs basicService) SetCreateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	result := acme.CreateAccountResult{account, erro}
	return engineactivity.SetResult(activity, &result)
}

func (bs basicService) GetDeactivateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, error))(activity)
}

func (bs basicService) SetDeactivateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Account, *api.Error) error)(account, erro)
}

func (bs basicService) GetRekeyAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetRekeyAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Account, *api.Error) error)(account, erro)
}

func (bs basicService) GetUpdateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, []string, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, []string, error))(activity)
}

func (bs basicService) SetUpdateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Account, *api.Error) error)(account, erro)
}

func (bs basicService) GetCreateOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, []*api.Identifier, string, string, error) {
	var input acme.CreateOrderInput
	if err := engineactivity.GetInput(activity, &input); err != nil {
		return nil, "", "", nil, "", "", err
	}
	return activity, input.DirectoryURL, input.AccountURL, input.Identifiers, input.NotBefore, input.NotAfter, nil
}

func (bs basicService) SetCreateOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	result := acme.CreateOrderResult{order, erro}
	return engineactivity.SetResult(activity, &result)
}

func (bs basicService) GetFinalizeOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error))(activity)
}

func (bs basicService) SetFinalizeOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Order, *api.Error) error)(order, erro)
}

func (bs basicService) GetGetOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetGetOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Order, *api.Error) error)(order, erro)
}

func (bs basicService) GetCreateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *api.Identifier, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, *api.Identifier, error))(activity)
}

func (bs basicService) SetCreateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Authorization, *api.Error) error)(authorization, erro)
}

func (bs basicService) GetDeactivateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetDeactivateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Authorization, *api.Error) error)(authorization, erro)
}

func (bs basicService) GetGetAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetGetAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Authorization, *api.Error) error)(authorization, erro)
}

func (bs basicService) GetGetChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetGetChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Challenge, *api.Error) error)(challenge, erro)
}

func (bs basicService) GetValidateChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetValidateChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Challenge, *api.Error) error)(challenge, erro)
}

func (bs basicService) GetGetCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, error))(activity)
}

func (bs basicService) SetGetCertificateResponse(ctx context.Context, activity *api.Activity, certificates []*x509.Certificate, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func([]*x509.Certificate, *api.Error) error)(certificates, erro)
}

func (bs basicService) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error) {
	f, err := engineactivity.GetRequest(activity)
	if err != nil {
		return nil, "", "", nil, nil, err
	}
	return f.(func(*api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error))(activity)
}

func (bs basicService) SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	f, err := engineactivity.SetResponse(activity)
	if err != nil {
		return err
	}
	return f.(func(*api.Error) error)(erro)
}
