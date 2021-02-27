package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/temporalserver/internal"
)

func newCmdRegisterNamespace() *cobra.Command {
	var tlsEnableHostVerification bool
	var address, namespace, tlsCertPath, tlsKeyPath, tlsCAPath, tlsServerName string

	cmd := &cobra.Command{
		Use:   "register-namespace",
		Short: "Run temporal register namespace",
		Args:  cobra.NoArgs,
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return internal.RunRegisterNamespace(address, namespace, tlsCertPath, tlsKeyPath, tlsCAPath, tlsServerName, tlsEnableHostVerification)
		}),
	}

	cmd.Flags().StringVar(&address, "address", "127.0.0.1:7233", "host:port for Temporal frontend service")
	cmd.Flags().StringVar(&namespace, "namespace", "powerssl", "Temporal workflow namespace")
	cmd.Flags().StringVar(&tlsCertPath, "tls-cert-path", "", "path to x509 certificate")
	cmd.Flags().StringVar(&tlsKeyPath, "tls-key-path", "", "path to private key")
	cmd.Flags().StringVar(&tlsCAPath, "tls-ca-path", "", "path to server CA certificate")
	cmd.Flags().BoolVar(&tlsEnableHostVerification, "tls-enable-host-verification", false, "validates hostname of temporal cluster against server certificate")
	cmd.Flags().StringVar(&tlsServerName, "tls-server-name", "", "override for target server name")

	return cmd
}
