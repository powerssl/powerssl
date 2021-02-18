package internal

import (
	"context"
	"io"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/common"
	"powerssl.dev/common/tracing"
	apiserverclient "powerssl.dev/sdk/apiserver/client"
)

const component = "powerssl-agent"

func Run(cfg *Config) (err error) {
	_, logger := common.NewZapAndKitLogger()

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return common.InterruptHandler(ctx, logger)
	})

	var tracer opentracing.Tracer
	{
		var closer io.Closer
		if tracer, closer, err = tracing.Init(component, "", log.With(logger, "component", "tracing")); err != nil {
			return err
		}
		defer common.ErrWrapCloser(closer, &err)
	}

	var apiserverClient *apiserverclient.GRPCClient
	{
		if apiserverClient, err = apiserverclient.NewGRPCClient(ctx, &cfg.APIServerClientConfig, cfg.AuthToken, logger, tracer); err != nil {
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
