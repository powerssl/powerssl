package temporalserver

import (
	"log"
	"fmt"

	"go.temporal.io/server/common/service/config"
	"go.temporal.io/server/temporal"
)

func Run(cfg *Config) {
	config, err := config.LoadConfig(cfg.Env, cfg.ConfigDir, cfg.Zone)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to load configuration: %v.\n", err))
	}

	s := temporal.NewServer(
		temporal.ForServices(cfg.Services),
		temporal.InterruptOn(temporal.InterruptCh()),
		//temporal.WithAuthorizer(newAuthorizer()),
		//temporal.WithClaimMapper(claimMapper),
		temporal.WithConfig(config),
	)

	if err := s.Start(); err != nil {
		log.Fatal(fmt.Sprintf("Unable to start server: %v.", err))
	}
	log.Print("All services are stopped.")
}
