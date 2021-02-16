package cmd

import (
	"github.com/spf13/cobra"

	"powerssl.dev/powerutil/internal"
	cmdutil "powerssl.dev/common/cmd"
)

func newCmdTemporal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "temporal",
		Short: "Temporal",
	}

	cmd.AddCommand(newCmdTemporalMigrate())
	cmd.AddCommand(newCmdTemporalRegisterNamespace())

	return cmd
}

func newCmdTemporalMigrate() *cobra.Command {
	var docker bool
	var host, password, plugin, port, temporalDatabase, user, visibilityDatabase string

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run temporal migrations",
		Args:  cobra.NoArgs,
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return internal.RunTemporalMigrate(docker, host, password, plugin, port, temporalDatabase, user, visibilityDatabase)
		}),
	}

	cmd.Flags().BoolVar(&docker, "docker", false, "Execute with docker")
	cmd.Flags().StringVar(&host, "host", "", "DB host")
	cmd.Flags().StringVar(&password, "password", "", "DB Password")
	cmd.Flags().StringVar(&plugin, "plugin", "", "DB Plugin")
	cmd.Flags().StringVar(&port, "port", "", "DB Port")
	cmd.Flags().StringVar(&temporalDatabase, "temporal-database", "temporal", "Temporal DB")
	cmd.Flags().StringVar(&user, "user", "", "DB User")
	cmd.Flags().StringVar(&visibilityDatabase, "visibility-database", "temporal_visibility", "Visibility DB")

	return cmd
}

func newCmdTemporalRegisterNamespace() *cobra.Command {
	var docker, tlsEnableHostVerification bool
	var address, namespace, tlsCertPath, tlsKeyPath, tlsCAPath, tlsServerName string

	cmd := &cobra.Command{
		Use:   "register-namespace",
		Short: "Run temporal register namespace",
		Args:  cobra.NoArgs,
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return internal.RunTemporalRegisterNamespace(docker, tlsEnableHostVerification, address, namespace, tlsCertPath, tlsKeyPath, tlsCAPath, tlsServerName)
		}),
	}

	cmd.Flags().BoolVar(&docker, "docker", false, "execute with docker")
	cmd.Flags().StringVar(&address, "address", "127.0.0.1:7233", "host:port for Temporal frontend service")
	cmd.Flags().StringVar(&namespace, "namespace", "powerssl", "Temporal workflow namespace")
	cmd.Flags().StringVar(&tlsCertPath, "tls-cert-path", "", "path to x509 certificate")
	cmd.Flags().StringVar(&tlsKeyPath, "tls-key-path", "", "path to private key")
	cmd.Flags().StringVar(&tlsCAPath, "tls-ca-path", "", "path to server CA certificate")
	cmd.Flags().BoolVar(&tlsEnableHostVerification, "tls-enable-host-verification", false, "validates hostname of temporal cluster against server certificate")
	cmd.Flags().StringVar(&tlsServerName, "tls-server-name", "", "override for target server name")

	return cmd
}
