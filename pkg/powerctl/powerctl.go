package powerctl

import (
	"errors"
	"os"

	"github.com/spf13/viper"

	apiserverclient "powerssl.io/pkg/apiserver/client"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/tracing"
)

func NewGRPCClient() (*apiserverclient.GRPCClient, error) {
	addr := viper.GetString("addr")
	authToken := viper.GetString("auth-token")
	certFile := viper.GetString("ca-file")
	insecure := viper.GetBool("insecure")
	insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
	serverNameOverride := viper.GetString("server-name-override")
	if addr == "" {
		return nil, errors.New("Provide addr")
	}
	if !insecure && !insecureSkipTLSVerify && certFile == "" {
		return nil, errors.New("Provide ca-file")
	}
	logger := util.NewLogger(os.Stdout)
	tracer, _, _ := tracing.NewNoopTracer("powerctl", logger)
	return apiserverclient.NewGRPCClient(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, authToken, logger, tracer)
}
