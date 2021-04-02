package main

import (
	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/auth/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
