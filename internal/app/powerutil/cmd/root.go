package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/powerssl/internal/pkg/cmd"
	"powerssl.dev/powerssl/internal/pkg/version"
)

var verbose bool

func Execute() {
	cmdutil.Execute(NewCmdRoot())
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
	cmd.AddCommand(newCmdMigrate())
	cmd.AddCommand(newCmdTemporal())
	cmd.AddCommand(newCmdVault())

	return cmd
}
