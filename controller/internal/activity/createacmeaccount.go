package activity

import (
	"context"

	temporalactivity "go.temporal.io/sdk/activity"

	"powerssl.dev/backend/vault"
	"powerssl.dev/sdk/controller/api"
	sharedactivity "powerssl.dev/workflow/activity"

	"powerssl.dev/controller/internal/acme/service"
	"powerssl.dev/controller/internal/integration/activity"
)

func CreateACMEAccount(ctx context.Context, params *sharedactivity.CreateACMEAccountParams) (_ *sharedactivity.CreateACMEAccountResults, err error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("CreateACMEAccount", params.ToKeyVals()...)

	var keyToken string
	if keyToken, err = vault.GetClient(ctx).ExportWrappedTransitKey(ctx, params.KeyName); err != nil {
		return nil, err
	}

	a := activity.New(ctx, api.ActivityACMECreateAccount, &service.CreateACMEAccountParams{
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
