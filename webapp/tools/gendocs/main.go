package main

import (
	"github.com/spangenberg/snakecharmer"

	"powerssl.dev/webapp/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot())
}
