package internal

import (
	"github.com/google/wire"

	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/server"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
)

var Provider = wire.NewSet(
	interrupthandler.Provider,
	log.Provider,
	metrics.Provider,
	oauth2.Provider,
	Provide,
	server.Provider,
	wire.FieldsOf(new(*Config), "Log", "Metrics", "OAuth2", "Server"),
)

func Provide(interruptHandlerF interrupthandler.F, metricsF metrics.F, serverF server.F) []func() error {
	return []func() error{
		interruptHandlerF,
		metricsF,
		serverF,
	}
}
