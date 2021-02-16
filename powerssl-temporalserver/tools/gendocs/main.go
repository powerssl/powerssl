package main

import (
	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/temporalserver/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
