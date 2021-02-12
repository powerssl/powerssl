package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/integration/acme/internal/version"
)

const component = "integration-acme"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	cmd := newCmdRoot()

	cobra.OnInitialize(initConfig(cmd, component, cfgFile, verbose))

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-integration-acme",
		Short: "powerssl-integration-acme provides PowerSSL ACME integration",
		Long: `powerssl-integration-acme provides PowerSSL ACME integration.

Find more information at: https://docs.powerssl.io/powerssl-integration-acme`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/integration-acme/config.yaml)")

	cmd.AddCommand(newCmdRun())

	return cmd
}

func initConfig(cmd *cobra.Command, component, cfgFile string, verbose bool) func() {
	return func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.AddConfigPath("/etc/powerssl/%s" + component)
			viper.SetConfigName("config")
		}

		viper.AutomaticEnv()
		viper.SetEnvPrefix("powerssl")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

		if err := viper.ReadInConfig(); err == nil && verbose {
			cmd.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

func handleError(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := f(cmd, args); err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
