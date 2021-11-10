package server

import (
	"context"
	"fmt"

	temporalconfig "go.temporal.io/server/common/config"
	"go.temporal.io/server/temporal"
	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg *Config, logger *zap.SugaredLogger) error {
	config, err := temporalconfig.LoadConfig(cfg.Env, cfg.ConfigDir, cfg.Zone)
	if err != nil {
		return fmt.Errorf("unable to load configuration: %w", err)
	}

	s := temporal.NewServer(
		temporal.ForServices(cfg.Services),
		temporal.InterruptOn(interruptCh(ctx)),
		temporal.WithConfig(config),
		temporal.WithLogger(newLogger(logger)),
	)

	if err = s.Start(); err != nil {
		return fmt.Errorf("unable to start server: %w", err)
	}
	return nil

}

func interruptCh(ctx context.Context) chan interface{} {
	ret := make(chan interface{}, 1)
	go func() {
		s := <-ctx.Done()
		ret <- s
		close(ret)
	}()
	return ret
}
