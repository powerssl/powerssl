package internal

import (
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/transport"
	"powerssl.dev/grpcgateway/internal/server"
)

type Config struct {
	APIServerClient transport.ClientConfig `flag:"apiServerClient"`
	Log             log.Config             `flag:"log"`
	Metrics         metrics.Config         `flag:"metrics"`
	Server          server.Config          `flag:"server"`
}

func (cfg *Config) Defaults() {
	return
}
