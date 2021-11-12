package main

import (
	"github.com/spangenberg/snakecharmer"
	"powerssl.dev/tools/dev-runner/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot())
}
