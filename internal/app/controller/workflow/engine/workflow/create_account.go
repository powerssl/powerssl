package workflow

import (
	"context"
	"fmt"

	"powerssl.io/internal/app/controller/workflow/engine/activity"
	"powerssl.io/internal/app/controller/workflow/engine/activity/acme"
	apiserverapi "powerssl.io/pkg/apiserver/api"
	apiserverclient "powerssl.io/pkg/apiserver/client"
)

type CreateAccount struct {
	Account              string
	DirectoryURL         string
	TermsOfServiceAgreed bool
	Contacts             []string

	Client *apiserverclient.GRPCClient
}

func (w CreateAccount) Kind() string {
	return "CreateAccount"
}

func (w CreateAccount) Run(ctx context.Context) {
	createAccountResult, err := w.CreateAccount(ctx, w.DirectoryURL, w.TermsOfServiceAgreed, w.Contacts)
	if err != nil {
		panic(err)
	}
	if createAccountResult.Error != nil {
		// panic(err)
	}
	acmeAccount := &apiserverapi.ACMEAccount{
		AccountURL: createAccountResult.Account.URL,
		// Contacts:             createAccountResult.Account.Contacts,
		// Status:               createAccountResult.Account.Status,
		// TermsOfServiceAgreed: createAccountResult.Account.TermsOfServiceAgreed,
	}
	updateMask := []string{"account_url"}
	if _, err := w.Client.ACMEAccount.Update(ctx, w.Account, updateMask, acmeAccount); err != nil {
		fmt.Printf("err: %#v\n", err)
	}
}

func (w CreateAccount) CreateAccount(ctx context.Context, directoryURL string, termsOfServiceAgreed bool, contacts []string) (*acme.CreateAccountResult, error) {
	a := activity.NewV2(acme.CreateAccount{})
	a.SetInput(acme.CreateAccountInput{
		DirectoryURL:         directoryURL,
		TermsOfServiceAgreed: termsOfServiceAgreed,
		Contacts:             contacts,
	})
	<-a.Execute(ctx)
	var result *acme.CreateAccountResult
	if err := a.GetResult(&result); err != nil {
		return nil, err
	}
	return result, nil
}
