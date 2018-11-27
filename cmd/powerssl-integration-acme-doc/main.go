package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra/doc"

	"powerssl.io/pkg/integrations/acme/cmd"
)

func main() {
	if err := doc.GenMarkdownTree(cmd.RootCmd(), "docs/powerssl-integration-acme"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
