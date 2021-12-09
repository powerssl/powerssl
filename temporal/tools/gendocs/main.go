package main

import (
	"github.com/spangenberg/snakecharmer"

	"powerssl.dev/temporal/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot())
}
