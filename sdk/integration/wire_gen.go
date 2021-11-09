// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package integration

import (
	"context"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"
	"powerssl.dev/common/metrics"
	"powerssl.dev/common/tracer"
	"powerssl.dev/sdk/controller"
	"powerssl.dev/sdk/integration/acme"
	"powerssl.dev/sdk/integration/dns"
)

// Injectors from wire.go:

func InitializeACME(ctx context.Context, cfg *Config, handler acme.Integration) ([]func() error, func(), error) {
	config := cfg.Log
	sugaredLogger, cleanup, err := log.Provide(config)
	if err != nil {
		return nil, nil, err
	}
	f := interrupthandler.Provide(ctx, sugaredLogger)
	metricsConfig := cfg.Metrics
	metricsF := metrics.Provide(ctx, metricsConfig, sugaredLogger)
	integrationConfig := &cfg.Integration
	controllerConfig := cfg.ControllerClient
	tracerConfig := cfg.Tracer
	opentracingTracer, cleanup2, err := tracer.Provide(tracerConfig, sugaredLogger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	grpcClient, err := controller.NewGRPCClient(ctx, controllerConfig, sugaredLogger, opentracingTracer)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	integrationF := ProvideACME(ctx, integrationConfig, sugaredLogger, grpcClient, handler)
	v := Provide(f, metricsF, integrationF)
	return v, func() {
		cleanup2()
		cleanup()
	}, nil
}

func InitializeDNS(ctx context.Context, cfg *Config, handler dns.Integration) ([]func() error, func(), error) {
	config := cfg.Log
	sugaredLogger, cleanup, err := log.Provide(config)
	if err != nil {
		return nil, nil, err
	}
	f := interrupthandler.Provide(ctx, sugaredLogger)
	metricsConfig := cfg.Metrics
	metricsF := metrics.Provide(ctx, metricsConfig, sugaredLogger)
	integrationConfig := &cfg.Integration
	controllerConfig := cfg.ControllerClient
	tracerConfig := cfg.Tracer
	opentracingTracer, cleanup2, err := tracer.Provide(tracerConfig, sugaredLogger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	grpcClient, err := controller.NewGRPCClient(ctx, controllerConfig, sugaredLogger, opentracingTracer)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	integrationF := ProvideDNS(ctx, integrationConfig, sugaredLogger, grpcClient, handler)
	v := Provide(f, metricsF, integrationF)
	return v, func() {
		cleanup2()
		cleanup()
	}, nil
}
