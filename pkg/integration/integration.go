package integration

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-kit/kit/log"

	acmetransport "powerssl.io/pkg/controller/acme/transport"
	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	controllerclient "powerssl.io/pkg/controller/client"
)

type kind uint

const (
	KindACME kind = iota
	KindDNS
)

type Integration interface {
	HandleActivity(activity *apiv1.Activity) error
}

type integration struct {
	client *controllerclient.GRPCClient
	logger log.Logger
	kind   kind
	name   string
}

func New(addr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool, kind kind, name string) *integration {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	client, err := controllerclient.NewGRPCClient(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, logger)
	if err != nil {
		panic(err)
	}

	return &integration{
		client: client,
		logger: logger,
		kind:   kind,
		name:   name,
	}
}

func (i *integration) Run(handler interface{}) {
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
				panic(err)
			}
			go i.handleActivity(activity)
		}
	}
	return nil
}

func (i *integration) handleActivity(activity *api.Activity) {
	i.logger.Log("activity", activity.Token, "name", activity.Name)
	var err error
	switch activity.Name {
	// case api.Activity_CA_AUTHORIZE_DOMAIN:
	// 	err = i.caAuthorizeDomain(activity)
	// case api.Activity_CA_REQUEST_CERTIFICATE:
	// 	err = i.caRequestCertificate(activity)
	// case api.Activity_CA_REVOKE_CERTIFICATE:
	// 	err = i.caRevokeCertificate(activity)
	// case api.Activity_CA_VERIFY_DOMAIN:
	// 	err = i.caVerifyDomain(activity)
	// case api.Activity_DNS_CREATE_RECORD:
	// 	err = i.dnsCreateRecord(activity)
	// case api.Activity_DNS_DELETE_RECORD:
	// 	err = i.dnsDeleteRecord(activity)
	// case api.Activity_DNS_VERIFY_DOMAIN:
	// 	err = i.dnsVerifyDomain(activity)
	default:
		err = fmt.Errorf("Activity %s not implemented", activity.Name)
	}
	i.logger.Log("activity", activity.Token, "err", err)
}
