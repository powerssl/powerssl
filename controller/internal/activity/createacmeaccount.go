package activity

import (
	"context"

	temporalactivity "go.temporal.io/sdk/activity"

	"powerssl.dev/backend/vault"
	"powerssl.dev/controller/internal/service/acme"
	"powerssl.dev/controller/internal/service/integration/activity"
	"powerssl.dev/sdk/controller/api"
	sharedactivity "powerssl.dev/workflow/activity"
)

func CreateACMEAccount(ctx context.Context, params *sharedactivity.CreateACMEAccountParams) (_ *sharedactivity.CreateACMEAccountResults, err error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("CreateACMEAccount", params.ToKeyVals()...)

	var keyToken string
	if keyToken, err = vault.GetClient(ctx).ExportWrappedTransitKey(ctx, params.KeyName); err != nil {
		return nil, err
	}

	a := activity.New(ctx, api.ActivityACMECreateAccount, &acme.CreateACMEAccountParams{
		Contacts:             params.Contacts,
		DirectoryURL:         params.DirectoryURL,
		KeyToken:             keyToken,
		TermsOfServiceAgreed: params.TermsOfServiceAgreed,
	})
	if err = a.Execute(ctx); err != nil {
		return nil, err
	}

	return nil, temporalactivity.ErrResultPending
}
