package oauth2

import "github.com/google/wire"

var Provider = wire.NewSet(
	New,
)
