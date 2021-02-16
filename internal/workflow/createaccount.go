package workflow

const CreateAccount = "CreateAccount"

type CreateAccountParams struct {
	Account              string
	DirectoryURL         string
	TermsOfServiceAgreed bool
	Contacts             []string
}

func (p *CreateAccountParams) ToKeyVals() []interface{} {
	return []interface{}{
		"Account", p.Account,
		"DirectoryURL", p.DirectoryURL,
		"TermsOfServiceAgreed", p.TermsOfServiceAgreed,
		"Contacts", p.Contacts,
	}
}
