package internal

import (
	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/server"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
)

type Config struct {
	Log     log.Config     `flag:"log"`
	Server  server.Config  `flag:"server"`
	OAuth2  oauth2.Config  `flag:"oauth2"`
	Metrics metrics.Config `flag:"metrics"`
}

func (c *Config) Defaults() {
	return
}
