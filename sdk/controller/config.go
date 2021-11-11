package controller // import "powerssl.dev/sdk/controller"

import "powerssl.dev/common/transport"

type Config struct {
	AuthToken string           `flag:"authToken;;;controller client addr"`
	Client    transport.Config `flag:"client"`
}
