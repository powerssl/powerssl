package cmd

import (
	"github.com/spf13/cobra"
	cmdutil "powerssl.dev/powerssl/internal/pkg/cmd"
	"powerssl.dev/powerssl/internal/pkg/version"
)

const component = "agent"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	cmdutil.ExecuteWithConfig(NewCmdRoot(), component, cfgFile, verbose)
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-agent",
		Short: "powerssl-agent provides PowerSSL Agent",
		Long: `powerssl-agent provides PowerSSL Agent.

Find more information at: https://docs.powerssl.io/powerssl-agent`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/agent/config.yaml)")

	cmd.AddCommand(newCmdRun())

	return cmd
}
