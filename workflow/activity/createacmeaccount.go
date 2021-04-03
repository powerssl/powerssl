package activity // import "powerssl.dev/workflow/activity"

import "powerssl.dev/sdk/controller/api"

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
	Account *api.Account
}
