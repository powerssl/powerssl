package powerctl

import (
	"context"
	"errors"
	"os"

	"github.com/spf13/viper"

	"powerssl.io/powerssl/internal/pkg/tracing"
	"powerssl.io/powerssl/internal/pkg/util"
	apiserverclient "powerssl.io/powerssl/pkg/apiserver/client"
)

func NewGRPCClient() (*apiserverclient.GRPCClient, error) {
	addr := viper.GetString("addr")
	authToken := viper.GetString("auth-token")
	caFle := viper.GetString("ca-file")
	insecure := viper.GetBool("insecure")
	insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
	serverNameOverride := viper.GetString("server-name-override")
	if addr == "" {
		return nil, errors.New("provide addr")
	}
	if !insecure && !insecureSkipTLSVerify && caFle == "" {
		return nil, errors.New("provide ca-file")
	}
	logger := util.NewLogger(os.Stdout)
	tracer, _, _ := tracing.NewNoopTracer("powerctl", logger)
	cfg := &util.ClientConfig{
		Addr:                  addr,
		CAFile:                caFle,
		Insecure:              insecure,
		InsecureSkipTLSVerify: insecureSkipTLSVerify,
		ServerNameOverride:    serverNameOverride,
	}
	return apiserverclient.NewGRPCClient(context.TODO(), cfg, authToken, logger, tracer)
}
