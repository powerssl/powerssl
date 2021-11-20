package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/transport"
	"powerssl.dev/grpcgateway/internal/server"
)

var ConfigFields = wire.FieldsOf(new(*Config), "APIServerClient", "Log", "Metrics", "Server")

type Config struct {
	APIServerClient transport.Config `flag:"apiServerClient"`
	Log             log.Config       `flag:"log"`
	Metrics         metrics.Config   `flag:"metrics"`
	Server          server.Config    `flag:"server"`
}

func (cfg *Config) Defaults() {
	return
}
