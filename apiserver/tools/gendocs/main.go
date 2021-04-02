package main

import (
	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/apiserver/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot(), "migrate")
}
