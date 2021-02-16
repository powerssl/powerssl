package acme // import "powerssl.dev/sdk/integration/acme"

import (
	"context"
	"crypto/x509"
	"fmt"

	"powerssl.dev/sdk/controller/acme"
	"powerssl.dev/sdk/controller/api"
)

type Integration interface {
	CreateAccount(ctx context.Context, directoryURL string, termsOfServiceAgreed bool, contacts []string) (*api.Account, error)
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

func (i *integration) HandleActivity(ctx context.Context, activity *api.Activity) error {
	var err error
	switch activity.Name {
	case api.ActivityACMECreateAccount:
		err = i.createAccount(ctx, activity)
	case api.ActivityACMECreateAuthorization:
		err = i.createAuthorization(ctx, activity)
	case api.ActivityACMECreateOrder:
		err = i.createOrder(ctx, activity)
	case api.ActivityACMEDeactivateAccount:
		err = i.deactivateAccount(ctx, activity)
	case api.ActivityACMEDeactivateAuthorization:
		err = i.deactivateAuthorization(ctx, activity)
	case api.ActivityACMEFinalizeOrder:
		err = i.finalizeOrder(ctx, activity)
	case api.ActivityACMEGetAuthorization:
		err = i.getAuthorization(ctx, activity)
	case api.ActivityACMEGetCertificate:
		err = i.getCertificate(ctx, activity)
	case api.ActivityACMEGetChallenge:
		err = i.getChallenge(ctx, activity)
	case api.ActivityACMEGetOrder:
		err = i.getOrder(ctx, activity)
	case api.ActivityACMERekeyAccount:
		err = i.rekeyAccount(ctx, activity)
	case api.ActivityACMERevokeCertificate:
		err = i.revokeCertificate(ctx, activity)
	case api.ActivityACMEUpdateAccount:
		err = i.updateAccount(ctx, activity)
	case api.ActivityACMEValidateChallenge:
		err = i.validateChallenge(ctx, activity)
	default:
		err = fmt.Errorf("activity %s not implemented", activity.Name)
	}
	return err
}

func (i *integration) createAccount(ctx context.Context, activity *api.Activity) error {
	activity, directoryURL, termsOfServiceAgreed, contacts, err := i.client.GetCreateAccountRequest(ctx, activity)
	if err != nil {
		return err
	}
	account, err := i.handler.CreateAccount(ctx, directoryURL, termsOfServiceAgreed, contacts)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetCreateAccountResponse(ctx, activity, account, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) createAuthorization(ctx context.Context, activity *api.Activity) error {
	activity, directoryURL, accountURL, identifier, err := i.client.GetCreateAuthorizationRequest(ctx, activity)
	if err != nil {
		return err
	}
	authorization, err := i.handler.CreateAuthorization(ctx, directoryURL, accountURL, identifier)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetCreateAuthorizationResponse(ctx, activity, authorization, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) createOrder(ctx context.Context, activity *api.Activity) error {
	activity, directoryURL, accountURL, identifiers, notBefore, notAfter, err := i.client.GetCreateOrderRequest(ctx, activity)
	if err != nil {
		return err
	}
	order, err := i.handler.CreateOrder(ctx, directoryURL, accountURL, identifiers, notBefore, notAfter)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetCreateOrderResponse(ctx, activity, order, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) deactivateAccount(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, err := i.client.GetDeactivateAccountRequest(ctx, activity)
	if err != nil {
		return err
	}
	account, err := i.handler.DeactivateAccount(ctx, accountURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetDeactivateAccountResponse(ctx, activity, account, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) deactivateAuthorization(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, authorizationURL, err := i.client.GetDeactivateAuthorizationRequest(ctx, activity)
	if err != nil {
		return err
	}
	authorization, err := i.handler.DeactivateAuthorization(ctx, accountURL, authorizationURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetDeactivateAuthorizationResponse(ctx, activity, authorization, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) finalizeOrder(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, orderURL, certificateSigningRequest, err := i.client.GetFinalizeOrderRequest(ctx, activity)
	if err != nil {
		return err
	}
	order, err := i.handler.FinalizeOrder(ctx, accountURL, orderURL, certificateSigningRequest)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetFinalizeOrderResponse(ctx, activity, order, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) getAuthorization(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, authorizationURL, err := i.client.GetGetAuthorizationRequest(ctx, activity)
	if err != nil {
		return err
	}
	authorization, err := i.handler.GetAuthorization(ctx, accountURL, authorizationURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetGetAuthorizationResponse(ctx, activity, authorization, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) getCertificate(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, certificateURL, err := i.client.GetGetCertificateRequest(ctx, activity)
	if err != nil {
		return err
	}
	certificates, err := i.handler.GetCertificate(ctx, accountURL, certificateURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetGetCertificateResponse(ctx, activity, certificates, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) getChallenge(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, challengeURL, err := i.client.GetGetChallengeRequest(ctx, activity)
	if err != nil {
		return err
	}
	challenge, err := i.handler.GetChallenge(ctx, accountURL, challengeURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetGetChallengeResponse(ctx, activity, challenge, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) getOrder(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, orderURL, err := i.client.GetGetOrderRequest(ctx, activity)
	if err != nil {
		return err
	}
	order, err := i.handler.GetOrder(ctx, accountURL, orderURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetGetOrderResponse(ctx, activity, order, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) rekeyAccount(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, directoryURL, err := i.client.GetRekeyAccountRequest(ctx, activity)
	if err != nil {
		return err
	}
	account, err := i.handler.RekeyAccount(ctx, accountURL, directoryURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetRekeyAccountResponse(ctx, activity, account, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) revokeCertificate(ctx context.Context, activity *api.Activity) error {
	activity, directoryURL, accountURL, certificate, reason, err := i.client.GetRevokeCertificateRequest(ctx, activity)
	if err != nil {
		return err
	}
	err = i.handler.RevokeCertificate(ctx, directoryURL, accountURL, certificate, reason)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetRevokeCertificateResponse(ctx, activity, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) updateAccount(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, contacts, err := i.client.GetUpdateAccountRequest(ctx, activity)
	if err != nil {
		return err
	}
	account, err := i.handler.UpdateAccount(ctx, accountURL, contacts)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetUpdateAccountResponse(ctx, activity, account, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) validateChallenge(ctx context.Context, activity *api.Activity) error {
	activity, accountURL, challengeURL, err := i.client.GetValidateChallengeRequest(ctx, activity)
	if err != nil {
		return err
	}
	challenge, err := i.handler.ValidateChallenge(ctx, accountURL, challengeURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetValidateChallengeResponse(ctx, activity, challenge, erro); err != nil {
		return err
	}
	return nil
}
