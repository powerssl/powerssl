package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/powerssl/internal/pkg/version"
)

var (
	cfgFile string
	verbose bool
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-agent",
		Short: "powerssl-agent provides PowerSSL Agent",
		Long: `powerssl-agent provides PowerSSL Agent.

Find more information at: https://docs.powerssl.io/powerssl-agent`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/agent/config.yaml)")

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
		viper.AddConfigPath("/etc/powerssl/agent")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("powerssl")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err == nil && verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
