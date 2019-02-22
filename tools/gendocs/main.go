package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra/doc"
	"gopkg.in/yaml.v2"

	agent "powerssl.io/internal/app/agent/cmd"
	apiserver "powerssl.io/internal/app/apiserver/cmd"
	auth "powerssl.io/internal/app/auth/cmd"
	controller "powerssl.io/internal/app/controller/cmd"
	acmeintegration "powerssl.io/internal/app/integrations/acme/cmd"
	cloudflareintegration "powerssl.io/internal/app/integrations/cloudflare/cmd"
	powerctl "powerssl.io/internal/app/powerctl/cmd"
	powerutil "powerssl.io/internal/app/powerutil/cmd"
	signer "powerssl.io/internal/app/signer/cmd"
	web "powerssl.io/internal/app/web/cmd"
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
	check(doc.GenMarkdownTreeCustom(web.NewCmdRoot(), "docs/powerssl-webapp", filePrepender, linkHandler))
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
	return name
}
