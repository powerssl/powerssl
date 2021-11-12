package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

const component = "integration-cloudflare"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	snakecharmer.ExecuteWithConfig(newCmdRoot(), "/etc/powerssl/"+component, "powerssl", &cfgFile, &verbose)
}

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-integration-cloudflare",
		Short: "powerssl-integration-cloudflare provides PowerSSL Cloudflare integration",
		Long: `powerssl-integration-cloudflare provides PowerSSL Cloudflare integration.

Find more information at: https://docs.powerssl.io/powerssl-integration-cloudflare`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/integration-cloudflare/config.yaml)")

	cmd.AddCommand(newCmdRun())

	return cmd
}
