package activity

import (
	"context"

	temporalactivity "go.temporal.io/sdk/activity"

	"powerssl.dev/backend/apiserver"
	"powerssl.dev/sdk/apiserver/api"
	"powerssl.dev/workflow/activity"
)

func UpdateAccount(ctx context.Context, params *activity.UpdateAccountParams) (_ *activity.UpdateAccountResults, err error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("UpdateAccount", params.ToKeyVals()...)

	var acmeAccount *api.ACMEAccount
	if acmeAccount, err = apiserver.GetClient(ctx).ACMEAccount.Update(ctx, params.Name, params.UpdateMask, params.ACMEAccount); err != nil {
		return nil, err
	}

	return &activity.UpdateAccountResults{
		ACMEAccount: acmeAccount,
	}, nil
}
