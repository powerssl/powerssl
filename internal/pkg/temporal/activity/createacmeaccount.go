package activity

import (
	"context"

	temporalactivity "go.temporal.io/sdk/activity"

	"powerssl.dev/powerssl/internal/app/controller/integration/activity"
	"powerssl.dev/powerssl/pkg/controller/api"
)

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

func CreateACMEAccount(ctx context.Context, params CreateACMEAccountParams) (*CreateACMEAccountResults, error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("CreateACMEAccount", params.ToKeyVals()...)

	a := activity.New(ctx, api.ActivityACMECreateAccount, params)
	if err := a.Execute(ctx); err != nil {
		return nil, err
	}

	return nil, temporalactivity.ErrResultPending
}
