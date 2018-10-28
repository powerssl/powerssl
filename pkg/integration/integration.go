package integration

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/go-kit/kit/log"

	acmetransport "powerssl.io/pkg/controller/acme/transport"
	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	controllerclient "powerssl.io/pkg/controller/client"
	integrationacme "powerssl.io/pkg/integration/acme"
	// integrationdns "powerssl.io/pkg/integration/dns"
)

type kind uint

const (
	KindACME kind = iota
	KindDNS
)

type Integration interface {
	HandleActivity(activity *api.Activity) error
}

type integration struct {
	client  *controllerclient.GRPCClient
	logger  log.Logger
	kind    kind
	name    string
	handler Integration
}

func New(addr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool, kind kind, name string, handler interface{}) *integration {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	client, err := controllerclient.NewGRPCClient(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, logger)
	if err != nil {
		logger.Log("transport", "gRPC", "err", err)
		os.Exit(1)
	}

	var integrationhandler Integration
	switch kind {
	case KindACME:
		integrationhandler = integrationacme.New(client.ACME, handler.(integrationacme.Integration))
	case KindDNS:
		logger.Log("kind", "DNS", "err", "Not yet supported")
		os.Exit(1)
		// integrationhandler = integrationdns.New(client.DNS, handler.(integrationdns.Integration))
	}

	return &integration{
		client:  client,
		logger:  logger,
		kind:    kind,
		name:    name,
		handler: integrationhandler,
	}
}

func (i *integration) Run() {
	for {
		i.logger.Log("err", i.run())
		time.Sleep(time.Second)
	}
}

func (i *integration) run() error {
	var kind apiv1.IntegrationKind
	switch i.kind {
	case KindACME:
		kind = apiv1.IntegrationKind_ACME
	case KindDNS:
		kind = apiv1.IntegrationKind_DNS
	}

	stream, err := i.client.Integration.Register(context.Background(), &apiv1.RegisterIntegrationRequest{
		Kind: kind,
		Name: i.name,
	})
	if err != nil {
		return err
	}
	i.logger.Log("connected", true)
	for {
		activity, err := stream.Recv()
		if err == io.EOF {
			i.logger.Log("err", "EOF")
			break
		}
		if err != nil {
			return err
		}
		{
			activity, err := acmetransport.DecodeGRPCActivity(activity)
			if err != nil {
				return err
			}
			go i.handleActivity(activity)
		}
	}
	return nil
}

func (i *integration) handleActivity(activity *api.Activity) {
	i.logger.Log("activity", activity.Token, "name", activity.Name)
	err := i.handler.HandleActivity(activity)
	i.logger.Log("activity", activity.Token, "err", err)
}
