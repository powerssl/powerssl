package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra/doc"

	"powerssl.io/pkg/powerctl/cmd"
)

func main() {
	if err := doc.GenMarkdownTree(cmd.RootCmd(), "doc/powerctl"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
