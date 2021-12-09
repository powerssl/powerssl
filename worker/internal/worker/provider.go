package worker

import (
	"context"

	"github.com/google/wire"
	temporalactivity "go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	temporalworkflow "go.temporal.io/sdk/workflow"

	backendapiserver "powerssl.dev/backend/context"
	"powerssl.dev/backend/temporal"
	"powerssl.dev/backend/vault"
	"powerssl.dev/sdk"
	"powerssl.dev/sdk/apiserver"
	sharedworkflow "powerssl.dev/workflow"
	sharedactivity "powerssl.dev/workflow/activity"

	"powerssl.dev/worker/internal/activity"
	"powerssl.dev/worker/internal/workflow"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, apiserverClient *apiserver.Client, vaultClient *vault.Client, temporalClient client.Client) F {
	return func() error {
		worker.EnableVerboseLogging(true)
		backgroundActivityContext := context.Background()
		backgroundActivityContext = sdk.SetAPIClient(backgroundActivityContext, apiserverClient)
		backgroundActivityContext = backendapiserver.SetVaultClient(backgroundActivityContext, vaultClient)
		w := worker.New(temporalClient, temporal.WorkerTaskQueue, worker.Options{
			BackgroundActivityContext: backgroundActivityContext,
		})
		w.RegisterActivityWithOptions(activity.CreateVaultTransitKey, temporalactivity.RegisterOptions{
			Name: sharedactivity.CreateVaultTransitKey,
		})
		w.RegisterActivityWithOptions(activity.UpdateAccount, temporalactivity.RegisterOptions{
			Name: sharedactivity.UpdateAccount,
		})
		w.RegisterWorkflowWithOptions(workflow.CreateAccount, temporalworkflow.RegisterOptions{
			Name: sharedworkflow.CreateAccount,
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
