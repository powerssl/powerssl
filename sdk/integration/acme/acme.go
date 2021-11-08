package acme // import "powerssl.dev/sdk/integration/acme"

import (
	"context"
	"crypto/x509"
	"fmt"

	apiv1 "powerssl.dev/api/controller/v1"
	"powerssl.dev/sdk/controller/acme"
	"powerssl.dev/sdk/controller/api"
)

type Integration interface {
	CreateAccount(ctx context.Context, keyToken, directoryURL string, termsOfServiceAgreed bool, contacts []string) (*api.Account, error)
	DeactivateAccount(ctx context.Context, accountURL string) (*api.Account, error)
	RekeyAccount(ctx context.Context, accountURL, directoryURL string) (*api.Account, error)
	UpdateAccount(ctx context.Context, accountURL string, contacts []string) (*api.Account, error)

	CreateOrder(ctx context.Context, directoryURL, accountURL string, identifiers []*api.Identifier, notBefore, notAfter string) (*api.Order, error)
	FinalizeOrder(ctx context.Context, accountURL, orderURL string, certificateSigningRequest *x509.CertificateRequest) (*api.Order, error)
	GetOrder(ctx context.Context, accountURL, orderURL string) (*api.Order, error)

	CreateAuthorization(ctx context.Context, directoryURL, accountURL string, identifier *api.Identifier) (*api.Authorization, error)
	DeactivateAuthorization(ctx context.Context, accountURL, authorizationURL string) (*api.Authorization, error)
	GetAuthorization(ctx context.Context, accountURL, authorizationURL string) (*api.Authorization, error)

	GetChallenge(ctx context.Context, accountURL, challengeURL string) (*api.Challenge, error)
	ValidateChallenge(ctx context.Context, accountURL, challengeURL string) (*api.Challenge, error)

	GetCertificate(ctx context.Context, accountURL, certificateURL string) ([]*x509.Certificate, error)
	RevokeCertificate(ctx context.Context, directoryURL, accountURL string, certificate *x509.Certificate, reason *api.RevocationReason) error
}

type integration struct {
	client  acme.Service
	handler Integration
}

func New(client acme.Service, handler Integration) *integration {
	return &integration{
		client:  client,
		handler: handler,
	}
}

func (i *integration) HandleActivity(ctx context.Context, activity *apiv1.Activity) error {
	var err error
	switch activity.GetName() {
	case apiv1.Activity_ACME_CREATE_ACCOUNT:
		err = i.createAccount(ctx, activity)
	case apiv1.Activity_ACME_CREATE_AUTHORIZATION:
		err = i.createAuthorization(ctx, activity)
	case apiv1.Activity_ACME_CREATE_ORDER:
		err = i.createOrder(ctx, activity)
	case apiv1.Activity_ACME_DEACTIVATE_ACCOUNT:
		err = i.deactivateAccount(ctx, activity)
	case apiv1.Activity_ACME_DEACTIVATE_AUTHORIZATION:
		err = i.deactivateAuthorization(ctx, activity)
	case apiv1.Activity_ACME_FINALIZE_ORDER:
		err = i.finalizeOrder(ctx, activity)
	case apiv1.Activity_ACME_GET_AUTHORIZATION:
		err = i.getAuthorization(ctx, activity)
	case apiv1.Activity_ACME_GET_CERTIFICATE:
		err = i.getCertificate(ctx, activity)
	case apiv1.Activity_ACME_GET_CHALLENGE:
		err = i.getChallenge(ctx, activity)
	case apiv1.Activity_ACME_GET_ORDER:
		err = i.getOrder(ctx, activity)
	case apiv1.Activity_ACME_REKEY_ACCOUNT:
		err = i.rekeyAccount(ctx, activity)
	case apiv1.Activity_ACME_REVOKE_CERTIFICATE:
		err = i.revokeCertificate(ctx, activity)
	case apiv1.Activity_ACME_UPDATE_ACCOUNT:
		err = i.updateAccount(ctx, activity)
	case apiv1.Activity_ACME_VALIDATE_CHALLENGE:
		err = i.validateChallenge(ctx, activity)
	default:
		err = fmt.Errorf("activity %s not implemented", activity.Name)
	}
	return err
}

