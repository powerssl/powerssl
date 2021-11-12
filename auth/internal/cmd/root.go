package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

const component = "auth"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/"+component, "powerssl", &cfgFile, &verbose)
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-auth",
		Short: "powerssl-auth provides PowerSSL Auth",
		Long: `powerssl-auth provides PowerSSL Auth.

Find more information at: https://docs.powerssl.io/powerssl-auth`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/auth/config.yaml)")

	cmd.AddCommand(newCmdServe())

	return cmd
}
