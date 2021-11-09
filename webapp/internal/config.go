package internal

import (
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/webapp/internal/server"
)

type Config struct {
	Log     log.Config     `flag:"log"`
	Metrics metrics.Config `flag:"metrics"`
	Server  server.Config  `flag:"server"`
}

func (cfg *Config) Defaults() {
	return
}
