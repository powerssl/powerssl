package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "powerssl.dev/common/cmd"
)

func newCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to PowerSSL",
		Args:  cobra.MaximumNArgs(1),
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 && args[0] != "" {
				viper.Set("addr", args[0])
			}

			if viper.GetString("addr") == "" {
				return errors.New("address must be set")
			}

			if viper.GetString("auth-token") == "" {
				return errors.New("auth token must be set")
			}

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

	return cmd
}