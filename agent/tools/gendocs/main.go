package main

import (
	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/agent/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
