package client

import (
	"github.com/google/wire"
	"go.temporal.io/sdk/client"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

func Provide(cfg Config, logger log.Logger) (client.Client, error) {
	logger = logger.With("component", "temporal")
	c, err := New(cfg, logger)
	if err != nil {
		return nil, err
	}
	return c, nil
}
