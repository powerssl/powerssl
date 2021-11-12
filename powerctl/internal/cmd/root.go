package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/common/version"
)

var (
	verbose bool
	cfgFile string
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerctl",
		Short: "powerctl controls PowerSSL",
		Long: `powerctl controls PowerSSL.

Find more information at: https://docs.powerssl.io/powerctl`,
		Version: version.String(),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			switch viper.GetString("output") {
			case "json", "table", "yaml":
			default:
				return fmt.Errorf("unknown output format")
			}
			return nil
		},
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.powerctl/config.yaml)")

	cmd.PersistentFlags().Bool("insecure", false, "Use insecure communication")
	cmd.PersistentFlags().Bool("insecure-skip-tls-verify", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.PersistentFlags().String("addr", "", "GRPC address of API server")
	cmd.PersistentFlags().String("auth-token", "", "Authentication token")
	cmd.PersistentFlags().String("ca-file", "", "Certificate authority file")
	cmd.PersistentFlags().String("server-name-override", "", "It will override the virtual host name of authority")
	cmd.PersistentFlags().StringP("output", "o", "table", "Output format")

	must(viper.BindPFlag("addr", cmd.PersistentFlags().Lookup("addr")))
	must(viper.BindPFlag("auth-token", cmd.PersistentFlags().Lookup("auth-token")))
	must(viper.BindPFlag("ca-file", cmd.PersistentFlags().Lookup("ca-file")))
	must(viper.BindPFlag("insecure", cmd.PersistentFlags().Lookup("insecure")))
	must(viper.BindPFlag("insecure-skip-tls-verify", cmd.PersistentFlags().Lookup("insecure-skip-tls-verify")))
	must(viper.BindPFlag("output", cmd.PersistentFlags().Lookup("output")))
	must(viper.BindPFlag("server-name-override", cmd.PersistentFlags().Lookup("server-name-override")))

	cmd.AddCommand(newCmdCompletion())
	cmd.AddCommand(newCmdCreate())
	cmd.AddCommand(newCmdDelete())
	cmd.AddCommand(newCmdDescribe())
	cmd.AddCommand(newCmdGet())
	cmd.AddCommand(newCmdLogin())
	cmd.AddCommand(newCmdUpdate())

	return cmd
}

func Execute() {
	cobra.OnInitialize(initConfig)

	snakecharmer.Execute(NewCmdRoot())
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		viper.AddConfigPath(filepath.Join(home, ".powerctl"))
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("powerssl")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil && strings.Contains("Not Found", err.Error()) {
		log.Fatal("Can't read config:", err)
	} else if err == nil && verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
