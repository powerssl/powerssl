package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"powerssl.io/pkg/powerctl"
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func pr(resource interface{}) {
	byt, err := yaml.Marshal(resource)
	if err != nil {
		er(err)
	}
	fmt.Println(string(byt))
}

func newGRPCClient() *powerctl.GrpcClient {
	grpcAddr := viper.GetString("grpcAddr")
	return powerctl.NewGRPCClient(grpcAddr)
}
