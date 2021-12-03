package telemetry

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Component string       `flag:"-"`
	Meter     MeterConfig  `flag:"meter"`
	Tracer    TracerConfig `flag:"tracer"`
}

func (cfg *Config) PreValidate(validate *validator.Validate) {
	cfg.Meter.PreValidate(validate)
	cfg.Tracer.PreValidate(validate)
}

type MeterConfig struct {
	Addr     string `flag:"addr" flag-desc:"metrics addr"`
	Exporter string `flag:"exporter" flag-desc:"metrics exporter"`
}

func (cfg *MeterConfig) PreValidate(_ *validator.Validate) {}

type TracerConfig struct {
	Disabled bool `flag:"disabled" flag-desc:"disable tracer"`
}

func (cfg *TracerConfig) PreValidate(_ *validator.Validate) {}
