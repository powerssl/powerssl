package apiserver // import "powerssl.dev/sdk/apiserver"

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common/transport"
)

type Config struct {
	AuthToken string           `flag:"authToken" flag-desc:"apiserver auth token" validate:"required"`
	Client    transport.Config `flag:"client"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Client.PreValidate(validate)
}
