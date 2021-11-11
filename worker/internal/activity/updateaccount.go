package activity

import (
	"context"

	temporalactivity "go.temporal.io/sdk/activity"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	apiv1 "powerssl.dev/api/apiserver/v1"
	context2 "powerssl.dev/backend/context"
	"powerssl.dev/workflow/activity"
)

func UpdateAccount(ctx context.Context, params *activity.UpdateAccountParams) (_ *activity.UpdateAccountResults, err error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("UpdateAccount", params.ToKeyVals()...)

	var acmeAccount *apiv1.ACMEAccount
	updateMask, err := fieldmaskpb.New(acmeAccount, params.UpdateMask...)
	if err != nil {
		return nil, err
	}
	if acmeAccount, err = context2.GetAPIClient(ctx).ACMEAccount.Update(ctx, &apiv1.UpdateACMEAccountRequest{
		Name:        params.Name,
		UpdateMask:  updateMask,
		AcmeAccount: params.ACMEAccount,
	}); err != nil {
		return nil, err
	}

	return &activity.UpdateAccountResults{
		ACMEAccount: acmeAccount,
	}, nil
}
