package main

import (
	"github.com/spangenberg/snakecharmer"

	"powerssl.dev/powerutil/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot(), "ca")
}
