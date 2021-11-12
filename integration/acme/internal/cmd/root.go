package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

const component = "integration-acme"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	snakecharmer.ExecuteWithConfig(newCmdRoot(), "/etc/powerssl/"+component, "powerssl", &cfgFile, &verbose)
}

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-integration-acme",
		Short: "powerssl-integration-acme provides PowerSSL ACME integration",
		Long: `powerssl-integration-acme provides PowerSSL ACME integration.

Find more information at: https://docs.powerssl.io/powerssl-integration-acme`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/integration-acme/config.yaml)")

	cmd.AddCommand(newCmdRun())

	return cmd
}
