package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/powerssl/internal/pkg/cmd"
	"powerssl.dev/powerssl/internal/pkg/version"
)

const component = "controller"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	cmdutil.ExecuteWithConfig(NewCmdRoot(), component, cfgFile, verbose)
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-controller",
		Short: "powerssl-controller provides PowerSSL Controller",
		Long: `powerssl-controller provides PowerSSL Controller.

Find more information at: https://docs.powerssl.io/powerssl-controller`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/controller/config.yaml)")

	cmd.AddCommand(newCmdServe())

	return cmd
}
