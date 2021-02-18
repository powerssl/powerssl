package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"gopkg.in/yaml.v3"
)

func GenDocs(newCmdRoot *cobra.Command) {
	var dir string

	cmd := &cobra.Command{
		Use:   "gendocs DIR",
		Short: "gendocs generates documentation for this component",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if dir == "" {
				return errors.New("directory needs to be set")
			}
			if _, err := os.Stat(dir); err != nil {
				if !os.IsNotExist(err) {
					return err
				}
				return errors.New("directory does not exist")
			}
			return nil
		},
		Run: HandleError(func(cmd *cobra.Command, args []string) (err error) {
			if err = doc.GenMarkdownTreeCustom(newCmdRoot, dir+"/"+newCmdRoot.Use, filePrepender, linkHandler); err != nil {
				return err
			}

			// NOTE: Strip timestamp from all generated docs.
			var files []string
			if files, err = filepath.Glob(dir + "/" + newCmdRoot.Use + "/*.md"); err != nil {
				return err
			}
			for _, file := range files {
				var input []byte
				if input, err = ioutil.ReadFile(file); err != nil {
					return err
				}
				lines := strings.Split(string(input), "\n")
				for i, line := range lines {
					if strings.Contains(line, "Find more information at:") {
						lines = append(lines[:i], lines[i+2:]...)
					}
				}
				output := strings.Join(lines[0:len(lines)-2], "\n")
				if err = ioutil.WriteFile(file, []byte(output), 0644); err != nil {
					return err
				}
			}
			return nil
		}),
	}

	cmd.Flags().StringVarP(&dir, "dir", "d", "", "output directory")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
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
		panic(err)
	}
	return fmt.Sprintf("---\n%s---\n", string(byt))
}

func linkHandler(name string) string {
	base := strings.TrimSuffix(name, path.Ext(name))
	s := strings.Split(base, "_")
	base = fmt.Sprintf("/%s", strings.Join(s, "/"))
	return base
}
