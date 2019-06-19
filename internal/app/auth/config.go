package auth

import (
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	Addr              string `validate:"required"`
	JWTPrivateKeyFile string `validate:"required"`
	MetricsAddr       string
	WebAppURI         string `validate:"required"`
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(cfg)
}
