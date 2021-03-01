package main

import (
	"powerssl.dev/apiserver/internal/cmd"
	cmdutil "powerssl.dev/common/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot(), "migrate")
}
