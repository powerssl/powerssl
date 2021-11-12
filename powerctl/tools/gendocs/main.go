package main

import (
	"github.com/spangenberg/snakecharmer"

	"powerssl.dev/powerctl/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot(), "create")
}
