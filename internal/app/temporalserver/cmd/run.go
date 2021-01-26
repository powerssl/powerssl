package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.temporal.io/server/temporal"

	"powerssl.dev/powerssl/internal/app/temporalserver"
)

func newCmdRun() *cobra.Command {
	var (
		env       string
		configDir string
		services  []string
		zone      string
	)

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the Temporal Server",
		PreRun: func(cmd *cobra.Command, args []string) {
			env = viper.GetString("env")
			configDir = viper.GetString("config-dir")
			// NOTE: viper.GetStringSlice doesn't do the trick
			services = strings.Split(viper.GetString("service")[1:len(viper.GetString("service"))-1], ",")
			zone = viper.GetString("zone")
		},
		Run: func(cmd *cobra.Command, args []string) {
			temporalserver.Run(&temporalserver.Config{
				Env:       env,
				ConfigDir: configDir,
				Services:  services,
				Zone:      zone,
			})
		},
	}

	cmd.Flags().StringP("env", "", "development", "Environment is one of the input params ex-development")
	cmd.Flags().StringP("config-dir", "", "config", "Config directory to load a set of yaml config files from")
	cmd.Flags().StringArrayP("service", "", temporal.Services, "Service(s) to start")
	cmd.Flags().StringP("zone", "", "", "Zone is another input param")

	viper.BindPFlag("env", cmd.Flags().Lookup("env"))
	viper.BindPFlag("config-dir", cmd.Flags().Lookup("config-dir"))
	viper.BindPFlag("service", cmd.Flags().Lookup("service"))
	viper.BindPFlag("zone", cmd.Flags().Lookup("zone"))

	return cmd
}
