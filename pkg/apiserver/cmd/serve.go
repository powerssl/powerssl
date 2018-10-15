package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/apiserver"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Run: func(cmd *cobra.Command, args []string) {
		grpcAddr := viper.GetString("grpc.addr")
		dbDialect := viper.GetString("db.dialect")
		dbConnection := viper.GetString("db.connection")

		apiserver.Run(grpcAddr, dbDialect, dbConnection)
	},
}

func init() {
	serveCmd.Flags().StringP("grpc-addr", "", ":8080", "GRPC Addr")
	serveCmd.Flags().StringP("db-dialect", "", "sqlite3", "DB Dialect")
	serveCmd.Flags().StringP("db-connection", "", "/tmp/powerssl.sqlie3", "DB connection")

	viper.BindPFlag("grpc.addr", serveCmd.Flags().Lookup("grpc-addr"))
	viper.BindPFlag("db.dialect", serveCmd.Flags().Lookup("db-dialect"))
	viper.BindPFlag("db.connection", serveCmd.Flags().Lookup("db-connection"))

	rootCmd.AddCommand(serveCmd)
}
