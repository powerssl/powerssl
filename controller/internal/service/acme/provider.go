package acme

import (
	"github.com/google/wire"

	apiv1 "powerssl.dev/api/controller/v1"
)

var Provider = wire.NewSet(
	New,
	wire.Bind(new(apiv1.ACMEServiceServer), new(*Service)),
)
