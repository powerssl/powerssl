package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/common/version"
)

const component = "apiserver"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	cmdutil.ExecuteWithConfig(NewCmdRoot(), component, &cfgFile, &verbose)
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-apiserver",
		Short: "powerssl-apiserver provides PowerSSL API",
		Long: `powerssl-apiserver provides PowerSSL API.

Find more information at: https://docs.powerssl.io/powerssl-apiserver`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/api/config.yaml)")

	cmd.AddCommand(newCmdMigrate())
	cmd.AddCommand(newCmdServe())

	return cmd
}
