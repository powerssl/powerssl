package main

import (
	"github.com/spangenberg/snakecharmer"

	"powerssl.dev/controller/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot())
}
