// Code generated for package template by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../web/index.html
package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x93\x5f\x6b\xdc\x30\x0c\xc0\xdf\xfb\x29\x34\xbf\xf4\xa5\xb9\xd0\x87\xc1\xe8\x1c\xc3\xd1\xc1\x28\xf4\xe1\x68\x19\x7b\x56\x6c\xdd\xc5\xab\xcf\x32\xb6\xd2\x23\x94\x7e\xf7\xe1\xcb\xe5\xfa\x67\xcc\x10\x14\x49\x3f\xfd\x41\xb2\xf5\x17\xc7\x56\xa6\x44\x30\xc8\x3e\x98\x0b\x5d\x05\x04\x8c\xbb\x4e\x51\x54\xd5\x40\xe8\xcc\x05\x00\x80\xde\x93\x20\xd8\x01\x73\x21\xe9\xd4\x28\xdb\xe6\x9b\xfa\xe0\xe2\x28\x14\xa5\x53\x07\xef\x64\xe8\x1c\x3d\x7b\x4b\xcd\x51\xb9\x02\x1f\xbd\x78\x0c\x4d\xb1\x18\xa8\xbb\xbe\x82\x32\x64\x1f\x9f\x1a\xe1\x66\xeb\xa5\x8b\xac\x20\xe2\x9e\x3a\xf5\xec\xe9\x90\x38\xcb\x92\x5a\xbc\x04\x32\x1b\x3e\x50\x7e\x7c\xbc\x87\x75\x4a\xba\x9d\x6d\xb3\xbf\xd8\xec\x93\xcc\x4a\x3d\xcf\x98\x21\x55\xba\x94\x00\x1d\xbc\x9c\x1d\xf5\x60\xf2\x6b\xe7\xf2\x0d\x5c\xbe\xbc\xc0\x6a\xbd\xb9\xab\x1a\xbc\xbe\x5e\x5e\x7d\xc4\x46\x19\x7e\x3d\xdc\x2d\xd8\xac\xfd\x8b\xed\x72\xb2\xbf\xa9\x7f\x23\x7f\x3e\x6c\x6e\x67\x43\x85\xcf\xec\xeb\xf7\xb9\xd5\xf6\x7d\xaf\xa7\xc6\xa1\x64\xdb\xa9\x16\x4b\x21\x29\x6d\x3f\x46\x17\x68\xf5\xa7\x28\xf3\x46\xeb\x76\x5e\x82\xee\xd9\x4d\xe6\x42\xef\xd1\x47\xb0\x01\x4b\xe9\x54\x9d\x39\xfa\x48\x59\x41\xe6\x40\x9d\xaa\x4e\x05\x45\xa6\xaa\x38\x5f\x52\xc0\xe9\x06\x22\x47\x5a\x06\x3a\x5c\x2f\xc1\x7b\x69\xbe\xaa\x4f\xa3\x1d\xae\x4f\x58\x5a\xa8\x40\xe8\x94\xd1\x25\xe1\xb9\xea\x96\xa3\x34\x07\xf2\xbb\x41\x9a\x9e\x83\x53\x66\x0d\xc2\x4f\x14\xc1\x17\x40\xd8\x72\xde\x03\x6f\x01\x21\x61\x29\x07\xce\x6e\xa5\xdb\x1a\x6f\xe0\x07\x43\x64\x81\x32\x60\x26\x98\x78\xcc\xb0\xde\xdc\xcd\xb1\xab\xd3\x90\xd2\xd2\xc0\x22\x33\x2d\x75\x53\xa6\xa6\xd8\xcc\x21\x60\x1f\x48\x19\x6d\xd9\x91\xd1\xed\x49\xa4\xbc\x5c\x8a\xcf\x49\x8e\xff\xfd\x28\xc2\xef\x26\x97\x26\xe8\x25\xd6\xaf\x29\x64\x39\x3a\xcc\x93\x02\x87\x82\x8d\x0d\x3e\xf5\x8c\xd9\x35\x82\x79\x57\xaf\x7b\x2d\xa1\xcc\x2d\xa7\x49\xb7\x73\xa2\xff\x66\x0e\xbc\xe3\x51\xce\xb9\x1d\xc6\x1d\x65\x65\xee\x8f\xe6\x8f\xd1\xc7\x46\x75\x5b\x97\x56\xe5\x69\xc1\xed\xfc\x18\xff\x06\x00\x00\xff\xff\xb4\xb6\x9b\x35\x9d\x03\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"index.html": indexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
