package activity // import "powerssl.dev/workflow/activity"

import (
	apiv1 "powerssl.dev/api/controller/v1"
)

const CreateACMEAccount = "CreateACMEAccount"

type CreateACMEAccountParams struct {
	Contacts             []string
	DirectoryURL         string
	KeyName              string
	TermsOfServiceAgreed bool
}

func (p *CreateACMEAccountParams) ToKeyVals() []interface{} {
	return []interface{}{
		"Contacts", p.Contacts,
		"DirectoryURL", p.DirectoryURL,
		"KeyName", p.KeyName,
		"TermsOfServiceAgreed", p.TermsOfServiceAgreed,
	}
}

type CreateACMEAccountResults struct {
	Account *apiv1.Account
}
