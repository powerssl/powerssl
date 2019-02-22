package cmd

import (
	"os"

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
	cmd := &cobra.Command{
		Use:   "bash",
		Short: "Generates bash completion scripts",
		Long: `To load completion run

. <(powerctl completion bash)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(powerctl completion bash)
`,
		Run: func(cmd *cobra.Command, args []string) {
			NewCmdRoot().GenBashCompletion(os.Stdout)
		},
	}

	return cmd
}

func newCmdZSHCompletion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zsh",
		Short: "Generates zsh completion scripts",
		Run: func(cmd *cobra.Command, args []string) {
			NewCmdRoot().GenZshCompletion(os.Stdout)
		},
	}

	return cmd
}
