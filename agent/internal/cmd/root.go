package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/agent", "powerssl")
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-agent",
		Short: "powerssl-agent provides PowerSSL Agent",
		Long: `powerssl-agent provides PowerSSL Agent.
Find more information at: https://docs.powerssl.io/powerssl-agent`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdRun())

	return cmd
}
