package service // import "powerssl.io/pkg/controller/ca/service"

import (
	"context"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/api"
	workflowengine "powerssl.io/pkg/controller/workflow/engine"
)

type Service interface {
	GetAuthorizeDomainRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error)
	SetAuthorizeDomainResponse(ctx context.Context, activity *api.Activity, erro *api.Error, challenges []*api.Challenge) error
	GetRequestCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error)
	SetRequestCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error, certificate string) error
	GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error)
	SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error
	GetVerifyDomainRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, api.ChallengeType, error)
	SetVerifyDomainResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error
}

func New(logger log.Logger, workflowengine *workflowengine.Engine) Service {
	var svc Service
	{
		svc = NewBasicService(logger, workflowengine)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	logger         log.Logger
	workflowengine *workflowengine.Engine
}

func NewBasicService(logger log.Logger, workflowengine *workflowengine.Engine) Service {
	return basicService{
		logger:         logger,
		workflowengine: workflowengine,
	}
}

func (bs basicService) GetAuthorizeDomainRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	return nil, "", nil
}

func (bs basicService) SetAuthorizeDomainResponse(ctx context.Context, activity *api.Activity, erro *api.Error, challenges []*api.Challenge) error {
	return nil
}

func (bs basicService) GetRequestCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	return nil, "", nil
}

func (bs basicService) SetRequestCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error, certificate string) error {
	return nil
}

func (bs basicService) GetRevokeCertificateRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, error) {
	return nil, "", nil
}

func (bs basicService) SetRevokeCertificateResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	return nil
}

func (bs basicService) GetVerifyDomainRequest(ctx context.Context, activity *api.Activity) (*api.Activity, string, api.ChallengeType, error) {
	return nil, "", 0, nil
}

func (bs basicService) SetVerifyDomainResponse(ctx context.Context, activity *api.Activity, erro *api.Error) error {
	return nil
}
