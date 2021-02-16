package activity

import (
	"powerssl.dev/sdk/controller/api"
)

const CreateACMEAccount = "CreateACMEAccount"

type CreateACMEAccountParams struct {
	DirectoryURL         string
	TermsOfServiceAgreed bool
	Contacts             []string
}

func (p *CreateACMEAccountParams) ToKeyVals() []interface{} {
	return []interface{}{
		"DirectoryURL", p.DirectoryURL,
		"TermsOfServiceAgreed", p.TermsOfServiceAgreed,
		"Contacts", p.Contacts,
	}
}

type CreateACMEAccountResults struct {
	Account *api.Account
}
