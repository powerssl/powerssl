package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/common/version"
)

const component = "agent"

var (
	verbose bool
	cfgFile string
)

func Execute() {
	cmdutil.ExecuteWithConfig(NewCmdRoot(), component, &cfgFile, &verbose)
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
