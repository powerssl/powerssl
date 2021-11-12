package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

const component = "worker"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/"+component, "powerssl", &cfgFile, &verbose)
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-worker",
		Short: "powerssl-worker provides PowerSSL Worker",
		Long: `powerssl-worker provides PowerSSL Worker.

Find more information at: https://docs.powerssl.io/powerssl-worker`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/worker/config.yaml)")

	cmd.AddCommand(newCmdRun())

	return cmd
}
