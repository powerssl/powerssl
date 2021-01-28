package activity

import (
	"context"

	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type CreateACMEAccountParams struct {
	DirectoryURL string
	TermsOfServiceAgreed bool
	Contacts []string
}

type CreateACMEAccountResults struct {
	ACMEAccount *api.ACMEAccount
}

func CreateACMEAccount(ctx context.Context, params CreateACMEAccountParams) (CreateACMEAccountResults, error) {
	return CreateACMEAccountResults{}, nil
}

type UpdateAccountParams struct {
	Name string
	UpdateMask []string
	ACMEAccount *api.ACMEAccount
}

func UpdateAccount(ctx context.Context, params UpdateAccountParams) error {
	return nil
}
