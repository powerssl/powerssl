// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"context"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
	"powerssl.dev/common/transport"
	"powerssl.dev/grpcgateway/internal/server"
)

// Injectors from wire.go:

func Initialize(ctx context.Context, cfg *Config) ([]func() error, func(), error) {
	config := cfg.Log
	logger, cleanup, err := log.Provide(config)
	if err != nil {
		return nil, nil, err
	}
	f := interrupthandler.Provide(ctx, logger)
	serverConfig := cfg.Server
	transportConfig := cfg.APIServerClient
	v := transport.NoDialOptions()
	clientConn, err := transport.New(ctx, transportConfig, v...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	serverF := server.Provide(ctx, serverConfig, logger, clientConn)
	telemetryConfig := cfg.Telemetry
	telemeter, cleanup2, err := telemetry.Provide(ctx, telemetryConfig, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	telemetryF := telemetry.ProvideF(ctx, telemeter)
	v2 := Provide(f, serverF, telemetryF)
	return v2, func() {
		cleanup2()
		cleanup()
	}, nil
}
