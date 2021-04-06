package service

import (
	"context"
	"crypto/x509"

	"powerssl.dev/common/log"
	"powerssl.dev/sdk/controller/acme"
	"powerssl.dev/sdk/controller/api"
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

func (mw loggingMiddleware) GetCreateAccountRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, _ bool, _ []string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetCreateAccountRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetCreateAccountRequest(ctx, activity)
}

func (mw loggingMiddleware) SetCreateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetCreateAccountResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetCreateAccountResponse(ctx, activity, account, erro)
}

func (mw loggingMiddleware) GetDeactivateAccountRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetDeactivateAccountRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetDeactivateAccountRequest(ctx, activity)
}

func (mw loggingMiddleware) SetDeactivateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetDeactivateAccountResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetDeactivateAccountResponse(ctx, activity, account, erro)
}

func (mw loggingMiddleware) GetRekeyAccountRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetRekeyAccountRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetRekeyAccountRequest(ctx, activity)
}

func (mw loggingMiddleware) SetRekeyAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetRekeyAccountResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetRekeyAccountResponse(ctx, activity, account, erro)
}

func (mw loggingMiddleware) GetUpdateAccountRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ []string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetUpdateAccountRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetUpdateAccountRequest(ctx, activity)
}

func (mw loggingMiddleware) SetUpdateAccountResponse(ctx context.Context, activity *api.Activity, account *api.Account, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetUpdateAccountResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetUpdateAccountResponse(ctx, activity, account, erro)
}

func (mw loggingMiddleware) GetCreateOrderRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, _ []*api.Identifier, _ string, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetCreateOrderRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetCreateOrderRequest(ctx, activity)
}

func (mw loggingMiddleware) SetCreateOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetCreateOrderResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetCreateOrderResponse(ctx, activity, order, erro)
}

func (mw loggingMiddleware) GetFinalizeOrderRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, _ *x509.CertificateRequest, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetFinalizeOrderRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetFinalizeOrderRequest(ctx, activity)
}

func (mw loggingMiddleware) SetFinalizeOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetFinalizeOrderResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetFinalizeOrderResponse(ctx, activity, order, erro)
}

func (mw loggingMiddleware) GetGetOrderRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetGetOrderRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetGetOrderRequest(ctx, activity)
}

func (mw loggingMiddleware) SetGetOrderResponse(ctx context.Context, activity *api.Activity, order *api.Order, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetGetOrderResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetGetOrderResponse(ctx, activity, order, erro)
}

func (mw loggingMiddleware) GetCreateAuthorizationRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, _ *api.Identifier, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetCreateAuthorizationRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetCreateAuthorizationRequest(ctx, activity)
}

func (mw loggingMiddleware) SetCreateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetCreateAuthorizationResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetCreateAuthorizationResponse(ctx, activity, authorization, erro)
}

func (mw loggingMiddleware) GetDeactivateAuthorizationRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetDeactivateAuthorizationRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetDeactivateAuthorizationRequest(ctx, activity)
}

func (mw loggingMiddleware) SetDeactivateAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetDeactivateAuthorizationResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetDeactivateAuthorizationResponse(ctx, activity, authorization, erro)
}

func (mw loggingMiddleware) GetGetAuthorizationRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetGetAuthorizationRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetGetAuthorizationRequest(ctx, activity)
}

func (mw loggingMiddleware) SetGetAuthorizationResponse(ctx context.Context, activity *api.Activity, authorization *api.Authorization, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetGetAuthorizationResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetGetAuthorizationResponse(ctx, activity, authorization, erro)
}

func (mw loggingMiddleware) GetGetChallengeRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetGetChallengeRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetGetChallengeRequest(ctx, activity)
}

func (mw loggingMiddleware) SetGetChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetGetChallengeResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetGetChallengeResponse(ctx, activity, challenge, erro)
}

func (mw loggingMiddleware) GetValidateChallengeRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetValidateChallengeRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetValidateChallengeRequest(ctx, activity)
}

func (mw loggingMiddleware) SetValidateChallengeResponse(ctx context.Context, activity *api.Activity, challenge *api.Challenge, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetValidateChallengeResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetValidateChallengeResponse(ctx, activity, challenge, erro)
}

func (mw loggingMiddleware) GetGetCertificateRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetGetCertificateRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetGetCertificateRequest(ctx, activity)
}

func (mw loggingMiddleware) SetGetCertificateResponse(ctx context.Context, activity *api.Activity, certificates []*x509.Certificate, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetGetCertificateResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetGetCertificateResponse(ctx, activity, certificates, erro)
}

func (mw loggingMiddleware) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (_ *api.Activity, _ string, _ string, _ *x509.Certificate, _ *api.RevocationReason, err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "GetRevokeCertificateRequest", "activity", activity, "err", err)
	}()
	return mw.next.GetRevokeCertificateRequest(ctx, activity)
}

func (mw loggingMiddleware) SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) (err error) {
	defer func() {
		mw.logger.Infow("Called method", "method", "SetRevokeCertificateResponse", "activity", activity, "err", err)
	}()
	return mw.next.SetRevokeCertificateResponse(ctx, activity, erro)
}
