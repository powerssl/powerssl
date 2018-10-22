package service // import "powerssl.io/pkg/controller/ca/service"

import (
	"context"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/api"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) GetAuthorizeDomainRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	defer func() {
		mw.logger.Log("method", "GetAuthorizeDomainRequest", "activity", activity)
	}()
	return mw.next.GetAuthorizeDomainRequest(ctx, activity)
}

func (mw loggingMiddleware) SetAuthorizeDomainResponse(ctx context.Context, activity *api.Activity, erro *api.Error, challenges []*api.Challenge) error {
	defer func() {
		mw.logger.Log("method", "SetAuthorizeDomainResponse", "activity", activity)
	}()
	return mw.next.SetAuthorizeDomainResponse(ctx, activity, erro, challenges)
}

func (mw loggingMiddleware) GetRequestCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	defer func() {
		mw.logger.Log("method", "GetRequestCertificateRequest", "activity", activity)
	}()
	return mw.next.GetRequestCertificateRequest(ctx, activity)
}

func (mw loggingMiddleware) SetRequestCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error, certificate string) error {
	defer func() {
		mw.logger.Log("method", "SetRequestCertificateResponse", "activity", activity)
	}()
	return mw.next.SetRequestCertificateResponse(ctx, activity, erro, certificate)
}

func (mw loggingMiddleware) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
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

func (mw loggingMiddleware) GetVerifyDomainRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, api.ChallengeType, error) {
	defer func() {
		mw.logger.Log("method", "GetVerifyDomainRequest", "activity", activity)
	}()
	return mw.next.GetVerifyDomainRequest(ctx, activity)
}

func (mw loggingMiddleware) SetVerifyDomainResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	defer func() {
		mw.logger.Log("method", "SetVerifyDomainResponse", "activity", activity)
	}()
	return mw.next.SetVerifyDomainResponse(ctx, activity, erro)
}
