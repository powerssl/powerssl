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
		dbConnection := viper.GetString("db.connection")
		dbDialect := viper.GetString("db.dialect")
		grpcAddr := viper.GetString("grpc.addr")
		insecure := viper.GetBool("insecure")
		tlsCertFile := viper.GetString("tls.cert-file")
		tlsPrivateKeyFile := viper.GetString("tls.private-key-file")

		apiserver.Run(grpcAddr, tlsCertFile, tlsPrivateKeyFile, insecure, dbDialect, dbConnection)
	},
}

func init() {
	serveCmd.Flags().BoolP("insecure", "", false, "Do not use TLS for the server")
	serveCmd.Flags().StringP("db-connection", "", "/tmp/powerssl.sqlie3", "DB connection")
	serveCmd.Flags().StringP("db-dialect", "", "sqlite3", "DB Dialect")
	serveCmd.Flags().StringP("grpc-addr", "", ":8080", "GRPC Addr")
	serveCmd.Flags().StringP("tls-cert-file", "", "", "File containing the default x509 Certificate for GRPC.")
	serveCmd.Flags().StringP("tls-private-key-file", "", "", "File containing the default x509 private key matching --tls-cert-file.")

	viper.BindPFlag("db.connection", serveCmd.Flags().Lookup("db-connection"))
	viper.BindPFlag("db.dialect", serveCmd.Flags().Lookup("db-dialect"))
	viper.BindPFlag("grpc.addr", serveCmd.Flags().Lookup("grpc-addr"))
	viper.BindPFlag("insecure", serveCmd.Flags().Lookup("insecure"))
	viper.BindPFlag("tls.cert-file", serveCmd.Flags().Lookup("tls-cert-file"))
	viper.BindPFlag("tls.private-key-file", serveCmd.Flags().Lookup("tls-private-key-file"))

	rootCmd.AddCommand(serveCmd)
}
