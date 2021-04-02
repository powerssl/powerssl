package main

import (
	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/controller/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
