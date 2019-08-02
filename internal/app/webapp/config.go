package webapp

import (
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	APIAddr     string `validate:"required"`
	Addr        string `validate:"required"`
	AuthURI     string `validate:"required"`
	GRPCWebURI  string `validate:"required"`
	MetricsAddr string
}

func (cfg *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(cfg)
}
