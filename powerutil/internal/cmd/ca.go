package cmd

import (
	"github.com/spf13/cobra"
)

func newCmdCA() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ca",
		Short: "CA",
	}

	cmd.AddCommand(newCmdCAGen())
	cmd.AddCommand(newCmdCAInit())
	cmd.AddCommand(newCmdCASign())

	return cmd
}
