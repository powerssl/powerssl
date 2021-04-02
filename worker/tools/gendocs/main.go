package main

import (
	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/worker/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
