package activity

import (
	"context"

	temporalactivity "go.temporal.io/sdk/activity"

	"powerssl.dev/powerssl/internal/pkg/apiserver"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type UpdateAccountParams struct {
	Name        string
	UpdateMask  []string
	ACMEAccount *api.ACMEAccount
}

func (p *UpdateAccountParams) ToKeyVals() []interface{} {
	return []interface{}{
		"Name", p.Name,
		"UpdateMask", p.UpdateMask,
		"ACMEAccount", p.ACMEAccount,
	}
}

type UpdateAccountResults struct {
	ACMEAccount *api.ACMEAccount
}

func UpdateAccount(ctx context.Context, params UpdateAccountParams) (_ *UpdateAccountResults, err error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("UpdateAccount", params.ToKeyVals()...)

	var acmeAccount *api.ACMEAccount
	if acmeAccount, err = apiserver.GetClient(ctx).ACMEAccount.Update(ctx, params.Name, params.UpdateMask, params.ACMEAccount); err != nil {
		return nil, err
	}

	return &UpdateAccountResults{
		ACMEAccount: acmeAccount,
	}, nil
}
