package api // import "powerssl.io/powerssl/pkg/controller/api"

//go:generate stringer -type=AccountStatus -trimprefix=AccountStatus
type AccountStatus uint

const (
	AccountStatusValid AccountStatus = iota + 1
	AccountStatusDeactivated
	AccountStatusRevoked
)

type Account struct {
	Contacts             []string
	Status               AccountStatus
	TermsOfServiceAgreed bool
	URL                  string
}

type Authorization struct {
	Challenges []Challenge
}

type IdentifierType string

const IdentifierTypeDNS IdentifierType = "dns"

type Identifier struct {
	Type  IdentifierType
	Value string
}
type Order struct {
	Authorizations []string
	URL            string
	CertificateURL string
}
type RevocationReason uint
type Challenge struct {
	URL string
}
