package main

import (
	"github.com/spangenberg/snakecharmer"

	"powerssl.dev/agent/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot())
}
