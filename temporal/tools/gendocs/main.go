package main

import (
	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/temporal/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
