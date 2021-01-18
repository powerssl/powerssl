package workflow

import (
	"time"

	temporalworkflow "go.temporal.io/sdk/workflow"

	"powerssl.dev/powerssl/internal/app/controller/activity"
)

func CreateAccount(ctx temporalworkflow.Context, directoryURL string, termsOfServiceAgreed bool, contacts []string) error {
	ctx = temporalworkflow.WithActivityOptions(ctx, temporalworkflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout: time.Hour,
	})
	if err := temporalworkflow.ExecuteActivity(ctx, activity.CreateAccount, directoryURL, termsOfServiceAgreed, contacts).Get(ctx, nil); err != nil {
		return err
	}

	return nil
}
