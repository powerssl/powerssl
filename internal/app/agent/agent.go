package agent

import (
	"context"
	"os"

	"golang.org/x/sync/errgroup"

	"powerssl.io/powerssl/internal/pkg/tracing"
	"powerssl.io/powerssl/internal/pkg/util"
	apiserverclient "powerssl.io/powerssl/pkg/apiserver/client"
)

func Run(cfg *Config) {
	logger := util.NewLogger(os.Stdout)

	util.ValidateConfig(cfg, logger)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	g.Go(func() error {
		tracer, _, _ := tracing.NewNoopTracer("powerssl-agent", logger)
		client, err := apiserverclient.NewGRPCClient(ctx, cfg.APIServerClientConfig, cfg.AuthToken, logger, tracer)
		if err != nil {
			logger.Log("transport", "gRPC", "err", err)
			os.Exit(1)
		}
		var _ = client
		return nil
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}
