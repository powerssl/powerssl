package acme

type CreateACMEAccountParams struct {
	Contacts             []string
	DirectoryURL         string
	KeyToken             string
	TermsOfServiceAgreed bool
}
