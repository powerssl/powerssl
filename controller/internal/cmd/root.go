package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/controller", "powerssl")
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-controller",
		Short: "powerssl-controller provides PowerSSL Controller",
		Long: `powerssl-controller provides PowerSSL Controller.

Find more information at: https://docs.powerssl.io/powerssl-controller`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdServe())

	return cmd
}
