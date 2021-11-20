package integration // import "powerssl.dev/sdk/integration"

import (
	"fmt"

	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/controller"
	"powerssl.dev/sdk/integration/internal"
)

var ConfigFields = wire.FieldsOf(new(*Config), "ControllerClient", "Integration", "Log", "Metrics", "Tracer")

type Config struct {
	ControllerClient controller.Config          `flag:"controllerClient"`
	Integration      internal.IntegrationConfig `flag:"integration"`
	Log              log.Config                 `flag:"log"`
	Metrics          metrics.Config             `flag:"metrics"`
	Tracer           tracer.Config              `flag:"tracer"`
}

func (cfg *Config) Defaults() {
	cfg.Tracer.Component = fmt.Sprintf("powerssl-integration-%s", cfg.Integration.Kind)
}
