package worker

import (
	"context"

	"github.com/google/wire"
	temporalactivity "go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	contextutil "powerssl.dev/backend/context"
	"powerssl.dev/backend/temporal"
	"powerssl.dev/backend/vault"
	"powerssl.dev/sdk"
	"powerssl.dev/sdk/apiserver"
	workflowactivity "powerssl.dev/workflow/activity"

	"powerssl.dev/controller/internal/activity"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, apiserverClient *apiserver.Client, vaultClient *vault.Client, temporalClient client.Client) F {
	return func() error {
		backgroundActivityContext := context.Background()
		backgroundActivityContext = sdk.SetAPIClient(backgroundActivityContext, apiserverClient)
		backgroundActivityContext = contextutil.SetVaultClient(backgroundActivityContext, vaultClient)
		worker.EnableVerboseLogging(true)
		w := worker.New(temporalClient, temporal.ControllerTaskQueue, worker.Options{
			BackgroundActivityContext: backgroundActivityContext,
		})
		w.RegisterActivityWithOptions(activity.CreateACMEAccount, temporalactivity.RegisterOptions{
			Name: workflowactivity.CreateACMEAccount,
		})
		if err := w.Start(); err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			w.Stop()
			return ctx.Err()
		}
	}
}
