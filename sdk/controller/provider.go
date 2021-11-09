package controller // import "powerssl.dev/sdk/controller"

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewGRPCClient,
)
