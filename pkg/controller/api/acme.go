package api

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

type Authorization struct{}
type Identifier struct{}
type Order struct{}
type RevocationReason uint
type Challenge struct{}
