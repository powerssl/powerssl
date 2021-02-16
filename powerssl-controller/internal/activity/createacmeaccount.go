package activity

import (
	"context"

	temporalactivity "go.temporal.io/sdk/activity"

	"powerssl.dev/controller/internal/integration/activity"
	"powerssl.dev/sdk/controller/api"
	sharedactivity "powerssl.dev/workflow/activity"
)

func CreateACMEAccount(ctx context.Context, params *sharedactivity.CreateACMEAccountParams) (*sharedactivity.CreateACMEAccountResults, error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("CreateACMEAccount", params.ToKeyVals()...)

	a := activity.New(ctx, api.ActivityACMECreateAccount, params)
	if err := a.Execute(ctx); err != nil {
		return nil, err
	}

	return nil, temporalactivity.ErrResultPending
}
