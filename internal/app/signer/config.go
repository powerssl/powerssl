package signer

import (
	"gopkg.in/go-playground/validator.v9"

	"powerssl.io/powerssl/internal/pkg/util"
)

type ServerConfig = util.ServerConfig

type Config struct {
	MetricsAddr  string
	ServerConfig *ServerConfig
	Tracer       string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(util.ServerConfigValidator, util.ServerConfig{})
	return validate.Struct(cfg)
}
