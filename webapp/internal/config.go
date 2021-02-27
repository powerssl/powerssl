package internal

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/common"
)

type Config struct {
	APIServer struct {
		Addr string `validate:"required,hostname_port"`
	}
	Addr string `validate:"required,hostname_port"`
	Auth struct {
		URI string `validate:"required,uri"`
	}
	GRPCWeb struct {
		URI string `validate:"required,uri"`
	}
	Metrics struct {
		Addr string
	}
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return common.ValidateConfig(validate, cfg)
}