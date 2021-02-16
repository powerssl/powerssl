package main

import (
	"powerssl.dev/agent/internal/cmd"
	cmdutil "powerssl.dev/common/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
