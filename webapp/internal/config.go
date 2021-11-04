package internal

import (
	"github.com/go-playground/validator/v10"

	validator2 "powerssl.dev/common/validator"
)

type Config struct {
	APIServer struct {
		Addr string
	}
	Addr string `validate:"required,hostname_port"`
	Auth struct {
		URI string
	}
	GRPCWeb struct {
		URI string
	}
	Insecure bool
	Metrics  struct {
		Addr string
	}
	TLS struct {
		CertFile       string `mapstructure:"cert-file"`
		PrivateKeyFile string `mapstructure:"private-key-file"`
	}
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return validator2.ValidateConfig(validate, cfg)
}
