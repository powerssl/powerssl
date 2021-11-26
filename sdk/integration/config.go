package integration // import "powerssl.dev/sdk/integration"

import (
	"fmt"

	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"

	"powerssl.dev/sdk/controller"
	"powerssl.dev/sdk/integration/internal"
)

var ConfigFields = wire.FieldsOf(new(*Config), "ControllerClient", "Integration", "Log", "Telemetry")

type Config struct {
	ControllerClient controller.Config          `flag:"controllerClient"`
	Integration      internal.IntegrationConfig `flag:"integration"`
	Log              log.Config                 `flag:"log"`
	Telemetry        telemetry.Config           `flag:"telemetry"`
}

func (cfg *Config) Defaults() {
	cfg.Telemetry.Component = fmt.Sprintf("powerssl-integration-%s", cfg.Integration.Kind)
}
