package util

import (
	"os"

	"github.com/go-kit/kit/log"
	"gopkg.in/go-playground/validator.v9"
)

type Config interface {
	Validate() error
}

func ValidateConfig(cfg Config, logger log.Logger) {
	if err := cfg.Validate(); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logger.Log(
				"namespace", err.Namespace(),
				"field", err.Field(),
				"struct_namespace", err.StructNamespace(),
				"struct_field", err.StructField(),
				"tag", err.Tag(),
				"actual_tag", err.ActualTag(),
				"kind", err.Kind(),
				"type", err.Type(),
				"value", err.Value(),
				"param", err.Param())
		}
		os.Exit(1)
	}
}
