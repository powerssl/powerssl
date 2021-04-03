package internal

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"go.temporal.io/api/workflowservice/v1"
	temporalclient "go.temporal.io/sdk/client"
)

func RunRegisterNamespace(address, namespace, description, ownerEmail string, workflowExecutionRetentionPeriod *time.Duration, tlsCertPath, tlsKeyPath, tlsCAPath, tlsServerName string, tlsEnableHostVerification bool) (err error) {
	var tlsConnectionOptions tls.Config

	if tlsCertPath != "" && tlsKeyPath != "" {
		var cert tls.Certificate
		if cert, err = tls.LoadX509KeyPair(tlsCertPath, tlsKeyPath); err != nil {
			return err
		}
		tlsConnectionOptions.Certificates = []tls.Certificate{cert}
	}

	if tlsCAPath != "" {
		var caData []byte
		if caData, err = ioutil.ReadFile(tlsCAPath); err != nil {
			return err
		}
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(caData)
		tlsConnectionOptions.RootCAs = certPool
	}

	if tlsServerName != "" {
		tlsConnectionOptions.ServerName = tlsServerName
	}

	tlsConnectionOptions.InsecureSkipVerify = tlsEnableHostVerification

	var namespaceClient temporalclient.NamespaceClient
	if namespaceClient, err = temporalclient.NewNamespaceClient(temporalclient.Options{
		HostPort: address,
		ConnectionOptions: temporalclient.ConnectionOptions{
			TLS: &tlsConnectionOptions,
		},
	}); err != nil {
		return err
	}

	return namespaceClient.Register(context.Background(), &workflowservice.RegisterNamespaceRequest{
		Namespace: namespace,
		Description: description,
		OwnerEmail: ownerEmail,
		WorkflowExecutionRetentionPeriod: workflowExecutionRetentionPeriod,
	})
}
