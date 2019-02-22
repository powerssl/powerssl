package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to PowerSSL",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 && args[0] != "" {
				viper.Set("addr", args[0])
			}

			if viper.GetString("addr") == "" {
				return fmt.Errorf("Address must be set")
			}

			if viper.GetString("auth-token") == "" {
				return fmt.Errorf("Auth token must be set")
			}

			if err := viper.WriteConfig(); err != nil {
				return err
			}

			fmt.Printf("Wrote config to %s\n", viper.ConfigFileUsed())

			return nil
		},
	}

	return cmd
}
