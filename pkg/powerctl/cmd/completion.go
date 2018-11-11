package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Hidden: true,
	Use:    "completion",
	Short:  "Generates completion scripts",
}

var bashCompletionCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generates bash completion scripts",
	Long: `To load completion run

. <(powerctl completion bash)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(powerctl completion bash)
`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

var zshCompletionCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generates zsh completion scripts",
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(bashCompletionCmd)
	completionCmd.AddCommand(zshCompletionCmd)

	rootCmd.AddCommand(completionCmd)
}
