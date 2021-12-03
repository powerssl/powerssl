package server

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	ConfigDir string   `flag:"configDir" flag-val:"config" flag-desc:"Config directory to load a set of yaml config files from" validate:"required"`
	Env       string   `flag:"env" flag-val:"development" flag-desc:"Environment is one of the input params ex-development" validate:"required"`
	Services  []string `flag:"services" flag-val:"frontend,history,matching,worker" flag-desc:"Service(s) to start" validate:"gt=0,required"`
	Zone      string   `flag:"zone" flag-desc:"Zone is another input param"`
}

func (cfg *Config) PreValidate(_ *validator.Validate) {}
