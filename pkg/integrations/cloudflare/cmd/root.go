package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-integration-cloudflare",
		Short: "powerssl-integration-cloudflare provides PowerSSL Cloudflare integration",
		Long: `powerssl-integration-cloudflare provides PowerSSL Cloudflare integration.

Find more information at: https://docs.powerssl.io/powerssl-integration-cloudflare/powerssl-integration-cloudflare.html`,
		Version: "0.1.0",
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/integration-cloudflare/config.yaml)")

	cmd.AddCommand(newCmdRun())

	return cmd
}

func Execute() {
	cobra.OnInitialize(initConfig)

	if err := NewCmdRoot().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/etc/powerssl/integration-cloudflare")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("powerssl")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err == nil && verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
