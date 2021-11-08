package integration // import "powerssl.dev/sdk/integration"

import (
	"context"
	"fmt"
	"time"

	"github.com/google/wire"
	"go.uber.org/zap"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracing"
	"powerssl.dev/sdk/integration/acme"
	"powerssl.dev/sdk/integration/dns"
	"powerssl.dev/sdk/integration/internal"

	"powerssl.dev/sdk/controller"
)

var provider = wire.NewSet(
	Provide,
	ProvideTracingComponent,
	controller.Provider,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	tracing.Provider,
	wire.FieldsOf(new(*Config), "AuthToken", "ControllerClientConfig", "Integration", "Metrics", "Tracer"),
)

var ProviderACME = wire.NewSet(
	provider,
	ProvideRunnerACMEF,
)

var ProviderDNS = wire.NewSet(
	provider,
	ProvideRunnerDNSF,
)

type F func() error

func Provide(interruptHandlerF interrupthandler.F, metricsServerF metrics.F, runnerF F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsServerF,
		runnerF,
	}
}

func ProvideRunnerACMEF(ctx context.Context, cfg *internal.IntegrationConfig, logger *zap.SugaredLogger, client *controller.GRPCClient, handler acme.Integration) F {
	name := cfg.Name
	return func() error {
		integration := internal.NewACME(name, logger, client, handler)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				logger.Error(integration.Run(ctx))
				time.Sleep(time.Second)
			}
		}
	}
}

func ProvideRunnerDNSF(ctx context.Context, cfg *internal.IntegrationConfig, logger *zap.SugaredLogger, client *controller.GRPCClient, handler dns.Integration) F {
	name := cfg.Name
	return func() error {
		integration := internal.NewDNS(name, logger, client, handler)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				logger.Error(integration.Run(ctx))
				time.Sleep(time.Second)
			}
		}
	}
}

func ProvideTracingComponent(cfg *internal.IntegrationConfig) tracing.TracerComponent {
	return tracing.TracerComponent(fmt.Sprintf("powerssl-integration-%s", cfg.Kind))
}
