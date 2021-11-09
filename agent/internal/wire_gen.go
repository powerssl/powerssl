// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"context"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/apiserver"
)

// Injectors from wire.go:

func Initialize(ctx context.Context, cfg *Config) ([]func() error, func(), error) {
	config := cfg.Log
	sugaredLogger, cleanup, err := log.Provide(config)
	if err != nil {
		return nil, nil, err
	}
	f := interrupthandler.Provide(ctx, sugaredLogger)
	apiserverConfig := cfg.APIServerClient
	tracerConfig := cfg.Tracer
	opentracingTracer, cleanup2, err := tracer.Provide(tracerConfig, sugaredLogger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	client, err := apiserver.Provide(ctx, apiserverConfig, sugaredLogger, opentracingTracer)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	v := Provide(f, client)
	return v, func() {
		cleanup2()
		cleanup()
	}, nil
}
