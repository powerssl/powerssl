package worker

import (
	"context"

	"github.com/google/wire"
	activity2 "go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	apiserver2 "powerssl.dev/backend/context"
	"powerssl.dev/backend/temporal"
	"powerssl.dev/backend/vault"
	"powerssl.dev/controller/internal/activity"
	"powerssl.dev/sdk/apiserver"
	activity3 "powerssl.dev/workflow/activity"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, apiserverClient *apiserver.Client, vaultClient *vault.Client, temporalClient client.Client) F {
	return func() error {
		backgroundActivityContext := context.Background()
		backgroundActivityContext = apiserver2.SetAPIClient(backgroundActivityContext, apiserverClient)
		backgroundActivityContext = apiserver2.SetVaultClient(backgroundActivityContext, vaultClient)
		worker.EnableVerboseLogging(true)
		w := worker.New(temporalClient, temporal.ControllerTaskQueue, worker.Options{
			BackgroundActivityContext: backgroundActivityContext,
		})
		w.RegisterActivityWithOptions(activity.CreateACMEAccount, activity2.RegisterOptions{
			Name: activity3.CreateACMEAccount,
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
