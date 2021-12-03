package integration // import "powerssl.dev/sdk/integration"

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"

	"powerssl.dev/sdk/controller"
	"powerssl.dev/sdk/integration/internal"
)

var ConfigFields = wire.FieldsOf(new(*Config), "ControllerClient", "Integration", "Log", "Telemetry")

type Config struct {
	ControllerClient controller.Config          `flag:"controllerClient"`
	Integration      internal.IntegrationConfig `flag:"-"`
	Log              log.Config                 `flag:"log"`
	Telemetry        telemetry.Config           `flag:"telemetry"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Telemetry.Component = fmt.Sprintf("powerssl-integration-%s", cfg.Integration.Kind)
	cfg.ControllerClient.PreValidate(validate)
	cfg.Integration.PreValidate(validate)
	cfg.Log.PreValidate(validate)
	cfg.Telemetry.PreValidate(validate)
}
