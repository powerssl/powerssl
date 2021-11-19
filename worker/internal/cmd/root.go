package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/worker", "powerssl")
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-worker",
		Short: "powerssl-worker provides PowerSSL Worker",
		Long: `powerssl-worker provides PowerSSL Worker.

Find more information at: https://docs.powerssl.io/powerssl-worker`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdRun())

	return cmd
}
