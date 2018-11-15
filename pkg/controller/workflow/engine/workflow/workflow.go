package workflow

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"

	apiserverapi "powerssl.io/pkg/apiserver/api"
	"powerssl.io/pkg/controller/api"
	"powerssl.io/pkg/controller/workflow/engine/activity"
	"powerssl.io/pkg/controller/workflow/engine/activity/acme"
)

type Definition interface {
	Kind() string
	Run(ctx context.Context)
}

type Workflow struct {
	UUID uuid.UUID

	Definition
}

func New(definition Definition) *Workflow {
	w := &Workflow{
		UUID:       uuid.New(),
		Definition: definition,
	}
	Workflows.Put(w)
	return w
}

func (w *Workflow) Name() string {
	return fmt.Sprintf("workflows/%s", w.UUID)
}

func (w *Workflow) ToAPI() *api.Workflow {
	return &api.Workflow{
		Name: w.Name(),
	}
}

func (w *Workflow) Execute(ctx context.Context) chan struct{} {
	c := make(chan struct{})
	go w.execute(ctx, c)
	return c
}

func (w *Workflow) execute(ctx context.Context, c chan struct{}) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Workflow")
	defer span.Finish()
	span.SetTag("kind", w.Kind())
	span.SetTag("name", w.Name())
	w.Run(ctx)
	// select {
	// case <-w.Run(ctx):
	// case <-ctx.Done():
	// 	fmt.Println(ctx.Err()) // TODO
	// }
	close(c)
}

type CreateAccount struct {
	AccountName          string
	DirectoryURL         string
	TermsOfServiceAgreed bool
	Contacts             []string
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
		panic(err)
	}
	acmeAccount := &apiserverapi.ACMEAccount{
		AccountURL: createAccountResult.Account.URL,
		// Contacts:             createAccountResult.Account.Contacts,
		// Status:               createAccountResult.Account.Status,
		// TermsOfServiceAgreed: createAccountResult.Account.TermsOfServiceAgreed,
	}
	var _ = acmeAccount
	//if _, err := client.ACMEAccount.Update(context.Background(), w.AccountName, acmeAccount); err != nil {
	//	panic(err)
	//}
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

type RequestCertificate struct {
	DirectoryURL string
	AccountURL   string
	Dnsnames     []string
	NotBefore    string
	NotAfter     string
}

func (w RequestCertificate) Kind() string {
	return "RequestCertificate"
}

func (w RequestCertificate) Run(ctx context.Context) {
	identifiers := make([]*api.Identifier, len(w.Dnsnames))
	for i, dnsname := range w.Dnsnames {
		identifiers[i] = &api.Identifier{Type: api.IdentifierTypeDNS, Value: dnsname}
	}

	{
		a := activity.New(api.Activity_ACME_CREATE_ORDER)
		a.SetInput(acme.CreateOrderInput{w.DirectoryURL, w.AccountURL, identifiers, w.NotBefore, w.NotAfter})
		a.Execute(ctx)
	}
	for range []struct{}{} {
		activity.New(api.Activity_ACME_GET_AUTHORIZATION)
		activity.New(api.Activity_ACME_VALIDATE_CHALLENGE)
	}
	activity.New(api.Activity_ACME_FINALIZE_ORDER)
	activity.New(api.Activity_ACME_GET_CERTIFICATE)
}
