package grpcgateway

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.io/powerssl/internal/pkg/util"
)

type APIServerClientConfig = util.ClientConfig

type Config struct {
	APIServerClientConfig *APIServerClientConfig
	Addr                  string `validate:"required"`
	MetricsAddr           string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(util.ClientConfigValidator, util.ClientConfig{})
	return validate.Struct(cfg)
}
