package main

import (
	"github.com/spangenberg/snakecharmer"

	"powerssl.dev/worker/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot())
}
