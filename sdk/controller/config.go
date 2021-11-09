package controller

import "powerssl.dev/common/transport"

type Config struct {
	AuthToken string                 `flag:"authToken;;;controller client addr"`
	Client    transport.ClientConfig `flag:"client"`
}
