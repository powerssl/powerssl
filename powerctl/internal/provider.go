package internal

import (
	"github.com/google/wire"

	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/sdk/apiserver"
)

var Provider = wire.NewSet(
	ConfigFields,
	apiserver.Provider,
	log.Provider,
	telemetry.Provider,
)
