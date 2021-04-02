package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/common/version"

	"powerssl.dev/tools/dev-runner/internal"
)

func Execute() {
	cmdutil.Execute(NewCmdRoot())
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "powerssl-dev-runner",
		Short:   "powerssl-dev-runner spins up the local dev environment",
		Version: version.String(),
		Args:    cobra.NoArgs,
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return internal.Run()
		}),
	}

	return cmd
}
