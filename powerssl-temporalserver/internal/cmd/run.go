package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.temporal.io/server/temporal"

	"powerssl.dev/temporalserver/internal"
	cmdutil "powerssl.dev/common/cmd"
)

func newCmdRun() *cobra.Command {
	var config internal.Config

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the Temporal Server",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.Unmarshal(&config); err != nil {
				return err
			}
			// NOTE: viper.GetStringSlice doesn't do the trick
			services := strings.Split(viper.GetString("services")[1:len(viper.GetString("services"))-1], ",")
			config.Services = []string{}
			for _, service := range services {
				if service != "" {
					config.Services = append(config.Services, service)
				}
			}
			return config.Validate()
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return internal.Run(&config)
		}),
	}

	cmd.Flags().String("env", "development", "Environment is one of the input params ex-development")
	cmd.Flags().String("config-dir", "config", "Config directory to load a set of yaml config files from")
	cmd.Flags().StringArray("service", temporal.Services, "Service(s) to start")
	cmd.Flags().String("zone", "", "Zone is another input param")

	cmdutil.Must(viper.BindPFlag("env", cmd.Flags().Lookup("env")))
	cmdutil.Must(viper.BindPFlag("config-dir", cmd.Flags().Lookup("config-dir")))
	cmdutil.Must(viper.BindPFlag("services", cmd.Flags().Lookup("service")))
	cmdutil.Must(viper.BindPFlag("zone", cmd.Flags().Lookup("zone")))

	return cmd
}
