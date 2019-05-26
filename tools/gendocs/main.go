package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra/doc"
	"gopkg.in/yaml.v2"

	agent "powerssl.io/powerssl/internal/app/agent/cmd"
	apiserver "powerssl.io/powerssl/internal/app/apiserver/cmd"
	auth "powerssl.io/powerssl/internal/app/auth/cmd"
	controller "powerssl.io/powerssl/internal/app/controller/cmd"
	acmeintegration "powerssl.io/powerssl/internal/app/integrations/acme/cmd"
	cloudflareintegration "powerssl.io/powerssl/internal/app/integrations/cloudflare/cmd"
	powerctl "powerssl.io/powerssl/internal/app/powerctl/cmd"
	powerutil "powerssl.io/powerssl/internal/app/powerutil/cmd"
	signer "powerssl.io/powerssl/internal/app/signer/cmd"
	webapp "powerssl.io/powerssl/internal/app/webapp/cmd"
)

func main() {
	check(doc.GenMarkdownTreeCustom(agent.NewCmdRoot(), "docs/powerssl-agent", filePrepender, linkHandler))
	check(doc.GenMarkdownTreeCustom(apiserver.NewCmdRoot(), "docs/powerssl-apiserver", filePrepender, linkHandler))
	check(doc.GenMarkdownTreeCustom(auth.NewCmdRoot(), "docs/powerssl-auth", filePrepender, linkHandler))
	check(doc.GenMarkdownTreeCustom(controller.NewCmdRoot(), "docs/powerssl-controller", filePrepender, linkHandler))
	check(doc.GenMarkdownTreeCustom(acmeintegration.NewCmdRoot(), "docs/powerssl-integration-acme", filePrepender, linkHandler))
	check(doc.GenMarkdownTreeCustom(cloudflareintegration.NewCmdRoot(), "docs/powerssl-integration-cloudflare", filePrepender, linkHandler))
	check(doc.GenMarkdownTreeCustom(powerctl.NewCmdRoot(), "docs/powerctl", filePrepender, linkHandler))
	check(doc.GenMarkdownTreeCustom(powerutil.NewCmdRoot(), "docs/powerutil", filePrepender, linkHandler))
	check(doc.GenMarkdownTreeCustom(signer.NewCmdRoot(), "docs/powerssl-signer", filePrepender, linkHandler))
	check(doc.GenMarkdownTreeCustom(webapp.NewCmdRoot(), "docs/powerssl-webapp", filePrepender, linkHandler))

	// NOTE: Strip timestamp from all generated docs.
	files, err := filepath.Glob("docs/**/*.md")
	check(err)
	for _, file := range files {
		input, err := ioutil.ReadFile(file)
		check(err)
		lines := strings.Split(string(input), "\n")
		for i, line := range lines {
			if strings.Contains(line, "Find more information at:") {
				lines = append(lines[:i], lines[i+2:]...)
			}
		}
		output := strings.Join(lines[0:len(lines)-2], "\n")
		check(ioutil.WriteFile(file, []byte(output), 0644))
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
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
		if s[1] == "create" || s[1] == "ca" {
			meta["has_children"] = true
		}
	case 3:
		meta["parent"] = s[len(s)-2]
		meta["grand_parent"] = s[len(s)-3]
	}
	byt, _ := yaml.Marshal(meta)
	return fmt.Sprintf("---\n%s---\n", string(byt))
}

func linkHandler(name string) string {
	base := strings.TrimSuffix(name, path.Ext(name))
	s := strings.Split(base, "_")
	base = fmt.Sprintf("/%s", strings.Join(s, "/"))
	return base
}
