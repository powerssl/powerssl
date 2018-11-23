package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	"google.golang.org/grpc/status"

	apiserverclient "powerssl.io/pkg/apiserver/client"
	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/tracing"
)

var DisplayName string

func er(msg interface{}) {
	err, ok := msg.(error)
	if ok {
		status, ok := status.FromError(err)
		if ok {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("RPC error:\n  Code:    %s\n  Message: %s\n", status.Code(), status.Message()))
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		fmt.Fprintln(os.Stderr, msg)
	}
	os.Exit(1)
}

func pr(resource interface{}) {
	var (
		err error
		out []byte
	)
	switch Output {
	case "yaml":
		out, err = yaml.Marshal(resource)
	case "json":
		out, err = json.Marshal(resource)
	case "table":
		var ok bool
		var resources []*Resource
		resources, ok = resource.([]*Resource)
		if !ok {
			if r, ok := resource.(*Resource); ok {
				resources = append(resources, r)
			}
		}
		resourceMap := make(map[string][]*Resource)
		for _, resource := range resources {
			resourceMap[resource.Kind] = append(resourceMap[resource.Kind], resource)
		}
		for kind, resources := range resourceMap {
			fmt.Printf("%s:\n", kind)
			table := tablewriter.NewWriter(os.Stdout)
			for _, resource := range resources {
				header, columns, err := resource.ToTable()
				if err != nil {
					er(err)
				}
				table.SetHeader(header)
				table.Append(columns)
			}
			table.Render()
		}
	default:
		er("Unknown output format")
	}
	if err != nil {
		er(err)
	}
	fmt.Fprintln(os.Stdout, string(out))
}

func NewGRPCClient() (*apiserverclient.GRPCClient, error) {
	certFile := viper.GetString("ca-file")
	addr := viper.GetString("addr")
	insecure := viper.GetBool("insecure")
	insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
	authToken := viper.GetString("auth-token")
	serverNameOverride := viper.GetString("server-name-override")
	if addr == "" {
		return nil, errors.New("Provide addr")
	}
	if !insecure && !insecureSkipTLSVerify && certFile == "" {
		return nil, errors.New("Provide ca-file")
	}
	logger := util.NewLogger(os.Stdout)
	tracer, _, _ := tracing.NewNoopTracer("powerctl", logger)
	client, err := apiserverclient.NewGRPCClient(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, authToken, logger, tracer)
	if err != nil {
		return nil, err
	}
	return client, nil
}
