package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra/doc"
	"gopkg.in/yaml.v2"

	"powerssl.io/internal/app/auth/cmd"
)

func main() {
	if err := doc.GenMarkdownTreeCustom(cmd.NewCmdRoot(), "docs/powerssl-auth", filePrepender, linkHandler); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func filePrepender(filename string) string {
	name := filepath.Base(filename)
	base := strings.TrimSuffix(name, path.Ext(name))
	s := strings.Split(base, "_")

	meta := make(map[string]interface{})
	meta["layout"] = "default"
	meta["title"] = s[len(s)-1]
	switch len(s) {
	case 1:
		meta["has_children"] = true
	case 2:
		meta["parent"] = s[len(s)-2]
	}
	byt, _ := yaml.Marshal(meta)
	return fmt.Sprintf("---\n%s---\n", string(byt))
}

func linkHandler(name string) string {
	return name
}
