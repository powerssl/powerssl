package cmd

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/powerctl/internal"
)

func newCmdLogin() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login to PowerSSL",
		Args:  cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return viper.Unmarshal(new(internal.Config))
		},
		Run: snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
			location := viper.ConfigFileUsed()
			if err := viper.WriteConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					r := regexp.MustCompile(`\[(.+)]`)
					dir := r.FindStringSubmatch(err.Error())[1]
					if err = os.MkdirAll(dir, 0755); err != nil {
						return err
					}
					location = filepath.Join(dir, "config.yaml")
					err = viper.WriteConfigAs(location)
				}
				if err != nil {
					return err
				}
			}
			cmd.Printf("Wrote config to %s\n", location)
			return nil
		}),
	}
}
