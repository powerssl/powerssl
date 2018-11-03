package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	acmetransport "powerssl.io/pkg/controller/acme/transport" // TODO: Wrong package
	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	controllerclient "powerssl.io/pkg/controller/client"
	integrationacme "powerssl.io/pkg/integration/acme"
	// integrationdns "powerssl.io/pkg/integration/dns"
	"powerssl.io/pkg/util/logging"
	"powerssl.io/pkg/util/tracing"
)

type kind string

const (
	KindACME kind = "acme"
	KindDNS  kind = "dns"
)

type Integration interface {
	HandleActivity(ctx context.Context, activity *api.Activity) error
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
		logger = logging.NewLogger()
	}

	var tracer stdopentracing.Tracer
	{
		if true { // TODO
			var err error
			tracer, _, err = tracing.NewJaegerTracer(fmt.Sprintf("powerssl-integration-%s", kind), logger)
			if err != nil {
				logger.Log("tracing", "jaeger", "during", "initialize", "err", err)
			}
		} else {
			tracer = stdopentracing.GlobalTracer()
		}
	}

	client, err := controllerclient.NewGRPCClient(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, logger, tracer)
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
	ctx := context.TODO()
	for {
		i.logger.Log("err", i.run(ctx))
		time.Sleep(time.Second)
	}
}

func (i *integration) run(ctx context.Context) error {
	var kind apiv1.IntegrationKind
	switch i.kind {
	case KindACME:
		kind = apiv1.IntegrationKind_ACME
	case KindDNS:
		kind = apiv1.IntegrationKind_DNS
	}

	stream, err := i.client.Integration.Register(ctx, &apiv1.RegisterIntegrationRequest{
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
	wireContext, err := extractWireContext(activity.Signature) // TODO: Do not use Signature for Span
	if err != nil {
		// TODO
	}
	serverSpan := opentracing.StartSpan(string(activity.Name), ext.RPCServerOption(wireContext))
	serverSpan.SetTag("token", activity.Token)
	ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)
	err = i.handler.HandleActivity(ctx, activity)
	serverSpan.Finish()
	i.logger.Log("activity", activity.Token, "err", err)
}

func extractWireContext(s string) (opentracing.SpanContext, error) {
	var tmc opentracing.TextMapCarrier
	if err := json.Unmarshal([]byte(s), &tmc); err != nil {
		return nil, err
	}
	wireContext, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, tmc)
	if err != nil {
		return nil, err
	}
	return wireContext, nil
}
