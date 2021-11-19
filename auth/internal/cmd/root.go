package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/auth", "powerssl")
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-auth",
		Short: "powerssl-auth provides PowerSSL Auth",
		Long: `powerssl-auth provides PowerSSL Auth.

Find more information at: https://docs.powerssl.io/powerssl-auth`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdServe())

	return cmd
}
