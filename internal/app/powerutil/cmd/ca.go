package cmd

import (
	"github.com/cloudflare/cfssl/log"
	"github.com/spf13/cobra"
)

func init() {
	log.Level = log.LevelError
}

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
