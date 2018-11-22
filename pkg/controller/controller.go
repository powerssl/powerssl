package controller

import (
	"context"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"

	apiserverclient "powerssl.io/pkg/apiserver/client"
	workflowengine "powerssl.io/pkg/controller/workflow/engine"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/auth"
	"powerssl.io/pkg/util/tracing"
)

func Run(grpcAddr, grpcCertFile, grpcKeyFile string, grpcInsecure bool, metricsAddr, tracerImpl, apiserverAddr, apiserverCertFile, apiserverServerNameOverride string, apiserverInsecure, apiserverInsecureSkipTLSVerify bool, jwksURL, apiserverAuthToken string) {
	logger := util.NewLogger(os.Stdout)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	tracer, closer, err := tracing.Init("powerssl-controller", tracerImpl, log.With(logger, "component", "tracing"))
	if err != nil {
		logger.Log("component", "tracing", "err", err)
		os.Exit(1)
	}
	defer closer.Close()

	duration := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "powerssl_io",
		Subsystem: "controller",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds.",
	}, []string{"method", "success"})

	var client *apiserverclient.GRPCClient
	{
		token, err := auth.NewServiceToken(apiserverAuthToken)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		if client, err = apiserverclient.NewGRPCClient(apiserverAddr, apiserverCertFile, apiserverServerNameOverride, apiserverInsecure, apiserverInsecureSkipTLSVerify, token, logger, tracer); err != nil {
			logger.Log("transport", "gRPC", "during", "Connect", "err", err)
			os.Exit(1)
		}
	}

	engine := workflowengine.New()

	services, err := makeServices(logger, tracer, duration, client, jwksURL)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	g.Go(func() error {
		return engine.Run(ctx)
	})

	if metricsAddr != "" {
		g.Go(func() error {
			return util.ServeMetrics(ctx, metricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return util.ServeGRPC(ctx, grpcAddr, grpcCertFile, grpcKeyFile, grpcInsecure, log.With(logger, "transport", "gRPC"), services)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}
