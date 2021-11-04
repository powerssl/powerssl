package internal

import (
	"github.com/go-playground/validator/v10"

	validator2 "powerssl.dev/common/validator"
)

type Config struct {
	ConfigDir string   `validate:"required" mapstructure:"config-dir"`
	Env       string   `validate:"required"`
	Services  []string `validate:"gt=0,required"`
	Zone      string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return validator2.ValidateConfig(validate, cfg)
}
