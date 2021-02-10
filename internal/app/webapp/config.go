package webapp

import (
	"github.com/go-playground/validator/v10"

	"powerssl.dev/powerssl/internal/pkg/util"
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
	return util.ValidateConfig(validate, cfg)
}
