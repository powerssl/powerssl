package service // import "powerssl.io/pkg/controller/acme/service"

import (
	"context"
	"crypto/x509"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/api"
	engineactivity "powerssl.io/pkg/controller/workflow/engine/activity"
)

type Service interface {
	GetCreateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, bool, []string, error)
	SetCreateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error

	GetDeactivateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error)
	SetDeactivateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error

	GetRekeyAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error)
	SetRekeyAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error

	GetUpdateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, []string, error)
	SetUpdateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error

	GetCreateOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, []*api.Identifier, string, string, error)
	SetCreateOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error

	GetFinalizeOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error)
	SetFinalizeOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error

	GetGetOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error)
	SetGetOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error

	GetCreateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *api.Identifier, error)
	SetCreateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error

	GetDeactivateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error)
	SetDeactivateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error

	GetGetAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error)
	SetGetAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error

	GetGetChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error)
	SetGetChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error

	GetValidateChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error)
	SetValidateChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error

	GetGetCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error)
	SetGetCertificateResponse(ctx context.Context, activity *api.Activity, certificates []*x509.Certificate, erro *api.Error) error

	GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error)
	SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error
}

func New(logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	logger log.Logger
}

func NewBasicService(logger log.Logger) Service {
	return basicService{
		logger: logger,
	}
}

func (bs basicService) GetCreateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, bool, []string, error) {
	a, err := engineactivity.Activities.GetByAPIActivity(activity)
	if err != nil {
		return nil, "", false, nil, err
	}
	return a.GetRequest.(func(*api.Activity) (*api.Activity, string, bool, []string, error))(activity)
}

func (bs basicService) SetCreateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	a, err := engineactivity.Activities.GetByAPIActivity(activity)
	if err != nil {
		return err
	}
	return a.SetResponse.(func(*api.Account, *api.Error) error)(account, erro)
}

func (bs basicService) GetDeactivateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	return nil, "", nil
}

func (bs basicService) SetDeactivateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	return nil
}

func (bs basicService) GetRekeyAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	return nil, "", "", nil
}

func (bs basicService) SetRekeyAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	return nil
}

func (bs basicService) GetUpdateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, []string, error) {
	return nil, "", nil, nil
}

func (bs basicService) SetUpdateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	return nil
}

func (bs basicService) GetCreateOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, []*api.Identifier, string, string, error) {
	return nil, "", "", nil, "", "", nil
}

func (bs basicService) SetCreateOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	return nil
}

func (bs basicService) GetFinalizeOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error) {
	return nil, "", "", nil, nil
}

func (bs basicService) SetFinalizeOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	return nil
}

func (bs basicService) GetGetOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	return nil, "", "", nil
}

func (bs basicService) SetGetOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	return nil
}

func (bs basicService) GetCreateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *api.Identifier, error) {
	return nil, "", "", nil, nil
}

func (bs basicService) SetCreateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	return nil
}

func (bs basicService) GetDeactivateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	return nil, "", "", nil
}

func (bs basicService) SetDeactivateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	return nil
}

func (bs basicService) GetGetAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	return nil, "", "", nil
}

func (bs basicService) SetGetAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	return nil
}

func (bs basicService) GetGetChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	return nil, "", "", nil
}

func (bs basicService) SetGetChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	return nil
}

func (bs basicService) GetValidateChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	return nil, "", "", nil
}

func (bs basicService) SetValidateChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	return nil
}

func (bs basicService) GetGetCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	return nil, "", "", nil
}

func (bs basicService) SetGetCertificateResponse(ctx context.Context, activity *api.Activity, certificates []*x509.Certificate, erro *api.Error) error {
	return nil
}

func (bs basicService) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error) {
	return nil, "", "", nil, nil, nil
}

func (bs basicService) SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	return nil
}
