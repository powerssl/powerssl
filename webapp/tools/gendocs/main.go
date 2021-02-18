package main

import (
	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/webapp/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
