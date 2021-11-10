package acme // import "powerssl.dev/sdk/integration/acme"

import (
	"context"
	"crypto/x509"
	"fmt"

	apiv1 "powerssl.dev/api/controller/v1"
)

type Integration interface {
	CreateAccount(ctx context.Context, keyToken, directoryURL string, termsOfServiceAgreed bool, contacts []string) (*apiv1.Account, error)
	DeactivateAccount(ctx context.Context, accountURL string) (*apiv1.Account, error)
	RekeyAccount(ctx context.Context, accountURL, directoryURL string) (*apiv1.Account, error)
	UpdateAccount(ctx context.Context, accountURL string, contacts []string) (*apiv1.Account, error)

	CreateOrder(ctx context.Context, directoryURL, accountURL string, identifiers []*apiv1.Identifier, notBefore, notAfter string) (*apiv1.Order, error)
	FinalizeOrder(ctx context.Context, accountURL, orderURL string, certificateSigningRequest *x509.CertificateRequest) (*apiv1.Order, error)
	GetOrder(ctx context.Context, accountURL, orderURL string) (*apiv1.Order, error)

	CreateAuthorization(ctx context.Context, directoryURL, accountURL string, identifier *apiv1.Identifier) (*apiv1.Authorization, error)
	DeactivateAuthorization(ctx context.Context, accountURL, authorizationURL string) (*apiv1.Authorization, error)
	GetAuthorization(ctx context.Context, accountURL, authorizationURL string) (*apiv1.Authorization, error)

	GetChallenge(ctx context.Context, accountURL, challengeURL string) (*apiv1.Challenge, error)
	ValidateChallenge(ctx context.Context, accountURL, challengeURL string) (*apiv1.Challenge, error)

	GetCertificate(ctx context.Context, accountURL, certificateURL string) ([]*x509.Certificate, error)
	RevokeCertificate(ctx context.Context, directoryURL, accountURL string, certificate *x509.Certificate, reason *apiv1.RevocationReason) error
}

type integration struct {
	client  apiv1.ACMEServiceClient
	handler Integration
}

func New(client apiv1.ACMEServiceClient, handler Integration) *integration {
	return &integration{
		client:  client,
		handler: handler,
	}
}

func (i *integration) HandleActivity(ctx context.Context, activity *apiv1.Activity) error {
	switch activity.GetName() {
	case apiv1.Activity_ACME_CREATE_ACCOUNT:
		return i.createAccount(ctx, activity)
	case apiv1.Activity_ACME_CREATE_AUTHORIZATION:
		return i.createAuthorization(ctx, activity)
	case apiv1.Activity_ACME_CREATE_ORDER:
		return i.createOrder(ctx, activity)
	case apiv1.Activity_ACME_DEACTIVATE_ACCOUNT:
		return i.deactivateAccount(ctx, activity)
	case apiv1.Activity_ACME_DEACTIVATE_AUTHORIZATION:
		return i.deactivateAuthorization(ctx, activity)
	case apiv1.Activity_ACME_FINALIZE_ORDER:
		return i.finalizeOrder(ctx, activity)
	case apiv1.Activity_ACME_GET_AUTHORIZATION:
		return i.getAuthorization(ctx, activity)
	case apiv1.Activity_ACME_GET_CERTIFICATE:
		return i.getCertificate(ctx, activity)
	case apiv1.Activity_ACME_GET_CHALLENGE:
		return i.getChallenge(ctx, activity)
	case apiv1.Activity_ACME_GET_ORDER:
		return i.getOrder(ctx, activity)
	case apiv1.Activity_ACME_REKEY_ACCOUNT:
		return i.rekeyAccount(ctx, activity)
	case apiv1.Activity_ACME_REVOKE_CERTIFICATE:
		return i.revokeCertificate(ctx, activity)
	case apiv1.Activity_ACME_UPDATE_ACCOUNT:
		return i.updateAccount(ctx, activity)
	case apiv1.Activity_ACME_VALIDATE_CHALLENGE:
		return i.validateChallenge(ctx, activity)
	default:
		return fmt.Errorf("activity %s not implemented", activity.GetName())
	}
}

func (i *integration) createAccount(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetCreateAccountRequest(ctx, activity)
	if err != nil {
		return err
	}
	account, err := i.handler.CreateAccount(ctx, request.KeyToken, request.DirectoryUrl, request.TermsOfServiceAgreed, request.Contacts)
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, err = i.client.SetCreateAccountResponse(ctx, &apiv1.SetCreateAccountResponseRequest{
		Activity: request.GetActivity(),
		Account:  account,
		Error:    e,
	})
	return err
}

