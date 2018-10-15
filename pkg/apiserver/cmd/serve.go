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
		grpcCertFile := viper.GetString("grpc.cert")
		grpcKeyFile := viper.GetString("grpc.key")
		grpcInsecure := viper.GetBool("grpc.insecure")
		dbDialect := viper.GetString("db.dialect")
		dbConnection := viper.GetString("db.connection")

		apiserver.Run(grpcAddr, grpcCertFile, grpcKeyFile, grpcInsecure, dbDialect, dbConnection)
	},
}

func init() {
	serveCmd.Flags().StringP("grpc-addr", "", ":8080", "GRPC Addr")
	serveCmd.Flags().StringP("cert-file", "", "", "TLS certificate")
	serveCmd.Flags().StringP("key-file", "", "", "TLS key")
	serveCmd.Flags().BoolP("insecure", "", false, "Do not use TLS")
	serveCmd.Flags().StringP("db-dialect", "", "sqlite3", "DB Dialect")
	serveCmd.Flags().StringP("db-connection", "", "/tmp/powerssl.sqlie3", "DB connection")

	viper.BindPFlag("grpc.addr", serveCmd.Flags().Lookup("grpc-addr"))
	viper.BindPFlag("grpc.cert", serveCmd.Flags().Lookup("cert-file"))
	viper.BindPFlag("grpc.key", serveCmd.Flags().Lookup("key-file"))
	viper.BindPFlag("grpc.insecure", serveCmd.Flags().Lookup("insecure"))
	viper.BindPFlag("db.dialect", serveCmd.Flags().Lookup("db-dialect"))
	viper.BindPFlag("db.connection", serveCmd.Flags().Lookup("db-connection"))

	rootCmd.AddCommand(serveCmd)
}
