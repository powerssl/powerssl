package acme

import (
	"context"
	"crypto/x509"
	"fmt"
	"time"

	acmeservice "powerssl.io/pkg/controller/acme/service"
	"powerssl.io/pkg/controller/api"
)

type Account struct {
	URL     string
	KeyName string // keys/123
}

type Authorization struct {
	Identifier Identifier          `json:"identifier"`
	Status     AuthorizationStatus `json:"status"`
	Expires    time.Time           `json:"expires"`
	Challenges []Challenge         `json:"challenges"`
	Wildcard   bool                `json:"wildcard"`

	URL string `json:"-"`
}

type AuthorizationStatus string

const (
	AuthorizationStatusPending     AuthorizationStatus = "pending"
	AuthorizationStatusValid       AuthorizationStatus = "valid"
	AuthorizationStatusInvalid     AuthorizationStatus = "invalid"
	AuthorizationStatusDeactivated AuthorizationStatus = "deactivated"
	AuthorizationStatusExpired     AuthorizationStatus = "expired"
	AuthorizationStatusRevoked     AuthorizationStatus = "revoked"
)

type Challenge struct {
	Type             ChallengeType   `json:"type"`
	URL              string          `json:"url"`
	Status           ChallengeStatus `json:"status"`
	Token            string          `json:"token"`
	Validated        time.Time       `json:"validated"`
	Error            Problem         `json:"error"`
	KeyAuthorization string          `json:"-"`
}

type ChallengeStatus string

const (
	ChallengeStatusPending    ChallengeStatus = "pending"
	ChallengeStatusProcessing ChallengeStatus = "processing"
	ChallengeStatusValid      ChallengeStatus = "valid"
	ChallengeStatusInvalid    ChallengeStatus = "invalid"
)

type ChallengeType string

const (
	ChallengeTypeHTTP01 ChallengeType = "http-01"
	ChallengeTypeDNS01  ChallengeType = "dns-01"
)

type Identifier struct {
	Type  IdentifierType `json:"type"`
	Value string         `json:"value"`
}

type IdentifierType string

const (
	IdentifierTypeDNS IdentifierType = "dns"
)

type Order struct {
	Status         OrderStatus
	Expires        string
	Identifiers    []Identifier
	NotBefore      string
	NotAfter       string
	Error          Problem
	Authorizations []Authorization
	Finalize       string
	Certificate    string
}

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusReady      OrderStatus = "ready"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusValid      OrderStatus = "valid"
	OrderStatusInvalid    OrderStatus = "invalid"
)

type Problem struct {
	Type        string
	Title       string
	Status      uint
	Detail      string
	Instance    string
	Subproblems []Subproblem
}

type RevocationReason uint

const (
	RevocationReasonUnspecified RevocationReason = iota
	RevocationReasonKeyCompromise
	RevocationReasonCACompromise
	RevocationReasonAffiliationChanged
	RevocationReasonSuperseded
	RevocationReasonCessationOfOperation
	RevocationReasonCertificateHold
	_ // Unused
	RevocationReasonRemoveFromCRL
	RevocationReasonPrivilegeWithdrawn
	RevocationReasonAACompromise
)

type Subproblem struct {
	Type       string
	Detaul     string
	Identifier Identifier
}

type Integration interface {
	CreateAccount(directoryURL string, termsOfServiceAgreed bool, contacts []string) (*api.Account, error)
	DeactivateAccount(accountURL string) (*api.Account, error)
	RekeyAccount(accountURL, directoryURL string) (*api.Account, error)
	UpdateAccount(accountURL string, contacts []string) (*api.Account, error)

	CreateOrder(directoryURL, accountURL string, identifiers []*api.Identifier, notBefore, notAfter string) (*api.Order, error)
	FinalizeOrder(accountURL, orderURL string, certificateSigningRequest *x509.CertificateRequest) (*api.Order, error)
	GetOrder(accountURL, orderURL string) (*api.Order, error)

	CreateAuthorization(directoryURL, accountURL string, identifier *api.Identifier) (*api.Authorization, error)
	DeactivateAuthorization(accountURL, authorizationURL string) (*api.Authorization, error)
	GetAuthorization(accountURL, authorizationURL string) (*api.Authorization, error)

	GetChallenge(accountURL, challengeURL string) (*api.Challenge, error)
	ValidateChallenge(accountURL, challengeURL string) (*api.Challenge, error)

	GetCertificate(accountURL, certificateURL string) ([]*x509.Certificate, error)
	RevokeCertificate(directoryURL, accountURL string, certificate *x509.Certificate, reason *api.RevocationReason) error
}

type integration struct {
	client  acmeservice.Service
	handler Integration
}

func New(client acmeservice.Service, handler Integration) *integration {
	return &integration{
		client:  client,
		handler: handler,
	}
}

