package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(newCmdRoot(), "/etc/powerssl/integration-cloudflare", "powerssl")
}

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-integration-cloudflare",
		Short: "powerssl-integration-cloudflare provides PowerSSL Cloudflare integration",
		Long: `powerssl-integration-cloudflare provides PowerSSL Cloudflare integration.

Find more information at: https://docs.powerssl.io/powerssl-integration-cloudflare`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdRun())

	return cmd
}
