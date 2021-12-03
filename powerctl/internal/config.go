package internal

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerctl"

var ConfigFields = wire.FieldsOf(new(*Config), "APIServer", "Log", "Telemetry")

type Config struct {
	APIServer apiserver.Config `flag:"apiServer"`
	Log       log.Config       `flag:"log"`
	Output    string           `flag:"output" flag-short:"o" flag-val:"table" flag-desc:"Output format" validate:"oneof=json table yaml"`
	Telemetry telemetry.Config `flag:"telemetry"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Telemetry.Component = component
	cfg.APIServer.PreValidate(validate)
	cfg.Log.PreValidate(validate)
	cfg.Telemetry.PreValidate(validate)
}
