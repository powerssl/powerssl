package integration // import "powerssl.dev/sdk/integration"

import (
	"context"
	"time"

	"github.com/google/wire"

	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/integration/acme"
	"powerssl.dev/sdk/integration/dns"
	"powerssl.dev/sdk/integration/internal"

	"powerssl.dev/sdk/controller"
)

var ProviderACME = wire.NewSet(
	ProvideACME,
	provider,
)

var ProviderDNS = wire.NewSet(
	ProvideDNS,
	provider,
)

var provider = wire.NewSet(
	ConfigFields,
	Provide,
	controller.Provider,
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	tracer.Provider,
)

type F func() error

func Provide(interruptHandlerF interrupthandler.F, metricsF metrics.F, runnerF F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsF,
		runnerF,
	}
}

func ProvideACME(ctx context.Context, cfg *internal.IntegrationConfig, logger log.Logger, client *controller.GRPCClient, handler acme.Integration) F {
	return func() error {
		integration := internal.NewACME(cfg.Name, logger, client, handler)
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

func ProvideDNS(ctx context.Context, cfg *internal.IntegrationConfig, logger log.Logger, client *controller.GRPCClient, handler dns.Integration) F {
	return func() error {
		integration := internal.NewDNS(cfg.Name, logger, client, handler)
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
