package client

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber-go/tally"
	temporalclient "go.temporal.io/sdk/client"
	temporalconverter "go.temporal.io/sdk/converter"
	temporallog "go.temporal.io/sdk/log"
	temporalworkflow "go.temporal.io/sdk/workflow"
)

type Client = temporalclient.Client

type Config struct {
	CAFile             string
	DataConverter      temporalconverter.DataConverter
	DisableHealthCheck bool
	HealthCheckTimeout time.Duration
	HostPort           string
	Identity           string
	Namespace          string
	TLSCertFile        string
	TLSKeyFile         string
}

func NewClient(cfg Config, logger temporallog.Logger, tracer opentracing.Tracer) (temporalclient.Client, error) {
	var err error
	var tlsConnectionOptions tls.Config

	if cfg.TLSCertFile != "" && cfg.TLSKeyFile != "" {
		var cert tls.Certificate
		if cert, err = tls.LoadX509KeyPair(cfg.TLSCertFile, cfg.TLSKeyFile); err != nil {
			return nil, err
		}
		tlsConnectionOptions.Certificates = []tls.Certificate{cert}
	}

	if cfg.CAFile != "" {
		var caData []byte
		if caData, err = ioutil.ReadFile(cfg.CAFile); err != nil {
			return nil, err
		}
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(caData)
		tlsConnectionOptions.RootCAs = certPool
	}

	scope, _ := tally.NewRootScope(tally.ScopeOptions{Separator: "_"}, time.Second)

	return temporalclient.NewClient(temporalclient.Options{
		HostPort:           cfg.HostPort,
		Namespace:          cfg.Namespace,
		Logger:             logger,
		MetricsScope:       scope,
		Identity:           cfg.Identity,
		DataConverter:      cfg.DataConverter,
		Tracer:             tracer,
		ContextPropagators: []temporalworkflow.ContextPropagator{},
		ConnectionOptions: temporalclient.ConnectionOptions{
			TLS:                &tlsConnectionOptions,
			DisableHealthCheck: cfg.DisableHealthCheck,
			HealthCheckTimeout: cfg.HealthCheckTimeout,
		},
	})
}
