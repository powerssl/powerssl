package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"

	temporalclient "go.temporal.io/sdk/client"
	temporalworkflow "go.temporal.io/sdk/workflow"

	"powerssl.dev/common/log"
)

func New(cfg Config, logger log.Logger) (client temporalclient.Client, err error) {
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

	identity := cfg.Identity
	if identity == "" {
		identity = fmt.Sprintf("%d@%s@%s", os.Getpid(), getHostName(), cfg.Component)
	}

	if client, err = temporalclient.NewClient(temporalclient.Options{
		HostPort:           cfg.HostPort,
		Namespace:          cfg.Namespace,
		Logger:             logger.TemporalLogger(),
		Identity:           identity,
		DataConverter:      cfg.DataConverter,
		ContextPropagators: []temporalworkflow.ContextPropagator{},
		ConnectionOptions: temporalclient.ConnectionOptions{
			TLS:                &tlsConnectionOptions,
			DisableHealthCheck: cfg.DisableHealthCheck,
			HealthCheckTimeout: cfg.HealthCheckTimeout,
		},
	}); err != nil {
		return nil, err
	}
	return client, nil
}

func getHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "Unknown"
	}
	return hostName
}
