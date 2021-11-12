package main

import (
	"github.com/spangenberg/snakecharmer"

	"powerssl.dev/auth/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot())
}
