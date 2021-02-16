package http

import (
	"net/http"
	"os"
)

func NewFileSystem(fs http.FileSystem) *FileSystem {
	return &FileSystem{fs: fs}
}

type FileSystem struct {
	fs http.FileSystem
}

func (fs *FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, _ := f.Stat()
	if s.IsDir() {
		return nil, os.ErrNotExist
	}

	return f, nil
}
