package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(newCmdRoot(), "/etc/powerssl/integration-acme", "powerssl")
}

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-integration-acme",
		Short: "powerssl-integration-acme provides PowerSSL ACME integration",
		Long: `powerssl-integration-acme provides PowerSSL ACME integration.

Find more information at: https://docs.powerssl.io/powerssl-integration-acme`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdRun())

	return cmd
}
