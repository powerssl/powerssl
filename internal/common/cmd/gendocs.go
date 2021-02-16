package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"gopkg.in/yaml.v3"
)

func GenDocs(cmd *cobra.Command) {
	if len(os.Args) < 2 {
		fail(fmt.Errorf("docs directory has to be provided"))
	}
	dir := os.Args[1]

	if err := doc.GenMarkdownTreeCustom(cmd, dir+"/"+cmd.Use, filePrepender, linkHandler); err != nil {
		fail(err)
	}

	// NOTE: Strip timestamp from all generated docs.
	files, err := filepath.Glob(dir + "/" + cmd.Use + "/*.md")
	if err != nil {
		fail(err)
	}
	for _, file := range files {
		var input []byte
		if input, err = ioutil.ReadFile(file); err != nil {
			fail(err)
		}
		lines := strings.Split(string(input), "\n")
		for i, line := range lines {
			if strings.Contains(line, "Find more information at:") {
				lines = append(lines[:i], lines[i+2:]...)
			}
		}
		output := strings.Join(lines[0:len(lines)-2], "\n")
		if err = ioutil.WriteFile(file, []byte(output), 0644); err != nil {
			fail(err)
		}
	}
}

func fail(err error) {
	log.Fatalf("error: %v\n", err)
}

func filePrepender(filename string) string {
	name := filepath.Base(filename)
	base := strings.TrimSuffix(name, path.Ext(name))
	s := strings.Split(base, "_")

	meta := make(map[string]interface{})
	meta["has_toc"] = false
	meta["permalink"] = fmt.Sprintf("/%s", strings.Join(s, "/"))
	meta["layout"] = "default"
	meta["title"] = s[len(s)-1]
	switch len(s) {
	case 1:
		meta["has_children"] = true
	case 2:
		meta["parent"] = s[len(s)-2]
		switch s[1] {
		case "create", "ca", "migrate", "temporal":
			meta["has_children"] = true
		}
	case 3:
		meta["parent"] = s[len(s)-2]
		meta["grand_parent"] = s[len(s)-3]
	}
	var byt []byte
	var err error
	if byt, err = yaml.Marshal(meta); err != nil {
		fail(err)
	}
	return fmt.Sprintf("---\n%s---\n", string(byt))
}

func linkHandler(name string) string {
	base := strings.TrimSuffix(name, path.Ext(name))
	s := strings.Split(base, "_")
	base = fmt.Sprintf("/%s", strings.Join(s, "/"))
	return base
}
