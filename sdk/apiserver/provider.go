package apiserver

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewClient,
)
