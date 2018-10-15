package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Verbose bool
	cfgFile string
	Output  string
)

var rootCmd = &cobra.Command{
	Use:   "powerctl",
	Short: "powerctl controls PowerSSL",
	Long: `powerctl controls PowerSSL.

Find more information at: https://powerssl.io`,
	Version: "0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.powerctl.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().StringVarP(&Output, "output", "o", "yaml", "Output format")

	rootCmd.PersistentFlags().StringP("ca-file", "", "", "Certificate authority file")
	rootCmd.PersistentFlags().StringP("grpc-addr", "", "", "GRPC address of API")
	rootCmd.PersistentFlags().BoolP("insecure", "", false, "Use insecure communication")
	rootCmd.PersistentFlags().BoolP("insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	rootCmd.PersistentFlags().StringP("server-name-override", "", "", "It will override the virtual host name of authority")

	viper.BindPFlag("ca-file", rootCmd.PersistentFlags().Lookup("ca-file"))
	viper.BindPFlag("grpc-addr", rootCmd.PersistentFlags().Lookup("grpc-addr"))
	viper.BindPFlag("insecure", rootCmd.PersistentFlags().Lookup("insecure"))
	viper.BindPFlag("insecure-skip-tls-verify", rootCmd.PersistentFlags().Lookup("insecure-skip-tls-verify"))
	viper.BindPFlag("server-name-override", rootCmd.PersistentFlags().Lookup("server-name-override"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".powerctl")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil && strings.Contains("Not Found", err.Error()) {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	} else if err == nil && Verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
