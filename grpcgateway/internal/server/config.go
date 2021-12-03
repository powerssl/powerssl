package server

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Addr string `flag:"addr" flag-desc:"server addr"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {}
