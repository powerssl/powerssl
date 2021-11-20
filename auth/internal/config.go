package internal

import (
	"github.com/google/wire"

	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/server"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
)

var ConfigFields = wire.FieldsOf(new(*Config), "Log", "Metrics", "OAuth2", "Server")

type Config struct {
	Log     log.Config     `flag:"log"`
	Server  server.Config  `flag:"server"`
	OAuth2  oauth2.Config  `flag:"oauth2"`
	Metrics metrics.Config `flag:"metrics"`
}

func (c *Config) Defaults() {
	return
}
