package main

import (
	"powerssl.dev/auth/internal/cmd"
	cmdutil "powerssl.dev/common/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
