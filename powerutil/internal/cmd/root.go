package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

var verbose bool

func Execute() {
	snakecharmer.Execute(NewCmdRoot())
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerutil",
		Short: "powerutil provides PowerSSL utilities",
		Long: `powerutil provides PowerSSL utilities.

Find more information at: https://docs.powerssl.io/powerutil`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	cmd.AddCommand(newCmdCA())
	cmd.AddCommand(newCmdVault())

	return cmd
}
