package acme // import "powerssl.dev/powerssl/pkg/controller/acme"

import (
	"context"
	"crypto/x509"

	"powerssl.dev/powerssl/pkg/controller/api"
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
