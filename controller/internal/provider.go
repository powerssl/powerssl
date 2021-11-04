package internal

import (
	"context"

	temporalactivity "go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"

	backendapiserver "powerssl.dev/backend/apiserver"
	"powerssl.dev/backend/temporal"
	"powerssl.dev/backend/temporal/client"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/backend/vault"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"
	"powerssl.dev/controller/internal/activity"
	"powerssl.dev/sdk/apiserver"
	sharedactivity "powerssl.dev/workflow/activity"
)

const component = "powerssl-controller"

func provideRunnerF(interruptHandlerF interrupthandler.F, metricsServerF metrics.F, serverF backendtransport.ServerF, workerF workerF) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsServerF,
		serverF,
		workerF,
	}
}

func provideTemporalClientComponent() client.TemporalClientComponent {
	return component
}

func provideTracingComponent() tracing.TracerComponent {
	return component
}

type workerF func() error

func provideWorkerF(ctx context.Context, apiserverClient *apiserver.Client, vaultClient *vault.Client, temporalClient client.Client) workerF {
	return func() error {
		backgroundActivityContext := context.Background()
		backgroundActivityContext = backendapiserver.SetClient(backgroundActivityContext, apiserverClient)
		backgroundActivityContext = vault.SetClient(backgroundActivityContext, vaultClient)
		worker.EnableVerboseLogging(true)
		w := worker.New(temporalClient, temporal.ControllerTaskQueue, worker.Options{
			BackgroundActivityContext: backgroundActivityContext,
		})
		w.RegisterActivityWithOptions(activity.CreateACMEAccount, temporalactivity.RegisterOptions{
			Name: sharedactivity.CreateACMEAccount,
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
