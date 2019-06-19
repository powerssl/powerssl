package integration

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.io/powerssl/internal/pkg/util"
)

type ControllerClientConfig = util.ClientConfig

type Config struct {
	AuthToken              string `validate:"required"`
	ControllerClientConfig *ControllerClientConfig
	MetricsAddr            string
	Tracer                 string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(util.ClientConfigValidator, util.ClientConfig{})
	return validate.Struct(cfg)
}
