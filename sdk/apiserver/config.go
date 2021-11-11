package apiserver // import "powerssl.dev/sdk/apiserver"

import "powerssl.dev/common/transport"

type Config struct {
	AuthToken string           `flag:"authToken;;;apiserver client addr"`
	Client    transport.Config `flag:"client"`
}
