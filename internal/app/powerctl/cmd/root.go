package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/powerssl/internal/app/powerctl/resource"
	"powerssl.io/powerssl/internal/pkg/version"
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

	cmd.PersistentFlags().BoolP("insecure", "", false, "Use insecure communication")
	cmd.PersistentFlags().BoolP("insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.PersistentFlags().StringP("addr", "", "", "GRPC address of API server")
	cmd.PersistentFlags().StringP("auth-token", "", "", "Authentication token")
	cmd.PersistentFlags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.PersistentFlags().StringP("output", "o", "table", "Output format")
	cmd.PersistentFlags().StringP("server-name-override", "", "", "It will override the virtual host name of authority")

	viper.BindPFlag("addr", cmd.PersistentFlags().Lookup("addr"))
	viper.BindPFlag("auth-token", cmd.PersistentFlags().Lookup("auth-token"))
	viper.BindPFlag("ca-file", cmd.PersistentFlags().Lookup("ca-file"))
	viper.BindPFlag("insecure", cmd.PersistentFlags().Lookup("insecure"))
	viper.BindPFlag("insecure-skip-tls-verify", cmd.PersistentFlags().Lookup("insecure-skip-tls-verify"))
	viper.BindPFlag("output", cmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("server-name-override", cmd.PersistentFlags().Lookup("server-name-override"))

	cmd.AddCommand(newCmdCompletion())
	cmdCreate := newCmdCreate()
	cmdCreate.AddCommand(resource.NewCmdCreateACMEAccount())
	cmdCreate.AddCommand(resource.NewCmdCreateACMEServer())
	cmdCreate.AddCommand(resource.NewCmdCreateCertificate())
	cmdCreate.AddCommand(resource.NewCmdCreateUser())
	cmd.AddCommand(cmdCreate)
	cmd.AddCommand(newCmdDelete())
	cmd.AddCommand(newCmdDescribe())
	cmd.AddCommand(newCmdGet())
	cmd.AddCommand(newCmdLogin())
	cmd.AddCommand(newCmdUpdate())

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
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(filepath.Join(home, ".powerctl"))
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("powerssl")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil && strings.Contains("Not Found", err.Error()) {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	} else if err == nil && verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
