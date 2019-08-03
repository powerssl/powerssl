package acme

import "powerssl.dev/powerssl/pkg/controller/api"

type CreateAccountInput struct {
	DirectoryURL         string
	TermsOfServiceAgreed bool
	Contacts             []string
}

type CreateAccountResult struct {
	Account *api.Account
	Error   *api.Error
}

type CreateAccount struct {
	Input  CreateAccountInput
	Result CreateAccountResult
}

func (a CreateAccount) ActivityName() api.ActivityName {
	return api.Activity_ACME_CREATE_ACCOUNT
}

func (a CreateAccount) GetInputType() interface{} {
	return CreateAccountInput{}
}

type CreateOrderInput struct {
	DirectoryURL string
	AccountURL   string
	Identifiers  []*api.Identifier
	NotBefore    string
	NotAfter     string
}

type CreateOrderResult struct {
	Order *api.Order
	Error *api.Error
}

type CreateOrder struct {
	Input  CreateOrderInput
	Result CreateOrderResult
}