func (i *integration) createAuthorization(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetCreateAuthorizationRequest(ctx, activity)
	if err != nil {
		return err
	}
	authorization, err := i.handler.CreateAuthorization(ctx, "", "", nil)
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = authorization, e
	_, err = i.client.SetCreateAuthorizationResponse(ctx, &apiv1.SetCreateAuthorizationResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) createOrder(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetCreateOrderRequest(ctx, activity)
	if err != nil {
		return err
	}
	order, err := i.handler.CreateOrder(ctx, "", "", nil, "", "")
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = order, e
	_, err = i.client.SetCreateOrderResponse(ctx, &apiv1.SetCreateOrderResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) deactivateAccount(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetDeactivateAccountRequest(ctx, activity)
	if err != nil {
		return err
	}
	account, err := i.handler.DeactivateAccount(ctx, "")
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = account, e
	_, err = i.client.SetDeactivateAccountResponse(ctx, &apiv1.SetDeactivateAccountResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) deactivateAuthorization(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetDeactivateAuthorizationRequest(ctx, activity)
	if err != nil {
		return err
	}
	authorization, err := i.handler.DeactivateAuthorization(ctx, "", "")
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = authorization, e
	_, err = i.client.SetDeactivateAuthorizationResponse(ctx, &apiv1.SetDeactivateAuthorizationResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) finalizeOrder(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetFinalizeOrderRequest(ctx, activity)
	if err != nil {
		return err
	}
	order, err := i.handler.FinalizeOrder(ctx, "", "", nil)
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = order, e
	_, err = i.client.SetFinalizeOrderResponse(ctx, &apiv1.SetFinalizeOrderResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) getAuthorization(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetGetAuthorizationRequest(ctx, activity)
	if err != nil {
		return err
	}
	authorization, err := i.handler.GetAuthorization(ctx, "", "")
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = authorization, e
	_, err = i.client.SetGetAuthorizationResponse(ctx, &apiv1.SetGetAuthorizationResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) getCertificate(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetGetCertificateRequest(ctx, activity)
	if err != nil {
		return err
	}
	certificates, err := i.handler.GetCertificate(ctx, "", "")
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = certificates, e
	_, err = i.client.SetGetCertificateResponse(ctx, &apiv1.SetGetCertificateResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) getChallenge(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetGetChallengeRequest(ctx, activity)
	if err != nil {
		return err
	}
	challenge, err := i.handler.GetChallenge(ctx, "", "")
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = challenge, e
	_, err = i.client.SetGetChallengeResponse(ctx, &apiv1.SetGetChallengeResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) getOrder(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetGetOrderRequest(ctx, activity)
	if err != nil {
		return err
	}
	order, err := i.handler.GetOrder(ctx, "", "")
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = order, e
	_, err = i.client.SetGetOrderResponse(ctx, &apiv1.SetGetOrderResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) rekeyAccount(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetRekeyAccountRequest(ctx, activity)
	if err != nil {
		return err
	}
	account, err := i.handler.RekeyAccount(ctx, "", "")
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = account, e
	_, err = i.client.SetRekeyAccountResponse(ctx, &apiv1.SetRekeyAccountResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) revokeCertificate(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetRevokeCertificateRequest(ctx, activity)
	if err != nil {
		return err
	}
	err = i.handler.RevokeCertificate(ctx, "", "", nil, nil)
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_ = e
	_, err = i.client.SetRevokeCertificateResponse(ctx, &apiv1.SetRevokeCertificateResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) updateAccount(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetUpdateAccountRequest(ctx, activity)
	if err != nil {
		return err
	}
	account, err := i.handler.UpdateAccount(ctx, "", nil)
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = account, e
	_, err = i.client.SetUpdateAccountResponse(ctx, &apiv1.SetUpdateAccountResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}

func (i *integration) validateChallenge(ctx context.Context, activity *apiv1.Activity) error {
	request, err := i.client.GetValidateChallengeRequest(ctx, activity)
	if err != nil {
		return err
	}
	challenge, err := i.handler.ValidateChallenge(ctx, "", "")
	var e *apiv1.Error
	if err != nil {
		e = &apiv1.Error{Message: err.Error()}
	}
	_, _ = challenge, e
	_, err = i.client.SetValidateChallengeResponse(ctx, &apiv1.SetValidateChallengeResponseRequest{
		Activity: request.GetActivity(),
	})
	return err
}
