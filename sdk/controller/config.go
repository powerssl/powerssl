package controller // import "powerssl.dev/sdk/controller"

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/transport"
)

type Config struct {
	AuthToken string           `flag:"authToken" flag-desc:"controller client addr" validate:"required"`
	Client    transport.Config `flag:"client"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Client.PreValidate(validate)
}
