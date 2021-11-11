package apiserver // import "powerssl.dev/sdk/apiserver"

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewClient,
)
