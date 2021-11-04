package certificate

import (
	"github.com/google/wire"

	apiv1 "powerssl.dev/api/apiserver/v1"
)

var Provider = wire.NewSet(
	New,
	wire.Bind(new(apiv1.CertificateServiceServer), new(*Service)),
)
