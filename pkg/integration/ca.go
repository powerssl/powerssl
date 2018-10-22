package integration

import (
	"context"

	"powerssl.io/pkg/controller/api"
)

type CAIntegration interface {
	GetName() string

	AuthorizeDomain(domain string) (challenges []*api.Challenge, err error)
	RequestCertificate(csr string) (cert string, err error)
	RevokeCertificate(cert string) (err error)
	VerifyDomain(domain string, challengeType api.ChallengeType) (err error)
}

func (i *integration) caAuthorizeDomain(activity *api.Activity) error {
	handler, ok := i.handler.(CAIntegration)
	if !ok {
		return errorNotImplemented
	}
	activity, domain, err := i.client.CA.GetAuthorizeDomainRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	challenges, err := handler.AuthorizeDomain(domain)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.CA.SetAuthorizeDomainResponse(context.Background(), activity, erro, challenges); err != nil {
		return err
	}
	return nil
}

func (i *integration) caRequestCertificate(activity *api.Activity) error {
	handler, ok := i.handler.(CAIntegration)
	if !ok {
		return errorNotImplemented
	}
	activity, certificateSigningRequest, err := i.client.CA.GetRequestCertificateRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	certificate, err := handler.RequestCertificate(certificateSigningRequest)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.CA.SetRequestCertificateResponse(context.Background(), activity, erro, certificate); err != nil {
		return err
	}
	return nil
}

func (i *integration) caRevokeCertificate(activity *api.Activity) error {
	handler, ok := i.handler.(CAIntegration)
	if !ok {
		return errorNotImplemented
	}
	activity, certificate, err := i.client.CA.GetRevokeCertificateRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	err = handler.RevokeCertificate(certificate)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.CA.SetRevokeCertificateResponse(context.Background(), activity, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) caVerifyDomain(activity *api.Activity) error {
	handler, ok := i.handler.(CAIntegration)
	if !ok {
		return errorNotImplemented
	}
	activity, domain, challengeType, err := i.client.CA.GetVerifyDomainRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	err = handler.VerifyDomain(domain, challengeType)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.CA.SetVerifyDomainResponse(context.Background(), activity, erro); err != nil {
		return err
	}
	return nil
}
