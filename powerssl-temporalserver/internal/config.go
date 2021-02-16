package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/util"
)

type Config struct {
	ConfigDir string   `validate:"required" mapstructure:"config-dir"`
	Env       string   `validate:"required"`
	Services  []string `validate:"gt=0,required"`
	Zone      string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return util.ValidateConfig(validate, cfg)
}
