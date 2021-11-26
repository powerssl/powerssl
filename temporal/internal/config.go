package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/temporal/internal/server"
)

var ConfigFields = wire.FieldsOf(new(*Config), "Log", "Server")

type Config struct {
	Log    log.Config    `flag:"log"`
	Server server.Config `flag:"server"`
}

func (cfg *Config) Defaults() {}
