package integration

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-kit/kit/log"

	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	catransport "powerssl.io/pkg/controller/ca/transport"
	controllerclient "powerssl.io/pkg/controller/client"
)

var errorNotImplemented = errors.New("Interface does not implement CAIntegration")

type Integration interface {
	HandleActivity(activity *apiv1.Activity) error
}

type integration struct {
	client  *controllerclient.GRPCClient
	logger  log.Logger
	handler interface{}
}

func New(addr, certFile, serverNameOverride string, insecure, insecureSkipTLSVerify bool, handler interface{}) *integration {
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
		client:  client,
		logger:  logger,
		handler: handler,
	}
}

func (i *integration) Run() {
	for {
		i.logger.Log("err", i.run())
		time.Sleep(time.Second)
	}
}

func (i *integration) run() error {
	var (
		kind apiv1.IntegrationKind
		name string
	)
	switch i.handler.(type) {
	case CAIntegration:
		kind = apiv1.IntegrationKind_CA
		name = i.handler.(CAIntegration).GetName()
	case DNSIntegration:
		kind = apiv1.IntegrationKind_DNS
		name = i.handler.(DNSIntegration).GetName()
	}

	stream, err := i.client.Integration.Register(context.Background(), &apiv1.RegisterIntegrationRequest{
		Kind: kind,
		Name: name,
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
			activity, err := catransport.DecodeGRPCActivity(activity)
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
	case api.Activity_CA_AUTHORIZE_DOMAIN:
		err = i.caAuthorizeDomain(activity)
	case api.Activity_CA_REQUEST_CERTIFICATE:
		err = i.caRequestCertificate(activity)
	case api.Activity_CA_REVOKE_CERTIFICATE:
		err = i.caRevokeCertificate(activity)
	case api.Activity_CA_VERIFY_DOMAIN:
		err = i.caVerifyDomain(activity)
	case api.Activity_DNS_CREATE_RECORD:
		err = i.dnsCreateRecord(activity)
	case api.Activity_DNS_DELETE_RECORD:
		err = i.dnsDeleteRecord(activity)
	case api.Activity_DNS_VERIFY_DOMAIN:
		err = i.dnsVerifyDomain(activity)
	default:
		err = fmt.Errorf("Activity %s not implemented", activity.Name)
	}
	i.logger.Log("activity", activity.Token, "err", err)
}
