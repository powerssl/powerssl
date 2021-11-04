package vault

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	New,
)