func apiActivity(activity *apiv1.Activity) *api.Activity {
	return &api.Activity{
		Name:      api.ActivityName(activity.GetName()),
		Signature: activity.GetSignature(),
		Token:     activity.GetToken(),
	}
}

func (i *integration) createAccount(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, keyToken, directoryURL, termsOfServiceAgreed, contacts, err := i.client.GetCreateAccountRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	account, err := i.handler.CreateAccount(ctx, keyToken, directoryURL, termsOfServiceAgreed, contacts)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetCreateAccountResponse(ctx, apiActivity, account, erro)
}

func (i *integration) createAuthorization(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, directoryURL, accountURL, identifier, err := i.client.GetCreateAuthorizationRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	authorization, err := i.handler.CreateAuthorization(ctx, directoryURL, accountURL, identifier)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetCreateAuthorizationResponse(ctx, apiActivity, authorization, erro)
}

func (i *integration) createOrder(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, directoryURL, accountURL, identifiers, notBefore, notAfter, err := i.client.GetCreateOrderRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	order, err := i.handler.CreateOrder(ctx, directoryURL, accountURL, identifiers, notBefore, notAfter)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetCreateOrderResponse(ctx, apiActivity, order, erro)
}

func (i *integration) deactivateAccount(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, err := i.client.GetDeactivateAccountRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	account, err := i.handler.DeactivateAccount(ctx, accountURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetDeactivateAccountResponse(ctx, apiActivity, account, erro)
}

func (i *integration) deactivateAuthorization(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, authorizationURL, err := i.client.GetDeactivateAuthorizationRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	authorization, err := i.handler.DeactivateAuthorization(ctx, accountURL, authorizationURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetDeactivateAuthorizationResponse(ctx, apiActivity, authorization, erro)
}

func (i *integration) finalizeOrder(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, orderURL, certificateSigningRequest, err := i.client.GetFinalizeOrderRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	order, err := i.handler.FinalizeOrder(ctx, accountURL, orderURL, certificateSigningRequest)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetFinalizeOrderResponse(ctx, apiActivity, order, erro)
}

func (i *integration) getAuthorization(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, authorizationURL, err := i.client.GetGetAuthorizationRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	authorization, err := i.handler.GetAuthorization(ctx, accountURL, authorizationURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetGetAuthorizationResponse(ctx, apiActivity, authorization, erro)
}

func (i *integration) getCertificate(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, certificateURL, err := i.client.GetGetCertificateRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	certificates, err := i.handler.GetCertificate(ctx, accountURL, certificateURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetGetCertificateResponse(ctx, apiActivity, certificates, erro)
}

func (i *integration) getChallenge(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, challengeURL, err := i.client.GetGetChallengeRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	challenge, err := i.handler.GetChallenge(ctx, accountURL, challengeURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetGetChallengeResponse(ctx, apiActivity, challenge, erro)
}

func (i *integration) getOrder(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, orderURL, err := i.client.GetGetOrderRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	order, err := i.handler.GetOrder(ctx, accountURL, orderURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetGetOrderResponse(ctx, apiActivity, order, erro)
}

func (i *integration) rekeyAccount(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, directoryURL, err := i.client.GetRekeyAccountRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	account, err := i.handler.RekeyAccount(ctx, accountURL, directoryURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetRekeyAccountResponse(ctx, apiActivity, account, erro)
}

func (i *integration) revokeCertificate(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, directoryURL, accountURL, certificate, reason, err := i.client.GetRevokeCertificateRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	err = i.handler.RevokeCertificate(ctx, directoryURL, accountURL, certificate, reason)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetRevokeCertificateResponse(ctx, apiActivity, erro)
}

func (i *integration) updateAccount(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, contacts, err := i.client.GetUpdateAccountRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	account, err := i.handler.UpdateAccount(ctx, accountURL, contacts)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetUpdateAccountResponse(ctx, apiActivity, account, erro)
}

func (i *integration) validateChallenge(ctx context.Context, activity *apiv1.Activity) error {
	apiActivity, accountURL, challengeURL, err := i.client.GetValidateChallengeRequest(ctx, apiActivity(activity))
	if err != nil {
		return err
	}
	challenge, err := i.handler.ValidateChallenge(ctx, accountURL, challengeURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	return i.client.SetValidateChallengeResponse(ctx, apiActivity, challenge, erro)
}