func (i *integration) HandleActivity(activity *api.Activity) error {
	var err error
	switch activity.Name {
	case api.Activity_ACME_CREATE_ACCOUNT:
		err = i.createAccount(activity)
	case api.Activity_ACME_CREATE_AUTHORIZATION:
		err = i.createAuthorization(activity)
	case api.Activity_ACME_CREATE_ORDER:
		err = i.createOrder(activity)
	case api.Activity_ACME_DEACTIVATE_ACCOUNT:
		err = i.deactivateAccount(activity)
	case api.Activity_ACME_DEACTIVATE_AUTHORIZATION:
		err = i.deactivateAuthorization(activity)
	case api.Activity_ACME_FINALIZE_ORDER:
		err = i.finalizeOrder(activity)
	case api.Activity_ACME_GET_AUTHORIZATION:
		err = i.getAuthorization(activity)
	case api.Activity_ACME_GET_CERTIFICATE:
		err = i.getCertificate(activity)
	case api.Activity_ACME_GET_CHALLENGE:
		err = i.getChallenge(activity)
	case api.Activity_ACME_GET_ORDER:
		err = i.getOrder(activity)
	case api.Activity_ACME_REKEY_ACCOUNT:
		err = i.rekeyAccount(activity)
	case api.Activity_ACME_REVOKE_CERTIFICATE:
		err = i.revokeCertificate(activity)
	case api.Activity_ACME_UPDATE_ACCOUNT:
		err = i.updateAccount(activity)
	case api.Activity_ACME_VALIDATE_CHALLENGE:
		err = i.validateChallenge(activity)
	default:
		err = fmt.Errorf("Activity %s not implemented", activity.Name)
	}
	return err
}

func (i *integration) createAccount(activity *api.Activity) error {
	activity, directoryURL, termsOfServiceAgreed, contacts, err := i.client.GetCreateAccountRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	account, err := i.handler.CreateAccount(directoryURL, termsOfServiceAgreed, contacts)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetCreateAccountResponse(context.Background(), activity, account, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) createAuthorization(activity *api.Activity) error {
	activity, directoryURL, accountURL, identifier, err := i.client.GetCreateAuthorizationRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	authorization, err := i.handler.CreateAuthorization(directoryURL, accountURL, identifier)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetCreateAuthorizationResponse(context.Background(), activity, authorization, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) createOrder(activity *api.Activity) error {
	activity, directoryURL, accountURL, identifiers, notBefore, notAfter, err := i.client.GetCreateOrderRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	order, err := i.handler.CreateOrder(directoryURL, accountURL, identifiers, notBefore, notAfter)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetCreateOrderResponse(context.Background(), activity, order, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) deactivateAccount(activity *api.Activity) error {
	activity, accountURL, err := i.client.GetDeactivateAccountRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	account, err := i.handler.DeactivateAccount(accountURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetDeactivateAccountResponse(context.Background(), activity, account, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) deactivateAuthorization(activity *api.Activity) error {
	activity, accountURL, authorizationURL, err := i.client.GetDeactivateAuthorizationRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	authorization, err := i.handler.DeactivateAuthorization(accountURL, authorizationURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetDeactivateAuthorizationResponse(context.Background(), activity, authorization, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) finalizeOrder(activity *api.Activity) error {
	activity, accountURL, orderURL, certificateSigningRequest, err := i.client.GetFinalizeOrderRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	order, err := i.handler.FinalizeOrder(accountURL, orderURL, certificateSigningRequest)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetFinalizeOrderResponse(context.Background(), activity, order, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) getAuthorization(activity *api.Activity) error {
	activity, accountURL, authorizationURL, err := i.client.GetGetAuthorizationRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	authorization, err := i.handler.GetAuthorization(accountURL, authorizationURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetGetAuthorizationResponse(context.Background(), activity, authorization, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) getCertificate(activity *api.Activity) error {
	activity, accountURL, certificateURL, err := i.client.GetGetCertificateRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	certificates, err := i.handler.GetCertificate(accountURL, certificateURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetGetCertificateResponse(context.Background(), activity, certificates, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) getChallenge(activity *api.Activity) error {
	activity, accountURL, challengeURL, err := i.client.GetGetChallengeRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	challenge, err := i.handler.GetChallenge(accountURL, challengeURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetGetChallengeResponse(context.Background(), activity, challenge, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) getOrder(activity *api.Activity) error {
	activity, accountURL, orderURL, err := i.client.GetGetOrderRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	order, err := i.handler.GetOrder(accountURL, orderURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetGetOrderResponse(context.Background(), activity, order, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) rekeyAccount(activity *api.Activity) error {
	activity, accountURL, directoryURL, err := i.client.GetRekeyAccountRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	account, err := i.handler.RekeyAccount(accountURL, directoryURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetRekeyAccountResponse(context.Background(), activity, account, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) revokeCertificate(activity *api.Activity) error {
	activity, directoryURL, accountURL, certificate, reason, err := i.client.GetRevokeCertificateRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	err = i.handler.RevokeCertificate(directoryURL, accountURL, certificate, reason)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetRevokeCertificateResponse(context.Background(), activity, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) updateAccount(activity *api.Activity) error {
	activity, accountURL, contacts, err := i.client.GetUpdateAccountRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	account, err := i.handler.UpdateAccount(accountURL, contacts)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetUpdateAccountResponse(context.Background(), activity, account, erro); err != nil {
		return err
	}
	return nil
}

func (i *integration) validateChallenge(activity *api.Activity) error {
	activity, accountURL, challengeURL, err := i.client.GetValidateChallengeRequest(context.Background(), activity)
	if err != nil {
		return err
	}
	challenge, err := i.handler.ValidateChallenge(accountURL, challengeURL)
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	if err := i.client.SetValidateChallengeResponse(context.Background(), activity, challenge, erro); err != nil {
		return err
	}
	return nil
}
