package acme // import "powerssl.dev/sdk/integration/acme"

import "time"

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
