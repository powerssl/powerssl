package log // import "powerssl.dev/common/log"

import "go.uber.org/zap"

const production = "production"

type Config struct {
	Env     string       `flag:"env;e;production;environment" validate:"required"`
	Options []zap.Option `flag:"-"`
}

func (c *Config) Production() bool {
	return c.Env == production
}
