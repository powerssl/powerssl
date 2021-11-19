package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/apiserver", "powerssl")
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-apiserver",
		Short: "powerssl-apiserver provides PowerSSL API",
		Long: `powerssl-apiserver provides PowerSSL API.

Find more information at: https://docs.powerssl.io/powerssl-apiserver`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdMigrate())
	cmd.AddCommand(newCmdServe())

	return cmd
}
