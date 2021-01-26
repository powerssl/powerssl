package workflow

import (
	"time"

	temporalworkflow "go.temporal.io/sdk/workflow"

	"powerssl.dev/powerssl/internal/app/worker/activity"
)

func CreateAccount(ctx temporalworkflow.Context, account, directoryURL string, termsOfServiceAgreed bool, contacts []string) error {
	ctx = temporalworkflow.WithActivityOptions(ctx, temporalworkflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout: time.Hour,
	})

	var acmeAccount interface{}
	if err := temporalworkflow.ExecuteActivity(ctx, activity.CreateACMEAccount, directoryURL, termsOfServiceAgreed, contacts).Get(ctx, &acmeAccount); err != nil {
		return err
	}

	updateMask := []string{"account_url"}
	if err := temporalworkflow.ExecuteActivity(ctx, activity.UpdateAccount, account, updateMask, account).Get(ctx, &acmeAccount); err != nil {
		return err
	}

	return nil
}
