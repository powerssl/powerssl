package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/temporal/internal"
)

func newCmdMigrate() *cobra.Command {
	var host, password, plugin, port, temporalDatabase, user, visibilityDatabase string

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run temporal migrations",
		Args:  cobra.NoArgs,
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return internal.RunMigrate(host, password, plugin, port, temporalDatabase, user, visibilityDatabase)
		}),
	}

	cmd.Flags().StringVar(&host, "host", "", "DB host")
	cmd.Flags().StringVar(&password, "password", "", "DB Password")
	cmd.Flags().StringVar(&plugin, "plugin", "", "DB Plugin")
	cmd.Flags().StringVar(&port, "port", "", "DB Port")
	cmd.Flags().StringVar(&temporalDatabase, "temporal-database", "temporal", "Temporal DB")
	cmd.Flags().StringVar(&user, "user", "", "DB User")
	cmd.Flags().StringVar(&visibilityDatabase, "visibility-database", "temporal_visibility", "Visibility DB")

	return cmd
}
