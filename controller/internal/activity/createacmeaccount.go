package activity

import (
	"context"

	temporalactivity "go.temporal.io/sdk/activity"

	apiv1 "powerssl.dev/api/controller/v1"
	ctxutil "powerssl.dev/backend/context"
	"powerssl.dev/controller/internal/service/acme"
	"powerssl.dev/controller/internal/service/integration/activity"
	sharedactivity "powerssl.dev/workflow/activity"
)

func CreateACMEAccount(ctx context.Context, params *sharedactivity.CreateACMEAccountParams) (_ *sharedactivity.CreateACMEAccountResults, err error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("CreateACMEAccount", params.ToKeyVals()...)

	var keyToken string
	if keyToken, err = ctxutil.GetVaultClient(ctx).ExportWrappedTransitKey(ctx, params.KeyName); err != nil {
		return nil, err
	}

	a := activity.New(ctx, apiv1.Activity_ACME_CREATE_ACCOUNT, &acme.CreateACMEAccountParams{
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
