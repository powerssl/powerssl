package client

import (
	"time"

	"github.com/go-playground/validator/v10"
	temporalconverter "go.temporal.io/sdk/converter"
)

type Config struct {
	Component          string                          `flag:"-"`
	CAFile             string                          `flag:"caFile" flag-desc:"temporal CA file"`
	DataConverter      temporalconverter.DataConverter `flag:"-"`
	DisableHealthCheck bool                            `flag:"disableHealthCheck" flag-desc:"temporal disable health check"`
	HealthCheckTimeout time.Duration                   `flag:"healthCheckTimeout" flag-desc:"temporal health check timeout"`
	HostPort           string                          `flag:"hostPort" flag-desc:"temporal host port" validate:"required,hostname_port"`
	Identity           string                          `flag:"identity" flag-desc:"temporal identity"`
	Namespace          string                          `flag:"namespace" flag-desc:"temporal namespace" validate:"required"`
	TLSCertFile        string                          `flag:"tlsCertFile" flag-desc:"temporal TLS cert file"`
	TLSKeyFile         string                          `flag:"tlsKeyFile" flag-desc:"temporal TLS key file"`
}

func (cfg *Config) PreValidate(_ *validator.Validate) {}
