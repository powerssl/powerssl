package acme

import (
	"context"
	"crypto/x509"
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
	RekeyAccount(accountURL string, directoryURL string) (*api.Account, error)
	UpdateAccount(accountURL string, contacts []string) (*api.Account, error)

	CreateOrder(directoryURL string, accountURL string, identifiers []*api.Identifier, notBefore, notAfter string) (*api.Order, error)
	FinalizeOrder(accountURL string, orderURL string, certificateSigningRequest *x509.CertificateRequest) (*api.Order, error)
	GetOrder(accountURL string, orderURL string) (*api.Order, error)

	CreateAuthorization(directoryURL string, accountURL string, identifier *api.Identifier) (*api.Authorization, error)
	DeactivateAuthorization(accountURL string, authorizationURL string) (*api.Authorization, error)
	GetAuthorization(accountURL string, authorizationURL string) (*api.Authorization, error)

	GetChallenge(accountURL string, challengeURL string) (*api.Challenge, error)
	ValidateChallenge(accountURL string, challengeURL string) (*api.Challenge, error)

	GetCertificate(accountURL string, certificateURL string) ([]*x509.Certificate, error)
	RevokeCertificate(directoryURL string, accountURL string, certificate *x509.Certificate, reason *api.RevocationReason) error
}

type integration struct {
	client  acmeservice.Service
	handler Integration
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
