package main

import (
	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/grpcgateway/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
