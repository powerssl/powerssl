package main

import (
	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/powerutil/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot(), "ca")
}
