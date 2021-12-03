package cmd

import (
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
	"powerssl.dev/powerctl/internal"
)

func Execute() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	snakecharmer.ExecuteWithConfig(NewCmdRoot(), filepath.Join(home, ".powerctl"), "powerssl")
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerctl",
		Short: "powerctl controls PowerSSL",
		Long: `powerctl controls PowerSSL.

Find more information at: https://docs.powerssl.io/powerctl`,
		Version: version.String(),
	}

	snakecharmer.GenerateFlags(cmd.PersistentFlags(), new(internal.Config))

	cmd.AddCommand(newCmdCompletion())
	cmd.AddCommand(newCmdCreate())
	cmd.AddCommand(newCmdDelete())
	cmd.AddCommand(newCmdDescribe())
	cmd.AddCommand(newCmdGet())
	cmd.AddCommand(newCmdLogin())
	cmd.AddCommand(newCmdUpdate())

	return cmd
}
