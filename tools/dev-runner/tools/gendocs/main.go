package main

import (
	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/tools/dev-runner/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
