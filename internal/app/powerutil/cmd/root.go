package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"powerssl.dev/powerssl/internal/pkg/version"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerutil",
		Short: "powerutil provides PowerSSL utilities",
		Long: `powerutil provides PowerSSL utilities.

Find more information at: https://docs.powerssl.io/powerutil`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdCA())
	cmd.AddCommand(newCmdMigrate())
	cmd.AddCommand(newCmdTemporal())
	cmd.AddCommand(newCmdVault())

	return cmd
}

func Execute() {
	if err := NewCmdRoot().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
