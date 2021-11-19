package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/temporal", "powerssl")
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-temporal",
		Short: "powerssl-temporal provides PowerSSL Temporal Server",
		Long: `powerssl-temporal provides PowerSSL Temporal Server.

Find more information at: https://docs.powerssl.io/powerssl-temporal`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdMigrate())
	cmd.AddCommand(newCmdRegisterNamespace())
	cmd.AddCommand(newCmdRun())

	return cmd
}
