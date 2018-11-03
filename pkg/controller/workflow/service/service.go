package service // import "powerssl.io/pkg/controller/workflow/service"

import (
	"context"
	"crypto/x509"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/api"
	"powerssl.io/pkg/controller/workflow/engine/activity"
	"powerssl.io/pkg/controller/workflow/engine/workflow"
)

type Service interface {
	Create(ctx context.Context, kind string) (*api.Workflow, error)
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

func (bs basicService) Create(ctx context.Context, kind string) (*api.Workflow, error) {
	var w *workflow.Workflow

	{
		w = workflow.New(ctx, kind)
		{
			directoryURL := "https://example.com/directory"
			termsOfServiceAgreed := true
			contacts := []string{"mailto:bob@example"}
			createAccount(w, directoryURL, termsOfServiceAgreed, contacts)
		}
		w.Execute()
	}

	{
		w = workflow.New(ctx, kind)
		{
			directoryURL := "https://example.com/directory"
			accountURL := "https://example.com/acct/123"
			dnsnames := []string{"example.com", "example.net"}
			notBefore := ""
			notAfter := ""
			requestCertificate(w, directoryURL, accountURL, dnsnames, notBefore, notAfter)
		}
		w.Execute()
	}

	return w.API(), nil
}

func createAccount(workflow *workflow.Workflow, directoryURL string, termsOfServiceAgreed bool, contacts []string) {
	{
		a := activity.New(api.Activity_ACME_CREATE_ACCOUNT)
		a.GetRequestFunc = func(activity *api.Activity) (*api.Activity, string, bool, []string, error) {
			return activity, directoryURL, termsOfServiceAgreed, contacts, nil
		}
		a.SetResponseFunc = func(account *api.Account, erro *api.Error) error {
			return nil
		}
		workflow.AddActivity(a)
	}
}

func requestCertificate(workflow *workflow.Workflow, directoryURL, accountURL string, dnsnames []string, notBefore, notAfter string) {
	identifiers := make([]*api.Identifier, len(dnsnames))
	for i, dnsname := range dnsnames {
		identifiers[i] = &api.Identifier{Type: api.IdentifierTypeDNS, Value: dnsname}
	}

	var (
		wOrder *api.Order
	)

	{
		a := activity.New(api.Activity_ACME_CREATE_ORDER)
		a.GetRequestFunc = func(activity *api.Activity) (*api.Activity, string, string, []*api.Identifier, string, string, error) {
			return activity, directoryURL, accountURL, identifiers, notBefore, notAfter, nil
		}
		a.SetResponseFunc = func(order *api.Order, erro *api.Error) error {
			wOrder = order
			return nil
		}
		workflow.AddActivity(a)
	}
	wOrder = &api.Order{Authorizations: []string{"123", "1234"}} // TODO
	for _, authorizationURL := range wOrder.Authorizations {
		var (
			wAuthorization *api.Authorization
		)

		{
			a := activity.New(api.Activity_ACME_GET_AUTHORIZATION)
			a.GetRequestFunc = func(activity *api.Activity) (*api.Activity, string, string, error) {
				return activity, accountURL, authorizationURL, nil
			}
			a.SetResponseFunc = func(authorization *api.Authorization, erro *api.Error) error {
				wAuthorization = authorization
				return nil
			}
			workflow.AddActivity(a)
		}
		challengeURL := "" //wAuthorization.Challenges[dns].URL
		{
			a := activity.New(api.Activity_ACME_VALIDATE_CHALLENGE)
			a.GetRequestFunc = func(activity *api.Activity) (*api.Activity, string, string, error) {
				return activity, accountURL, challengeURL, nil
			}
			a.SetResponseFunc = func(challenge *api.Challenge, erro *api.Error) error {
				return nil
			}
			workflow.AddActivity(a)
		}
	}
	orderURL := wOrder.URL
	var certificateSigningRequest *x509.CertificateRequest
	{
		a := activity.New(api.Activity_ACME_FINALIZE_ORDER)
		a.GetRequestFunc = func(activity *api.Activity) (*api.Activity, string, string, *x509.CertificateRequest, error) {
			return activity, directoryURL, orderURL, certificateSigningRequest, nil
		}
		a.SetResponseFunc = func(order *api.Order, erro *api.Error) error {
			wOrder = order
			return nil
		}
		workflow.AddActivity(a)
	}
	certificateURL := wOrder.CertificateURL
	{
		a := activity.New(api.Activity_ACME_GET_CERTIFICATE)
		a.GetRequestFunc = func(activity *api.Activity) (*api.Activity, string, string, error) {
			return activity, accountURL, certificateURL, nil
		}
		a.SetResponseFunc = func(certificates []*x509.Certificate, erro *api.Error) error {
			return nil
		}
		workflow.AddActivity(a)
	}
}
