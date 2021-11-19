package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/webapp", "powerssl")
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-webapp",
		Short: "powerssl-webapp provides PowerSSL WebApp",
		Long: `powerssl-webapp provides PowerSSL WebApp.

Find more information at: https://docs.powerssl.io/powerssl-webapp`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdServe())

	return cmd
}
