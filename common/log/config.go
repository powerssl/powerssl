package log // import "powerssl.dev/common/log"

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const production = "production"

type Config struct {
	Env     string       `flag:"env" flag-short:"e" flag-val:"production" flag-desc:"environment" validate:"required"`
	Options []zap.Option `flag:"-"`
}

func (cfg *Config) PreValidate(_ *validator.Validate) {}

func (cfg *Config) Production() bool {
	return cfg.Env == production
}
