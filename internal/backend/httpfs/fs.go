package httpfs

import (
	"net/http"
	"os"
)

type FileSystem struct {
	fs http.FileSystem
}

func NewFileSystem(fs http.FileSystem) *FileSystem {
	return &FileSystem{fs: fs}
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
