package workflow

import (
	"powerssl.dev/powerssl/pkg/apiserver/api"
	"time"

	"go.temporal.io/sdk/temporal"
	temporalworkflow "go.temporal.io/sdk/workflow"

	temporalutil "powerssl.dev/powerssl/internal/pkg/temporal"
	"powerssl.dev/powerssl/internal/pkg/temporal/activity"
)

type CreateAccountParams struct {
	Account              string
	DirectoryURL         string
	TermsOfServiceAgreed bool
	Contacts             []string
}

func (p *CreateAccountParams) ToKeyVals() []interface{} {
	return []interface{}{
		"Account", p.Account,
		"DirectoryURL", p.DirectoryURL,
		"TermsOfServiceAgreed", p.TermsOfServiceAgreed,
		"Contacts", p.Contacts,
	}
}

func CreateAccount(ctx temporalworkflow.Context, params CreateAccountParams) error {
	logger := temporalworkflow.GetLogger(ctx)
	logger.Info("CreateAccount", params.ToKeyVals()...)

	ctx = temporalworkflow.WithActivityOptions(ctx, temporalworkflow.ActivityOptions{
		ScheduleToStartTimeout: 5 * time.Second,
		StartToCloseTimeout:    10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    10 * time.Second,
			MaximumAttempts:    10,
		},
	})
	controllerCtx := temporalworkflow.WithTaskQueue(ctx, temporalutil.ControllerTaskQueue)
	workerCtx := temporalworkflow.WithTaskQueue(ctx, temporalutil.WorkerTaskQueue)

	if err := temporalworkflow.ExecuteActivity(workerCtx, activity.CreateVaultTransitKey, activity.CreateVaultTransitKeyParams{
		Name: params.Account,
	}).Get(ctx, nil); err != nil {
		return err
	}

	var createACMEAccountResult activity.CreateACMEAccountResults
	if err := temporalworkflow.ExecuteActivity(controllerCtx, activity.CreateACMEAccount, &activity.CreateACMEAccountParams{
		DirectoryURL:         params.DirectoryURL,
		TermsOfServiceAgreed: params.TermsOfServiceAgreed,
		Contacts:             params.Contacts,
	}).Get(ctx, &createACMEAccountResult); err != nil {
		return err
	}

	if err := temporalworkflow.ExecuteActivity(workerCtx, activity.UpdateAccount, activity.UpdateAccountParams{
		Name:       params.Account,
		UpdateMask: []string{"account_url"},
		ACMEAccount: &api.ACMEAccount{
			AccountURL: createACMEAccountResult.Account.URL,
		},
	}).Get(ctx, nil); err != nil {
		return err
	}

	return nil
}
