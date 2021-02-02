package workflow

import (
	"time"

	temporal "go.temporal.io/sdk/temporal"
	temporalworkflow "go.temporal.io/sdk/workflow"

	"powerssl.dev/powerssl/internal/pkg/temporal/activity"
)

type CreateAccountParams struct {
	Account              string
	DirectoryURL         string
	TermsOfServiceAgreed bool
	Contacts             []string
}

func CreateAccount(ctx temporalworkflow.Context, params CreateAccountParams) error {
	ctx = temporalworkflow.WithActivityOptions(ctx, temporalworkflow.ActivityOptions{
		// ScheduleToStartTimeout - The queue timeout before the activity starts executed.
		ScheduleToStartTimeout: time.Hour,

		// StartToCloseTimeout - The timeout from the start of execution to end of it.
		StartToCloseTimeout: time.Hour,

		// RetryPolicy specifies how to retry an Activity if an error occurs.
		// More details are available at docs.temporal.io.
		// RetryPolicy is optional. If one is not specified a default RetryPolicy is provided by the server.
		// The default RetryPolicy provided by the server specifies:
		// - InitialInterval of 1 second
		// - BackoffCoefficient of 2.0
		// - MaximumInterval of 100 x InitialInterval
		// - MaximumAttempts of 0 (unlimited)
		// To disable retries set MaximumAttempts to 1.
		// The default RetryPolicy provided by the server can be overridden by the dynamic config.
		RetryPolicy: &temporal.RetryPolicy{
			// Backoff interval for the first retry. If BackoffCoefficient is 1.0 then it is used for all retries.
			// Required, no default value.
			InitialInterval: time.Second,

			// Coefficient used to calculate the next retry backoff interval.
			// The next retry interval is previous interval multiplied by this coefficient.
			// Must be 1 or larger. Default is 2.0.
			BackoffCoefficient: 2.0,

			// Maximum backoff interval between retries. Exponential backoff leads to interval increase.
			// This value is the cap of the interval. Default is 100x of initial interval.
			MaximumInterval: 100 * time.Second,

			// Maximum number of attempts. When exceeded the retries stop even if not expired yet.
			// If not set or set to 0, it means unlimited, and rely on activity ScheduleToCloseTimeout to stop.
			MaximumAttempts: 0,
		},
	})

	var createACMEAccountResult activity.CreateACMEAccountResults
	if err := temporalworkflow.ExecuteActivity(ctx, activity.CreateACMEAccount, activity.CreateACMEAccountParams{
		DirectoryURL:         params.DirectoryURL,
		TermsOfServiceAgreed: params.TermsOfServiceAgreed,
		Contacts:             params.Contacts,
	}).Get(ctx, &createACMEAccountResult); err != nil {
		return err
	}

	if err := temporalworkflow.ExecuteActivity(ctx, activity.UpdateAccount, activity.UpdateAccountParams{
		Name:        params.Account,
		UpdateMask:  []string{"account_url"},
		ACMEAccount: createACMEAccountResult.ACMEAccount,
	}).Get(ctx, nil); err != nil {
		return err
	}

	return nil
}
