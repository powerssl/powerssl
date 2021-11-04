package cmd // import "powerssl.dev/common/cmd"

import (
	"context"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/common/runner"
)

func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func ExecuteWithConfig(cmd *cobra.Command, component string, cfgFile *string, verbose *bool) {
	cobra.OnInitialize(initConfig(cmd, component, cfgFile, verbose))

	Execute(cmd)
}

func HandleError(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := f(cmd, args); err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
	}
}

func Run(f func(ctx context.Context) ([]func() error, func(), error)) func(cmd *cobra.Command, args []string) {
	return HandleError(func(cmd *cobra.Command, args []string) error {
		return runner.Run(f)
	})
}

func initConfig(cmd *cobra.Command, component string, cfgFile *string, verbose *bool) func() {
	return func() {
		if *cfgFile != "" {
			viper.SetConfigFile(*cfgFile)
		} else {
			viper.AddConfigPath("/etc/powerssl/" + component)
			viper.SetConfigName("config")
		}

		viper.AutomaticEnv()
		viper.SetEnvPrefix("powerssl")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

		if err := viper.ReadInConfig(); err == nil && *verbose {
			cmd.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
