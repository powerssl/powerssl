package internal

import (
	"fmt"

	temporalconfig "go.temporal.io/server/common/service/config"
	"go.temporal.io/server/temporal"
)

func Run(cfg *Config) error {
	config, err := temporalconfig.LoadConfig(cfg.Env, cfg.ConfigDir, cfg.Zone)
	if err != nil {
		return fmt.Errorf("unable to load configuration: %v", err)
	}

	s := temporal.NewServer(
		temporal.ForServices(cfg.Services),
		temporal.InterruptOn(temporal.InterruptCh()),
		temporal.WithConfig(config),
	)

	if err = s.Start(); err != nil {
		return fmt.Errorf("unable to start server: %v", err)
	}
	return nil
}
