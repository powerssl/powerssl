package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/common/version"
)

const component = "temporal"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	cmdutil.ExecuteWithConfig(NewCmdRoot(), component, cfgFile, verbose)
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-temporal",
		Short: "powerssl-temporal provides PowerSSL Temporal Server",
		Long: `powerssl-temporal provides PowerSSL Temporal Server.

Find more information at: https://docs.powerssl.io/powerssl-temporal`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/temporal/config.yaml)")

	cmd.AddCommand(newCmdMigrate())
	cmd.AddCommand(newCmdRegisterNamespace())
	cmd.AddCommand(newCmdRun())

	return cmd
}
