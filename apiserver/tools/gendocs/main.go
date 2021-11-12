package main

import (
	"github.com/spangenberg/snakecharmer"
	"powerssl.dev/apiserver/internal/cmd"
)

func main() {
	snakecharmer.GenDocs(cmd.NewCmdRoot(), "migrate")
}
