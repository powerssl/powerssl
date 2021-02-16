package main

import (
	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/signer/internal/cmd"
)

func main() {
	cmdutil.GenDocs(cmd.NewCmdRoot())
}
