package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	file     = "internal/asset/index.html"
	template = "<script>\nvar powerssl = {\n  apiURI: \"{{ .APIAddr }}\",\n  authURI: \"{{ .AuthURI }}\",\n  grpcWebURI: \"{{ .GRPCWebURI }}\"\n};\n</script>"
)

func main() {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	s := string(content)
	if strings.Contains(s, template) {
		return
	}
	i := strings.Index(s, "<script")
	if i == -1 {
		log.Fatal("<script is not present in index.html")
	}
	s = s[:i] + template + s[i:]
	stat, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile(file, []byte(s), stat.Mode()); err != nil {
		log.Fatal(err)
	}
}
