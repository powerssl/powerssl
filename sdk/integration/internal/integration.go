package internal

import (
	"context"
	"io"

	"go.uber.org/zap"

	apiv1 "powerssl.dev/api/controller/v1"
	"powerssl.dev/sdk/controller"
	"powerssl.dev/sdk/integration/acme"
	"powerssl.dev/sdk/integration/dns"
	"powerssl.dev/sdk/integration/middleware"
	// integrationdns "powerssl.dev/sdk/integration/dns"
)

type Integration interface {
	HandleActivity(ctx context.Context, activity *apiv1.Activity) error
}

type IntegrationName string

type IntegrationConfig struct {
	Kind apiv1.IntegrationKind
	Name string
}

type integration struct {
	cfg     *IntegrationConfig
	client  *controller.GRPCClient
	handler Integration
	logger  *zap.SugaredLogger
}

func New(cfg *IntegrationConfig, logger *zap.SugaredLogger, client *controller.GRPCClient, handler Integration) *integration {
	return &integration{
		cfg:     cfg,
		client:  client,
		handler: handler,
		logger:  logger,
	}
}

func NewACME(name string, logger *zap.SugaredLogger, client *controller.GRPCClient, handler acme.Integration) *integration {
	cfg := &IntegrationConfig{
		Kind: apiv1.IntegrationKind_ACME,
		Name: name,
	}
	acmeHandler := acme.New(client.ACME, handler)
	return New(cfg, logger, client, acmeHandler)
}

func NewDNS(name string, logger *zap.SugaredLogger, client *controller.GRPCClient, handler dns.Integration) *integration {
	cfg := &IntegrationConfig{
		Kind: apiv1.IntegrationKind_DNS,
		Name: name,
	}
	dnsHandler := dns.New(client, handler)
	return New(cfg, logger, client, dnsHandler)
}

func (i *integration) Run(ctx context.Context) error {
	stream, err := i.register(ctx)
	if err != nil {
		return err
	}
	handler, err := i.makeHandler(ctx)
	if err != nil {
		return err
	}
	for {
		if err = i.receive(stream, handler); err != nil {
			i.logger.Error(err)
		}
	}
}

func (i *integration) makeHandler(ctx context.Context) (func(*apiv1.Activity), error) {
	handler := middleware.ActivityHandler(func(activity *apiv1.Activity) {
		if err := i.handler.HandleActivity(ctx, activity); err != nil {
			i.logger.Error(err)
		}
	})
	handler = middleware.LoggingMiddleware(i.logger)(handler)
	handler = middleware.TracingMiddleware()(handler)
	return handler.Handle, nil
}

func (i *integration) receive(stream apiv1.IntegrationService_RegisterClient, handle func(*apiv1.Activity)) error {
	for {
		activity, err := stream.Recv()
		if err == io.EOF {
			i.logger.Info("EOF")
			return nil
		}
		if err != nil {
			return err
		}
		go handle(activity)
	}
}

func (i *integration) register(ctx context.Context) (apiv1.IntegrationService_RegisterClient, error) {
	stream, err := i.client.Integration.Register(ctx, &apiv1.RegisterIntegrationRequest{
		Kind: i.cfg.Kind,
		Name: string(i.cfg.Name),
	})
	if err != nil {
		return nil, err
	}
	i.logger.Info("connected")
	return stream, nil
}
