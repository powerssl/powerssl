package main

import (
	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/powerctl/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot(), "create")
}
