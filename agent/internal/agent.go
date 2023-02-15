package internal

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/common"
	"powerssl.dev/common/log"
	"powerssl.dev/common/tracing"
	"powerssl.dev/sdk/apiserver"
)

const component = "powerssl-agent"

func Run(cfg *Config) (err error) {
	var logger log.Logger
	if logger, err = log.NewLogger(false); err != nil {
		return err
	}
	defer common.ErrWrapSync(logger, &err)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return common.InterruptHandler(ctx, logger)
	})

	var tracer opentracing.Tracer
	{
		var closer io.Closer
		if tracer, closer, err = tracing.Init(component, "", logger.With("component", "tracing")); err != nil {
			return err
		}
		defer common.ErrWrapCloser(closer, &err)
	}

	var apiserverClient *apiserver.Client
	{
		if apiserverClient, err = apiserver.NewClient(ctx, &cfg.APIServerClientConfig, cfg.AuthToken, logger, tracer); err != nil {
			return err
		}
	}
	var _ = apiserverClient

	g.Go(func() error {
		return nil
	})

	if err = g.Wait(); err != nil {
		switch err.(type) {
		case common.InterruptError:
		default:
			return err
		}
	}
	return nil
}