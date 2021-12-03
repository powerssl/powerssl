package repository

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	ConnString string `flag:"connString" flag-desc:"db conn string" validate:"required"`
}

func (cfg *Config) PreValidate(_ *validator.Validate) {}
