package service

import (
	"context"
	"crypto/x509"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/acme"
	"powerssl.io/pkg/controller/api"
)

type Middleware func(acme.Service) acme.Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next acme.Service) acme.Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   acme.Service
}

func (mw loggingMiddleware) GetCreateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, bool, []string, error) {
	defer func() {
		mw.logger.Log("method", "GetCreateAccountRequest", "activity", activity)
	}()
	return mw.next.GetCreateAccountRequest(ctx, activity)
}

func (mw loggingMiddleware) SetCreateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetCreateAccountResponse", "activity", activity)
	}()
	return mw.next.SetCreateAccountResponse(ctx, activity, account, erro)
}

func (mw loggingMiddleware) GetDeactivateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	defer func() {
		mw.logger.Log("method", "GetDeactivateAccountRequest", "activity", activity)
	}()
	return mw.next.GetDeactivateAccountRequest(ctx, activity)
}

func (mw loggingMiddleware) SetDeactivateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetDeactivateAccountResponse", "activity", activity)
	}()
	return mw.next.SetDeactivateAccountResponse(ctx, activity, account, erro)
}

func (mw loggingMiddleware) GetRekeyAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	defer func() {
		mw.logger.Log("method", "GetRekeyAccountRequest", "activity", activity)
	}()
	return mw.next.GetRekeyAccountRequest(ctx, activity)
}

func (mw loggingMiddleware) SetRekeyAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetRekeyAccountResponse", "activity", activity)
	}()
	return mw.next.SetRekeyAccountResponse(ctx, activity, account, erro)
}

func (mw loggingMiddleware) GetUpdateAccountRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, []string, error) {
	defer func() {
		mw.logger.Log("method", "GetUpdateAccountRequest", "activity", activity)
	}()
	return mw.next.GetUpdateAccountRequest(ctx, activity)
}

func (mw loggingMiddleware) SetUpdateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetUpdateAccountResponse", "activity", activity)
	}()
	return mw.next.SetUpdateAccountResponse(ctx, activity, account, erro)
}

func (mw loggingMiddleware) GetCreateOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, []*api.Identifier, string, string, error) {
	defer func() {
		mw.logger.Log("method", "GetCreateOrderRequest", "activity", activity)
	}()
	return mw.next.GetCreateOrderRequest(ctx, activity)
}

func (mw loggingMiddleware) SetCreateOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetCreateOrderResponse", "activity", activity)
	}()
	return mw.next.SetCreateOrderResponse(ctx, activity, order, erro)
}

func (mw loggingMiddleware) GetFinalizeOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error) {
	defer func() {
		mw.logger.Log("method", "GetFinalizeOrderRequest", "activity", activity)
	}()
	return mw.next.GetFinalizeOrderRequest(ctx, activity)
}

func (mw loggingMiddleware) SetFinalizeOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetFinalizeOrderResponse", "activity", activity)
	}()
	return mw.next.SetFinalizeOrderResponse(ctx, activity, order, erro)
}

func (mw loggingMiddleware) GetGetOrderRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	defer func() {
		mw.logger.Log("method", "GetGetOrderRequest", "activity", activity)
	}()
	return mw.next.GetGetOrderRequest(ctx, activity)
}

func (mw loggingMiddleware) SetGetOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetGetOrderResponse", "activity", activity)
	}()
	return mw.next.SetGetOrderResponse(ctx, activity, order, erro)
}

func (mw loggingMiddleware) GetCreateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *api.Identifier, error) {
	defer func() {
		mw.logger.Log("method", "GetCreateAuthorizationRequest", "activity", activity)
	}()
	return mw.next.GetCreateAuthorizationRequest(ctx, activity)
}

func (mw loggingMiddleware) SetCreateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetCreateAuthorizationResponse", "activity", activity)
	}()
	return mw.next.SetCreateAuthorizationResponse(ctx, activity, authorization, erro)
}

func (mw loggingMiddleware) GetDeactivateAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	defer func() {
		mw.logger.Log("method", "GetDeactivateAuthorizationRequest", "activity", activity)
	}()
	return mw.next.GetDeactivateAuthorizationRequest(ctx, activity)
}

func (mw loggingMiddleware) SetDeactivateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetDeactivateAuthorizationResponse", "activity", activity)
	}()
	return mw.next.SetDeactivateAuthorizationResponse(ctx, activity, authorization, erro)
}

func (mw loggingMiddleware) GetGetAuthorizationRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	defer func() {
		mw.logger.Log("method", "GetGetAuthorizationRequest", "activity", activity)
	}()
	return mw.next.GetGetAuthorizationRequest(ctx, activity)
}

func (mw loggingMiddleware) SetGetAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetGetAuthorizationResponse", "activity", activity)
	}()
	return mw.next.SetGetAuthorizationResponse(ctx, activity, authorization, erro)
}

func (mw loggingMiddleware) GetGetChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	defer func() {
		mw.logger.Log("method", "GetGetChallengeRequest", "activity", activity)
	}()
	return mw.next.GetGetChallengeRequest(ctx, activity)
}

func (mw loggingMiddleware) SetGetChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetGetChallengeResponse", "activity", activity)
	}()
	return mw.next.SetGetChallengeResponse(ctx, activity, challenge, erro)
}

func (mw loggingMiddleware) GetValidateChallengeRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	defer func() {
		mw.logger.Log("method", "GetValidateChallengeRequest", "activity", activity)
	}()
	return mw.next.GetValidateChallengeRequest(ctx, activity)
}

func (mw loggingMiddleware) SetValidateChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetValidateChallengeResponse", "activity", activity)
	}()
	return mw.next.SetValidateChallengeResponse(ctx, activity, challenge, erro)
}

func (mw loggingMiddleware) GetGetCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, error) {
	defer func() {
		mw.logger.Log("method", "GetGetCertificateRequest", "activity", activity)
	}()
	return mw.next.GetGetCertificateRequest(ctx, activity)
}

func (mw loggingMiddleware) SetGetCertificateResponse(ctx context.Context, activity *api.Activity, certificates []*x509.Certificate, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetGetCertificateResponse", "activity", activity)
	}()
	return mw.next.SetGetCertificateResponse(ctx, activity, certificates, erro)
}

func (mw loggingMiddleware) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, string, *x509.Certificate, *api.RevocationReason, error) {
	defer func() {
		mw.logger.Log("method", "GetRevokeCertificateRequest", "activity", activity)
	}()
	return mw.next.GetRevokeCertificateRequest(ctx, activity)
}

func (mw loggingMiddleware) SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetRevokeCertificateResponse", "activity", activity)
	}()
	return mw.next.SetRevokeCertificateResponse(ctx, activity, erro)
}
