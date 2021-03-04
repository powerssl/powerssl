package policy

import (
	"embed"
	"io/fs"
)

//go:embed *.hcl
var embedded embed.FS

// Open opens the named file for reading and returns it as an fs.File.
func Open(name string) (fs.File, error) {
	return embedded.Open(name)
}

// ReadDir reads and returns the entire named directory.
func ReadDir(name string) ([]fs.DirEntry, error) {
	return embedded.ReadDir(name)
}

// ReadFile reads and returns the content of the named file.
func ReadFile(name string) ([]byte, error) {
	return embedded.ReadFile(name)
}
