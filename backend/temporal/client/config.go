package client

import (
	"time"

	temporalconverter "go.temporal.io/sdk/converter"
)

type Config struct {
	Component          string                          `flag:"-"`
	CAFile             string                          `flag:"caFile;;;temporal CA file"`
	DataConverter      temporalconverter.DataConverter `flag:"-"`
	DisableHealthCheck bool                            `flag:"disableHealthCheck;;;temporal disable health check"`
	HealthCheckTimeout time.Duration                   `flag:"healthCheckTimeout;;;temporal health check timeout"`
	HostPort           string                          `flag:"hostPort;;;temporal host port" validate:"required,hostname_port"`
	Identity           string                          `flag:"identity;;;temporal identity"`
	Namespace          string                          `flag:"namespace;;;temporal namespace" validate:"required"`
	TLSCertFile        string                          `flag:"tlsCertFile;;;temporal TLS cert file"`
	TLSKeyFile         string                          `flag:"tlsKeyFile;;;temporal TLS key file"`
}
