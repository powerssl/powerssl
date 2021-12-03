package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"
)

func newCmdCompletion() *cobra.Command {
	cmd := &cobra.Command{
		Hidden: true,
		Use:    "completion",
		Short:  "Generates completion scripts",
	}

	cmd.AddCommand(newCmdBashCompletion())
	cmd.AddCommand(newCmdZSHCompletion())

	return cmd
}

func newCmdBashCompletion() *cobra.Command {
	return &cobra.Command{
		Use:   "bash",
		Short: "Generates bash completion scripts",
		Long: `To load completion run

. <(powerctl completion bash)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(powerctl completion bash)
`,
		Args: cobra.NoArgs,
		Run: snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
			return NewCmdRoot().GenBashCompletion(cmd.OutOrStdout())
		}),
	}
}

func newCmdZSHCompletion() *cobra.Command {
	return &cobra.Command{
		Use:   "zsh",
		Short: "Generates zsh completion scripts",
		Args:  cobra.NoArgs,
		Run: snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
			return NewCmdRoot().GenZshCompletion(cmd.OutOrStdout())
		}),
	}
}
