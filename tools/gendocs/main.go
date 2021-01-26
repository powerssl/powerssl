package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"gopkg.in/yaml.v2"

	agent "powerssl.dev/powerssl/internal/app/agent/cmd"
	apiserver "powerssl.dev/powerssl/internal/app/apiserver/cmd"
	auth "powerssl.dev/powerssl/internal/app/auth/cmd"
	controller "powerssl.dev/powerssl/internal/app/controller/cmd"
	grpcgateway "powerssl.dev/powerssl/internal/app/grpcgateway/cmd"
	acmeintegration "powerssl.dev/powerssl/internal/app/integrations/acme/cmd"
	cloudflareintegration "powerssl.dev/powerssl/internal/app/integrations/cloudflare/cmd"
	powerctl "powerssl.dev/powerssl/internal/app/powerctl/cmd"
	powerutil "powerssl.dev/powerssl/internal/app/powerutil/cmd"
	signer "powerssl.dev/powerssl/internal/app/signer/cmd"
	temporalserver "powerssl.dev/powerssl/internal/app/temporalserver/cmd"
	webapp "powerssl.dev/powerssl/internal/app/webapp/cmd"
)

func main() {
	for _, f := range []func() *cobra.Command{
		acmeintegration.NewCmdRoot,
		agent.NewCmdRoot,
		apiserver.NewCmdRoot,
		auth.NewCmdRoot,
		cloudflareintegration.NewCmdRoot,
		controller.NewCmdRoot,
		grpcgateway.NewCmdRoot,
		powerctl.NewCmdRoot,
		powerutil.NewCmdRoot,
		signer.NewCmdRoot,
		temporalserver.NewCmdRoot,
		webapp.NewCmdRoot,
	} {
		cmd := f()
		if err := doc.GenMarkdownTreeCustom(cmd, "docs/"+cmd.Use, filePrepender, linkHandler); err != nil {
			fail(err)
		}
	}

	// NOTE: Strip timestamp from all generated docs.
	files, err := filepath.Glob("docs/**/*.md")
	if err != nil {
		fail(err)
	}
	for _, file := range files {
		input, err := ioutil.ReadFile(file)
		if err != nil {
			fail(err)
		}
		lines := strings.Split(string(input), "\n")
		for i, line := range lines {
			if strings.Contains(line, "Find more information at:") {
				lines = append(lines[:i], lines[i+2:]...)
			}
		}
		output := strings.Join(lines[0:len(lines)-2], "\n")
		if err := ioutil.WriteFile(file, []byte(output), 0644); err != nil {
			fail(err)
		}
	}
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
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
