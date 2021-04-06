package client // import "powerssl.dev/backend/temporal/client"

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber-go/tally"
	temporalclient "go.temporal.io/sdk/client"
	temporalconverter "go.temporal.io/sdk/converter"
	temporalworkflow "go.temporal.io/sdk/workflow"

	"powerssl.dev/common/log"
)

func getHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "Unknown"
	}
	return hostName
}

type Client = temporalclient.Client

type Config struct {
	CAFile             string `mapstructure:"ca-file"`
	DataConverter      temporalconverter.DataConverter
	DisableHealthCheck bool
	HealthCheckTimeout time.Duration
	HostPort           string `mapstructure:"host-port" validate:"required,hostname_port"`
	Identity           string
	Namespace          string `validate:"required"`
	TLSCertFile        string
	TLSKeyFile         string
}

func NewClient(cfg Config, logger log.Logger, tracer opentracing.Tracer, component string) (temporalclient.Client, io.Closer, error) {
	var err error
	var tlsConnectionOptions tls.Config

	if cfg.TLSCertFile != "" && cfg.TLSKeyFile != "" {
		var cert tls.Certificate
		if cert, err = tls.LoadX509KeyPair(cfg.TLSCertFile, cfg.TLSKeyFile); err != nil {
			return nil, nil, err
		}
		tlsConnectionOptions.Certificates = []tls.Certificate{cert}
	}

	if cfg.CAFile != "" {
		var caData []byte
		if caData, err = ioutil.ReadFile(cfg.CAFile); err != nil {
			return nil, nil, err
		}
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(caData)
		tlsConnectionOptions.RootCAs = certPool
	}

	identity := cfg.Identity
	if identity == "" {
		identity = fmt.Sprintf("%d@%s@%s", os.Getpid(), getHostName(), component)
	}

	scope, closer := tally.NewRootScope(tally.ScopeOptions{Separator: "_"}, time.Second)

	var client temporalclient.Client
	if client, err = temporalclient.NewClient(temporalclient.Options{
		HostPort:           cfg.HostPort,
		Namespace:          cfg.Namespace,
		Logger:             temporalLogger{logger},
		MetricsScope:       scope,
		Identity:           identity,
		DataConverter:      cfg.DataConverter,
		Tracer:             tracer,
		ContextPropagators: []temporalworkflow.ContextPropagator{},
		ConnectionOptions: temporalclient.ConnectionOptions{
			TLS:                &tlsConnectionOptions,
			DisableHealthCheck: cfg.DisableHealthCheck,
			HealthCheckTimeout: cfg.HealthCheckTimeout,
		},
	}); err != nil {
		return nil, nil, err
	}
	return client, closer, nil
}
