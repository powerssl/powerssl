package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerutil",
		Short: "powerutil provides PowerSSL utilities",
		Long: `powerutil provides PowerSSL utilities.

Find more information at: https://docs.powerssl.io/powerutil`,
		Version: "0.1.0",
	}

	cmd.AddCommand(newCmdCA())
	cmd.AddCommand(newCmdVault())

	return cmd
}

func Execute() {
	if err := NewCmdRoot().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
